package nsip6

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
			"advertiseondefaultpartition": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Advertise VIPs from Shared VLAN on Default Partition",
			},
			"decrementhoplimit": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Decrement Hop Limit by 1 when ENABLED.This setting is applicable only for UDP traffic.",
			},
			"dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Allow dynamic routing on this IP address. Specific to Subnet IPv6 (SNIP6) address.",
			},
			"ftp": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Allow File Transfer Protocol (FTP) access to this IP address.",
			},
			"gui": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Allow graphical user interface (GUI) access to this IP address.",
			},
			"hostroute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to push the VIP6 to ZebOS routing table for Kernel route redistribution through dynamic routing protocols.",
			},
			"icmp": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Respond to ICMP requests for this IP address.",
			},
			"icmpresponse": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("5"),
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
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Allow access to management applications on this IP address.",
			},
			"mptcpadvertise": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If enabled, this IP will be advertised by Citrix ADC to MPTCP enabled clients as part of ADD_ADDR option.",
			},
			"nd": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Respond to Neighbor Discovery (ND) requests for this IP address.",
			},
			"ndowner": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(255),
				Description: "NdOwner in Cluster for VIPS and Striped SNIPS",
			},
			"networkroute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to push the SNIP6 subnet to ZebOS routing table for Kernel route redistribution through dynamic routing protocol.",
			},
			"ospf6lsatype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("EXTERNAL"),
				Description: "Type of LSAs to be used by the IPv6 OSPF protocol, running on the Citrix ADC, for advertising the route for the VIP6 address.",
			},
			"ospfarea": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(-1),
				Description: "ID of the area in which the Intra-Area-Prefix LSAs are to be advertised for the VIP6 address by the IPv6 OSPF protocol running on the Citrix ADC. When ospfArea is not set, VIP6 is advertised on all areas.",
			},
			"ownerdownresponse": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "in cluster system, if the owner node is down, whether should it respond to icmp/arp",
			},
			"ownernode": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(255),
				Description: "ID of the cluster node for which you are adding the IP address. Must be used if you want the IP address to be active only on the specific node. Can be configured only through the cluster IP address. Cannot be changed after the IP address is created.",
			},
			"restrictaccess": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Block access to nonmanagement applications on this IP address. This option is applicable forMIP6s, SNIP6s, and NSIP6s, and is disabled by default. Nonmanagement applications can run on the underlying Citrix ADC Free BSD operating system.",
			},
			"scope": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("global"),
				Description: "Scope of the IPv6 address to be created. Cannot be changed after the IP address is created.",
			},
			"snmp": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Allow Simple Network Management Protocol (SNMP) access to this IP address.",
			},
			"ssh": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Allow secure Shell (SSH) access to this IP address.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
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
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Allow Telnet access to this IP address.",
			},
			"type": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("SNIP"),
				Description: "Type of IP address to be created on the Citrix ADC. Cannot be changed after the IP address is created.",
			},
			"vlan": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
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
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable the state of all the virtual servers associated with this VIP6 address.",
			},
			"vserverrhilevel": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ONE_VSERVER"),
				Description: "Advertise or do not advertise the route for the Virtual IP (VIP6) address on the basis of the state of the virtual servers associated with that VIP6.\n* NONE - Advertise the route for the VIP6 address, irrespective of the state of the virtual servers associated with the address.\n* ONE VSERVER - Advertise the route for the VIP6 address if at least one of the associated virtual servers is in UP state.\n* ALL VSERVER - Advertise the route for the VIP6 address if all of the associated virtual servers are in UP state.\n* VSVR_CNTRLD.   Advertise the route for the VIP address according to the  RHIstate (RHI STATE) parameter setting on all the associated virtual servers of the VIP address along with their states.\n\nWhen Vserver RHI Level (RHI) parameter is set to VSVR_CNTRLD, the following are different RHI behaviors for the VIP address on the basis of RHIstate (RHI STATE) settings on the virtual servers associated with the VIP address:\n * If you set RHI STATE to PASSIVE on all virtual servers, the Citrix ADC always advertises the route for the VIP address.\n * If you set RHI STATE to ACTIVE on all virtual servers, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers is in UP state.\n *If you set RHI STATE to ACTIVE on some and PASSIVE on others, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers, whose RHI STATE set to ACTIVE, is in UP state.",
			},
		},
	}
}

