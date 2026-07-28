package sslvserver_sslcipher_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslvserverSslcipherBindingResourceModel describes the resource data model.
type SslvserverSslcipherBindingResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Cipheraliasname types.String `tfsdk:"cipheraliasname"`
	Ciphername      types.String `tfsdk:"ciphername"`
	Description     types.String `tfsdk:"description"`
	Vservername     types.String `tfsdk:"vservername"`
}

func (r *SslvserverSslcipherBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslvserver_sslcipher_binding resource.",
			},
			"cipheraliasname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the cipher group/alias/individual cipheri bindings.",
			},
			"ciphername": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the individual cipher, user-defined cipher group, or predefined (built-in) cipher alias.",
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the SSL virtual server.",
			},
		},
	}
}

func sslvserver_sslcipher_bindingGetThePayloadFromthePlan(ctx context.Context, data *SslvserverSslcipherBindingResourceModel) ssl.Sslvserversslcipherbinding {
	tflog.Debug(ctx, "In sslvserver_sslcipher_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslvserver_sslcipher_binding := ssl.Sslvserversslcipherbinding{}
	if !data.Cipheraliasname.IsNull() && !data.Cipheraliasname.IsUnknown() {
		sslvserver_sslcipher_binding.Cipheraliasname = data.Cipheraliasname.ValueString()
	}
	if !data.Ciphername.IsNull() && !data.Ciphername.IsUnknown() {
		sslvserver_sslcipher_binding.Ciphername = data.Ciphername.ValueString()
	}
	// "description" is a read-only/response-only field (present in GET but not in the
	// add payload per the NITRO doc); it is never sent in the bind payload (Pattern 15).
	if !data.Vservername.IsNull() && !data.Vservername.IsUnknown() {
		sslvserver_sslcipher_binding.Vservername = data.Vservername.ValueString()
	}

	return sslvserver_sslcipher_binding
}

func sslvserver_sslcipher_bindingSetAttrFromGet(ctx context.Context, data *SslvserverSslcipherBindingResourceModel, getResponseData map[string]interface{}) *SslvserverSslcipherBindingResourceModel {
	tflog.Debug(ctx, "In sslvserver_sslcipher_bindingSetAttrFromGet Function")

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
	if val, ok := getResponseData["vservername"]; ok && val != nil {
		data.Vservername = types.StringValue(val.(string))
	} else {
		data.Vservername = types.StringNull()
	}

	// ID is set once in Create (resource) / Read (datasource); do not recompute here
	// to avoid wiping it when a key field is absent from the GET response (Pattern 6).

	return data
}
