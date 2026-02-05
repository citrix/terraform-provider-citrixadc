package sslcertreq

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslcertreqResourceModel describes the resource data model.
type SslcertreqResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Challengepassword    types.String `tfsdk:"challengepassword"`
	Commonname           types.String `tfsdk:"commonname"`
	Companyname          types.String `tfsdk:"companyname"`
	Countryname          types.String `tfsdk:"countryname"`
	Digestmethod         types.String `tfsdk:"digestmethod"`
	Emailaddress         types.String `tfsdk:"emailaddress"`
	Fipskeyname          types.String `tfsdk:"fipskeyname"`
	Keyfile              types.String `tfsdk:"keyfile"`
	Keyform              types.String `tfsdk:"keyform"`
	Localityname         types.String `tfsdk:"localityname"`
	Organizationname     types.String `tfsdk:"organizationname"`
	Organizationunitname types.String `tfsdk:"organizationunitname"`
	Pempassphrase        types.String `tfsdk:"pempassphrase"`
	Reqfile              types.String `tfsdk:"reqfile"`
	Statename            types.String `tfsdk:"statename"`
	Subjectaltname       types.String `tfsdk:"subjectaltname"`
}

func (r *SslcertreqResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslcertreq resource.",
			},
			"challengepassword": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Pass phrase, embedded in the certificate signing request that is shared only between the client or server requesting the certificate and the SSL certificate issuer (typically the certificate authority). This pass phrase can be used to authenticate a client or server that is requesting a certificate from the certificate authority.",
			},
			"commonname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Fully qualified domain name for the company or web site. The common name must match the name used by DNS servers to do a DNS lookup of your server. Most browsers use this information for authenticating the server's certificate during the SSL handshake. If the server name in the URL does not match the common name as given in the server certificate, the browser terminates the SSL handshake or prompts the user with a warning message.\nDo not use wildcard characters, such as asterisk (*) or question mark (?), and do not use an IP address as the common name. The common name must not contain the protocol specifier <http://> or <https://>.",
			},
			"companyname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Additional name for the company or web site.",
			},
			"countryname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Two letter ISO code for your country. For example, US for United States.",
			},
			"digestmethod": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Digest algorithm used in creating CSR",
			},
			"emailaddress": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Contact person's e-mail address. This address is publically displayed as part of the certificate. Provide an e-mail address that is monitored by an administrator who can be contacted about the certificate.",
			},
			"fipskeyname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the FIPS key used to create the certificate signing request. FIPS keys are created inside the Hardware Security Module of the FIPS card.",
			},
			"keyfile": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the private key used to create the certificate signing request, which then becomes part of the certificate-key pair. The private key can be either an RSA or a DSA key. The key must be present in the appliance's local storage. /nsconfig/ssl is the default path.",
			},
			"keyform": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("PEM"),
				Description: "Format in which the key is stored on the appliance.",
			},
			"localityname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the city or town in which your organization's head office is located.",
			},
			"organizationname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the organization that will use this certificate. The organization name (corporation, limited partnership, university, or government agency) must be registered with some authority at the national, state, or city level. Use the legal name under which the organization is registered.\nDo not abbreviate the organization name and do not use the following characters in the name:\nAngle brackets (< >) tilde (~), exclamation mark, at (@), pound (#), zero (0), caret (^), asterisk (*), forward slash (/), square brackets ([ ]), question mark (?).",
			},
			"organizationunitname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the division or section in the organization that will use the certificate.",
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
				Description: "Name for and, optionally, path to the certificate signing request (CSR). /nsconfig/ssl/ is the default path.",
			},
			"statename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Full name of the state or province where your organization is located.\nDo not abbreviate.",
			},
			"subjectaltname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Subject Alternative Name (SAN) is an extension to X.509 that allows various values to be associated with a security certificate using a subjectAltName field. These values are called \"Subject Alternative Names\" (SAN). Names include:\n      1. Email addresses\n      2. IP addresses\n      3. URIs\n      4. DNS names (this is usually also provided as the Common Name RDN within the Subject field of the main certificate.)\n      5. Directory names (alternative Distinguished Names to that given in the Subject)\nExample:-subjectAltName \"DNS:*.example.com DNS:www.example.org\"",
			},
		},
	}
}

