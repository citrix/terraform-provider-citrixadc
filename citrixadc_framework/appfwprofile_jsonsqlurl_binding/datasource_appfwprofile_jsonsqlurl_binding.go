package appfwprofile_jsonsqlurl_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileJsonsqlurlBindingDataSource)(nil)

func APpfwprofileJsonsqlurlBindingDataSource() datasource.DataSource {
	return &AppfwprofileJsonsqlurlBindingDataSource{}
}

type AppfwprofileJsonsqlurlBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileJsonsqlurlBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_jsonsqlurl_binding"
}

func (d *AppfwprofileJsonsqlurlBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileJsonsqlurlBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileJsonsqlurlBindingDataSourceSchema()
}

func (d *AppfwprofileJsonsqlurlBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileJsonsqlurlBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	asvalueexprjsonsql_Name := data.AsValueExprJsonSql
	asvaluetypejsonsql_Name := data.AsValueTypeJsonSql
	jsonsqlurl_Name := data.Jsonsqlurl
	keynamejsonsql_Name := data.KeynameJsonSql

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_jsonsqlurl_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_jsonsqlurl_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_jsonsqlurl_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check as_value_expr_json_sql
		if val, ok := v["as_value_expr_json_sql"].(string); ok {
			if asvalueexprjsonsql_Name.IsNull() || val != asvalueexprjsonsql_Name.ValueString() {
				match = false
				continue
			}
		} else if !asvalueexprjsonsql_Name.IsNull() {
			match = false
			continue
		}

		// Check as_value_type_json_sql
		if val, ok := v["as_value_type_json_sql"].(string); ok {
			if asvaluetypejsonsql_Name.IsNull() || val != asvaluetypejsonsql_Name.ValueString() {
				match = false
				continue
			}
		} else if !asvaluetypejsonsql_Name.IsNull() {
			match = false
			continue
		}

		// Check jsonsqlurl
		if val, ok := v["jsonsqlurl"].(string); ok {
			if jsonsqlurl_Name.IsNull() || val != jsonsqlurl_Name.ValueString() {
				match = false
				continue
			}
		} else if !jsonsqlurl_Name.IsNull() {
			match = false
			continue
		}

		// Check keyname_json_sql
		if val, ok := v["keyname_json_sql"].(string); ok {
			if keynamejsonsql_Name.IsNull() || val != keynamejsonsql_Name.ValueString() {
				match = false
				continue
			}
		} else if !keynamejsonsql_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_jsonsqlurl_binding with as_value_expr_json_sql %s not found", asvalueexprjsonsql_Name))
		return
	}

	appfwprofile_jsonsqlurl_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
