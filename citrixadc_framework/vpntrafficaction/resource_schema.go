package vpntrafficaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpntrafficactionResourceModel describes the resource data model.
type VpntrafficactionResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Apptimeout       types.Int64  `tfsdk:"apptimeout"`
	Formssoaction    types.String `tfsdk:"formssoaction"`
	Fta              types.String `tfsdk:"fta"`
	Hdx              types.String `tfsdk:"hdx"`
	Kcdaccount       types.String `tfsdk:"kcdaccount"`
	Name             types.String `tfsdk:"name"`
	Passwdexpression types.String `tfsdk:"passwdexpression"`
	Proxy            types.String `tfsdk:"proxy"`
	Qual             types.String `tfsdk:"qual"`
	Samlssoprofile   types.String `tfsdk:"samlssoprofile"`
	Sso              types.String `tfsdk:"sso"`
	Userexpression   types.String `tfsdk:"userexpression"`
	Wanscaler        types.String `tfsdk:"wanscaler"`
}

func (r *VpntrafficactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpntrafficaction resource.",
			},
			"apptimeout": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				// Optional+Computed without a documented default: retain the prior
				// state value when omitted from config so an unrelated update
				// (e.g. an unset of another attribute) does not spuriously mark it
				// as "known after apply" and force an empty update payload.
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Description: "Maximum amount of time, in minutes, a user can stay logged on to the web application.",
			},
			"formssoaction": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Name of the form-based single sign-on profile. Form-based single sign-on allows users to log on one time to all protected applications in your network, instead of requiring them to log on separately to access each one.",
			},
			"fta": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Specify file type association, which is a list of file extensions that users are allowed to open.",
			},
			"hdx": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Provide hdx proxy to the ICA traffic",
			},
			"kcdaccount": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Kerberos constrained delegation account name",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the traffic action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after a traffic action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"passwdexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(""),
				Description: "expression that will be evaluated to obtain password for SingleSignOn",
			},
			"proxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(""),
				Description: "IP address and Port of the proxy server to be used for HTTP access for this request.",
			},
			"qual": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					// qual is is_updateable:false in NITRO (UpdateResource rejects it
					// with errorcode 278 "Invalid argument [qual]"). A change to qual
					// must force recreation, mirroring the SDK v2 intent and NITRO reality.
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Protocol, either HTTP or TCP, to be used with the action.",
			},
			"samlssoprofile": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Profile to be used for doing SAML SSO to remote relying party",
			},
			"sso": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Provide single sign-on to the web application.\n	    NOTE : Authentication mechanisms like Basic-authentication  require the user credentials to be sent in plaintext which is not secure if the server is running on HTTP (instead of HTTPS).",
			},
			"userexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(""),
				Description: "expression that will be evaluated to obtain username for SingleSignOn",
			},
			"wanscaler": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Use the Repeater Plug-in to optimize network traffic.",
			},
		},
	}
}

