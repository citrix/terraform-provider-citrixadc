package lbmonitor_sslcertkey_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LbmonitorSslcertkeyBindingResourceModel describes the resource data model.
type LbmonitorSslcertkeyBindingResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Ca          types.Bool   `tfsdk:"ca"`
	Certkeyname types.String `tfsdk:"certkeyname"`
	Crlcheck    types.String `tfsdk:"crlcheck"`
	Monitorname types.String `tfsdk:"monitorname"`
	Ocspcheck   types.String `tfsdk:"ocspcheck"`
}

func (r *LbmonitorSslcertkeyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbmonitor_sslcertkey_binding resource.",
			},
			"ca": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The rule for use of CRL corresponding to this CA certificate during client authentication. If crlCheck is set to Mandatory, the system will deny all SSL clients if the CRL is missing, expired - NextUpdate date is in the past, or is incomplete with remote CRL refresh enabled. If crlCheck is set to optional, the system will allow SSL clients in the above error cases.However, in any case if the client certificate is revoked in the CRL, the SSL client will be denied access.",
			},
			"certkeyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the certificate bound to the monitor.",
			},
			"crlcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The state of the CRL check parameter. (Mandatory/Optional)",
			},
			"monitorname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the monitor.",
			},
			"ocspcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The state of the OCSP check parameter. (Mandatory/Optional)",
			},
		},
	}
}

func lbmonitor_sslcertkey_bindingGetThePayloadFromtheConfig(ctx context.Context, data *LbmonitorSslcertkeyBindingResourceModel) lb.Lbmonitorsslcertkeybinding {
	tflog.Debug(ctx, "In lbmonitor_sslcertkey_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lbmonitor_sslcertkey_binding := lb.Lbmonitorsslcertkeybinding{}
	if !data.Ca.IsNull() {
		lbmonitor_sslcertkey_binding.Ca = data.Ca.ValueBool()
	}
	if !data.Certkeyname.IsNull() {
		lbmonitor_sslcertkey_binding.Certkeyname = data.Certkeyname.ValueString()
	}
	if !data.Crlcheck.IsNull() {
		lbmonitor_sslcertkey_binding.Crlcheck = data.Crlcheck.ValueString()
	}
	if !data.Monitorname.IsNull() {
		lbmonitor_sslcertkey_binding.Monitorname = data.Monitorname.ValueString()
	}
	if !data.Ocspcheck.IsNull() {
		lbmonitor_sslcertkey_binding.Ocspcheck = data.Ocspcheck.ValueString()
	}

	return lbmonitor_sslcertkey_binding
}

func lbmonitor_sslcertkey_bindingSetAttrFromGet(ctx context.Context, data *LbmonitorSslcertkeyBindingResourceModel, getResponseData map[string]interface{}) *LbmonitorSslcertkeyBindingResourceModel {
	tflog.Debug(ctx, "In lbmonitor_sslcertkey_bindingSetAttrFromGet Function")

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
	if val, ok := getResponseData["monitorname"]; ok && val != nil {
		data.Monitorname = types.StringValue(val.(string))
	} else {
		data.Monitorname = types.StringNull()
	}
	if val, ok := getResponseData["ocspcheck"]; ok && val != nil {
		data.Ocspcheck = types.StringValue(val.(string))
	} else {
		data.Ocspcheck = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ca:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Ca.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("certkeyname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Certkeyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("monitorname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Monitorname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
