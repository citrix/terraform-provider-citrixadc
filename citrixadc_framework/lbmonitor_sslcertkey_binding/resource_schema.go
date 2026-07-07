package lbmonitor_sslcertkey_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "The rule for use of CRL corresponding to this CA certificate during client authentication. If crlCheck is set to Mandatory, the system will deny all SSL clients if the CRL is missing, expired - NextUpdate date is in the past, or is incomplete with remote CRL refresh enabled. If crlCheck is set to optional, the system will allow SSL clients in the above error cases.However, in any case if the client certificate is revoked in the CRL, the SSL client will be denied access.",
			},
			"certkeyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the certificate bound to the monitor.",
			},
			"crlcheck": schema.StringAttribute{
				// NITRO GET never echoes crlcheck back, so Computed would leave the
				// value unknown after apply ("inconsistent result"). Optional-only. (Pattern 13)
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The state of the CRL check parameter. (Mandatory/Optional)",
			},
			"monitorname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the monitor.",
			},
			"ocspcheck": schema.StringAttribute{
				// NITRO GET echoes ocspcheck only for CA bindings, so Computed would
				// leave it unknown after apply for non-CA bindings. Optional-only. (Pattern 13)
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The state of the OCSP check parameter. (Mandatory/Optional)",
			},
		},
	}
}

func lbmonitor_sslcertkey_bindingGetThePayloadFromthePlan(ctx context.Context, data *LbmonitorSslcertkeyBindingResourceModel) lb.Lbmonitorsslcertkeybinding {
	tflog.Debug(ctx, "In lbmonitor_sslcertkey_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	lbmonitor_sslcertkey_binding := lb.Lbmonitorsslcertkeybinding{}
	if !data.Ca.IsNull() && !data.Ca.IsUnknown() {
		lbmonitor_sslcertkey_binding.Ca = data.Ca.ValueBool()
	}
	if !data.Certkeyname.IsNull() && !data.Certkeyname.IsUnknown() {
		lbmonitor_sslcertkey_binding.Certkeyname = data.Certkeyname.ValueString()
	}
	if !data.Crlcheck.IsNull() && !data.Crlcheck.IsUnknown() {
		lbmonitor_sslcertkey_binding.Crlcheck = data.Crlcheck.ValueString()
	}
	if !data.Monitorname.IsNull() && !data.Monitorname.IsUnknown() {
		lbmonitor_sslcertkey_binding.Monitorname = data.Monitorname.ValueString()
	}
	if !data.Ocspcheck.IsNull() && !data.Ocspcheck.IsUnknown() {
		lbmonitor_sslcertkey_binding.Ocspcheck = data.Ocspcheck.ValueString()
	}

	return lbmonitor_sslcertkey_binding
}

// lbmonitor_sslcertkey_bindingSetAttrFromGet is the RESOURCE-side setter.
// It preserves the prior plan/state value for fields the NITRO GET response does
// not echo back (ca/crlcheck/ocspcheck are Optional+Computed and the appliance may
// omit them when left at default). Overwriting them with null would trip Terraform's
// "inconsistent result after apply" check. (Pattern 7)
// It does NOT touch data.Id — the ID is composed once in Create and preserved in Update.
func lbmonitor_sslcertkey_bindingSetAttrFromGet(ctx context.Context, data *LbmonitorSslcertkeyBindingResourceModel, getResponseData map[string]interface{}) *LbmonitorSslcertkeyBindingResourceModel {
	tflog.Debug(ctx, "In lbmonitor_sslcertkey_bindingSetAttrFromGet Function")

	// ca is Computed and is always echoed by GET (true/false) — adopt it.
	if val, ok := getResponseData["ca"]; ok && val != nil {
		data.Ca = types.BoolValue(val.(bool))
	}
	if val, ok := getResponseData["certkeyname"]; ok && val != nil {
		data.Certkeyname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["monitorname"]; ok && val != nil {
		data.Monitorname = types.StringValue(val.(string))
	}
	// crlcheck and ocspcheck are Optional-only inputs. NITRO either omits them or
	// returns a server default (e.g. ocspcheck="Optional" for CA bindings) that the
	// user never configured. Adopting that would trip "inconsistent result after
	// apply" for non-CA bindings. Preserve the existing plan/state value. (Pattern 7)

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("monitorname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Monitorname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("certkeyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Certkeyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// lbmonitor_sslcertkey_bindingSetAttrFromGetForDatasource is the DATASOURCE-side
// setter. The datasource has no prior plan/state to preserve, so it faithfully
// copies every field from the GET response and sets its own composite ID. (Pattern 7)
func lbmonitor_sslcertkey_bindingSetAttrFromGetForDatasource(ctx context.Context, data *LbmonitorSslcertkeyBindingResourceModel, getResponseData map[string]interface{}) *LbmonitorSslcertkeyBindingResourceModel {
	tflog.Debug(ctx, "In lbmonitor_sslcertkey_bindingSetAttrFromGetForDatasource Function")

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

	// Set the composite ID for the datasource (no Create to do it).
	// Legacy SDK v2 ID order: monitorname,certkeyname (see resource_id_mapping.json).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("monitorname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Monitorname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("certkeyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Certkeyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
