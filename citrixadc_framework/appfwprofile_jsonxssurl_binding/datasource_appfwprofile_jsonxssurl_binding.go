package appfwprofile_jsonxssurl_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileJsonxssurlBindingDataSource)(nil)

func APpfwprofileJsonxssurlBindingDataSource() datasource.DataSource {
	return &AppfwprofileJsonxssurlBindingDataSource{}
}

type AppfwprofileJsonxssurlBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileJsonxssurlBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_jsonxssurl_binding"
}

func (d *AppfwprofileJsonxssurlBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileJsonxssurlBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileJsonxssurlBindingDataSourceSchema()
}

func (d *AppfwprofileJsonxssurlBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileJsonxssurlBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	asvalueexprjsonxss_Name := data.AsValueExprJsonXss
	asvaluetypejsonxss_Name := data.AsValueTypeJsonXss
	jsonxssurl_Name := data.Jsonxssurl
	keynamejsonxss_Name := data.KeynameJsonXss

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_jsonxssurl_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_jsonxssurl_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_jsonxssurl_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check as_value_expr_json_xss
		if val, ok := v["as_value_expr_json_xss"].(string); ok {
			if asvalueexprjsonxss_Name.IsNull() || val != asvalueexprjsonxss_Name.ValueString() {
				match = false
				continue
			}
		} else if !asvalueexprjsonxss_Name.IsNull() {
			match = false
			continue
		}

		// Check as_value_type_json_xss
		if val, ok := v["as_value_type_json_xss"].(string); ok {
			if asvaluetypejsonxss_Name.IsNull() || val != asvaluetypejsonxss_Name.ValueString() {
				match = false
				continue
			}
		} else if !asvaluetypejsonxss_Name.IsNull() {
			match = false
			continue
		}

		// Check jsonxssurl
		if val, ok := v["jsonxssurl"].(string); ok {
			if jsonxssurl_Name.IsNull() || val != jsonxssurl_Name.ValueString() {
				match = false
				continue
			}
		} else if !jsonxssurl_Name.IsNull() {
			match = false
			continue
		}

		// Check keyname_json_xss
		if val, ok := v["keyname_json_xss"].(string); ok {
			if keynamejsonxss_Name.IsNull() || val != keynamejsonxss_Name.ValueString() {
				match = false
				continue
			}
		} else if !keynamejsonxss_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_jsonxssurl_binding with as_value_expr_json_xss %s not found", asvalueexprjsonxss_Name))
		return
	}

	appfwprofile_jsonxssurl_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
