package nsip6

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Nsip6ResourceModel describes the resource data model.
type Nsip6ResourceModel struct {
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
}

func (r *Nsip6Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsip6 resource.",
			},
			// SDK v2: Optional+Computed, no default (value read from ADC).
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
			// SDK v2: Required + ForceNew.
			"ipv6address": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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
			// SDK v2: Optional+Computed+ForceNew. UseStateForUnknown keeps the value
			// stable across refreshes; RequiresReplaceIfConfigured reproduces ForceNew
			// only when the user actually configured the attribute.
			"ownernode": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "ID of the cluster node for which you are adding the IP address. Must be used if you want the IP address to be active only on the specific node. Can be configured only through the cluster IP address. Cannot be changed after the IP address is created.",
			},
			"restrictaccess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Block access to nonmanagement applications on this IP address. This option is applicable forMIP6s, SNIP6s, and NSIP6s, and is disabled by default. Nonmanagement applications can run on the underlying Citrix ADC Free BSD operating system.",
			},
			// SDK v2: Optional+Computed+ForceNew.
			"scope": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
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
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"telnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Telnet access to this IP address.",
			},
			// SDK v2: Optional+Computed+ForceNew. The read-back value is derived from
			// the NITRO read-only "iptype" array (see nsip6SetAttrFromGet).
			"type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Type of IP address to be created on the Citrix ADC. Cannot be changed after the IP address is created.",
			},
			// SDK v2: Optional+Computed+ForceNew.
			"vlan": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
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
		},
	}
}

// nsip6IptypeToString extracts the IP type value from the NITRO read-only
// "iptype" field, which is returned as an array (e.g. ["VIP"]) rather than the
// request-side "type" field (mirrors the SDK v2 read of iptype[0]).
func nsip6IptypeToString(val interface{}) (string, bool) {
	switch v := val.(type) {
	case string:
		return v, true
	case []interface{}:
		if len(v) > 0 {
			if s, ok := v[0].(string); ok {
				return s, true
			}
		}
	}
	return "", false
}

