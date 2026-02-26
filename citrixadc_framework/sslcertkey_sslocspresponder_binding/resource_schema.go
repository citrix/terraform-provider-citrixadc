package sslcertkey_sslocspresponder_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "The certificate-key pair being unbound is a Certificate Authority (CA) certificate. If you choose this option, the certificate-key pair is unbound from the list of CA certificates that were bound to the specified SSL virtual server or SSL service.",
			},
			"certkey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the certificate-key pair.",
			},
			"ocspresponder": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "OCSP responders bound to this certkey",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ocsp priority",
			},
		},
	}
}

func sslcertkey_sslocspresponder_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SslcertkeySslocspresponderBindingResourceModel) ssl.Sslcertkeysslocspresponderbinding {
	tflog.Debug(ctx, "In sslcertkey_sslocspresponder_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslcertkey_sslocspresponder_binding := ssl.Sslcertkeysslocspresponderbinding{}
	if !data.Ca.IsNull() {
		sslcertkey_sslocspresponder_binding.Ca = data.Ca.ValueBool()
	}
	if !data.Certkey.IsNull() {
		sslcertkey_sslocspresponder_binding.Certkey = data.Certkey.ValueString()
	}
	if !data.Ocspresponder.IsNull() {
		sslcertkey_sslocspresponder_binding.Ocspresponder = data.Ocspresponder.ValueString()
	}
	if !data.Priority.IsNull() {
		sslcertkey_sslocspresponder_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return sslcertkey_sslocspresponder_binding
}

func sslcertkey_sslocspresponder_bindingSetAttrFromGet(ctx context.Context, data *SslcertkeySslocspresponderBindingResourceModel, getResponseData map[string]interface{}) *SslcertkeySslocspresponderBindingResourceModel {
	tflog.Debug(ctx, "In sslcertkey_sslocspresponder_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ca"]; ok && val != nil {
		data.Ca = types.BoolValue(val.(bool))
	} else {
		data.Ca = types.BoolNull()
	}
	if val, ok := getResponseData["certkey"]; ok && val != nil {
		data.Certkey = types.StringValue(val.(string))
	} else {
		data.Certkey = types.StringNull()
	}
	if val, ok := getResponseData["ocspresponder"]; ok && val != nil {
		data.Ocspresponder = types.StringValue(val.(string))
	} else {
		data.Ocspresponder = types.StringNull()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else {
		data.Priority = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("certkey:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Certkey.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("ocspresponder:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Ocspresponder.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
