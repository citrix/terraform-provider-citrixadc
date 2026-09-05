package lbglobal_lbpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LbglobalLbpolicyBindingDataSourceModel is the data-source-specific model,
// decoupled from LbglobalLbpolicyBindingResourceModel. A data source is a pure
// read surface, so it can expose the FULL GET projection: the configurable
// attributes (as Computed outputs) AND the read-only metadata attributes the
// resource deliberately omits.
type LbglobalLbpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"` // Required lookup key
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"` // Required lookup key

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/lbglobal_lbpolicy_binding.json).
	Flowtype types.Int64 `tfsdk:"flowtype"`
	Numpol   types.Int64 `tfsdk:"numpol"`
}

func LbglobalLbpolicyBindingDataSourceSchema() schema.Schema {
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
				Description: "Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of invocation, Available settings function as follows:\n* vserver - Invokes the unnamed policy label associated with the specified virtual server.\n* policylabel - Invoke the specified policy label.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the LB policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "0",
			},

			// Read-only (GET-only) metadata surfaced by the data source. All Computed.
			"flowtype": schema.Int64Attribute{
				Computed:    true,
				Description: "flowtype of the bound LB policy.",
			},
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "number of polices bound to label.",
			},
		},
	}
}

// lbglobal_lbpolicy_bindingDataSourceSetAttrFromGet projects a NITRO
// lbglobal_lbpolicy_binding GET response onto the data-source model.
func lbglobal_lbpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *LbglobalLbpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In lbglobal_lbpolicy_bindingDataSourceSetAttrFromGet Function")

	// Preserve the config-provided lookup keys when the GET omits them.
	policyname := data.Policyname
	typ := data.Type

	if v, ok := g["policyname"]; ok && v != nil {
		data.Policyname = types.StringValue(utils.AnyToString(v))
	} else {
		data.Policyname = policyname
	}
	if v, ok := g["type"]; ok && v != nil {
		data.Type = types.StringValue(utils.AnyToString(v))
	} else {
		data.Type = typ
	}

	data.Globalbindtype = utils.MapGetString(g, "globalbindtype")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Invoke = utils.MapGetBool(g, "invoke")
	data.Labelname = utils.MapGetString(g, "labelname")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Priority = utils.MapGetInt64(g, "priority")

	// Read-only (GET-only) metadata.
	data.Flowtype = utils.MapGetInt64(g, "flowtype")
	data.Numpol = utils.MapGetInt64(g, "numpol")

	// Set the composite ID (policyname:<v>,type:<v>).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(data.Policyname.ValueString())))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(data.Type.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
