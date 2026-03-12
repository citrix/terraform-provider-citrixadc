package sslvserver_sslcertkey_binding

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

// SslvserverSslcertkeyBindingResourceModel describes the resource data model.
type SslvserverSslcertkeyBindingResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Ca          types.Bool   `tfsdk:"ca"`
	Certkeyname types.String `tfsdk:"certkeyname"`
	Crlcheck    types.String `tfsdk:"crlcheck"`
	Ocspcheck   types.String `tfsdk:"ocspcheck"`
	Skipcaname  types.Bool   `tfsdk:"skipcaname"`
	Snicert     types.Bool   `tfsdk:"snicert"`
	Vservername types.String `tfsdk:"vservername"`
}

func (r *SslvserverSslcertkeyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslvserver_sslcertkey_binding resource.",
			},
			"ca": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CA certificate.",
			},
			"certkeyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the certificate key pair binding.",
			},
			"crlcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The state of the CRL check parameter. (Mandatory/Optional)",
			},
			"ocspcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The state of the OCSP check parameter. (Mandatory/Optional)",
			},
			"skipcaname": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting for client certificate in a SSL handshake",
			},
			"snicert": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.",
			},
			"vservername": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL virtual server.",
			},
		},
	}
}

func sslvserver_sslcertkey_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SslvserverSslcertkeyBindingResourceModel) ssl.Sslvserversslcertkeybinding {
	tflog.Debug(ctx, "In sslvserver_sslcertkey_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslvserver_sslcertkey_binding := ssl.Sslvserversslcertkeybinding{}
	if !data.Ca.IsNull() {
		sslvserver_sslcertkey_binding.Ca = data.Ca.ValueBool()
	}
	if !data.Certkeyname.IsNull() {
		sslvserver_sslcertkey_binding.Certkeyname = data.Certkeyname.ValueString()
	}
	if !data.Crlcheck.IsNull() {
		sslvserver_sslcertkey_binding.Crlcheck = data.Crlcheck.ValueString()
	}
	if !data.Ocspcheck.IsNull() {
		sslvserver_sslcertkey_binding.Ocspcheck = data.Ocspcheck.ValueString()
	}
	if !data.Skipcaname.IsNull() {
		sslvserver_sslcertkey_binding.Skipcaname = data.Skipcaname.ValueBool()
	}
	if !data.Snicert.IsNull() {
		sslvserver_sslcertkey_binding.Snicert = data.Snicert.ValueBool()
	}
	if !data.Vservername.IsNull() {
		sslvserver_sslcertkey_binding.Vservername = data.Vservername.ValueString()
	}

	return sslvserver_sslcertkey_binding
}

func sslvserver_sslcertkey_bindingSetAttrFromGet(ctx context.Context, data *SslvserverSslcertkeyBindingResourceModel, getResponseData map[string]interface{}) *SslvserverSslcertkeyBindingResourceModel {
	tflog.Debug(ctx, "In sslvserver_sslcertkey_bindingSetAttrFromGet Function")

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
	if val, ok := getResponseData["vservername"]; ok && val != nil {
		data.Vservername = types.StringValue(val.(string))
	} else {
		data.Vservername = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ca:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Ca.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("certkeyname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Certkeyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("crlcheck:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Crlcheck.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("ocspcheck:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Ocspcheck.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("snicert:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Snicert.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("vservername:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Vservername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
