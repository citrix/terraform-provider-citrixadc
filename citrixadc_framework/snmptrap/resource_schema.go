package snmptrap

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
			// SDK v2: Optional + Computed (no Default) -- value read back from ADC.
			"allpartitions": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Send traps of all partitions to this destination.",
			},
			// SDK v2: Optional + Computed (no Default).
			"communityname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password (string) sent with the trap messages, so that the trap listener can authenticate them. Can include 1 to 31 uppercase or lowercase letters, numbers, and hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters.  \nYou must specify the same community string on the trap listener device. Otherwise, the trap listener drops the trap messages.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the string includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my string\" or 'my string').",
			},
			// SDK v2: Optional + Computed (no Default).
			"destport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(162),
				Description: "UDP port at which the trap listener listens for trap messages. This setting must match the setting on the trap listener. Otherwise, the listener drops the trap messages.",
			},
			// SDK v2: Optional + Computed (no Default).
			"severity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("Unknown"),
				Description: "Severity level at or above which the Citrix ADC sends trap messages to this trap listener. The severity levels, in increasing order of severity, are Informational, Warning, Minor, Major, Critical. This parameter can be set for trap listeners of type SPECIFIC only. The default is to send all levels of trap messages. \nImportant: Trap messages are not assigned severity levels unless you specify severity levels when configuring SNMP alarms.",
			},
			// SDK v2: Optional + Computed (no Default).
			"srcip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 or IPv6 address that the Citrix ADC inserts as the source IP address in all SNMP trap messages that it sends to this trap listener. By default this is the appliance's NSIP or NSIP6 address, but you can specify an IPv4 MIP or SNIP/SNIP6 address. In cluster setup, the default value is the individual node's NSIP, but it can be set to CLIP or Striped SNIP address. In non default partition, this parameter must be set to the SNIP/SNIP6 address.",
			},
			// SDK v2: Optional + Computed (no Default). Not ForceNew.
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			// SDK v2: Required + ForceNew -> RequiresReplace.
			"trapclass": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of trap messages that the Citrix ADC sends to the trap listener: Generic or the enterprise-specific messages defined in the MIB file.",
			},
			// SDK v2: Required + ForceNew -> RequiresReplace.
			"trapdestination": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IPv4 or the IPv6 address of the trap listener to which the Citrix ADC is to send SNMP trap messages.",
			},
			// SDK v2: Optional + Default("V2") + ForceNew. Default requires Computed in
			// the Framework; ForceNew -> RequiresReplace. UseStateForUnknown keeps the
			// value stable for the Computed attribute.
			"version": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("V2"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
				Description: "SNMP version, which determines the format of trap messages sent to the trap listener. \nThis setting must match the setting on the trap listener. Otherwise, the listener drops the trap messages.",
			},
		},
	}
}

func snmptrapGetThePayloadFromthePlan(ctx context.Context, data *SnmptrapResourceModel) snmp.Snmptrap {
	tflog.Debug(ctx, "In snmptrapGetThePayloadFromthePlan Function")

	// Create API request body from the model
	snmptrap := snmp.Snmptrap{}
	if !data.Allpartitions.IsNull() && !data.Allpartitions.IsUnknown() {
		snmptrap.Allpartitions = data.Allpartitions.ValueString()
	}
	if !data.Communityname.IsNull() && !data.Communityname.IsUnknown() {
		snmptrap.Communityname = data.Communityname.ValueString()
	}
	if !data.Destport.IsNull() && !data.Destport.IsUnknown() {
		snmptrap.Destport = utils.IntPtr(int(data.Destport.ValueInt64()))
	}
	if !data.Severity.IsNull() && !data.Severity.IsUnknown() {
		snmptrap.Severity = data.Severity.ValueString()
	}
	if !data.Srcip.IsNull() && !data.Srcip.IsUnknown() {
		snmptrap.Srcip = data.Srcip.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		snmptrap.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Trapclass.IsNull() && !data.Trapclass.IsUnknown() {
		snmptrap.Trapclass = data.Trapclass.ValueString()
	}
	if !data.Trapdestination.IsNull() && !data.Trapdestination.IsUnknown() {
		snmptrap.Trapdestination = data.Trapdestination.ValueString()
	}
	if !data.Version.IsNull() && !data.Version.IsUnknown() {
		snmptrap.Version = data.Version.ValueString()
	}

	return snmptrap
}

func snmptrapSetAttrFromGet(ctx context.Context, data *SnmptrapResourceModel, getResponseData map[string]interface{}) *SnmptrapResourceModel {
	tflog.Debug(ctx, "In snmptrapSetAttrFromGet Function")

	// Convert API response to model.
	// omit-on-default guard: when NITRO omits a value from GET, only null the
	// attribute if it is currently unknown -- never clobber a known configured value.
	if val, ok := getResponseData["allpartitions"]; ok && val != nil {
		data.Allpartitions = types.StringValue(val.(string))
	} else if data.Allpartitions.IsUnknown() {
		data.Allpartitions = types.StringNull()
	}
	if val, ok := getResponseData["communityname"]; ok && val != nil {
		data.Communityname = types.StringValue(val.(string))
	} else if data.Communityname.IsUnknown() {
		data.Communityname = types.StringNull()
	}
	if val, ok := getResponseData["destport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Destport = types.Int64Value(intVal)
		}
	} else if data.Destport.IsUnknown() {
		data.Destport = types.Int64Null()
	}
	if val, ok := getResponseData["severity"]; ok && val != nil {
		data.Severity = types.StringValue(val.(string))
	} else if data.Severity.IsUnknown() {
		data.Severity = types.StringNull()
	}
	if val, ok := getResponseData["srcip"]; ok && val != nil {
		data.Srcip = types.StringValue(val.(string))
	} else if data.Srcip.IsUnknown() {
		data.Srcip = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else if data.Td.IsUnknown() {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["trapclass"]; ok && val != nil {
		data.Trapclass = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["trapdestination"]; ok && val != nil {
		data.Trapdestination = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["version"]; ok && val != nil {
		data.Version = types.StringValue(val.(string))
	}

	// Set ID for the resource.
	// SDK v2 backward-compatible composite ID: "trapclass,trapdestination,version".
	data.Id = types.StringValue(fmt.Sprintf("%s,%s,%s", data.Trapclass.ValueString(), data.Trapdestination.ValueString(), data.Version.ValueString()))

	return data
}
