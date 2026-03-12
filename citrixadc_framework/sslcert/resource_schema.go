package sslcert

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

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslcertResourceModel describes the resource data model.
type SslcertResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Cacert         types.String `tfsdk:"cacert"`
	Cacertform     types.String `tfsdk:"cacertform"`
	Cakey          types.String `tfsdk:"cakey"`
	Cakeyform      types.String `tfsdk:"cakeyform"`
	Caserial       types.String `tfsdk:"caserial"`
	Certfile       types.String `tfsdk:"certfile"`
	Certform       types.String `tfsdk:"certform"`
	Certtype       types.String `tfsdk:"certtype"`
	Days           types.Int64  `tfsdk:"days"`
	Keyfile        types.String `tfsdk:"keyfile"`
	Keyform        types.String `tfsdk:"keyform"`
	Pempassphrase  types.String `tfsdk:"pempassphrase"`
	Reqfile        types.String `tfsdk:"reqfile"`
	Subjectaltname types.String `tfsdk:"subjectaltname"`
}

func (r *SslcertResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslcert resource.",
			},
			"cacert": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the CA certificate file that issues and signs the Intermediate-CA certificate or the end-user client and server certificates.",
			},
			"cacertform": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("PEM"),
				Description: "Format of the CA certificate.",
			},
			"cakey": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Private key, associated with the CA certificate that is used to sign the Intermediate-CA certificate or the end-user client and server certificate. If the CA key file is password protected, the user is prompted to enter the pass phrase that was used to encrypt the key.",
			},
			"cakeyform": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("PEM"),
				Description: "Format for the CA certificate.",
			},
			"caserial": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Serial number file maintained for the CA certificate. This file contains the serial number of the next certificate to be issued or signed by the CA. If the specified file does not exist, a new file is created, with /nsconfig/ssl/ as the default path. If you do not specify a proper path for the existing serial file, a new serial file is created. This might change the certificate serial numbers assigned by the CA certificate to each of the certificates it signs.",
			},
			"certfile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for and, optionally, path to the generated certificate file. /nsconfig/ssl/ is the default path.",
			},
			"certform": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("PEM"),
				Description: "Format in which the certificate is stored on the appliance.",
			},
			"certtype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of certificate to generate. Specify one of the following:\n* ROOT_CERT - Self-signed Root-CA certificate. You must specify the key file name. The generated Root-CA certificate can be used for signing end-user client or server certificates or to create Intermediate-CA certificates.\n* INTM_CERT - Intermediate-CA certificate.\n* CLNT_CERT - End-user client certificate used for client authentication.\n* SRVR_CERT - SSL server certificate used on SSL servers for end-to-end encryption.",
			},
			"days": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(365),
				Description: "Number of days for which the certificate will be valid, beginning with the time and day (system time) of creation.",
			},
			"keyfile": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for and, optionally, path to the private key. You can either use an existing RSA or DSA key that you own or create a new private key on the Citrix ADC. This file is required only when creating a self-signed Root-CA certificate. The key file is stored in the /nsconfig/ssl directory by default.\nIf the input key specified is an encrypted key, you are prompted to enter the PEM pass phrase that was used for encrypting the key.",
			},
			"keyform": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("PEM"),
				Description: "Format in which the key is stored on the appliance.",
			},
			"pempassphrase": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
			"reqfile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for and, optionally, path to the certificate-signing request (CSR). /nsconfig/ssl/ is the default path.",
			},
			"subjectaltname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Subject Alternative Name (SAN) is an extension to X.509 that allows various values to be associated with a security certificate using a subjectAltName field. These values are called \"Subject Alternative Names\" (SAN). Names include:\n      1. Email addresses\n      2. IP addresses\n      3. URIs\n      4. DNS names (This is usually also provided as the Common Name RDN within the Subject field of the main certificate.)\n      5. directory names (alternative Distinguished Names to that given in the Subject)",
			},
		},
	}
}

