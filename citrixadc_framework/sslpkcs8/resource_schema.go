package sslpkcs8

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Sslpkcs8ResourceModel describes the resource data model.
type Sslpkcs8ResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Keyfile           types.String `tfsdk:"keyfile"`
	Keyform           types.String `tfsdk:"keyform"`
	Password          types.String `tfsdk:"password"`
	PasswordWo        types.String `tfsdk:"password_wo"`
	PasswordWoVersion types.Int64  `tfsdk:"password_wo_version"`
	Pkcs8file         types.String `tfsdk:"pkcs8file"`
}

func (r *Sslpkcs8Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslpkcs8 resource.",
			},
			"keyfile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the input key file to be converted from PEM or DER format to PKCS#8 format. /nsconfig/ssl/ is the default path.",
			},
			"keyform": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("PEM"),
				Description: "Format in which the key file is stored on the appliance.",
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Password to assign to the file if the key is encrypted. Applies only for PEM format files.",
			},
			"password_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Password to assign to the file if the key is encrypted. Applies only for PEM format files.",
			},
			"password_wo_version": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a password_wo update.",
			},
			"pkcs8file": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for and, optionally, path to, the output file where the PKCS#8 format key file is stored. /nsconfig/ssl/ is the default path.",
			},
		},
	}
}

func sslpkcs8GetThePayloadFromthePlan(ctx context.Context, data *Sslpkcs8ResourceModel) ssl.Sslpkcs8 {
	tflog.Debug(ctx, "In sslpkcs8GetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslpkcs8 := ssl.Sslpkcs8{}
	if !data.Keyfile.IsNull() && !data.Keyfile.IsUnknown() {
		sslpkcs8.Keyfile = data.Keyfile.ValueString()
	}
	if !data.Keyform.IsNull() && !data.Keyform.IsUnknown() {
		sslpkcs8.Keyform = data.Keyform.ValueString()
	}
	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		sslpkcs8.Password = data.Password.ValueString()
	}
	// Skip write-only attribute: password_wo
	// Skip version tracker attribute: password_wo_version
	if !data.Pkcs8file.IsNull() && !data.Pkcs8file.IsUnknown() {
		sslpkcs8.Pkcs8file = data.Pkcs8file.ValueString()
	}

	return sslpkcs8
}

func sslpkcs8GetThePayloadFromtheConfig(ctx context.Context, data *Sslpkcs8ResourceModel, payload *ssl.Sslpkcs8) {
	tflog.Debug(ctx, "In sslpkcs8GetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: password_wo -> password
	if !data.PasswordWo.IsNull() {
		passwordWo := data.PasswordWo.ValueString()
		if passwordWo != "" {
			payload.Password = passwordWo
		}
	}
}

func sslpkcs8SetAttrFromGet(ctx context.Context, data *Sslpkcs8ResourceModel, getResponseData map[string]interface{}) *Sslpkcs8ResourceModel {
	tflog.Debug(ctx, "In sslpkcs8SetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["keyfile"]; ok && val != nil {
		data.Keyfile = types.StringValue(val.(string))
	} else {
		data.Keyfile = types.StringNull()
	}
	if val, ok := getResponseData["keyform"]; ok && val != nil {
		data.Keyform = types.StringValue(val.(string))
	} else {
		data.Keyform = types.StringNull()
	}
	// password is not returned by NITRO API (secret/ephemeral) - retain from config
	// password_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// password_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["pkcs8file"]; ok && val != nil {
		data.Pkcs8file = types.StringValue(val.(string))
	} else {
		data.Pkcs8file = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("sslpkcs8-config")

	return data
}
