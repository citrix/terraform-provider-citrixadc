package nspbr6

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Nspbr6ResourceModel describes the resource data model.
type Nspbr6ResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Interface      types.String `tfsdk:"interface"`
	Action         types.String `tfsdk:"action"`
	Destipop       types.String `tfsdk:"destipop"`
	Destipv6       types.Bool   `tfsdk:"destipv6"`
	Destipv6val    types.String `tfsdk:"destipv6val"`
	Destport       types.Bool   `tfsdk:"destport"`
	Destportop     types.String `tfsdk:"destportop"`
	Destportval    types.String `tfsdk:"destportval"`
	Detail         types.Bool   `tfsdk:"detail"`
	Iptunnel       types.String `tfsdk:"iptunnel"`
	Monitor        types.String `tfsdk:"monitor"`
	Msr            types.String `tfsdk:"msr"`
	Name           types.String `tfsdk:"name"`
	Nexthop        types.Bool   `tfsdk:"nexthop"`
	Nexthopval     types.String `tfsdk:"nexthopval"`
	Nexthopvlan    types.Int64  `tfsdk:"nexthopvlan"`
	Ownergroup     types.String `tfsdk:"ownergroup"`
	Priority       types.Int64  `tfsdk:"priority"`
	Protocol       types.String `tfsdk:"protocol"`
	Protocolnumber types.Int64  `tfsdk:"protocolnumber"`
	Srcipop        types.String `tfsdk:"srcipop"`
	Srcipv6        types.Bool   `tfsdk:"srcipv6"`
	Srcipv6val     types.String `tfsdk:"srcipv6val"`
	Srcmac         types.String `tfsdk:"srcmac"`
	Srcmacmask     types.String `tfsdk:"srcmacmask"`
	Srcport        types.Bool   `tfsdk:"srcport"`
	Srcportop      types.String `tfsdk:"srcportop"`
	Srcportval     types.String `tfsdk:"srcportval"`
	State          types.String `tfsdk:"state"`
	Td             types.Int64  `tfsdk:"td"`
	Vlan           types.Int64  `tfsdk:"vlan"`
	Vxlan          types.Int64  `tfsdk:"vxlan"`
	Vxlanvlanmap   types.String `tfsdk:"vxlanvlanmap"`
}

