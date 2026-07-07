package vpnglobal_intranetip_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpnglobalIntranetipBindingResourceModel describes the resource data model.
type VpnglobalIntranetipBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Intranetip             types.String `tfsdk:"intranetip"`
	Netmask                types.String `tfsdk:"netmask"`
}

func (r *VpnglobalIntranetipBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_intranetip_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				// Pattern 7/13: this field is NOT echoed back by the NITRO GET response for
				// this binding. Marking it Computed would cause "known after apply" churn /
				// "inconsistent result after apply". It is preserved from plan/state in
				// vpnglobal_intranetip_bindingSetAttrFromGet. RequiresReplace mirrors the
				// SDK v2 ForceNew on this attribute.
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"intranetip": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The intranet ip address or range.",
			},
			"netmask": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The intranet ip address or range's netmask.",
			},
		},
	}
}

func vpnglobal_intranetip_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnglobalIntranetipBindingResourceModel) vpn.Vpnglobalintranetipbinding {
	tflog.Debug(ctx, "In vpnglobal_intranetip_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnglobal_intranetip_binding := vpn.Vpnglobalintranetipbinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnglobal_intranetip_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Intranetip.IsNull() && !data.Intranetip.IsUnknown() {
		vpnglobal_intranetip_binding.Intranetip = data.Intranetip.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		vpnglobal_intranetip_binding.Netmask = data.Netmask.ValueString()
	}

	return vpnglobal_intranetip_binding
}

// vpnglobal_intranetip_bindingSetAttrFromGet is the RESOURCE setter. It preserves
// plan/state values for fields that NITRO does not echo back (gotopriorityexpression)
// and does NOT recompute the ID (the ID is set exactly once in Create / preserved in
// Update — see Pattern 6 and Pattern 7).
func vpnglobal_intranetip_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalIntranetipBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalIntranetipBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_intranetip_bindingSetAttrFromGet Function")

	// Convert API response to model.
	// gotopriorityexpression is NOT returned by the NITRO GET for this binding; preserve
	// the existing plan/state value instead of nulling it (Pattern 7).
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["intranetip"]; ok && val != nil {
		data.Intranetip = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	}

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("intranetip:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Intranetip.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("netmask:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Netmask.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// vpnglobal_intranetip_bindingSetAttrFromGetForDatasource is the DATASOURCE setter. The
// datasource has no prior plan/state to preserve, so it faithfully copies every field
// from the GET response and sets the composite ID itself (Pattern 7 datasource split).
func vpnglobal_intranetip_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VpnglobalIntranetipBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalIntranetipBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_intranetip_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["intranetip"]; ok && val != nil {
		data.Intranetip = types.StringValue(val.(string))
	} else {
		data.Intranetip = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}

	// Set ID for the datasource (no Create to do it).
	// Multiple unique attributes - comma-separated key:UrlEncode(value) pairs.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("intranetip:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Intranetip.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("netmask:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Netmask.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
