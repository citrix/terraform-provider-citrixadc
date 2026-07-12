package dnskey

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// DnskeyResourceModel describes the resource data model.
type DnskeyResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Algorithm          types.String `tfsdk:"algorithm"`
	Autorollover       types.String `tfsdk:"autorollover"`
	Expires            types.Int64  `tfsdk:"expires"`
	Filenameprefix     types.String `tfsdk:"filenameprefix"`
	Keyname            types.String `tfsdk:"keyname"`
	Keysize            types.Int64  `tfsdk:"keysize"`
	Keytype            types.String `tfsdk:"keytype"`
	Notificationperiod types.Int64  `tfsdk:"notificationperiod"`
	Password           types.String `tfsdk:"password"`
	PasswordWo         types.String `tfsdk:"password_wo"`
	PasswordWoVersion  types.Int64  `tfsdk:"password_wo_version"`
	Privatekey         types.String `tfsdk:"privatekey"`
	Publickey          types.String `tfsdk:"publickey"`
	Revoke             types.Bool   `tfsdk:"revoke"`
	Rollovermethod     types.String `tfsdk:"rollovermethod"`
	Src                types.String `tfsdk:"src"`
	Ttl                types.Int64  `tfsdk:"ttl"`
	Units1             types.String `tfsdk:"units1"`
	Units2             types.String `tfsdk:"units2"`
	Zonename           types.String `tfsdk:"zonename"`
}

func (r *DnskeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnskey resource.",
			},
			"algorithm": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Algorithm to generate the key.",
			},
			"autorollover": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag to enable/disable key rollover automatically.\nNote:\n* Key name will be appended with _AR1 for successor key. For e.g. current key=k1, successor key=k1_AR1.\n* Key name can be truncated if current name length is more than 58 bytes to accomodate the suffix.",
			},
			"expires": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time period for which to consider the key valid, after the key is used to sign a zone.",
			},
			"filenameprefix": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Common prefix for the names of the generated public and private key files and the Delegation Signer (DS) resource record. During key generation, the .key, .private, and .ds suffixes are appended automatically to the file name prefix to produce the names of the public key, the private key, and the DS record, respectively.",
			},
			"keyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the public-private key pair to publish in the zone.",
			},
			"keysize": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Size of the key, in bits.",
			},
			"keytype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of key to create.",
			},
			"notificationperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time at which to generate notification of key expiration, specified as number of days, hours, or minutes before expiry. Must be less than the expiry period. The notification is an SNMP trap sent to an SNMP manager. To enable the appliance to send the trap, enable the DNSKEY-EXPIRY SNMP alarm. \nIn case autorollover option is enabled, rollover for successor key will be intiated at this time. No notification trap will be sent.",
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Passphrase for reading the encrypted public/private DNS keys",
			},
			"password_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Passphrase for reading the encrypted public/private DNS keys",
			},
			"password_wo_version": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Increment this version to signal a password_wo update.",
			},
			"privatekey": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "File name of the private key.",
			},
			"publickey": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "File name of the public key.",
			},
			"revoke": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Revoke the key. Note: This operation is non-reversible.",
			},
			"rollovermethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Method used for automatic rollover.\n* Key type: ZSK, Method: PrePublication or DoubleSignature.\n* Key type: KSK, Method: DoubleRRSet.",
			},
			"src": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL (protocol, host, path, and file name) from where the DNS key file will be imported. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access. This is a mandatory argument",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to Live (TTL), in seconds, for the DNSKEY resource record created in the zone. TTL is the time for which the record must be cached by the DNS proxies. If the TTL is not specified, either the DNS zone's minimum TTL or the default value of 3600 is used.",
			},
			"units1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Units for the expiry period.",
			},
			"units2": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Units for the notification period.",
			},
			"zonename": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the zone for which to create a key.",
			},
		},
	}
}

