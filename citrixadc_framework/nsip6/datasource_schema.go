package nsip6

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Nsip6DataSourceModel is the data-source-specific model, decoupled from
// Nsip6ResourceModel.
//
// A data source is a pure read surface, so it exposes the read/write attributes
// (as Computed outputs) AND the read-only runtime/status attributes the resource
// deliberately omits (iptype, curstate, VIP vserver counters, systemtype, ...),
// all of which are populated only from a GET.
type Nsip6DataSourceModel struct {
	Id                          types.String `tfsdk:"id"`
	Advertiseondefaultpartition types.String `tfsdk:"advertiseondefaultpartition"`
	Decrementhoplimit           types.String `tfsdk:"decrementhoplimit"`
	Dynamicrouting              types.String `tfsdk:"dynamicrouting"`
	Ftp                         types.String `tfsdk:"ftp"`
	Gui                         types.String `tfsdk:"gui"`
	Hostroute                   types.String `tfsdk:"hostroute"`
	Icmp                        types.String `tfsdk:"icmp"`
	Icmpresponse                types.String `tfsdk:"icmpresponse"`
	Ip6hostrtgw                 types.String `tfsdk:"ip6hostrtgw"`
	Ipv6address                 types.String `tfsdk:"ipv6address"`
	Map                         types.String `tfsdk:"map"`
	Metric                      types.Int64  `tfsdk:"metric"`
	Mgmtaccess                  types.String `tfsdk:"mgmtaccess"`
	Mptcpadvertise              types.String `tfsdk:"mptcpadvertise"`
	Nd                          types.String `tfsdk:"nd"`
	Ndowner                     types.Int64  `tfsdk:"ndowner"`
	Networkroute                types.String `tfsdk:"networkroute"`
	Ospf6lsatype                types.String `tfsdk:"ospf6lsatype"`
	Ospfarea                    types.Int64  `tfsdk:"ospfarea"`
	Ownerdownresponse           types.String `tfsdk:"ownerdownresponse"`
	Ownernode                   types.Int64  `tfsdk:"ownernode"`
	Restrictaccess              types.String `tfsdk:"restrictaccess"`
	Scope                       types.String `tfsdk:"scope"`
	Snmp                        types.String `tfsdk:"snmp"`
	Ssh                         types.String `tfsdk:"ssh"`
	State                       types.String `tfsdk:"state"`
	Tag                         types.Int64  `tfsdk:"tag"`
	Td                          types.Int64  `tfsdk:"td"`
	Telnet                      types.String `tfsdk:"telnet"`
	Type                        types.String `tfsdk:"type"`
	Vlan                        types.Int64  `tfsdk:"vlan"`
	Vrid6                       types.Int64  `tfsdk:"vrid6"`
	Vserver                     types.String `tfsdk:"vserver"`
	Vserverrhilevel             types.String `tfsdk:"vserverrhilevel"`

	// Read-only (GET-only) runtime/status attributes from the NITRO read-only
	// set. Never settable; populated from GET.
	Iptype             types.List   `tfsdk:"iptype"`
	Curstate           types.String `tfsdk:"curstate"`
	Viprtadv2bsd       types.Bool   `tfsdk:"viprtadv2bsd"`
	Vipvsercount       types.Int64  `tfsdk:"vipvsercount"`
	Vipvserdowncount   types.Int64  `tfsdk:"vipvserdowncount"`
	Systemtype         types.String `tfsdk:"systemtype"`
	Operationalndowner types.Int64  `tfsdk:"operationalndowner"`
}

