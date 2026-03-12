package snmptrap_snmpuser_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SnmptrapSnmpuserBindingResourceModel describes the resource data model.
type SnmptrapSnmpuserBindingResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Securitylevel   types.String `tfsdk:"securitylevel"`
	Td              types.Int64  `tfsdk:"td"`
	Trapclass       types.String `tfsdk:"trapclass"`
	Trapdestination types.String `tfsdk:"trapdestination"`
	Username        types.String `tfsdk:"username"`
	Version         types.String `tfsdk:"version"`
}

func (r *SnmptrapSnmpuserBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the snmptrap_snmpuser_binding resource.",
			},
			"securitylevel": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("authNoPriv"),
				Description: "Security level of the SNMPv3 trap.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"trapclass": schema.StringAttribute{
				Required:    true,
				Description: "Type of trap messages that the Citrix ADC sends to the trap listener: Generic or the enterprise-specific messages defined in the MIB file.",
			},
			"trapdestination": schema.StringAttribute{
				Required:    true,
				Description: "IPv4 or the IPv6 address of the trap listener to which the Citrix ADC is to send SNMP trap messages.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SNMP user that will send the SNMPv3 traps.",
			},
			"version": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("V3"),
				Description: "SNMP version, which determines the format of trap messages sent to the trap listener. \nThis setting must match the setting on the trap listener. Otherwise, the listener drops the trap messages.",
			},
		},
	}
}

func snmptrap_snmpuser_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SnmptrapSnmpuserBindingResourceModel) snmp.Snmptrapsnmpuserbinding {
	tflog.Debug(ctx, "In snmptrap_snmpuser_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	snmptrap_snmpuser_binding := snmp.Snmptrapsnmpuserbinding{}
	if !data.Securitylevel.IsNull() {
		snmptrap_snmpuser_binding.Securitylevel = data.Securitylevel.ValueString()
	}
	if !data.Td.IsNull() {
		snmptrap_snmpuser_binding.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Trapclass.IsNull() {
		snmptrap_snmpuser_binding.Trapclass = data.Trapclass.ValueString()
	}
	if !data.Trapdestination.IsNull() {
		snmptrap_snmpuser_binding.Trapdestination = data.Trapdestination.ValueString()
	}
	if !data.Username.IsNull() {
		snmptrap_snmpuser_binding.Username = data.Username.ValueString()
	}
	if !data.Version.IsNull() {
		snmptrap_snmpuser_binding.Version = data.Version.ValueString()
	}

	return snmptrap_snmpuser_binding
}

func snmptrap_snmpuser_bindingSetAttrFromGet(ctx context.Context, data *SnmptrapSnmpuserBindingResourceModel, getResponseData map[string]interface{}) *SnmptrapSnmpuserBindingResourceModel {
	tflog.Debug(ctx, "In snmptrap_snmpuser_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["securitylevel"]; ok && val != nil {
		data.Securitylevel = types.StringValue(val.(string))
	} else {
		data.Securitylevel = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["trapclass"]; ok && val != nil {
		data.Trapclass = types.StringValue(val.(string))
	} else {
		data.Trapclass = types.StringNull()
	}
	if val, ok := getResponseData["trapdestination"]; ok && val != nil {
		data.Trapdestination = types.StringValue(val.(string))
	} else {
		data.Trapdestination = types.StringNull()
	}
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	} else {
		data.Username = types.StringNull()
	}
	if val, ok := getResponseData["version"]; ok && val != nil {
		data.Version = types.StringValue(val.(string))
	} else {
		data.Version = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("td:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Td.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("trapclass:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Trapclass.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("trapdestination:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Trapdestination.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("username:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Username.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("version:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Version.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
