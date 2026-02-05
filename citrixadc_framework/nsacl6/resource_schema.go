package nsacl6

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Nsacl6ResourceModel describes the resource data model.
type Nsacl6ResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Interface      types.String `tfsdk:"interface"`
	Acl6action     types.String `tfsdk:"acl6action"`
	Acl6name       types.String `tfsdk:"acl6name"`
	Aclaction      types.String `tfsdk:"aclaction"`
	Destipop       types.String `tfsdk:"destipop"`
	Destipv6       types.Bool   `tfsdk:"destipv6"`
	Destipv6val    types.String `tfsdk:"destipv6val"`
	Destport       types.Bool   `tfsdk:"destport"`
	Destportop     types.String `tfsdk:"destportop"`
	Destportval    types.String `tfsdk:"destportval"`
	Dfdhash        types.String `tfsdk:"dfdhash"`
	Dfdprefix      types.Int64  `tfsdk:"dfdprefix"`
	Established    types.Bool   `tfsdk:"established"`
	Icmpcode       types.Int64  `tfsdk:"icmpcode"`
	Icmptype       types.Int64  `tfsdk:"icmptype"`
	Logstate       types.String `tfsdk:"logstate"`
	Newname        types.String `tfsdk:"newname"`
	Nodeid         types.Int64  `tfsdk:"nodeid"`
	Priority       types.Int64  `tfsdk:"priority"`
	Protocol       types.String `tfsdk:"protocol"`
	Protocolnumber types.Int64  `tfsdk:"protocolnumber"`
	Ratelimit      types.Int64  `tfsdk:"ratelimit"`
	Srcipop        types.String `tfsdk:"srcipop"`
	Srcipv6        types.Bool   `tfsdk:"srcipv6"`
	Srcipv6val     types.String `tfsdk:"srcipv6val"`
	Srcmac         types.String `tfsdk:"srcmac"`
	Srcmacmask     types.String `tfsdk:"srcmacmask"`
	Srcport        types.Bool   `tfsdk:"srcport"`
	Srcportop      types.String `tfsdk:"srcportop"`
	Srcportval     types.String `tfsdk:"srcportval"`
	State          types.String `tfsdk:"state"`
	Stateful       types.String `tfsdk:"stateful"`
	Td             types.Int64  `tfsdk:"td"`
	Ttl            types.Int64  `tfsdk:"ttl"`
	Type           types.String `tfsdk:"type"`
	Vlan           types.Int64  `tfsdk:"vlan"`
	Vxlan          types.Int64  `tfsdk:"vxlan"`
}

