package hanode_routemonitor6_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*HanodeRoutemonitor6BindingDataSource)(nil)

func HAnodeRoutemonitor6BindingDataSource() datasource.DataSource {
	return &HanodeRoutemonitor6BindingDataSource{}
}

type HanodeRoutemonitor6BindingDataSource struct {
	client *service.NitroClient
}

func (d *HanodeRoutemonitor6BindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hanode_routemonitor6_binding"
}

func (d *HanodeRoutemonitor6BindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *HanodeRoutemonitor6BindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = HanodeRoutemonitor6BindingDataSourceSchema()
}

func (d *HanodeRoutemonitor6BindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data HanodeRoutemonitor6BindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	id_Name := fmt.Sprintf("%d", data.Hanodeid.ValueInt64())
	routemonitor_Name := data.Routemonitor

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Hanode_routemonitor6_binding.Type(),
		ResourceName:             id_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read hanode_routemonitor6_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "hanode_routemonitor6_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("hanode_routemonitor6_binding with routemonitor %s not found", routemonitor_Name))
		return
	}

	hanode_routemonitor6_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
