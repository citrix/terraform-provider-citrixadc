package tmtrafficaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/tm"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// TmtrafficactionResourceModel describes the resource data model.
type TmtrafficactionResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Apptimeout       types.Int64  `tfsdk:"apptimeout"`
	Forcedtimeout    types.String `tfsdk:"forcedtimeout"`
	Forcedtimeoutval types.Int64  `tfsdk:"forcedtimeoutval"`
	Formssoaction    types.String `tfsdk:"formssoaction"`
	Initiatelogout   types.String `tfsdk:"initiatelogout"`
	Kcdaccount       types.String `tfsdk:"kcdaccount"`
	Name             types.String `tfsdk:"name"`
	Passwdexpression types.String `tfsdk:"passwdexpression"`
	Persistentcookie types.String `tfsdk:"persistentcookie"`
	Samlssoprofile   types.String `tfsdk:"samlssoprofile"`
	Sso              types.String `tfsdk:"sso"`
	Userexpression   types.String `tfsdk:"userexpression"`
}

func (r *TmtrafficactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the tmtrafficaction resource.",
			},
			"apptimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time interval, in minutes, of user inactivity after which the connection is closed.",
			},
			"forcedtimeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Setting to start, stop or reset TM session force timer",
			},
			"forcedtimeoutval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time interval, in minutes, for which force timer should be set.",
			},
			"formssoaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the configured form-based single sign-on profile.",
			},
			"initiatelogout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initiate logout for the traffic management (TM) session if the policy evaluates to true. The session is then terminated after two minutes.",
			},
			"kcdaccount": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("None"),
				Description: "Kerberos constrained delegation account name",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the traffic action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after a traffic action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"passwdexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "expression that will be evaluated to obtain password for SingleSignOn",
			},
			"persistentcookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use persistent cookies for the traffic session. A persistent cookie remains on the user device and is sent with each HTTP request. The cookie becomes stale if the session ends.",
			},
			"samlssoprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Profile to be used for doing SAML SSO to remote relying party",
			},
			"sso": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use single sign-on for the resource that the user is accessing now.",
			},
			"userexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "expression that will be evaluated to obtain username for SingleSignOn",
			},
		},
	}
}

func tmtrafficactionGetThePayloadFromtheConfig(ctx context.Context, data *TmtrafficactionResourceModel) tm.Tmtrafficaction {
	tflog.Debug(ctx, "In tmtrafficactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	tmtrafficaction := tm.Tmtrafficaction{}
	if !data.Apptimeout.IsNull() {
		tmtrafficaction.Apptimeout = utils.IntPtr(int(data.Apptimeout.ValueInt64()))
	}
	if !data.Forcedtimeout.IsNull() {
		tmtrafficaction.Forcedtimeout = data.Forcedtimeout.ValueString()
	}
	if !data.Forcedtimeoutval.IsNull() {
		tmtrafficaction.Forcedtimeoutval = utils.IntPtr(int(data.Forcedtimeoutval.ValueInt64()))
	}
	if !data.Formssoaction.IsNull() {
		tmtrafficaction.Formssoaction = data.Formssoaction.ValueString()
	}
	if !data.Initiatelogout.IsNull() {
		tmtrafficaction.Initiatelogout = data.Initiatelogout.ValueString()
	}
	if !data.Kcdaccount.IsNull() {
		tmtrafficaction.Kcdaccount = data.Kcdaccount.ValueString()
	}
	if !data.Name.IsNull() {
		tmtrafficaction.Name = data.Name.ValueString()
	}
	if !data.Passwdexpression.IsNull() {
		tmtrafficaction.Passwdexpression = data.Passwdexpression.ValueString()
	}
	if !data.Persistentcookie.IsNull() {
		tmtrafficaction.Persistentcookie = data.Persistentcookie.ValueString()
	}
	if !data.Samlssoprofile.IsNull() {
		tmtrafficaction.Samlssoprofile = data.Samlssoprofile.ValueString()
	}
	if !data.Sso.IsNull() {
		tmtrafficaction.Sso = data.Sso.ValueString()
	}
	if !data.Userexpression.IsNull() {
		tmtrafficaction.Userexpression = data.Userexpression.ValueString()
	}

	return tmtrafficaction
}

func tmtrafficactionSetAttrFromGet(ctx context.Context, data *TmtrafficactionResourceModel, getResponseData map[string]interface{}) *TmtrafficactionResourceModel {
	tflog.Debug(ctx, "In tmtrafficactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["apptimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Apptimeout = types.Int64Value(intVal)
		}
	} else {
		data.Apptimeout = types.Int64Null()
	}
	if val, ok := getResponseData["forcedtimeout"]; ok && val != nil {
		data.Forcedtimeout = types.StringValue(val.(string))
	} else {
		data.Forcedtimeout = types.StringNull()
	}
	if val, ok := getResponseData["forcedtimeoutval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Forcedtimeoutval = types.Int64Value(intVal)
		}
	} else {
		data.Forcedtimeoutval = types.Int64Null()
	}
	if val, ok := getResponseData["formssoaction"]; ok && val != nil {
		data.Formssoaction = types.StringValue(val.(string))
	} else {
		data.Formssoaction = types.StringNull()
	}
	if val, ok := getResponseData["initiatelogout"]; ok && val != nil {
		data.Initiatelogout = types.StringValue(val.(string))
	} else {
		data.Initiatelogout = types.StringNull()
	}
	if val, ok := getResponseData["kcdaccount"]; ok && val != nil {
		data.Kcdaccount = types.StringValue(val.(string))
	} else {
		data.Kcdaccount = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["passwdexpression"]; ok && val != nil {
		data.Passwdexpression = types.StringValue(val.(string))
	} else {
		data.Passwdexpression = types.StringNull()
	}
	if val, ok := getResponseData["persistentcookie"]; ok && val != nil {
		data.Persistentcookie = types.StringValue(val.(string))
	} else {
		data.Persistentcookie = types.StringNull()
	}
	if val, ok := getResponseData["samlssoprofile"]; ok && val != nil {
		data.Samlssoprofile = types.StringValue(val.(string))
	} else {
		data.Samlssoprofile = types.StringNull()
	}
	if val, ok := getResponseData["sso"]; ok && val != nil {
		data.Sso = types.StringValue(val.(string))
	} else {
		data.Sso = types.StringNull()
	}
	if val, ok := getResponseData["userexpression"]; ok && val != nil {
		data.Userexpression = types.StringValue(val.(string))
	} else {
		data.Userexpression = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
