package appfwprofile_cmdinjection_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwprofileCmdinjectionBindingDataSource)(nil)

func APpfwprofileCmdinjectionBindingDataSource() datasource.DataSource {
	return &AppfwprofileCmdinjectionBindingDataSource{}
}

type AppfwprofileCmdinjectionBindingDataSource struct {
	client *service.NitroClient
}

func (d *AppfwprofileCmdinjectionBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_cmdinjection_binding"
}

func (d *AppfwprofileCmdinjectionBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwprofileCmdinjectionBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwprofileCmdinjectionBindingDataSourceSchema()
}

func (d *AppfwprofileCmdinjectionBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwprofileCmdinjectionBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	asscanlocationcmd_Name := data.AsScanLocationCmd
	asvalueexprcmd_Name := data.AsValueExprCmd
	asvaluetypecmd_Name := data.AsValueTypeCmd
	cmdinjection_Name := data.Cmdinjection
	formactionurlcmd_Name := data.FormactionurlCmd

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Appfwprofile_cmdinjection_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_cmdinjection_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_cmdinjection_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check as_scan_location_cmd
		if val, ok := v["as_scan_location_cmd"].(string); ok {
			if asscanlocationcmd_Name.IsNull() || val != asscanlocationcmd_Name.ValueString() {
				match = false
				continue
			}
		} else if !asscanlocationcmd_Name.IsNull() {
			match = false
			continue
		}

		// Check as_value_expr_cmd
		if val, ok := v["as_value_expr_cmd"].(string); ok {
			if asvalueexprcmd_Name.IsNull() || val != asvalueexprcmd_Name.ValueString() {
				match = false
				continue
			}
		} else if !asvalueexprcmd_Name.IsNull() {
			match = false
			continue
		}

		// Check as_value_type_cmd
		if val, ok := v["as_value_type_cmd"].(string); ok {
			if asvaluetypecmd_Name.IsNull() || val != asvaluetypecmd_Name.ValueString() {
				match = false
				continue
			}
		} else if !asvaluetypecmd_Name.IsNull() {
			match = false
			continue
		}

		// Check cmdinjection
		if val, ok := v["cmdinjection"].(string); ok {
			if cmdinjection_Name.IsNull() || val != cmdinjection_Name.ValueString() {
				match = false
				continue
			}
		} else if !cmdinjection_Name.IsNull() {
			match = false
			continue
		}

		// Check formactionurl_cmd
		if val, ok := v["formactionurl_cmd"].(string); ok {
			if formactionurlcmd_Name.IsNull() || val != formactionurlcmd_Name.ValueString() {
				match = false
				continue
			}
		} else if !formactionurlcmd_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("appfwprofile_cmdinjection_binding with as_scan_location_cmd %s not found", asscanlocationcmd_Name))
		return
	}

	appfwprofile_cmdinjection_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
