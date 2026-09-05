package transformglobal_transformpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// TransformglobalTransformpolicyBindingDataSourceModel is the data-source-specific
// model, decoupled from the resource model. A data source is a pure read surface
// (Read only; no plan/apply lifecycle), so it can expose the full GET projection:
// the read/write attributes (as Computed outputs) AND the read-only attributes
// that the resource deliberately omits (flowtype, numpol).
type TransformglobalTransformpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/transformglobal_transformpolicy_binding.json).
	Flowtype types.Int64 `tfsdk:"flowtype"`
	Numpol   types.Int64 `tfsdk:"numpol"`
}

func TransformglobalTransformpolicyBindingDataSourceSchema() schema.Schema {
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
				Description: "If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forwards the request or response to the specified virtual server or evaluates the specified policy label.",
			},
			"labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and the label type is Policy Label.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of invocation. Available settings function as follows:\n* reqvserver - Send the request to the specified request virtual server.\n* resvserver - Send the response to the specified response virtual server.\n* policylabel - Invoke the specified policy label.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the transform policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Specifies the bind point to which to bind the policy. Available settings function as follows:\n* REQ_OVERRIDE. Request override. Binds the policy to the priority request queue.\n* REQ_DEFAULT. Binds the policy to the default request queue.\n* HTTPQUIC_REQ_OVERRIDE - Binds the policy to the HTTP_QUIC override request queue.\n* HTTPQUIC_REQ_DEFAULT - Binds the policy to the HTTP_QUIC default request queue.",
			},

			// Read-only (GET-only) metadata surfaced by the data source. All Computed.
			"flowtype": schema.Int64Attribute{
				Computed:    true,
				Description: "flowtype of the bound transform policy.",
			},
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of policies bound to the bindpoint.",
			},
		},
	}
}

// transformglobal_transformpolicy_bindingDataSourceSetAttrFromGet projects a NITRO
// GET response onto the data-source model using the shared utils.MapGet* helpers.
func transformglobal_transformpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *TransformglobalTransformpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In transformglobal_transformpolicy_bindingDataSourceSetAttrFromGet Function")

	data.Globalbindtype = utils.MapGetString(g, "globalbindtype")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Invoke = utils.MapGetBool(g, "invoke")
	data.Labelname = utils.MapGetString(g, "labelname")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Policyname = utils.MapGetString(g, "policyname")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Type = utils.MapGetString(g, "type")

	// Read-only (GET-only) metadata.
	data.Flowtype = utils.MapGetInt64(g, "flowtype")
	data.Numpol = utils.MapGetInt64(g, "numpol")

	// Set composite ID (policyname,type) preserving the resource ID format.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(data.Policyname.ValueString())))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(data.Type.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
