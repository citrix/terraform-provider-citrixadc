package vpnformssoaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpnformssoactionResourceModel describes the resource data model.
type VpnformssoactionResourceModel struct {
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

func (r *VpnformssoactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnformssoaction resource.",
			},
			"actionurl": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Root-relative URL to which the completed form is submitted.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the form based single sign-on profile.",
			},
			"namevaluepair": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(""),
				Description: "Other name-value pair attributes to send to the server, in addition to sending the user name and password. Value names are separated by an ampersand (&), such as in name1=value1&name2=value2.",
			},
			"nvtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DYNAMIC"),
				Description: "How to process the name-value pair. Available settings function as follows:\n* STATIC - The administrator-configured values are used.\n* DYNAMIC - The response is parsed, the form is extracted, and then submitted.",
			},
			"passwdfield": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the form field in which the user types in the password.",
			},
			"responsesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(8096),
				Description: "Maximum number of bytes to allow in the response size. Specifies the number of bytes in the response to be parsed for extracting the forms.",
			},
			"ssosuccessrule": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Expression that defines the criteria for SSO success. Expression such as checking for cookie in the response is a common example.",
			},
			"submitmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("GET"),
				Description: "HTTP method (GET or POST) used by the single sign-on form to send the logon credentials to the logon server.",
			},
			"userfield": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the form field in which the user types in the user ID.",
			},
		},
	}
}

func vpnformssoactionGetThePayloadFromthePlan(ctx context.Context, data *VpnformssoactionResourceModel) vpn.Vpnformssoaction {
	tflog.Debug(ctx, "In vpnformssoactionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnformssoaction := vpn.Vpnformssoaction{}
	if !data.Actionurl.IsNull() && !data.Actionurl.IsUnknown() {
		vpnformssoaction.Actionurl = data.Actionurl.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnformssoaction.Name = data.Name.ValueString()
	}
	if !data.Namevaluepair.IsNull() && !data.Namevaluepair.IsUnknown() {
		vpnformssoaction.Namevaluepair = data.Namevaluepair.ValueString()
	}
	if !data.Nvtype.IsNull() && !data.Nvtype.IsUnknown() {
		vpnformssoaction.Nvtype = data.Nvtype.ValueString()
	}
	if !data.Passwdfield.IsNull() && !data.Passwdfield.IsUnknown() {
		vpnformssoaction.Passwdfield = data.Passwdfield.ValueString()
	}
	if !data.Responsesize.IsNull() && !data.Responsesize.IsUnknown() {
		vpnformssoaction.Responsesize = utils.IntPtr(int(data.Responsesize.ValueInt64()))
	}
	if !data.Ssosuccessrule.IsNull() && !data.Ssosuccessrule.IsUnknown() {
		vpnformssoaction.Ssosuccessrule = data.Ssosuccessrule.ValueString()
	}
	if !data.Submitmethod.IsNull() && !data.Submitmethod.IsUnknown() {
		vpnformssoaction.Submitmethod = data.Submitmethod.ValueString()
	}
	if !data.Userfield.IsNull() && !data.Userfield.IsUnknown() {
		vpnformssoaction.Userfield = data.Userfield.ValueString()
	}

	return vpnformssoaction
}

func vpnformssoactionSetAttrFromGet(ctx context.Context, data *VpnformssoactionResourceModel, getResponseData map[string]interface{}) *VpnformssoactionResourceModel {
	tflog.Debug(ctx, "In vpnformssoactionSetAttrFromGet Function")

	// Convert API response to model.
	// Guard else-branches: only null a field when the current value is Unknown,
	// so a configured value that NITRO omits from GET (omit-on-default) is preserved.
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
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