func nsip6GetThePayloadFromtheConfig(ctx context.Context, data *Nsip6ResourceModel) ns.Nsip6 {
	tflog.Debug(ctx, "In nsip6GetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsip6 := ns.Nsip6{}
	if !data.Advertiseondefaultpartition.IsNull() {
		nsip6.Advertiseondefaultpartition = data.Advertiseondefaultpartition.ValueString()
	}
	if !data.Decrementhoplimit.IsNull() {
		nsip6.Decrementhoplimit = data.Decrementhoplimit.ValueString()
	}
	if !data.Dynamicrouting.IsNull() {
		nsip6.Dynamicrouting = data.Dynamicrouting.ValueString()
	}
	if !data.Ftp.IsNull() {
		nsip6.Ftp = data.Ftp.ValueString()
	}
	if !data.Gui.IsNull() {
		nsip6.Gui = data.Gui.ValueString()
	}
	if !data.Hostroute.IsNull() {
		nsip6.Hostroute = data.Hostroute.ValueString()
	}
	if !data.Icmp.IsNull() {
		nsip6.Icmp = data.Icmp.ValueString()
	}
	if !data.Icmpresponse.IsNull() {
		nsip6.Icmpresponse = data.Icmpresponse.ValueString()
	}
	if !data.Ip6hostrtgw.IsNull() {
		nsip6.Ip6hostrtgw = data.Ip6hostrtgw.ValueString()
	}
	if !data.Ipv6address.IsNull() {
		nsip6.Ipv6address = data.Ipv6address.ValueString()
	}
	if !data.Map.IsNull() {
		nsip6.Map = data.Map.ValueString()
	}
	if !data.Metric.IsNull() {
		nsip6.Metric = utils.IntPtr(int(data.Metric.ValueInt64()))
	}
	if !data.Mgmtaccess.IsNull() {
		nsip6.Mgmtaccess = data.Mgmtaccess.ValueString()
	}
	if !data.Mptcpadvertise.IsNull() {
		nsip6.Mptcpadvertise = data.Mptcpadvertise.ValueString()
	}
	if !data.Nd.IsNull() {
		nsip6.Nd = data.Nd.ValueString()
	}
	if !data.Ndowner.IsNull() {
		nsip6.Ndowner = utils.IntPtr(int(data.Ndowner.ValueInt64()))
	}
	if !data.Networkroute.IsNull() {
		nsip6.Networkroute = data.Networkroute.ValueString()
	}
	if !data.Ospf6lsatype.IsNull() {
		nsip6.Ospf6lsatype = data.Ospf6lsatype.ValueString()
	}
	if !data.Ospfarea.IsNull() {
		nsip6.Ospfarea = utils.IntPtr(int(data.Ospfarea.ValueInt64()))
	}
	if !data.Ownerdownresponse.IsNull() {
		nsip6.Ownerdownresponse = data.Ownerdownresponse.ValueString()
	}
	if !data.Ownernode.IsNull() {
		nsip6.Ownernode = utils.IntPtr(int(data.Ownernode.ValueInt64()))
	}
	if !data.Restrictaccess.IsNull() {
		nsip6.Restrictaccess = data.Restrictaccess.ValueString()
	}
	if !data.Scope.IsNull() {
		nsip6.Scope = data.Scope.ValueString()
	}
	if !data.Snmp.IsNull() {
		nsip6.Snmp = data.Snmp.ValueString()
	}
	if !data.Ssh.IsNull() {
		nsip6.Ssh = data.Ssh.ValueString()
	}
	if !data.State.IsNull() {
		nsip6.State = data.State.ValueString()
	}
	if !data.Tag.IsNull() {
		nsip6.Tag = utils.IntPtr(int(data.Tag.ValueInt64()))
	}
	if !data.Td.IsNull() {
		nsip6.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Telnet.IsNull() {
		nsip6.Telnet = data.Telnet.ValueString()
	}
	if !data.Type.IsNull() {
		nsip6.Type = data.Type.ValueString()
	}
	if !data.Vlan.IsNull() {
		nsip6.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}
	if !data.Vrid6.IsNull() {
		nsip6.Vrid6 = utils.IntPtr(int(data.Vrid6.ValueInt64()))
	}
	if !data.Vserver.IsNull() {
		nsip6.Vserver = data.Vserver.ValueString()
	}
	if !data.Vserverrhilevel.IsNull() {
		nsip6.Vserverrhilevel = data.Vserverrhilevel.ValueString()
	}

	return nsip6
}

func nsip6SetAttrFromGet(ctx context.Context, data *Nsip6ResourceModel, getResponseData map[string]interface{}) *Nsip6ResourceModel {
	tflog.Debug(ctx, "In nsip6SetAttrFromGet Function")

	// Convert API response to model
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
		}
	} else {
		data.Tag = types.Int64Null()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
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
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	} else {
		data.Vlan = types.Int64Null()
	}
	if val, ok := getResponseData["vrid6"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vrid6 = types.Int64Value(intVal)
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

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s,%d", data.Ipv6address.ValueString(), data.Td.ValueInt64()))

	return data
}