func dnskeyGetThePayloadFromthePlan(ctx context.Context, data *DnskeyResourceModel) dns.Dnskey {
	tflog.Debug(ctx, "In dnskeyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	dnskey := dns.Dnskey{}
	if !data.Algorithm.IsNull() && !data.Algorithm.IsUnknown() {
		dnskey.Algorithm = data.Algorithm.ValueString()
	}
	if !data.Autorollover.IsNull() && !data.Autorollover.IsUnknown() {
		dnskey.Autorollover = data.Autorollover.ValueString()
	}
	if !data.Expires.IsNull() && !data.Expires.IsUnknown() {
		dnskey.Expires = utils.IntPtr(int(data.Expires.ValueInt64()))
	}
	if !data.Filenameprefix.IsNull() && !data.Filenameprefix.IsUnknown() {
		dnskey.Filenameprefix = data.Filenameprefix.ValueString()
	}
	if !data.Keyname.IsNull() && !data.Keyname.IsUnknown() {
		dnskey.Keyname = data.Keyname.ValueString()
	}
	if !data.Keysize.IsNull() && !data.Keysize.IsUnknown() {
		dnskey.Keysize = utils.IntPtr(int(data.Keysize.ValueInt64()))
	}
	if !data.Keytype.IsNull() && !data.Keytype.IsUnknown() {
		dnskey.Keytype = data.Keytype.ValueString()
	}
	if !data.Notificationperiod.IsNull() && !data.Notificationperiod.IsUnknown() {
		dnskey.Notificationperiod = utils.IntPtr(int(data.Notificationperiod.ValueInt64()))
	}
	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		dnskey.Password = data.Password.ValueString()
	}
	// Skip write-only attribute: password_wo
	// Skip version tracker attribute: password_wo_version
	if !data.Privatekey.IsNull() && !data.Privatekey.IsUnknown() {
		dnskey.Privatekey = data.Privatekey.ValueString()
	}
	if !data.Publickey.IsNull() && !data.Publickey.IsUnknown() {
		dnskey.Publickey = data.Publickey.ValueString()
	}
	if !data.Revoke.IsNull() && !data.Revoke.IsUnknown() {
		dnskey.Revoke = data.Revoke.ValueBool()
	}
	if !data.Rollovermethod.IsNull() && !data.Rollovermethod.IsUnknown() {
		dnskey.Rollovermethod = data.Rollovermethod.ValueString()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		dnskey.Src = data.Src.ValueString()
	}
	if !data.Ttl.IsNull() && !data.Ttl.IsUnknown() {
		dnskey.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}
	if !data.Units1.IsNull() && !data.Units1.IsUnknown() {
		dnskey.Units1 = data.Units1.ValueString()
	}
	if !data.Units2.IsNull() && !data.Units2.IsUnknown() {
		dnskey.Units2 = data.Units2.ValueString()
	}
	if !data.Zonename.IsNull() && !data.Zonename.IsUnknown() {
		dnskey.Zonename = data.Zonename.ValueString()
	}

	return dnskey
}

func dnskeyGetThePayloadFromtheConfig(ctx context.Context, data *DnskeyResourceModel, payload *dns.Dnskey) {
	tflog.Debug(ctx, "In dnskeyGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: password_wo -> password
	if !data.PasswordWo.IsNull() {
		passwordWo := data.PasswordWo.ValueString()
		if passwordWo != "" {
			payload.Password = passwordWo
		}
	}
}

func dnskeySetAttrFromGet(ctx context.Context, data *DnskeyResourceModel, getResponseData map[string]interface{}) *DnskeyResourceModel {
	tflog.Debug(ctx, "In dnskeySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["algorithm"]; ok && val != nil {
		data.Algorithm = types.StringValue(val.(string))
	} else {
		data.Algorithm = types.StringNull()
	}
	if val, ok := getResponseData["autorollover"]; ok && val != nil {
		data.Autorollover = types.StringValue(val.(string))
	} else {
		data.Autorollover = types.StringNull()
	}
	if val, ok := getResponseData["expires"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Expires = types.Int64Value(intVal)
		}
	} else {
		data.Expires = types.Int64Null()
	}
	if val, ok := getResponseData["filenameprefix"]; ok && val != nil {
		data.Filenameprefix = types.StringValue(val.(string))
	} else {
		data.Filenameprefix = types.StringNull()
	}
	if val, ok := getResponseData["keyname"]; ok && val != nil {
		data.Keyname = types.StringValue(val.(string))
	} else {
		data.Keyname = types.StringNull()
	}
	if val, ok := getResponseData["keysize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Keysize = types.Int64Value(intVal)
		}
	} else {
		data.Keysize = types.Int64Null()
	}
	if val, ok := getResponseData["keytype"]; ok && val != nil {
		data.Keytype = types.StringValue(val.(string))
	} else {
		data.Keytype = types.StringNull()
	}
	if val, ok := getResponseData["notificationperiod"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Notificationperiod = types.Int64Value(intVal)
		}
	} else {
		data.Notificationperiod = types.Int64Null()
	}
	// password is not returned by NITRO API (secret/ephemeral) - retain from config
	// password_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// password_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["privatekey"]; ok && val != nil {
		data.Privatekey = types.StringValue(val.(string))
	} else {
		data.Privatekey = types.StringNull()
	}
	if val, ok := getResponseData["publickey"]; ok && val != nil {
		data.Publickey = types.StringValue(val.(string))
	} else {
		data.Publickey = types.StringNull()
	}
	if val, ok := getResponseData["revoke"]; ok && val != nil {
		data.Revoke = types.BoolValue(val.(bool))
	} else {
		data.Revoke = types.BoolNull()
	}
	if val, ok := getResponseData["rollovermethod"]; ok && val != nil {
		data.Rollovermethod = types.StringValue(val.(string))
	} else {
		data.Rollovermethod = types.StringNull()
	}
	if val, ok := getResponseData["src"]; ok && val != nil {
		data.Src = types.StringValue(val.(string))
	} else {
		data.Src = types.StringNull()
	}
	if val, ok := getResponseData["ttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ttl = types.Int64Value(intVal)
		}
	} else {
		data.Ttl = types.Int64Null()
	}
	if val, ok := getResponseData["units1"]; ok && val != nil {
		data.Units1 = types.StringValue(val.(string))
	} else {
		data.Units1 = types.StringNull()
	}
	if val, ok := getResponseData["units2"]; ok && val != nil {
		data.Units2 = types.StringValue(val.(string))
	} else {
		data.Units2 = types.StringNull()
	}
	if val, ok := getResponseData["zonename"]; ok && val != nil {
		data.Zonename = types.StringValue(val.(string))
	} else {
		data.Zonename = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Keyname.ValueString()))

	return data
}
