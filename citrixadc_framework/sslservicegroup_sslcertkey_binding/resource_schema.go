package sslservicegroup_sslcertkey_binding

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

// SslservicegroupSslcertkeyBindingResourceModel describes the resource data model.
type SslservicegroupSslcertkeyBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Ca               types.Bool   `tfsdk:"ca"`
	Certkeyname      types.String `tfsdk:"certkeyname"`
	Crlcheck         types.String `tfsdk:"crlcheck"`
	Ocspcheck        types.String `tfsdk:"ocspcheck"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
	Snicert          types.Bool   `tfsdk:"snicert"`
}

func (r *SslservicegroupSslcertkeyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslservicegroup_sslcertkey_binding resource.",
			},
			"ca": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CA certificate.",
			},
			"certkeyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the certificate bound to the SSL service group.",
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
			"servicegroupname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the SSL service to which the SSL policy needs to be bound.",
			},
			"snicert": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.",
			},
		},
	}
}

func sslservicegroup_sslcertkey_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SslservicegroupSslcertkeyBindingResourceModel) ssl.Sslservicegroupsslcertkeybinding {
	tflog.Debug(ctx, "In sslservicegroup_sslcertkey_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslservicegroup_sslcertkey_binding := ssl.Sslservicegroupsslcertkeybinding{}
	if !data.Ca.IsNull() {
		sslservicegroup_sslcertkey_binding.Ca = data.Ca.ValueBool()
	}
	if !data.Certkeyname.IsNull() {
		sslservicegroup_sslcertkey_binding.Certkeyname = data.Certkeyname.ValueString()
	}
	if !data.Crlcheck.IsNull() {
		sslservicegroup_sslcertkey_binding.Crlcheck = data.Crlcheck.ValueString()
	}
	if !data.Ocspcheck.IsNull() {
		sslservicegroup_sslcertkey_binding.Ocspcheck = data.Ocspcheck.ValueString()
	}
	if !data.Servicegroupname.IsNull() {
		sslservicegroup_sslcertkey_binding.Servicegroupname = data.Servicegroupname.ValueString()
	}
	if !data.Snicert.IsNull() {
		sslservicegroup_sslcertkey_binding.Snicert = data.Snicert.ValueBool()
	}

	return sslservicegroup_sslcertkey_binding
}

func sslservicegroup_sslcertkey_bindingSetAttrFromGet(ctx context.Context, data *SslservicegroupSslcertkeyBindingResourceModel, getResponseData map[string]interface{}) *SslservicegroupSslcertkeyBindingResourceModel {
	tflog.Debug(ctx, "In sslservicegroup_sslcertkey_bindingSetAttrFromGet Function")

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
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	} else {
		data.Servicegroupname = types.StringNull()
	}
	if val, ok := getResponseData["snicert"]; ok && val != nil {
		data.Snicert = types.BoolValue(val.(bool))
	} else {
		data.Snicert = types.BoolNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ca:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ca.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("certkeyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Certkeyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
