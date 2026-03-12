package sslvserver_sslciphersuite_binding

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

// SslvserverSslciphersuiteBindingResourceModel describes the resource data model.
type SslvserverSslciphersuiteBindingResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Ciphername  types.String `tfsdk:"ciphername"`
	Description types.String `tfsdk:"description"`
	Vservername types.String `tfsdk:"vservername"`
}

func (r *SslvserverSslciphersuiteBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslvserver_sslciphersuite_binding resource.",
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
			"vservername": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL virtual server.",
			},
		},
	}
}

func sslvserver_sslciphersuite_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SslvserverSslciphersuiteBindingResourceModel) ssl.Sslvserversslciphersuitebinding {
	tflog.Debug(ctx, "In sslvserver_sslciphersuite_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslvserver_sslciphersuite_binding := ssl.Sslvserversslciphersuitebinding{}
	if !data.Ciphername.IsNull() {
		sslvserver_sslciphersuite_binding.Ciphername = data.Ciphername.ValueString()
	}
	if !data.Description.IsNull() {
		sslvserver_sslciphersuite_binding.Description = data.Description.ValueString()
	}
	if !data.Vservername.IsNull() {
		sslvserver_sslciphersuite_binding.Vservername = data.Vservername.ValueString()
	}

	return sslvserver_sslciphersuite_binding
}

func sslvserver_sslciphersuite_bindingSetAttrFromGet(ctx context.Context, data *SslvserverSslciphersuiteBindingResourceModel, getResponseData map[string]interface{}) *SslvserverSslciphersuiteBindingResourceModel {
	tflog.Debug(ctx, "In sslvserver_sslciphersuite_bindingSetAttrFromGet Function")

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
	if val, ok := getResponseData["vservername"]; ok && val != nil {
		data.Vservername = types.StringValue(val.(string))
	} else {
		data.Vservername = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ciphername:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Ciphername.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vservername:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Vservername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
