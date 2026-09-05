package policypatset_pattern_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*PolicypatsetPatternBindingDataSource)(nil)

func POlicypatsetPatternBindingDataSource() datasource.DataSource {
	return &PolicypatsetPatternBindingDataSource{}
}

type PolicypatsetPatternBindingDataSource struct {
	client *service.NitroClient
}

func (d *PolicypatsetPatternBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policypatset_pattern_binding"
}

func (d *PolicypatsetPatternBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *PolicypatsetPatternBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = PolicypatsetPatternBindingDataSourceSchema()
}

func (d *PolicypatsetPatternBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data PolicypatsetPatternBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Array filter with parent ID - look up the patset by name, filter array by String.
	name_Name := data.Name.ValueString()
	stringText := data.String

	findParams := service.FindParams{
		ResourceType:             service.Policypatset_pattern_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 2823,
	}
	dataArr, err := d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read policypatset_pattern_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "policypatset_pattern_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right String
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["String"].(string); ok {
			if stringText.IsNull() || val == stringText.ValueString() {
				foundIndex = i
				break
			}
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("policypatset_pattern_binding with string %s not found", stringText))
		return
	}

	policypatset_pattern_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
