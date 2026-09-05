package lbpolicylabel

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LbpolicylabelDataSourceModel is the data-source-specific model, decoupled from
// LbpolicylabelResourceModel. A data source is a pure read surface (Read only),
// so it exposes the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits (numpol,
// hits, gotopriorityexpression, labeltype, invoke_labelname).
type LbpolicylabelDataSourceModel struct {
	Id              types.String `tfsdk:"id"`
	Comment         types.String `tfsdk:"comment"`
	Labelname       types.String `tfsdk:"labelname"` // Required lookup key
	Newname         types.String `tfsdk:"newname"`
	Policylabeltype types.String `tfsdk:"policylabeltype"`

	// Read-only (GET-only) metadata from the NITRO read-only set
	// (zion73x_readonly/lbpolicylabel.json). Never settable; populated from GET.
	Numpol                 types.Int64  `tfsdk:"numpol"`
	Hits                   types.Int64  `tfsdk:"hits"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Labeltype              types.String `tfsdk:"labeltype"`
	InvokeLabelname        types.String `tfsdk:"invoke_labelname"`
}

func LbpolicylabelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this LB policy label.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LB policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my lb policy label\" or 'my lb policy label').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the LB policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"policylabeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocols supported by the policylabel. Available Types are :\n* HTTP - HTTP requests.\n* DNS - DNS request.\n* OTHERTCP - OTHERTCP request.\n* SIP_UDP - SIP_UDP request.\n* SIP_TCP - SIP_TCP request.\n* MYSQL - MYSQL request.\n* MSSQL - MSSQL request.\n* ORACLE - ORACLE request.\n* NAT - NAT request.\n* DIAMETER - DIAMETER request.\n* RADIUS - RADIUS request.\n* MQTT - MQTT request.\n* QUIC_BRIDGE - QUIC_BRIDGE request.\n* HTTP_QUIC - HTTP_QUIC request.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of policies bound to the label.",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of times the policy label was invoked.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"labeltype": schema.StringAttribute{
				Computed:    true,
				Description: "Type of policy label to invoke. Possible values: [ reqvserver, policylabel ].",
			},
			"invoke_labelname": schema.StringAttribute{
				Computed:    true,
				Description: "If labelType is policylabel, name of the policy label to invoke; if labelType is reqvserver, name of the virtual server.",
			},
		},
	}
}

// lbpolicylabelDataSourceSetAttrFromGet projects a NITRO lbpolicylabel GET
// response onto the data-source model. Attributes are simply filled from the GET
// (or left Null when the GET omits them) via the shared utils.MapGet* helpers.
func lbpolicylabelDataSourceSetAttrFromGet(ctx context.Context, data *LbpolicylabelDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In lbpolicylabelDataSourceSetAttrFromGet Function")

	if v, ok := g["labelname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Labelname = types.StringValue(utils.AnyToString(v))
	} else {
		data.Id = data.Labelname
	}

	// Read/write attributes as read-back outputs.
	data.Comment = utils.MapGetString(g, "comment")
	data.Policylabeltype = utils.MapGetString(g, "policylabeltype")

	// newname is rename-only and never returned by GET -> Null.
	data.Newname = types.StringNull()

	// Read-only metadata.
	data.Numpol = utils.MapGetInt64(g, "numpol")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.InvokeLabelname = utils.MapGetString(g, "invoke_labelname")
}
