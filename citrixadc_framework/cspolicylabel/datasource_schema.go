package cspolicylabel

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CspolicylabelDataSourceModel is the data-source-specific model, decoupled from
// CspolicylabelResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (numpol, hits, policyname, priority, targetvserver, gotopriorityexpression,
// labeltype, invoke_labelname). Every non-key attribute is Computed; the
// Framework's per-attribute model <-> schema reflection requires this model to
// have exactly the attributes the data-source schema declares, which is why it
// cannot reuse the resource model.
type CspolicylabelDataSourceModel struct {
	Id                types.String `tfsdk:"id"`
	Labelname         types.String `tfsdk:"labelname"` // Required lookup key
	Cspolicylabeltype types.String `tfsdk:"cspolicylabeltype"`
	Newname           types.String `tfsdk:"newname"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/cspolicylabel.json). Never settable; populated from GET.
	Numpol                 types.Int64  `tfsdk:"numpol"`
	Hits                   types.Int64  `tfsdk:"hits"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Targetvserver          types.String `tfsdk:"targetvserver"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Labeltype              types.String `tfsdk:"labeltype"`
	InvokeLabelname        types.String `tfsdk:"invoke_labelname"`
}

func CspolicylabelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cspolicylabeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol supported by the policy label. All policies bound to the policy label must either match the specified protocol or be a subtype of that protocol. Available settings function as follows:\n* HTTP - Supports policies that process HTTP traffic. Used to access unencrypted Web sites. (The default.)\n* SSL - Supports policies that process HTTPS/SSL encrypted traffic. Used to access encrypted Web sites.\n* TCP - Supports policies that process any type of TCP traffic, including HTTP.\n* SSL_TCP - Supports policies that process SSL-encrypted TCP traffic, including SSL.\n* UDP - Supports policies that process any type of UDP-based traffic, including DNS.\n* DNS - Supports policies that process DNS traffic.\n* ANY - Supports all types of policies except HTTP, SSL, and TCP.\n* SIP_UDP - Supports policies that process UDP based Session Initiation Protocol (SIP) traffic. SIP initiates, manages, and terminates multimedia communications sessions, and has emerged as the standard for Internet telephony (VoIP).\n* RTSP - Supports policies that process Real Time Streaming Protocol (RTSP) traffic. RTSP provides delivery of multimedia and other streaming data, such as audio, video, and other types of streamed media.\n* RADIUS - Supports policies that process Remote Authentication Dial In User Service (RADIUS) traffic. RADIUS supports combined authentication, authorization, and auditing services for network management.\n* MYSQL - Supports policies that process MYSQL traffic.\n* MSSQL - Supports policies that process Microsoft SQL traffic.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the policy label. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.\nThe label name must be unique within the list of policy labels for content switching.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policylabel\" or 'my policylabel').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new name of the content switching policylabel.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of policies bound to the label.",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of times the policy label was invoked.",
			},
			"policyname": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the content switching policy.",
			},
			"priority": schema.Int64Attribute{
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
			"targetvserver": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the virtual server to which to forward requests that match the policy.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"labeltype": schema.StringAttribute{
				Computed:    true,
				Description: "Type of policy label invocation. Possible values = policylabel.",
			},
			"invoke_labelname": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the label to invoke if the current policy rule evaluates to TRUE.",
			},
		},
	}
}

// cspolicylabelDataSourceSetAttrFromGet projects a NITRO cspolicylabel GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func cspolicylabelDataSourceSetAttrFromGet(ctx context.Context, data *CspolicylabelDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In cspolicylabelDataSourceSetAttrFromGet Function")

	if v, ok := g["labelname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Labelname = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Cspolicylabeltype = utils.MapGetString(g, "cspolicylabeltype")
	data.Newname = utils.MapGetString(g, "newname")

	// Read-only attributes.
	data.Numpol = utils.MapGetInt64(g, "numpol")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Policyname = utils.MapGetString(g, "policyname")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Targetvserver = utils.MapGetString(g, "targetvserver")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.InvokeLabelname = utils.MapGetString(g, "invoke_labelname")
}
