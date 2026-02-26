package appfwprofile_jsoncmdurl_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileJsoncmdurlBindingDataSource)(nil)

func APpfwprofileJsoncmdurlBindingDataSource() datasource.DataSource {
	return &AppfwprofileJsoncmdurlBindingDataSource{}
}

type AppfwprofileJsoncmdurlBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileJsoncmdurlBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_jsoncmdurl_binding"
}

func (d *AppfwprofileJsoncmdurlBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileJsoncmdurlBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileJsoncmdurlBindingDataSourceSchema()
}

func (d *AppfwprofileJsoncmdurlBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileJsoncmdurlBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	asvalueexprjsoncmd_Name := data.AsValueExprJsonCmd
	asvaluetypejsoncmd_Name := data.AsValueTypeJsonCmd
	jsoncmdurl_Name := data.Jsoncmdurl
	keynamejsoncmd_Name := data.KeynameJsonCmd

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_jsoncmdurl_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_jsoncmdurl_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_jsoncmdurl_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check as_value_expr_json_cmd
		if val, ok := v["as_value_expr_json_cmd"].(string); ok {
			if asvalueexprjsoncmd_Name.IsNull() || val != asvalueexprjsoncmd_Name.ValueString() {
				match = false
				continue
			}
		} else if !asvalueexprjsoncmd_Name.IsNull() {
			match = false
			continue
		}

		// Check as_value_type_json_cmd
		if val, ok := v["as_value_type_json_cmd"].(string); ok {
			if asvaluetypejsoncmd_Name.IsNull() || val != asvaluetypejsoncmd_Name.ValueString() {
				match = false
				continue
			}
		} else if !asvaluetypejsoncmd_Name.IsNull() {
			match = false
			continue
		}

		// Check jsoncmdurl
		if val, ok := v["jsoncmdurl"].(string); ok {
			if jsoncmdurl_Name.IsNull() || val != jsoncmdurl_Name.ValueString() {
				match = false
				continue
			}
		} else if !jsoncmdurl_Name.IsNull() {
			match = false
			continue
		}

		// Check keyname_json_cmd
		if val, ok := v["keyname_json_cmd"].(string); ok {
			if keynamejsoncmd_Name.IsNull() || val != keynamejsoncmd_Name.ValueString() {
				match = false
				continue
			}
		} else if !keynamejsoncmd_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_jsoncmdurl_binding with as_value_expr_json_cmd %s not found", asvalueexprjsoncmd_Name))
		return
	}

	appfwprofile_jsoncmdurl_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
