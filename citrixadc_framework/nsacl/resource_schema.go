package nsacl

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

// NsaclResourceModel describes the resource data model.
type NsaclResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Interface       types.String `tfsdk:"interface"`
	Aclaction       types.String `tfsdk:"aclaction"`
	Aclname         types.String `tfsdk:"aclname"`
	Destip          types.Bool   `tfsdk:"destip"`
	Destipdataset   types.String `tfsdk:"destipdataset"`
	Destipop        types.String `tfsdk:"destipop"`
	Destipval       types.String `tfsdk:"destipval"`
	Destport        types.Bool   `tfsdk:"destport"`
	Destportdataset types.String `tfsdk:"destportdataset"`
	Destportop      types.String `tfsdk:"destportop"`
	Destportval     types.String `tfsdk:"destportval"`
	Dfdhash         types.String `tfsdk:"dfdhash"`
	Established     types.Bool   `tfsdk:"established"`
	Icmpcode        types.Int64  `tfsdk:"icmpcode"`
	Icmptype        types.Int64  `tfsdk:"icmptype"`
	Logstate        types.String `tfsdk:"logstate"`
	Newname         types.String `tfsdk:"newname"`
	Nodeid          types.Int64  `tfsdk:"nodeid"`
	Priority        types.Int64  `tfsdk:"priority"`
	Protocol        types.String `tfsdk:"protocol"`
	Protocolnumber  types.Int64  `tfsdk:"protocolnumber"`
	Ratelimit       types.Int64  `tfsdk:"ratelimit"`
	Srcip           types.Bool   `tfsdk:"srcip"`
	Srcipdataset    types.String `tfsdk:"srcipdataset"`
	Srcipop         types.String `tfsdk:"srcipop"`
	Srcipval        types.String `tfsdk:"srcipval"`
	Srcmac          types.String `tfsdk:"srcmac"`
	Srcmacmask      types.String `tfsdk:"srcmacmask"`
	Srcport         types.Bool   `tfsdk:"srcport"`
	Srcportdataset  types.String `tfsdk:"srcportdataset"`
	Srcportop       types.String `tfsdk:"srcportop"`
	Srcportval      types.String `tfsdk:"srcportval"`
	State           types.String `tfsdk:"state"`
	Stateful        types.String `tfsdk:"stateful"`
	Td              types.Int64  `tfsdk:"td"`
	Ttl             types.Int64  `tfsdk:"ttl"`
	Type            types.String `tfsdk:"type"`
	Vlan            types.Int64  `tfsdk:"vlan"`
	Vxlan           types.Int64  `tfsdk:"vxlan"`
}

