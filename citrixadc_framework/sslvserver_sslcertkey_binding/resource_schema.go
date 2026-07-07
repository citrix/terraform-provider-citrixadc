package sslvserver_sslcertkey_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
				Optional: true,
				Computed: true,
				// ca participates in the binding identity and the ADC default is
				// false; pin a default so an omitted ca is a known value at plan
				// time (the NITRO GET does not always echo ca:false back).
				Default: booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "CA certificate.",
			},
			"certkeyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the certificate key pair binding.",
			},
			"crlcheck": schema.StringAttribute{
				// Optional-only (no Computed): the NITRO GET does not echo a
				// server default for crlcheck, so a Computed value could never be
				// resolved when omitted -> "unknown value after apply". An omitted
				// crlcheck is simply null.
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The state of the CRL check parameter. (Mandatory/Optional)",
			},
			"ocspcheck": schema.StringAttribute{
				// Optional-only (see crlcheck rationale).
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The state of the OCSP check parameter. (Mandatory/Optional)",
			},
			"skipcaname": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				// ADC default is false; pin it so an omitted skipcaname is a known
				// value at plan time (NITRO GET does not echo skipcaname:false).
				Default: booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting for client certificate in a SSL handshake",
			},
			"snicert": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				// SDK v2 had Default: false on snicert. Preserve that contract so
				// configs that omit snicert get a deterministic, known value.
				Default: booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.",
			},
			"vservername": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the SSL virtual server.",
			},
		},
	}
}

func sslvserver_sslcertkey_bindingGetThePayloadFromthePlan(ctx context.Context, data *SslvserverSslcertkeyBindingResourceModel) ssl.Sslvserversslcertkeybinding {
	tflog.Debug(ctx, "In sslvserver_sslcertkey_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslvserver_sslcertkey_binding := ssl.Sslvserversslcertkeybinding{}
	if !data.Ca.IsNull() && !data.Ca.IsUnknown() {
		sslvserver_sslcertkey_binding.Ca = data.Ca.ValueBool()
	}
	if !data.Certkeyname.IsNull() && !data.Certkeyname.IsUnknown() {
		sslvserver_sslcertkey_binding.Certkeyname = data.Certkeyname.ValueString()
	}
	if !data.Crlcheck.IsNull() && !data.Crlcheck.IsUnknown() {
		sslvserver_sslcertkey_binding.Crlcheck = data.Crlcheck.ValueString()
	}
	if !data.Ocspcheck.IsNull() && !data.Ocspcheck.IsUnknown() {
		sslvserver_sslcertkey_binding.Ocspcheck = data.Ocspcheck.ValueString()
	}
	if !data.Skipcaname.IsNull() && !data.Skipcaname.IsUnknown() {
		sslvserver_sslcertkey_binding.Skipcaname = data.Skipcaname.ValueBool()
	}
	if !data.Snicert.IsNull() && !data.Snicert.IsUnknown() {
		sslvserver_sslcertkey_binding.Snicert = data.Snicert.ValueBool()
	}
	if !data.Vservername.IsNull() && !data.Vservername.IsUnknown() {
		sslvserver_sslcertkey_binding.Vservername = data.Vservername.ValueString()
	}

	return sslvserver_sslcertkey_binding
}

// sslvserver_sslcertkey_bindingSetAttrFromGet is the RESOURCE-side setter. The
// NITRO GET for a binding only echoes back fields that were actually set on the
// bound certkey, and never returns false bools / unset strings. To avoid
// "inconsistent result after apply" (Pattern 7) and ID-wipe (Pattern 6), it
// preserves the existing plan/state value for any field the API omits, and it
// does NOT recompute data.Id (the ID is set exactly once in Create).
func sslvserver_sslcertkey_bindingSetAttrFromGet(ctx context.Context, data *SslvserverSslcertkeyBindingResourceModel, getResponseData map[string]interface{}) *SslvserverSslcertkeyBindingResourceModel {
	tflog.Debug(ctx, "In sslvserver_sslcertkey_bindingSetAttrFromGet Function")

	// Convert API response to model, preserving the plan/state value when the
	// API does not echo the field back.
	if val, ok := getResponseData["ca"]; ok && val != nil {
		data.Ca = types.BoolValue(val.(bool))
	}
	if val, ok := getResponseData["certkeyname"]; ok && val != nil {
		data.Certkeyname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["crlcheck"]; ok && val != nil {
		data.Crlcheck = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["ocspcheck"]; ok && val != nil {
		data.Ocspcheck = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["skipcaname"]; ok && val != nil {
		data.Skipcaname = types.BoolValue(val.(bool))
	}
	if val, ok := getResponseData["snicert"]; ok && val != nil {
		data.Snicert = types.BoolValue(val.(bool))
	}
	if val, ok := getResponseData["vservername"]; ok && val != nil {
		data.Vservername = types.StringValue(val.(string))
	}

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("vservername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vservername.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("certkeyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Certkeyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("snicert:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Snicert.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("ca:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ca.ValueBool()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// sslvserver_sslcertkey_bindingSetAttrFromGetForDatasource is the DATASOURCE-side
// setter. The datasource has no prior plan/state, so it faithfully copies every
// field from the GET response and sets data.Id itself (Pattern 7 split).
func sslvserver_sslcertkey_bindingSetAttrFromGetForDatasource(ctx context.Context, data *SslvserverSslcertkeyBindingResourceModel, getResponseData map[string]interface{}) *SslvserverSslcertkeyBindingResourceModel {
	tflog.Debug(ctx, "In sslvserver_sslcertkey_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["ca"]; ok && val != nil {
		data.Ca = types.BoolValue(val.(bool))
	} else {
		data.Ca = types.BoolValue(false)
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
		data.Snicert = types.BoolValue(false)
	}
	if val, ok := getResponseData["vservername"]; ok && val != nil {
		data.Vservername = types.StringValue(val.(string))
	} else {
		data.Vservername = types.StringNull()
	}

	// Set ID for the datasource (no Create to set it).
	// Multiple unique attributes - comma-separated key:UrlEncode(value) pairs,
	// in legacy resource_id_mapping.json order.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("vservername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vservername.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("certkeyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Certkeyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("snicert:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Snicert.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("ca:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ca.ValueBool()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
