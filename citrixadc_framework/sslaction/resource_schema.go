package sslaction

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

// SslactionResourceModel describes the resource data model.
type SslactionResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Cacertgrpname          types.String `tfsdk:"cacertgrpname"`
	Certfingerprintdigest  types.String `tfsdk:"certfingerprintdigest"`
	Certfingerprintheader  types.String `tfsdk:"certfingerprintheader"`
	Certhashheader         types.String `tfsdk:"certhashheader"`
	Certheader             types.String `tfsdk:"certheader"`
	Certissuerheader       types.String `tfsdk:"certissuerheader"`
	Certnotafterheader     types.String `tfsdk:"certnotafterheader"`
	Certnotbeforeheader    types.String `tfsdk:"certnotbeforeheader"`
	Certserialheader       types.String `tfsdk:"certserialheader"`
	Certsubjectheader      types.String `tfsdk:"certsubjectheader"`
	Cipher                 types.String `tfsdk:"cipher"`
	Cipherheader           types.String `tfsdk:"cipherheader"`
	Clientauth             types.String `tfsdk:"clientauth"`
	Clientcert             types.String `tfsdk:"clientcert"`
	Clientcertfingerprint  types.String `tfsdk:"clientcertfingerprint"`
	Clientcerthash         types.String `tfsdk:"clientcerthash"`
	Clientcertissuer       types.String `tfsdk:"clientcertissuer"`
	Clientcertnotafter     types.String `tfsdk:"clientcertnotafter"`
	Clientcertnotbefore    types.String `tfsdk:"clientcertnotbefore"`
	Clientcertserialnumber types.String `tfsdk:"clientcertserialnumber"`
	Clientcertsubject      types.String `tfsdk:"clientcertsubject"`
	Clientcertverification types.String `tfsdk:"clientcertverification"`
	Forward                types.String `tfsdk:"forward"`
	Name                   types.String `tfsdk:"name"`
	Owasupport             types.String `tfsdk:"owasupport"`
	Sessionid              types.String `tfsdk:"sessionid"`
	Sessionidheader        types.String `tfsdk:"sessionidheader"`
	Ssllogprofile          types.String `tfsdk:"ssllogprofile"`
}

func (r *SslactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslaction resource.",
			},
			"cacertgrpname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "This action will allow to pick CA(s) from the specific CA group, to verify the client certificate.",
			},
			"certfingerprintdigest": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Digest algorithm used to compute the fingerprint of the client certificate.",
			},
			"certfingerprintheader": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the header into which to insert the client certificate fingerprint.",
			},
			"certhashheader": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the header into which to insert the client certificate signature (hash).",
			},
			"certheader": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the header into which to insert the client certificate.",
			},
			"certissuerheader": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the header into which to insert the client certificate issuer details.",
			},
			"certnotafterheader": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the header into which to insert the certificate's expiry date.",
			},
			"certnotbeforeheader": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the header into which to insert the date and time from which the certificate is valid.",
			},
			"certserialheader": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the header into which to insert the client serial number.",
			},
			"certsubjectheader": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the header into which to insert the client certificate subject.",
			},
			"cipher": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Insert the cipher suite that the client and the Citrix ADC negotiated for the SSL session into the HTTP header of the request being sent to the web server. The appliance inserts the cipher-suite name, SSL protocol, export or non-export string, and cipher strength bit, depending on the type of browser connecting to the SSL virtual server or service (for example, Cipher-Suite: RC4- MD5 SSLv3 Non-Export 128-bit).",
			},
			"cipherheader": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the header into which to insert the name of the cipher suite.",
			},
			"clientauth": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Perform client certificate authentication.",
			},
			"clientcert": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Insert the entire client certificate into the HTTP header of the request being sent to the web server. The certificate is inserted in ASCII (PEM) format.",
			},
			"clientcertfingerprint": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Insert the certificate's fingerprint into the HTTP header of the request being sent to the web server. The fingerprint is derived by computing the specified hash value (SHA256, for example) of the DER-encoding of the client certificate.",
			},
			"clientcerthash": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Insert the certificate's signature into the HTTP header of the request being sent to the web server. The signature is the value extracted directly from the X.509 certificate signature field. All X.509 certificates contain a signature field.",
			},
			"clientcertissuer": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Insert the certificate issuer details into the HTTP header of the request being sent to the web server.",
			},
			"clientcertnotafter": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Insert the date of expiry of the certificate into the HTTP header of the request being sent to the web server. Every certificate is configured with the date and time at which the certificate expires.",
			},
			"clientcertnotbefore": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Insert the date from which the certificate is valid into the HTTP header of the request being sent to the web server. Every certificate is configured with the date and time from which it is valid.",
			},
			"clientcertserialnumber": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Insert the entire client serial number into the HTTP header of the request being sent to the web server.",
			},
			"clientcertsubject": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Insert the client certificate subject, also known as the distinguished name (DN), into the HTTP header of the request being sent to the web server.",
			},
			"clientcertverification": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("Mandatory"),
				Description: "Client certificate verification is mandatory or optional.",
			},
			"forward": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "This action takes an argument a vserver name, to this vserver one will be able to forward all the packets.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the SSL action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"owasupport": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "If the appliance is in front of an Outlook Web Access (OWA) server, insert a special header field, FRONT-END-HTTPS: ON, into the HTTP requests going to the OWA server. This header communicates to the server that the transaction is HTTPS and not HTTP.",
			},
			"sessionid": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Insert the SSL session ID into the HTTP header of the request being sent to the web server. Every SSL connection that the client and the Citrix ADC share has a unique ID that identifies the specific connection.",
			},
			"sessionidheader": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the header into which to insert the Session ID.",
			},
			"ssllogprofile": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the ssllogprofile.",
			},
		},
	}
}

