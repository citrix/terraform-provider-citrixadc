package snmpoption

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SnmpoptionResourceModel describes the resource data model.
type SnmpoptionResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Partitionnameintrap  types.String `tfsdk:"partitionnameintrap"`
	Severityinfointrap   types.String `tfsdk:"severityinfointrap"`
	Snmpset              types.String `tfsdk:"snmpset"`
	Snmptraplogging      types.String `tfsdk:"snmptraplogging"`
	Snmptraplogginglevel types.String `tfsdk:"snmptraplogginglevel"`
}

func (r *SnmpoptionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the snmpoption resource.",
			},
			"partitionnameintrap": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Send partition name as a varbind in traps. By default the partition names are not sent as a varbind.",
			},
			"severityinfointrap": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "By default, the severity level info of the trap is not mentioned in the trap message. Enable this option to send severity level of trap as one of the varbind in the trap message.",
			},
			"snmpset": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Accept SNMP SET requests sent to the Citrix ADC, and allow SNMP managers to write values to MIB objects that are configured for write access.",
			},
			"snmptraplogging": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Log any SNMP trap events (for SNMP alarms in which logging is enabled) even if no trap listeners are configured. With the default setting, SNMP trap events are logged if at least one trap listener is configured on the appliance.",
			},
			"snmptraplogginglevel": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("INFORMATIONAL"),
				Description: "Audit log level of SNMP trap logs. The default value is INFORMATIONAL.",
			},
		},
	}
}

func snmpoptionGetThePayloadFromtheConfig(ctx context.Context, data *SnmpoptionResourceModel) snmp.Snmpoption {
	tflog.Debug(ctx, "In snmpoptionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	snmpoption := snmp.Snmpoption{}
	if !data.Partitionnameintrap.IsNull() {
		snmpoption.Partitionnameintrap = data.Partitionnameintrap.ValueString()
	}
	if !data.Severityinfointrap.IsNull() {
		snmpoption.Severityinfointrap = data.Severityinfointrap.ValueString()
	}
	if !data.Snmpset.IsNull() {
		snmpoption.Snmpset = data.Snmpset.ValueString()
	}
	if !data.Snmptraplogging.IsNull() {
		snmpoption.Snmptraplogging = data.Snmptraplogging.ValueString()
	}
	if !data.Snmptraplogginglevel.IsNull() {
		snmpoption.Snmptraplogginglevel = data.Snmptraplogginglevel.ValueString()
	}

	return snmpoption
}

func snmpoptionSetAttrFromGet(ctx context.Context, data *SnmpoptionResourceModel, getResponseData map[string]interface{}) *SnmpoptionResourceModel {
	tflog.Debug(ctx, "In snmpoptionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["partitionnameintrap"]; ok && val != nil {
		data.Partitionnameintrap = types.StringValue(val.(string))
	} else {
		data.Partitionnameintrap = types.StringNull()
	}
	if val, ok := getResponseData["severityinfointrap"]; ok && val != nil {
		data.Severityinfointrap = types.StringValue(val.(string))
	} else {
		data.Severityinfointrap = types.StringNull()
	}
	if val, ok := getResponseData["snmpset"]; ok && val != nil {
		data.Snmpset = types.StringValue(val.(string))
	} else {
		data.Snmpset = types.StringNull()
	}
	if val, ok := getResponseData["snmptraplogging"]; ok && val != nil {
		data.Snmptraplogging = types.StringValue(val.(string))
	} else {
		data.Snmptraplogging = types.StringNull()
	}
	if val, ok := getResponseData["snmptraplogginglevel"]; ok && val != nil {
		data.Snmptraplogginglevel = types.StringValue(val.(string))
	} else {
		data.Snmptraplogginglevel = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("snmpoption-config")

	return data
}
