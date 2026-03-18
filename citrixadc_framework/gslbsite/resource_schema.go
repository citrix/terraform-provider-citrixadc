package gslbsite

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// GslbsiteResourceModel describes the resource data model.
type GslbsiteResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Backupparentlist       types.List   `tfsdk:"backupparentlist"`
	Clip                   types.String `tfsdk:"clip"`
	Metricexchange         types.String `tfsdk:"metricexchange"`
	Naptrreplacementsuffix types.String `tfsdk:"naptrreplacementsuffix"`
	Newname                types.String `tfsdk:"newname"`
	Nwmetricexchange       types.String `tfsdk:"nwmetricexchange"`
	Parentsite             types.String `tfsdk:"parentsite"`
	Publicclip             types.String `tfsdk:"publicclip"`
	Publicip               types.String `tfsdk:"publicip"`
	Sessionexchange        types.String `tfsdk:"sessionexchange"`
	Siteipaddress          types.String `tfsdk:"siteipaddress"`
	Sitename               types.String `tfsdk:"sitename"`
	Sitepassword           types.String `tfsdk:"sitepassword"`
	Sitetype               types.String `tfsdk:"sitetype"`
	Triggermonitor         types.String `tfsdk:"triggermonitor"`
}

func (r *GslbsiteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the gslbsite resource.",
			},
			"backupparentlist": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "The list of backup gslb sites configured in preferred order. Need to be parent gsb sites.",
			},
			"clip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Cluster IP address. Specify this parameter to connect to the remote cluster site for GSLB auto-sync. Note: The cluster IP address is defined when creating the cluster.",
			},
			"metricexchange": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Exchange metrics with other sites. Metrics are exchanged by using Metric Exchange Protocol (MEP). The appliances in the GSLB setup exchange health information once every second.\n\nIf you disable metrics exchange, you can use only static load balancing methods (such as round robin, static proximity, or the hash-based methods), and if you disable metrics exchange when a dynamic load balancing method (such as least connection) is in operation, the appliance falls back to round robin. Also, if you disable metrics exchange, you must use a monitor to determine the state of GSLB services. Otherwise, the service is marked as DOWN.",
			},
			"naptrreplacementsuffix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The naptr replacement suffix configured here will be used to construct the naptr replacement field in NAPTR record.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the GSLB site.",
			},
			"nwmetricexchange": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Exchange, with other GSLB sites, network metrics such as round-trip time (RTT), learned from communications with various local DNS (LDNS) servers used by clients. RTT information is used in the dynamic RTT load balancing method, and is exchanged every 5 seconds.",
			},
			"parentsite": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parent site of the GSLB site, in a parent-child topology.",
			},
			"publicclip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP address to be used to globally access the remote cluster when it is deployed behind a NAT. It can be same as the normal cluster IP address.",
			},
			"publicip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Public IP address for the local site. Required only if the appliance is deployed in a private address space and the site has a public IP address hosted on an external firewall or a NAT device.",
			},
			"sessionexchange": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Exchange persistent session entries with other GSLB sites every five seconds.",
			},
			"siteipaddress": schema.StringAttribute{
				Required:    true,
				Description: "IP address for the GSLB site. The GSLB site uses this IP address to communicate with other GSLB sites. For a local site, use any IP address that is owned by the appliance (for example, a SNIP or MIP address, or the IP address of the ADNS service).",
			},
			"sitename": schema.StringAttribute{
				Required:    true,
				Description: "Name for the GSLB site. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the virtual server is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my gslbsite\" or 'my gslbsite').",
			},
			"sitepassword": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Password to be used for mep communication between gslb site nodes.",
			},
			"sitetype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("NONE"),
				Description: "Type of site to create. If the type is not specified, the appliance automatically detects and sets the type on the basis of the IP address being assigned to the site. If the specified site IP address is owned by the appliance (for example, a MIP address or SNIP address), the site is a local site. Otherwise, it is a remote site.",
			},
			"triggermonitor": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ALWAYS"),
				Description: "Specify the conditions under which the GSLB service must be monitored by a monitor, if one is bound. Available settings function as follows:\n* ALWAYS - Monitor the GSLB service at all times.\n* MEPDOWN - Monitor the GSLB service only when the exchange of metrics through the Metrics Exchange Protocol (MEP) is disabled.\nMEPDOWN_SVCDOWN - Monitor the service in either of the following situations:\n* The exchange of metrics through MEP is disabled.\n* The exchange of metrics through MEP is enabled but the status of the service, learned through metrics exchange, is DOWN.",
			},
		},
	}
}

