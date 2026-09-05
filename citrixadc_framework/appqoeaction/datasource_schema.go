package appqoeaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppqoeactionDataSourceModel is the data-source-specific model, decoupled from
// AppqoeactionResourceModel. A data source is a pure read surface, so it can
// expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits.
type AppqoeactionDataSourceModel struct {
	Id                types.String `tfsdk:"id"`
	Altcontentpath    types.String `tfsdk:"altcontentpath"`
	Altcontentsvcname types.String `tfsdk:"altcontentsvcname"`
	Customfile        types.String `tfsdk:"customfile"`
	Delay             types.Int64  `tfsdk:"delay"`
	Dosaction         types.String `tfsdk:"dosaction"`
	Dostrigexpression types.String `tfsdk:"dostrigexpression"`
	Maxconn           types.Int64  `tfsdk:"maxconn"`
	Name              types.String `tfsdk:"name"`
	Numretries        types.String `tfsdk:"numretries"`
	Polqdepth         types.Int64  `tfsdk:"polqdepth"`
	Priority          types.String `tfsdk:"priority"`
	Priqdepth         types.Int64  `tfsdk:"priqdepth"`
	Respondwith       types.String `tfsdk:"respondwith"`
	Retryonreset      types.String `tfsdk:"retryonreset"`
	Retryontimeout    types.Int64  `tfsdk:"retryontimeout"`
	Tcpprofile        types.String `tfsdk:"tcpprofile"`

	// Read-only (GET-only) attribute from the NITRO doc read-only set
	// (zion73x_readonly/appqoeaction.json). Never settable; populated from GET.
	Hits types.Int64 `tfsdk:"hits"`
}

func AppqoeactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"altcontentpath": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Path to the alternative content service to be used in the ACS",
			},
			"altcontentsvcname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the alternative content service to be used in the ACS",
			},
			"customfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "name of the HTML page object to use as the response",
			},
			"delay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Delay threshold, in microseconds, for requests that match the policy's rule. If the delay statistics gathered for the matching request exceed the specified delay, configured action triggered for that request, if there is no action then requests are dropped to the lowest priority level",
			},
			"dosaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DoS Action to take when vserver will be considered under DoS attack and corresponding rule matches. Mandatory if AppQoE actions are to be used for DoS attack prevention.",
			},
			"dostrigexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Optional expression to add second level check to trigger DoS actions. Specifically used for Analytics based DoS response generation",
			},
			"maxconn": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of concurrent connections that can be open for requests that matches with rule.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the AppQoE action. Must begin with a letter, number, or the underscore symbol (_). Other characters allowed, after the first character, are the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), and colon (:) characters. This is a mandatory argument",
			},
			"numretries": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Retry count",
			},
			"polqdepth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Policy queue depth threshold value. When the policy queue size (number of requests queued for the policy binding this action is attached to) increases to the specified polqDepth value, subsequent requests are dropped to the lowest priority level.",
			},
			"priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority for queuing the request. If server resources are not available for a request that matches the configured rule, this option specifies a priority for queuing the request until the server resources are available again. If priority is not configured then Lowest priority will be used to queue the request.",
			},
			"priqdepth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Queue depth threshold value per priorirty level. If the queue size (number of requests in the queue of that particular priorirty) on the virtual server to which this policy is bound, increases to the specified qDepth value, subsequent requests are dropped to the lowest priority level.",
			},
			"respondwith": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Responder action to be taken when the threshold is reached. Available settings function as follows:\n            ACS - Serve content from an alternative content service\n                  Threshold : maxConn or delay\n            NS - Serve from the Citrix ADC (built-in response)\n                 Threshold : maxConn or delay",
			},
			"retryonreset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Retry on TCP Reset",
			},
			"retryontimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Retry on request Timeout(in millisec) upon sending request to backend servers",
			},
			"tcpprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bind TCP Profile based on L2/L3/L7 parameters.",
			},

			// Read-only (GET-only) attribute surfaced by the data source.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of times the AppQoE action was applied.",
			},
		},
	}
}

// appqoeactionDataSourceSetAttrFromGet projects a NITRO appqoeaction GET
// response onto the data-source model. Attributes are simply filled from the GET
// (or left Null when the GET omits them) via the shared utils.MapGet* helpers.
func appqoeactionDataSourceSetAttrFromGet(ctx context.Context, data *AppqoeactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appqoeactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Altcontentpath = utils.MapGetString(g, "altcontentpath")
	data.Altcontentsvcname = utils.MapGetString(g, "altcontentsvcname")
	data.Customfile = utils.MapGetString(g, "customfile")
	data.Delay = utils.MapGetInt64(g, "delay")
	data.Dosaction = utils.MapGetString(g, "dosaction")
	data.Dostrigexpression = utils.MapGetString(g, "dostrigexpression")
	data.Maxconn = utils.MapGetInt64(g, "maxconn")
	data.Numretries = utils.MapGetString(g, "numretries")
	data.Polqdepth = utils.MapGetInt64(g, "polqdepth")
	data.Priority = utils.MapGetString(g, "priority")
	data.Priqdepth = utils.MapGetInt64(g, "priqdepth")
	data.Respondwith = utils.MapGetString(g, "respondwith")
	data.Retryonreset = utils.MapGetString(g, "retryonreset")
	data.Retryontimeout = utils.MapGetInt64(g, "retryontimeout")
	data.Tcpprofile = utils.MapGetString(g, "tcpprofile")

	// Read-only (GET-only) attribute.
	data.Hits = utils.MapGetInt64(g, "hits")
}