func Nsip6DataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"advertiseondefaultpartition": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Advertise VIPs from Shared VLAN on Default Partition",
			},
			"decrementhoplimit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Decrement Hop Limit by 1 when ENABLED.This setting is applicable only for UDP traffic.",
			},
			"dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow dynamic routing on this IP address. Specific to Subnet IPv6 (SNIP6) address.",
			},
			"ftp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow File Transfer Protocol (FTP) access to this IP address.",
			},
			"gui": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow graphical user interface (GUI) access to this IP address.",
			},
			"hostroute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to push the VIP6 to ZebOS routing table for Kernel route redistribution through dynamic routing protocols.",
			},
			"icmp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Respond to ICMP requests for this IP address.",
			},
			"icmpresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Respond to ICMPv6 requests for a Virtual IP (VIP) address on the basis of the states of the virtual servers associated with that VIP",
			},
			"ip6hostrtgw": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv6 address of the gateway for the route. If Gateway is not set, VIP uses :: as the gateway.",
			},
			"ipv6address": schema.StringAttribute{
				Required:    true,
				Description: "IPv6 address to create on the Citrix ADC.",
			},
			"map": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mapped IPV4 address for the IPV6 address.",
			},
			"metric": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value to add to or subtract from the cost of the route advertised for the VIP6 address.",
			},
			"mgmtaccess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow access to management applications on this IP address.",
			},
			"mptcpadvertise": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If enabled, this IP will be advertised by Citrix ADC to MPTCP enabled clients as part of ADD_ADDR option.",
			},
			"nd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Respond to Neighbor Discovery (ND) requests for this IP address.",
			},
			"ndowner": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NdOwner in Cluster for VIPS and Striped SNIPS",
			},
			"networkroute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to push the SNIP6 subnet to ZebOS routing table for Kernel route redistribution through dynamic routing protocol.",
			},
			"ospf6lsatype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of LSAs to be used by the IPv6 OSPF protocol, running on the Citrix ADC, for advertising the route for the VIP6 address.",
			},
			"ospfarea": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the area in which the Intra-Area-Prefix LSAs are to be advertised for the VIP6 address by the IPv6 OSPF protocol running on the Citrix ADC. When ospfArea is not set, VIP6 is advertised on all areas.",
			},
			"ownerdownresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "in cluster system, if the owner node is down, whether should it respond to icmp/arp",
			},
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the cluster node for which you are adding the IP address. Must be used if you want the IP address to be active only on the specific node. Can be configured only through the cluster IP address. Cannot be changed after the IP address is created.",
			},
			"restrictaccess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Block access to nonmanagement applications on this IP address. This option is applicable forMIP6s, SNIP6s, and NSIP6s, and is disabled by default. Nonmanagement applications can run on the underlying Citrix ADC Free BSD operating system.",
			},
			"scope": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Scope of the IPv6 address to be created. Cannot be changed after the IP address is created.",
			},
			"snmp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Simple Network Management Protocol (SNMP) access to this IP address.",
			},
			"ssh": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow secure Shell (SSH) access to this IP address.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the IP address.",
			},
			"tag": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Tag value for the network/host route associated with this IP.",
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"telnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Telnet access to this IP address.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of IP address to be created on the Citrix ADC. Cannot be changed after the IP address is created.",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The VLAN number.",
			},
			"vrid6": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "A positive integer that uniquely identifies a VMAC address for binding to this VIP address. This binding is used to set up Citrix ADCs in an active-active configuration using VRRP.",
			},
			"vserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the state of all the virtual servers associated with this VIP6 address.",
			},
			"vserverrhilevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Advertise or do not advertise the route for the Virtual IP (VIP6) address on the basis of the state of the virtual servers associated with that VIP6.\n* NONE - Advertise the route for the VIP6 address, irrespective of the state of the virtual servers associated with the address.\n* ONE VSERVER - Advertise the route for the VIP6 address if at least one of the associated virtual servers is in UP state.\n* ALL VSERVER - Advertise the route for the VIP6 address if all of the associated virtual servers are in UP state.\n* VSVR_CNTRLD.   Advertise the route for the VIP address according to the  RHIstate (RHI STATE) parameter setting on all the associated virtual servers of the VIP address along with their states.\n\nWhen Vserver RHI Level (RHI) parameter is set to VSVR_CNTRLD, the following are different RHI behaviors for the VIP address on the basis of RHIstate (RHI STATE) settings on the virtual servers associated with the VIP address:\n * If you set RHI STATE to PASSIVE on all virtual servers, the Citrix ADC always advertises the route for the VIP address.\n * If you set RHI STATE to ACTIVE on all virtual servers, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers is in UP state.\n *If you set RHI STATE to ACTIVE on some and PASSIVE on others, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers, whose RHI STATE set to ACTIVE, is in UP state.",
			},

			// Read-only (GET-only) runtime/status attributes surfaced by the data
			// source (these are intentionally NOT modeled on the resource). All Computed.
			"iptype": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "The type of the IPv6 address. Possible values: [ NSIP, VIP, SNIP, GSLBsiteIP, ADNSsvcIP, RADIUSListenersvcIP, CLIP ]",
			},
			"curstate": schema.StringAttribute{
				Computed:    true,
				Description: "Current state of this IP. Possible values: [ DISABLED, ENABLED ]",
			},
			"viprtadv2bsd": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether this route is advertised to FreeBSD.",
			},
			"vipvsercount": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of vservers bound to this VIP.",
			},
			"vipvserdowncount": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of vservers bound to this VIP, which are down.",
			},
			"systemtype": schema.StringAttribute{
				Computed:    true,
				Description: "The type of the System. Used for display purpose. Possible values: [ Stand-alone, HA, Cluster ]",
			},
			"operationalndowner": schema.Int64Attribute{
				Computed:    true,
				Description: "Operational ND6 Owner.",
			},
		},
	}
}

