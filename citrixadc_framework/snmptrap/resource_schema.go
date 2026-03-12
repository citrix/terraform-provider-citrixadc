package snmptrap

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SnmptrapResourceModel describes the resource data model.
type SnmptrapResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Allpartitions   types.String `tfsdk:"allpartitions"`
	Communityname   types.String `tfsdk:"communityname"`
	Destport        types.Int64  `tfsdk:"destport"`
	Severity        types.String `tfsdk:"severity"`
	Srcip           types.String `tfsdk:"srcip"`
	Td              types.Int64  `tfsdk:"td"`
	Trapclass       types.String `tfsdk:"trapclass"`
	Trapdestination types.String `tfsdk:"trapdestination"`
	Version         types.String `tfsdk:"version"`
}

func (r *SnmptrapResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the snmptrap resource.",
			},
			"allpartitions": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Send traps of all partitions to this destination.",
			},
			"communityname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password (string) sent with the trap messages, so that the trap listener can authenticate them. Can include 1 to 31 uppercase or lowercase letters, numbers, and hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters.  \nYou must specify the same community string on the trap listener device. Otherwise, the trap listener drops the trap messages.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the string includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my string\" or 'my string').",
			},
			"destport": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(162),
				Description: "UDP port at which the trap listener listens for trap messages. This setting must match the setting on the trap listener. Otherwise, the listener drops the trap messages.",
			},
			"severity": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("Unknown"),
				Description: "Severity level at or above which the Citrix ADC sends trap messages to this trap listener. The severity levels, in increasing order of severity, are Informational, Warning, Minor, Major, Critical. This parameter can be set for trap listeners of type SPECIFIC only. The default is to send all levels of trap messages. \nImportant: Trap messages are not assigned severity levels unless you specify severity levels when configuring SNMP alarms.",
			},
			"srcip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 or IPv6 address that the Citrix ADC inserts as the source IP address in all SNMP trap messages that it sends to this trap listener. By default this is the appliance's NSIP or NSIP6 address, but you can specify an IPv4 MIP or SNIP/SNIP6 address. In cluster setup, the default value is the individual node's NSIP, but it can be set to CLIP or Striped SNIP address. In non default partition, this parameter must be set to the SNIP/SNIP6 address.",
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
			"version": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("V2"),
				Description: "SNMP version, which determines the format of trap messages sent to the trap listener. \nThis setting must match the setting on the trap listener. Otherwise, the listener drops the trap messages.",
			},
		},
	}
}

func snmptrapGetThePayloadFromtheConfig(ctx context.Context, data *SnmptrapResourceModel) snmp.Snmptrap {
	tflog.Debug(ctx, "In snmptrapGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	snmptrap := snmp.Snmptrap{}
	if !data.Allpartitions.IsNull() {
		snmptrap.Allpartitions = data.Allpartitions.ValueString()
	}
	if !data.Communityname.IsNull() {
		snmptrap.Communityname = data.Communityname.ValueString()
	}
	if !data.Destport.IsNull() {
		snmptrap.Destport = utils.IntPtr(int(data.Destport.ValueInt64()))
	}
	if !data.Severity.IsNull() {
		snmptrap.Severity = data.Severity.ValueString()
	}
	if !data.Srcip.IsNull() {
		snmptrap.Srcip = data.Srcip.ValueString()
	}
	if !data.Td.IsNull() {
		snmptrap.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Trapclass.IsNull() {
		snmptrap.Trapclass = data.Trapclass.ValueString()
	}
	if !data.Trapdestination.IsNull() {
		snmptrap.Trapdestination = data.Trapdestination.ValueString()
	}
	if !data.Version.IsNull() {
		snmptrap.Version = data.Version.ValueString()
	}

	return snmptrap
}

func snmptrapSetAttrFromGet(ctx context.Context, data *SnmptrapResourceModel, getResponseData map[string]interface{}) *SnmptrapResourceModel {
	tflog.Debug(ctx, "In snmptrapSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["allpartitions"]; ok && val != nil {
		data.Allpartitions = types.StringValue(val.(string))
	} else {
		data.Allpartitions = types.StringNull()
	}
	if val, ok := getResponseData["communityname"]; ok && val != nil {
		data.Communityname = types.StringValue(val.(string))
	} else {
		data.Communityname = types.StringNull()
	}
	if val, ok := getResponseData["destport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Destport = types.Int64Value(intVal)
		}
	} else {
		data.Destport = types.Int64Null()
	}
	if val, ok := getResponseData["severity"]; ok && val != nil {
		data.Severity = types.StringValue(val.(string))
	} else {
		data.Severity = types.StringNull()
	}
	if val, ok := getResponseData["srcip"]; ok && val != nil {
		data.Srcip = types.StringValue(val.(string))
	} else {
		data.Srcip = types.StringNull()
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
	if val, ok := getResponseData["version"]; ok && val != nil {
		data.Version = types.StringValue(val.(string))
	} else {
		data.Version = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%d,%s,%s,%s", data.Td.ValueInt64(), data.Trapclass.ValueString(), data.Trapdestination.ValueString(), data.Version.ValueString()))

	return data
}
