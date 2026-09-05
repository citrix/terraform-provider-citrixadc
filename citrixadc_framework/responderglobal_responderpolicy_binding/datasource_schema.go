package responderglobal_responderpolicy_binding

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ResponderglobalResponderpolicyBindingDataSourceModel is the data-source-specific
// model, decoupled from the resource model. A data source is a pure read surface
// (Read only; no plan/apply lifecycle), so it can expose the FULL GET projection:
// the read/write attributes (as Computed outputs) AND the read-only attributes
// that the resource deliberately omits. Every non-key attribute is Computed.
type ResponderglobalResponderpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/responderglobal_responderpolicy_binding.json). Never
	// settable; populated from GET, Null when the appliance omits them.
	Numpol   types.Int64 `tfsdk:"numpol"`
	Flowtype types.Int64 `tfsdk:"flowtype"`
}

func ResponderglobalResponderpolicyBindingDataSourceSchema() schema.Schema {
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
				Description: "Name of the policy label to invoke. If the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is policylabel.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of invocation, Available settings function as follows:\n* vserver - Forward the request to the specified virtual server.\n* policylabel - Invoke the specified policy label.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the responder policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Specifies the bind point whose policies you want to display. Available settings function as follows:\n* REQ_OVERRIDE - Request override. Binds the policy to the priority request queue.\n* REQ_DEFAULT - Binds the policy to the default request queue.\n* OTHERTCP_REQ_OVERRIDE - Binds the policy to the non-HTTP TCP priority request queue.\n* OTHERTCP_REQ_DEFAULT - Binds the policy to the non-HTTP TCP default request queue..\n* SIPUDP_REQ_OVERRIDE - Binds the policy to the SIP UDP priority response queue..\n* SIPUDP_REQ_DEFAULT - Binds the policy to the SIP UDP default response queue.\n* RADIUS_REQ_OVERRIDE - Binds the policy to the RADIUS priority response queue..\n* RADIUS_REQ_DEFAULT - Binds the policy to the RADIUS default response queue.\n* MSSQL_REQ_OVERRIDE - Binds the policy to the Microsoft SQL priority response queue..\n* MSSQL_REQ_DEFAULT - Binds the policy to the Microsoft SQL default response queue.\n* MYSQL_REQ_OVERRIDE - Binds the policy to the MySQL priority response queue.\n* MYSQL_REQ_DEFAULT - Binds the policy to the MySQL default response queue.\n* HTTPQUIC_REQ_OVERRIDE - Binds the policy to the HTTP_QUIC override response queue.\n* HTTPQUIC_REQ_DEFAULT - Binds the policy to the HTTP_QUIC default response queue.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of policies bound to label.",
			},
			"flowtype": schema.Int64Attribute{
				Computed:    true,
				Description: "Flowtype of the bound responder policy.",
			},
		},
	}
}

// responderglobal_responderpolicy_bindingDataSourceSetAttrFromGet projects a NITRO
// responderglobal_responderpolicy_binding GET response onto the data-source model.
func responderglobal_responderpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *ResponderglobalResponderpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In responderglobal_responderpolicy_bindingDataSourceSetAttrFromGet Function")

	// Set ID/key. Backward-compatible with SDK v2: the legacy resource used
	// d.SetId(policyname).
	if v, ok := g["policyname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Policyname = types.StringValue(utils.AnyToString(v))
	}

	data.Globalbindtype = utils.MapGetString(g, "globalbindtype")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Invoke = utils.MapGetBool(g, "invoke")
	data.Labelname = utils.MapGetString(g, "labelname")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Type = utils.MapGetString(g, "type")

	// Read-only (GET-only) attributes.
	data.Numpol = utils.MapGetInt64(g, "numpol")
	data.Flowtype = utils.MapGetInt64(g, "flowtype")
}