func (r *NsaclResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsacl resource.",
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of an interface. The Citrix ADC applies the ACL rule only to the incoming packets from the specified interface. If you do not specify any value, the appliance applies the ACL rule to the incoming packets of all interfaces.",
			},
			"aclaction": schema.StringAttribute{
				Required:    true,
				Description: "Action to perform on incoming IPv4 packets that match the extended ACL rule.\nAvailable settings function as follows:\n* ALLOW - The Citrix ADC processes the packet.\n* BRIDGE - The Citrix ADC bridges the packet to the destination without processing it.\n* DENY - The Citrix ADC drops the packet.",
			},
			"aclname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the extended ACL rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"destip": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the destination IP address of an incoming IPv4 packet.  In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.",
			},
			"destipdataset": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Policy dataset which can have multiple IP ranges bound to it.",
			},
			"destipop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Either the equals (=) or does not equal (!=) logical operator.",
			},
			"destipval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the destination IP address of an incoming IPv4 packet.  In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.",
			},
			"destport": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number or range of port numbers to match against the destination port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.\n\nNote: The destination port can be specified only for TCP and UDP protocols.",
			},
			"destportdataset": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Policy dataset which can have multiple port ranges bound to it.",
			},
			"destportop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Either the equals (=) or does not equal (!=) logical operator.",
			},
			"destportval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number or range of port numbers to match against the destination port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.\n\nNote: The destination port can be specified only for TCP and UDP protocols.",
			},
			"dfdhash": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the type hashmethod to be applied, to steer the packet to the FP of the packet.",
			},
			"established": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow only incoming TCP packets that have the ACK or RST bit set, if the action set for the ACL rule is ALLOW and these packets match the other conditions in the ACL rule.",
			},
			"icmpcode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Code of a particular ICMP message type to match against the ICMP code of an incoming ICMP packet.  For example, to block DESTINATION HOST UNREACHABLE messages, specify 3 as the ICMP type and 1 as the ICMP code.\n\nIf you set this parameter, you must set the ICMP Type parameter.",
			},
			"icmptype": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ICMP Message type to match against the message type of an incoming ICMP packet. For example, to block DESTINATION UNREACHABLE messages, you must specify 3 as the ICMP type.\n\nNote: This parameter can be specified only for the ICMP protocol.",
			},
			"logstate": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable logging of events related to the extended ACL rule. The log messages are stored in the configured syslog or auditlog server.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the extended ACL rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the NodeId to steer the packet to the provided FP.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority for the extended ACL rule that determines the order in which it is evaluated relative to the other extended ACL rules. If you do not specify priorities while creating extended ACL rules, the ACL rules are evaluated in the order in which they are created.",
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol to match against the protocol of an incoming IPv4 packet.",
			},
			"protocolnumber": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol to match against the protocol of an incoming IPv4 packet.",
			},
			"ratelimit": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(100),
				Description: "Maximum number of log messages to be generated per second. If you set this parameter, you must enable the Log State parameter.",
			},
			"srcip": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the source IP address of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.",
			},
			"srcipdataset": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Policy dataset which can have multiple IP ranges bound to it.",
			},
			"srcipop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Either the equals (=) or does not equal (!=) logical operator.",
			},
			"srcipval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the source IP address of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example:10.102.29.30-10.102.29.189.",
			},
			"srcmac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "MAC address to match against the source MAC address of an incoming IPv4 packet.",
			},
			"srcmacmask": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("000000000000"),
				Description: "Used to define range of Source MAC address. It takes string of 0 and 1, 0s are for exact match and 1s for wildcard. For matching first 3 bytes of MAC address, srcMacMask value \"000000111111\".",
			},
			"srcport": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number or range of port numbers to match against the source port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.",
			},
			"srcportdataset": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Policy dataset which can have multiple port ranges bound to it.",
			},
			"srcportop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Either the equals (=) or does not equal (!=) logical operator.",
			},
			"srcportval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number or range of port numbers to match against the source port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable the extended ACL rule. After you apply the extended ACL rules, the Citrix ADC compares incoming packets against the enabled extended ACL rules.",
			},
			"stateful": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If stateful option is enabled, transparent sessions are created for the traffic hitting this ACL and not hitting any other features like LB, INAT etc.",
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
				Description: "Number of seconds, in multiples of four, after which the extended ACL rule expires. If you do not want the extended ACL rule to expire, do not specify a TTL value.",
			},
			"type": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("CLASSIC"),
				Description: "Type of the acl ,default will be CLASSIC.\nAvailable options as follows:\n* CLASSIC - specifies the regular extended acls.\n* DFD - cluster specific acls,specifies hashmethod for steering of the packet in cluster .",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the VLAN. The Citrix ADC applies the ACL rule only to the incoming packets of the specified VLAN. If you do not specify a VLAN ID, the appliance applies the ACL rule to the incoming packets on all VLANs.",
			},
			"vxlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the VXLAN. The Citrix ADC applies the ACL rule only to the incoming packets of the specified VXLAN. If you do not specify a VXLAN ID, the appliance applies the ACL rule to the incoming packets on all VXLANs.",
			},
		},
	}
}