func sslcertGetThePayloadFromtheConfig(ctx context.Context, data *SslcertResourceModel) ssl.Sslcert {
	tflog.Debug(ctx, "In sslcertGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslcert := ssl.Sslcert{}
	if !data.Cacert.IsNull() {
		sslcert.Cacert = data.Cacert.ValueString()
	}
	if !data.Cacertform.IsNull() {
		sslcert.Cacertform = data.Cacertform.ValueString()
	}
	if !data.Cakey.IsNull() {
		sslcert.Cakey = data.Cakey.ValueString()
	}
	if !data.Cakeyform.IsNull() {
		sslcert.Cakeyform = data.Cakeyform.ValueString()
	}
	if !data.Caserial.IsNull() {
		sslcert.Caserial = data.Caserial.ValueString()
	}
	if !data.Certfile.IsNull() {
		sslcert.Certfile = data.Certfile.ValueString()
	}
	if !data.Certform.IsNull() {
		sslcert.Certform = data.Certform.ValueString()
	}
	if !data.Certtype.IsNull() {
		sslcert.Certtype = data.Certtype.ValueString()
	}
	if !data.Days.IsNull() {
		sslcert.Days = utils.IntPtr(int(data.Days.ValueInt64()))
	}
	if !data.Keyfile.IsNull() {
		sslcert.Keyfile = data.Keyfile.ValueString()
	}
	if !data.Keyform.IsNull() {
		sslcert.Keyform = data.Keyform.ValueString()
	}
	if !data.Pempassphrase.IsNull() {
		sslcert.Pempassphrase = data.Pempassphrase.ValueString()
	}
	if !data.Reqfile.IsNull() {
		sslcert.Reqfile = data.Reqfile.ValueString()
	}
	if !data.Subjectaltname.IsNull() {
		sslcert.Subjectaltname = data.Subjectaltname.ValueString()
	}

	return sslcert
}

func sslcertSetAttrFromGet(ctx context.Context, data *SslcertResourceModel, getResponseData map[string]interface{}) *SslcertResourceModel {
	tflog.Debug(ctx, "In sslcertSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["cacert"]; ok && val != nil {
		data.Cacert = types.StringValue(val.(string))
	} else {
		data.Cacert = types.StringNull()
	}
	if val, ok := getResponseData["cacertform"]; ok && val != nil {
		data.Cacertform = types.StringValue(val.(string))
	} else {
		data.Cacertform = types.StringNull()
	}
	if val, ok := getResponseData["cakey"]; ok && val != nil {
		data.Cakey = types.StringValue(val.(string))
	} else {
		data.Cakey = types.StringNull()
	}
	if val, ok := getResponseData["cakeyform"]; ok && val != nil {
		data.Cakeyform = types.StringValue(val.(string))
	} else {
		data.Cakeyform = types.StringNull()
	}
	if val, ok := getResponseData["caserial"]; ok && val != nil {
		data.Caserial = types.StringValue(val.(string))
	} else {
		data.Caserial = types.StringNull()
	}
	if val, ok := getResponseData["certfile"]; ok && val != nil {
		data.Certfile = types.StringValue(val.(string))
	} else {
		data.Certfile = types.StringNull()
	}
	if val, ok := getResponseData["certform"]; ok && val != nil {
		data.Certform = types.StringValue(val.(string))
	} else {
		data.Certform = types.StringNull()
	}
	if val, ok := getResponseData["certtype"]; ok && val != nil {
		data.Certtype = types.StringValue(val.(string))
	} else {
		data.Certtype = types.StringNull()
	}
	if val, ok := getResponseData["days"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Days = types.Int64Value(intVal)
		}
	} else {
		data.Days = types.Int64Null()
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
	if val, ok := getResponseData["pempassphrase"]; ok && val != nil {
		data.Pempassphrase = types.StringValue(val.(string))
	} else {
		data.Pempassphrase = types.StringNull()
	}
	if val, ok := getResponseData["reqfile"]; ok && val != nil {
		data.Reqfile = types.StringValue(val.(string))
	} else {
		data.Reqfile = types.StringNull()
	}
	if val, ok := getResponseData["subjectaltname"]; ok && val != nil {
		data.Subjectaltname = types.StringValue(val.(string))
	} else {
		data.Subjectaltname = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("sslcert-config")

	return data
}
