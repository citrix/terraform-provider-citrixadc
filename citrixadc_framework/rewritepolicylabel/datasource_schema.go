package rewritepolicylabel

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RewritepolicylabelDataSourceModel is the data-source-specific model, decoupled
// from RewritepolicylabelResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (numpol, hits, priority, isdefault, ...). Every non-key attribute is
// Computed; the Framework's per-attribute model <-> schema reflection requires
// this model to have exactly the attributes the data-source schema declares.
type RewritepolicylabelDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Labelname types.String `tfsdk:"labelname"` // Required lookup key
	Newname   types.String `tfsdk:"newname"`
	Transform types.String `tfsdk:"transform"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/rewritepolicylabel.json). Never settable; populated from GET.
	Numpol                 types.Int64  `tfsdk:"numpol"`
	Hits                   types.Int64  `tfsdk:"hits"`
	Priority               types.Int64  `tfsdk:"priority"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Labeltype              types.String `tfsdk:"labeltype"`
	InvokeLabelname        types.String `tfsdk:"invoke_labelname"`
	Flowtype               types.Int64  `tfsdk:"flowtype"`
	Description            types.String `tfsdk:"description"`
	Isdefault              types.Bool   `tfsdk:"isdefault"`
	Builtin                types.List   `tfsdk:"builtin"`
	Feature                types.String `tfsdk:"feature"`
}

func RewritepolicylabelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this rewrite policy label.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the rewrite policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the rewrite policy label is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my rewrite policy label\" or 'my rewrite policy label').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the rewrite policy label. \nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy label\" or 'my policy label').",
			},
			"transform": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Types of transformations allowed by the policies bound to the label. For Rewrite, the following types are supported:\n* http_req - HTTP requests\n* http_res - HTTP responses\n* othertcp_req - Non-HTTP TCP requests\n* othertcp_res - Non-HTTP TCP responses\n* url - URLs\n* text - Text strings\n* clientless_vpn_req - Citrix ADC clientless VPN requests\n* clientless_vpn_res - Citrix ADC clientless VPN responses\n* sipudp_req - SIP requests\n* sipudp_res - SIP responses\n* diameter_req - DIAMETER requests\n* diameter_res - DIAMETER responses\n* radius_req - RADIUS requests\n* radius_res - RADIUS responses\n* dns_req - DNS requests\n* dns_res - DNS responses\n* mqtt_req - MQTT requests\n* mqtt_res - MQTT responses",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of polices bound to label.",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of times policy label was invoked.",
			},
			"priority": schema.Int64Attribute{
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"labeltype": schema.StringAttribute{
				Computed:    true,
				Description: "Type of invocation. Possible values = reqvserver, resvserver, resHttpEventvserver, policylabel.",
			},
			"invoke_labelname": schema.StringAttribute{
				Computed:    true,
				Description: "If labelType is policylabel, name of the policy label to invoke. If labelType is reqvserver or resvserver, name of the virtual server to which to forward the request or response.",
			},
			"flowtype": schema.Int64Attribute{
				Computed:    true,
				Description: "Flowtype of the bound rewrite policy.",
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: "Description of the policylabel.",
			},
			"isdefault": schema.BoolAttribute{
				Computed:    true,
				Description: "A value of true is returned if it is a default rewritepolicylabel.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if rewrite policy label is built-in or not. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// rewritepolicylabelDataSourceSetAttrFromGet projects a NITRO rewritepolicylabel
// GET response onto the data-source model. Because a data source has no
// plan/apply reconciliation, attributes are simply filled from the GET (or left
// Null when the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func rewritepolicylabelDataSourceSetAttrFromGet(ctx context.Context, data *RewritepolicylabelDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In rewritepolicylabelDataSourceSetAttrFromGet Function")

	if v, ok := g["labelname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Labelname = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Comment = utils.MapGetString(g, "comment")
	data.Transform = utils.MapGetString(g, "transform")

	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only metadata.
	data.Numpol = utils.MapGetInt64(g, "numpol")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.InvokeLabelname = utils.MapGetString(g, "invoke_labelname")
	data.Flowtype = utils.MapGetInt64(g, "flowtype")
	data.Description = utils.MapGetString(g, "description")
	data.Isdefault = utils.MapGetBool(g, "isdefault")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
