package lbvserver_service_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LbvserverServiceBindingResourceModel describes the resource data model.
type LbvserverServiceBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Order            types.Int64  `tfsdk:"order"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
	Servicename      types.String `tfsdk:"servicename"`
	Weight           types.Int64  `tfsdk:"weight"`
}

func (r *LbvserverServiceBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbvserver_service_binding resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vserver\" or 'my vserver').",
			},
			"order": schema.Int64Attribute{
				// Not echoed by the GET response for a service binding (only "orderstr"
				// is returned), so it cannot be Computed without perpetual
				// unknown-after-apply errors (Pattern 13). NITRO cannot update a binding
				// in place (errorcode 273 on re-bind), so a change forces replace,
				// matching the SDK v2 ForceNew contract.
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Order number to be assigned to the service when it is bound to the lb vserver.",
			},
			"servicegroupname": schema.StringAttribute{
				// Not echoed by GET for a service binding; drop Computed (Pattern 13).
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the service group.",
			},
			"servicename": schema.StringAttribute{
				// SDK v2 ForceNew: a binding cannot be re-bound in place (errorcode 273),
				// so a servicename change forces replace.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Service to bind to the virtual server.",
			},
			"weight": schema.Int64Attribute{
				// SDK v2 ForceNew: a binding cannot be updated in place (errorcode 273
				// on re-bind), so a weight change forces replace (unbind + rebind).
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Weight to assign to the specified service.",
			},
		},
	}
}

func lbvserver_service_bindingGetThePayloadFromthePlan(ctx context.Context, data *LbvserverServiceBindingResourceModel) lb.Lbvserverservicebinding {
	tflog.Debug(ctx, "In lbvserver_service_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	lbvserver_service_binding := lb.Lbvserverservicebinding{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		lbvserver_service_binding.Name = data.Name.ValueString()
	}
	if !data.Order.IsNull() && !data.Order.IsUnknown() {
		lbvserver_service_binding.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}
	if !data.Servicegroupname.IsNull() && !data.Servicegroupname.IsUnknown() {
		lbvserver_service_binding.Servicegroupname = data.Servicegroupname.ValueString()
	}
	if !data.Servicename.IsNull() && !data.Servicename.IsUnknown() {
		lbvserver_service_binding.Servicename = data.Servicename.ValueString()
	}
	if !data.Weight.IsNull() && !data.Weight.IsUnknown() {
		lbvserver_service_binding.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return lbvserver_service_binding
}

// lbvserver_service_bindingSetAttrFromGet is the RESOURCE-side setter. It preserves
// the configured plan/state values for non-echoed inputs (Pattern 7) and does NOT
// recompute the ID (the ID is set exactly once in Create — Pattern 6).
func lbvserver_service_bindingSetAttrFromGet(ctx context.Context, data *LbvserverServiceBindingResourceModel, getResponseData map[string]interface{}) *LbvserverServiceBindingResourceModel {
	tflog.Debug(ctx, "In lbvserver_service_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["order"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Order = types.Int64Value(intVal)
		}
	}
	// servicegroupname is not echoed by GET for a service binding; preserve the
	// configured plan/state value rather than nulling it (Pattern 7).
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["servicename"]; ok && val != nil {
		data.Servicename = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		}
	}

	// NOTE: ID is composed once in Create (and preserved in Update/Read); do not
	// recompute it here.
	return data
}

// lbvserver_service_bindingSetAttrFromGetForDatasource is the DATASOURCE-side setter.
// It faithfully copies every field from the GET response and sets the ID, since a
// datasource has no prior plan/state and never calls Create (Pattern 7 split).
func lbvserver_service_bindingSetAttrFromGetForDatasource(ctx context.Context, data *LbvserverServiceBindingResourceModel, getResponseData map[string]interface{}) *LbvserverServiceBindingResourceModel {
	tflog.Debug(ctx, "In lbvserver_service_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["order"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Order = types.Int64Value(intVal)
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
		}
	} else {
		data.Weight = types.Int64Null()
	}

	// Set ID for the datasource: identity is "name,servicename".
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicename.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
