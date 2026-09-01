package gslbservice_lbmonitor_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// GslbserviceLbmonitorBindingDataSourceModel is the data-source-specific model,
// decoupled from GslbserviceLbmonitorBindingResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares.
type GslbserviceLbmonitorBindingDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	MonitorName types.String `tfsdk:"monitor_name"`
	Monstate    types.String `tfsdk:"monstate"`
	Servicename types.String `tfsdk:"servicename"`
	Weight      types.Int64  `tfsdk:"weight"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/gslbservice_lbmonitor_binding.json). Never settable;
	// populated from GET.
	Monstatparam2              types.Int64  `tfsdk:"monstatparam2"`
	Failedprobes               types.Int64  `tfsdk:"failedprobes"`
	Totalfailedprobes          types.Int64  `tfsdk:"totalfailedprobes"`
	Lastresponse               types.String `tfsdk:"lastresponse"`
	MonitorState               types.String `tfsdk:"monitor_state"`
	Monitortotalfailedprobes   types.Int64  `tfsdk:"monitortotalfailedprobes"`
	Monstatcode                types.Int64  `tfsdk:"monstatcode"`
	Monitorcurrentfailedprobes types.Int64  `tfsdk:"monitorcurrentfailedprobes"`
	Monstatparam1              types.Int64  `tfsdk:"monstatparam1"`
	Monstatparam3              types.Int64  `tfsdk:"monstatparam3"`
	Responsetime               types.Int64  `tfsdk:"responsetime"`
	Monitortotalprobes         types.Int64  `tfsdk:"monitortotalprobes"`
}

func GslbserviceLbmonitorBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"monitor_name": schema.StringAttribute{
				Required:    true,
				Description: "Monitor name.",
			},
			"monstate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of the monitor bound to gslb service.",
			},
			"servicename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the GSLB service.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the monitor-service binding. A larger number specifies a greater weight. Contributes to the monitoring threshold, which determines the state of the service.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"monstatparam2": schema.Int64Attribute{
				Computed:    true,
				Description: "Second parameter for use with message code.",
			},
			"failedprobes": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of the current failed monitoring probes.",
			},
			"totalfailedprobes": schema.Int64Attribute{
				Computed:    true,
				Description: "The total number of failed probs.",
			},
			"lastresponse": schema.StringAttribute{
				Computed:    true,
				Description: "Displays the gslb monitor status in string format.",
			},
			"monitor_state": schema.StringAttribute{
				Computed:    true,
				Description: "The running state of the monitor on this service. Possible values = UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, GOING OUT OF SERVICE, DOWN WHEN GOING OUT OF SERVICE, NS_EMPTY_STR, Unknown, DISABLED.",
			},
			"monitortotalfailedprobes": schema.Int64Attribute{
				Computed:    true,
				Description: "Total number of failed probes.",
			},
			"monstatcode": schema.Int64Attribute{
				Computed:    true,
				Description: "The code indicating the monitor response.",
			},
			"monitorcurrentfailedprobes": schema.Int64Attribute{
				Computed:    true,
				Description: "Total number of currently failed probes.",
			},
			"monstatparam1": schema.Int64Attribute{
				Computed:    true,
				Description: "First parameter for use with message code.",
			},
			"monstatparam3": schema.Int64Attribute{
				Computed:    true,
				Description: "Third parameter for use with message code.",
			},
			"responsetime": schema.Int64Attribute{
				Computed:    true,
				Description: "Response time of this monitor.",
			},
			"monitortotalprobes": schema.Int64Attribute{
				Computed:    true,
				Description: "Total number of probes sent to monitor this service.",
			},
		},
	}
}

// gslbservice_lbmonitor_bindingDataSourceSetAttrFromGet projects a NITRO
// gslbservice_lbmonitor_binding GET response onto the data-source model. Because
// a data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func gslbservice_lbmonitor_bindingDataSourceSetAttrFromGet(ctx context.Context, data *GslbserviceLbmonitorBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In gslbservice_lbmonitor_bindingDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.MonitorName = utils.MapGetString(g, "monitor_name")
	data.Monstate = utils.MapGetString(g, "monstate")
	data.Servicename = utils.MapGetString(g, "servicename")
	data.Weight = utils.MapGetInt64(g, "weight")

	// Read-only attributes.
	data.Monstatparam2 = utils.MapGetInt64(g, "monstatparam2")
	data.Failedprobes = utils.MapGetInt64(g, "failedprobes")
	data.Totalfailedprobes = utils.MapGetInt64(g, "totalfailedprobes")
	data.Lastresponse = utils.MapGetString(g, "lastresponse")
	data.MonitorState = utils.MapGetString(g, "monitor_state")
	data.Monitortotalfailedprobes = utils.MapGetInt64(g, "monitortotalfailedprobes")
	data.Monstatcode = utils.MapGetInt64(g, "monstatcode")
	data.Monitorcurrentfailedprobes = utils.MapGetInt64(g, "monitorcurrentfailedprobes")
	data.Monstatparam1 = utils.MapGetInt64(g, "monstatparam1")
	data.Monstatparam3 = utils.MapGetInt64(g, "monstatparam3")
	data.Responsetime = utils.MapGetInt64(g, "responsetime")
	data.Monitortotalprobes = utils.MapGetInt64(g, "monitortotalprobes")

	// Composite binding ID: comma-separated key:UrlEncode(value) pairs.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("monitor_name:%s", utils.UrlEncode(data.MonitorName.ValueString())))
	idParts = append(idParts, fmt.Sprintf("servicename:%s", utils.UrlEncode(data.Servicename.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