// nsip6GetThePayloadFromtheConfig builds the full add/create payload from the
// plan. Skip unknown/null values so that Optional+Computed attributes the user
// omitted fall through to the appliance defaults.
func nsip6GetThePayloadFromtheConfig(ctx context.Context, data *Nsip6ResourceModel) ns.Nsip6 {
	tflog.Debug(ctx, "In nsip6GetThePayloadFromtheConfig Function")

	nsip6 := ns.Nsip6{}
	if !data.Advertiseondefaultpartition.IsNull() && !data.Advertiseondefaultpartition.IsUnknown() {
		nsip6.Advertiseondefaultpartition = data.Advertiseondefaultpartition.ValueString()
	}
	if !data.Decrementhoplimit.IsNull() && !data.Decrementhoplimit.IsUnknown() {
		nsip6.Decrementhoplimit = data.Decrementhoplimit.ValueString()
	}
	if !data.Dynamicrouting.IsNull() && !data.Dynamicrouting.IsUnknown() {
		nsip6.Dynamicrouting = data.Dynamicrouting.ValueString()
	}
	if !data.Ftp.IsNull() && !data.Ftp.IsUnknown() {
		nsip6.Ftp = data.Ftp.ValueString()
	}
	if !data.Gui.IsNull() && !data.Gui.IsUnknown() {
		nsip6.Gui = data.Gui.ValueString()
	}
	if !data.Hostroute.IsNull() && !data.Hostroute.IsUnknown() {
		nsip6.Hostroute = data.Hostroute.ValueString()
	}
	if !data.Icmp.IsNull() && !data.Icmp.IsUnknown() {
		nsip6.Icmp = data.Icmp.ValueString()
	}
	if !data.Icmpresponse.IsNull() && !data.Icmpresponse.IsUnknown() {
		nsip6.Icmpresponse = data.Icmpresponse.ValueString()
	}
	if !data.Ip6hostrtgw.IsNull() && !data.Ip6hostrtgw.IsUnknown() {
		nsip6.Ip6hostrtgw = data.Ip6hostrtgw.ValueString()
	}
	if !data.Ipv6address.IsNull() && !data.Ipv6address.IsUnknown() {
		nsip6.Ipv6address = data.Ipv6address.ValueString()
	}
	if !data.Map.IsNull() && !data.Map.IsUnknown() {
		nsip6.Map = data.Map.ValueString()
	}
	if !data.Metric.IsNull() && !data.Metric.IsUnknown() {
		nsip6.Metric = utils.IntPtr(int(data.Metric.ValueInt64()))
	}
	if !data.Mgmtaccess.IsNull() && !data.Mgmtaccess.IsUnknown() {
		nsip6.Mgmtaccess = data.Mgmtaccess.ValueString()
	}
	if !data.Mptcpadvertise.IsNull() && !data.Mptcpadvertise.IsUnknown() {
		nsip6.Mptcpadvertise = data.Mptcpadvertise.ValueString()
	}
	if !data.Nd.IsNull() && !data.Nd.IsUnknown() {
		nsip6.Nd = data.Nd.ValueString()
	}
	if !data.Ndowner.IsNull() && !data.Ndowner.IsUnknown() {
		nsip6.Ndowner = utils.IntPtr(int(data.Ndowner.ValueInt64()))
	}
	if !data.Networkroute.IsNull() && !data.Networkroute.IsUnknown() {
		nsip6.Networkroute = data.Networkroute.ValueString()
	}
	if !data.Ospf6lsatype.IsNull() && !data.Ospf6lsatype.IsUnknown() {
		nsip6.Ospf6lsatype = data.Ospf6lsatype.ValueString()
	}
	if !data.Ospfarea.IsNull() && !data.Ospfarea.IsUnknown() {
		nsip6.Ospfarea = utils.IntPtr(int(data.Ospfarea.ValueInt64()))
	}
	if !data.Ownerdownresponse.IsNull() && !data.Ownerdownresponse.IsUnknown() {
		nsip6.Ownerdownresponse = data.Ownerdownresponse.ValueString()
	}
	if !data.Ownernode.IsNull() && !data.Ownernode.IsUnknown() {
		nsip6.Ownernode = utils.IntPtr(int(data.Ownernode.ValueInt64()))
	}
	if !data.Restrictaccess.IsNull() && !data.Restrictaccess.IsUnknown() {
		nsip6.Restrictaccess = data.Restrictaccess.ValueString()
	}
	if !data.Scope.IsNull() && !data.Scope.IsUnknown() {
		nsip6.Scope = data.Scope.ValueString()
	}
	if !data.Snmp.IsNull() && !data.Snmp.IsUnknown() {
		nsip6.Snmp = data.Snmp.ValueString()
	}
	if !data.Ssh.IsNull() && !data.Ssh.IsUnknown() {
		nsip6.Ssh = data.Ssh.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		nsip6.State = data.State.ValueString()
	}
	if !data.Tag.IsNull() && !data.Tag.IsUnknown() {
		nsip6.Tag = utils.IntPtr(int(data.Tag.ValueInt64()))
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		nsip6.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Telnet.IsNull() && !data.Telnet.IsUnknown() {
		nsip6.Telnet = data.Telnet.ValueString()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		nsip6.Type = data.Type.ValueString()
	}
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		nsip6.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}
	if !data.Vrid6.IsNull() && !data.Vrid6.IsUnknown() {
		nsip6.Vrid6 = utils.IntPtr(int(data.Vrid6.ValueInt64()))
	}
	if !data.Vserver.IsNull() && !data.Vserver.IsUnknown() {
		nsip6.Vserver = data.Vserver.ValueString()
	}
	if !data.Vserverrhilevel.IsNull() && !data.Vserverrhilevel.IsUnknown() {
		nsip6.Vserverrhilevel = data.Vserverrhilevel.ValueString()
	}

	return nsip6
}