func sslactionGetThePayloadFromtheConfig(ctx context.Context, data *SslactionResourceModel) ssl.Sslaction {
	tflog.Debug(ctx, "In sslactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslaction := ssl.Sslaction{}
	if !data.Cacertgrpname.IsNull() {
		sslaction.Cacertgrpname = data.Cacertgrpname.ValueString()
	}
	if !data.Certfingerprintdigest.IsNull() {
		sslaction.Certfingerprintdigest = data.Certfingerprintdigest.ValueString()
	}
	if !data.Certfingerprintheader.IsNull() {
		sslaction.Certfingerprintheader = data.Certfingerprintheader.ValueString()
	}
	if !data.Certhashheader.IsNull() {
		sslaction.Certhashheader = data.Certhashheader.ValueString()
	}
	if !data.Certheader.IsNull() {
		sslaction.Certheader = data.Certheader.ValueString()
	}
	if !data.Certissuerheader.IsNull() {
		sslaction.Certissuerheader = data.Certissuerheader.ValueString()
	}
	if !data.Certnotafterheader.IsNull() {
		sslaction.Certnotafterheader = data.Certnotafterheader.ValueString()
	}
	if !data.Certnotbeforeheader.IsNull() {
		sslaction.Certnotbeforeheader = data.Certnotbeforeheader.ValueString()
	}
	if !data.Certserialheader.IsNull() {
		sslaction.Certserialheader = data.Certserialheader.ValueString()
	}
	if !data.Certsubjectheader.IsNull() {
		sslaction.Certsubjectheader = data.Certsubjectheader.ValueString()
	}
	if !data.Cipher.IsNull() {
		sslaction.Cipher = data.Cipher.ValueString()
	}
	if !data.Cipherheader.IsNull() {
		sslaction.Cipherheader = data.Cipherheader.ValueString()
	}
	if !data.Clientauth.IsNull() {
		sslaction.Clientauth = data.Clientauth.ValueString()
	}
	if !data.Clientcert.IsNull() {
		sslaction.Clientcert = data.Clientcert.ValueString()
	}
	if !data.Clientcertfingerprint.IsNull() {
		sslaction.Clientcertfingerprint = data.Clientcertfingerprint.ValueString()
	}
	if !data.Clientcerthash.IsNull() {
		sslaction.Clientcerthash = data.Clientcerthash.ValueString()
	}
	if !data.Clientcertissuer.IsNull() {
		sslaction.Clientcertissuer = data.Clientcertissuer.ValueString()
	}
	if !data.Clientcertnotafter.IsNull() {
		sslaction.Clientcertnotafter = data.Clientcertnotafter.ValueString()
	}
	if !data.Clientcertnotbefore.IsNull() {
		sslaction.Clientcertnotbefore = data.Clientcertnotbefore.ValueString()
	}
	if !data.Clientcertserialnumber.IsNull() {
		sslaction.Clientcertserialnumber = data.Clientcertserialnumber.ValueString()
	}
	if !data.Clientcertsubject.IsNull() {
		sslaction.Clientcertsubject = data.Clientcertsubject.ValueString()
	}
	if !data.Clientcertverification.IsNull() {
		sslaction.Clientcertverification = data.Clientcertverification.ValueString()
	}
	if !data.Forward.IsNull() {
		sslaction.Forward = data.Forward.ValueString()
	}
	if !data.Name.IsNull() {
		sslaction.Name = data.Name.ValueString()
	}
	if !data.Owasupport.IsNull() {
		sslaction.Owasupport = data.Owasupport.ValueString()
	}
	if !data.Sessionid.IsNull() {
		sslaction.Sessionid = data.Sessionid.ValueString()
	}
	if !data.Sessionidheader.IsNull() {
		sslaction.Sessionidheader = data.Sessionidheader.ValueString()
	}
	if !data.Ssllogprofile.IsNull() {
		sslaction.Ssllogprofile = data.Ssllogprofile.ValueString()
	}

	return sslaction
}

func sslactionSetAttrFromGet(ctx context.Context, data *SslactionResourceModel, getResponseData map[string]interface{}) *SslactionResourceModel {
	tflog.Debug(ctx, "In sslactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["cacertgrpname"]; ok && val != nil {
		data.Cacertgrpname = types.StringValue(val.(string))
	} else {
		data.Cacertgrpname = types.StringNull()
	}
	if val, ok := getResponseData["certfingerprintdigest"]; ok && val != nil {
		data.Certfingerprintdigest = types.StringValue(val.(string))
	} else {
		data.Certfingerprintdigest = types.StringNull()
	}
	if val, ok := getResponseData["certfingerprintheader"]; ok && val != nil {
		data.Certfingerprintheader = types.StringValue(val.(string))
	} else {
		data.Certfingerprintheader = types.StringNull()
	}
	if val, ok := getResponseData["certhashheader"]; ok && val != nil {
		data.Certhashheader = types.StringValue(val.(string))
	} else {
		data.Certhashheader = types.StringNull()
	}
	if val, ok := getResponseData["certheader"]; ok && val != nil {
		data.Certheader = types.StringValue(val.(string))
	} else {
		data.Certheader = types.StringNull()
	}
	if val, ok := getResponseData["certissuerheader"]; ok && val != nil {
		data.Certissuerheader = types.StringValue(val.(string))
	} else {
		data.Certissuerheader = types.StringNull()
	}
	if val, ok := getResponseData["certnotafterheader"]; ok && val != nil {
		data.Certnotafterheader = types.StringValue(val.(string))
	} else {
		data.Certnotafterheader = types.StringNull()
	}
	if val, ok := getResponseData["certnotbeforeheader"]; ok && val != nil {
		data.Certnotbeforeheader = types.StringValue(val.(string))
	} else {
		data.Certnotbeforeheader = types.StringNull()
	}
	if val, ok := getResponseData["certserialheader"]; ok && val != nil {
		data.Certserialheader = types.StringValue(val.(string))
	} else {
		data.Certserialheader = types.StringNull()
	}
	if val, ok := getResponseData["certsubjectheader"]; ok && val != nil {
		data.Certsubjectheader = types.StringValue(val.(string))
	} else {
		data.Certsubjectheader = types.StringNull()
	}
	if val, ok := getResponseData["cipher"]; ok && val != nil {
		data.Cipher = types.StringValue(val.(string))
	} else {
		data.Cipher = types.StringNull()
	}
	if val, ok := getResponseData["cipherheader"]; ok && val != nil {
		data.Cipherheader = types.StringValue(val.(string))
	} else {
		data.Cipherheader = types.StringNull()
	}
	if val, ok := getResponseData["clientauth"]; ok && val != nil {
		data.Clientauth = types.StringValue(val.(string))
	} else {
		data.Clientauth = types.StringNull()
	}
	if val, ok := getResponseData["clientcert"]; ok && val != nil {
		data.Clientcert = types.StringValue(val.(string))
	} else {
		data.Clientcert = types.StringNull()
	}
	if val, ok := getResponseData["clientcertfingerprint"]; ok && val != nil {
		data.Clientcertfingerprint = types.StringValue(val.(string))
	} else {
		data.Clientcertfingerprint = types.StringNull()
	}
	if val, ok := getResponseData["clientcerthash"]; ok && val != nil {
		data.Clientcerthash = types.StringValue(val.(string))
	} else {
		data.Clientcerthash = types.StringNull()
	}
	if val, ok := getResponseData["clientcertissuer"]; ok && val != nil {
		data.Clientcertissuer = types.StringValue(val.(string))
	} else {
		data.Clientcertissuer = types.StringNull()
	}
	if val, ok := getResponseData["clientcertnotafter"]; ok && val != nil {
		data.Clientcertnotafter = types.StringValue(val.(string))
	} else {
		data.Clientcertnotafter = types.StringNull()
	}
	if val, ok := getResponseData["clientcertnotbefore"]; ok && val != nil {
		data.Clientcertnotbefore = types.StringValue(val.(string))
	} else {
		data.Clientcertnotbefore = types.StringNull()
	}
	if val, ok := getResponseData["clientcertserialnumber"]; ok && val != nil {
		data.Clientcertserialnumber = types.StringValue(val.(string))
	} else {
		data.Clientcertserialnumber = types.StringNull()
	}
	if val, ok := getResponseData["clientcertsubject"]; ok && val != nil {
		data.Clientcertsubject = types.StringValue(val.(string))
	} else {
		data.Clientcertsubject = types.StringNull()
	}
	if val, ok := getResponseData["clientcertverification"]; ok && val != nil {
		data.Clientcertverification = types.StringValue(val.(string))
	} else {
		data.Clientcertverification = types.StringNull()
	}
	if val, ok := getResponseData["forward"]; ok && val != nil {
		data.Forward = types.StringValue(val.(string))
	} else {
		data.Forward = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["owasupport"]; ok && val != nil {
		data.Owasupport = types.StringValue(val.(string))
	} else {
		data.Owasupport = types.StringNull()
	}
	if val, ok := getResponseData["sessionid"]; ok && val != nil {
		data.Sessionid = types.StringValue(val.(string))
	} else {
		data.Sessionid = types.StringNull()
	}
	if val, ok := getResponseData["sessionidheader"]; ok && val != nil {
		data.Sessionidheader = types.StringValue(val.(string))
	} else {
		data.Sessionidheader = types.StringNull()
	}
	if val, ok := getResponseData["ssllogprofile"]; ok && val != nil {
		data.Ssllogprofile = types.StringValue(val.(string))
	} else {
		data.Ssllogprofile = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
