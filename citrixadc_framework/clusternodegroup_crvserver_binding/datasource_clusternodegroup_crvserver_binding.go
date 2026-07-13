package clusternodegroup_crvserver_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*ClusternodegroupCrvserverBindingDataSource)(nil)

func CLusternodegroupCrvserverBindingDataSource() datasource.DataSource {
	return &ClusternodegroupCrvserverBindingDataSource{}
}

type ClusternodegroupCrvserverBindingDataSource struct {
	client *service.NitroClient
}

func (d *ClusternodegroupCrvserverBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusternodegroup_crvserver_binding"
}

func (d *ClusternodegroupCrvserverBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *ClusternodegroupCrvserverBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ClusternodegroupCrvserverBindingDataSourceSchema()
}

func (d *ClusternodegroupCrvserverBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ClusternodegroupCrvserverBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	vserver_Name := data.Vserver

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Clusternodegroup_crvserver_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read clusternodegroup_crvserver_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "clusternodegroup_crvserver_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check vserver
		if val, ok := v["vserver"].(string); ok {
			if vserver_Name.IsNull() || val != vserver_Name.ValueString() {
				match = false
				continue
			}
		} else if !vserver_Name.IsNull() {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("clusternodegroup_crvserver_binding with vserver %s not found", vserver_Name))
		return
	}

	clusternodegroup_crvserver_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
