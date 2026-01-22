package sslcertkey

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslCertKeyResourceModel describes the resource data model.
type SslCertKeyResourceModel struct {
	Id                          types.String `tfsdk:"id"`
	Certkey                     types.String `tfsdk:"certkey"`
	Cert                        types.String `tfsdk:"cert"`
	Key                         types.String `tfsdk:"key"`
	Password                    types.Bool   `tfsdk:"password"`
	Fipskey                     types.String `tfsdk:"fipskey"`
	Hsmkey                      types.String `tfsdk:"hsmkey"`
	Inform                      types.String `tfsdk:"inform"`
	Expirymonitor               types.String `tfsdk:"expirymonitor"`
	NotificationPeriod          types.Int64  `tfsdk:"notificationperiod"`
	Bundle                      types.String `tfsdk:"bundle"`
	LinkCertKeyName             types.String `tfsdk:"linkcertkeyname"`
	NoDomainCheck               types.Bool   `tfsdk:"nodomaincheck"`
	OcspStaplingCache           types.Bool   `tfsdk:"ocspstaplingcache"`
	DeleteFromDevice            types.Bool   `tfsdk:"deletefromdevice"`
	DeleteCertKeyFilesOnRemoval types.String `tfsdk:"deletecertkeyfilesonremoval"`
	Passplain                   types.String `tfsdk:"passplain"`
	PassplainWo                 types.String `tfsdk:"passplain_wo"`
	PassplainWoVersion          types.Int64  `tfsdk:"passplain_wo_version"`
	CertHash                    types.String `tfsdk:"cert_hash"`
	KeyHash                     types.String `tfsdk:"key_hash"`
}

func (r *SslCertKeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the SSL certificate key pair. This is the same as the certkey attribute.",
			},
			"certkey": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the certificate and private-key pair. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following fields cannot be changed after creation: certkey, bundle, hsmkey.",
			},
			"cert": schema.StringAttribute{
				Required:    true,
				Description: "Name of and, optionally, path to the X509 certificate file that is used to form the certificate-key pair. The certificate file should be present on the appliance's hard-disk drive or solid-state drive. Storing a certificate in any location other than the default might cause inconsistency in a high availability setup. /nsconfig/ssl/ is the default path.",
			},
			"key": schema.StringAttribute{
				Optional:    true,
				Description: "Name of and, optionally, path to the private-key file that is used to form the certificate-key pair. The certificate file should be present on the appliance's hard-disk drive or solid-state drive. Storing a certificate in any location other than the default might cause inconsistency in a high availability setup. /nsconfig/ssl/ is the default path.",
			},
			"password": schema.BoolAttribute{
				Optional:    true,
				Description: "Passphrase that was used to encrypt the private-key. Use this option to load encrypted private-keys in PEM format.",
			},
			"fipskey": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the FIPS key that was created inside the Hardware Security Module (HSM) of a FIPS appliance, or a key that was imported into the HSM.",
			},
			"hsmkey": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the HSM key that was created in the External Hardware Security Module (HSM) of a FIPS appliance. The following fields cannot be changed after creation: certkey, bundle, hsmkey.",
			},
			"inform": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Input format of the certificate and the private-key files. The three formats supported by the appliance are: PEM - Privacy Enhanced Mail, DER - Distinguished Encoding Rule, PFX - Personal Information Exchange. Default: PEM",
			},
			"expirymonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Issue an alert when the certificate is about to expire. Possible values: ENABLED, DISABLED",
			},
			"notificationperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in number of days, before certificate expiration, at which to generate an alert that the certificate is about to expire. Minimum value: 10, Maximum value: 100",
			},
			"bundle": schema.StringAttribute{
				Optional:    true,
				Description: "Parse the certificate chain as a single file after linking the server certificate to its issuer's certificate within the file. Possible values: YES, NO. The following fields cannot be changed after creation: certkey, bundle, hsmkey.",
			},
			"linkcertkeyname": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the Certificate Authority certificate-key pair to which to link a certificate-key pair.",
			},
			"nodomaincheck": schema.BoolAttribute{
				Optional:    true,
				Description: "Override the check for matching domain names during a certificate update operation.",
			},
			"ocspstaplingcache": schema.BoolAttribute{
				Optional:    true,
				Description: "Clear cached ocspStapling response in case of an update operation.",
			},
			"deletefromdevice": schema.BoolAttribute{
				Optional:    true,
				Description: "Delete cert/key file from file system. Possible values: true, false",
			},
			"deletecertkeyfilesonremoval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Delete certificate and key files when the certificate is removed. Possible values: YES, NO",
			},
			"passplain": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Pass phrase used to encrypt the private-key. Required when adding an encrypted private-key in PEM format.",
			},
			"passplain_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Pass phrase used to encrypt the private-key. Required when adding an encrypted private-key in PEM format.",
			},
			"passplain_wo_version": schema.Int64Attribute{
				Description: "Increment this version to signal a passplain_wo update.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
			},
			"cert_hash": schema.StringAttribute{
				Optional:    true,
				Description: "Hash of the certificate file content. Used internally to detect certificate file changes.",
			},
			"key_hash": schema.StringAttribute{
				Optional:    true,
				Description: "Hash of the private key file content. Used internally to detect key file changes.",
			},
		},
	}
}

