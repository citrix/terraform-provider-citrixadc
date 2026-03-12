package vpnintranetapplication

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpnintranetapplicationResourceModel describes the resource data model.
type VpnintranetapplicationResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Clientapplication   types.List   `tfsdk:"clientapplication"`
	Destip              types.String `tfsdk:"destip"`
	Destport            types.String `tfsdk:"destport"`
	Hostname            types.String `tfsdk:"hostname"`
	Interception        types.String `tfsdk:"interception"`
	Intranetapplication types.String `tfsdk:"intranetapplication"`
	Iprange             types.String `tfsdk:"iprange"`
	Netmask             types.String `tfsdk:"netmask"`
	Protocol            types.String `tfsdk:"protocol"`
	Spoofiip            types.String `tfsdk:"spoofiip"`
	Srcip               types.String `tfsdk:"srcip"`
	Srcport             types.Int64  `tfsdk:"srcport"`
}

func (r *VpnintranetapplicationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnintranetapplication resource.",
			},
			"clientapplication": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Names of the client applications, such as PuTTY and Xshell.",
			},
			"destip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Destination IP address, IP range, or host name of the intranet application. This address is the server IP address.",
			},
			"destport": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Destination TCP or UDP port number for the intranet application. Use a hyphen to specify a range of port numbers, for example 90-95.",
			},
			"hostname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the host for which to configure interception. The names are resolved during interception when users log on with the Citrix Gateway Plug-in.",
			},
			"interception": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Interception mode for the intranet application or resource. Correct value depends on the type of client software used to make connections. If the interception mode is set to TRANSPARENT, users connect with the Citrix Gateway Plug-in for Windows. With the PROXY setting, users connect with the Citrix Gateway Plug-in for Java.",
			},
			"intranetapplication": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the intranet application.",
			},
			"iprange": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "If you have multiple servers in your network, such as web, email, and file shares, configure an intranet application that includes the IP range for all the network applications. This allows users to access all the intranet applications contained in the IP address range.",
			},
			"netmask": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Destination subnet mask for the intranet application.",
			},
			"protocol": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Protocol used by the intranet application. If protocol is set to BOTH, TCP and UDP traffic is allowed.",
			},
			"spoofiip": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("True"),
				Description: "IP address that the intranet application will use to route the connection through the virtual adapter.",
			},
			"srcip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Source IP address. Required if interception mode is set to PROXY. Default is the loopback address, 127.0.0.1.",
			},
			"srcport": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Source port for the application for which the Citrix Gateway virtual server proxies the traffic. If users are connecting from a device that uses the Citrix Gateway Plug-in for Java, applications must be configured manually by using the source IP address and TCP port values specified in the intranet application profile. If a port value is not set, the destination port value is used.",
			},
		},
	}
}

func vpnintranetapplicationGetThePayloadFromtheConfig(ctx context.Context, data *VpnintranetapplicationResourceModel) vpn.Vpnintranetapplication {
	tflog.Debug(ctx, "In vpnintranetapplicationGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnintranetapplication := vpn.Vpnintranetapplication{}
	if !data.Destip.IsNull() {
		vpnintranetapplication.Destip = data.Destip.ValueString()
	}
	if !data.Destport.IsNull() {
		vpnintranetapplication.Destport = data.Destport.ValueString()
	}
	if !data.Hostname.IsNull() {
		vpnintranetapplication.Hostname = data.Hostname.ValueString()
	}
	if !data.Interception.IsNull() {
		vpnintranetapplication.Interception = data.Interception.ValueString()
	}
	if !data.Intranetapplication.IsNull() {
		vpnintranetapplication.Intranetapplication = data.Intranetapplication.ValueString()
	}
	if !data.Iprange.IsNull() {
		vpnintranetapplication.Iprange = data.Iprange.ValueString()
	}
	if !data.Netmask.IsNull() {
		vpnintranetapplication.Netmask = data.Netmask.ValueString()
	}
	if !data.Protocol.IsNull() {
		vpnintranetapplication.Protocol = data.Protocol.ValueString()
	}
	if !data.Spoofiip.IsNull() {
		vpnintranetapplication.Spoofiip = data.Spoofiip.ValueString()
	}
	if !data.Srcip.IsNull() {
		vpnintranetapplication.Srcip = data.Srcip.ValueString()
	}
	if !data.Srcport.IsNull() {
		vpnintranetapplication.Srcport = utils.IntPtr(int(data.Srcport.ValueInt64()))
	}

	return vpnintranetapplication
}

func vpnintranetapplicationSetAttrFromGet(ctx context.Context, data *VpnintranetapplicationResourceModel, getResponseData map[string]interface{}) *VpnintranetapplicationResourceModel {
	tflog.Debug(ctx, "In vpnintranetapplicationSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["destip"]; ok && val != nil {
		data.Destip = types.StringValue(val.(string))
	} else {
		data.Destip = types.StringNull()
	}
	if val, ok := getResponseData["destport"]; ok && val != nil {
		data.Destport = types.StringValue(val.(string))
	} else {
		data.Destport = types.StringNull()
	}
	if val, ok := getResponseData["hostname"]; ok && val != nil {
		data.Hostname = types.StringValue(val.(string))
	} else {
		data.Hostname = types.StringNull()
	}
	if val, ok := getResponseData["interception"]; ok && val != nil {
		data.Interception = types.StringValue(val.(string))
	} else {
		data.Interception = types.StringNull()
	}
	if val, ok := getResponseData["intranetapplication"]; ok && val != nil {
		data.Intranetapplication = types.StringValue(val.(string))
	} else {
		data.Intranetapplication = types.StringNull()
	}
	if val, ok := getResponseData["iprange"]; ok && val != nil {
		data.Iprange = types.StringValue(val.(string))
	} else {
		data.Iprange = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}
	if val, ok := getResponseData["protocol"]; ok && val != nil {
		data.Protocol = types.StringValue(val.(string))
	} else {
		data.Protocol = types.StringNull()
	}
	if val, ok := getResponseData["spoofiip"]; ok && val != nil {
		data.Spoofiip = types.StringValue(val.(string))
	} else {
		data.Spoofiip = types.StringNull()
	}
	if val, ok := getResponseData["srcip"]; ok && val != nil {
		data.Srcip = types.StringValue(val.(string))
	} else {
		data.Srcip = types.StringNull()
	}
	if val, ok := getResponseData["srcport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Srcport = types.Int64Value(intVal)
		}
	} else {
		data.Srcport = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Intranetapplication.ValueString())

	return data
}