func nsaclGetThePayloadFromtheConfig(ctx context.Context, data *NsaclResourceModel) ns.Nsacl {
	tflog.Debug(ctx, "In nsaclGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsacl := ns.Nsacl{}
	if !data.Interface.IsNull() {
		nsacl.Interface = data.Interface.ValueString()
	}
	if !data.Aclaction.IsNull() {
		nsacl.Aclaction = data.Aclaction.ValueString()
	}
	if !data.Aclname.IsNull() {
		nsacl.Aclname = data.Aclname.ValueString()
	}
	if !data.Destip.IsNull() {
		nsacl.Destip = data.Destip.ValueBool()
	}
	if !data.Destipdataset.IsNull() {
		nsacl.Destipdataset = data.Destipdataset.ValueString()
	}
	if !data.Destipop.IsNull() {
		nsacl.Destipop = data.Destipop.ValueString()
	}
	if !data.Destipval.IsNull() {
		nsacl.Destipval = data.Destipval.ValueString()
	}
	if !data.Destport.IsNull() {
		nsacl.Destport = data.Destport.ValueBool()
	}
	if !data.Destportdataset.IsNull() {
		nsacl.Destportdataset = data.Destportdataset.ValueString()
	}
	if !data.Destportop.IsNull() {
		nsacl.Destportop = data.Destportop.ValueString()
	}
	if !data.Destportval.IsNull() {
		nsacl.Destportval = data.Destportval.ValueString()
	}
	if !data.Dfdhash.IsNull() {
		nsacl.Dfdhash = data.Dfdhash.ValueString()
	}
	if !data.Established.IsNull() {
		nsacl.Established = data.Established.ValueBool()
	}
	if !data.Icmpcode.IsNull() {
		nsacl.Icmpcode = utils.IntPtr(int(data.Icmpcode.ValueInt64()))
	}
	if !data.Icmptype.IsNull() {
		nsacl.Icmptype = utils.IntPtr(int(data.Icmptype.ValueInt64()))
	}
	if !data.Logstate.IsNull() {
		nsacl.Logstate = data.Logstate.ValueString()
	}
	if !data.Newname.IsNull() {
		nsacl.Newname = data.Newname.ValueString()
	}
	if !data.Nodeid.IsNull() {
		nsacl.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Priority.IsNull() {
		nsacl.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Protocol.IsNull() {
		nsacl.Protocol = data.Protocol.ValueString()
	}
	if !data.Protocolnumber.IsNull() {
		nsacl.Protocolnumber = utils.IntPtr(int(data.Protocolnumber.ValueInt64()))
	}
	if !data.Ratelimit.IsNull() {
		nsacl.Ratelimit = utils.IntPtr(int(data.Ratelimit.ValueInt64()))
	}
	if !data.Srcip.IsNull() {
		nsacl.Srcip = data.Srcip.ValueBool()
	}
	if !data.Srcipdataset.IsNull() {
		nsacl.Srcipdataset = data.Srcipdataset.ValueString()
	}
	if !data.Srcipop.IsNull() {
		nsacl.Srcipop = data.Srcipop.ValueString()
	}
	if !data.Srcipval.IsNull() {
		nsacl.Srcipval = data.Srcipval.ValueString()
	}
	if !data.Srcmac.IsNull() {
		nsacl.Srcmac = data.Srcmac.ValueString()
	}
	if !data.Srcmacmask.IsNull() {
		nsacl.Srcmacmask = data.Srcmacmask.ValueString()
	}
	if !data.Srcport.IsNull() {
		nsacl.Srcport = data.Srcport.ValueBool()
	}
	if !data.Srcportdataset.IsNull() {
		nsacl.Srcportdataset = data.Srcportdataset.ValueString()
	}
	if !data.Srcportop.IsNull() {
		nsacl.Srcportop = data.Srcportop.ValueString()
	}
	if !data.Srcportval.IsNull() {
		nsacl.Srcportval = data.Srcportval.ValueString()
	}
	if !data.State.IsNull() {
		nsacl.State = data.State.ValueString()
	}
	if !data.Stateful.IsNull() {
		nsacl.Stateful = data.Stateful.ValueString()
	}
	if !data.Td.IsNull() {
		nsacl.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Ttl.IsNull() {
		nsacl.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}
	if !data.Type.IsNull() {
		nsacl.Type = data.Type.ValueString()
	}
	if !data.Vlan.IsNull() {
		nsacl.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}
	if !data.Vxlan.IsNull() {
		nsacl.Vxlan = utils.IntPtr(int(data.Vxlan.ValueInt64()))
	}

	return nsacl
}

func nsaclSetAttrFromGet(ctx context.Context, data *NsaclResourceModel, getResponseData map[string]interface{}) *NsaclResourceModel {
	tflog.Debug(ctx, "In nsaclSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["Interface"]; ok && val != nil {
		data.Interface = types.StringValue(val.(string))
	} else {
		data.Interface = types.StringNull()
	}
	if val, ok := getResponseData["aclaction"]; ok && val != nil {
		data.Aclaction = types.StringValue(val.(string))
	} else {
		data.Aclaction = types.StringNull()
	}
	if val, ok := getResponseData["aclname"]; ok && val != nil {
		data.Aclname = types.StringValue(val.(string))
	} else {
		data.Aclname = types.StringNull()
	}
	if val, ok := getResponseData["destip"]; ok && val != nil {
		data.Destip = types.BoolValue(val.(bool))
	} else {
		data.Destip = types.BoolNull()
	}
	if val, ok := getResponseData["destipdataset"]; ok && val != nil {
		data.Destipdataset = types.StringValue(val.(string))
	} else {
		data.Destipdataset = types.StringNull()
	}
	if val, ok := getResponseData["destipop"]; ok && val != nil {
		data.Destipop = types.StringValue(val.(string))
	} else {
		data.Destipop = types.StringNull()
	}
	if val, ok := getResponseData["destipval"]; ok && val != nil {
		data.Destipval = types.StringValue(val.(string))
	} else {
		data.Destipval = types.StringNull()
	}
	if val, ok := getResponseData["destport"]; ok && val != nil {
		data.Destport = types.BoolValue(val.(bool))
	} else {
		data.Destport = types.BoolNull()
	}
	if val, ok := getResponseData["destportdataset"]; ok && val != nil {
		data.Destportdataset = types.StringValue(val.(string))
	} else {
		data.Destportdataset = types.StringNull()
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
	if val, ok := getResponseData["srcip"]; ok && val != nil {
		data.Srcip = types.BoolValue(val.(bool))
	} else {
		data.Srcip = types.BoolNull()
	}
	if val, ok := getResponseData["srcipdataset"]; ok && val != nil {
		data.Srcipdataset = types.StringValue(val.(string))
	} else {
		data.Srcipdataset = types.StringNull()
	}
	if val, ok := getResponseData["srcipop"]; ok && val != nil {
		data.Srcipop = types.StringValue(val.(string))
	} else {
		data.Srcipop = types.StringNull()
	}
	if val, ok := getResponseData["srcipval"]; ok && val != nil {
		data.Srcipval = types.StringValue(val.(string))
	} else {
		data.Srcipval = types.StringNull()
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
	if val, ok := getResponseData["srcportdataset"]; ok && val != nil {
		data.Srcportdataset = types.StringValue(val.(string))
	} else {
		data.Srcportdataset = types.StringNull()
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
	data.Id = types.StringValue(fmt.Sprintf("%s,%s", data.Aclname.ValueString(), data.Type.ValueString()))

	return data
}
