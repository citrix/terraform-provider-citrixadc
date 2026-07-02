package sslpkcs12

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Sslpkcs12ResourceModel describes the resource data model.
type Sslpkcs12ResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Import                 types.Bool   `tfsdk:"import"`
	Aes256                 types.Bool   `tfsdk:"aes256"`
	Certfile               types.String `tfsdk:"certfile"`
	Des                    types.Bool   `tfsdk:"des"`
	Des3                   types.Bool   `tfsdk:"des3"`
	Export                 types.Bool   `tfsdk:"export"`
	Keyfile                types.String `tfsdk:"keyfile"`
	Outfile                types.String `tfsdk:"outfile"`
	Password               types.String `tfsdk:"password"`
	PasswordWo             types.String `tfsdk:"password_wo"`
	PasswordWoVersion      types.Int64  `tfsdk:"password_wo_version"`
	Pempassphrase          types.String `tfsdk:"pempassphrase"`
	PempassphraseWo        types.String `tfsdk:"pempassphrase_wo"`
	PempassphraseWoVersion types.Int64  `tfsdk:"pempassphrase_wo_version"`
	Pkcs12file             types.String `tfsdk:"pkcs12file"`
}

func (r *Sslpkcs12Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslpkcs12 resource.",
			},
			"import": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Convert the certificate and private-key from PKCS#12 format to PEM format.",
			},
			"aes256": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Encrypt the private key by using the AES algorithm (256-bit key) during the import operation. On the command line, you are prompted to enter the pass phrase.",
			},
			"certfile": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Certificate file to be converted from PEM to PKCS#12 format.",
			},
			"des": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Encrypt the private key by using the DES algorithm in CBC mode during the import operation. On the command line, you are prompted to enter the pass phrase.",
			},
			"des3": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Encrypt the private key by using the Triple-DES algorithm in EDE CBC mode (168-bit key) during the import operation. On the command line, you are prompted to enter the pass phrase.",
			},
			"export": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Convert the certificate and private key from PEM format to PKCS#12 format. On the command line, you are prompted to enter the pass phrase.",
			},
			"keyfile": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the private key file to be converted from PEM to PKCS#12 format. If the key file is encrypted, you are prompted to enter the pass phrase used for encrypting the key.",
			},
			"outfile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for and, optionally, path to, the output file that contains the certificate and the private key after converting from PKCS#12 to PEM format. /nsconfig/ssl/ is the default path.\nIf importing, the certificate-key pair is stored in PEM format. If exporting, the certificate-key pair is stored in PKCS#12 format.",
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
			"password_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
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
			"pempassphrase": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
			"pempassphrase_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
			"pempassphrase_wo_version": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a pempassphrase_wo update.",
			},
			"pkcs12file": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for and, optionally, path to, the PKCS#12 file. If importing, specify the input file name that contains the certificate and the private key in PKCS#12 format. If exporting, specify the output file name that contains the certificate and the private key after converting from PEM to\nPKCS#12 format. /nsconfig/ssl/ is the default path.\nDuring the import operation, if the key is encrypted, you are prompted to enter the pass phrase used for encrypting the key.",
			},
		},
	}
}

