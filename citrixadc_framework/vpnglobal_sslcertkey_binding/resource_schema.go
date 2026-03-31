package vpnglobal_sslcertkey_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpnglobalSslcertkeyBindingResourceModel describes the resource data model.
type VpnglobalSslcertkeyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Cacert                 types.String `tfsdk:"cacert"`
	Certkeyname            types.String `tfsdk:"certkeyname"`
	Crlcheck               types.String `tfsdk:"crlcheck"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Ocspcheck              types.String `tfsdk:"ocspcheck"`
	Userdataencryptionkey  types.String `tfsdk:"userdataencryptionkey"`
}

func (r *VpnglobalSslcertkeyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_sslcertkey_binding resource.",
			},
			"cacert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the CA certificate binding.",
			},
			"certkeyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SSL certkey to use in signing tokens. Only RSA cert key is allowed",
			},
			"crlcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The state of the CRL check parameter (Mandatory/Optional).",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"ocspcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The state of the OCSP check parameter (Mandatory/Optional).",
			},
			"userdataencryptionkey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Certificate to be used for encrypting user data like KB Question and Answers, Alternate Email Address, etc.",
			},
		},
	}
}

func vpnglobal_sslcertkey_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnglobalSslcertkeyBindingResourceModel) vpn.Vpnglobalsslcertkeybinding {
	tflog.Debug(ctx, "In vpnglobal_sslcertkey_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnglobal_sslcertkey_binding := vpn.Vpnglobalsslcertkeybinding{}
	if !data.Cacert.IsNull() {
		vpnglobal_sslcertkey_binding.Cacert = data.Cacert.ValueString()
	}
	if !data.Certkeyname.IsNull() {
		vpnglobal_sslcertkey_binding.Certkeyname = data.Certkeyname.ValueString()
	}
	if !data.Crlcheck.IsNull() {
		vpnglobal_sslcertkey_binding.Crlcheck = data.Crlcheck.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() {
		vpnglobal_sslcertkey_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Ocspcheck.IsNull() {
		vpnglobal_sslcertkey_binding.Ocspcheck = data.Ocspcheck.ValueString()
	}
	if !data.Userdataencryptionkey.IsNull() {
		vpnglobal_sslcertkey_binding.Userdataencryptionkey = data.Userdataencryptionkey.ValueString()
	}

	return vpnglobal_sslcertkey_binding
}

func vpnglobal_sslcertkey_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalSslcertkeyBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalSslcertkeyBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_sslcertkey_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["cacert"]; ok && val != nil {
		data.Cacert = types.StringValue(val.(string))
	} else {
		data.Cacert = types.StringNull()
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
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["ocspcheck"]; ok && val != nil {
		data.Ocspcheck = types.StringValue(val.(string))
	} else {
		data.Ocspcheck = types.StringNull()
	}
	if val, ok := getResponseData["userdataencryptionkey"]; ok && val != nil {
		data.Userdataencryptionkey = types.StringValue(val.(string))
	} else {
		data.Userdataencryptionkey = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("cacert:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Cacert.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("certkeyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Certkeyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("userdataencryptionkey:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Userdataencryptionkey.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