// nsip6SetAttrFromGet maps the GET response onto the resource model. The
// else-branches only null a value when the current model value is still
// Unknown; a known configured value that NITRO happens to omit from GET
// (omit-on-default trap) is preserved to avoid a spurious diff /
// "inconsistent result after apply".
func nsip6SetAttrFromGet(ctx context.Context, data *Nsip6ResourceModel, getResponseData map[string]interface{}) *Nsip6ResourceModel {
	tflog.Debug(ctx, "In nsip6SetAttrFromGet Function")

	if val, ok := getResponseData["advertiseondefaultpartition"]; ok && val != nil {
		data.Advertiseondefaultpartition = types.StringValue(val.(string))
	} else if data.Advertiseondefaultpartition.IsUnknown() {
		data.Advertiseondefaultpartition = types.StringNull()
	}
	if val, ok := getResponseData["decrementhoplimit"]; ok && val != nil {
		data.Decrementhoplimit = types.StringValue(val.(string))
	} else if data.Decrementhoplimit.IsUnknown() {
		data.Decrementhoplimit = types.StringNull()
	}
	if val, ok := getResponseData["dynamicrouting"]; ok && val != nil {
		data.Dynamicrouting = types.StringValue(val.(string))
	} else if data.Dynamicrouting.IsUnknown() {
		data.Dynamicrouting = types.StringNull()
	}
	if val, ok := getResponseData["ftp"]; ok && val != nil {
		data.Ftp = types.StringValue(val.(string))
	} else if data.Ftp.IsUnknown() {
		data.Ftp = types.StringNull()
	}
	if val, ok := getResponseData["gui"]; ok && val != nil {
		data.Gui = types.StringValue(val.(string))
	} else if data.Gui.IsUnknown() {
		data.Gui = types.StringNull()
	}
	if val, ok := getResponseData["hostroute"]; ok && val != nil {
		data.Hostroute = types.StringValue(val.(string))
	} else if data.Hostroute.IsUnknown() {
		data.Hostroute = types.StringNull()
	}
	if val, ok := getResponseData["icmp"]; ok && val != nil {
		data.Icmp = types.StringValue(val.(string))
	} else if data.Icmp.IsUnknown() {
		data.Icmp = types.StringNull()
	}
	if val, ok := getResponseData["icmpresponse"]; ok && val != nil {
		data.Icmpresponse = types.StringValue(val.(string))
	} else if data.Icmpresponse.IsUnknown() {
		data.Icmpresponse = types.StringNull()
	}
	if val, ok := getResponseData["ip6hostrtgw"]; ok && val != nil {
		data.Ip6hostrtgw = types.StringValue(val.(string))
	} else if data.Ip6hostrtgw.IsUnknown() {
		data.Ip6hostrtgw = types.StringNull()
	}
	if val, ok := getResponseData["ipv6address"]; ok && val != nil {
		data.Ipv6address = types.StringValue(val.(string))
	} else if data.Ipv6address.IsUnknown() {
		data.Ipv6address = types.StringNull()
	}
	if val, ok := getResponseData["map"]; ok && val != nil {
		data.Map = types.StringValue(val.(string))
	} else if data.Map.IsUnknown() {
		data.Map = types.StringNull()
	}
	if val, ok := getResponseData["metric"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Metric = types.Int64Value(intVal)
		}
	} else if data.Metric.IsUnknown() {
		data.Metric = types.Int64Null()
	}
	if val, ok := getResponseData["mgmtaccess"]; ok && val != nil {
		data.Mgmtaccess = types.StringValue(val.(string))
	} else if data.Mgmtaccess.IsUnknown() {
		data.Mgmtaccess = types.StringNull()
	}
	if val, ok := getResponseData["mptcpadvertise"]; ok && val != nil {
		data.Mptcpadvertise = types.StringValue(val.(string))
	} else if data.Mptcpadvertise.IsUnknown() {
		data.Mptcpadvertise = types.StringNull()
	}
	if val, ok := getResponseData["nd"]; ok && val != nil {
		data.Nd = types.StringValue(val.(string))
	} else if data.Nd.IsUnknown() {
		data.Nd = types.StringNull()
	}
	if val, ok := getResponseData["ndowner"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ndowner = types.Int64Value(intVal)
		}
	} else if data.Ndowner.IsUnknown() {
		data.Ndowner = types.Int64Null()
	}
	if val, ok := getResponseData["networkroute"]; ok && val != nil {
		data.Networkroute = types.StringValue(val.(string))
	} else if data.Networkroute.IsUnknown() {
		data.Networkroute = types.StringNull()
	}
	if val, ok := getResponseData["ospf6lsatype"]; ok && val != nil {
		data.Ospf6lsatype = types.StringValue(val.(string))
	} else if data.Ospf6lsatype.IsUnknown() {
		data.Ospf6lsatype = types.StringNull()
	}
	if val, ok := getResponseData["ospfarea"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ospfarea = types.Int64Value(intVal)
		}
	} else if data.Ospfarea.IsUnknown() {
		data.Ospfarea = types.Int64Null()
	}
	if val, ok := getResponseData["ownerdownresponse"]; ok && val != nil {
		data.Ownerdownresponse = types.StringValue(val.(string))
	} else if data.Ownerdownresponse.IsUnknown() {
		data.Ownerdownresponse = types.StringNull()
	}
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		}
	} else if data.Ownernode.IsUnknown() {
		data.Ownernode = types.Int64Null()
	}
	if val, ok := getResponseData["restrictaccess"]; ok && val != nil {
		data.Restrictaccess = types.StringValue(val.(string))
	} else if data.Restrictaccess.IsUnknown() {
		data.Restrictaccess = types.StringNull()
	}
	if val, ok := getResponseData["scope"]; ok && val != nil {
		data.Scope = types.StringValue(val.(string))
	} else if data.Scope.IsUnknown() {
		data.Scope = types.StringNull()
	}
	if val, ok := getResponseData["snmp"]; ok && val != nil {
		data.Snmp = types.StringValue(val.(string))
	} else if data.Snmp.IsUnknown() {
		data.Snmp = types.StringNull()
	}
	if val, ok := getResponseData["ssh"]; ok && val != nil {
		data.Ssh = types.StringValue(val.(string))
	} else if data.Ssh.IsUnknown() {
		data.Ssh = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else if data.State.IsUnknown() {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["tag"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tag = types.Int64Value(intVal)
		}
	} else if data.Tag.IsUnknown() {
		data.Tag = types.Int64Null()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else if data.Td.IsUnknown() {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["telnet"]; ok && val != nil {
		data.Telnet = types.StringValue(val.(string))
	} else if data.Telnet.IsUnknown() {
		data.Telnet = types.StringNull()
	}
	// type is returned by NITRO under the read-only "iptype" array field.
	if val, ok := getResponseData["iptype"]; ok && val != nil {
		if s, ok2 := nsip6IptypeToString(val); ok2 {
			data.Type = types.StringValue(s)
		} else if data.Type.IsUnknown() {
			data.Type = types.StringNull()
		}
	} else if data.Type.IsUnknown() {
		data.Type = types.StringNull()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	} else if data.Vlan.IsUnknown() {
		data.Vlan = types.Int64Null()
	}
	if val, ok := getResponseData["vrid6"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vrid6 = types.Int64Value(intVal)
		}
	} else if data.Vrid6.IsUnknown() {
		data.Vrid6 = types.Int64Null()
	}
	if val, ok := getResponseData["vserver"]; ok && val != nil {
		data.Vserver = types.StringValue(val.(string))
	} else if data.Vserver.IsUnknown() {
		data.Vserver = types.StringNull()
	}
	if val, ok := getResponseData["vserverrhilevel"]; ok && val != nil {
		data.Vserverrhilevel = types.StringValue(val.(string))
	} else if data.Vserverrhilevel.IsUnknown() {
		data.Vserverrhilevel = types.StringNull()
	}

	// NOTE: data.Id is set once in Create (plain ipv6address) and preserved from
	// prior state in Read/Update; it is intentionally not recomputed here.

	return data
}

// nsip6SetAttrFromGetForDatasource faithfully copies every field from the GET
// response into the model for the datasource (no prior plan/state to preserve)
// and sets the resource ID.
func nsip6SetAttrFromGetForDatasource(ctx context.Context, data *Nsip6ResourceModel, getResponseData map[string]interface{}) *Nsip6ResourceModel {
	tflog.Debug(ctx, "In nsip6SetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["advertiseondefaultpartition"]; ok && val != nil {
		data.Advertiseondefaultpartition = types.StringValue(val.(string))
	} else {
		data.Advertiseondefaultpartition = types.StringNull()
	}
	if val, ok := getResponseData["decrementhoplimit"]; ok && val != nil {
		data.Decrementhoplimit = types.StringValue(val.(string))
	} else {
		data.Decrementhoplimit = types.StringNull()
	}
	if val, ok := getResponseData["dynamicrouting"]; ok && val != nil {
		data.Dynamicrouting = types.StringValue(val.(string))
	} else {
		data.Dynamicrouting = types.StringNull()
	}
	if val, ok := getResponseData["ftp"]; ok && val != nil {
		data.Ftp = types.StringValue(val.(string))
	} else {
		data.Ftp = types.StringNull()
	}
	if val, ok := getResponseData["gui"]; ok && val != nil {
		data.Gui = types.StringValue(val.(string))
	} else {
		data.Gui = types.StringNull()
	}
	if val, ok := getResponseData["hostroute"]; ok && val != nil {
		data.Hostroute = types.StringValue(val.(string))
	} else {
		data.Hostroute = types.StringNull()
	}
	if val, ok := getResponseData["icmp"]; ok && val != nil {
		data.Icmp = types.StringValue(val.(string))
	} else {
		data.Icmp = types.StringNull()
	}
	if val, ok := getResponseData["icmpresponse"]; ok && val != nil {
		data.Icmpresponse = types.StringValue(val.(string))
	} else {
		data.Icmpresponse = types.StringNull()
	}
	if val, ok := getResponseData["ip6hostrtgw"]; ok && val != nil {
		data.Ip6hostrtgw = types.StringValue(val.(string))
	} else {
		data.Ip6hostrtgw = types.StringNull()
	}
	if val, ok := getResponseData["ipv6address"]; ok && val != nil {
		data.Ipv6address = types.StringValue(val.(string))
	} else {
		data.Ipv6address = types.StringNull()
	}
	if val, ok := getResponseData["map"]; ok && val != nil {
		data.Map = types.StringValue(val.(string))
	} else {
		data.Map = types.StringNull()
	}
	if val, ok := getResponseData["metric"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Metric = types.Int64Value(intVal)
		} else {
			data.Metric = types.Int64Null()
		}
	} else {
		data.Metric = types.Int64Null()
	}
	if val, ok := getResponseData["mgmtaccess"]; ok && val != nil {
		data.Mgmtaccess = types.StringValue(val.(string))
	} else {
		data.Mgmtaccess = types.StringNull()
	}
	if val, ok := getResponseData["mptcpadvertise"]; ok && val != nil {
		data.Mptcpadvertise = types.StringValue(val.(string))
	} else {
		data.Mptcpadvertise = types.StringNull()
	}
	if val, ok := getResponseData["nd"]; ok && val != nil {
		data.Nd = types.StringValue(val.(string))
	} else {
		data.Nd = types.StringNull()
	}
	if val, ok := getResponseData["ndowner"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ndowner = types.Int64Value(intVal)
		} else {
			data.Ndowner = types.Int64Null()
		}
	} else {
		data.Ndowner = types.Int64Null()
	}
	if val, ok := getResponseData["networkroute"]; ok && val != nil {
		data.Networkroute = types.StringValue(val.(string))
	} else {
		data.Networkroute = types.StringNull()
	}
	if val, ok := getResponseData["ospf6lsatype"]; ok && val != nil {
		data.Ospf6lsatype = types.StringValue(val.(string))
	} else {
		data.Ospf6lsatype = types.StringNull()
	}
	if val, ok := getResponseData["ospfarea"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ospfarea = types.Int64Value(intVal)
		} else {
			data.Ospfarea = types.Int64Null()
		}
	} else {
		data.Ospfarea = types.Int64Null()
	}
	if val, ok := getResponseData["ownerdownresponse"]; ok && val != nil {
		data.Ownerdownresponse = types.StringValue(val.(string))
	} else {
		data.Ownerdownresponse = types.StringNull()
	}
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		} else {
			data.Ownernode = types.Int64Null()
		}
	} else {
		data.Ownernode = types.Int64Null()
	}
	if val, ok := getResponseData["restrictaccess"]; ok && val != nil {
		data.Restrictaccess = types.StringValue(val.(string))
	} else {
		data.Restrictaccess = types.StringNull()
	}
	if val, ok := getResponseData["scope"]; ok && val != nil {
		data.Scope = types.StringValue(val.(string))
	} else {
		data.Scope = types.StringNull()
	}
	if val, ok := getResponseData["snmp"]; ok && val != nil {
		data.Snmp = types.StringValue(val.(string))
	} else {
		data.Snmp = types.StringNull()
	}
	if val, ok := getResponseData["ssh"]; ok && val != nil {
		data.Ssh = types.StringValue(val.(string))
	} else {
		data.Ssh = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["tag"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tag = types.Int64Value(intVal)
		} else {
			data.Tag = types.Int64Null()
		}
	} else {
		data.Tag = types.Int64Null()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		} else {
			data.Td = types.Int64Null()
		}
	} else {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["telnet"]; ok && val != nil {
		data.Telnet = types.StringValue(val.(string))
	} else {
		data.Telnet = types.StringNull()
	}
	// type is returned by NITRO under the read-only "iptype" array field.
	if val, ok := getResponseData["iptype"]; ok && val != nil {
		if s, ok2 := nsip6IptypeToString(val); ok2 {
			data.Type = types.StringValue(s)
		} else {
			data.Type = types.StringNull()
		}
	} else {
		data.Type = types.StringNull()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		} else {
			data.Vlan = types.Int64Null()
		}
	} else {
		data.Vlan = types.Int64Null()
	}
	if val, ok := getResponseData["vrid6"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vrid6 = types.Int64Value(intVal)
		} else {
			data.Vrid6 = types.Int64Null()
		}
	} else {
		data.Vrid6 = types.Int64Null()
	}
	if val, ok := getResponseData["vserver"]; ok && val != nil {
		data.Vserver = types.StringValue(val.(string))
	} else {
		data.Vserver = types.StringNull()
	}
	if val, ok := getResponseData["vserverrhilevel"]; ok && val != nil {
		data.Vserverrhilevel = types.StringValue(val.(string))
	} else {
		data.Vserverrhilevel = types.StringNull()
	}

	// Datasource has no Create; set the ID to the plain ipv6address value,
	// matching the resource Create ID scheme.
	data.Id = types.StringValue(data.Ipv6address.ValueString())

	return data
}
