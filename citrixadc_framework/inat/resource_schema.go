package inat

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// InatResourceModel describes the resource data model.
type InatResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Connfailover types.String `tfsdk:"connfailover"`
	Ftp          types.String `tfsdk:"ftp"`
	Mode         types.String `tfsdk:"mode"`
	Name         types.String `tfsdk:"name"`
	Privateip    types.String `tfsdk:"privateip"`
	Proxyip      types.String `tfsdk:"proxyip"`
	Publicip     types.String `tfsdk:"publicip"`
	Tcpproxy     types.String `tfsdk:"tcpproxy"`
	Td           types.Int64  `tfsdk:"td"`
	Tftp         types.String `tfsdk:"tftp"`
	Useproxyport types.String `tfsdk:"useproxyport"`
	Usip         types.String `tfsdk:"usip"`
	Usnip        types.String `tfsdk:"usnip"`
}

func (r *InatResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the inat resource.",
			},
			"connfailover": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Synchronize connection information with the secondary appliance in a high availability (HA) pair. That is, synchronize all connection-related information for the INAT session",
			},
			"ftp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the FTP protocol on the server for transferring files between the client and the server.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Stateless translation.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the Inbound NAT (INAT) entry. Leading character must be a number or letter. Other characters allowed, after the first character, are @ _ - . (period) : (colon) # and space ( ).",
			},
			"privateip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the server to which the packet is sent by the Citrix ADC. Can be an IPv4 or IPv6 address.",
			},
			"proxyip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique IP address used as the source IP address in packets sent to the server. Must be a MIP or SNIP address.",
			},
			"publicip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Public IP address of packets received on the Citrix ADC. Can be aNetScaler-owned VIP or VIP6 address.",
			},
			"tcpproxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable TCP proxy, which enables the Citrix ADC to optimize the RNAT TCP traffic by using Layer 4 features.",
			},
			"td": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"tftp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "To enable/disable TFTP (Default DISABLED).",
			},
			"useproxyport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the Citrix ADC to proxy the source port of packets before sending the packets to the server.",
			},
			"usip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the Citrix ADC to retain the source IP address of packets before sending the packets to the server.",
			},
			"usnip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the Citrix ADC to use a SNIP address as the source IP address of packets before sending the packets to the server.",
			},
		},
	}
}

// inatGetThePayloadFromthePlan builds the full add payload (create).
func inatGetThePayloadFromthePlan(ctx context.Context, data *InatResourceModel) network.Inat {
	tflog.Debug(ctx, "In inatGetThePayloadFromthePlan Function")

	// Create API request body from the model
	inat := network.Inat{}
	if !data.Connfailover.IsNull() && !data.Connfailover.IsUnknown() {
		inat.Connfailover = data.Connfailover.ValueString()
	}
	if !data.Ftp.IsNull() && !data.Ftp.IsUnknown() {
		inat.Ftp = data.Ftp.ValueString()
	}
	if !data.Mode.IsNull() && !data.Mode.IsUnknown() {
		inat.Mode = data.Mode.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		inat.Name = data.Name.ValueString()
	}
	if !data.Privateip.IsNull() && !data.Privateip.IsUnknown() {
		inat.Privateip = data.Privateip.ValueString()
	}
	if !data.Proxyip.IsNull() && !data.Proxyip.IsUnknown() {
		inat.Proxyip = data.Proxyip.ValueString()
	}
	if !data.Publicip.IsNull() && !data.Publicip.IsUnknown() {
		inat.Publicip = data.Publicip.ValueString()
	}
	if !data.Tcpproxy.IsNull() && !data.Tcpproxy.IsUnknown() {
		inat.Tcpproxy = data.Tcpproxy.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		inat.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Tftp.IsNull() && !data.Tftp.IsUnknown() {
		inat.Tftp = data.Tftp.ValueString()
	}
	if !data.Useproxyport.IsNull() && !data.Useproxyport.IsUnknown() {
		inat.Useproxyport = data.Useproxyport.ValueString()
	}
	if !data.Usip.IsNull() && !data.Usip.IsUnknown() {
		inat.Usip = data.Usip.ValueString()
	}
	if !data.Usnip.IsNull() && !data.Usnip.IsUnknown() {
		inat.Usnip = data.Usnip.ValueString()
	}

	return inat
}

// inatGetTheUpdatablePayloadFromThePlan builds the update (PUT) payload restricted
// to the NITRO-updatable fields. publicip and td are ForceNew/RequiresReplace and
// are NOT accepted by the inat update operation, so they are deliberately excluded.
func inatGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *InatResourceModel) network.Inat {
	tflog.Debug(ctx, "In inatGetTheUpdatablePayloadFromThePlan Function")

	inat := network.Inat{}
	// name is the primary key and is required in the update payload.
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		inat.Name = data.Name.ValueString()
	}
	if !data.Connfailover.IsNull() && !data.Connfailover.IsUnknown() {
		inat.Connfailover = data.Connfailover.ValueString()
	}
	if !data.Ftp.IsNull() && !data.Ftp.IsUnknown() {
		inat.Ftp = data.Ftp.ValueString()
	}
	if !data.Mode.IsNull() && !data.Mode.IsUnknown() {
		inat.Mode = data.Mode.ValueString()
	}
	if !data.Privateip.IsNull() && !data.Privateip.IsUnknown() {
		inat.Privateip = data.Privateip.ValueString()
	}
	if !data.Proxyip.IsNull() && !data.Proxyip.IsUnknown() {
		inat.Proxyip = data.Proxyip.ValueString()
	}
	if !data.Tcpproxy.IsNull() && !data.Tcpproxy.IsUnknown() {
		inat.Tcpproxy = data.Tcpproxy.ValueString()
	}
	if !data.Tftp.IsNull() && !data.Tftp.IsUnknown() {
		inat.Tftp = data.Tftp.ValueString()
	}
	if !data.Useproxyport.IsNull() && !data.Useproxyport.IsUnknown() {
		inat.Useproxyport = data.Useproxyport.ValueString()
	}
	if !data.Usip.IsNull() && !data.Usip.IsUnknown() {
		inat.Usip = data.Usip.ValueString()
	}
	if !data.Usnip.IsNull() && !data.Usnip.IsUnknown() {
		inat.Usnip = data.Usnip.ValueString()
	}

	return inat
}

