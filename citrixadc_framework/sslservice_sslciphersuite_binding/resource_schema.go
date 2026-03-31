package sslservice_sslciphersuite_binding

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

// SslserviceSslciphersuiteBindingResourceModel describes the resource data model.
type SslserviceSslciphersuiteBindingResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Cipherdefaulton types.Int64  `tfsdk:"cipherdefaulton"`
	Ciphername      types.String `tfsdk:"ciphername"`
	Description     types.String `tfsdk:"description"`
	Servicename     types.String `tfsdk:"servicename"`
}

func (r *SslserviceSslciphersuiteBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslservice_sslciphersuite_binding resource.",
			},
			"cipherdefaulton": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Flag indicating whether the bound cipher was the DEFAULT cipher, bound at boot time, or any other cipher from the CLI",
			},
			"ciphername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The cipher group/alias/individual cipher configuration",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The cipher suite description.",
			},
			"servicename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL service for which to set advanced configuration.",
			},
		},
	}
}

func sslservice_sslciphersuite_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SslserviceSslciphersuiteBindingResourceModel) ssl.Sslservicesslciphersuitebinding {
	tflog.Debug(ctx, "In sslservice_sslciphersuite_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslservice_sslciphersuite_binding := ssl.Sslservicesslciphersuitebinding{}
	if !data.Cipherdefaulton.IsNull() {
		sslservice_sslciphersuite_binding.Cipherdefaulton = utils.IntPtr(int(data.Cipherdefaulton.ValueInt64()))
	}
	if !data.Ciphername.IsNull() {
		sslservice_sslciphersuite_binding.Ciphername = data.Ciphername.ValueString()
	}
	if !data.Description.IsNull() {
		sslservice_sslciphersuite_binding.Description = data.Description.ValueString()
	}
	if !data.Servicename.IsNull() {
		sslservice_sslciphersuite_binding.Servicename = data.Servicename.ValueString()
	}

	return sslservice_sslciphersuite_binding
}

func sslservice_sslciphersuite_bindingSetAttrFromGet(ctx context.Context, data *SslserviceSslciphersuiteBindingResourceModel, getResponseData map[string]interface{}) *SslserviceSslciphersuiteBindingResourceModel {
	tflog.Debug(ctx, "In sslservice_sslciphersuite_bindingSetAttrFromGet Function")

	// Convert API response to model
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
