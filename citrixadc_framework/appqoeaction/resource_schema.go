package appqoeaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appqoe"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppqoeactionResourceModel describes the resource data model.
type AppqoeactionResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Altcontentpath    types.String `tfsdk:"altcontentpath"`
	Altcontentsvcname types.String `tfsdk:"altcontentsvcname"`
	Customfile        types.String `tfsdk:"customfile"`
	Delay             types.Int64  `tfsdk:"delay"`
	Dosaction         types.String `tfsdk:"dosaction"`
	Dostrigexpression types.String `tfsdk:"dostrigexpression"`
	Maxconn           types.Int64  `tfsdk:"maxconn"`
	Name              types.String `tfsdk:"name"`
	Numretries        types.Int64  `tfsdk:"numretries"`
	Polqdepth         types.Int64  `tfsdk:"polqdepth"`
	Priority          types.String `tfsdk:"priority"`
	Priqdepth         types.Int64  `tfsdk:"priqdepth"`
	Respondwith       types.String `tfsdk:"respondwith"`
	Retryonreset      types.String `tfsdk:"retryonreset"`
	Retryontimeout    types.Int64  `tfsdk:"retryontimeout"`
	Tcpprofile        types.String `tfsdk:"tcpprofile"`
}

func (r *AppqoeactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appqoeaction resource.",
			},
			"altcontentpath": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Path to the alternative content service to be used in the ACS",
			},
			"altcontentsvcname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the alternative content service to be used in the ACS",
			},
			"customfile": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "name of the HTML page object to use as the response",
			},
			"delay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Delay threshold, in microseconds, for requests that match the policy's rule. If the delay statistics gathered for the matching request exceed the specified delay, configured action triggered for that request, if there is no action then requests are dropped to the lowest priority level",
			},
			"dosaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DoS Action to take when vserver will be considered under DoS attack and corresponding rule matches. Mandatory if AppQoE actions are to be used for DoS attack prevention.",
			},
			"dostrigexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Optional expression to add second level check to trigger DoS actions. Specifically used for Analytics based DoS response generation",
			},
			"maxconn": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of concurrent connections that can be open for requests that matches with rule.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the AppQoE action. Must begin with a letter, number, or the underscore symbol (_). Other characters allowed, after the first character, are the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), and colon (:) characters. This is a mandatory argument",
			},
			"numretries": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3),
				Description: "Retry count",
			},
			"polqdepth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Policy queue depth threshold value. When the policy queue size (number of requests queued for the policy binding this action is attached to) increases to the specified polqDepth value, subsequent requests are dropped to the lowest priority level.",
			},
			"priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority for queuing the request. If server resources are not available for a request that matches the configured rule, this option specifies a priority for queuing the request until the server resources are available again. If priority is not configured then Lowest priority will be used to queue the request.",
			},
			"priqdepth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Queue depth threshold value per priorirty level. If the queue size (number of requests in the queue of that particular priorirty) on the virtual server to which this policy is bound, increases to the specified qDepth value, subsequent requests are dropped to the lowest priority level.",
			},
			"respondwith": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Responder action to be taken when the threshold is reached. Available settings function as follows:\n            ACS - Serve content from an alternative content service\n                  Threshold : maxConn or delay\n            NS - Serve from the Citrix ADC (built-in response)\n                 Threshold : maxConn or delay",
			},
			"retryonreset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Retry on TCP Reset",
			},
			"retryontimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Retry on request Timeout(in millisec) upon sending request to backend servers",
			},
			"tcpprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bind TCP Profile based on L2/L3/L7 parameters.",
			},
		},
	}
}

