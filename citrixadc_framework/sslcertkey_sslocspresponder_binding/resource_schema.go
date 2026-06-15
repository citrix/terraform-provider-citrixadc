package sslcertkey_sslocspresponder_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslcertkeySslocspresponderBindingResourceModel describes the resource data model.
type SslcertkeySslocspresponderBindingResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Ca            types.Bool   `tfsdk:"ca"`
	Certkey       types.String `tfsdk:"certkey"`
	Ocspresponder types.String `tfsdk:"ocspresponder"`
	Priority      types.Int64  `tfsdk:"priority"`
}

func (r *SslcertkeySslocspresponderBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslcertkey_sslocspresponder_binding resource.",
			},
			"ca": schema.BoolAttribute{
				// "ca" is never echoed by the NITRO GET response for this binding, so it
				// cannot be Computed-resolved at apply time - keep it Optional only
				// (Pattern 13) to avoid "unknown value" / perpetual-diff errors.
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "The certificate-key pair being unbound is a Certificate Authority (CA) certificate. If you choose this option, the certificate-key pair is unbound from the list of CA certificates that were bound to the specified SSL virtual server or SSL service.",
			},
			"certkey": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the certificate-key pair.",
			},
			"ocspresponder": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "OCSP responders bound to this certkey",
			},
			"priority": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "ocsp priority",
			},
		},
	}
}

func sslcertkey_sslocspresponder_bindingGetThePayloadFromthePlan(ctx context.Context, data *SslcertkeySslocspresponderBindingResourceModel) ssl.Sslcertkeysslocspresponderbinding {
	tflog.Debug(ctx, "In sslcertkey_sslocspresponder_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslcertkey_sslocspresponder_binding := ssl.Sslcertkeysslocspresponderbinding{}
	if !data.Ca.IsNull() && !data.Ca.IsUnknown() {
		sslcertkey_sslocspresponder_binding.Ca = data.Ca.ValueBool()
	}
	if !data.Certkey.IsNull() && !data.Certkey.IsUnknown() {
		sslcertkey_sslocspresponder_binding.Certkey = data.Certkey.ValueString()
	}
	if !data.Ocspresponder.IsNull() && !data.Ocspresponder.IsUnknown() {
		sslcertkey_sslocspresponder_binding.Ocspresponder = data.Ocspresponder.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		sslcertkey_sslocspresponder_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return sslcertkey_sslocspresponder_binding
}

// sslcertkey_sslocspresponder_bindingSetAttrFromGet is the resource-side state setter.
// It PRESERVES the existing plan/state values for identity attributes (certkey,
// ocspresponder) and for the non-echoed "ca" flag (the NITRO GET response for this
// binding never returns "ca"; see Pattern 7 / Pattern 13). It only copies back the
// server-echoed "priority" field. The ID is composed once in Create; it is not
// recomputed here (Pattern 6).
func sslcertkey_sslocspresponder_bindingSetAttrFromGet(ctx context.Context, data *SslcertkeySslocspresponderBindingResourceModel, getResponseData map[string]interface{}) *SslcertkeySslocspresponderBindingResourceModel {
	tflog.Debug(ctx, "In sslcertkey_sslocspresponder_bindingSetAttrFromGet Function")

	// certkey / ocspresponder are identity attributes carried in the ID and config -
	// preserve the existing value. "ca" is never echoed by the NITRO GET response, so
	// preserve the configured/state value rather than nulling it.
	if val, ok := getResponseData["certkey"]; ok && val != nil {
		data.Certkey = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["ocspresponder"]; ok && val != nil {
		data.Ocspresponder = types.StringValue(val.(string))
	}
	// priority is server-echoed (returned as a string) - copy it back.
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	}

	return data
}

// sslcertkey_sslocspresponder_bindingSetAttrFromGetForDatasource is the datasource-side
// state setter. The datasource has no prior plan/state to preserve, so it faithfully
// copies every field from the GET response and sets the ID itself (Pattern 7). The
// "ca" flag is not echoed by NITRO, so it is left as supplied in the datasource config.
func sslcertkey_sslocspresponder_bindingSetAttrFromGetForDatasource(ctx context.Context, data *SslcertkeySslocspresponderBindingResourceModel, getResponseData map[string]interface{}) *SslcertkeySslocspresponderBindingResourceModel {
	tflog.Debug(ctx, "In sslcertkey_sslocspresponder_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["certkey"]; ok && val != nil {
		data.Certkey = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["ocspresponder"]; ok && val != nil {
		data.Ocspresponder = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	}

	// Compose the legacy-compatible composite ID (certkey,ocspresponder).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("certkey:%s", utils.UrlEncode(data.Certkey.ValueString())))
	idParts = append(idParts, fmt.Sprintf("ocspresponder:%s", utils.UrlEncode(data.Ocspresponder.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