func sslcertkeyGetThePayloadFromtheConfig(ctx context.Context, data *SslCertKeyResourceModel) ssl.Sslcertkey {
	tflog.Debug(ctx, "In sslcertkeyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslcertkey := ssl.Sslcertkey{}

	if !data.Certkey.IsNull() {
		sslcertkey.Certkey = data.Certkey.ValueString()
	}
	if !data.Cert.IsNull() {
		sslcertkey.Cert = data.Cert.ValueString()
	}
	if !data.Key.IsNull() {
		sslcertkey.Key = data.Key.ValueString()
	}
	if !data.Password.IsNull() {
		sslcertkey.Password = data.Password.ValueBool()
	}
	if !data.Fipskey.IsNull() {
		sslcertkey.Fipskey = data.Fipskey.ValueString()
	}
	if !data.Hsmkey.IsNull() {
		sslcertkey.Hsmkey = data.Hsmkey.ValueString()
	}
	if !data.Inform.IsNull() {
		sslcertkey.Inform = data.Inform.ValueString()
	}
	if !data.Passplain.IsNull() {
		sslcertkey.Passplain = data.Passplain.ValueString()
	}
	if !data.Expirymonitor.IsNull() {
		sslcertkey.Expirymonitor = data.Expirymonitor.ValueString()
	}
	if !data.NotificationPeriod.IsNull() {
		sslcertkey.Notificationperiod = utils.IntPtr(int(data.NotificationPeriod.ValueInt64()))
	}
	if !data.Bundle.IsNull() {
		sslcertkey.Bundle = data.Bundle.ValueString()
	}
	if !data.DeleteCertKeyFilesOnRemoval.IsNull() {
		sslcertkey.Deletecertkeyfilesonremoval = data.DeleteCertKeyFilesOnRemoval.ValueString()
	}
	if !data.PassplainWo.IsNull() {
		passplainWo := data.PassplainWo.ValueString()
		if passplainWo != "" {
			sslcertkey.Passplain = passplainWo
		}
	}

	return sslcertkey
}

func sslcertkeySetAttrFromGet(ctx context.Context, data *SslCertKeyResourceModel, getResponseData map[string]interface{}) *SslCertKeyResourceModel {
	tflog.Debug(ctx, "In sslcertkeySetAttrFromGet Function")

	// Set ID for the resource - use certkey as the identifier
	if val, ok := getResponseData["certkey"]; ok && val != nil {
		data.Id = types.StringValue(val.(string))
		data.Certkey = types.StringValue(val.(string))
	}

	// Convert API response to model
	if val, ok := getResponseData["cert"]; ok && val != nil {
		data.Cert = types.StringValue(val.(string))
	} else {
		data.Cert = types.StringNull()
	}
	if val, ok := getResponseData["key"]; ok && val != nil {
		data.Key = types.StringValue(val.(string))
	} else {
		data.Key = types.StringNull()
	}
	// Password and passplain are not returned by NITRO API - keep existing state
	// The API returns hashed values which would cause drift, so we skip updating these fields
	// They will retain their configured values from the plan
	if val, ok := getResponseData["fipskey"]; ok && val != nil {
		data.Fipskey = types.StringValue(val.(string))
	} else {
		data.Fipskey = types.StringNull()
	}
	if val, ok := getResponseData["hsmkey"]; ok && val != nil {
		data.Hsmkey = types.StringValue(val.(string))
	} else {
		data.Hsmkey = types.StringNull()
	}
	if val, ok := getResponseData["inform"]; ok && val != nil {
		data.Inform = types.StringValue(val.(string))
	} else {
		data.Inform = types.StringNull()
	}
	if val, ok := getResponseData["expirymonitor"]; ok && val != nil {
		data.Expirymonitor = types.StringValue(val.(string))
	} else {
		data.Expirymonitor = types.StringNull()
	}
	if val, ok := getResponseData["notificationperiod"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.NotificationPeriod = types.Int64Value(intVal)
		}
	} else {
		data.NotificationPeriod = types.Int64Null()
	}
	if val, ok := getResponseData["bundle"]; ok && val != nil {
		data.Bundle = types.StringValue(val.(string))
	} else {
		data.Bundle = types.StringNull()
	}
	if val, ok := getResponseData["linkcertkeyname"]; ok && val != nil {
		data.LinkCertKeyName = types.StringValue(val.(string))
	} else {
		data.LinkCertKeyName = types.StringNull()
	}
	// if val, ok := getResponseData["nodomaincheck"]; ok && val != nil {
	// 	data.NoDomainCheck = types.BoolValue(val.(bool))
	// } else {
	// 	data.NoDomainCheck = types.BoolNull()
	// }
	if val, ok := getResponseData["ocspstaplingcache"]; ok && val != nil {
		data.OcspStaplingCache = types.BoolValue(val.(bool))
	} else {
		data.OcspStaplingCache = types.BoolNull()
	}
	if val, ok := getResponseData["deletecertkeyfilesonremoval"]; ok && val != nil {
		data.DeleteCertKeyFilesOnRemoval = types.StringValue(val.(string))
	} else {
		data.DeleteCertKeyFilesOnRemoval = types.StringNull()
	}
	if val, ok := getResponseData["deletefromdevice"]; ok && val != nil {
		data.DeleteFromDevice = types.BoolValue(val.(bool))
	} else {
		data.DeleteFromDevice = types.BoolNull()
	}

	return data
}
