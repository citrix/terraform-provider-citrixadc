package authenticationazurekeyvault

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

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

// AuthenticationazurekeyvaultResourceModel describes the resource data model.
type AuthenticationazurekeyvaultResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Authentication             types.String `tfsdk:"authentication"`
	Clientid                   types.String `tfsdk:"clientid"`
	Clientsecret               types.String `tfsdk:"clientsecret"`
	ClientsecretWo             types.String `tfsdk:"clientsecret_wo"`
	ClientsecretWoVersion      types.Int64  `tfsdk:"clientsecret_wo_version"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Name                       types.String `tfsdk:"name"`
	Pushservice                types.String `tfsdk:"pushservice"`
	Refreshinterval            types.Int64  `tfsdk:"refreshinterval"`
	Servicekeyname             types.String `tfsdk:"servicekeyname"`
	Signaturealg               types.String `tfsdk:"signaturealg"`
	Tenantid                   types.String `tfsdk:"tenantid"`
	Tokenendpoint              types.String `tfsdk:"tokenendpoint"`
	Vaultname                  types.String `tfsdk:"vaultname"`
}

func (r *AuthenticationazurekeyvaultResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationazurekeyvault resource.",
			},
			"authentication": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "If authentication is disabled, otp checks are not performed after azure vault keys are obtained. This is useful to distinguish whether user has registered devices.",
			},
			"clientid": schema.StringAttribute{
				Required:    true,
				Description: "Unique identity of the relying party requesting for authentication.",
			},
			"clientsecret": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Unique secret string to authorize relying party at authorization server.",
			},
			"clientsecret_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Unique secret string to authorize relying party at authorization server.",
			},
			"clientsecret_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a clientsecret_wo update.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the group that is added to user sessions that match current IdP policy. It can be used in policies to identify relying party trust.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the new Azure Key Vault profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"pushservice": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the service used to send push notifications",
			},
			"refreshinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(50),
				Description: "Interval at which access token in obtained.",
			},
			"servicekeyname": schema.StringAttribute{
				Required:    true,
				Description: "Friendly name of the Key to be used to compute signature.",
			},
			"signaturealg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("RS256"),
				Description: "Algorithm to be used to sign/verify transactions",
			},
			"tenantid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "TenantID of the application. This is usually specific to providers such as Microsoft and usually refers to the deployment identifier.",
			},
			"tokenendpoint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL endpoint on relying party to which the OAuth token is to be sent.",
			},
			"vaultname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Azure vault account as configured in azure portal.",
			},
		},
	}
}

func authenticationazurekeyvaultGetThePayloadFromthePlan(ctx context.Context, data *AuthenticationazurekeyvaultResourceModel) authentication.Authenticationazurekeyvault {
	tflog.Debug(ctx, "In authenticationazurekeyvaultGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authenticationazurekeyvault := authentication.Authenticationazurekeyvault{}
	if !data.Authentication.IsNull() && !data.Authentication.IsUnknown() {
		authenticationazurekeyvault.Authentication = data.Authentication.ValueString()
	}
	if !data.Clientid.IsNull() && !data.Clientid.IsUnknown() {
		authenticationazurekeyvault.Clientid = data.Clientid.ValueString()
	}
	if !data.Clientsecret.IsNull() && !data.Clientsecret.IsUnknown() {
		authenticationazurekeyvault.Clientsecret = data.Clientsecret.ValueString()
	}
	// Skip write-only attribute: clientsecret_wo
	// Skip version tracker attribute: clientsecret_wo_version
	if !data.Defaultauthenticationgroup.IsNull() && !data.Defaultauthenticationgroup.IsUnknown() {
		authenticationazurekeyvault.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authenticationazurekeyvault.Name = data.Name.ValueString()
	}
	if !data.Pushservice.IsNull() && !data.Pushservice.IsUnknown() {
		authenticationazurekeyvault.Pushservice = data.Pushservice.ValueString()
	}
	if !data.Refreshinterval.IsNull() && !data.Refreshinterval.IsUnknown() {
		authenticationazurekeyvault.Refreshinterval = utils.IntPtr(int(data.Refreshinterval.ValueInt64()))
	}
	if !data.Servicekeyname.IsNull() && !data.Servicekeyname.IsUnknown() {
		authenticationazurekeyvault.Servicekeyname = data.Servicekeyname.ValueString()
	}
	if !data.Signaturealg.IsNull() && !data.Signaturealg.IsUnknown() {
		authenticationazurekeyvault.Signaturealg = data.Signaturealg.ValueString()
	}
	if !data.Tenantid.IsNull() && !data.Tenantid.IsUnknown() {
		authenticationazurekeyvault.Tenantid = data.Tenantid.ValueString()
	}
	if !data.Tokenendpoint.IsNull() && !data.Tokenendpoint.IsUnknown() {
		authenticationazurekeyvault.Tokenendpoint = data.Tokenendpoint.ValueString()
	}
	if !data.Vaultname.IsNull() && !data.Vaultname.IsUnknown() {
		authenticationazurekeyvault.Vaultname = data.Vaultname.ValueString()
	}

	return authenticationazurekeyvault
}

func authenticationazurekeyvaultGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationazurekeyvaultResourceModel, payload *authentication.Authenticationazurekeyvault) {
	tflog.Debug(ctx, "In authenticationazurekeyvaultGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: clientsecret_wo -> clientsecret
	if !data.ClientsecretWo.IsNull() {
		clientsecretWo := data.ClientsecretWo.ValueString()
		if clientsecretWo != "" {
			payload.Clientsecret = clientsecretWo
		}
	}
}

func authenticationazurekeyvaultSetAttrFromGet(ctx context.Context, data *AuthenticationazurekeyvaultResourceModel, getResponseData map[string]interface{}) *AuthenticationazurekeyvaultResourceModel {
	tflog.Debug(ctx, "In authenticationazurekeyvaultSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["authentication"]; ok && val != nil {
		data.Authentication = types.StringValue(val.(string))
	} else {
		data.Authentication = types.StringNull()
	}
	if val, ok := getResponseData["clientid"]; ok && val != nil {
		data.Clientid = types.StringValue(val.(string))
	} else {
		data.Clientid = types.StringNull()
	}
	// clientsecret is not returned by NITRO API (secret/ephemeral) - retain from config
	// clientsecret_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// clientsecret_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["defaultauthenticationgroup"]; ok && val != nil {
		data.Defaultauthenticationgroup = types.StringValue(val.(string))
	} else {
		data.Defaultauthenticationgroup = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["pushservice"]; ok && val != nil {
		data.Pushservice = types.StringValue(val.(string))
	} else {
		data.Pushservice = types.StringNull()
	}
	if val, ok := getResponseData["refreshinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Refreshinterval = types.Int64Value(intVal)
		}
	} else {
		data.Refreshinterval = types.Int64Null()
	}
	if val, ok := getResponseData["servicekeyname"]; ok && val != nil {
		data.Servicekeyname = types.StringValue(val.(string))
	} else {
		data.Servicekeyname = types.StringNull()
	}
	if val, ok := getResponseData["signaturealg"]; ok && val != nil {
		data.Signaturealg = types.StringValue(val.(string))
	} else {
		data.Signaturealg = types.StringNull()
	}
	if val, ok := getResponseData["tenantid"]; ok && val != nil {
		data.Tenantid = types.StringValue(val.(string))
	} else {
		data.Tenantid = types.StringNull()
	}
	if val, ok := getResponseData["tokenendpoint"]; ok && val != nil {
		data.Tokenendpoint = types.StringValue(val.(string))
	} else {
		data.Tokenendpoint = types.StringNull()
	}
	if val, ok := getResponseData["vaultname"]; ok && val != nil {
		data.Vaultname = types.StringValue(val.(string))
	} else {
		data.Vaultname = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
