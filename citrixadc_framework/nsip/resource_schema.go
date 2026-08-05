package nsip

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

// NsipResourceModel describes the resource data model.
type NsipResourceModel struct {
	Id                          types.String `tfsdk:"id"`
	Advertiseondefaultpartition types.String `tfsdk:"advertiseondefaultpartition"`
	Arp                         types.String `tfsdk:"arp"`
	Arpowner                    types.Int64  `tfsdk:"arpowner"`
	Arpresponse                 types.String `tfsdk:"arpresponse"`
	Bgp                         types.String `tfsdk:"bgp"`
	Decrementttl                types.String `tfsdk:"decrementttl"`
	Dynamicrouting              types.String `tfsdk:"dynamicrouting"`
	Ftp                         types.String `tfsdk:"ftp"`
	Gui                         types.String `tfsdk:"gui"`
	Hostroute                   types.String `tfsdk:"hostroute"`
	Hostrtgw                    types.String `tfsdk:"hostrtgw"`
	Icmp                        types.String `tfsdk:"icmp"`
	Icmpresponse                types.String `tfsdk:"icmpresponse"`
	Ipaddress                   types.String `tfsdk:"ipaddress"`
	Metric                      types.Int64  `tfsdk:"metric"`
	Mgmtaccess                  types.String `tfsdk:"mgmtaccess"`
	Mptcpadvertise              types.String `tfsdk:"mptcpadvertise"`
	Netmask                     types.String `tfsdk:"netmask"`
	Networkroute                types.String `tfsdk:"networkroute"`
	Ospf                        types.String `tfsdk:"ospf"`
	Ospfarea                    types.Int64  `tfsdk:"ospfarea"`
	Ospflsatype                 types.String `tfsdk:"ospflsatype"`
	Ownerdownresponse           types.String `tfsdk:"ownerdownresponse"`
	Ownernode                   types.Int64  `tfsdk:"ownernode"`
	Restrictaccess              types.String `tfsdk:"restrictaccess"`
	Rip                         types.String `tfsdk:"rip"`
	Snmp                        types.String `tfsdk:"snmp"`
	Ssh                         types.String `tfsdk:"ssh"`
	State                       types.String `tfsdk:"state"`
	Tag                         types.Int64  `tfsdk:"tag"`
	Td                          types.Int64  `tfsdk:"td"`
	Telnet                      types.String `tfsdk:"telnet"`
	Type                        types.String `tfsdk:"type"`
	Vrid                        types.Int64  `tfsdk:"vrid"`
	Vserver                     types.String `tfsdk:"vserver"`
	Vserverrhilevel             types.String `tfsdk:"vserverrhilevel"`
}

func (r *NsipResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsip resource.",
			},
			"advertiseondefaultpartition": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Advertise VIPs from Shared VLAN on Default Partition.",
			},
			"arp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Respond to ARP requests for this IP address.",
			},
			"arpowner": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The arp owner in a Cluster for this IP address. It can vary from 0 to 31.",
			},
			"arpresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Respond to ARP requests for a Virtual IP (VIP) address on the basis of the states of the virtual servers associated with that VIP.",
			},
			"bgp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this option to enable or disable BGP on this IP address for the entity.",
			},
			"decrementttl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Decrement TTL by 1 when ENABLED.This setting is applicable only for UDP traffic.",
			},
			"dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow dynamic routing on this IP address. Specific to Subnet IP (SNIP) address.",
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
				Description: "Option to push the VIP to ZebOS routing table for Kernel route redistribution through dynamic routing protocols",
			},
			"hostrtgw": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the gateway of the route for this VIP address.",
			},
			"icmp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Respond to ICMP requests for this IP address.",
			},
			"icmpresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Respond to ICMP requests for a Virtual IP (VIP) address on the basis of the states of the virtual servers associated with that VIP.",
			},
			"ipaddress": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IPv4 address to create on the Citrix ADC. Cannot be changed after the IP address is created.",
			},
			"metric": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value to add to or subtract from the cost of the route advertised for the VIP address.",
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
			"netmask": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Subnet mask associated with the IP address.",
			},
			"networkroute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to push the SNIP subnet to ZebOS routing table for Kernel route redistribution through dynamic routing protocol.",
			},
			"ospf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this option to enable or disable OSPF on this IP address for the entity.",
			},
			"ospfarea": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the area in which the type1 link-state advertisements (LSAs) are to be advertised for this virtual IP (VIP) address by the OSPF protocol running on the Citrix ADC. When this parameter is not set, the VIP is advertised on all areas.",
			},
			"ospflsatype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of LSAs to be used by the OSPF protocol, running on the Citrix ADC, for advertising the route for this VIP address.",
			},
			"ownerdownresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "in cluster system, if the owner node is down, whether should it respond to icmp/arp",
			},
			"ownernode": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "The owner node in a Cluster for this IP address. Owner node can vary from 0 to 31. If ownernode is not specified then the IP is treated as Striped IP.",
			},
			"restrictaccess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Block access to nonmanagement applications on this IP. This option is applicable for MIPs, SNIPs, and NSIP, and is disabled by default.",
			},
			"rip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this option to enable or disable RIP on this IP address for the entity.",
			},
			"snmp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Simple Network Management Protocol (SNMP) access to this IP address.",
			},
			"ssh": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow secure shell (SSH) access to this IP address.",
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
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. TD id 4095 is used reserved for LSN use",
			},
			"telnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Telnet access to this IP address.",
			},
			"type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Type of the IP address to create on the Citrix ADC. Cannot be changed after the IP address is created.",
			},
			"vrid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "A positive integer that uniquely identifies a VMAC address for binding to this VIP address. This binding is used to set up Citrix ADCs in an active-active configuration using VRRP.",
			},
			"vserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this option to set (enable or disable) the virtual server attribute for this IP address.",
			},
			"vserverrhilevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Advertise the route for the Virtual IP (VIP) address on the basis of the state of the virtual servers associated with that VIP.",
			},
		},
	}
}

