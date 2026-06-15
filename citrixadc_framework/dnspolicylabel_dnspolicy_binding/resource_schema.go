package dnspolicylabel_dnspolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

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

// DnspolicylabelDnspolicyBindingResourceModel describes the resource data model.
type DnspolicylabelDnspolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	InvokeLabelname        types.String `tfsdk:"invokelabelname"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
}

func (r *DnspolicylabelDnspolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnspolicylabel_dnspolicy_binding resource.",
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
				Description: "Invoke flag.",
			},
			"invokelabelname": schema.StringAttribute{
				// Optional only (no Computed): the binding GET never echoes this field,
				// so a Computed value can never be resolved after apply (Pattern 13).
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the label to invoke if the current policy rule evaluates to TRUE.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the dns policy label.",
			},
			"labeltype": schema.StringAttribute{
				// Optional only (no Computed): the binding GET never echoes this field,
				// so a Computed value can never be resolved after apply (Pattern 13).
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of policy label invocation.",
			},
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The dns policy name.",
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

func dnspolicylabel_dnspolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *DnspolicylabelDnspolicyBindingResourceModel) dns.Dnspolicylabeldnspolicybinding {
	tflog.Debug(ctx, "In dnspolicylabel_dnspolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	dnspolicylabel_dnspolicy_binding := dns.Dnspolicylabeldnspolicybinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		dnspolicylabel_dnspolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() && !data.Invoke.IsUnknown() {
		dnspolicylabel_dnspolicy_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.InvokeLabelname.IsNull() && !data.InvokeLabelname.IsUnknown() {
		dnspolicylabel_dnspolicy_binding.Invokelabelname = data.InvokeLabelname.ValueString()
	}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		dnspolicylabel_dnspolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() && !data.Labeltype.IsUnknown() {
		dnspolicylabel_dnspolicy_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		dnspolicylabel_dnspolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		dnspolicylabel_dnspolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return dnspolicylabel_dnspolicy_binding
}

// dnspolicylabel_dnspolicy_bindingSetAttrFromGet is the resource-side state setter.
// All attributes except the unique keys (labelname, policyname) are RequiresReplace and
// several (gotopriorityexpression, invoke, invokelabelname, labeltype) are server-derived
// / not echoed back for a binding GET. Overwriting them from a sparse GET response would
// either wipe the user's configured values or set "inconsistent result after apply"
// errors (Pattern 7 / Pattern 13). So preserve the existing plan/state values; only
// adopt a GET value when it is present, otherwise leave the model field untouched.
func dnspolicylabel_dnspolicy_bindingSetAttrFromGet(ctx context.Context, data *DnspolicylabelDnspolicyBindingResourceModel, getResponseData map[string]interface{}) *DnspolicylabelDnspolicyBindingResourceModel {
	tflog.Debug(ctx, "In dnspolicylabel_dnspolicy_bindingSetAttrFromGet Function")

	// Preserve user-supplied / server-overridden inputs. Only adopt a GET value when
	// it is present and non-nil; never null out an existing value.
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["invoke"]; ok && val != nil {
		data.Invoke = types.BoolValue(val.(bool))
	}
	if val, ok := getResponseData["invoke_labelname"]; ok && val != nil {
		data.InvokeLabelname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["labeltype"]; ok && val != nil {
		data.Labeltype = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	}

	// ID is set once in Create; do not recompute here.
	return data
}

// dnspolicylabel_dnspolicy_bindingSetAttrFromGetForDatasource faithfully copies every
// field from the GET response (the datasource has no prior plan/state to preserve) and
// sets the composite ID (Pattern 7 datasource split).
func dnspolicylabel_dnspolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *DnspolicylabelDnspolicyBindingResourceModel, getResponseData map[string]interface{}) *DnspolicylabelDnspolicyBindingResourceModel {
	tflog.Debug(ctx, "In dnspolicylabel_dnspolicy_bindingSetAttrFromGetForDatasource Function")

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

	// Set ID for the datasource (no Create path).
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("labelname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Labelname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
