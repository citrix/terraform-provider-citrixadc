package vrid6

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Vrid6DataSourceModel is the data-source-specific model, decoupled from
// Vrid6ResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (ifaces, ifnum, type, state, flags, ipaddress, effectivepriority,
// operationalownernode). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type Vrid6DataSourceModel struct {
	Id       types.String `tfsdk:"id"`
	Vrid6_id types.Int64  `tfsdk:"vrid6_id"` // Required lookup key

	// Read/write attributes, surfaced here as Computed outputs.
	All                  types.Bool   `tfsdk:"all"`
	Ownernode            types.Int64  `tfsdk:"ownernode"`
	Preemption           types.String `tfsdk:"preemption"`
	Preemptiondelaytimer types.Int64  `tfsdk:"preemptiondelaytimer"`
	Priority             types.Int64  `tfsdk:"priority"`
	Sharing              types.String `tfsdk:"sharing"`
	Trackifnumpriority   types.Int64  `tfsdk:"trackifnumpriority"`
	Tracking             types.String `tfsdk:"tracking"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/vrid6.json). Never settable; populated from GET.
	Ifaces               types.String `tfsdk:"ifaces"`
	Ifnum                types.String `tfsdk:"ifnum"`
	Type                 types.String `tfsdk:"type"`
	State                types.Int64  `tfsdk:"state"`
	Flags                types.Int64  `tfsdk:"flags"`
	Ipaddress            types.String `tfsdk:"ipaddress"`
	Effectivepriority    types.Int64  `tfsdk:"effectivepriority"`
	Operationalownernode types.Int64  `tfsdk:"operationalownernode"`
}

func Vrid6DataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"all": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Remove all configured VMAC6 addresses from the Citrix ADC.",
			},
			"vrid6_id": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies a VMAC6 address.",
			},
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "In a cluster setup, assign a cluster node as the owner of this VMAC address for IP based VRRP configuration. If no owner is configured, ow ner node is displayed as ALL and one node is dynamically elected as the owner.",
			},
			"preemption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "In an active-active mode configuration, make a backup VIP address the master if its priority becomes higher than that of a master VIP address bound to this VMAC address.\n             If you disable pre-emption while a backup VIP address is the master, the backup VIP address remains master until the original master VIP's priority becomes higher than that of the current master.",
			},
			"preemptiondelaytimer": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Preemption delay time in seconds, in an active-active configuration. If any high priority node will come in network, it will wait for these many seconds before becoming master.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Base priority (BP), in an active-active mode configuration, which ordinarily determines the master VIP address.",
			},
			"sharing": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "In an active-active mode configuration, enable the backup VIP address to process any traffic instead of dropping it.",
			},
			"trackifnumpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority by which the Effective priority will be reduced if any of the tracked interfaces goes down in an active-active configuration.",
			},
			"tracking": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The effective priority (EP) value, relative to the base priority (BP) value in an active-active mode configuration. When EP is set to a value other than None, it is EP, not BP, which determines the master VIP address.\nAvailable settings function as follows:\n* NONE - No tracking. EP = BP\n* ALL -  If the status of all virtual servers is UP, EP = BP. Otherwise, EP = 0.\n* ONE - If the status of at least one virtual server is UP, EP = BP. Otherwise, EP = 0.\n* PROGRESSIVE - If the status of all virtual servers is UP, EP = BP. If the status of all virtual servers is DOWN, EP = 0. Otherwise EP = BP (1 - K/N), where N is the total number of virtual servers associated with the VIP address and K is the number of virtual servers for which the status is DOWN.\nDefault: NONE.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"ifaces": schema.StringAttribute{
				Computed:    true,
				Description: "Interfaces bound to this VRID.",
			},
			"ifnum": schema.StringAttribute{
				Computed:    true,
				Description: "Interfaces to bind to the VMAC6, specified in (slot/port) notation (for example, 1/2). Use spaces to separate multiple entries.",
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "Type (static or dynamic) of this VRID. Possible values = STATIC, DYNAMIC.",
			},
			"state": schema.Int64Attribute{
				Computed:    true,
				Description: "State of this VRID.",
			},
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Flags.",
			},
			"ipaddress": schema.StringAttribute{
				Computed:    true,
				Description: "The IP address bound to the VRID6.",
			},
			"effectivepriority": schema.Int64Attribute{
				Computed:    true,
				Description: "The effective priority of this VRID.",
			},
			"operationalownernode": schema.Int64Attribute{
				Computed:    true,
				Description: "Run time owner node of the vrid.",
			},
		},
	}
}

// vrid6DataSourceSetAttrFromGet projects a NITRO vrid6 GET response onto the
// data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them) — no unknown->null resolution or plan preservation is required. The
// shared utils.MapGet* helpers implement that projection.
func vrid6DataSourceSetAttrFromGet(ctx context.Context, data *Vrid6DataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vrid6DataSourceSetAttrFromGet Function")

	if v, ok := g["id"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
	}
	data.Vrid6_id = utils.MapGetInt64(g, "id")

	// Read/write attributes as read-back outputs.
	data.All = utils.MapGetBool(g, "all")
	data.Ownernode = utils.MapGetInt64(g, "ownernode")
	data.Preemption = utils.MapGetString(g, "preemption")
	data.Preemptiondelaytimer = utils.MapGetInt64(g, "preemptiondelaytimer")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Sharing = utils.MapGetString(g, "sharing")
	data.Trackifnumpriority = utils.MapGetInt64(g, "trackifnumpriority")
	data.Tracking = utils.MapGetString(g, "tracking")

	// Read-only metadata.
	data.Ifaces = utils.MapGetString(g, "ifaces")
	data.Ifnum = utils.MapGetString(g, "ifnum")
	data.Type = utils.MapGetString(g, "type")
	data.State = utils.MapGetInt64(g, "state")
	data.Flags = utils.MapGetInt64(g, "flags")
	data.Ipaddress = utils.MapGetString(g, "ipaddress")
	data.Effectivepriority = utils.MapGetInt64(g, "effectivepriority")
	data.Operationalownernode = utils.MapGetInt64(g, "operationalownernode")
}