func (r *Nspbr6Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nspbr6 resource.",
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of an interface. The Citrix ADC compares the PBR6 only to the outgoing packets on the specified interface. If you do not specify a value, the appliance compares the PBR6 to the outgoing packets on all interfaces.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Action to perform on the outgoing IPv6 packets that match the PBR6.\n\nAvailable settings function as follows:\n* ALLOW - The Citrix ADC sends the packet to the designated next-hop router.\n* DENY - The Citrix ADC applies the routing table for normal destination-based routing.",
			},
			"destipop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Either the equals (=) or does not equal (!=) logical operator.",
			},
			"destipv6": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the destination IP address of an outgoing IPv6 packet.  In the command line interface, separate the range with a hyphen.",
			},
			"destipv6val": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the destination IP address of an outgoing IPv6 packet.  In the command line interface, separate the range with a hyphen.",
			},
			"destport": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number or range of port numbers to match against the destination port number of an outgoing IPv6 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.\n\nNote: The destination port can be specified only for TCP and UDP protocols.",
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
			"detail": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "To get a detailed view.",
			},
			"iptunnel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The iptunnel name where packets need to be forwarded upon.",
			},
			"monitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the monitor.(Can be only of type ping or ARP )",
			},
			"msr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Monitor the route specified by the Next Hop parameter.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the PBR6. Must begin with an ASCII alphabetic or underscore \\(_\\) character, and must contain only ASCII alphanumeric, underscore, hash \\(\\#\\), period \\(.\\), space, colon \\(:\\), at \\(@\\), equals \\(=\\), and hyphen \\(-\\) characters. Cannot be changed after the PBR6 is created.",
			},
			"nexthop": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the next hop router to which to send matching packets if action is set to ALLOW. This next hop should be directly reachable from the appliance.",
			},
			"nexthopval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The Next Hop IPv6 address.",
			},
			"nexthopvlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "VLAN number to be used for link local nexthop .",
			},
			"ownergroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The owner node group in a Cluster for this pbr rule. If owner node group is not specified then the pbr rule is treated as Striped pbr rule.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority of the PBR6, which determines the order in which it is evaluated relative to the other PBR6s. If you do not specify priorities while creating PBR6s, the PBR6s are evaluated in the order in which they are created.",
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol, identified by protocol name, to match against the protocol of an outgoing IPv6 packet.",
			},
			"protocolnumber": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol, identified by protocol number, to match against the protocol of an outgoing IPv6 packet.",
			},
			"srcipop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Either the equals (=) or does not equal (!=) logical operator.",
			},
			"srcipv6": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the source IP address of an outgoing IPv6 packet. In the command line interface, separate the range with a hyphen.",
			},
			"srcipv6val": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or range of IP addresses to match against the source IP address of an outgoing IPv6 packet. In the command line interface, separate the range with a hyphen.",
			},
			"srcmac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "MAC address to match against the source MAC address of an outgoing IPv6 packet.",
			},
			"srcmacmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Used to define range of Source MAC address. It takes string of 0 and 1, 0s are for exact match and 1s for wildcard. For matching first 3 bytes of MAC address, srcMacMask value \"000000111111\".",
			},
			"srcport": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number or range of port numbers to match against the source port number of an outgoing IPv6 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.",
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
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the PBR6. After you apply the PBR6s, the Citrix ADC compares outgoing packets to the enabled PBR6s.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the VLAN. The Citrix ADC compares the PBR6 only to the outgoing packets on the specified VLAN. If you do not specify an interface ID, the appliance compares the PBR6 to the outgoing packets on all VLANs.",
			},
			"vxlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the VXLAN. The Citrix ADC compares the PBR6 only to the outgoing packets on the specified VXLAN. If you do not specify an interface ID, the appliance compares the PBR6 to the outgoing packets on all VXLANs.",
			},
			"vxlanvlanmap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The vlan to vxlan mapping to be applied for incoming packets over this pbr tunnel.",
			},
		},
	}
}

// nspbr6GetThePayloadFromthePlan builds the full Create payload (including state, matching SDK v2 Create).
func nspbr6GetThePayloadFromthePlan(ctx context.Context, data *Nspbr6ResourceModel) ns.Nspbr6 {
	tflog.Debug(ctx, "In nspbr6GetThePayloadFromthePlan Function")

	nspbr6 := nspbr6GetTheUpdatablePayloadFromThePlan(ctx, data)
	if !data.State.IsNull() && !data.State.IsUnknown() {
		nspbr6.State = data.State.ValueString()
	}
	return nspbr6
}

