package appfwprofile_sqlinjection_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileSqlinjectionBindingDataSource)(nil)

func APpfwprofileSqlinjectionBindingDataSource() datasource.DataSource {
	return &AppfwprofileSqlinjectionBindingDataSource{}
}

type AppfwprofileSqlinjectionBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileSqlinjectionBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_sqlinjection_binding"
}

func (d *AppfwprofileSqlinjectionBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileSqlinjectionBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileSqlinjectionBindingDataSourceSchema()
}

func (d *AppfwprofileSqlinjectionBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileSqlinjectionBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	asscanlocationsql_Name := data.AsScanLocationSql
	asvalueexprsql_Name := data.AsValueExprSql
	asvaluetypesql_Name := data.AsValueTypeSql
	formactionurlsql_Name := data.FormactionurlSql
	sqlinjection_Name := data.Sqlinjection

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_sqlinjection_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_sqlinjection_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_sqlinjection_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check as_scan_location_sql
		if val, ok := v["as_scan_location_sql"].(string); ok {
			if asscanlocationsql_Name.IsNull() || val != asscanlocationsql_Name.ValueString() {
				match = false
				continue
			}
		} else if !asscanlocationsql_Name.IsNull() {
			match = false
			continue
		}

		// Check as_value_expr_sql
		if val, ok := v["as_value_expr_sql"].(string); ok {
			if asvalueexprsql_Name.IsNull() || val != asvalueexprsql_Name.ValueString() {
				match = false
				continue
			}
		} else if !asvalueexprsql_Name.IsNull() {
			match = false
			continue
		}

		// Check as_value_type_sql
		if val, ok := v["as_value_type_sql"].(string); ok {
			if asvaluetypesql_Name.IsNull() || val != asvaluetypesql_Name.ValueString() {
				match = false
				continue
			}
		} else if !asvaluetypesql_Name.IsNull() {
			match = false
			continue
		}

		// Check formactionurl_sql
		if val, ok := v["formactionurl_sql"].(string); ok {
			if formactionurlsql_Name.IsNull() || val != formactionurlsql_Name.ValueString() {
				match = false
				continue
			}
		} else if !formactionurlsql_Name.IsNull() {
			match = false
			continue
		}

		// Check sqlinjection
		if val, ok := v["sqlinjection"].(string); ok {
			if sqlinjection_Name.IsNull() || val != sqlinjection_Name.ValueString() {
				match = false
				continue
			}
		} else if !sqlinjection_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_sqlinjection_binding with as_scan_location_sql %s not found", asscanlocationsql_Name))
		return
	}

	appfwprofile_sqlinjection_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
