package apiprofile_apispec_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*ApiprofileApispecBindingDataSource)(nil)

func APiprofileApispecBindingDataSource() datasource.DataSource {
	return &ApiprofileApispecBindingDataSource{}
}

type ApiprofileApispecBindingDataSource struct {
	client *service.NitroClient
}

func (d *ApiprofileApispecBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apiprofile_apispec_binding"
}

func (d *ApiprofileApispecBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *ApiprofileApispecBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ApiprofileApispecBindingDataSourceSchema()
}

func (d *ApiprofileApispecBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ApiprofileApispecBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	apispec_Name := data.Apispec

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Apiprofile_apispec_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read apiprofile_apispec_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "apiprofile_apispec_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check apispec
		if val, ok := v["apispec"].(string); ok {
			if apispec_Name.IsNull() || val != apispec_Name.ValueString() {
				match = false
				continue
			}
		} else if !apispec_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("apiprofile_apispec_binding with apispec %s not found", apispec_Name))
		return
	}

	apiprofile_apispec_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}