// nspbr6GetTheUpdatablePayloadFromThePlan builds the payload for AddResource/UpdateResource,
// excluding the state attribute (state is driven by the enable/disable action, matching SDK v2).
func nspbr6GetTheUpdatablePayloadFromThePlan(ctx context.Context, data *Nspbr6ResourceModel) ns.Nspbr6 {
	tflog.Debug(ctx, "In nspbr6GetTheUpdatablePayloadFromThePlan Function")

	// Create API request body from the model
	nspbr6 := ns.Nspbr6{}
	if !data.Interface.IsNull() && !data.Interface.IsUnknown() {
		nspbr6.Interface = data.Interface.ValueString()
	}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		nspbr6.Action = data.Action.ValueString()
	}
	if !data.Destipop.IsNull() && !data.Destipop.IsUnknown() {
		nspbr6.Destipop = data.Destipop.ValueString()
	}
	if !data.Destipv6.IsNull() && !data.Destipv6.IsUnknown() {
		nspbr6.Destipv6 = data.Destipv6.ValueBool()
	}
	if !data.Destipv6val.IsNull() && !data.Destipv6val.IsUnknown() {
		nspbr6.Destipv6val = data.Destipv6val.ValueString()
	}
	if !data.Destport.IsNull() && !data.Destport.IsUnknown() {
		nspbr6.Destport = data.Destport.ValueBool()
	}
	if !data.Destportop.IsNull() && !data.Destportop.IsUnknown() {
		nspbr6.Destportop = data.Destportop.ValueString()
	}
	if !data.Destportval.IsNull() && !data.Destportval.IsUnknown() {
		nspbr6.Destportval = data.Destportval.ValueString()
	}
	if !data.Detail.IsNull() && !data.Detail.IsUnknown() {
		nspbr6.Detail = data.Detail.ValueBool()
	}
	if !data.Iptunnel.IsNull() && !data.Iptunnel.IsUnknown() {
		nspbr6.Iptunnel = data.Iptunnel.ValueString()
	}
	if !data.Monitor.IsNull() && !data.Monitor.IsUnknown() {
		nspbr6.Monitor = data.Monitor.ValueString()
	}
	if !data.Msr.IsNull() && !data.Msr.IsUnknown() {
		nspbr6.Msr = data.Msr.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		nspbr6.Name = data.Name.ValueString()
	}
	if !data.Nexthop.IsNull() && !data.Nexthop.IsUnknown() {
		nspbr6.Nexthop = data.Nexthop.ValueBool()
	}
	if !data.Nexthopval.IsNull() && !data.Nexthopval.IsUnknown() {
		nspbr6.Nexthopval = data.Nexthopval.ValueString()
	}
	if !data.Nexthopvlan.IsNull() && !data.Nexthopvlan.IsUnknown() {
		nspbr6.Nexthopvlan = utils.IntPtr(int(data.Nexthopvlan.ValueInt64()))
	}
	if !data.Ownergroup.IsNull() && !data.Ownergroup.IsUnknown() {
		nspbr6.Ownergroup = data.Ownergroup.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		nspbr6.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Protocol.IsNull() && !data.Protocol.IsUnknown() {
		nspbr6.Protocol = data.Protocol.ValueString()
	}
	if !data.Protocolnumber.IsNull() && !data.Protocolnumber.IsUnknown() {
		nspbr6.Protocolnumber = utils.IntPtr(int(data.Protocolnumber.ValueInt64()))
	}
	if !data.Srcipop.IsNull() && !data.Srcipop.IsUnknown() {
		nspbr6.Srcipop = data.Srcipop.ValueString()
	}
	if !data.Srcipv6.IsNull() && !data.Srcipv6.IsUnknown() {
		nspbr6.Srcipv6 = data.Srcipv6.ValueBool()
	}
	if !data.Srcipv6val.IsNull() && !data.Srcipv6val.IsUnknown() {
		nspbr6.Srcipv6val = data.Srcipv6val.ValueString()
	}
	if !data.Srcmac.IsNull() && !data.Srcmac.IsUnknown() {
		nspbr6.Srcmac = data.Srcmac.ValueString()
	}
	if !data.Srcmacmask.IsNull() && !data.Srcmacmask.IsUnknown() {
		nspbr6.Srcmacmask = data.Srcmacmask.ValueString()
	}
	if !data.Srcport.IsNull() && !data.Srcport.IsUnknown() {
		nspbr6.Srcport = data.Srcport.ValueBool()
	}
	if !data.Srcportop.IsNull() && !data.Srcportop.IsUnknown() {
		nspbr6.Srcportop = data.Srcportop.ValueString()
	}
	if !data.Srcportval.IsNull() && !data.Srcportval.IsUnknown() {
		nspbr6.Srcportval = data.Srcportval.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		nspbr6.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		nspbr6.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}
	if !data.Vxlan.IsNull() && !data.Vxlan.IsUnknown() {
		nspbr6.Vxlan = utils.IntPtr(int(data.Vxlan.ValueInt64()))
	}
	if !data.Vxlanvlanmap.IsNull() && !data.Vxlanvlanmap.IsUnknown() {
		nspbr6.Vxlanvlanmap = data.Vxlanvlanmap.ValueString()
	}

	return nspbr6
}

