package transformaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/transform"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// TransformactionResourceModel describes the resource data model.
type TransformactionResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Comment          types.String `tfsdk:"comment"`
	Cookiedomainfrom types.String `tfsdk:"cookiedomainfrom"`
	Cookiedomaininto types.String `tfsdk:"cookiedomaininto"`
	Name             types.String `tfsdk:"name"`
	Priority         types.Int64  `tfsdk:"priority"`
	Profilename      types.String `tfsdk:"profilename"`
	Requrlfrom       types.String `tfsdk:"requrlfrom"`
	Requrlinto       types.String `tfsdk:"requrlinto"`
	Resurlfrom       types.String `tfsdk:"resurlfrom"`
	Resurlinto       types.String `tfsdk:"resurlinto"`
	State            types.String `tfsdk:"state"`
}

func (r *TransformactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the transformaction resource.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this URL Transformation action.",
			},
			"cookiedomainfrom": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Pattern that matches the domain to be transformed in Set-Cookie headers.",
			},
			"cookiedomaininto": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "PCRE-format regular expression that describes the transformation to be performed on cookie domains that match the cookieDomainFrom pattern. \nNOTE: The cookie domain to be transformed is extracted from the request.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the URL transformation action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the URL Transformation action is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, my transform action or my transform action).",
			},
			"priority": schema.Int64Attribute{
				Required:    true,
				Description: "Positive integer specifying the priority of the action within the profile. A lower number specifies a higher priority. Must be unique within the list of actions bound to the profile. Policies are evaluated in the order of their priority numbers, and the first policy that matches is applied.",
			},
			"profilename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the URL Transformation profile with which to associate this action.",
			},
			"requrlfrom": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "PCRE-format regular expression that describes the request URL pattern to be transformed.",
			},
			"requrlinto": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "PCRE-format regular expression that describes the transformation to be performed on URLs that match the reqUrlFrom pattern.",
			},
			"resurlfrom": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "PCRE-format regular expression that describes the response URL pattern to be transformed.",
			},
			"resurlinto": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "PCRE-format regular expression that describes the transformation to be performed on URLs that match the resUrlFrom pattern.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable this action.",
			},
		},
	}
}

func transformactionGetThePayloadFromtheConfig(ctx context.Context, data *TransformactionResourceModel) transform.Transformaction {
	tflog.Debug(ctx, "In transformactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	transformaction := transform.Transformaction{}
	if !data.Comment.IsNull() {
		transformaction.Comment = data.Comment.ValueString()
	}
	if !data.Cookiedomainfrom.IsNull() {
		transformaction.Cookiedomainfrom = data.Cookiedomainfrom.ValueString()
	}
	if !data.Cookiedomaininto.IsNull() {
		transformaction.Cookiedomaininto = data.Cookiedomaininto.ValueString()
	}
	if !data.Name.IsNull() {
		transformaction.Name = data.Name.ValueString()
	}
	if !data.Priority.IsNull() {
		transformaction.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Profilename.IsNull() {
		transformaction.Profilename = data.Profilename.ValueString()
	}
	if !data.Requrlfrom.IsNull() {
		transformaction.Requrlfrom = data.Requrlfrom.ValueString()
	}
	if !data.Requrlinto.IsNull() {
		transformaction.Requrlinto = data.Requrlinto.ValueString()
	}
	if !data.Resurlfrom.IsNull() {
		transformaction.Resurlfrom = data.Resurlfrom.ValueString()
	}
	if !data.Resurlinto.IsNull() {
		transformaction.Resurlinto = data.Resurlinto.ValueString()
	}
	if !data.State.IsNull() {
		transformaction.State = data.State.ValueString()
	}

	return transformaction
}

func transformactionSetAttrFromGet(ctx context.Context, data *TransformactionResourceModel, getResponseData map[string]interface{}) *TransformactionResourceModel {
	tflog.Debug(ctx, "In transformactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["cookiedomainfrom"]; ok && val != nil {
		data.Cookiedomainfrom = types.StringValue(val.(string))
	} else {
		data.Cookiedomainfrom = types.StringNull()
	}
	if val, ok := getResponseData["cookiedomaininto"]; ok && val != nil {
		data.Cookiedomaininto = types.StringValue(val.(string))
	} else {
		data.Cookiedomaininto = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["profilename"]; ok && val != nil {
		data.Profilename = types.StringValue(val.(string))
	} else {
		data.Profilename = types.StringNull()
	}
	if val, ok := getResponseData["requrlfrom"]; ok && val != nil {
		data.Requrlfrom = types.StringValue(val.(string))
	} else {
		data.Requrlfrom = types.StringNull()
	}
	if val, ok := getResponseData["requrlinto"]; ok && val != nil {
		data.Requrlinto = types.StringValue(val.(string))
	} else {
		data.Requrlinto = types.StringNull()
	}
	if val, ok := getResponseData["resurlfrom"]; ok && val != nil {
		data.Resurlfrom = types.StringValue(val.(string))
	} else {
		data.Resurlfrom = types.StringNull()
	}
	if val, ok := getResponseData["resurlinto"]; ok && val != nil {
		data.Resurlinto = types.StringValue(val.(string))
	} else {
		data.Resurlinto = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
