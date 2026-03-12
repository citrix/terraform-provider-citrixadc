package sslservice_sslcertkey_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslserviceSslcertkeyBindingResourceModel describes the resource data model.
type SslserviceSslcertkeyBindingResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Ca          types.Bool   `tfsdk:"ca"`
	Certkeyname types.String `tfsdk:"certkeyname"`
	Crlcheck    types.String `tfsdk:"crlcheck"`
	Ocspcheck   types.String `tfsdk:"ocspcheck"`
	Servicename types.String `tfsdk:"servicename"`
	Skipcaname  types.Bool   `tfsdk:"skipcaname"`
	Snicert     types.Bool   `tfsdk:"snicert"`
}

func (r *SslserviceSslcertkeyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslservice_sslcertkey_binding resource.",
			},
			"ca": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CA certificate.",
			},
			"certkeyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The certificate key pair binding.",
			},
			"crlcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The state of the CRL check parameter. (Mandatory/Optional)",
			},
			"ocspcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Rule to use for the OCSP responder associated with the CA certificate during client authentication. If MANDATORY is specified, deny all SSL clients if the OCSP check fails because of connectivity issues with the remote OCSP server, or any other reason that prevents the OCSP check. With the OPTIONAL setting, allow SSL clients even if the OCSP check fails except when the client certificate is revoked.",
			},
			"servicename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL service for which to set advanced configuration.",
			},
			"skipcaname": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting      for client certificate in a SSL handshake",
			},
			"snicert": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.",
			},
		},
	}
}

func sslservice_sslcertkey_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SslserviceSslcertkeyBindingResourceModel) ssl.Sslservicesslcertkeybinding {
	tflog.Debug(ctx, "In sslservice_sslcertkey_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslservice_sslcertkey_binding := ssl.Sslservicesslcertkeybinding{}
	if !data.Ca.IsNull() {
		sslservice_sslcertkey_binding.Ca = data.Ca.ValueBool()
	}
	if !data.Certkeyname.IsNull() {
		sslservice_sslcertkey_binding.Certkeyname = data.Certkeyname.ValueString()
	}
	if !data.Crlcheck.IsNull() {
		sslservice_sslcertkey_binding.Crlcheck = data.Crlcheck.ValueString()
	}
	if !data.Ocspcheck.IsNull() {
		sslservice_sslcertkey_binding.Ocspcheck = data.Ocspcheck.ValueString()
	}
	if !data.Servicename.IsNull() {
		sslservice_sslcertkey_binding.Servicename = data.Servicename.ValueString()
	}
	if !data.Skipcaname.IsNull() {
		sslservice_sslcertkey_binding.Skipcaname = data.Skipcaname.ValueBool()
	}
	if !data.Snicert.IsNull() {
		sslservice_sslcertkey_binding.Snicert = data.Snicert.ValueBool()
	}

	return sslservice_sslcertkey_binding
}

func sslservice_sslcertkey_bindingSetAttrFromGet(ctx context.Context, data *SslserviceSslcertkeyBindingResourceModel, getResponseData map[string]interface{}) *SslserviceSslcertkeyBindingResourceModel {
	tflog.Debug(ctx, "In sslservice_sslcertkey_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ca"]; ok && val != nil {
		data.Ca = types.BoolValue(val.(bool))
	} else {
		data.Ca = types.BoolNull()
	}
	if val, ok := getResponseData["certkeyname"]; ok && val != nil {
		data.Certkeyname = types.StringValue(val.(string))
	} else {
		data.Certkeyname = types.StringNull()
	}
	if val, ok := getResponseData["crlcheck"]; ok && val != nil {
		data.Crlcheck = types.StringValue(val.(string))
	} else {
		data.Crlcheck = types.StringNull()
	}
	if val, ok := getResponseData["ocspcheck"]; ok && val != nil {
		data.Ocspcheck = types.StringValue(val.(string))
	} else {
		data.Ocspcheck = types.StringNull()
	}
	if val, ok := getResponseData["servicename"]; ok && val != nil {
		data.Servicename = types.StringValue(val.(string))
	} else {
		data.Servicename = types.StringNull()
	}
	if val, ok := getResponseData["skipcaname"]; ok && val != nil {
		data.Skipcaname = types.BoolValue(val.(bool))
	} else {
		data.Skipcaname = types.BoolNull()
	}
	if val, ok := getResponseData["snicert"]; ok && val != nil {
		data.Snicert = types.BoolValue(val.(bool))
	} else {
		data.Snicert = types.BoolNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ca:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Ca.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("certkeyname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Certkeyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicename:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Servicename.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
