package snmptrap_snmpuser_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SnmptrapSnmpuserBindingDataSource)(nil)

func SNmptrapSnmpuserBindingDataSource() datasource.DataSource {
	return &SnmptrapSnmpuserBindingDataSource{}
}

type SnmptrapSnmpuserBindingDataSource struct {
	client *service.NitroClient
}

func (d *SnmptrapSnmpuserBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmptrap_snmpuser_binding"
}

func (d *SnmptrapSnmpuserBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SnmptrapSnmpuserBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SnmptrapSnmpuserBindingDataSourceSchema()
}

func (d *SnmptrapSnmpuserBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SnmptrapSnmpuserBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID
	td_Name := data.Td
	trapclass_Name := data.Trapclass
	trapdestination_Name := data.Trapdestination
	username_Name := data.Username
	version_Name := data.Version

	var dataArr []map[string]interface{}
	var err error

	// Build ArgsMap with required parameters for binding
	argsMap := make(map[string]string)
	if !trapclass_Name.IsNull() {
		argsMap["trapclass"] = trapclass_Name.ValueString()
	}
	if !trapdestination_Name.IsNull() {
		argsMap["trapdestination"] = trapdestination_Name.ValueString()
	}
	if !version_Name.IsNull() {
		argsMap["version"] = version_Name.ValueString()
	}
	if !td_Name.IsNull() {
		argsMap["td"] = fmt.Sprintf("%d", td_Name.ValueInt64())
	}

	findParams := service.FindParams{
		ResourceType:             service.Snmptrap_snmpuser_binding.Type(),
		ArgsMap:                  argsMap,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read snmptrap_snmpuser_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "snmptrap_snmpuser_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check td
		if val, ok := v["td"]; ok {
			val, _ = utils.ConvertToInt64(val)
			if td_Name.IsNull() || val != td_Name.ValueInt64() {
				match = false
				continue
			}
		} else if !td_Name.IsNull() {
			match = false
			continue
		}

		// Check trapclass
		if val, ok := v["trapclass"].(string); ok {
			if trapclass_Name.IsNull() || val != trapclass_Name.ValueString() {
				match = false
				continue
			}
		} else if !trapclass_Name.IsNull() {
			match = false
			continue
		}

		// Check trapdestination
		if val, ok := v["trapdestination"].(string); ok {
			if trapdestination_Name.IsNull() || val != trapdestination_Name.ValueString() {
				match = false
				continue
			}
		} else if !trapdestination_Name.IsNull() {
			match = false
			continue
		}

		// Check username
		if val, ok := v["username"].(string); ok {
			if username_Name.IsNull() || val != username_Name.ValueString() {
				match = false
				continue
			}
		} else if !username_Name.IsNull() {
			match = false
			continue
		}

		// Check version
		if val, ok := v["version"].(string); ok {
			if version_Name.IsNull() || val != version_Name.ValueString() {
				match = false
				continue
			}
		} else if !version_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("snmptrap_snmpuser_binding with td %s not found", td_Name))
		return
	}

	snmptrap_snmpuser_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