func vpntrafficactionGetThePayloadFromtheConfig(ctx context.Context, data *VpntrafficactionResourceModel) vpn.Vpntrafficaction {
	tflog.Debug(ctx, "In vpntrafficactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpntrafficaction := vpn.Vpntrafficaction{}
	if !data.Apptimeout.IsNull() && !data.Apptimeout.IsUnknown() {
		vpntrafficaction.Apptimeout = utils.IntPtr(int(data.Apptimeout.ValueInt64()))
	}
	if !data.Formssoaction.IsNull() && !data.Formssoaction.IsUnknown() {
		vpntrafficaction.Formssoaction = data.Formssoaction.ValueString()
	}
	if !data.Fta.IsNull() && !data.Fta.IsUnknown() {
		vpntrafficaction.Fta = data.Fta.ValueString()
	}
	if !data.Hdx.IsNull() && !data.Hdx.IsUnknown() {
		vpntrafficaction.Hdx = data.Hdx.ValueString()
	}
	if !data.Kcdaccount.IsNull() && !data.Kcdaccount.IsUnknown() {
		vpntrafficaction.Kcdaccount = data.Kcdaccount.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpntrafficaction.Name = data.Name.ValueString()
	}
	if !data.Passwdexpression.IsNull() && !data.Passwdexpression.IsUnknown() {
		vpntrafficaction.Passwdexpression = data.Passwdexpression.ValueString()
	}
	if !data.Proxy.IsNull() && !data.Proxy.IsUnknown() {
		vpntrafficaction.Proxy = data.Proxy.ValueString()
	}
	if !data.Qual.IsNull() && !data.Qual.IsUnknown() {
		vpntrafficaction.Qual = data.Qual.ValueString()
	}
	if !data.Samlssoprofile.IsNull() && !data.Samlssoprofile.IsUnknown() {
		vpntrafficaction.Samlssoprofile = data.Samlssoprofile.ValueString()
	}
	if !data.Sso.IsNull() && !data.Sso.IsUnknown() {
		vpntrafficaction.Sso = data.Sso.ValueString()
	}
	if !data.Userexpression.IsNull() && !data.Userexpression.IsUnknown() {
		vpntrafficaction.Userexpression = data.Userexpression.ValueString()
	}
	if !data.Wanscaler.IsNull() && !data.Wanscaler.IsUnknown() {
		vpntrafficaction.Wanscaler = data.Wanscaler.ValueString()
	}

	return vpntrafficaction
}

// vpntrafficactionGetTheUpdatablePayloadFromThePlan builds the payload for the
// UpdateResource call, restricted to NITRO-updatable fields. qual is
// is_updateable:false and NITRO rejects it on update (errorcode 278 "Invalid
// argument [qual]"), so it is intentionally excluded here. It stays in the
// create payload (vpntrafficactionGetThePayloadFromtheConfig) and carries a
// RequiresReplace plan modifier, so any qual change is a recreate, never an update.
func vpntrafficactionGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *VpntrafficactionResourceModel) vpn.Vpntrafficaction {
	tflog.Debug(ctx, "In vpntrafficactionGetTheUpdatablePayloadFromThePlan Function")

	vpntrafficaction := vpn.Vpntrafficaction{}
	if !data.Apptimeout.IsNull() && !data.Apptimeout.IsUnknown() {
		vpntrafficaction.Apptimeout = utils.IntPtr(int(data.Apptimeout.ValueInt64()))
	}
	if !data.Formssoaction.IsNull() && !data.Formssoaction.IsUnknown() {
		vpntrafficaction.Formssoaction = data.Formssoaction.ValueString()
	}
	if !data.Fta.IsNull() && !data.Fta.IsUnknown() {
		vpntrafficaction.Fta = data.Fta.ValueString()
	}
	if !data.Hdx.IsNull() && !data.Hdx.IsUnknown() {
		vpntrafficaction.Hdx = data.Hdx.ValueString()
	}
	if !data.Kcdaccount.IsNull() && !data.Kcdaccount.IsUnknown() {
		vpntrafficaction.Kcdaccount = data.Kcdaccount.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpntrafficaction.Name = data.Name.ValueString()
	}
	if !data.Passwdexpression.IsNull() && !data.Passwdexpression.IsUnknown() {
		vpntrafficaction.Passwdexpression = data.Passwdexpression.ValueString()
	}
	if !data.Proxy.IsNull() && !data.Proxy.IsUnknown() {
		vpntrafficaction.Proxy = data.Proxy.ValueString()
	}
	// Skip non-updatable attribute: qual (NITRO errorcode 278 on update)
	if !data.Samlssoprofile.IsNull() && !data.Samlssoprofile.IsUnknown() {
		vpntrafficaction.Samlssoprofile = data.Samlssoprofile.ValueString()
	}
	if !data.Sso.IsNull() && !data.Sso.IsUnknown() {
		vpntrafficaction.Sso = data.Sso.ValueString()
	}
	if !data.Userexpression.IsNull() && !data.Userexpression.IsUnknown() {
		vpntrafficaction.Userexpression = data.Userexpression.ValueString()
	}
	if !data.Wanscaler.IsNull() && !data.Wanscaler.IsUnknown() {
		vpntrafficaction.Wanscaler = data.Wanscaler.ValueString()
	}

	return vpntrafficaction
}