func sslpkcs12GetThePayloadFromthePlan(ctx context.Context, data *Sslpkcs12ResourceModel) ssl.Sslpkcs12 {
	tflog.Debug(ctx, "In sslpkcs12GetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslpkcs12 := ssl.Sslpkcs12{}
	if !data.Import.IsNull() && !data.Import.IsUnknown() {
		sslpkcs12.Import = data.Import.ValueBool()
	}
	if !data.Aes256.IsNull() && !data.Aes256.IsUnknown() {
		sslpkcs12.Aes256 = data.Aes256.ValueBool()
	}
	if !data.Certfile.IsNull() && !data.Certfile.IsUnknown() {
		sslpkcs12.Certfile = data.Certfile.ValueString()
	}
	if !data.Des.IsNull() && !data.Des.IsUnknown() {
		sslpkcs12.Des = data.Des.ValueBool()
	}
	if !data.Des3.IsNull() && !data.Des3.IsUnknown() {
		sslpkcs12.Des3 = data.Des3.ValueBool()
	}
	if !data.Export.IsNull() && !data.Export.IsUnknown() {
		sslpkcs12.Export = data.Export.ValueBool()
	}
	if !data.Keyfile.IsNull() && !data.Keyfile.IsUnknown() {
		sslpkcs12.Keyfile = data.Keyfile.ValueString()
	}
	if !data.Outfile.IsNull() && !data.Outfile.IsUnknown() {
		sslpkcs12.Outfile = data.Outfile.ValueString()
	}
	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		sslpkcs12.Password = data.Password.ValueString()
	}
	// Skip write-only attribute: password_wo
	// Skip version tracker attribute: password_wo_version
	if !data.Pempassphrase.IsNull() && !data.Pempassphrase.IsUnknown() {
		sslpkcs12.Pempassphrase = data.Pempassphrase.ValueString()
	}
	// Skip write-only attribute: pempassphrase_wo
	// Skip version tracker attribute: pempassphrase_wo_version
	if !data.Pkcs12file.IsNull() && !data.Pkcs12file.IsUnknown() {
		sslpkcs12.Pkcs12file = data.Pkcs12file.ValueString()
	}

	return sslpkcs12
}

func sslpkcs12GetThePayloadFromtheConfig(ctx context.Context, data *Sslpkcs12ResourceModel, payload *ssl.Sslpkcs12) {
	tflog.Debug(ctx, "In sslpkcs12GetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: password_wo -> password
	if !data.PasswordWo.IsNull() {
		passwordWo := data.PasswordWo.ValueString()
		if passwordWo != "" {
			payload.Password = passwordWo
		}
	}
	// Handle write-only secret attribute: pempassphrase_wo -> pempassphrase
	if !data.PempassphraseWo.IsNull() {
		pempassphraseWo := data.PempassphraseWo.ValueString()
		if pempassphraseWo != "" {
			payload.Pempassphrase = pempassphraseWo
		}
	}
}

func sslpkcs12SetAttrFromGet(ctx context.Context, data *Sslpkcs12ResourceModel, getResponseData map[string]interface{}) *Sslpkcs12ResourceModel {
	tflog.Debug(ctx, "In sslpkcs12SetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["Import"]; ok && val != nil {
		data.Import = types.BoolValue(val.(bool))
	} else {
		data.Import = types.BoolNull()
	}
	if val, ok := getResponseData["aes256"]; ok && val != nil {
		data.Aes256 = types.BoolValue(val.(bool))
	} else {
		data.Aes256 = types.BoolNull()
	}
	if val, ok := getResponseData["certfile"]; ok && val != nil {
		data.Certfile = types.StringValue(val.(string))
	} else {
		data.Certfile = types.StringNull()
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
	if val, ok := getResponseData["export"]; ok && val != nil {
		data.Export = types.BoolValue(val.(bool))
	} else {
		data.Export = types.BoolNull()
	}
	if val, ok := getResponseData["keyfile"]; ok && val != nil {
		data.Keyfile = types.StringValue(val.(string))
	} else {
		data.Keyfile = types.StringNull()
	}
	if val, ok := getResponseData["outfile"]; ok && val != nil {
		data.Outfile = types.StringValue(val.(string))
	} else {
		data.Outfile = types.StringNull()
	}
	// password is not returned by NITRO API (secret/ephemeral) - retain from config
	// password_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// password_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	// pempassphrase is not returned by NITRO API (secret/ephemeral) - retain from config
	// pempassphrase_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// pempassphrase_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["pkcs12file"]; ok && val != nil {
		data.Pkcs12file = types.StringValue(val.(string))
	} else {
		data.Pkcs12file = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("sslpkcs12-config")

	return data
}
