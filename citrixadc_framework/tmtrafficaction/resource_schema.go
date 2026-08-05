package tmtrafficaction

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
				Computed:    true,
				Description: "Kerberos constrained delegation account name",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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

func tmtrafficactionGetThePayloadFromthePlan(ctx context.Context, data *TmtrafficactionResourceModel) tm.Tmtrafficaction {
	tflog.Debug(ctx, "In tmtrafficactionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	tmtrafficaction := tm.Tmtrafficaction{}
	if !data.Apptimeout.IsNull() && !data.Apptimeout.IsUnknown() {
		tmtrafficaction.Apptimeout = utils.IntPtr(int(data.Apptimeout.ValueInt64()))
	}
	if !data.Forcedtimeout.IsNull() && !data.Forcedtimeout.IsUnknown() {
		tmtrafficaction.Forcedtimeout = data.Forcedtimeout.ValueString()
	}
	if !data.Forcedtimeoutval.IsNull() && !data.Forcedtimeoutval.IsUnknown() {
		tmtrafficaction.Forcedtimeoutval = utils.IntPtr(int(data.Forcedtimeoutval.ValueInt64()))
	}
	if !data.Formssoaction.IsNull() && !data.Formssoaction.IsUnknown() {
		tmtrafficaction.Formssoaction = data.Formssoaction.ValueString()
	}
	if !data.Initiatelogout.IsNull() && !data.Initiatelogout.IsUnknown() {
		tmtrafficaction.Initiatelogout = data.Initiatelogout.ValueString()
	}
	if !data.Kcdaccount.IsNull() && !data.Kcdaccount.IsUnknown() {
		tmtrafficaction.Kcdaccount = data.Kcdaccount.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		tmtrafficaction.Name = data.Name.ValueString()
	}
	if !data.Passwdexpression.IsNull() && !data.Passwdexpression.IsUnknown() {
		tmtrafficaction.Passwdexpression = data.Passwdexpression.ValueString()
	}
	if !data.Persistentcookie.IsNull() && !data.Persistentcookie.IsUnknown() {
		tmtrafficaction.Persistentcookie = data.Persistentcookie.ValueString()
	}
	if !data.Samlssoprofile.IsNull() && !data.Samlssoprofile.IsUnknown() {
		tmtrafficaction.Samlssoprofile = data.Samlssoprofile.ValueString()
	}
	if !data.Sso.IsNull() && !data.Sso.IsUnknown() {
		tmtrafficaction.Sso = data.Sso.ValueString()
	}
	if !data.Userexpression.IsNull() && !data.Userexpression.IsUnknown() {
		tmtrafficaction.Userexpression = data.Userexpression.ValueString()
	}

	return tmtrafficaction
}

// tmtrafficactionSetAttrFromGet maps the NITRO GET response onto the resource
// state model. To avoid the omit-on-default trap (NITRO omits a configured
// value from GET), the else-branches only null a value that is still Unknown
// (needed to resolve Computed attrs on Create); a known configured value is
// preserved.
func tmtrafficactionSetAttrFromGet(ctx context.Context, data *TmtrafficactionResourceModel, getResponseData map[string]interface{}) *TmtrafficactionResourceModel {
	tflog.Debug(ctx, "In tmtrafficactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["apptimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Apptimeout = types.Int64Value(intVal)
		} else if data.Apptimeout.IsUnknown() {
			data.Apptimeout = types.Int64Null()
		}
	} else if data.Apptimeout.IsUnknown() {
		data.Apptimeout = types.Int64Null()
	}
	if val, ok := getResponseData["forcedtimeout"]; ok && val != nil {
		data.Forcedtimeout = types.StringValue(val.(string))
	} else if data.Forcedtimeout.IsUnknown() {
		data.Forcedtimeout = types.StringNull()
	}
	if val, ok := getResponseData["forcedtimeoutval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Forcedtimeoutval = types.Int64Value(intVal)
		} else if data.Forcedtimeoutval.IsUnknown() {
			data.Forcedtimeoutval = types.Int64Null()
		}
	} else if data.Forcedtimeoutval.IsUnknown() {
		data.Forcedtimeoutval = types.Int64Null()
	}
	if val, ok := getResponseData["formssoaction"]; ok && val != nil {
		data.Formssoaction = types.StringValue(val.(string))
	} else if data.Formssoaction.IsUnknown() {
		data.Formssoaction = types.StringNull()
	}
	if val, ok := getResponseData["initiatelogout"]; ok && val != nil {
		data.Initiatelogout = types.StringValue(val.(string))
	} else if data.Initiatelogout.IsUnknown() {
		data.Initiatelogout = types.StringNull()
	}
	if val, ok := getResponseData["kcdaccount"]; ok && val != nil {
		data.Kcdaccount = types.StringValue(val.(string))
	} else if data.Kcdaccount.IsUnknown() {
		data.Kcdaccount = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["passwdexpression"]; ok && val != nil {
		data.Passwdexpression = types.StringValue(val.(string))
	} else if data.Passwdexpression.IsUnknown() {
		data.Passwdexpression = types.StringNull()
	}
	if val, ok := getResponseData["persistentcookie"]; ok && val != nil {
		data.Persistentcookie = types.StringValue(val.(string))
	} else if data.Persistentcookie.IsUnknown() {
		data.Persistentcookie = types.StringNull()
	}
	if val, ok := getResponseData["samlssoprofile"]; ok && val != nil {
		data.Samlssoprofile = types.StringValue(val.(string))
	} else if data.Samlssoprofile.IsUnknown() {
		data.Samlssoprofile = types.StringNull()
	}
	if val, ok := getResponseData["sso"]; ok && val != nil {
		data.Sso = types.StringValue(val.(string))
	} else if data.Sso.IsUnknown() {
		data.Sso = types.StringNull()
	}
	if val, ok := getResponseData["userexpression"]; ok && val != nil {
		data.Userexpression = types.StringValue(val.(string))
	} else if data.Userexpression.IsUnknown() {
		data.Userexpression = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}

// tmtrafficactionSetAttrFromGetForDatasource maps the NITRO GET response onto
// the model for the datasource, unconditionally copying every returned value
// (nulling anything absent) and setting the ID.
func tmtrafficactionSetAttrFromGetForDatasource(ctx context.Context, data *TmtrafficactionResourceModel, getResponseData map[string]interface{}) *TmtrafficactionResourceModel {
	tflog.Debug(ctx, "In tmtrafficactionSetAttrFromGetForDatasource Function")

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

	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