// vpntrafficactionSetAttrFromGet maps the NITRO GET response onto the resource
// model. The else-branches only null a value when it was Unknown (never
// configured / never in prior state). This avoids the omit-on-default trap:
// NITRO omits some values from GET (e.g. default enum toggles), and blindly
// nulling them would clobber a known configured value and cause an
// "inconsistent result after apply" error.
func vpntrafficactionSetAttrFromGet(ctx context.Context, data *VpntrafficactionResourceModel, getResponseData map[string]interface{}) *VpntrafficactionResourceModel {
	tflog.Debug(ctx, "In vpntrafficactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["apptimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Apptimeout = types.Int64Value(intVal)
		}
	} else if data.Apptimeout.IsUnknown() {
		data.Apptimeout = types.Int64Null()
	}
	if val, ok := getResponseData["formssoaction"]; ok && val != nil {
		data.Formssoaction = types.StringValue(val.(string))
	} else if data.Formssoaction.IsUnknown() {
		data.Formssoaction = types.StringNull()
	}
	if val, ok := getResponseData["fta"]; ok && val != nil {
		data.Fta = types.StringValue(val.(string))
	} else if data.Fta.IsUnknown() {
		data.Fta = types.StringNull()
	}
	if val, ok := getResponseData["hdx"]; ok && val != nil {
		data.Hdx = types.StringValue(val.(string))
	} else if data.Hdx.IsUnknown() {
		data.Hdx = types.StringNull()
	}
	if val, ok := getResponseData["kcdaccount"]; ok && val != nil {
		data.Kcdaccount = types.StringValue(val.(string))
	} else if data.Kcdaccount.IsUnknown() {
		data.Kcdaccount = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["passwdexpression"]; ok && val != nil {
		data.Passwdexpression = types.StringValue(val.(string))
	} else if data.Passwdexpression.IsUnknown() {
		data.Passwdexpression = types.StringNull()
	}
	if val, ok := getResponseData["proxy"]; ok && val != nil {
		data.Proxy = types.StringValue(val.(string))
	} else if data.Proxy.IsUnknown() {
		data.Proxy = types.StringNull()
	}
	if val, ok := getResponseData["qual"]; ok && val != nil {
		data.Qual = types.StringValue(val.(string))
	} else if data.Qual.IsUnknown() {
		data.Qual = types.StringNull()
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
	if val, ok := getResponseData["wanscaler"]; ok && val != nil {
		data.Wanscaler = types.StringValue(val.(string))
	} else if data.Wanscaler.IsUnknown() {
		data.Wanscaler = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}

// vpntrafficactionSetAttrFromGetForDatasource faithfully copies every field from
// the GET response (nulling absent fields) for the datasource path, which has no
// prior plan/state to preserve. It also sets the ID since the datasource never
// runs Create.
func vpntrafficactionSetAttrFromGetForDatasource(ctx context.Context, data *VpntrafficactionResourceModel, getResponseData map[string]interface{}) *VpntrafficactionResourceModel {
	tflog.Debug(ctx, "In vpntrafficactionSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["apptimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Apptimeout = types.Int64Value(intVal)
		} else {
			data.Apptimeout = types.Int64Null()
		}
	} else {
		data.Apptimeout = types.Int64Null()
	}
	if val, ok := getResponseData["formssoaction"]; ok && val != nil {
		data.Formssoaction = types.StringValue(val.(string))
	} else {
		data.Formssoaction = types.StringNull()
	}
	if val, ok := getResponseData["fta"]; ok && val != nil {
		data.Fta = types.StringValue(val.(string))
	} else {
		data.Fta = types.StringNull()
	}
	if val, ok := getResponseData["hdx"]; ok && val != nil {
		data.Hdx = types.StringValue(val.(string))
	} else {
		data.Hdx = types.StringNull()
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
	if val, ok := getResponseData["proxy"]; ok && val != nil {
		data.Proxy = types.StringValue(val.(string))
	} else {
		data.Proxy = types.StringNull()
	}
	if val, ok := getResponseData["qual"]; ok && val != nil {
		data.Qual = types.StringValue(val.(string))
	} else {
		data.Qual = types.StringNull()
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
	if val, ok := getResponseData["wanscaler"]; ok && val != nil {
		data.Wanscaler = types.StringValue(val.(string))
	} else {
		data.Wanscaler = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