func appqoeactionGetThePayloadFromtheConfig(ctx context.Context, data *AppqoeactionResourceModel) appqoe.Appqoeaction {
	tflog.Debug(ctx, "In appqoeactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appqoeaction := appqoe.Appqoeaction{}
	if !data.Altcontentpath.IsNull() {
		appqoeaction.Altcontentpath = data.Altcontentpath.ValueString()
	}
	if !data.Altcontentsvcname.IsNull() {
		appqoeaction.Altcontentsvcname = data.Altcontentsvcname.ValueString()
	}
	if !data.Customfile.IsNull() {
		appqoeaction.Customfile = data.Customfile.ValueString()
	}
	if !data.Delay.IsNull() {
		appqoeaction.Delay = utils.IntPtr(int(data.Delay.ValueInt64()))
	}
	if !data.Dosaction.IsNull() {
		appqoeaction.Dosaction = data.Dosaction.ValueString()
	}
	if !data.Dostrigexpression.IsNull() {
		appqoeaction.Dostrigexpression = data.Dostrigexpression.ValueString()
	}
	if !data.Maxconn.IsNull() {
		appqoeaction.Maxconn = utils.IntPtr(int(data.Maxconn.ValueInt64()))
	}
	if !data.Name.IsNull() {
		appqoeaction.Name = data.Name.ValueString()
	}
	if !data.Numretries.IsNull() {
		appqoeaction.Numretries = utils.IntPtr(int(data.Numretries.ValueInt64()))
	}
	if !data.Polqdepth.IsNull() {
		appqoeaction.Polqdepth = utils.IntPtr(int(data.Polqdepth.ValueInt64()))
	}
	if !data.Priority.IsNull() {
		appqoeaction.Priority = data.Priority.ValueString()
	}
	if !data.Priqdepth.IsNull() {
		appqoeaction.Priqdepth = utils.IntPtr(int(data.Priqdepth.ValueInt64()))
	}
	if !data.Respondwith.IsNull() {
		appqoeaction.Respondwith = data.Respondwith.ValueString()
	}
	if !data.Retryonreset.IsNull() {
		appqoeaction.Retryonreset = data.Retryonreset.ValueString()
	}
	if !data.Retryontimeout.IsNull() {
		appqoeaction.Retryontimeout = utils.IntPtr(int(data.Retryontimeout.ValueInt64()))
	}
	if !data.Tcpprofile.IsNull() {
		appqoeaction.Tcpprofile = data.Tcpprofile.ValueString()
	}

	return appqoeaction
}

func appqoeactionSetAttrFromGet(ctx context.Context, data *AppqoeactionResourceModel, getResponseData map[string]interface{}) *AppqoeactionResourceModel {
	tflog.Debug(ctx, "In appqoeactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["altcontentpath"]; ok && val != nil {
		data.Altcontentpath = types.StringValue(val.(string))
	} else {
		data.Altcontentpath = types.StringNull()
	}
	if val, ok := getResponseData["altcontentsvcname"]; ok && val != nil {
		data.Altcontentsvcname = types.StringValue(val.(string))
	} else {
		data.Altcontentsvcname = types.StringNull()
	}
	if val, ok := getResponseData["customfile"]; ok && val != nil {
		data.Customfile = types.StringValue(val.(string))
	} else {
		data.Customfile = types.StringNull()
	}
	if val, ok := getResponseData["delay"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Delay = types.Int64Value(intVal)
		}
	} else {
		data.Delay = types.Int64Null()
	}
	if val, ok := getResponseData["dosaction"]; ok && val != nil {
		data.Dosaction = types.StringValue(val.(string))
	} else {
		data.Dosaction = types.StringNull()
	}
	if val, ok := getResponseData["dostrigexpression"]; ok && val != nil {
		data.Dostrigexpression = types.StringValue(val.(string))
	} else {
		data.Dostrigexpression = types.StringNull()
	}
	if val, ok := getResponseData["maxconn"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxconn = types.Int64Value(intVal)
		}
	} else {
		data.Maxconn = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["numretries"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Numretries = types.Int64Value(intVal)
		}
	} else {
		data.Numretries = types.Int64Null()
	}
	if val, ok := getResponseData["polqdepth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Polqdepth = types.Int64Value(intVal)
		}
	} else {
		data.Polqdepth = types.Int64Null()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		data.Priority = types.StringValue(val.(string))
	} else {
		data.Priority = types.StringNull()
	}
	if val, ok := getResponseData["priqdepth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priqdepth = types.Int64Value(intVal)
		}
	} else {
		data.Priqdepth = types.Int64Null()
	}
	if val, ok := getResponseData["respondwith"]; ok && val != nil {
		data.Respondwith = types.StringValue(val.(string))
	} else {
		data.Respondwith = types.StringNull()
	}
	if val, ok := getResponseData["retryonreset"]; ok && val != nil {
		data.Retryonreset = types.StringValue(val.(string))
	} else {
		data.Retryonreset = types.StringNull()
	}
	if val, ok := getResponseData["retryontimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Retryontimeout = types.Int64Value(intVal)
		}
	} else {
		data.Retryontimeout = types.Int64Null()
	}
	if val, ok := getResponseData["tcpprofile"]; ok && val != nil {
		data.Tcpprofile = types.StringValue(val.(string))
	} else {
		data.Tcpprofile = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
