package policydataset_value_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*PolicydatasetValueBindingDataSource)(nil)

func POlicydatasetValueBindingDataSource() datasource.DataSource {
	return &PolicydatasetValueBindingDataSource{}
}

type PolicydatasetValueBindingDataSource struct {
	client *service.NitroClient
}

func (d *PolicydatasetValueBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policydataset_value_binding"
}

func (d *PolicydatasetValueBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *PolicydatasetValueBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = PolicydatasetValueBindingDataSourceSchema()
}

func (d *PolicydatasetValueBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data PolicydatasetValueBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	endrange_Name := data.Endrange
	value_Name := data.Value

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Policydataset_value_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read policydataset_value_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "policydataset_value_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check endrange
		if val, ok := v["endrange"].(string); ok {
			if endrange_Name.IsNull() || val != endrange_Name.ValueString() {
				match = false
				continue
			}
		} else if !endrange_Name.IsNull() {
			match = false
			continue
		}

		// Check value
		if val, ok := v["value"].(string); ok {
			if value_Name.IsNull() || val != value_Name.ValueString() {
				match = false
				continue
			}
		} else if !value_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("policydataset_value_binding with endrange %s not found", endrange_Name))
		return
	}

	policydataset_value_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
