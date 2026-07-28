package linkset_interface_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*LinksetInterfaceBindingDataSource)(nil)

func LInksetInterfaceBindingDataSource() datasource.DataSource {
	return &LinksetInterfaceBindingDataSource{}
}

type LinksetInterfaceBindingDataSource struct {
	client *service.NitroClient
}

func (d *LinksetInterfaceBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_linkset_interface_binding"
}

func (d *LinksetInterfaceBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LinksetInterfaceBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LinksetInterfaceBindingDataSourceSchema()
}

func (d *LinksetInterfaceBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LinksetInterfaceBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	linksetid := data.Linksetid.ValueString()
	wantIfnum := data.Ifnum

	// The direct linkset_interface_binding endpoint returns a keyless empty body on
	// this firmware; read the bound interfaces from the aggregate parent endpoint.
	dataArr, err := linkset_interface_bindingAggregateRead(d.client, linksetid)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read linkset_interface_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "linkset_interface_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right ifnum.
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["ifnum"].(string); ok {
			if wantIfnum.IsNull() || val == wantIfnum.ValueString() {
				foundIndex = i
				break
			}
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("linkset_interface_binding with ifnum %s not found", wantIfnum))
		return
	}

	linkset_interface_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