func (r *Nsacl6Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsacl6 resource.",
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of an interface. The Citrix ADC applies the ACL6 rule only to the incoming packets from the specified interface. If you do not specify any value, the appliance applies the ACL6 rule to the incoming packets from all interfaces.",
			},
			"acl6action": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Action to perform on the incoming IPv6 packets that match the ACL6 rule.\nAvailable settings function as follows:\n* ALLOW - The Citrix ADC processes the packet.\n* BRIDGE - The Citrix ADC bridges the packet to the destination without processing it.\n* DENY - The Citrix ADC drops the packet.",
			},
			"acl6name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the ACL6 rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"aclaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action associated with the ACL6.",
			},
			"destipop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Either the equals (=) or does not equal (!=) logical operator.",
			},
			"destipv6": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the destination IP address of an incoming IPv6 packet.  In the command line interface, separate the range with a hyphen.",
			},
			"destipv6val": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Destination IPv6 address (range).",
			},
			"destport": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number or range of port numbers to match against the destination port number of an incoming IPv6 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.\n\nNote: The destination port can be specified only for TCP and UDP protocols.",
			},
			"destportop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Either the equals (=) or does not equal (!=) logical operator.",
			},
			"destportval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Destination port (range).",
			},
			"dfdhash": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the type of hashmethod to be applied, to steer the packet to the FP of the packet.",
			},
			"dfdprefix": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(128),
				Description: "hashprefix to be applied to SIP/DIP to generate rsshash FP.eg 128 => hash calculated on the complete IP",
			},
			"established": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow only incoming TCP packets that have the ACK or RST bit set if the action set for the ACL6 rule is ALLOW and these packets match the other conditions in the ACL6 rule.",
			},
			"icmpcode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Code of a particular ICMP message type to match against the ICMP code of an incoming IPv6 ICMP packet.  For example, to block DESTINATION HOST UNREACHABLE messages, specify 3 as the ICMP type and 1 as the ICMP code.\n\nIf you set this parameter, you must set the ICMP Type parameter.",
			},
			"icmptype": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ICMP Message type to match against the message type of an incoming IPv6 ICMP packet. For example, to block DESTINATION UNREACHABLE messages, you must specify 3 as the ICMP type.\n\nNote: This parameter can be specified only for the ICMP protocol.",
			},
			"logstate": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable logging of events related to the ACL6 rule. The log messages are stored in the configured syslog or auditlog server.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the ACL6 rule. Must begin with an ASCII alphabetic or underscore \\(_\\) character, and must contain only ASCII alphanumeric, underscore, hash \\(\\#\\), period \\(.\\), space, colon \\(:\\), at \\(@\\), equals \\(=\\), and hyphen \\(-\\) characters.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the NodeId to steer the packet to the provided FP.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority for the ACL6 rule, which determines the order in which it is evaluated relative to the other ACL6 rules. If you do not specify priorities while creating ACL6 rules, the ACL6 rules are evaluated in the order in which they are created.",
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol, identified by protocol name, to match against the protocol of an incoming IPv6 packet.",
			},
			"protocolnumber": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol, identified by protocol number, to match against the protocol of an incoming IPv6 packet.",
			},
			"ratelimit": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(100),
				Description: "Maximum number of log messages to be generated per second. If you set this parameter, you must enable the Log State parameter.",
			},
			"srcipop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Either the equals (=) or does not equal (!=) logical operator.",
			},
			"srcipv6": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the source IP address of an incoming IPv6 packet. In the command line interface, separate the range with a hyphen.",
			},
			"srcipv6val": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source IPv6 address (range).",
			},
			"srcmac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "MAC address to match against the source MAC address of an incoming IPv6 packet.",
			},
			"srcmacmask": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("000000000000"),
				Description: "Used to define range of Source MAC address. It takes string of 0 and 1, 0s are for exact match and 1s for wildcard. For matching first 3 bytes of MAC address, srcMacMask value \"000000111111\".",
			},
			"srcport": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number or range of port numbers to match against the source port number of an incoming IPv6 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.\n\nNote: The destination port can be specified only for TCP and UDP protocols.",
			},
			"srcportop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Either the equals (=) or does not equal (!=) logical operator.",
			},
			"srcportval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source port (range).",
			},
			"state": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of the ACL6.",
			},
			"stateful": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If stateful option is enabled, transparent sessions are created for the traffic hitting this ACL6 and not hitting any other features like LB, INAT etc.",
			},
			"td": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"ttl": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Time to expire this ACL6 (in seconds).",
			},
			"type": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("CLASSIC"),
				Description: "Type of the acl6 ,default will be CLASSIC.\nAvailable options as follows:\n* CLASSIC - specifies the regular extended acls.\n* DFD - cluster specific acls,specifies hashmethod for steering of the packet in cluster .",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the VLAN. The Citrix ADC applies the ACL6 rule only to the incoming packets on the specified VLAN. If you do not specify a VLAN ID, the appliance applies the ACL6 rule to the incoming packets on all VLANs.",
			},
			"vxlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the VXLAN. The Citrix ADC applies the ACL6 rule only to the incoming packets on the specified VXLAN. If you do not specify a VXLAN ID, the appliance applies the ACL6 rule to the incoming packets on all VXLANs.",
			},
		},
	}
}