// nspbr6GetTheChangedPayloadFromThePlan builds an UpdateResource (SET) payload that contains
// ONLY the attributes that actually changed between the prior state and the plan, mirroring
// SDK v2's per-attribute d.HasChange() gating. This is required because NITRO rejects a PBR6
// SET that carries certain unchanged fields; in particular a SET that includes "iptunnel"
// fails with errorcode 383 ("SET operations on PBR for changing iptunnel not allowed. Remove
// and add PBR with new iptunnel"), even when the value is unchanged. The "state" attribute is
// excluded here — enable/disable is driven separately via the NITRO action.
func nspbr6GetTheChangedPayloadFromThePlan(ctx context.Context, data *Nspbr6ResourceModel, state *Nspbr6ResourceModel) ns.Nspbr6 {
	tflog.Debug(ctx, "In nspbr6GetTheChangedPayloadFromThePlan Function")

	// Name is the key and is always required in the SET body.
	nspbr6 := ns.Nspbr6{
		Name: data.Name.ValueString(),
	}
	if !data.Interface.Equal(state.Interface) && !data.Interface.IsNull() && !data.Interface.IsUnknown() {
		nspbr6.Interface = data.Interface.ValueString()
	}
	if !data.Action.Equal(state.Action) && !data.Action.IsNull() && !data.Action.IsUnknown() {
		nspbr6.Action = data.Action.ValueString()
	}
	if !data.Destipop.Equal(state.Destipop) && !data.Destipop.IsNull() && !data.Destipop.IsUnknown() {
		nspbr6.Destipop = data.Destipop.ValueString()
	}
	if !data.Destipv6.Equal(state.Destipv6) && !data.Destipv6.IsNull() && !data.Destipv6.IsUnknown() {
		nspbr6.Destipv6 = data.Destipv6.ValueBool()
	}
	if !data.Destipv6val.Equal(state.Destipv6val) && !data.Destipv6val.IsNull() && !data.Destipv6val.IsUnknown() {
		nspbr6.Destipv6val = data.Destipv6val.ValueString()
	}
	if !data.Destport.Equal(state.Destport) && !data.Destport.IsNull() && !data.Destport.IsUnknown() {
		nspbr6.Destport = data.Destport.ValueBool()
	}
	if !data.Destportop.Equal(state.Destportop) && !data.Destportop.IsNull() && !data.Destportop.IsUnknown() {
		nspbr6.Destportop = data.Destportop.ValueString()
	}
	if !data.Destportval.Equal(state.Destportval) && !data.Destportval.IsNull() && !data.Destportval.IsUnknown() {
		nspbr6.Destportval = data.Destportval.ValueString()
	}
	if !data.Detail.Equal(state.Detail) && !data.Detail.IsNull() && !data.Detail.IsUnknown() {
		nspbr6.Detail = data.Detail.ValueBool()
	}
	if !data.Iptunnel.Equal(state.Iptunnel) && !data.Iptunnel.IsNull() && !data.Iptunnel.IsUnknown() {
		nspbr6.Iptunnel = data.Iptunnel.ValueString()
	}
	if !data.Monitor.Equal(state.Monitor) && !data.Monitor.IsNull() && !data.Monitor.IsUnknown() {
		nspbr6.Monitor = data.Monitor.ValueString()
	}
	if !data.Msr.Equal(state.Msr) && !data.Msr.IsNull() && !data.Msr.IsUnknown() {
		nspbr6.Msr = data.Msr.ValueString()
	}
	if !data.Nexthop.Equal(state.Nexthop) && !data.Nexthop.IsNull() && !data.Nexthop.IsUnknown() {
		nspbr6.Nexthop = data.Nexthop.ValueBool()
	}
	if !data.Nexthopval.Equal(state.Nexthopval) && !data.Nexthopval.IsNull() && !data.Nexthopval.IsUnknown() {
		nspbr6.Nexthopval = data.Nexthopval.ValueString()
	}
	if !data.Nexthopvlan.Equal(state.Nexthopvlan) && !data.Nexthopvlan.IsNull() && !data.Nexthopvlan.IsUnknown() {
		nspbr6.Nexthopvlan = utils.IntPtr(int(data.Nexthopvlan.ValueInt64()))
	}
	if !data.Ownergroup.Equal(state.Ownergroup) && !data.Ownergroup.IsNull() && !data.Ownergroup.IsUnknown() {
		nspbr6.Ownergroup = data.Ownergroup.ValueString()
	}
	if !data.Priority.Equal(state.Priority) && !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		nspbr6.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Protocol.Equal(state.Protocol) && !data.Protocol.IsNull() && !data.Protocol.IsUnknown() {
		nspbr6.Protocol = data.Protocol.ValueString()
	}
	if !data.Protocolnumber.Equal(state.Protocolnumber) && !data.Protocolnumber.IsNull() && !data.Protocolnumber.IsUnknown() {
		nspbr6.Protocolnumber = utils.IntPtr(int(data.Protocolnumber.ValueInt64()))
	}
	if !data.Srcipop.Equal(state.Srcipop) && !data.Srcipop.IsNull() && !data.Srcipop.IsUnknown() {
		nspbr6.Srcipop = data.Srcipop.ValueString()
	}
	if !data.Srcipv6.Equal(state.Srcipv6) && !data.Srcipv6.IsNull() && !data.Srcipv6.IsUnknown() {
		nspbr6.Srcipv6 = data.Srcipv6.ValueBool()
	}
	if !data.Srcipv6val.Equal(state.Srcipv6val) && !data.Srcipv6val.IsNull() && !data.Srcipv6val.IsUnknown() {
		nspbr6.Srcipv6val = data.Srcipv6val.ValueString()
	}
	if !data.Srcmac.Equal(state.Srcmac) && !data.Srcmac.IsNull() && !data.Srcmac.IsUnknown() {
		nspbr6.Srcmac = data.Srcmac.ValueString()
	}
	if !data.Srcmacmask.Equal(state.Srcmacmask) && !data.Srcmacmask.IsNull() && !data.Srcmacmask.IsUnknown() {
		nspbr6.Srcmacmask = data.Srcmacmask.ValueString()
	}
	if !data.Srcport.Equal(state.Srcport) && !data.Srcport.IsNull() && !data.Srcport.IsUnknown() {
		nspbr6.Srcport = data.Srcport.ValueBool()
	}
	if !data.Srcportop.Equal(state.Srcportop) && !data.Srcportop.IsNull() && !data.Srcportop.IsUnknown() {
		nspbr6.Srcportop = data.Srcportop.ValueString()
	}
	if !data.Srcportval.Equal(state.Srcportval) && !data.Srcportval.IsNull() && !data.Srcportval.IsUnknown() {
		nspbr6.Srcportval = data.Srcportval.ValueString()
	}
	if !data.Td.Equal(state.Td) && !data.Td.IsNull() && !data.Td.IsUnknown() {
		nspbr6.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Vlan.Equal(state.Vlan) && !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		nspbr6.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}
	if !data.Vxlan.Equal(state.Vxlan) && !data.Vxlan.IsNull() && !data.Vxlan.IsUnknown() {
		nspbr6.Vxlan = utils.IntPtr(int(data.Vxlan.ValueInt64()))
	}
	if !data.Vxlanvlanmap.Equal(state.Vxlanvlanmap) && !data.Vxlanvlanmap.IsNull() && !data.Vxlanvlanmap.IsUnknown() {
		nspbr6.Vxlanvlanmap = data.Vxlanvlanmap.ValueString()
	}

	return nspbr6
}

