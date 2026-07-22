package sslservicegroup_sslcipher_binding

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

// SslservicegroupSslcipherBindingResourceModel describes the resource data model.
type SslservicegroupSslcipherBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Cipheraliasname  types.String `tfsdk:"cipheraliasname"`
	Ciphername       types.String `tfsdk:"ciphername"`
	Description      types.String `tfsdk:"description"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
}

func (r *SslservicegroupSslcipherBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslservicegroup_sslcipher_binding resource.",
			},
			"cipheraliasname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the cipher group/alias/name configured for the SSL service group.",
			},
			"ciphername": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "A cipher-suite can consist of an individual cipher name, the system predefined cipher-alias name, or user defined cipher-group name.",
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the SSL service to which the SSL policy needs to be bound.",
			},
		},
	}
}

func sslservicegroup_sslcipher_bindingGetThePayloadFromthePlan(ctx context.Context, data *SslservicegroupSslcipherBindingResourceModel) ssl.Sslservicegroupsslcipherbinding {
	tflog.Debug(ctx, "In sslservicegroup_sslcipher_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslservicegroup_sslcipher_binding := ssl.Sslservicegroupsslcipherbinding{}
	// Pattern 15: cipheraliasname and description are read-only/GET-only
	// properties; the bind (add) endpoint does not accept them.
	if !data.Ciphername.IsNull() && !data.Ciphername.IsUnknown() {
		sslservicegroup_sslcipher_binding.Ciphername = data.Ciphername.ValueString()
	}
	if !data.Servicegroupname.IsNull() && !data.Servicegroupname.IsUnknown() {
		sslservicegroup_sslcipher_binding.Servicegroupname = data.Servicegroupname.ValueString()
	}

	return sslservicegroup_sslcipher_binding
}

func sslservicegroup_sslcipher_bindingSetAttrFromGet(ctx context.Context, data *SslservicegroupSslcipherBindingResourceModel, getResponseData map[string]interface{}) *SslservicegroupSslcipherBindingResourceModel {
	tflog.Debug(ctx, "In sslservicegroup_sslcipher_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["cipheraliasname"]; ok && val != nil {
		data.Cipheraliasname = types.StringValue(val.(string))
	} else {
		data.Cipheraliasname = types.StringNull()
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