func sslcertreqGetThePayloadFromtheConfig(ctx context.Context, data *SslcertreqResourceModel) ssl.Sslcertreq {
	tflog.Debug(ctx, "In sslcertreqGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslcertreq := ssl.Sslcertreq{}
	if !data.Challengepassword.IsNull() {
		sslcertreq.Challengepassword = data.Challengepassword.ValueString()
	}
	if !data.Commonname.IsNull() {
		sslcertreq.Commonname = data.Commonname.ValueString()
	}
	if !data.Companyname.IsNull() {
		sslcertreq.Companyname = data.Companyname.ValueString()
	}
	if !data.Countryname.IsNull() {
		sslcertreq.Countryname = data.Countryname.ValueString()
	}
	if !data.Digestmethod.IsNull() {
		sslcertreq.Digestmethod = data.Digestmethod.ValueString()
	}
	if !data.Emailaddress.IsNull() {
		sslcertreq.Emailaddress = data.Emailaddress.ValueString()
	}
	if !data.Fipskeyname.IsNull() {
		sslcertreq.Fipskeyname = data.Fipskeyname.ValueString()
	}
	if !data.Keyfile.IsNull() {
		sslcertreq.Keyfile = data.Keyfile.ValueString()
	}
	if !data.Keyform.IsNull() {
		sslcertreq.Keyform = data.Keyform.ValueString()
	}
	if !data.Localityname.IsNull() {
		sslcertreq.Localityname = data.Localityname.ValueString()
	}
	if !data.Organizationname.IsNull() {
		sslcertreq.Organizationname = data.Organizationname.ValueString()
	}
	if !data.Organizationunitname.IsNull() {
		sslcertreq.Organizationunitname = data.Organizationunitname.ValueString()
	}
	if !data.Pempassphrase.IsNull() {
		sslcertreq.Pempassphrase = data.Pempassphrase.ValueString()
	}
	if !data.Reqfile.IsNull() {
		sslcertreq.Reqfile = data.Reqfile.ValueString()
	}
	if !data.Statename.IsNull() {
		sslcertreq.Statename = data.Statename.ValueString()
	}
	if !data.Subjectaltname.IsNull() {
		sslcertreq.Subjectaltname = data.Subjectaltname.ValueString()
	}

	return sslcertreq
}

func sslcertreqSetAttrFromGet(ctx context.Context, data *SslcertreqResourceModel, getResponseData map[string]interface{}) *SslcertreqResourceModel {
	tflog.Debug(ctx, "In sslcertreqSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["challengepassword"]; ok && val != nil {
		data.Challengepassword = types.StringValue(val.(string))
	} else {
		data.Challengepassword = types.StringNull()
	}
	if val, ok := getResponseData["commonname"]; ok && val != nil {
		data.Commonname = types.StringValue(val.(string))
	} else {
		data.Commonname = types.StringNull()
	}
	if val, ok := getResponseData["companyname"]; ok && val != nil {
		data.Companyname = types.StringValue(val.(string))
	} else {
		data.Companyname = types.StringNull()
	}
	if val, ok := getResponseData["countryname"]; ok && val != nil {
		data.Countryname = types.StringValue(val.(string))
	} else {
		data.Countryname = types.StringNull()
	}
	if val, ok := getResponseData["digestmethod"]; ok && val != nil {
		data.Digestmethod = types.StringValue(val.(string))
	} else {
		data.Digestmethod = types.StringNull()
	}
	if val, ok := getResponseData["emailaddress"]; ok && val != nil {
		data.Emailaddress = types.StringValue(val.(string))
	} else {
		data.Emailaddress = types.StringNull()
	}
	if val, ok := getResponseData["fipskeyname"]; ok && val != nil {
		data.Fipskeyname = types.StringValue(val.(string))
	} else {
		data.Fipskeyname = types.StringNull()
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
	if val, ok := getResponseData["localityname"]; ok && val != nil {
		data.Localityname = types.StringValue(val.(string))
	} else {
		data.Localityname = types.StringNull()
	}
	if val, ok := getResponseData["organizationname"]; ok && val != nil {
		data.Organizationname = types.StringValue(val.(string))
	} else {
		data.Organizationname = types.StringNull()
	}
	if val, ok := getResponseData["organizationunitname"]; ok && val != nil {
		data.Organizationunitname = types.StringValue(val.(string))
	} else {
		data.Organizationunitname = types.StringNull()
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
	if val, ok := getResponseData["statename"]; ok && val != nil {
		data.Statename = types.StringValue(val.(string))
	} else {
		data.Statename = types.StringNull()
	}
	if val, ok := getResponseData["subjectaltname"]; ok && val != nil {
		data.Subjectaltname = types.StringValue(val.(string))
	} else {
		data.Subjectaltname = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("sslcertreq-config")

	return data
}
