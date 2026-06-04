package appfwprofile_restvalidation_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileRestvalidationBindingDataSource)(nil)

func APpfwprofileRestvalidationBindingDataSource() datasource.DataSource {
	return &AppfwprofileRestvalidationBindingDataSource{}
}

type AppfwprofileRestvalidationBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileRestvalidationBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_restvalidation_binding"
}

func (d *AppfwprofileRestvalidationBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileRestvalidationBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileRestvalidationBindingDataSourceSchema()
}

func (d *AppfwprofileRestvalidationBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileRestvalidationBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	restvalidationaction_Name := data.RestValidationAction
	restvalidation_Name := data.Restvalidation

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_restvalidation_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_restvalidation_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_restvalidation_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check rest_validation_action
		if val, ok := v["rest_validation_action"].(string); ok {
			if restvalidationaction_Name.IsNull() || val != restvalidationaction_Name.ValueString() {
				match = false
				continue
			}
		} else if !restvalidationaction_Name.IsNull() {
			match = false
			continue
		}

		// Check restvalidation
		if val, ok := v["restvalidation"].(string); ok {
			if restvalidation_Name.IsNull() || val != restvalidation_Name.ValueString() {
				match = false
				continue
			}
		} else if !restvalidation_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_restvalidation_binding with rest_validation_action %s not found", restvalidationaction_Name))
		return
	}

	appfwprofile_restvalidation_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}