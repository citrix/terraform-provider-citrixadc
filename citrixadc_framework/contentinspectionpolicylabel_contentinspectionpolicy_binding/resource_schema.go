package contentinspectionpolicylabel_contentinspectionpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"

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

// ContentinspectionpolicylabelContentinspectionpolicyBindingResourceModel describes the resource data model.
type ContentinspectionpolicylabelContentinspectionpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	InvokeLabelname        types.String `tfsdk:"invokelabelname"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
}

func (r *ContentinspectionpolicylabelContentinspectionpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the contentinspectionpolicylabel_contentinspectionpolicy_binding resource.",
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
				Description: "Suspend evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.",
			},
			"invokelabelname": schema.StringAttribute{
				// SDK v2 contract: Optional (ForceNew). The NITRO GET does NOT echo
				// this field back, so it is NOT Computed (Computed would force an
				// unknown-after-apply value that GET can never resolve -> Pattern 13).
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "* If labelType is policylabel, name of the policy label to invoke.\n* If labelType is reqvserver or resvserver, name of the virtual server to which to forward the request or response.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the contentInspection policy label to which to bind the policy.",
			},
			"labeltype": schema.StringAttribute{
				// SDK v2 contract: Optional (ForceNew). The NITRO GET does NOT echo
				// this field back, so it is NOT Computed (Pattern 13).
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of invocation. Available settings function as follows:\n* reqvserver - Forward the request to the specified request virtual server.\n* resvserver - Forward the response to the specified response virtual server.\n* policylabel - Invoke the specified policy label.",
			},
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the contentInspection policy to bind to the policy label.",
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

func contentinspectionpolicylabel_contentinspectionpolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *ContentinspectionpolicylabelContentinspectionpolicyBindingResourceModel) contentinspection.Contentinspectionpolicylabelcontentinspectionpolicybinding {
	tflog.Debug(ctx, "In contentinspectionpolicylabel_contentinspectionpolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	contentinspectionpolicylabel_contentinspectionpolicy_binding := contentinspection.Contentinspectionpolicylabelcontentinspectionpolicybinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		contentinspectionpolicylabel_contentinspectionpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() && !data.Invoke.IsUnknown() {
		contentinspectionpolicylabel_contentinspectionpolicy_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.InvokeLabelname.IsNull() && !data.InvokeLabelname.IsUnknown() {
		contentinspectionpolicylabel_contentinspectionpolicy_binding.Invokelabelname = data.InvokeLabelname.ValueString()
	}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		contentinspectionpolicylabel_contentinspectionpolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() && !data.Labeltype.IsUnknown() {
		contentinspectionpolicylabel_contentinspectionpolicy_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		contentinspectionpolicylabel_contentinspectionpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		contentinspectionpolicylabel_contentinspectionpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return contentinspectionpolicylabel_contentinspectionpolicy_binding
}

// contentinspectionpolicylabel_contentinspectionpolicy_bindingSetAttrFromGet is
// the RESOURCE-side state setter. The NITRO GET for this binding echoes back
// labelname, policyname, priority, gotopriorityexpression and invoke (the
// server may override gotopriorityexpression to "END" — adopt it). It does NOT
// echo invokelabelname (invokelabelname) or labeltype, so for those we PRESERVE
// the existing plan/state value rather than nulling it (Pattern 7 / pattern (e)).
// The ID is NOT recomputed here — Create sets it once (Pattern 6).
func contentinspectionpolicylabel_contentinspectionpolicy_bindingSetAttrFromGet(ctx context.Context, data *ContentinspectionpolicylabelContentinspectionpolicyBindingResourceModel, getResponseData map[string]interface{}) *ContentinspectionpolicylabelContentinspectionpolicyBindingResourceModel {
	tflog.Debug(ctx, "In contentinspectionpolicylabel_contentinspectionpolicy_bindingSetAttrFromGet Function")

	// Echoed by GET — safe to adopt (gotopriorityexpression is server-overridden, e.g. "END").
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["invoke"]; ok && val != nil {
		data.Invoke = types.BoolValue(val.(bool))
	}
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	}

	// NOT echoed by the NITRO GET (invokelabelname, labeltype) — preserve the
	// configured plan/state value rather than nulling it.

	return data
}

// contentinspectionpolicylabel_contentinspectionpolicy_bindingSetAttrFromGetForDatasource
// is the DATASOURCE-side setter. A datasource has no prior plan/state to
// preserve, so it faithfully copies every field the GET response returns and
// sets data.Id (Pattern 7 datasource split).
func contentinspectionpolicylabel_contentinspectionpolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *ContentinspectionpolicylabelContentinspectionpolicyBindingResourceModel, getResponseData map[string]interface{}) *ContentinspectionpolicylabelContentinspectionpolicyBindingResourceModel {
	tflog.Debug(ctx, "In contentinspectionpolicylabel_contentinspectionpolicy_bindingSetAttrFromGetForDatasource Function")

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

	// Set ID for the datasource (matches the resource Create ID format: labelname,policyname)
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("labelname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Labelname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