func nsacl6GetThePayloadFromtheConfig(ctx context.Context, data *Nsacl6ResourceModel) ns.Nsacl6 {
	tflog.Debug(ctx, "In nsacl6GetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsacl6 := ns.Nsacl6{}
	if !data.Interface.IsNull() {
		nsacl6.Interface = data.Interface.ValueString()
	}
	if !data.Acl6action.IsNull() {
		nsacl6.Acl6action = data.Acl6action.ValueString()
	}
	if !data.Acl6name.IsNull() {
		nsacl6.Acl6name = data.Acl6name.ValueString()
	}
	if !data.Aclaction.IsNull() {
		nsacl6.Aclaction = data.Aclaction.ValueString()
	}
	if !data.Destipop.IsNull() {
		nsacl6.Destipop = data.Destipop.ValueString()
	}
	if !data.Destipv6.IsNull() {
		nsacl6.Destipv6 = data.Destipv6.ValueBool()
	}
	if !data.Destipv6val.IsNull() {
		nsacl6.Destipv6val = data.Destipv6val.ValueString()
	}
	if !data.Destport.IsNull() {
		nsacl6.Destport = data.Destport.ValueBool()
	}
	if !data.Destportop.IsNull() {
		nsacl6.Destportop = data.Destportop.ValueString()
	}
	if !data.Destportval.IsNull() {
		nsacl6.Destportval = data.Destportval.ValueString()
	}
	if !data.Dfdhash.IsNull() {
		nsacl6.Dfdhash = data.Dfdhash.ValueString()
	}
	if !data.Dfdprefix.IsNull() {
		nsacl6.Dfdprefix = utils.IntPtr(int(data.Dfdprefix.ValueInt64()))
	}
	if !data.Established.IsNull() {
		nsacl6.Established = data.Established.ValueBool()
	}
	if !data.Icmpcode.IsNull() {
		nsacl6.Icmpcode = utils.IntPtr(int(data.Icmpcode.ValueInt64()))
	}
	if !data.Icmptype.IsNull() {
		nsacl6.Icmptype = utils.IntPtr(int(data.Icmptype.ValueInt64()))
	}
	if !data.Logstate.IsNull() {
		nsacl6.Logstate = data.Logstate.ValueString()
	}
	if !data.Newname.IsNull() {
		nsacl6.Newname = data.Newname.ValueString()
	}
	if !data.Nodeid.IsNull() {
		nsacl6.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Priority.IsNull() {
		nsacl6.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Protocol.IsNull() {
		nsacl6.Protocol = data.Protocol.ValueString()
	}
	if !data.Protocolnumber.IsNull() {
		nsacl6.Protocolnumber = utils.IntPtr(int(data.Protocolnumber.ValueInt64()))
	}
	if !data.Ratelimit.IsNull() {
		nsacl6.Ratelimit = utils.IntPtr(int(data.Ratelimit.ValueInt64()))
	}
	if !data.Srcipop.IsNull() {
		nsacl6.Srcipop = data.Srcipop.ValueString()
	}
	if !data.Srcipv6.IsNull() {
		nsacl6.Srcipv6 = data.Srcipv6.ValueBool()
	}
	if !data.Srcipv6val.IsNull() {
		nsacl6.Srcipv6val = data.Srcipv6val.ValueString()
	}
	if !data.Srcmac.IsNull() {
		nsacl6.Srcmac = data.Srcmac.ValueString()
	}
	if !data.Srcmacmask.IsNull() {
		nsacl6.Srcmacmask = data.Srcmacmask.ValueString()
	}
	if !data.Srcport.IsNull() {
		nsacl6.Srcport = data.Srcport.ValueBool()
	}
	if !data.Srcportop.IsNull() {
		nsacl6.Srcportop = data.Srcportop.ValueString()
	}
	if !data.Srcportval.IsNull() {
		nsacl6.Srcportval = data.Srcportval.ValueString()
	}
	if !data.State.IsNull() {
		nsacl6.State = data.State.ValueString()
	}
	if !data.Stateful.IsNull() {
		nsacl6.Stateful = data.Stateful.ValueString()
	}
	if !data.Td.IsNull() {
		nsacl6.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Ttl.IsNull() {
		nsacl6.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}
	if !data.Type.IsNull() {
		nsacl6.Type = data.Type.ValueString()
	}
	if !data.Vlan.IsNull() {
		nsacl6.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}
	if !data.Vxlan.IsNull() {
		nsacl6.Vxlan = utils.IntPtr(int(data.Vxlan.ValueInt64()))
	}

	return nsacl6
}

func nsacl6SetAttrFromGet(ctx context.Context, data *Nsacl6ResourceModel, getResponseData map[string]interface{}) *Nsacl6ResourceModel {
	tflog.Debug(ctx, "In nsacl6SetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["Interface"]; ok && val != nil {
		data.Interface = types.StringValue(val.(string))
	} else {
		data.Interface = types.StringNull()
	}
	if val, ok := getResponseData["acl6action"]; ok && val != nil {
		data.Acl6action = types.StringValue(val.(string))
	} else {
		data.Acl6action = types.StringNull()
	}
	if val, ok := getResponseData["acl6name"]; ok && val != nil {
		data.Acl6name = types.StringValue(val.(string))
	} else {
		data.Acl6name = types.StringNull()
	}
	if val, ok := getResponseData["aclaction"]; ok && val != nil {
		data.Aclaction = types.StringValue(val.(string))
	} else {
		data.Aclaction = types.StringNull()
	}
	if val, ok := getResponseData["destipop"]; ok && val != nil {
		data.Destipop = types.StringValue(val.(string))
	} else {
		data.Destipop = types.StringNull()
	}
	if val, ok := getResponseData["destipv6"]; ok && val != nil {
		data.Destipv6 = types.BoolValue(val.(bool))
	} else {
		data.Destipv6 = types.BoolNull()
	}
	if val, ok := getResponseData["destipv6val"]; ok && val != nil {
		data.Destipv6val = types.StringValue(val.(string))
	} else {
		data.Destipv6val = types.StringNull()
	}
	if val, ok := getResponseData["destport"]; ok && val != nil {
		data.Destport = types.BoolValue(val.(bool))
	} else {
		data.Destport = types.BoolNull()
	}
	if val, ok := getResponseData["destportop"]; ok && val != nil {
		data.Destportop = types.StringValue(val.(string))
	} else {
		data.Destportop = types.StringNull()
	}
	if val, ok := getResponseData["destportval"]; ok && val != nil {
		data.Destportval = types.StringValue(val.(string))
	} else {
		data.Destportval = types.StringNull()
	}
	if val, ok := getResponseData["dfdhash"]; ok && val != nil {
		data.Dfdhash = types.StringValue(val.(string))
	} else {
		data.Dfdhash = types.StringNull()
	}
	if val, ok := getResponseData["dfdprefix"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Dfdprefix = types.Int64Value(intVal)
		}
	} else {
		data.Dfdprefix = types.Int64Null()
	}
	if val, ok := getResponseData["established"]; ok && val != nil {
		data.Established = types.BoolValue(val.(bool))
	} else {
		data.Established = types.BoolNull()
	}
	if val, ok := getResponseData["icmpcode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Icmpcode = types.Int64Value(intVal)
		}
	} else {
		data.Icmpcode = types.Int64Null()
	}
	if val, ok := getResponseData["icmptype"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Icmptype = types.Int64Value(intVal)
		}
	} else {
		data.Icmptype = types.Int64Null()
	}
	if val, ok := getResponseData["logstate"]; ok && val != nil {
		data.Logstate = types.StringValue(val.(string))
	} else {
		data.Logstate = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["protocol"]; ok && val != nil {
		data.Protocol = types.StringValue(val.(string))
	} else {
		data.Protocol = types.StringNull()
	}
	if val, ok := getResponseData["protocolnumber"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Protocolnumber = types.Int64Value(intVal)
		}
	} else {
		data.Protocolnumber = types.Int64Null()
	}
	if val, ok := getResponseData["ratelimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ratelimit = types.Int64Value(intVal)
		}
	} else {
		data.Ratelimit = types.Int64Null()
	}
	if val, ok := getResponseData["srcipop"]; ok && val != nil {
		data.Srcipop = types.StringValue(val.(string))
	} else {
		data.Srcipop = types.StringNull()
	}
	if val, ok := getResponseData["srcipv6"]; ok && val != nil {
		data.Srcipv6 = types.BoolValue(val.(bool))
	} else {
		data.Srcipv6 = types.BoolNull()
	}
	if val, ok := getResponseData["srcipv6val"]; ok && val != nil {
		data.Srcipv6val = types.StringValue(val.(string))
	} else {
		data.Srcipv6val = types.StringNull()
	}
	if val, ok := getResponseData["srcmac"]; ok && val != nil {
		data.Srcmac = types.StringValue(val.(string))
	} else {
		data.Srcmac = types.StringNull()
	}
	if val, ok := getResponseData["srcmacmask"]; ok && val != nil {
		data.Srcmacmask = types.StringValue(val.(string))
	} else {
		data.Srcmacmask = types.StringNull()
	}
	if val, ok := getResponseData["srcport"]; ok && val != nil {
		data.Srcport = types.BoolValue(val.(bool))
	} else {
		data.Srcport = types.BoolNull()
	}
	if val, ok := getResponseData["srcportop"]; ok && val != nil {
		data.Srcportop = types.StringValue(val.(string))
	} else {
		data.Srcportop = types.StringNull()
	}
	if val, ok := getResponseData["srcportval"]; ok && val != nil {
		data.Srcportval = types.StringValue(val.(string))
	} else {
		data.Srcportval = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["stateful"]; ok && val != nil {
		data.Stateful = types.StringValue(val.(string))
	} else {
		data.Stateful = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["ttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ttl = types.Int64Value(intVal)
		}
	} else {
		data.Ttl = types.Int64Null()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	} else {
		data.Vlan = types.Int64Null()
	}
	if val, ok := getResponseData["vxlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vxlan = types.Int64Value(intVal)
		}
	} else {
		data.Vxlan = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s,%s", data.Acl6name.ValueString(), data.Type.ValueString()))

	return data
}
