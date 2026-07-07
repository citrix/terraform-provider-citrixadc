package lbvserver_dnspolicy64_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LbvserverDnspolicy64BindingResourceModel describes the resource data model.
type LbvserverDnspolicy64BindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Name                   types.String `tfsdk:"name"`
	Order                  types.Int64  `tfsdk:"order"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
}

func (r *LbvserverDnspolicy64BindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbvserver_dnspolicy64_binding resource.",
			},
			"bindpoint": schema.StringAttribute{
				// NITRO GET does not echo bindpoint; Optional-only so it is null
				// (not unknown) after apply when unset. (Pattern 13)
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bind point to which to bind the policy.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"invoke": schema.BoolAttribute{
				// NITRO GET does not echo invoke; Optional-only. (Pattern 13)
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Invoke policies bound to a virtual server or policy label.",
			},
			"labelname": schema.StringAttribute{
				// NITRO GET does not echo labelname; Optional-only. (Pattern 13)
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE.",
			},
			"labeltype": schema.StringAttribute{
				// NITRO GET does not echo labeltype; Optional-only. (Pattern 13)
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of policy label to invoke. Applicable only to rewrite, videooptimization and cache policies. Available settings function as follows:\n* reqvserver - Evaluate the request against the request-based policies bound to the specified virtual server.\n* resvserver - Evaluate the response against the response-based policies bound to the specified virtual server.\n* policylabel - invoke the request or response against the specified user-defined policy label.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vserver\" or 'my vserver').",
			},
			"order": schema.Int64Attribute{
				// NITRO GET does not echo order; Optional-only. (Pattern 13)
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing vserver's bindings. Determines the priority given to the service among all the services bound.",
			},
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the policy bound to the LB vserver.",
			},
			"priority": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
					int64planmodifier.UseStateForUnknown(),
				},
				Description: "Priority.",
			},
		},
	}
}

func lbvserver_dnspolicy64_bindingGetThePayloadFromthePlan(ctx context.Context, data *LbvserverDnspolicy64BindingResourceModel) lb.Lbvserverdnspolicy64binding {
	tflog.Debug(ctx, "In lbvserver_dnspolicy64_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	lbvserver_dnspolicy64_binding := lb.Lbvserverdnspolicy64binding{}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() {
		lbvserver_dnspolicy64_binding.Bindpoint = data.Bindpoint.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		lbvserver_dnspolicy64_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() && !data.Invoke.IsUnknown() {
		lbvserver_dnspolicy64_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		lbvserver_dnspolicy64_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() && !data.Labeltype.IsUnknown() {
		lbvserver_dnspolicy64_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		lbvserver_dnspolicy64_binding.Name = data.Name.ValueString()
	}
	if !data.Order.IsNull() && !data.Order.IsUnknown() {
		lbvserver_dnspolicy64_binding.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		lbvserver_dnspolicy64_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		lbvserver_dnspolicy64_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return lbvserver_dnspolicy64_binding
}

// lbvserver_dnspolicy64_bindingSetAttrFromGet is the resource-side state setter.
// It preserves prior plan/state values for inputs the NITRO GET response does not
// echo back (or echoes back in a normalized form), to avoid "inconsistent result
// after apply" diffs. Identity attributes (name, policyname) are taken from the
// GET response. (Pattern 7 / Pattern 13)
func lbvserver_dnspolicy64_bindingSetAttrFromGet(ctx context.Context, data *LbvserverDnspolicy64BindingResourceModel, getResponseData map[string]interface{}) *LbvserverDnspolicy64BindingResourceModel {
	tflog.Debug(ctx, "In lbvserver_dnspolicy64_bindingSetAttrFromGet Function")

	// Identity attributes are always present in the GET response.
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	}

	// The following attributes are write-time inputs that NITRO does not reliably
	// echo back (or echoes in a server-overridden form). Only adopt the GET value
	// when present; otherwise preserve the existing plan/state value so Terraform
	// does not see a spurious change.
	if val, ok := getResponseData["bindpoint"]; ok && val != nil {
		data.Bindpoint = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["invoke"]; ok && val != nil {
		data.Invoke = types.BoolValue(val.(bool))
	}
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["labeltype"]; ok && val != nil {
		data.Labeltype = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["order"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Order = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	}

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// lbvserver_dnspolicy64_bindingSetAttrFromGetForDatasource faithfully copies every
// field from the GET response (the datasource has no prior plan/state to preserve)
// and sets the composite ID. (Pattern 7 datasource split)
func lbvserver_dnspolicy64_bindingSetAttrFromGetForDatasource(ctx context.Context, data *LbvserverDnspolicy64BindingResourceModel, getResponseData map[string]interface{}) *LbvserverDnspolicy64BindingResourceModel {
	tflog.Debug(ctx, "In lbvserver_dnspolicy64_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["bindpoint"]; ok && val != nil {
		data.Bindpoint = types.StringValue(val.(string))
	} else {
		data.Bindpoint = types.StringNull()
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["invoke"]; ok && val != nil {
		data.Invoke = types.BoolValue(val.(bool))
	} else {
		data.Invoke = types.BoolNull()
	}
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	} else {
		data.Labelname = types.StringNull()
	}
	if val, ok := getResponseData["labeltype"]; ok && val != nil {
		data.Labeltype = types.StringValue(val.(string))
	} else {
		data.Labeltype = types.StringNull()
	}
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
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	} else {
		data.Policyname = types.StringNull()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else {
		data.Priority = types.Int64Null()
	}

	// Set ID for the datasource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
