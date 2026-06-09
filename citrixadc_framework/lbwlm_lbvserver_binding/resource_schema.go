package lbwlm_lbvserver_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LbwlmLbvserverBindingResourceModel describes the resource data model.
type LbwlmLbvserverBindingResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Vservername types.String `tfsdk:"vservername"`
	Wlmname     types.String `tfsdk:"wlmname"`
}

func (r *LbwlmLbvserverBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbwlm_lbvserver_binding resource.",
			},
			"vservername": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the virtual server which is to be bound to the WLM.",
			},
			"wlmname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the Work Load Manager.",
			},
		},
	}
}

func lbwlm_lbvserver_bindingGetThePayloadFromthePlan(ctx context.Context, data *LbwlmLbvserverBindingResourceModel) lb.Lbwlmlbvserverbinding {
	tflog.Debug(ctx, "In lbwlm_lbvserver_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	lbwlm_lbvserver_binding := lb.Lbwlmlbvserverbinding{}
	if !data.Vservername.IsNull() && !data.Vservername.IsUnknown() {
		lbwlm_lbvserver_binding.Vservername = data.Vservername.ValueString()
	}
	if !data.Wlmname.IsNull() && !data.Wlmname.IsUnknown() {
		lbwlm_lbvserver_binding.Wlmname = data.Wlmname.ValueString()
	}

	return lbwlm_lbvserver_binding
}

func lbwlm_lbvserver_bindingSetAttrFromGet(ctx context.Context, data *LbwlmLbvserverBindingResourceModel, getResponseData map[string]interface{}) *LbwlmLbvserverBindingResourceModel {
	tflog.Debug(ctx, "In lbwlm_lbvserver_bindingSetAttrFromGet Function")

	// Resource path: both attributes are RequiresReplace composite-key fields.
	// Preserve plan/state values and do NOT recompute the ID (set once in Create).
	return data
}

func lbwlm_lbvserver_bindingSetAttrFromGetForDatasource(ctx context.Context, data *LbwlmLbvserverBindingResourceModel, getResponseData map[string]interface{}) *LbwlmLbvserverBindingResourceModel {
	tflog.Debug(ctx, "In lbwlm_lbvserver_bindingSetAttrFromGetForDatasource Function")

	// Datasource has no prior state; faithfully copy every field from the GET response.
	if val, ok := getResponseData["vservername"]; ok && val != nil {
		data.Vservername = types.StringValue(val.(string))
	} else {
		data.Vservername = types.StringNull()
	}
	if val, ok := getResponseData["wlmname"]; ok && val != nil {
		data.Wlmname = types.StringValue(val.(string))
	} else {
		data.Wlmname = types.StringNull()
	}

	// Compose the composite ID (datasource has no Create to set it)
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("vservername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vservername.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("wlmname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Wlmname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