// nspbr6HasNonStateChange reports whether any updateable attribute other than "state" changed
// between the prior state and the plan (state is handled via the enable/disable action).
func nspbr6HasNonStateChange(data *Nspbr6ResourceModel, state *Nspbr6ResourceModel) bool {
	return !data.Interface.Equal(state.Interface) ||
		!data.Action.Equal(state.Action) ||
		!data.Destipop.Equal(state.Destipop) ||
		!data.Destipv6.Equal(state.Destipv6) ||
		!data.Destipv6val.Equal(state.Destipv6val) ||
		!data.Destport.Equal(state.Destport) ||
		!data.Destportop.Equal(state.Destportop) ||
		!data.Destportval.Equal(state.Destportval) ||
		!data.Detail.Equal(state.Detail) ||
		!data.Iptunnel.Equal(state.Iptunnel) ||
		!data.Monitor.Equal(state.Monitor) ||
		!data.Msr.Equal(state.Msr) ||
		!data.Nexthop.Equal(state.Nexthop) ||
		!data.Nexthopval.Equal(state.Nexthopval) ||
		!data.Nexthopvlan.Equal(state.Nexthopvlan) ||
		!data.Ownergroup.Equal(state.Ownergroup) ||
		!data.Priority.Equal(state.Priority) ||
		!data.Protocol.Equal(state.Protocol) ||
		!data.Protocolnumber.Equal(state.Protocolnumber) ||
		!data.Srcipop.Equal(state.Srcipop) ||
		!data.Srcipv6.Equal(state.Srcipv6) ||
		!data.Srcipv6val.Equal(state.Srcipv6val) ||
		!data.Srcmac.Equal(state.Srcmac) ||
		!data.Srcmacmask.Equal(state.Srcmacmask) ||
		!data.Srcport.Equal(state.Srcport) ||
		!data.Srcportop.Equal(state.Srcportop) ||
		!data.Srcportval.Equal(state.Srcportval) ||
		!data.Td.Equal(state.Td) ||
		!data.Vlan.Equal(state.Vlan) ||
		!data.Vxlan.Equal(state.Vxlan) ||
		!data.Vxlanvlanmap.Equal(state.Vxlanvlanmap)
}

