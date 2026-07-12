package sslecdsakey

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslecdsakeyResourceModel describes the resource data model.
type SslecdsakeyResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Aes256            types.Bool   `tfsdk:"aes256"`
	Curve             types.String `tfsdk:"curve"`
	Des               types.Bool   `tfsdk:"des"`
	Des3              types.Bool   `tfsdk:"des3"`
	Keyfile           types.String `tfsdk:"keyfile"`
	Keyform           types.String `tfsdk:"keyform"`
	Password          types.String `tfsdk:"password"`
	PasswordWo        types.String `tfsdk:"password_wo"`
	PasswordWoVersion types.Int64  `tfsdk:"password_wo_version"`
	Pkcs8             types.Bool   `tfsdk:"pkcs8"`
}

func (r *SslecdsakeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslecdsakey resource.",
			},
			"aes256": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Encrypt the generated ECDSA key by using the AES algorithm.",
			},
			"curve": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Curve id to generate ECDSA key. Only P_256 and P_384 are supported",
			},
			"des": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Encrypt the generated ECDSA key by using the DES algorithm. On the command line, you are prompted to enter the pass phrase (password) that is used to encrypt the key.",
			},
			"des3": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Encrypt the generated ECDSA key by using the Triple-DES algorithm. On the command line, you are prompted to enter the pass phrase (password) that is used to encrypt the key.",
			},
			"keyfile": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for and, optionally, path to the ECDSA key file. /nsconfig/ssl/ is the default path.",
			},
			"keyform": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Format in which the ECDSA key file is stored on the appliance.",
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Pass phrase to use for encryption if DES or DES3 option is selected.",
			},
			"password_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Pass phrase to use for encryption if DES or DES3 option is selected.",
			},
			"password_wo_version": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Increment this version to signal a password_wo update.",
			},
			"pkcs8": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Create the private key in PKCS#8 format.",
			},
		},
	}
}

func sslecdsakeyGetThePayloadFromthePlan(ctx context.Context, data *SslecdsakeyResourceModel) ssl.Sslecdsakey {
	tflog.Debug(ctx, "In sslecdsakeyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslecdsakey := ssl.Sslecdsakey{}
	if !data.Aes256.IsNull() && !data.Aes256.IsUnknown() {
		sslecdsakey.Aes256 = data.Aes256.ValueBool()
	}
	if !data.Curve.IsNull() && !data.Curve.IsUnknown() {
		sslecdsakey.Curve = data.Curve.ValueString()
	}
	if !data.Des.IsNull() && !data.Des.IsUnknown() {
		sslecdsakey.Des = data.Des.ValueBool()
	}
	if !data.Des3.IsNull() && !data.Des3.IsUnknown() {
		sslecdsakey.Des3 = data.Des3.ValueBool()
	}
	if !data.Keyfile.IsNull() && !data.Keyfile.IsUnknown() {
		sslecdsakey.Keyfile = data.Keyfile.ValueString()
	}
	if !data.Keyform.IsNull() && !data.Keyform.IsUnknown() {
		sslecdsakey.Keyform = data.Keyform.ValueString()
	}
	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		sslecdsakey.Password = data.Password.ValueString()
	}
	// Skip write-only attribute: password_wo
	// Skip version tracker attribute: password_wo_version
	if !data.Pkcs8.IsNull() && !data.Pkcs8.IsUnknown() {
		sslecdsakey.Pkcs8 = data.Pkcs8.ValueBool()
	}

	return sslecdsakey
}

func sslecdsakeyGetThePayloadFromtheConfig(ctx context.Context, data *SslecdsakeyResourceModel, payload *ssl.Sslecdsakey) {
	tflog.Debug(ctx, "In sslecdsakeyGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: password_wo -> password
	if !data.PasswordWo.IsNull() {
		passwordWo := data.PasswordWo.ValueString()
		if passwordWo != "" {
			payload.Password = passwordWo
		}
	}
}

func sslecdsakeySetAttrFromGet(ctx context.Context, data *SslecdsakeyResourceModel, getResponseData map[string]interface{}) *SslecdsakeyResourceModel {
	tflog.Debug(ctx, "In sslecdsakeySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["aes256"]; ok && val != nil {
		data.Aes256 = types.BoolValue(val.(bool))
	} else {
		data.Aes256 = types.BoolNull()
	}
	if val, ok := getResponseData["curve"]; ok && val != nil {
		data.Curve = types.StringValue(val.(string))
	} else {
		data.Curve = types.StringNull()
	}
	if val, ok := getResponseData["des"]; ok && val != nil {
		data.Des = types.BoolValue(val.(bool))
	} else {
		data.Des = types.BoolNull()
	}
	if val, ok := getResponseData["des3"]; ok && val != nil {
		data.Des3 = types.BoolValue(val.(bool))
	} else {
		data.Des3 = types.BoolNull()
	}
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
	if val, ok := getResponseData["pkcs8"]; ok && val != nil {
		data.Pkcs8 = types.BoolValue(val.(bool))
	} else {
		data.Pkcs8 = types.BoolNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("sslecdsakey-config")

	return data
}
