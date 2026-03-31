package sslservicegroup_sslciphersuite_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslservicegroupSslciphersuiteBindingResourceModel describes the resource data model.
type SslservicegroupSslciphersuiteBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Ciphername       types.String `tfsdk:"ciphername"`
	Description      types.String `tfsdk:"description"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
}

func (r *SslservicegroupSslciphersuiteBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslservicegroup_sslciphersuite_binding resource.",
			},
			"ciphername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the cipher group/alias/name configured for the SSL service group.",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The description of the cipher.",
			},
			"servicegroupname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the SSL service to which the SSL policy needs to be bound.",
			},
		},
	}
}

func sslservicegroup_sslciphersuite_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SslservicegroupSslciphersuiteBindingResourceModel) ssl.Sslservicegroupsslciphersuitebinding {
	tflog.Debug(ctx, "In sslservicegroup_sslciphersuite_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslservicegroup_sslciphersuite_binding := ssl.Sslservicegroupsslciphersuitebinding{}
	if !data.Ciphername.IsNull() {
		sslservicegroup_sslciphersuite_binding.Ciphername = data.Ciphername.ValueString()
	}
	if !data.Description.IsNull() {
		sslservicegroup_sslciphersuite_binding.Description = data.Description.ValueString()
	}
	if !data.Servicegroupname.IsNull() {
		sslservicegroup_sslciphersuite_binding.Servicegroupname = data.Servicegroupname.ValueString()
	}

	return sslservicegroup_sslciphersuite_binding
}

func sslservicegroup_sslciphersuite_bindingSetAttrFromGet(ctx context.Context, data *SslservicegroupSslciphersuiteBindingResourceModel, getResponseData map[string]interface{}) *SslservicegroupSslciphersuiteBindingResourceModel {
	tflog.Debug(ctx, "In sslservicegroup_sslciphersuite_bindingSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	} else {
		data.Servicegroupname = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ciphername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ciphername.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
