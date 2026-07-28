package channel_interface_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*ChannelInterfaceBindingDataSource)(nil)

func CHannelInterfaceBindingDataSource() datasource.DataSource {
	return &ChannelInterfaceBindingDataSource{}
}

type ChannelInterfaceBindingDataSource struct {
	client *service.NitroClient
}

func (d *ChannelInterfaceBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_channel_interface_binding"
}

func (d *ChannelInterfaceBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *ChannelInterfaceBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ChannelInterfaceBindingDataSourceSchema()
}

func (d *ChannelInterfaceBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ChannelInterfaceBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	channelid := data.Channelid.ValueString()
	wantIfnum := datasourceFirstIfnum(ctx, &data)

	// The direct channel_interface_binding endpoint returns a keyless empty body on
	// this firmware; read the bound interfaces from the aggregate parent endpoint.
	dataArr, err := channel_interface_bindingAggregateRead(d.client, channelid)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read channel_interface_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "channel_interface_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right ifnum.
	// The aggregate response represents "ifnum" as a JSON array (e.g. ["1/2"]);
	// match against any member (tolerating the scalar form defensively).
	foundIndex := -1
	for i, v := range dataArr {
		if channelInterfaceRowHasIfnum(v["ifnum"], wantIfnum) {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("channel_interface_binding with ifnum %s not found", wantIfnum))
		return
	}

	channel_interface_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
