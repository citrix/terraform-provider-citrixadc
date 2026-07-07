package csvserver_cspolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/cs"

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

// CsvserverCspolicyBindingResourceModel describes the resource data model.
type CsvserverCspolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Name                   types.String `tfsdk:"name"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Targetlbvserver        types.String `tfsdk:"targetlbvserver"`
}

func (r *CsvserverCspolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the csvserver_cspolicy_binding resource.",
			},
			"bindpoint": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The bindpoint to which the policy is bound.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
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
			"labelname": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the label invoked.",
			},
			"labeltype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The invocation type.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the content switching virtual server to which the content switching policy applies.",
			},
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Policies bound to this vserver.",
			},
			"priority": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Priority for the policy.",
			},
			"targetlbvserver": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "target vserver name.",
			},
		},
	}
}

func csvserver_cspolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *CsvserverCspolicyBindingResourceModel) cs.Csvservercspolicybinding {
	tflog.Debug(ctx, "In csvserver_cspolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	csvserver_cspolicy_binding := cs.Csvservercspolicybinding{}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() {
		csvserver_cspolicy_binding.Bindpoint = data.Bindpoint.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		csvserver_cspolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() && !data.Invoke.IsUnknown() {
		csvserver_cspolicy_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		csvserver_cspolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() && !data.Labeltype.IsUnknown() {
		csvserver_cspolicy_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		csvserver_cspolicy_binding.Name = data.Name.ValueString()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		csvserver_cspolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		csvserver_cspolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Targetlbvserver.IsNull() && !data.Targetlbvserver.IsUnknown() {
		csvserver_cspolicy_binding.Targetlbvserver = data.Targetlbvserver.ValueString()
	}

	return csvserver_cspolicy_binding
}

// csvserver_cspolicy_bindingSetAttrFromGet is the resource-side state setter.
// It preserves the prior plan/state value for fields the GET response does not
// echo back (server-overridden / non-echoed inputs) so that Terraform does not
// see a spurious diff or "inconsistent result after apply". It does NOT recompute
// the ID (the ID is set exactly once in Create — Pattern 6).
func csvserver_cspolicy_bindingSetAttrFromGet(ctx context.Context, data *CsvserverCspolicyBindingResourceModel, getResponseData map[string]interface{}) *CsvserverCspolicyBindingResourceModel {
	tflog.Debug(ctx, "In csvserver_cspolicy_bindingSetAttrFromGet Function")

	// Convert API response to model. Only overwrite when the field is present in
	// the GET response; otherwise preserve the existing plan/state value.
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
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["targetlbvserver"]; ok && val != nil {
		data.Targetlbvserver = types.StringValue(val.(string))
	}

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// csvserver_cspolicy_bindingSetAttrFromGetForDatasource faithfully copies every
// field from the GET response (nulling fields the response omits) and sets the
// datasource ID, since the datasource has no Create to set it (Pattern 7).
func csvserver_cspolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *CsvserverCspolicyBindingResourceModel, getResponseData map[string]interface{}) *CsvserverCspolicyBindingResourceModel {
	tflog.Debug(ctx, "In csvserver_cspolicy_bindingSetAttrFromGetForDatasource Function")

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
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	} else {
		data.Policyname = types.StringNull()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		} else {
			data.Priority = types.Int64Null()
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["targetlbvserver"]; ok && val != nil {
		data.Targetlbvserver = types.StringValue(val.(string))
	} else {
		data.Targetlbvserver = types.StringNull()
	}

	// Set ID for the datasource (no Create to set it).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
