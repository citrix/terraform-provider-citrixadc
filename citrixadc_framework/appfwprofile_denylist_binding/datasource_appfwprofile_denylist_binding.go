package appfwprofile_denylist_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileDenylistBindingDataSource)(nil)

func APpfwprofileDenylistBindingDataSource() datasource.DataSource {
	return &AppfwprofileDenylistBindingDataSource{}
}

type AppfwprofileDenylistBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileDenylistBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_denylist_binding"
}

func (d *AppfwprofileDenylistBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileDenylistBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileDenylistBindingDataSourceSchema()
}

func (d *AppfwprofileDenylistBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileDenylistBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	asdenylist_Name := data.AsDenyList
	asdenylistlocation_Name := data.AsDenyListLocation
	asdenylistvaluetype_Name := data.AsDenyListValueType

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_denylist_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_denylist_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_denylist_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check as_deny_list
		if val, ok := v["as_deny_list"].(string); ok {
			if asdenylist_Name.IsNull() || val != asdenylist_Name.ValueString() {
				match = false
				continue
			}
		} else if !asdenylist_Name.IsNull() {
			match = false
			continue
		}

		// Check as_deny_list_location
		if val, ok := v["as_deny_list_location"].(string); ok {
			if asdenylistlocation_Name.IsNull() || val != asdenylistlocation_Name.ValueString() {
				match = false
				continue
			}
		} else if !asdenylistlocation_Name.IsNull() {
			match = false
			continue
		}

		// Check as_deny_list_value_type
		if val, ok := v["as_deny_list_value_type"].(string); ok {
			if asdenylistvaluetype_Name.IsNull() || val != asdenylistvaluetype_Name.ValueString() {
				match = false
				continue
			}
		} else if !asdenylistvaluetype_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_denylist_binding with as_deny_list %s not found", asdenylist_Name))
		return
	}

	appfwprofile_denylist_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
