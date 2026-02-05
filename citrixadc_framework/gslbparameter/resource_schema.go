package gslbparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// GslbparameterResourceModel describes the resource data model.
type GslbparameterResourceModel struct {
	Id                        types.String `tfsdk:"id"`
	Automaticconfigsync       types.String `tfsdk:"automaticconfigsync"`
	Dropldnsreq               types.String `tfsdk:"dropldnsreq"`
	Gslbconfigsyncmonitor     types.String `tfsdk:"gslbconfigsyncmonitor"`
	Gslbsvcstatedelaytime     types.Int64  `tfsdk:"gslbsvcstatedelaytime"`
	Gslbsyncinterval          types.Int64  `tfsdk:"gslbsyncinterval"`
	Gslbsynclocfiles          types.String `tfsdk:"gslbsynclocfiles"`
	Gslbsyncmode              types.String `tfsdk:"gslbsyncmode"`
	Gslbsyncsaveconfigcommand types.String `tfsdk:"gslbsyncsaveconfigcommand"`
	Ldnsentrytimeout          types.Int64  `tfsdk:"ldnsentrytimeout"`
	Ldnsmask                  types.String `tfsdk:"ldnsmask"`
	Ldnsprobeorder            types.List   `tfsdk:"ldnsprobeorder"`
	Mepkeepalivetimeout       types.Int64  `tfsdk:"mepkeepalivetimeout"`
	Rtttolerance              types.Int64  `tfsdk:"rtttolerance"`
	Svcstatelearningtime      types.Int64  `tfsdk:"svcstatelearningtime"`
	Undefaction               types.String `tfsdk:"undefaction"`
	V6ldnsmasklen             types.Int64  `tfsdk:"v6ldnsmasklen"`
}

func (r *GslbparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the gslbparameter resource.",
			},
			"automaticconfigsync": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "GSLB configuration will be synced automatically to remote gslb sites if enabled.",
			},
			"dropldnsreq": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Drop LDNS requests if round-trip time (RTT) information is not available.",
			},
			"gslbconfigsyncmonitor": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "If enabled, remote gslb site's rsync port will be monitored and site is considered for configuration sync only when the monitor is successful.",
			},
			"gslbsvcstatedelaytime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Amount of delay in updating the state of GSLB service to DOWN when MEP goes down.\n			This parameter is applicable only if monitors are not bound to GSLB services",
			},
			"gslbsyncinterval": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10),
				Description: "Time duartion (in seconds) for which the gslb sync process will wait before checking for config changes.",
			},
			"gslbsynclocfiles": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "If disabled, Location files will not be synced to the remote sites as part of manual sync and automatic sync.",
			},
			"gslbsyncmode": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("IncrementalSync"),
				Description: "Mode in which configuration will be synced from master site to remote sites.",
			},
			"gslbsyncsaveconfigcommand": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "If enabled, 'save ns config' command will be treated as other GSLB commands and synced to GSLB nodes when auto gslb sync option is enabled.",
			},
			"ldnsentrytimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(180),
				Description: "Time, in seconds, after which an inactive LDNS entry is removed.",
			},
			"ldnsmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The IPv4 network mask with which to create LDNS entries.",
			},
			"ldnsprobeorder": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Order in which monitors should be initiated to calculate RTT.",
			},
			"mepkeepalivetimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10),
				Description: "Time duartion (in seconds) during which if no new packets received by Local gslb site from Remote gslb site then mark the MEP connection DOWN",
			},
			"rtttolerance": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(5),
				Description: "Tolerance, in milliseconds, for newly learned round-trip time (RTT) values. If the difference between the old RTT value and the newly computed RTT value is less than or equal to the specified tolerance value, the LDNS entry in the network metric table is not updated with the new RTT value. Prevents the exchange of metrics when variations in RTT values are negligible.",
			},
			"svcstatelearningtime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time (in seconds) within which local or child site services remain in learning phase. GSLB site will enter the learning phase after reboot, HA failover, Cluster GSLB owner node changes or MEP being enabled on local node.  Backup parent (if configured) will selectively move the adopted children's GSLB services to learning phase when primary parent goes down. While a service is in learning period, remote site will not honour the state and stats got through MEP for that service. State can be learnt from health monitor if bound explicitly.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("NOLBACTION"),
				Description: "Action to perform when policy evaluation creates an UNDEF condition. Available settings function as follows:\n* NOLBACTION - Does not consider LB action in making LB decision.\n* RESET - Reset the request and notify the user, so that the user can resend the request.\n* DROP - Drop the request without sending a response to the user.",
			},
			"v6ldnsmasklen": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(128),
				Description: "Mask for creating LDNS entries for IPv6 source addresses. The mask is defined as the number of leading bits to consider, in the source IP address, when creating an LDNS entry.",
			},
		},
	}
}

