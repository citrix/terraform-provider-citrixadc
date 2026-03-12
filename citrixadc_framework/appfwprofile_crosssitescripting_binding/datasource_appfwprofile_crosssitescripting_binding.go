package appfwprofile_crosssitescripting_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileCrosssitescriptingBindingDataSource)(nil)

func APpfwprofileCrosssitescriptingBindingDataSource() datasource.DataSource {
	return &AppfwprofileCrosssitescriptingBindingDataSource{}
}

type AppfwprofileCrosssitescriptingBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileCrosssitescriptingBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_crosssitescripting_binding"
}

func (d *AppfwprofileCrosssitescriptingBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileCrosssitescriptingBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileCrosssitescriptingBindingDataSourceSchema()
}

func (d *AppfwprofileCrosssitescriptingBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileCrosssitescriptingBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	asscanlocationxss_Name := data.AsScanLocationXss
	asvalueexprxss_Name := data.AsValueExprXss
	asvaluetypexss_Name := data.AsValueTypeXss
	crosssitescripting_Name := data.Crosssitescripting
	formactionurlxss_Name := data.FormactionurlXss

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_crosssitescripting_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_crosssitescripting_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_crosssitescripting_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check as_scan_location_xss
		if val, ok := v["as_scan_location_xss"].(string); ok {
			if asscanlocationxss_Name.IsNull() || val != asscanlocationxss_Name.ValueString() {
				match = false
				continue
			}
		} else if !asscanlocationxss_Name.IsNull() {
			match = false
			continue
		}

		// Check as_value_expr_xss
		if val, ok := v["as_value_expr_xss"].(string); ok {
			if asvalueexprxss_Name.IsNull() || val != asvalueexprxss_Name.ValueString() {
				match = false
				continue
			}
		} else if !asvalueexprxss_Name.IsNull() {
			match = false
			continue
		}

		// Check as_value_type_xss
		if val, ok := v["as_value_type_xss"].(string); ok {
			if asvaluetypexss_Name.IsNull() || val != asvaluetypexss_Name.ValueString() {
				match = false
				continue
			}
		} else if !asvaluetypexss_Name.IsNull() {
			match = false
			continue
		}

		// Check crosssitescripting
		if val, ok := v["crosssitescripting"].(string); ok {
			if crosssitescripting_Name.IsNull() || val != crosssitescripting_Name.ValueString() {
				match = false
				continue
			}
		} else if !crosssitescripting_Name.IsNull() {
			match = false
			continue
		}

		// Check formactionurl_xss
		if val, ok := v["formactionurl_xss"].(string); ok {
			if formactionurlxss_Name.IsNull() || val != formactionurlxss_Name.ValueString() {
				match = false
				continue
			}
		} else if !formactionurlxss_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_crosssitescripting_binding with as_scan_location_xss %s not found", asscanlocationxss_Name))
		return
	}

	appfwprofile_crosssitescripting_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
