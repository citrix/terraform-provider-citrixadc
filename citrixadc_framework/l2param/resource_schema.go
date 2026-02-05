package l2param

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// L2paramResourceModel describes the resource data model.
type L2paramResourceModel struct {
	Id                      types.String `tfsdk:"id"`
	Bdggrpproxyarp          types.String `tfsdk:"bdggrpproxyarp"`
	Bdgsetting              types.String `tfsdk:"bdgsetting"`
	Bridgeagetimeout        types.Int64  `tfsdk:"bridgeagetimeout"`
	Garponvridintf          types.String `tfsdk:"garponvridintf"`
	Garpreply               types.String `tfsdk:"garpreply"`
	Macmodefwdmypkt         types.String `tfsdk:"macmodefwdmypkt"`
	Maxbridgecollision      types.Int64  `tfsdk:"maxbridgecollision"`
	Mbfinstlearning         types.String `tfsdk:"mbfinstlearning"`
	Mbfpeermacupdate        types.Int64  `tfsdk:"mbfpeermacupdate"`
	Proxyarp                types.String `tfsdk:"proxyarp"`
	Returntoethernetsender  types.String `tfsdk:"returntoethernetsender"`
	Rstintfonhafo           types.String `tfsdk:"rstintfonhafo"`
	Skipproxyingbsdtraffic  types.String `tfsdk:"skipproxyingbsdtraffic"`
	Stopmacmoveupdate       types.String `tfsdk:"stopmacmoveupdate"`
	Usemymac                types.String `tfsdk:"usemymac"`
	Usenetprofilebsdtraffic types.String `tfsdk:"usenetprofilebsdtraffic"`
}

func (r *L2paramResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the l2param resource.",
			},
			"bdggrpproxyarp": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Set/reset proxy ARP in bridge group deployment",
			},
			"bdgsetting": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Bridging settings for C2C behavior. If enabled, each PE will learn MAC entries independently. Otherwise, when L2 mode is ON, learned MAC entries on a PE will be broadcasted to all other PEs.",
			},
			"bridgeagetimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(300),
				Description: "Time-out value for the bridge table entries, in seconds. The new value applies only to the entries that are dynamically learned after the new value is set. Previously existing bridge table entries expire after the previously configured time-out value.",
			},
			"garponvridintf": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Send GARP messagess on VRID-configured interfaces upon failover",
			},
			"garpreply": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Set/reset REPLY form of GARP",
			},
			"macmodefwdmypkt": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Allows MAC mode vserver to pick and forward the packets even if it is destined to Citrix ADC owned VIP.",
			},
			"maxbridgecollision": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(20),
				Description: "Maximum bridge collision for loop detection",
			},
			"mbfinstlearning": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable instant learning of MAC changes in MBF mode.",
			},
			"mbfpeermacupdate": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10),
				Description: "When mbf_instant_learning is enabled, learn any changes in peer's MAC after this time interval, which is in 10ms ticks.",
			},
			"proxyarp": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Proxies the ARP as Citrix ADC MAC for FreeBSD.",
			},
			"returntoethernetsender": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Return to ethernet sender.",
			},
			"rstintfonhafo": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable the reset interface upon HA failover.",
			},
			"skipproxyingbsdtraffic": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Control source parameters (IP and Port) for FreeBSD initiated traffic. If Enabled, source parameters are retained. Else proxy the source parameters based on next hop.",
			},
			"stopmacmoveupdate": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Stop Update of server mac change to NAT sessions.",
			},
			"usemymac": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Use Citrix ADC MAC for all outgoing packets.",
			},
			"usenetprofilebsdtraffic": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Control source parameters (IP and Port) for FreeBSD initiated traffic. If enabled proxy the source parameters based on netprofile source ip. If netprofile does not have ip configured, then it will continue to use NSIP as earlier.",
			},
		},
	}
}

