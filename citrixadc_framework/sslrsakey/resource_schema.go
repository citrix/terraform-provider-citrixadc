package sslrsakey

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslrsakeyResourceModel describes the resource data model.
type SslrsakeyResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Aes256   types.Bool   `tfsdk:"aes256"`
	Bits     types.Int64  `tfsdk:"bits"`
	Des      types.Bool   `tfsdk:"des"`
	Des3     types.Bool   `tfsdk:"des3"`
	Exponent types.String `tfsdk:"exponent"`
	Keyfile  types.String `tfsdk:"keyfile"`
	Keyform  types.String `tfsdk:"keyform"`
	Password types.String `tfsdk:"password"`
	Pkcs8    types.Bool   `tfsdk:"pkcs8"`
}

func (r *SslrsakeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslrsakey resource.",
			},
			"aes256": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Encrypt the generated RSA key by using the AES algorithm.",
			},
			"bits": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Size, in bits, of the RSA key.",
			},
			"des": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Encrypt the generated RSA key by using the DES algorithm. On the command line, you are prompted to enter the pass phrase (password) that is used to encrypt the key.",
			},
			"des3": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Encrypt the generated RSA key by using the Triple-DES algorithm. On the command line, you are prompted to enter the pass phrase (password) that is used to encrypt the key.",
			},
			"exponent": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("F4"),
				Description: "Public exponent for the RSA key. The exponent is part of the cipher algorithm and is required for creating the RSA key.",
			},
			"keyfile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for and, optionally, path to the RSA key file. /nsconfig/ssl/ is the default path.",
			},
			"keyform": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("PEM"),
				Description: "Format in which the RSA key file is stored on the appliance.",
			},
			"password": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Pass phrase to use for encryption if DES or DES3 option is selected.",
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

func sslrsakeyGetThePayloadFromtheConfig(ctx context.Context, data *SslrsakeyResourceModel) ssl.Sslrsakey {
	tflog.Debug(ctx, "In sslrsakeyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslrsakey := ssl.Sslrsakey{}
	if !data.Aes256.IsNull() {
		sslrsakey.Aes256 = data.Aes256.ValueBool()
	}
	if !data.Bits.IsNull() {
		sslrsakey.Bits = utils.IntPtr(int(data.Bits.ValueInt64()))
	}
	if !data.Des.IsNull() {
		sslrsakey.Des = data.Des.ValueBool()
	}
	if !data.Des3.IsNull() {
		sslrsakey.Des3 = data.Des3.ValueBool()
	}
	if !data.Exponent.IsNull() {
		sslrsakey.Exponent = data.Exponent.ValueString()
	}
	if !data.Keyfile.IsNull() {
		sslrsakey.Keyfile = data.Keyfile.ValueString()
	}
	if !data.Keyform.IsNull() {
		sslrsakey.Keyform = data.Keyform.ValueString()
	}
	if !data.Password.IsNull() {
		sslrsakey.Password = data.Password.ValueString()
	}
	if !data.Pkcs8.IsNull() {
		sslrsakey.Pkcs8 = data.Pkcs8.ValueBool()
	}

	return sslrsakey
}

func sslrsakeySetAttrFromGet(ctx context.Context, data *SslrsakeyResourceModel, getResponseData map[string]interface{}) *SslrsakeyResourceModel {
	tflog.Debug(ctx, "In sslrsakeySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["aes256"]; ok && val != nil {
		data.Aes256 = types.BoolValue(val.(bool))
	} else {
		data.Aes256 = types.BoolNull()
	}
	if val, ok := getResponseData["bits"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Bits = types.Int64Value(intVal)
		}
	} else {
		data.Bits = types.Int64Null()
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
	if val, ok := getResponseData["exponent"]; ok && val != nil {
		data.Exponent = types.StringValue(val.(string))
	} else {
		data.Exponent = types.StringNull()
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
	if val, ok := getResponseData["password"]; ok && val != nil {
		data.Password = types.StringValue(val.(string))
	} else {
		data.Password = types.StringNull()
	}
	if val, ok := getResponseData["pkcs8"]; ok && val != nil {
		data.Pkcs8 = types.BoolValue(val.(bool))
	} else {
		data.Pkcs8 = types.BoolNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("sslrsakey-config")

	return data
}
