package appfwprofile_bypasslist_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileBypasslistBindingDataSource)(nil)

func APpfwprofileBypasslistBindingDataSource() datasource.DataSource {
	return &AppfwprofileBypasslistBindingDataSource{}
}

type AppfwprofileBypasslistBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileBypasslistBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_bypasslist_binding"
}

func (d *AppfwprofileBypasslistBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileBypasslistBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileBypasslistBindingDataSourceSchema()
}

func (d *AppfwprofileBypasslistBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileBypasslistBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	asbypasslist_Name := data.AsBypassList
	asbypasslistlocation_Name := data.AsBypassListLocation
	asbypasslistvaluetype_Name := data.AsBypassListValueType

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_bypasslist_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_bypasslist_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_bypasslist_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check as_bypass_list
		if val, ok := v["as_bypass_list"].(string); ok {
			if asbypasslist_Name.IsNull() || val != asbypasslist_Name.ValueString() {
				match = false
				continue
			}
		} else if !asbypasslist_Name.IsNull() {
			match = false
			continue
		}

		// Check as_bypass_list_location
		if val, ok := v["as_bypass_list_location"].(string); ok {
			if asbypasslistlocation_Name.IsNull() || val != asbypasslistlocation_Name.ValueString() {
				match = false
				continue
			}
		} else if !asbypasslistlocation_Name.IsNull() {
			match = false
			continue
		}

		// Check as_bypass_list_value_type
		if val, ok := v["as_bypass_list_value_type"].(string); ok {
			if asbypasslistvaluetype_Name.IsNull() || val != asbypasslistvaluetype_Name.ValueString() {
				match = false
				continue
			}
		} else if !asbypasslistvaluetype_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_bypasslist_binding with as_bypass_list %s not found", asbypasslist_Name))
		return
	}

	appfwprofile_bypasslist_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
