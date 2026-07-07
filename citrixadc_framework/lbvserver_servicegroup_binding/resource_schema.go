package lbvserver_servicegroup_binding

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

// LbvserverServicegroupBindingResourceModel describes the resource data model.
type LbvserverServicegroupBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Order            types.Int64  `tfsdk:"order"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
	Servicename      types.String `tfsdk:"servicename"`
	Weight           types.Int64  `tfsdk:"weight"`
}

func (r *LbvserverServicegroupBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbvserver_servicegroup_binding resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vserver\" or 'my vserver').",
			},
			"order": schema.Int64Attribute{
				// SDK v2: Optional (no server default echoed unless set). Dropping Computed
				// avoids an unresolved "unknown value" when the binding is created without
				// order and the GET response omits it.
				Optional:    true,
				Description: "Order number to be assigned to the service when it is bound to the lb vserver.",
			},
			"servicegroupname": schema.StringAttribute{
				// SDK v2 contract: Required + ForceNew. Also part of the resource identity.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The service group name bound to the selected load balancing virtual server.",
			},
			"servicename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Service to bind to the virtual server.",
			},
			"weight": schema.Int64Attribute{
				// weight is never echoed by the binding GET response; keep it Optional-only
				// (no Computed) so an unset value resolves to null instead of an unknown.
				Optional:    true,
				Description: "Integer specifying the weight of the service. A larger number specifies a greater weight. Defines the capacity of the service relative to the other services in the load balancing configuration. Determines the priority given to the service in load balancing decisions.",
			},
		},
	}
}

func lbvserver_servicegroup_bindingGetThePayloadFromthePlan(ctx context.Context, data *LbvserverServicegroupBindingResourceModel) lb.Lbvserverservicegroupbinding {
	tflog.Debug(ctx, "In lbvserver_servicegroup_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	lbvserver_servicegroup_binding := lb.Lbvserverservicegroupbinding{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		lbvserver_servicegroup_binding.Name = data.Name.ValueString()
	}
	if !data.Order.IsNull() && !data.Order.IsUnknown() {
		lbvserver_servicegroup_binding.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}
	if !data.Servicegroupname.IsNull() && !data.Servicegroupname.IsUnknown() {
		lbvserver_servicegroup_binding.Servicegroupname = data.Servicegroupname.ValueString()
	}
	if !data.Servicename.IsNull() && !data.Servicename.IsUnknown() {
		lbvserver_servicegroup_binding.Servicename = data.Servicename.ValueString()
	}
	if !data.Weight.IsNull() && !data.Weight.IsUnknown() {
		lbvserver_servicegroup_binding.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return lbvserver_servicegroup_binding
}

// lbvserver_servicegroup_bindingSetAttrFromGet is the RESOURCE-side setter. It preserves
// plan/state values for inputs the NITRO GET does not echo back (e.g. weight, and order
// when unset), so apply does not fail with "inconsistent result after apply". The ID is set
// exactly once in Create (and preserved in Update) — this function must not recompute it.
func lbvserver_servicegroup_bindingSetAttrFromGet(ctx context.Context, data *LbvserverServicegroupBindingResourceModel, getResponseData map[string]interface{}) *LbvserverServicegroupBindingResourceModel {
	tflog.Debug(ctx, "In lbvserver_servicegroup_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	// order: GET echoes it only when it was explicitly set; preserve prior value otherwise.
	if val, ok := getResponseData["order"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Order = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["servicename"]; ok && val != nil {
		data.Servicename = types.StringValue(val.(string))
	}
	// weight: not echoed by GET; preserve the prior plan/state value.
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		}
	}

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// lbvserver_servicegroup_bindingSetAttrFromGetForDatasource is the DATASOURCE-side setter.
// A datasource has no prior plan/state to preserve, so it faithfully copies every field
// from the GET response and sets the composite ID itself (no Create runs for a datasource).
func lbvserver_servicegroup_bindingSetAttrFromGetForDatasource(ctx context.Context, data *LbvserverServicegroupBindingResourceModel, getResponseData map[string]interface{}) *LbvserverServicegroupBindingResourceModel {
	tflog.Debug(ctx, "In lbvserver_servicegroup_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["order"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Order = types.Int64Value(intVal)
		} else {
			data.Order = types.Int64Null()
		}
	} else {
		data.Order = types.Int64Null()
	}
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	} else {
		data.Servicegroupname = types.StringNull()
	}
	if val, ok := getResponseData["servicename"]; ok && val != nil {
		data.Servicename = types.StringValue(val.(string))
	} else {
		data.Servicename = types.StringNull()
	}
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		} else {
			data.Weight = types.Int64Null()
		}
	} else {
		data.Weight = types.Int64Null()
	}

	// Composite ID: legacy order "name,servicegroupname".
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
