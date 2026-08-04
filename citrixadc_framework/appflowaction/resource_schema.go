package appflowaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appflow"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppflowactionResourceModel describes the resource data model.
type AppflowactionResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Botinsight             types.String `tfsdk:"botinsight"`
	Ciinsight              types.String `tfsdk:"ciinsight"`
	Clientsidemeasurements types.String `tfsdk:"clientsidemeasurements"`
	Collectors             types.List   `tfsdk:"collectors"`
	Comment                types.String `tfsdk:"comment"`
	Distributionalgorithm  types.String `tfsdk:"distributionalgorithm"`
	Metricslog             types.Bool   `tfsdk:"metricslog"`
	Name                   types.String `tfsdk:"name"`
	Newname                types.String `tfsdk:"newname"`
	Pagetracking           types.String `tfsdk:"pagetracking"`
	Securityinsight        types.String `tfsdk:"securityinsight"`
	Transactionlog         types.String `tfsdk:"transactionlog"`
	Videoanalytics         types.String `tfsdk:"videoanalytics"`
	Webinsight             types.String `tfsdk:"webinsight"`
}

func (r *AppflowactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appflowaction resource.",
			},
			"botinsight": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will send the bot insight records to the configured collectors.",
			},
			"ciinsight": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will send the ContentInspection Insight records to the configured collectors.",
			},
			"clientsidemeasurements": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will collect the time required to load and render the mainpage on the client.",
			},
			"collectors": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Name(s) of collector(s) to be associated with the AppFlow action.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about this action.  In the CLI, if including spaces between words, enclose the comment in quotation marks. (The quotation marks are not required in the configuration utility.)",
			},
			"distributionalgorithm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will distribute records among the collectors. Else, all records will be sent to all the collectors.",
			},
			"metricslog": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "If only the stats records are to be exported, turn on this option.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow action\" or 'my appflow action').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the AppFlow action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow action\" or 'my appflow action').",
			},
			"pagetracking": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will start tracking the page for waterfall chart by inserting a NS_ESNS cookie in the response.",
			},
			"securityinsight": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will send the security insight records to the configured collectors.",
			},
			"transactionlog": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Log ANOMALOUS or ALL transactions",
			},
			"videoanalytics": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will send the videoinsight records to the configured collectors.",
			},
			"webinsight": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will send the webinsight records to the configured collectors.",
			},
		},
	}
}

func appflowactionGetThePayloadFromthePlan(ctx context.Context, data *AppflowactionResourceModel) appflow.Appflowaction {
	tflog.Debug(ctx, "In appflowactionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appflowaction := appflow.Appflowaction{}
	if !data.Botinsight.IsNull() && !data.Botinsight.IsUnknown() {
		appflowaction.Botinsight = data.Botinsight.ValueString()
	}
	if !data.Ciinsight.IsNull() && !data.Ciinsight.IsUnknown() {
		appflowaction.Ciinsight = data.Ciinsight.ValueString()
	}
	if !data.Clientsidemeasurements.IsNull() && !data.Clientsidemeasurements.IsUnknown() {
		appflowaction.Clientsidemeasurements = data.Clientsidemeasurements.ValueString()
	}
	if !data.Collectors.IsNull() && !data.Collectors.IsUnknown() {
		var collectorsList []string
		data.Collectors.ElementsAs(ctx, &collectorsList, false)
		appflowaction.Collectors = collectorsList
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appflowaction.Comment = data.Comment.ValueString()
	}
	if !data.Distributionalgorithm.IsNull() && !data.Distributionalgorithm.IsUnknown() {
		appflowaction.Distributionalgorithm = data.Distributionalgorithm.ValueString()
	}
	if !data.Metricslog.IsNull() && !data.Metricslog.IsUnknown() {
		appflowaction.Metricslog = data.Metricslog.ValueBool()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appflowaction.Name = data.Name.ValueString()
	}
	if !data.Pagetracking.IsNull() && !data.Pagetracking.IsUnknown() {
		appflowaction.Pagetracking = data.Pagetracking.ValueString()
	}
	if !data.Securityinsight.IsNull() && !data.Securityinsight.IsUnknown() {
		appflowaction.Securityinsight = data.Securityinsight.ValueString()
	}
	if !data.Transactionlog.IsNull() && !data.Transactionlog.IsUnknown() {
		appflowaction.Transactionlog = data.Transactionlog.ValueString()
	}
	if !data.Videoanalytics.IsNull() && !data.Videoanalytics.IsUnknown() {
		appflowaction.Videoanalytics = data.Videoanalytics.ValueString()
	}
	if !data.Webinsight.IsNull() && !data.Webinsight.IsUnknown() {
		appflowaction.Webinsight = data.Webinsight.ValueString()
	}

	return appflowaction
}

func appflowactionGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *AppflowactionResourceModel) appflow.Appflowaction {
	tflog.Debug(ctx, "In appflowactionGetTheUpdatablePayloadFromThePlan Function")

	// Create API request body from the model, restricted to NITRO-updatable fields.
	// metricslog and transactionlog are not updatable (ForceNew/RequiresReplace).
	appflowaction := appflow.Appflowaction{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appflowaction.Name = data.Name.ValueString()
	}
	if !data.Botinsight.IsNull() && !data.Botinsight.IsUnknown() {
		appflowaction.Botinsight = data.Botinsight.ValueString()
	}
	if !data.Ciinsight.IsNull() && !data.Ciinsight.IsUnknown() {
		appflowaction.Ciinsight = data.Ciinsight.ValueString()
	}
	if !data.Clientsidemeasurements.IsNull() && !data.Clientsidemeasurements.IsUnknown() {
		appflowaction.Clientsidemeasurements = data.Clientsidemeasurements.ValueString()
	}
	if !data.Collectors.IsNull() && !data.Collectors.IsUnknown() {
		var collectorsList []string
		data.Collectors.ElementsAs(ctx, &collectorsList, false)
		appflowaction.Collectors = collectorsList
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appflowaction.Comment = data.Comment.ValueString()
	}
	if !data.Distributionalgorithm.IsNull() && !data.Distributionalgorithm.IsUnknown() {
		appflowaction.Distributionalgorithm = data.Distributionalgorithm.ValueString()
	}
	if !data.Pagetracking.IsNull() && !data.Pagetracking.IsUnknown() {
		appflowaction.Pagetracking = data.Pagetracking.ValueString()
	}
	if !data.Securityinsight.IsNull() && !data.Securityinsight.IsUnknown() {
		appflowaction.Securityinsight = data.Securityinsight.ValueString()
	}
	if !data.Videoanalytics.IsNull() && !data.Videoanalytics.IsUnknown() {
		appflowaction.Videoanalytics = data.Videoanalytics.ValueString()
	}
	if !data.Webinsight.IsNull() && !data.Webinsight.IsUnknown() {
		appflowaction.Webinsight = data.Webinsight.ValueString()
	}

	return appflowaction
}