// nsipGetThePayloadFromtheConfig builds the full add/create payload from the plan.
func nsipGetThePayloadFromtheConfig(ctx context.Context, data *NsipResourceModel) ns.Nsip {
	tflog.Debug(ctx, "In nsipGetThePayloadFromtheConfig Function")

	nsip := ns.Nsip{}
	if !data.Advertiseondefaultpartition.IsNull() && !data.Advertiseondefaultpartition.IsUnknown() {
		nsip.Advertiseondefaultpartition = data.Advertiseondefaultpartition.ValueString()
	}
	if !data.Arp.IsNull() && !data.Arp.IsUnknown() {
		nsip.Arp = data.Arp.ValueString()
	}
	if !data.Arpowner.IsNull() && !data.Arpowner.IsUnknown() {
		nsip.Arpowner = utils.IntPtr(int(data.Arpowner.ValueInt64()))
	}
	if !data.Arpresponse.IsNull() && !data.Arpresponse.IsUnknown() {
		nsip.Arpresponse = data.Arpresponse.ValueString()
	}
	if !data.Bgp.IsNull() && !data.Bgp.IsUnknown() {
		nsip.Bgp = data.Bgp.ValueString()
	}
	if !data.Decrementttl.IsNull() && !data.Decrementttl.IsUnknown() {
		nsip.Decrementttl = data.Decrementttl.ValueString()
	}
	if !data.Dynamicrouting.IsNull() && !data.Dynamicrouting.IsUnknown() {
		nsip.Dynamicrouting = data.Dynamicrouting.ValueString()
	}
	if !data.Ftp.IsNull() && !data.Ftp.IsUnknown() {
		nsip.Ftp = data.Ftp.ValueString()
	}
	if !data.Gui.IsNull() && !data.Gui.IsUnknown() {
		nsip.Gui = data.Gui.ValueString()
	}
	if !data.Hostroute.IsNull() && !data.Hostroute.IsUnknown() {
		nsip.Hostroute = data.Hostroute.ValueString()
	}
	if !data.Hostrtgw.IsNull() && !data.Hostrtgw.IsUnknown() {
		nsip.Hostrtgw = data.Hostrtgw.ValueString()
	}
	if !data.Icmp.IsNull() && !data.Icmp.IsUnknown() {
		nsip.Icmp = data.Icmp.ValueString()
	}
	if !data.Icmpresponse.IsNull() && !data.Icmpresponse.IsUnknown() {
		nsip.Icmpresponse = data.Icmpresponse.ValueString()
	}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		nsip.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Metric.IsNull() && !data.Metric.IsUnknown() {
		nsip.Metric = utils.IntPtr(int(data.Metric.ValueInt64()))
	}
	if !data.Mgmtaccess.IsNull() && !data.Mgmtaccess.IsUnknown() {
		nsip.Mgmtaccess = data.Mgmtaccess.ValueString()
	}
	if !data.Mptcpadvertise.IsNull() && !data.Mptcpadvertise.IsUnknown() {
		nsip.Mptcpadvertise = data.Mptcpadvertise.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		nsip.Netmask = data.Netmask.ValueString()
	}
	if !data.Networkroute.IsNull() && !data.Networkroute.IsUnknown() {
		nsip.Networkroute = data.Networkroute.ValueString()
	}
	if !data.Ospf.IsNull() && !data.Ospf.IsUnknown() {
		nsip.Ospf = data.Ospf.ValueString()
	}
	if !data.Ospfarea.IsNull() && !data.Ospfarea.IsUnknown() {
		nsip.Ospfarea = utils.IntPtr(int(data.Ospfarea.ValueInt64()))
	}
	if !data.Ospflsatype.IsNull() && !data.Ospflsatype.IsUnknown() {
		nsip.Ospflsatype = data.Ospflsatype.ValueString()
	}
	if !data.Ownerdownresponse.IsNull() && !data.Ownerdownresponse.IsUnknown() {
		nsip.Ownerdownresponse = data.Ownerdownresponse.ValueString()
	}
	if !data.Ownernode.IsNull() && !data.Ownernode.IsUnknown() {
		nsip.Ownernode = utils.IntPtr(int(data.Ownernode.ValueInt64()))
	}
	if !data.Restrictaccess.IsNull() && !data.Restrictaccess.IsUnknown() {
		nsip.Restrictaccess = data.Restrictaccess.ValueString()
	}
	if !data.Rip.IsNull() && !data.Rip.IsUnknown() {
		nsip.Rip = data.Rip.ValueString()
	}
	if !data.Snmp.IsNull() && !data.Snmp.IsUnknown() {
		nsip.Snmp = data.Snmp.ValueString()
	}
	if !data.Ssh.IsNull() && !data.Ssh.IsUnknown() {
		nsip.Ssh = data.Ssh.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		nsip.State = data.State.ValueString()
	}
	if !data.Tag.IsNull() && !data.Tag.IsUnknown() {
		nsip.Tag = utils.IntPtr(int(data.Tag.ValueInt64()))
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		nsip.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Telnet.IsNull() && !data.Telnet.IsUnknown() {
		nsip.Telnet = data.Telnet.ValueString()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		nsip.Type = data.Type.ValueString()
	}
	if !data.Vrid.IsNull() && !data.Vrid.IsUnknown() {
		nsip.Vrid = utils.IntPtr(int(data.Vrid.ValueInt64()))
	}
	if !data.Vserver.IsNull() && !data.Vserver.IsUnknown() {
		nsip.Vserver = data.Vserver.ValueString()
	}
	if !data.Vserverrhilevel.IsNull() && !data.Vserverrhilevel.IsUnknown() {
		nsip.Vserverrhilevel = data.Vserverrhilevel.ValueString()
	}

	return nsip
}

// nsipSetAttrFromGet maps the GET response onto the resource model.
// It preserves any known configured/state value when a field is omitted by
// the NITRO GET response (omit-on-default trap): the else-branch only nulls a
// value that is still Unknown, never clobbering a known value.
func nsipSetAttrFromGet(ctx context.Context, data *NsipResourceModel, getResponseData map[string]interface{}) *NsipResourceModel {
	tflog.Debug(ctx, "In nsipSetAttrFromGet Function")

	if val, ok := getResponseData["advertiseondefaultpartition"]; ok && val != nil {
		data.Advertiseondefaultpartition = types.StringValue(val.(string))
	} else if data.Advertiseondefaultpartition.IsUnknown() {
		data.Advertiseondefaultpartition = types.StringNull()
	}
	if val, ok := getResponseData["arp"]; ok && val != nil {
		data.Arp = types.StringValue(val.(string))
	} else if data.Arp.IsUnknown() {
		data.Arp = types.StringNull()
	}
	if val, ok := getResponseData["arpowner"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Arpowner = types.Int64Value(intVal)
		}
	} else if data.Arpowner.IsUnknown() {
		data.Arpowner = types.Int64Null()
	}
	if val, ok := getResponseData["arpresponse"]; ok && val != nil {
		data.Arpresponse = types.StringValue(val.(string))
	} else if data.Arpresponse.IsUnknown() {
		data.Arpresponse = types.StringNull()
	}
	if val, ok := getResponseData["bgp"]; ok && val != nil {
		data.Bgp = types.StringValue(val.(string))
	} else if data.Bgp.IsUnknown() {
		data.Bgp = types.StringNull()
	}
	if val, ok := getResponseData["decrementttl"]; ok && val != nil {
		data.Decrementttl = types.StringValue(val.(string))
	} else if data.Decrementttl.IsUnknown() {
		data.Decrementttl = types.StringNull()
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
	if val, ok := getResponseData["hostrtgw"]; ok && val != nil {
		data.Hostrtgw = types.StringValue(val.(string))
	} else if data.Hostrtgw.IsUnknown() {
		data.Hostrtgw = types.StringNull()
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
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else if data.Ipaddress.IsUnknown() {
		data.Ipaddress = types.StringNull()
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
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else if data.Netmask.IsUnknown() {
		data.Netmask = types.StringNull()
	}
	if val, ok := getResponseData["networkroute"]; ok && val != nil {
		data.Networkroute = types.StringValue(val.(string))
	} else if data.Networkroute.IsUnknown() {
		data.Networkroute = types.StringNull()
	}
	if val, ok := getResponseData["ospf"]; ok && val != nil {
		data.Ospf = types.StringValue(val.(string))
	} else if data.Ospf.IsUnknown() {
		data.Ospf = types.StringNull()
	}
	if val, ok := getResponseData["ospfarea"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ospfarea = types.Int64Value(intVal)
		}
	} else if data.Ospfarea.IsUnknown() {
		data.Ospfarea = types.Int64Null()
	}
	if val, ok := getResponseData["ospflsatype"]; ok && val != nil {
		data.Ospflsatype = types.StringValue(val.(string))
	} else if data.Ospflsatype.IsUnknown() {
		data.Ospflsatype = types.StringNull()
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
	if val, ok := getResponseData["rip"]; ok && val != nil {
		data.Rip = types.StringValue(val.(string))
	} else if data.Rip.IsUnknown() {
		data.Rip = types.StringNull()
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
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else if data.Type.IsUnknown() {
		data.Type = types.StringNull()
	}
	if val, ok := getResponseData["vrid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vrid = types.Int64Value(intVal)
		}
	} else if data.Vrid.IsUnknown() {
		data.Vrid = types.Int64Null()
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

	// NOTE: data.Id is set once in Create (plain ipaddress) and preserved from
	// prior state in Read/Update; it is intentionally not recomputed here.

	return data
}

// nsipSetAttrFromGetForDatasource faithfully copies every field from the GET
// response into the model for the datasource (no prior plan/state to preserve),
// and sets the resource ID.
func nsipSetAttrFromGetForDatasource(ctx context.Context, data *NsipResourceModel, getResponseData map[string]interface{}) *NsipResourceModel {
	tflog.Debug(ctx, "In nsipSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["advertiseondefaultpartition"]; ok && val != nil {
		data.Advertiseondefaultpartition = types.StringValue(val.(string))
	} else {
		data.Advertiseondefaultpartition = types.StringNull()
	}
	if val, ok := getResponseData["arp"]; ok && val != nil {
		data.Arp = types.StringValue(val.(string))
	} else {
		data.Arp = types.StringNull()
	}
	if val, ok := getResponseData["arpowner"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Arpowner = types.Int64Value(intVal)
		} else {
			data.Arpowner = types.Int64Null()
		}
	} else {
		data.Arpowner = types.Int64Null()
	}
	if val, ok := getResponseData["arpresponse"]; ok && val != nil {
		data.Arpresponse = types.StringValue(val.(string))
	} else {
		data.Arpresponse = types.StringNull()
	}
	if val, ok := getResponseData["bgp"]; ok && val != nil {
		data.Bgp = types.StringValue(val.(string))
	} else {
		data.Bgp = types.StringNull()
	}
	if val, ok := getResponseData["decrementttl"]; ok && val != nil {
		data.Decrementttl = types.StringValue(val.(string))
	} else {
		data.Decrementttl = types.StringNull()
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
	if val, ok := getResponseData["hostrtgw"]; ok && val != nil {
		data.Hostrtgw = types.StringValue(val.(string))
	} else {
		data.Hostrtgw = types.StringNull()
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
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
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
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}
	if val, ok := getResponseData["networkroute"]; ok && val != nil {
		data.Networkroute = types.StringValue(val.(string))
	} else {
		data.Networkroute = types.StringNull()
	}
	if val, ok := getResponseData["ospf"]; ok && val != nil {
		data.Ospf = types.StringValue(val.(string))
	} else {
		data.Ospf = types.StringNull()
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
	if val, ok := getResponseData["ospflsatype"]; ok && val != nil {
		data.Ospflsatype = types.StringValue(val.(string))
	} else {
		data.Ospflsatype = types.StringNull()
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
	if val, ok := getResponseData["rip"]; ok && val != nil {
		data.Rip = types.StringValue(val.(string))
	} else {
		data.Rip = types.StringNull()
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
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}
	if val, ok := getResponseData["vrid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vrid = types.Int64Value(intVal)
		} else {
			data.Vrid = types.Int64Null()
		}
	} else {
		data.Vrid = types.Int64Null()
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

	// Datasource has no Create; set the ID to the plain ipaddress value,
	// matching the resource Create ID scheme.
	data.Id = types.StringValue(data.Ipaddress.ValueString())

	return data
}
