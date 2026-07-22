package metricsprofile

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// MetricsprofileResourceModel describes the resource data model.
type MetricsprofileResourceModel struct {
	Id                        types.String `tfsdk:"id"`
	Collector                 types.String `tfsdk:"collector"`
	Metrics                   types.String `tfsdk:"metrics"`
	Metricsauthtoken          types.String `tfsdk:"metricsauthtoken"`
	MetricsauthtokenWo        types.String `tfsdk:"metricsauthtoken_wo"`
	MetricsauthtokenWoVersion types.Int64  `tfsdk:"metricsauthtoken_wo_version"`
	Metricsendpointurl        types.String `tfsdk:"metricsendpointurl"`
	Metricsexportfrequency    types.Int64  `tfsdk:"metricsexportfrequency"`
	Name                      types.String `tfsdk:"name"`
	Outputmode                types.String `tfsdk:"outputmode"`
	Schemafile                types.String `tfsdk:"schemafile"`
	Servemode                 types.String `tfsdk:"servemode"`
}

func (r *MetricsprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the metricsprofile resource.",
			},
			"collector": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The collector should be a HTTP/HTTPS service.",
			},
			"metrics": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is used enable or disable metrics",
			},
			"metricsauthtoken": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Token for authenticating with the endpoint. If the endpoint requires the Authorization header in a particular format, specify the complete format as the value to this parameter. For eg., in case of splunk, the Authorizaiton header is required to be of the form - Splunk <auth-token>.",
			},
			"metricsauthtoken_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Token for authenticating with the endpoint. If the endpoint requires the Authorization header in a particular format, specify the complete format as the value to this parameter. For eg., in case of splunk, the Authorizaiton header is required to be of the form - Splunk <auth-token>.",
			},
			"metricsauthtoken_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a metricsauthtoken_wo update.",
			},
			"metricsendpointurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The URL at which to upload the metrics data on the endpoint",
			},
			"metricsexportfrequency": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is for configuring the metrics export frequency in seconds, frequency value must be in [30,300] seconds range",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the metrics profile. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters.!",
			},
			"outputmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option indicates the format in which metrics data is generated",
			},
			"schemafile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is for configuring json schema file containing a list of counters to be exported by metricscollector. Schema file should be present under /var/metrics_conf path",
			},
			"servemode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is to configure metrics pull or push mode. In push mode metricscollector exports metrics to configured collector. In pull mode, metricscollector only generates the metrics which will be pulled by external agent. No collector configuration is required in pull mode and it is applicable only for output mode Prometheus",
			},
		},
	}
}

func metricsprofileGetThePayloadFromthePlan(ctx context.Context, data *MetricsprofileResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In metricsprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	// No vendored metrics.Metricsprofile struct exists; build a generic payload map.
	metricsprofile := make(map[string]interface{})
	if !data.Collector.IsNull() && !data.Collector.IsUnknown() {
		metricsprofile["collector"] = data.Collector.ValueString()
	}
	if !data.Metrics.IsNull() && !data.Metrics.IsUnknown() {
		metricsprofile["metrics"] = data.Metrics.ValueString()
	}
	if !data.Metricsauthtoken.IsNull() && !data.Metricsauthtoken.IsUnknown() {
		metricsprofile["metricsauthtoken"] = data.Metricsauthtoken.ValueString()
	}
	// Skip write-only attribute: metricsauthtoken_wo
	// Skip version tracker attribute: metricsauthtoken_wo_version
	if !data.Metricsendpointurl.IsNull() && !data.Metricsendpointurl.IsUnknown() {
		metricsprofile["metricsendpointurl"] = data.Metricsendpointurl.ValueString()
	}
	if !data.Metricsexportfrequency.IsNull() && !data.Metricsexportfrequency.IsUnknown() {
		metricsprofile["metricsexportfrequency"] = int(data.Metricsexportfrequency.ValueInt64())
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		metricsprofile["name"] = data.Name.ValueString()
	}
	if !data.Outputmode.IsNull() && !data.Outputmode.IsUnknown() {
		metricsprofile["outputmode"] = data.Outputmode.ValueString()
	}
	if !data.Schemafile.IsNull() && !data.Schemafile.IsUnknown() {
		metricsprofile["schemafile"] = data.Schemafile.ValueString()
	}
	if !data.Servemode.IsNull() && !data.Servemode.IsUnknown() {
		metricsprofile["servemode"] = data.Servemode.ValueString()
	}

	return metricsprofile
}

func metricsprofileGetThePayloadFromtheConfig(ctx context.Context, data *MetricsprofileResourceModel, payload map[string]interface{}) {
	tflog.Debug(ctx, "In metricsprofileGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: metricsauthtoken_wo -> metricsauthtoken
	if !data.MetricsauthtokenWo.IsNull() {
		metricsauthtokenWo := data.MetricsauthtokenWo.ValueString()
		if metricsauthtokenWo != "" {
			payload["metricsauthtoken"] = metricsauthtokenWo
		}
	}
}

func metricsprofileSetAttrFromGet(ctx context.Context, data *MetricsprofileResourceModel, getResponseData map[string]interface{}) *MetricsprofileResourceModel {
	tflog.Debug(ctx, "In metricsprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["collector"]; ok && val != nil {
		data.Collector = types.StringValue(val.(string))
	} else {
		data.Collector = types.StringNull()
	}
	if val, ok := getResponseData["metrics"]; ok && val != nil {
		data.Metrics = types.StringValue(val.(string))
	} else {
		data.Metrics = types.StringNull()
	}
	// metricsauthtoken is not returned by NITRO API (secret/ephemeral) - retain from config
	// metricsauthtoken_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// metricsauthtoken_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["metricsendpointurl"]; ok && val != nil {
		data.Metricsendpointurl = types.StringValue(val.(string))
	} else {
		data.Metricsendpointurl = types.StringNull()
	}
	if val, ok := getResponseData["metricsexportfrequency"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Metricsexportfrequency = types.Int64Value(intVal)
		}
	} else {
		data.Metricsexportfrequency = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["outputmode"]; ok && val != nil {
		data.Outputmode = types.StringValue(val.(string))
	} else {
		data.Outputmode = types.StringNull()
	}
	if val, ok := getResponseData["schemafile"]; ok && val != nil {
		data.Schemafile = types.StringValue(val.(string))
	} else {
		data.Schemafile = types.StringNull()
	}
	if val, ok := getResponseData["servemode"]; ok && val != nil {
		data.Servemode = types.StringValue(val.(string))
	} else {
		data.Servemode = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