func inatSetAttrFromGet(ctx context.Context, data *InatResourceModel, getResponseData map[string]interface{}) *InatResourceModel {
	tflog.Debug(ctx, "In inatSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["connfailover"]; ok && val != nil {
		data.Connfailover = types.StringValue(val.(string))
	} else {
		data.Connfailover = types.StringNull()
	}
	if val, ok := getResponseData["ftp"]; ok && val != nil {
		data.Ftp = types.StringValue(val.(string))
	} else {
		data.Ftp = types.StringNull()
	}
	if val, ok := getResponseData["mode"]; ok && val != nil {
		data.Mode = types.StringValue(val.(string))
	} else {
		data.Mode = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["privateip"]; ok && val != nil {
		data.Privateip = types.StringValue(val.(string))
	} else {
		data.Privateip = types.StringNull()
	}
	if val, ok := getResponseData["proxyip"]; ok && val != nil {
		data.Proxyip = types.StringValue(val.(string))
	} else {
		data.Proxyip = types.StringNull()
	}
	if val, ok := getResponseData["publicip"]; ok && val != nil {
		data.Publicip = types.StringValue(val.(string))
	} else {
		data.Publicip = types.StringNull()
	}
	if val, ok := getResponseData["tcpproxy"]; ok && val != nil {
		data.Tcpproxy = types.StringValue(val.(string))
	} else {
		data.Tcpproxy = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else if data.Td.IsUnknown() {
		// td default (0) is omitted from GET; only null-out when unresolved (plan
		// unknown). A configured td=0 held in state is preserved to avoid an
		// "inconsistent result after apply" error.
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["tftp"]; ok && val != nil {
		data.Tftp = types.StringValue(val.(string))
	} else {
		data.Tftp = types.StringNull()
	}
	if val, ok := getResponseData["useproxyport"]; ok && val != nil {
		data.Useproxyport = types.StringValue(val.(string))
	} else {
		data.Useproxyport = types.StringNull()
	}
	if val, ok := getResponseData["usip"]; ok && val != nil {
		data.Usip = types.StringValue(val.(string))
	} else {
		data.Usip = types.StringNull()
	}
	if val, ok := getResponseData["usnip"]; ok && val != nil {
		data.Usnip = types.StringValue(val.(string))
	} else {
		data.Usnip = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID (matches SDK v2 d.SetId(name))
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
