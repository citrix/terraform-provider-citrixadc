package sslservice_sslcipher_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslserviceSslcipherBindingResourceModel describes the resource data model.
type SslserviceSslcipherBindingResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Cipheraliasname types.String `tfsdk:"cipheraliasname"`
	Cipherdefaulton types.Int64  `tfsdk:"cipherdefaulton"`
	Ciphername      types.String `tfsdk:"ciphername"`
	Description     types.String `tfsdk:"description"`
	Servicename     types.String `tfsdk:"servicename"`
}

func (r *SslserviceSslcipherBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslservice_sslcipher_binding resource.",
			},
			"cipheraliasname": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "The cipher group/alias/individual cipher configuration.",
			},
			"cipherdefaulton": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Description: "Flag indicating whether the bound cipher was the DEFAULT cipher, bound at boot time, or any other cipher from the CLI",
			},
			"ciphername": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the individual cipher, user-defined cipher group, or predefined (built-in) cipher alias.",
			},
			"description": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "The cipher suite description.",
			},
			"servicename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the SSL service for which to set advanced configuration.",
			},
		},
	}
}

func sslservice_sslcipher_bindingGetThePayloadFromthePlan(ctx context.Context, data *SslserviceSslcipherBindingResourceModel) ssl.Sslservicesslcipherbinding {
	tflog.Debug(ctx, "In sslservice_sslcipher_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	// Pattern 15: cipheraliasname, cipherdefaulton, description are read-only / GET-only
	// attributes (NITRO rejects them in the bind payload). They are surfaced as Computed
	// only and must never be sent in the add/bind payload.
	sslservice_sslcipher_binding := ssl.Sslservicesslcipherbinding{}
	if !data.Ciphername.IsNull() && !data.Ciphername.IsUnknown() {
		sslservice_sslcipher_binding.Ciphername = data.Ciphername.ValueString()
	}
	if !data.Servicename.IsNull() && !data.Servicename.IsUnknown() {
		sslservice_sslcipher_binding.Servicename = data.Servicename.ValueString()
	}

	return sslservice_sslcipher_binding
}

func sslservice_sslcipher_bindingSetAttrFromGet(ctx context.Context, data *SslserviceSslcipherBindingResourceModel, getResponseData map[string]interface{}) *SslserviceSslcipherBindingResourceModel {
	tflog.Debug(ctx, "In sslservice_sslcipher_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["cipheraliasname"]; ok && val != nil {
		data.Cipheraliasname = types.StringValue(val.(string))
	} else {
		data.Cipheraliasname = types.StringNull()
	}
	if val, ok := getResponseData["cipherdefaulton"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cipherdefaulton = types.Int64Value(intVal)
		}
	} else {
		data.Cipherdefaulton = types.Int64Null()
	}
	if val, ok := getResponseData["ciphername"]; ok && val != nil {
		data.Ciphername = types.StringValue(val.(string))
	} else {
		data.Ciphername = types.StringNull()
	}
	if val, ok := getResponseData["description"]; ok && val != nil {
		data.Description = types.StringValue(val.(string))
	} else {
		data.Description = types.StringNull()
	}
	if val, ok := getResponseData["servicename"]; ok && val != nil {
		data.Servicename = types.StringValue(val.(string))
	} else {
		data.Servicename = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ciphername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ciphername.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicename.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
