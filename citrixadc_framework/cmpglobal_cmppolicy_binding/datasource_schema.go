package cmpglobal_cmppolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CmpglobalCmppolicyBindingDataSourceModel is the data-source-specific model,
// decoupled from CmpglobalCmppolicyBindingResourceModel so the data source can
// expose read-only (GET-only) attributes the resource omits.
type CmpglobalCmppolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`

	// Read-only (GET-only) attributes from the NITRO read-only set
	// (zion73x_readonly/cmpglobal_cmppolicy_binding.json). Never settable;
	// populated from GET and null when the appliance omits them.
	Numpol types.Int64 `tfsdk:"numpol"`
}

func CmpglobalCmppolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"globalbindtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression or other value specifying the priority of the next policy, within the policy label, to evaluate if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher numbered priority.\n* END - Stop evaluation.\n* USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT.\n* An expression that evaluates to a number.\nIf you specify an expression, it's evaluation result determines the next policy to evaluate, as follows:\n* If the expression evaluates to a higher numbered priority, that policy is evaluated next.\n* If the expression evaluates to the priority of the current policy, the policy with the next higher priority number is evaluated next.\n* If the expression evaluates to a priority number that is numerically higher than the highest priority number, policy evaluation ends.\nAn UNDEF event is triggered if:\n* The expression is invalid.\n* The expression evaluates to a priority number that is numerically lower than the current policy's priority.\n* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"invoke": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Invoke policies bound to a virtual server or a policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority.",
			},
			"labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the label to invoke if the current policy rule evaluates to TRUE.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of policy label invocation.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the globally bound HTTP compression policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Positive integer specifying the priority of the policy. The lower the number, the higher the priority. By default, polices within a label are evaluated in the order of their priority numbers.\nIn the configuration utility, you can click the Priority field and edit the priority level or drag the entry to a new position in the list. If you drag the entry to a new position, the priority level is updated automatically.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Bind point to which the policy is bound.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of policies bound to the bindpoint.",
			},
		},
	}
}

// cmpglobal_cmppolicy_bindingDataSourceSetAttrFromGet projects a NITRO
// cmpglobal_cmppolicy_binding GET response onto the data-source model. The
// shared utils.MapGet* helpers fill each attribute from the GET (or leave it
// Null when the GET omits it).
func cmpglobal_cmppolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *CmpglobalCmppolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In cmpglobal_cmppolicy_bindingDataSourceSetAttrFromGet Function")

	data.Globalbindtype = utils.MapGetString(g, "globalbindtype")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Invoke = utils.MapGetBool(g, "invoke")
	data.Labelname = utils.MapGetString(g, "labelname")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Priority = utils.MapGetInt64(g, "priority")

	// Lookup keys: prefer the GET value, but preserve the configured value when
	// the appliance omits it from the binding response.
	if v := utils.MapGetString(g, "policyname"); !v.IsNull() {
		data.Policyname = v
	}
	if v := utils.MapGetString(g, "type"); !v.IsNull() {
		data.Type = v
	}

	// Read-only (GET-only) attributes.
	data.Numpol = utils.MapGetInt64(g, "numpol")

	// Composite key -> id (key:UrlEncode(value) pairs).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Type.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