// nsip6DataSourceSetAttrFromGet projects a NITRO nsip6 GET response onto the
// data-source model. There is no prior plan/state to preserve, so the shared
// utils.MapGet* helpers fill each attribute (or leave it Null when the GET omits
// it). The ID is the plain ipv6address value, matching the resource ID scheme.
func nsip6DataSourceSetAttrFromGet(ctx context.Context, data *Nsip6DataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nsip6DataSourceSetAttrFromGet Function")

	if v, ok := g["ipv6address"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Ipv6address = types.StringValue(utils.AnyToString(v))
	}

	data.Advertiseondefaultpartition = utils.MapGetString(g, "advertiseondefaultpartition")
	data.Decrementhoplimit = utils.MapGetString(g, "decrementhoplimit")
	data.Dynamicrouting = utils.MapGetString(g, "dynamicrouting")
	data.Ftp = utils.MapGetString(g, "ftp")
	data.Gui = utils.MapGetString(g, "gui")
	data.Hostroute = utils.MapGetString(g, "hostroute")
	data.Icmp = utils.MapGetString(g, "icmp")
	data.Icmpresponse = utils.MapGetString(g, "icmpresponse")
	data.Ip6hostrtgw = utils.MapGetString(g, "ip6hostrtgw")
	data.Map = utils.MapGetString(g, "map")
	data.Metric = utils.MapGetInt64(g, "metric")
	data.Mgmtaccess = utils.MapGetString(g, "mgmtaccess")
	data.Mptcpadvertise = utils.MapGetString(g, "mptcpadvertise")
	data.Nd = utils.MapGetString(g, "nd")
	data.Ndowner = utils.MapGetInt64(g, "ndowner")
	data.Networkroute = utils.MapGetString(g, "networkroute")
	data.Ospf6lsatype = utils.MapGetString(g, "ospf6lsatype")
	data.Ospfarea = utils.MapGetInt64(g, "ospfarea")
	data.Ownerdownresponse = utils.MapGetString(g, "ownerdownresponse")
	data.Ownernode = utils.MapGetInt64(g, "ownernode")
	data.Restrictaccess = utils.MapGetString(g, "restrictaccess")
	data.Scope = utils.MapGetString(g, "scope")
	data.Snmp = utils.MapGetString(g, "snmp")
	data.Ssh = utils.MapGetString(g, "ssh")
	data.State = utils.MapGetString(g, "state")
	data.Tag = utils.MapGetInt64(g, "tag")
	// td is a config-supplied key; NITRO omits it for the default traffic
	// domain (0), so preserve the configured value instead of nulling it.
	if tdv, tdok := g["td"]; tdok && tdv != nil {
		if iv, err := utils.ConvertToInt64(tdv); err == nil {
			data.Td = types.Int64Value(iv)
		}
	}
	data.Telnet = utils.MapGetString(g, "telnet")
	data.Type = utils.MapGetString(g, "type")
	data.Vlan = utils.MapGetInt64(g, "vlan")
	data.Vrid6 = utils.MapGetInt64(g, "vrid6")
	data.Vserver = utils.MapGetString(g, "vserver")
	data.Vserverrhilevel = utils.MapGetString(g, "vserverrhilevel")

	// Read-only runtime/status attributes.
	data.Iptype = utils.MapGetStringList(g, "iptype")
	data.Curstate = utils.MapGetString(g, "curstate")
	data.Viprtadv2bsd = utils.MapGetBool(g, "viprtadv2bsd")
	data.Vipvsercount = utils.MapGetInt64(g, "vipvsercount")
	data.Vipvserdowncount = utils.MapGetInt64(g, "vipvserdowncount")
	data.Systemtype = utils.MapGetString(g, "systemtype")
	data.Operationalndowner = utils.MapGetInt64(g, "operationalndowner")
}