func nspbr6SetAttrFromGet(ctx context.Context, data *Nspbr6ResourceModel, getResponseData map[string]interface{}) *Nspbr6ResourceModel {
	tflog.Debug(ctx, "In nspbr6SetAttrFromGet Function")

	// Convert API response to model.
	// The else-branches only null a value when it is Unknown (i.e. an unset Optional+Computed
	// attribute), so a known/configured value that NITRO omits from GET (omit-on-default) is
	// preserved instead of being clobbered.

	// The NITRO write tag for interface is "Interface" (capitalised); GET responses have been
	// observed to use either case, so both keys are checked.
	if val, ok := getResponseData["Interface"]; ok && val != nil {
		data.Interface = types.StringValue(val.(string))
	} else if val, ok := getResponseData["interface"]; ok && val != nil {
		data.Interface = types.StringValue(val.(string))
	} else if data.Interface.IsUnknown() {
		data.Interface = types.StringNull()
	}
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else if data.Action.IsUnknown() {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["destipop"]; ok && val != nil {
		data.Destipop = types.StringValue(val.(string))
	} else if data.Destipop.IsUnknown() {
		data.Destipop = types.StringNull()
	}
	if val, ok := getResponseData["destipv6"]; ok && val != nil {
		data.Destipv6 = types.BoolValue(val.(bool))
	} else if data.Destipv6.IsUnknown() {
		data.Destipv6 = types.BoolNull()
	}
	if val, ok := getResponseData["destipv6val"]; ok && val != nil {
		data.Destipv6val = types.StringValue(val.(string))
	} else if data.Destipv6val.IsUnknown() {
		data.Destipv6val = types.StringNull()
	}
	if val, ok := getResponseData["destport"]; ok && val != nil {
		data.Destport = types.BoolValue(val.(bool))
	} else if data.Destport.IsUnknown() {
		data.Destport = types.BoolNull()
	}
	if val, ok := getResponseData["destportop"]; ok && val != nil {
		data.Destportop = types.StringValue(val.(string))
	} else if data.Destportop.IsUnknown() {
		data.Destportop = types.StringNull()
	}
	if val, ok := getResponseData["destportval"]; ok && val != nil {
		data.Destportval = types.StringValue(val.(string))
	} else if data.Destportval.IsUnknown() {
		data.Destportval = types.StringNull()
	}
	if val, ok := getResponseData["detail"]; ok && val != nil {
		data.Detail = types.BoolValue(val.(bool))
	} else if data.Detail.IsUnknown() {
		data.Detail = types.BoolNull()
	}
	if val, ok := getResponseData["iptunnel"]; ok && val != nil {
		data.Iptunnel = types.StringValue(val.(string))
	} else if data.Iptunnel.IsUnknown() {
		data.Iptunnel = types.StringNull()
	}
	if val, ok := getResponseData["monitor"]; ok && val != nil {
		data.Monitor = types.StringValue(val.(string))
	} else if data.Monitor.IsUnknown() {
		data.Monitor = types.StringNull()
	}
	if val, ok := getResponseData["msr"]; ok && val != nil {
		data.Msr = types.StringValue(val.(string))
	} else if data.Msr.IsUnknown() {
		data.Msr = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["nexthop"]; ok && val != nil {
		data.Nexthop = types.BoolValue(val.(bool))
	} else if data.Nexthop.IsUnknown() {
		data.Nexthop = types.BoolNull()
	}
	if val, ok := getResponseData["nexthopval"]; ok && val != nil {
		data.Nexthopval = types.StringValue(val.(string))
	} else if data.Nexthopval.IsUnknown() {
		data.Nexthopval = types.StringNull()
	}
	if val, ok := getResponseData["nexthopvlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nexthopvlan = types.Int64Value(intVal)
		}
	} else if data.Nexthopvlan.IsUnknown() {
		data.Nexthopvlan = types.Int64Null()
	}
	if val, ok := getResponseData["ownergroup"]; ok && val != nil {
		data.Ownergroup = types.StringValue(val.(string))
	} else if data.Ownergroup.IsUnknown() {
		data.Ownergroup = types.StringNull()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else if data.Priority.IsUnknown() {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["protocol"]; ok && val != nil {
		data.Protocol = types.StringValue(val.(string))
	} else if data.Protocol.IsUnknown() {
		data.Protocol = types.StringNull()
	}
	if val, ok := getResponseData["protocolnumber"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Protocolnumber = types.Int64Value(intVal)
		}
	} else if data.Protocolnumber.IsUnknown() {
		data.Protocolnumber = types.Int64Null()
	}
	if val, ok := getResponseData["srcipop"]; ok && val != nil {
		data.Srcipop = types.StringValue(val.(string))
	} else if data.Srcipop.IsUnknown() {
		data.Srcipop = types.StringNull()
	}
	if val, ok := getResponseData["srcipv6"]; ok && val != nil {
		data.Srcipv6 = types.BoolValue(val.(bool))
	} else if data.Srcipv6.IsUnknown() {
		data.Srcipv6 = types.BoolNull()
	}
	if val, ok := getResponseData["srcipv6val"]; ok && val != nil {
		data.Srcipv6val = types.StringValue(val.(string))
	} else if data.Srcipv6val.IsUnknown() {
		data.Srcipv6val = types.StringNull()
	}
	if val, ok := getResponseData["srcmac"]; ok && val != nil {
		data.Srcmac = types.StringValue(val.(string))
	} else if data.Srcmac.IsUnknown() {
		data.Srcmac = types.StringNull()
	}
	if val, ok := getResponseData["srcmacmask"]; ok && val != nil {
		data.Srcmacmask = types.StringValue(val.(string))
	} else if data.Srcmacmask.IsUnknown() {
		data.Srcmacmask = types.StringNull()
	}
	if val, ok := getResponseData["srcport"]; ok && val != nil {
		data.Srcport = types.BoolValue(val.(bool))
	} else if data.Srcport.IsUnknown() {
		data.Srcport = types.BoolNull()
	}
	if val, ok := getResponseData["srcportop"]; ok && val != nil {
		data.Srcportop = types.StringValue(val.(string))
	} else if data.Srcportop.IsUnknown() {
		data.Srcportop = types.StringNull()
	}
	if val, ok := getResponseData["srcportval"]; ok && val != nil {
		data.Srcportval = types.StringValue(val.(string))
	} else if data.Srcportval.IsUnknown() {
		data.Srcportval = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else if data.State.IsUnknown() {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else if data.Td.IsUnknown() {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	} else if data.Vlan.IsUnknown() {
		data.Vlan = types.Int64Null()
	}
	if val, ok := getResponseData["vxlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vxlan = types.Int64Value(intVal)
		}
	} else if data.Vxlan.IsUnknown() {
		data.Vxlan = types.Int64Null()
	}
	if val, ok := getResponseData["vxlanvlanmap"]; ok && val != nil {
		data.Vxlanvlanmap = types.StringValue(val.(string))
	} else if data.Vxlanvlanmap.IsUnknown() {
		data.Vxlanvlanmap = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
