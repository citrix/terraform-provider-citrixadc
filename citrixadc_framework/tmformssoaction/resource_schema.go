package tmformssoaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/tm"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// TmformssoactionResourceModel describes the resource data model.
type TmformssoactionResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Actionurl      types.String `tfsdk:"actionurl"`
	Name           types.String `tfsdk:"name"`
	Namevaluepair  types.String `tfsdk:"namevaluepair"`
	Nvtype         types.String `tfsdk:"nvtype"`
	Passwdfield    types.String `tfsdk:"passwdfield"`
	Responsesize   types.Int64  `tfsdk:"responsesize"`
	Ssosuccessrule types.String `tfsdk:"ssosuccessrule"`
	Submitmethod   types.String `tfsdk:"submitmethod"`
	Userfield      types.String `tfsdk:"userfield"`
}

func (r *TmformssoactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the tmformssoaction resource.",
			},
			"actionurl": schema.StringAttribute{
				Required:    true,
				Description: "URL to which the completed form is submitted.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the new form-based single sign-on profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an SSO action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"namevaluepair": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name-value pair attributes to send to the server in addition to sending the username and password. Value names are separated by an ampersand (&) (for example, name1=value1&name2=value2).",
			},
			"nvtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of processing of the name-value pair. If you specify STATIC, the values configured by the administrator are used. For DYNAMIC, the response is parsed, and the form is extracted and then submitted.",
			},
			"passwdfield": schema.StringAttribute{
				Required:    true,
				Description: "Name of the form field in which the user types in the password.",
			},
			"responsesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of bytes, in the response, to parse for extracting the forms.",
			},
			"ssosuccessrule": schema.StringAttribute{
				Required:    true,
				Description: "Expression, that checks to see if single sign-on is successful.",
			},
			"submitmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP method used by the single sign-on form to send the logon credentials to the logon server. Applies only to STATIC name-value type.",
			},
			"userfield": schema.StringAttribute{
				Required:    true,
				Description: "Name of the form field in which the user types in the user ID.",
			},
		},
	}
}

func tmformssoactionGetThePayloadFromthePlan(ctx context.Context, data *TmformssoactionResourceModel) tm.Tmformssoaction {
	tflog.Debug(ctx, "In tmformssoactionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	tmformssoaction := tm.Tmformssoaction{}
	if !data.Actionurl.IsNull() && !data.Actionurl.IsUnknown() {
		tmformssoaction.Actionurl = data.Actionurl.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		tmformssoaction.Name = data.Name.ValueString()
	}
	if !data.Namevaluepair.IsNull() && !data.Namevaluepair.IsUnknown() {
		tmformssoaction.Namevaluepair = data.Namevaluepair.ValueString()
	}
	if !data.Nvtype.IsNull() && !data.Nvtype.IsUnknown() {
		tmformssoaction.Nvtype = data.Nvtype.ValueString()
	}
	if !data.Passwdfield.IsNull() && !data.Passwdfield.IsUnknown() {
		tmformssoaction.Passwdfield = data.Passwdfield.ValueString()
	}
	if !data.Responsesize.IsNull() && !data.Responsesize.IsUnknown() {
		tmformssoaction.Responsesize = utils.IntPtr(int(data.Responsesize.ValueInt64()))
	}
	if !data.Ssosuccessrule.IsNull() && !data.Ssosuccessrule.IsUnknown() {
		tmformssoaction.Ssosuccessrule = data.Ssosuccessrule.ValueString()
	}
	if !data.Submitmethod.IsNull() && !data.Submitmethod.IsUnknown() {
		tmformssoaction.Submitmethod = data.Submitmethod.ValueString()
	}
	if !data.Userfield.IsNull() && !data.Userfield.IsUnknown() {
		tmformssoaction.Userfield = data.Userfield.ValueString()
	}

	return tmformssoaction
}

func tmformssoactionSetAttrFromGet(ctx context.Context, data *TmformssoactionResourceModel, getResponseData map[string]interface{}) *TmformssoactionResourceModel {
	tflog.Debug(ctx, "In tmformssoactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["actionurl"]; ok && val != nil {
		data.Actionurl = types.StringValue(val.(string))
	} else if data.Actionurl.IsUnknown() {
		data.Actionurl = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}
	// Guard against omit-on-default: only null when the value is Unknown,
	// never clobber a known configured value that NITRO omitted from GET.
	if val, ok := getResponseData["namevaluepair"]; ok && val != nil {
		data.Namevaluepair = types.StringValue(val.(string))
	} else if data.Namevaluepair.IsUnknown() {
		data.Namevaluepair = types.StringNull()
	}
	if val, ok := getResponseData["nvtype"]; ok && val != nil {
		data.Nvtype = types.StringValue(val.(string))
	} else if data.Nvtype.IsUnknown() {
		data.Nvtype = types.StringNull()
	}
	if val, ok := getResponseData["passwdfield"]; ok && val != nil {
		data.Passwdfield = types.StringValue(val.(string))
	} else if data.Passwdfield.IsUnknown() {
		data.Passwdfield = types.StringNull()
	}
	if val, ok := getResponseData["responsesize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Responsesize = types.Int64Value(intVal)
		}
	} else if data.Responsesize.IsUnknown() {
		data.Responsesize = types.Int64Null()
	}
	if val, ok := getResponseData["ssosuccessrule"]; ok && val != nil {
		data.Ssosuccessrule = types.StringValue(val.(string))
	} else if data.Ssosuccessrule.IsUnknown() {
		data.Ssosuccessrule = types.StringNull()
	}
	if val, ok := getResponseData["submitmethod"]; ok && val != nil {
		data.Submitmethod = types.StringValue(val.(string))
	} else if data.Submitmethod.IsUnknown() {
		data.Submitmethod = types.StringNull()
	}
	if val, ok := getResponseData["userfield"]; ok && val != nil {
		data.Userfield = types.StringValue(val.(string))
	} else if data.Userfield.IsUnknown() {
		data.Userfield = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