func gslbparameterGetThePayloadFromtheConfig(ctx context.Context, data *GslbparameterResourceModel) gslb.Gslbparameter {
	tflog.Debug(ctx, "In gslbparameterGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	gslbparameter := gslb.Gslbparameter{}
	if !data.Automaticconfigsync.IsNull() {
		gslbparameter.Automaticconfigsync = data.Automaticconfigsync.ValueString()
	}
	if !data.Dropldnsreq.IsNull() {
		gslbparameter.Dropldnsreq = data.Dropldnsreq.ValueString()
	}
	if !data.Gslbconfigsyncmonitor.IsNull() {
		gslbparameter.Gslbconfigsyncmonitor = data.Gslbconfigsyncmonitor.ValueString()
	}
	if !data.Gslbsvcstatedelaytime.IsNull() {
		gslbparameter.Gslbsvcstatedelaytime = utils.IntPtr(int(data.Gslbsvcstatedelaytime.ValueInt64()))
	}
	if !data.Gslbsyncinterval.IsNull() {
		gslbparameter.Gslbsyncinterval = utils.IntPtr(int(data.Gslbsyncinterval.ValueInt64()))
	}
	if !data.Gslbsynclocfiles.IsNull() {
		gslbparameter.Gslbsynclocfiles = data.Gslbsynclocfiles.ValueString()
	}
	if !data.Gslbsyncmode.IsNull() {
		gslbparameter.Gslbsyncmode = data.Gslbsyncmode.ValueString()
	}
	if !data.Gslbsyncsaveconfigcommand.IsNull() {
		gslbparameter.Gslbsyncsaveconfigcommand = data.Gslbsyncsaveconfigcommand.ValueString()
	}
	if !data.Ldnsentrytimeout.IsNull() {
		gslbparameter.Ldnsentrytimeout = utils.IntPtr(int(data.Ldnsentrytimeout.ValueInt64()))
	}
	if !data.Ldnsmask.IsNull() {
		gslbparameter.Ldnsmask = data.Ldnsmask.ValueString()
	}
	if !data.Mepkeepalivetimeout.IsNull() {
		gslbparameter.Mepkeepalivetimeout = utils.IntPtr(int(data.Mepkeepalivetimeout.ValueInt64()))
	}
	if !data.Rtttolerance.IsNull() {
		gslbparameter.Rtttolerance = utils.IntPtr(int(data.Rtttolerance.ValueInt64()))
	}
	if !data.Svcstatelearningtime.IsNull() {
		gslbparameter.Svcstatelearningtime = utils.IntPtr(int(data.Svcstatelearningtime.ValueInt64()))
	}
	if !data.Undefaction.IsNull() {
		gslbparameter.Undefaction = data.Undefaction.ValueString()
	}
	if !data.V6ldnsmasklen.IsNull() {
		gslbparameter.V6ldnsmasklen = utils.IntPtr(int(data.V6ldnsmasklen.ValueInt64()))
	}

	return gslbparameter
}

func gslbparameterSetAttrFromGet(ctx context.Context, data *GslbparameterResourceModel, getResponseData map[string]interface{}) *GslbparameterResourceModel {
	tflog.Debug(ctx, "In gslbparameterSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["automaticconfigsync"]; ok && val != nil {
		data.Automaticconfigsync = types.StringValue(val.(string))
	} else {
		data.Automaticconfigsync = types.StringNull()
	}
	if val, ok := getResponseData["dropldnsreq"]; ok && val != nil {
		data.Dropldnsreq = types.StringValue(val.(string))
	} else {
		data.Dropldnsreq = types.StringNull()
	}
	if val, ok := getResponseData["gslbconfigsyncmonitor"]; ok && val != nil {
		data.Gslbconfigsyncmonitor = types.StringValue(val.(string))
	} else {
		data.Gslbconfigsyncmonitor = types.StringNull()
	}
	if val, ok := getResponseData["gslbsvcstatedelaytime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Gslbsvcstatedelaytime = types.Int64Value(intVal)
		}
	} else {
		data.Gslbsvcstatedelaytime = types.Int64Null()
	}
	if val, ok := getResponseData["gslbsyncinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Gslbsyncinterval = types.Int64Value(intVal)
		}
	} else {
		data.Gslbsyncinterval = types.Int64Null()
	}
	if val, ok := getResponseData["gslbsynclocfiles"]; ok && val != nil {
		data.Gslbsynclocfiles = types.StringValue(val.(string))
	} else {
		data.Gslbsynclocfiles = types.StringNull()
	}
	if val, ok := getResponseData["gslbsyncmode"]; ok && val != nil {
		data.Gslbsyncmode = types.StringValue(val.(string))
	} else {
		data.Gslbsyncmode = types.StringNull()
	}
	if val, ok := getResponseData["gslbsyncsaveconfigcommand"]; ok && val != nil {
		data.Gslbsyncsaveconfigcommand = types.StringValue(val.(string))
	} else {
		data.Gslbsyncsaveconfigcommand = types.StringNull()
	}
	if val, ok := getResponseData["ldnsentrytimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ldnsentrytimeout = types.Int64Value(intVal)
		}
	} else {
		data.Ldnsentrytimeout = types.Int64Null()
	}
	if val, ok := getResponseData["ldnsmask"]; ok && val != nil {
		data.Ldnsmask = types.StringValue(val.(string))
	} else {
		data.Ldnsmask = types.StringNull()
	}
	if val, ok := getResponseData["mepkeepalivetimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mepkeepalivetimeout = types.Int64Value(intVal)
		}
	} else {
		data.Mepkeepalivetimeout = types.Int64Null()
	}
	if val, ok := getResponseData["rtttolerance"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Rtttolerance = types.Int64Value(intVal)
		}
	} else {
		data.Rtttolerance = types.Int64Null()
	}
	if val, ok := getResponseData["svcstatelearningtime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Svcstatelearningtime = types.Int64Value(intVal)
		}
	} else {
		data.Svcstatelearningtime = types.Int64Null()
	}
	if val, ok := getResponseData["undefaction"]; ok && val != nil {
		data.Undefaction = types.StringValue(val.(string))
	} else {
		data.Undefaction = types.StringNull()
	}
	if val, ok := getResponseData["v6ldnsmasklen"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.V6ldnsmasklen = types.Int64Value(intVal)
		}
	} else {
		data.V6ldnsmasklen = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("gslbparameter-config")

	return data
}
