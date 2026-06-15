package cachepolicylabel_cachepolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/cache"

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

// CachepolicylabelCachepolicyBindingResourceModel describes the resource data model.
type CachepolicylabelCachepolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	InvokeLabelname        types.String `tfsdk:"invoke_labelname"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
}

func (r *CachepolicylabelCachepolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cachepolicylabel_cachepolicy_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"invoke": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Invoke policies bound to a virtual server or a user-defined policy label. After the invoked policies are evaluated, the flow returns to the policy with the next-lower priority.",
			},
			"invoke_labelname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Name of the policy label to invoke if the current policy rule evaluates to TRUE.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the cache policy label to which to bind the policy.",
			},
			"labeltype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Type of policy label to invoke: an unnamed label associated with a virtual server, or user-defined policy label.",
			},
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the cache policy to bind to the policy label.",
			},
			"priority": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specifies the priority of the policy.",
			},
		},
	}
}

func cachepolicylabel_cachepolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *CachepolicylabelCachepolicyBindingResourceModel) cache.Cachepolicylabelcachepolicybinding {
	tflog.Debug(ctx, "In cachepolicylabel_cachepolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	cachepolicylabel_cachepolicy_binding := cache.Cachepolicylabelcachepolicybinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		cachepolicylabel_cachepolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() && !data.Invoke.IsUnknown() {
		cachepolicylabel_cachepolicy_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.InvokeLabelname.IsNull() && !data.InvokeLabelname.IsUnknown() {
		cachepolicylabel_cachepolicy_binding.Invokelabelname = data.InvokeLabelname.ValueString()
	}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		cachepolicylabel_cachepolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() && !data.Labeltype.IsUnknown() {
		cachepolicylabel_cachepolicy_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		cachepolicylabel_cachepolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		cachepolicylabel_cachepolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return cachepolicylabel_cachepolicy_binding
}

// cachepolicylabel_cachepolicy_bindingSetAttrFromGet is the RESOURCE-side state setter.
// All non-identity attributes (gotopriorityexpression, invoke, invoke_labelname,
// labeltype, priority) are RequiresReplace inputs that the NITRO GET response may not
// echo back, or may return server-normalized values for. To avoid "inconsistent result
// after apply"/perpetual diffs, the existing plan/state value is preserved when the GET
// response omits the field; it is only adopted from GET when not already set in state.
// Identity attributes (labelname, policyname) are always adopted from the GET response.
// The ID is NOT recomputed here - it is set once in Create (and in the datasource setter).
func cachepolicylabel_cachepolicy_bindingSetAttrFromGet(ctx context.Context, data *CachepolicylabelCachepolicyBindingResourceModel, getResponseData map[string]interface{}) *CachepolicylabelCachepolicyBindingResourceModel {
	tflog.Debug(ctx, "In cachepolicylabel_cachepolicy_bindingSetAttrFromGet Function")

	// Identity attributes - always adopt from GET (also makes imported state usable)
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	}

	// Server-overridden / non-echoed RequiresReplace inputs:
	// adopt from GET only when present; otherwise preserve the existing plan/state value.
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["invoke"]; ok && val != nil {
		data.Invoke = types.BoolValue(val.(bool))
	}
	// invoke_labelname and labeltype are Optional+Computed but are NOT echoed by the
	// NITRO GET response. Adopt from GET when present; otherwise keep the existing
	// known state value, and fall back to a (known) null so the Computed attribute is
	// never left unknown after apply.
	if val, ok := getResponseData["invoke_labelname"]; ok && val != nil {
		data.InvokeLabelname = types.StringValue(val.(string))
	} else if data.InvokeLabelname.IsUnknown() {
		data.InvokeLabelname = types.StringNull()
	}
	if val, ok := getResponseData["labeltype"]; ok && val != nil {
		data.Labeltype = types.StringValue(val.(string))
	} else if data.Labeltype.IsUnknown() {
		data.Labeltype = types.StringNull()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	}

	return data
}

// cachepolicylabel_cachepolicy_bindingSetAttrFromGetForDatasource is the DATASOURCE-side
// state setter. The datasource has no prior plan/state to preserve, so it faithfully
// copies every field from the GET response and sets the composite ID itself.
func cachepolicylabel_cachepolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *CachepolicylabelCachepolicyBindingResourceModel, getResponseData map[string]interface{}) *CachepolicylabelCachepolicyBindingResourceModel {
	tflog.Debug(ctx, "In cachepolicylabel_cachepolicy_bindingSetAttrFromGetForDatasource Function")

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
	if val, ok := getResponseData["invoke_labelname"]; ok && val != nil {
		data.InvokeLabelname = types.StringValue(val.(string))
	} else {
		data.InvokeLabelname = types.StringNull()
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
	idParts = append(idParts, fmt.Sprintf("labelname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Labelname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
