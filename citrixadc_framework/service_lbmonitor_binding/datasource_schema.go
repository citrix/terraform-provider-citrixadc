package service_lbmonitor_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ServiceLbmonitorBindingDataSourceModel is the data-source-specific model,
// decoupled from the resource model. A data source is a pure read surface (Read
// only; no plan/apply lifecycle), so it can expose the FULL GET projection: the
// read/write attributes (as Computed outputs) AND the read-only monitor runtime
// metadata that the resource deliberately omits. Every non-key attribute is
// Computed.
type ServiceLbmonitorBindingDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	MonitorName types.String `tfsdk:"monitor_name"`
	Monstate    types.String `tfsdk:"monstate"`
	Name        types.String `tfsdk:"name"`
	Passive     types.Bool   `tfsdk:"passive"`
	Weight      types.Int64  `tfsdk:"weight"`

	// Read-only (GET-only) monitor runtime metadata from the NITRO doc read-only
	// set (zion73x_readonly/service_lbmonitor_binding.json). Never settable;
	// populated from GET, Null when the appliance omits them.
	Monitortotalfailedprobes   types.Int64  `tfsdk:"monitortotalfailedprobes"`
	Lastresponse               types.String `tfsdk:"lastresponse"`
	Failedprobes               types.Int64  `tfsdk:"failedprobes"`
	Monstatparam2              types.Int64  `tfsdk:"monstatparam2"`
	Totalprobes                types.Int64  `tfsdk:"totalprobes"`
	DupWeight                  types.Int64  `tfsdk:"dup_weight"`
	Monitortotalprobes         types.Int64  `tfsdk:"monitortotalprobes"`
	Monstatparam1              types.Int64  `tfsdk:"monstatparam1"`
	Monitorcurrentfailedprobes types.Int64  `tfsdk:"monitorcurrentfailedprobes"`
	Monstatcode                types.Int64  `tfsdk:"monstatcode"`
	MonitorState               types.String `tfsdk:"monitor_state"`
	Totalfailedprobes          types.Int64  `tfsdk:"totalfailedprobes"`
	DupState                   types.String `tfsdk:"dup_state"`
	Responsetime               types.Int64  `tfsdk:"responsetime"`
	Monstatparam3              types.Int64  `tfsdk:"monstatparam3"`
}

func ServiceLbmonitorBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"monitor_name": schema.StringAttribute{
				Required:    true,
				Description: "The monitor Names.",
			},
			"monstate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The configured state (enable/disable) of the monitor on this server.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the service to which to bind a monitor.",
			},
			"passive": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the monitor-service binding. When a monitor is UP, the weight assigned to its binding with the service determines how much the monitor contributes toward keeping the health of the service above the value configured for the Monitor Threshold parameter.",
			},

			// Read-only (GET-only) monitor runtime metadata surfaced by the data source.
			"monitortotalfailedprobes": schema.Int64Attribute{
				Computed:    true,
				Description: "Total number of failed probes.",
			},
			"lastresponse": schema.StringAttribute{
				Computed:    true,
				Description: "The string form of monstatcode.",
			},
			"failedprobes": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of the current failed monitoring probes.",
			},
			"monstatparam2": schema.Int64Attribute{
				Computed:    true,
				Description: "Second parameter for use with message code.",
			},
			"totalprobes": schema.Int64Attribute{
				Computed:    true,
				Description: "The total number of probes sent.",
			},
			"dup_weight": schema.Int64Attribute{
				Computed:    true,
				Description: "The weight of the monitor.",
			},
			"monitortotalprobes": schema.Int64Attribute{
				Computed:    true,
				Description: "Total number of probes sent to monitor this service.",
			},
			"monstatparam1": schema.Int64Attribute{
				Computed:    true,
				Description: "First parameter for use with message code.",
			},
			"monitorcurrentfailedprobes": schema.Int64Attribute{
				Computed:    true,
				Description: "Total number of currently failed probes.",
			},
			"monstatcode": schema.Int64Attribute{
				Computed:    true,
				Description: "The code indicating the monitor response.",
			},
			"monitor_state": schema.StringAttribute{
				Computed:    true,
				Description: "The running state of the monitor on this service (for example UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, DISABLED).",
			},
			"totalfailedprobes": schema.Int64Attribute{
				Computed:    true,
				Description: "The total number of failed probes.",
			},
			"dup_state": schema.StringAttribute{
				Computed:    true,
				Description: "State value from table (ENABLED or DISABLED).",
			},
			"responsetime": schema.Int64Attribute{
				Computed:    true,
				Description: "Response time of this monitor.",
			},
			"monstatparam3": schema.Int64Attribute{
				Computed:    true,
				Description: "Third parameter for use with message code.",
			},
		},
	}
}

// service_lbmonitor_bindingDataSourceSetAttrFromGet projects a NITRO
// service_lbmonitor_binding GET response onto the data-source model.
func service_lbmonitor_bindingDataSourceSetAttrFromGet(ctx context.Context, data *ServiceLbmonitorBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In service_lbmonitor_bindingDataSourceSetAttrFromGet Function")

	data.MonitorName = utils.MapGetString(g, "monitor_name")
	data.Monstate = utils.MapGetString(g, "monstate")
	data.Name = utils.MapGetString(g, "name")
	data.Passive = utils.MapGetBool(g, "passive")
	data.Weight = utils.MapGetInt64(g, "weight")

	// Read-only (GET-only) monitor runtime metadata.
	data.Monitortotalfailedprobes = utils.MapGetInt64(g, "monitortotalfailedprobes")
	data.Lastresponse = utils.MapGetString(g, "lastresponse")
	data.Failedprobes = utils.MapGetInt64(g, "failedprobes")
	data.Monstatparam2 = utils.MapGetInt64(g, "monstatparam2")
	data.Totalprobes = utils.MapGetInt64(g, "totalprobes")
	data.DupWeight = utils.MapGetInt64(g, "dup_weight")
	data.Monitortotalprobes = utils.MapGetInt64(g, "monitortotalprobes")
	data.Monstatparam1 = utils.MapGetInt64(g, "monstatparam1")
	data.Monitorcurrentfailedprobes = utils.MapGetInt64(g, "monitorcurrentfailedprobes")
	data.Monstatcode = utils.MapGetInt64(g, "monstatcode")
	data.MonitorState = utils.MapGetString(g, "monitor_state")
	data.Totalfailedprobes = utils.MapGetInt64(g, "totalfailedprobes")
	data.DupState = utils.MapGetString(g, "dup_state")
	data.Responsetime = utils.MapGetInt64(g, "responsetime")
	data.Monstatparam3 = utils.MapGetInt64(g, "monstatparam3")

	// Set composite ID. Backward-compatible with SDK v2: identity is
	// "monitor_name,name" (comma-separated key:UrlEncode(value) pairs).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("monitor_name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.MonitorName.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
