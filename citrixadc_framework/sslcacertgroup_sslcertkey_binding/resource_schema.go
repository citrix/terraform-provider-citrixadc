package sslcacertgroup_sslcertkey_binding

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

// SslcacertgroupSslcertkeyBindingResourceModel describes the resource data model.
type SslcacertgroupSslcertkeyBindingResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Cacertgroupname types.String `tfsdk:"cacertgroupname"`
	Certkeyname     types.String `tfsdk:"certkeyname"`
	Crlcheck        types.String `tfsdk:"crlcheck"`
	Ocspcheck       types.String `tfsdk:"ocspcheck"`
}

func (r *SslcacertgroupSslcertkeyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslcacertgroup_sslcertkey_binding resource.",
			},
			"cacertgroupname": schema.StringAttribute{
				Required:    true,
				Description: "Name given to the CA certificate group. The name will be used to add the CA certificates to the group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my file\" or 'my file').",
			},
			"certkeyname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the certkey added to the Citrix ADC. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the certificate-key pair is created.The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cert\" or 'my cert').",
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
		},
	}
}

func sslcacertgroup_sslcertkey_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SslcacertgroupSslcertkeyBindingResourceModel) ssl.Sslcacertgroupsslcertkeybinding {
	tflog.Debug(ctx, "In sslcacertgroup_sslcertkey_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslcacertgroup_sslcertkey_binding := ssl.Sslcacertgroupsslcertkeybinding{}
	if !data.Cacertgroupname.IsNull() {
		sslcacertgroup_sslcertkey_binding.Cacertgroupname = data.Cacertgroupname.ValueString()
	}
	if !data.Certkeyname.IsNull() {
		sslcacertgroup_sslcertkey_binding.Certkeyname = data.Certkeyname.ValueString()
	}
	if !data.Crlcheck.IsNull() {
		sslcacertgroup_sslcertkey_binding.Crlcheck = data.Crlcheck.ValueString()
	}
	if !data.Ocspcheck.IsNull() {
		sslcacertgroup_sslcertkey_binding.Ocspcheck = data.Ocspcheck.ValueString()
	}

	return sslcacertgroup_sslcertkey_binding
}

func sslcacertgroup_sslcertkey_bindingSetAttrFromGet(ctx context.Context, data *SslcacertgroupSslcertkeyBindingResourceModel, getResponseData map[string]interface{}) *SslcacertgroupSslcertkeyBindingResourceModel {
	tflog.Debug(ctx, "In sslcacertgroup_sslcertkey_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["cacertgroupname"]; ok && val != nil {
		data.Cacertgroupname = types.StringValue(val.(string))
	} else {
		data.Cacertgroupname = types.StringNull()
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

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("cacertgroupname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Cacertgroupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("certkeyname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Certkeyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
