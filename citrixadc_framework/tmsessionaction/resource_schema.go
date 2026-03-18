package tmsessionaction

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

// TmsessionactionResourceModel describes the resource data model.
type TmsessionactionResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Defaultauthorizationaction types.String `tfsdk:"defaultauthorizationaction"`
	Homepage                   types.String `tfsdk:"homepage"`
	Httponlycookie             types.String `tfsdk:"httponlycookie"`
	Kcdaccount                 types.String `tfsdk:"kcdaccount"`
	Name                       types.String `tfsdk:"name"`
	Persistentcookie           types.String `tfsdk:"persistentcookie"`
	Persistentcookievalidity   types.Int64  `tfsdk:"persistentcookievalidity"`
	Sesstimeout                types.Int64  `tfsdk:"sesstimeout"`
	Sso                        types.String `tfsdk:"sso"`
	Ssocredential              types.String `tfsdk:"ssocredential"`
	Ssodomain                  types.String `tfsdk:"ssodomain"`
}

func (r *TmsessionactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the tmsessionaction resource.",
			},
			"defaultauthorizationaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow or deny access to content for which there is no specific authorization policy.",
			},
			"homepage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Web address of the home page that a user is displayed when authentication vserver is bookmarked and used to login.",
			},
			"httponlycookie": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Allow only an HTTP session cookie, in which case the cookie cannot be accessed by scripts.",
			},
			"kcdaccount": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Kerberos constrained delegation account name",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the session action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after a session action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"persistentcookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable persistent SSO cookies for the traffic management (TM) session. A persistent cookie remains on the user device and is sent with each HTTP request. The cookie becomes stale if the session ends. This setting is overwritten if a traffic action sets persistent cookie to OFF.\nNote: If persistent cookie is enabled, make sure you set the persistent cookie validity.",
			},
			"persistentcookievalidity": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the number of minutes for which the persistent cookie remains valid. Can be set only if the persistent cookie setting is enabled.",
			},
			"sesstimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Session timeout, in minutes. If there is no traffic during the timeout period, the user is disconnected and must reauthenticate to access intranet resources.",
			},
			"sso": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use single sign-on (SSO) to log users on to all web applications automatically after they authenticate, or pass users to the web application logon page to authenticate to each application individually. Note that this configuration does not honor the following authentication types for security reason. BASIC, DIGEST, and NTLM (without Negotiate NTLM2 Key or Negotiate Sign Flag). Use TM TrafficAction to configure SSO for these authentication types.",
			},
			"ssocredential": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the primary or secondary authentication credentials for single sign-on (SSO).",
			},
			"ssodomain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain to use for single sign-on (SSO).",
			},
		},
	}
}

func tmsessionactionGetThePayloadFromtheConfig(ctx context.Context, data *TmsessionactionResourceModel) tm.Tmsessionaction {
	tflog.Debug(ctx, "In tmsessionactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	tmsessionaction := tm.Tmsessionaction{}
	if !data.Defaultauthorizationaction.IsNull() {
		tmsessionaction.Defaultauthorizationaction = data.Defaultauthorizationaction.ValueString()
	}
	if !data.Homepage.IsNull() {
		tmsessionaction.Homepage = data.Homepage.ValueString()
	}
	if !data.Httponlycookie.IsNull() {
		tmsessionaction.Httponlycookie = data.Httponlycookie.ValueString()
	}
	if !data.Kcdaccount.IsNull() {
		tmsessionaction.Kcdaccount = data.Kcdaccount.ValueString()
	}
	if !data.Name.IsNull() {
		tmsessionaction.Name = data.Name.ValueString()
	}
	if !data.Persistentcookie.IsNull() {
		tmsessionaction.Persistentcookie = data.Persistentcookie.ValueString()
	}
	if !data.Persistentcookievalidity.IsNull() {
		tmsessionaction.Persistentcookievalidity = utils.IntPtr(int(data.Persistentcookievalidity.ValueInt64()))
	}
	if !data.Sesstimeout.IsNull() {
		tmsessionaction.Sesstimeout = utils.IntPtr(int(data.Sesstimeout.ValueInt64()))
	}
	if !data.Sso.IsNull() {
		tmsessionaction.Sso = data.Sso.ValueString()
	}
	if !data.Ssocredential.IsNull() {
		tmsessionaction.Ssocredential = data.Ssocredential.ValueString()
	}
	if !data.Ssodomain.IsNull() {
		tmsessionaction.Ssodomain = data.Ssodomain.ValueString()
	}

	return tmsessionaction
}

func tmsessionactionSetAttrFromGet(ctx context.Context, data *TmsessionactionResourceModel, getResponseData map[string]interface{}) *TmsessionactionResourceModel {
	tflog.Debug(ctx, "In tmsessionactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["defaultauthorizationaction"]; ok && val != nil {
		data.Defaultauthorizationaction = types.StringValue(val.(string))
	} else {
		data.Defaultauthorizationaction = types.StringNull()
	}
	if val, ok := getResponseData["homepage"]; ok && val != nil {
		data.Homepage = types.StringValue(val.(string))
	} else {
		data.Homepage = types.StringNull()
	}
	if val, ok := getResponseData["httponlycookie"]; ok && val != nil {
		data.Httponlycookie = types.StringValue(val.(string))
	} else {
		data.Httponlycookie = types.StringNull()
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
	if val, ok := getResponseData["persistentcookie"]; ok && val != nil {
		data.Persistentcookie = types.StringValue(val.(string))
	} else {
		data.Persistentcookie = types.StringNull()
	}
	if val, ok := getResponseData["persistentcookievalidity"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Persistentcookievalidity = types.Int64Value(intVal)
		}
	} else {
		data.Persistentcookievalidity = types.Int64Null()
	}
	if val, ok := getResponseData["sesstimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sesstimeout = types.Int64Value(intVal)
		}
	} else {
		data.Sesstimeout = types.Int64Null()
	}
	if val, ok := getResponseData["sso"]; ok && val != nil {
		data.Sso = types.StringValue(val.(string))
	} else {
		data.Sso = types.StringNull()
	}
	if val, ok := getResponseData["ssocredential"]; ok && val != nil {
		data.Ssocredential = types.StringValue(val.(string))
	} else {
		data.Ssocredential = types.StringNull()
	}
	if val, ok := getResponseData["ssodomain"]; ok && val != nil {
		data.Ssodomain = types.StringValue(val.(string))
	} else {
		data.Ssodomain = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
