package mapbmr_bmrv4network_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*MapbmrBmrv4networkBindingDataSource)(nil)

func MApbmrBmrv4networkBindingDataSource() datasource.DataSource {
	return &MapbmrBmrv4networkBindingDataSource{}
}

type MapbmrBmrv4networkBindingDataSource struct {
	client *service.NitroClient
}

func (d *MapbmrBmrv4networkBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mapbmr_bmrv4network_binding"
}

func (d *MapbmrBmrv4networkBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *MapbmrBmrv4networkBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = MapbmrBmrv4networkBindingDataSourceSchema()
}

func (d *MapbmrBmrv4networkBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data MapbmrBmrv4networkBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	network_Name := data.Network

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Mapbmr_bmrv4network_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read mapbmr_bmrv4network_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "mapbmr_bmrv4network_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check network
		if val, ok := v["network"].(string); ok {
			if network_Name.IsNull() || val != network_Name.ValueString() {
				match = false
				continue
			}
		} else if !network_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("mapbmr_bmrv4network_binding with network %s not found", network_Name.ValueString()))
		return
	}

	mapbmr_bmrv4network_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