func appflowactionSetAttrFromGet(ctx context.Context, data *AppflowactionResourceModel, getResponseData map[string]interface{}) *AppflowactionResourceModel {
	tflog.Debug(ctx, "In appflowactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["botinsight"]; ok && val != nil {
		data.Botinsight = types.StringValue(val.(string))
	} else {
		data.Botinsight = types.StringNull()
	}
	if val, ok := getResponseData["ciinsight"]; ok && val != nil {
		data.Ciinsight = types.StringValue(val.(string))
	} else {
		data.Ciinsight = types.StringNull()
	}
	if val, ok := getResponseData["clientsidemeasurements"]; ok && val != nil {
		data.Clientsidemeasurements = types.StringValue(val.(string))
	} else {
		data.Clientsidemeasurements = types.StringNull()
	}
	if val, ok := getResponseData["collectors"]; ok && val != nil {
		switch v := val.(type) {
		case []interface{}:
			stringList := utils.ToStringList(v)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Collectors = listValue
		case string:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, []string{v})
			data.Collectors = listValue
		default:
			data.Collectors = types.ListNull(types.StringType)
		}
	} else {
		data.Collectors = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["distributionalgorithm"]; ok && val != nil {
		data.Distributionalgorithm = types.StringValue(val.(string))
	} else {
		data.Distributionalgorithm = types.StringNull()
	}
	if val, ok := getResponseData["metricslog"]; ok && val != nil {
		data.Metricslog = types.BoolValue(val.(bool))
	} else {
		data.Metricslog = types.BoolNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	// newname is not returned by NITRO API (rename-only parameter) - resolve to null
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["pagetracking"]; ok && val != nil {
		data.Pagetracking = types.StringValue(val.(string))
	} else {
		data.Pagetracking = types.StringNull()
	}
	if val, ok := getResponseData["securityinsight"]; ok && val != nil {
		data.Securityinsight = types.StringValue(val.(string))
	} else {
		data.Securityinsight = types.StringNull()
	}
	if val, ok := getResponseData["transactionlog"]; ok && val != nil {
		data.Transactionlog = types.StringValue(val.(string))
	} else {
		data.Transactionlog = types.StringNull()
	}
	if val, ok := getResponseData["videoanalytics"]; ok && val != nil {
		data.Videoanalytics = types.StringValue(val.(string))
	} else {
		data.Videoanalytics = types.StringNull()
	}
	if val, ok := getResponseData["webinsight"]; ok && val != nil {
		data.Webinsight = types.StringValue(val.(string))
	} else {
		data.Webinsight = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
