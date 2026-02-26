package hanode_routemonitor_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*HanodeRoutemonitorBindingDataSource)(nil)

func HAnodeRoutemonitorBindingDataSource() datasource.DataSource {
	return &HanodeRoutemonitorBindingDataSource{}
}

type HanodeRoutemonitorBindingDataSource struct {
	client *service.NitroClient
}

func (d *HanodeRoutemonitorBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hanode_routemonitor_binding"
}

func (d *HanodeRoutemonitorBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *HanodeRoutemonitorBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = HanodeRoutemonitorBindingDataSourceSchema()
}

func (d *HanodeRoutemonitorBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data HanodeRoutemonitorBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	id_Name := fmt.Sprintf("%d", data.Hanodeid.ValueInt64())
	netmask_Name := data.Netmask
	routemonitor_Name := data.Routemonitor

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Hanode_routemonitor_binding.Type(),
		ResourceName:             id_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read hanode_routemonitor_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "hanode_routemonitor_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check netmask
		if val, ok := v["netmask"].(string); ok {
			if netmask_Name.IsNull() || val != netmask_Name.ValueString() {
				match = false
				continue
			}
		} else if !netmask_Name.IsNull() {
			match = false
			continue
		}

		// Check routemonitor
		if val, ok := v["routemonitor"].(string); ok {
			if routemonitor_Name.IsNull() || val != routemonitor_Name.ValueString() {
				match = false
				continue
			}
		} else if !routemonitor_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("hanode_routemonitor_binding with netmask %s not found", netmask_Name))
		return
	}

	hanode_routemonitor_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