func gslbsiteGetThePayloadFromtheConfig(ctx context.Context, data *GslbsiteResourceModel) gslb.Gslbsite {
	tflog.Debug(ctx, "In gslbsiteGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	gslbsite := gslb.Gslbsite{}
	if !data.Clip.IsNull() {
		gslbsite.Clip = data.Clip.ValueString()
	}
	if !data.Metricexchange.IsNull() {
		gslbsite.Metricexchange = data.Metricexchange.ValueString()
	}
	if !data.Naptrreplacementsuffix.IsNull() {
		gslbsite.Naptrreplacementsuffix = data.Naptrreplacementsuffix.ValueString()
	}
	if !data.Newname.IsNull() {
		gslbsite.Newname = data.Newname.ValueString()
	}
	if !data.Nwmetricexchange.IsNull() {
		gslbsite.Nwmetricexchange = data.Nwmetricexchange.ValueString()
	}
	if !data.Parentsite.IsNull() {
		gslbsite.Parentsite = data.Parentsite.ValueString()
	}
	if !data.Publicclip.IsNull() {
		gslbsite.Publicclip = data.Publicclip.ValueString()
	}
	if !data.Publicip.IsNull() {
		gslbsite.Publicip = data.Publicip.ValueString()
	}
	if !data.Sessionexchange.IsNull() {
		gslbsite.Sessionexchange = data.Sessionexchange.ValueString()
	}
	if !data.Siteipaddress.IsNull() {
		gslbsite.Siteipaddress = data.Siteipaddress.ValueString()
	}
	if !data.Sitename.IsNull() {
		gslbsite.Sitename = data.Sitename.ValueString()
	}
	if !data.Sitepassword.IsNull() {
		gslbsite.Sitepassword = data.Sitepassword.ValueString()
	}
	if !data.Sitetype.IsNull() {
		gslbsite.Sitetype = data.Sitetype.ValueString()
	}
	if !data.Triggermonitor.IsNull() {
		gslbsite.Triggermonitor = data.Triggermonitor.ValueString()
	}

	return gslbsite
}

func gslbsiteSetAttrFromGet(ctx context.Context, data *GslbsiteResourceModel, getResponseData map[string]interface{}) *GslbsiteResourceModel {
	tflog.Debug(ctx, "In gslbsiteSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["clip"]; ok && val != nil {
		data.Clip = types.StringValue(val.(string))
	} else {
		data.Clip = types.StringNull()
	}
	if val, ok := getResponseData["metricexchange"]; ok && val != nil {
		data.Metricexchange = types.StringValue(val.(string))
	} else {
		data.Metricexchange = types.StringNull()
	}
	if val, ok := getResponseData["naptrreplacementsuffix"]; ok && val != nil {
		data.Naptrreplacementsuffix = types.StringValue(val.(string))
	} else {
		data.Naptrreplacementsuffix = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["nwmetricexchange"]; ok && val != nil {
		data.Nwmetricexchange = types.StringValue(val.(string))
	} else {
		data.Nwmetricexchange = types.StringNull()
	}
	if val, ok := getResponseData["parentsite"]; ok && val != nil {
		data.Parentsite = types.StringValue(val.(string))
	} else {
		data.Parentsite = types.StringNull()
	}
	if val, ok := getResponseData["publicclip"]; ok && val != nil {
		data.Publicclip = types.StringValue(val.(string))
	} else {
		data.Publicclip = types.StringNull()
	}
	if val, ok := getResponseData["publicip"]; ok && val != nil {
		data.Publicip = types.StringValue(val.(string))
	} else {
		data.Publicip = types.StringNull()
	}
	if val, ok := getResponseData["sessionexchange"]; ok && val != nil {
		data.Sessionexchange = types.StringValue(val.(string))
	} else {
		data.Sessionexchange = types.StringNull()
	}
	if val, ok := getResponseData["siteipaddress"]; ok && val != nil {
		data.Siteipaddress = types.StringValue(val.(string))
	} else {
		data.Siteipaddress = types.StringNull()
	}
	if val, ok := getResponseData["sitename"]; ok && val != nil {
		data.Sitename = types.StringValue(val.(string))
	} else {
		data.Sitename = types.StringNull()
	}
	if val, ok := getResponseData["sitepassword"]; ok && val != nil {
		data.Sitepassword = types.StringValue(val.(string))
	} else {
		data.Sitepassword = types.StringNull()
	}
	if val, ok := getResponseData["sitetype"]; ok && val != nil {
		data.Sitetype = types.StringValue(val.(string))
	} else {
		data.Sitetype = types.StringNull()
	}
	if val, ok := getResponseData["triggermonitor"]; ok && val != nil {
		data.Triggermonitor = types.StringValue(val.(string))
	} else {
		data.Triggermonitor = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Sitename.ValueString())

	return data
}