func l2paramGetThePayloadFromtheConfig(ctx context.Context, data *L2paramResourceModel) network.L2param {
	tflog.Debug(ctx, "In l2paramGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	l2param := network.L2param{}
	if !data.Bdggrpproxyarp.IsNull() {
		l2param.Bdggrpproxyarp = data.Bdggrpproxyarp.ValueString()
	}
	if !data.Bdgsetting.IsNull() {
		l2param.Bdgsetting = data.Bdgsetting.ValueString()
	}
	if !data.Bridgeagetimeout.IsNull() {
		l2param.Bridgeagetimeout = utils.IntPtr(int(data.Bridgeagetimeout.ValueInt64()))
	}
	if !data.Garponvridintf.IsNull() {
		l2param.Garponvridintf = data.Garponvridintf.ValueString()
	}
	if !data.Garpreply.IsNull() {
		l2param.Garpreply = data.Garpreply.ValueString()
	}
	if !data.Macmodefwdmypkt.IsNull() {
		l2param.Macmodefwdmypkt = data.Macmodefwdmypkt.ValueString()
	}
	if !data.Maxbridgecollision.IsNull() {
		l2param.Maxbridgecollision = utils.IntPtr(int(data.Maxbridgecollision.ValueInt64()))
	}
	if !data.Mbfinstlearning.IsNull() {
		l2param.Mbfinstlearning = data.Mbfinstlearning.ValueString()
	}
	if !data.Mbfpeermacupdate.IsNull() {
		l2param.Mbfpeermacupdate = utils.IntPtr(int(data.Mbfpeermacupdate.ValueInt64()))
	}
	if !data.Proxyarp.IsNull() {
		l2param.Proxyarp = data.Proxyarp.ValueString()
	}
	if !data.Returntoethernetsender.IsNull() {
		l2param.Returntoethernetsender = data.Returntoethernetsender.ValueString()
	}
	if !data.Rstintfonhafo.IsNull() {
		l2param.Rstintfonhafo = data.Rstintfonhafo.ValueString()
	}
	if !data.Skipproxyingbsdtraffic.IsNull() {
		l2param.Skipproxyingbsdtraffic = data.Skipproxyingbsdtraffic.ValueString()
	}
	if !data.Stopmacmoveupdate.IsNull() {
		l2param.Stopmacmoveupdate = data.Stopmacmoveupdate.ValueString()
	}
	if !data.Usemymac.IsNull() {
		l2param.Usemymac = data.Usemymac.ValueString()
	}
	if !data.Usenetprofilebsdtraffic.IsNull() {
		l2param.Usenetprofilebsdtraffic = data.Usenetprofilebsdtraffic.ValueString()
	}

	return l2param
}

func l2paramSetAttrFromGet(ctx context.Context, data *L2paramResourceModel, getResponseData map[string]interface{}) *L2paramResourceModel {
	tflog.Debug(ctx, "In l2paramSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bdggrpproxyarp"]; ok && val != nil {
		data.Bdggrpproxyarp = types.StringValue(val.(string))
	} else {
		data.Bdggrpproxyarp = types.StringNull()
	}
	if val, ok := getResponseData["bdgsetting"]; ok && val != nil {
		data.Bdgsetting = types.StringValue(val.(string))
	} else {
		data.Bdgsetting = types.StringNull()
	}
	if val, ok := getResponseData["bridgeagetimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Bridgeagetimeout = types.Int64Value(intVal)
		}
	} else {
		data.Bridgeagetimeout = types.Int64Null()
	}
	if val, ok := getResponseData["garponvridintf"]; ok && val != nil {
		data.Garponvridintf = types.StringValue(val.(string))
	} else {
		data.Garponvridintf = types.StringNull()
	}
	if val, ok := getResponseData["garpreply"]; ok && val != nil {
		data.Garpreply = types.StringValue(val.(string))
	} else {
		data.Garpreply = types.StringNull()
	}
	if val, ok := getResponseData["macmodefwdmypkt"]; ok && val != nil {
		data.Macmodefwdmypkt = types.StringValue(val.(string))
	} else {
		data.Macmodefwdmypkt = types.StringNull()
	}
	if val, ok := getResponseData["maxbridgecollision"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxbridgecollision = types.Int64Value(intVal)
		}
	} else {
		data.Maxbridgecollision = types.Int64Null()
	}
	if val, ok := getResponseData["mbfinstlearning"]; ok && val != nil {
		data.Mbfinstlearning = types.StringValue(val.(string))
	} else {
		data.Mbfinstlearning = types.StringNull()
	}
	if val, ok := getResponseData["mbfpeermacupdate"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mbfpeermacupdate = types.Int64Value(intVal)
		}
	} else {
		data.Mbfpeermacupdate = types.Int64Null()
	}
	if val, ok := getResponseData["proxyarp"]; ok && val != nil {
		data.Proxyarp = types.StringValue(val.(string))
	} else {
		data.Proxyarp = types.StringNull()
	}
	if val, ok := getResponseData["returntoethernetsender"]; ok && val != nil {
		data.Returntoethernetsender = types.StringValue(val.(string))
	} else {
		data.Returntoethernetsender = types.StringNull()
	}
	if val, ok := getResponseData["rstintfonhafo"]; ok && val != nil {
		data.Rstintfonhafo = types.StringValue(val.(string))
	} else {
		data.Rstintfonhafo = types.StringNull()
	}
	if val, ok := getResponseData["skipproxyingbsdtraffic"]; ok && val != nil {
		data.Skipproxyingbsdtraffic = types.StringValue(val.(string))
	} else {
		data.Skipproxyingbsdtraffic = types.StringNull()
	}
	if val, ok := getResponseData["stopmacmoveupdate"]; ok && val != nil {
		data.Stopmacmoveupdate = types.StringValue(val.(string))
	} else {
		data.Stopmacmoveupdate = types.StringNull()
	}
	if val, ok := getResponseData["usemymac"]; ok && val != nil {
		data.Usemymac = types.StringValue(val.(string))
	} else {
		data.Usemymac = types.StringNull()
	}
	if val, ok := getResponseData["usenetprofilebsdtraffic"]; ok && val != nil {
		data.Usenetprofilebsdtraffic = types.StringValue(val.(string))
	} else {
		data.Usenetprofilebsdtraffic = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("l2param-config")

	return data
}
