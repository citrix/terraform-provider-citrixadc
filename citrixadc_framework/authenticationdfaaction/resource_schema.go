package authenticationdfaaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationdfaactionResourceModel describes the resource data model.
type AuthenticationdfaactionResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Clientid                   types.String `tfsdk:"clientid"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Name                       types.String `tfsdk:"name"`
	Passphrase                 types.String `tfsdk:"passphrase"`
	PassphraseWo               types.String `tfsdk:"passphrase_wo"`
	PassphraseWoVersion        types.Int64  `tfsdk:"passphrase_wo_version"`
	Serverurl                  types.String `tfsdk:"serverurl"`
}

func (r *AuthenticationdfaactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationdfaaction resource.",
			},
			"clientid": schema.StringAttribute{
				Required:    true,
				Description: "If configured, this string is sent to the DFA server as the X-Citrix-Exchange header value.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the DFA action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the DFA action is added.",
			},
			"passphrase": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Key shared between the DFA server and the Citrix ADC.\nRequired to allow the Citrix ADC to communicate with the DFA server.",
			},
			"passphrase_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Key shared between the DFA server and the Citrix ADC.\nRequired to allow the Citrix ADC to communicate with the DFA server.",
			},
			"passphrase_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a passphrase_wo update.",
			},
			"serverurl": schema.StringAttribute{
				Required:    true,
				Description: "DFA Server URL",
			},
		},
	}
}

func authenticationdfaactionGetThePayloadFromthePlan(ctx context.Context, data *AuthenticationdfaactionResourceModel) authentication.Authenticationdfaaction {
	tflog.Debug(ctx, "In authenticationdfaactionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authenticationdfaaction := authentication.Authenticationdfaaction{}
	if !data.Clientid.IsNull() && !data.Clientid.IsUnknown() {
		authenticationdfaaction.Clientid = data.Clientid.ValueString()
	}
	if !data.Defaultauthenticationgroup.IsNull() && !data.Defaultauthenticationgroup.IsUnknown() {
		authenticationdfaaction.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authenticationdfaaction.Name = data.Name.ValueString()
	}
	if !data.Passphrase.IsNull() && !data.Passphrase.IsUnknown() {
		authenticationdfaaction.Passphrase = data.Passphrase.ValueString()
	}
	// Skip write-only attribute: passphrase_wo
	// Skip version tracker attribute: passphrase_wo_version
	if !data.Serverurl.IsNull() && !data.Serverurl.IsUnknown() {
		authenticationdfaaction.Serverurl = data.Serverurl.ValueString()
	}

	return authenticationdfaaction
}

func authenticationdfaactionGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationdfaactionResourceModel, payload *authentication.Authenticationdfaaction) {
	tflog.Debug(ctx, "In authenticationdfaactionGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: passphrase_wo -> passphrase
	if !data.PassphraseWo.IsNull() {
		passphraseWo := data.PassphraseWo.ValueString()
		if passphraseWo != "" {
			payload.Passphrase = passphraseWo
		}
	}
}

func authenticationdfaactionSetAttrFromGet(ctx context.Context, data *AuthenticationdfaactionResourceModel, getResponseData map[string]interface{}) *AuthenticationdfaactionResourceModel {
	tflog.Debug(ctx, "In authenticationdfaactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["clientid"]; ok && val != nil {
		data.Clientid = types.StringValue(val.(string))
	} else {
		data.Clientid = types.StringNull()
	}
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
	// passphrase is not returned by NITRO API (secret/ephemeral) - retain from config
	// passphrase_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// passphrase_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["serverurl"]; ok && val != nil {
		data.Serverurl = types.StringValue(val.(string))
	} else {
		data.Serverurl = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
