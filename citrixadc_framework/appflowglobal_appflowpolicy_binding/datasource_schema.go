package appflowglobal_appflowpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppflowglobalAppflowpolicyBindingDataSourceModel is the data-source-specific
// model. It carries every attribute the data source already exposed PLUS the
// read-only (GET-only) attributes from the NITRO doc read-only set
// (zion73x_readonly/appflowglobal_appflowpolicy_binding.json). All non-key
// attributes are Computed; the read-only attributes are never settable and are
// populated from the GET response (Null when the appliance omits them).
type AppflowglobalAppflowpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`

	// Read-only (GET-only) attributes surfaced by the data source.
	Flowtype types.Int64 `tfsdk:"flowtype"`
	Numpol   types.Int64 `tfsdk:"numpol"`
}

func AppflowglobalAppflowpolicyBindingDataSourceSchema() schema.Schema {
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
				Description: "Invoke policies bound to a virtual server or a user-defined policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority.",
			},
			"labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the label to invoke if the current policy evaluates to TRUE.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of policy label to invoke. Specify vserver for a policy label associated with a virtual server, or policylabel for a user-defined policy label.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the AppFlow policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Global bind point for which to show detailed information about the policies bound to the bind point.",
			},

			// Read-only (GET-only) attributes. Computed; null when the
			// appliance omits them.
			"flowtype": schema.Int64Attribute{
				Computed:    true,
				Description: "Flow type of the bound AppFlow policy.",
			},
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of policies bound to the bindpoint.",
			},
		},
	}
}

// appflowglobal_appflowpolicy_bindingDataSourceSetAttrFromGet projects a NITRO
// appflowglobal_appflowpolicy_binding GET response onto the data-source model. A
// data source has no plan/apply reconciliation, so attributes are simply filled
// from the GET (or left Null when the GET omits them) via the shared
// utils.MapGet* helpers.
func appflowglobal_appflowpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *AppflowglobalAppflowpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appflowglobal_appflowpolicy_bindingDataSourceSetAttrFromGet Function")

	data.Globalbindtype = utils.MapGetString(g, "globalbindtype")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Invoke = utils.MapGetBool(g, "invoke")
	data.Labelname = utils.MapGetString(g, "labelname")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Policyname = utils.MapGetString(g, "policyname")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Type = utils.MapGetString(g, "type")

	// Read-only (GET-only) attributes.
	data.Flowtype = utils.MapGetInt64(g, "flowtype")
	data.Numpol = utils.MapGetInt64(g, "numpol")

	// Set the composite ID (policyname:<v>,type:<v>). A single appflowpolicy can be
	// bound at multiple bind points (type) simultaneously, so policyname alone is not
	// a unique key.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(data.Policyname.ValueString())))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(data.Type.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
