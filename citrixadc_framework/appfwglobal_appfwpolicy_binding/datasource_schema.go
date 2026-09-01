package appfwglobal_appfwpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwglobalAppfwpolicyBindingDataSourceModel is the data-source-specific model,
// decoupled from AppfwglobalAppfwpolicyBindingResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits
// (flowtype, numpol, policytype). Every non-key attribute is Computed; the
// Framework's per-attribute model <-> schema reflection requires this model to
// have exactly the attributes the data-source schema declares.
type AppfwglobalAppfwpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	State                  types.String `tfsdk:"state"`
	Type                   types.String `tfsdk:"type"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/appfwglobal_appfwpolicy_binding.json). Never settable;
	// populated from GET.
	Flowtype   types.Int64  `tfsdk:"flowtype"`
	Numpol     types.Int64  `tfsdk:"numpol"`
	Policytype types.String `tfsdk:"policytype"`
}

func AppfwglobalAppfwpolicyBindingDataSourceSchema() schema.Schema {
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
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"invoke": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.",
			},
			"labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is set to Policy Label.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of policy label invocation.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The priority of the policy.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the binding to activate or deactivate the policy. This is applicable to classic policies only.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Bind point to which to policy is bound.",
			},

			// Read-only (GET-only) attributes surfaced by the data source (these
			// are intentionally NOT modeled on the resource). All Computed.
			"flowtype": schema.Int64Attribute{
				Computed:    true,
				Description: "flowtype of the bound application firewall policy.",
			},
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of policies bound to the bindpoint.",
			},
			"policytype": schema.StringAttribute{
				Computed:    true,
				Description: "Policy type. Possible values: [ Classic Policy, Advanced Policy ]",
			},
		},
	}
}

// appfwglobal_appfwpolicy_bindingDataSourceSetAttrFromGet projects a NITRO
// appfwglobal_appfwpolicy_binding GET response onto the data-source model. A data
// source has no plan/apply reconciliation, so attributes are simply filled from
// the GET (or left Null when the GET omits them) via the shared utils.MapGet*
// helpers. It faithfully copies every field (including 'state') and sets the
// composite ID.
func appfwglobal_appfwpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *AppfwglobalAppfwpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appfwglobal_appfwpolicy_bindingDataSourceSetAttrFromGet Function")

	data.Globalbindtype = utils.MapGetString(g, "globalbindtype")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Invoke = utils.MapGetBool(g, "invoke")
	data.Labelname = utils.MapGetString(g, "labelname")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Policyname = utils.MapGetString(g, "policyname")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.State = utils.MapGetString(g, "state")
	data.Type = utils.MapGetString(g, "type")

	// Read-only (GET-only) attributes.
	data.Flowtype = utils.MapGetInt64(g, "flowtype")
	data.Numpol = utils.MapGetInt64(g, "numpol")
	data.Policytype = utils.MapGetString(g, "policytype")

	// Set composite ID for the datasource.
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Type.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
