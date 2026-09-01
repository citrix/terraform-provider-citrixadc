package systemextramgmtcpu

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SystemextramgmtcpuDataSourceModel is the data-source-specific model, decoupled
// from SystemextramgmtcpuResourceModel. systemextramgmtcpu is a singleton
// (keyless) config object. A data source is a pure read surface, so it can
// expose the FULL GET projection: the resource-facing knobs (as Computed
// outputs) AND the read-only state the appliance returns on GET (configuredstate,
// effectivestate) that the resource omits.
type SystemextramgmtcpuDataSourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Enabled               types.Bool   `tfsdk:"enabled"`
	Reboot                types.Bool   `tfsdk:"reboot"`
	ReachableTimeout      types.String `tfsdk:"reachable_timeout"`
	ReachablePollDelay    types.String `tfsdk:"reachable_poll_delay"`
	ReachablePollInterval types.String `tfsdk:"reachable_poll_interval"`
	ReachablePollTimeout  types.String `tfsdk:"reachable_poll_timeout"`

	// Read-only (GET-only) state from the NITRO read-only set
	// (zion73x_readonly/systemextramgmtcpu.json). Never settable; from GET.
	Configuredstate types.String `tfsdk:"configuredstate"`
	Effectivestate  types.String `tfsdk:"effectivestate"`
}

func SystemextramgmtcpuDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Boolean value indicating the effective state of the extra management CPU.",
			},
			// The datasource surfaces these resource-only write knobs (never returned
			// by NITRO / meaningless for a datasource read) so the model stays
			// aligned with the resource user-facing contract.
			"reboot": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"reachable_timeout": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"reachable_poll_delay": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"reachable_poll_interval": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"reachable_poll_timeout": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},

			// Read-only (GET-only) state surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"configuredstate": schema.StringAttribute{
				Computed:    true,
				Description: "Configured state of extra management cpu. Possible values: ENABLED, DISABLED.",
			},
			"effectivestate": schema.StringAttribute{
				Computed:    true,
				Description: "Current running state of extra management cpu. Possible values: ENABLED, DISABLED.",
			},
		},
	}
}

// systemextramgmtcpuDataSourceSetAttrFromGet projects a NITRO systemextramgmtcpu
// GET response onto the data-source model. The read-only state fields are filled
// from the GET via the shared utils.MapGet* helpers; the action-only knobs the
// GET never returns are set Null.
func systemextramgmtcpuDataSourceSetAttrFromGet(ctx context.Context, data *SystemextramgmtcpuDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In systemextramgmtcpuDataSourceSetAttrFromGet Function")

	// effectivestate is always present in the GET response; mirror the resource
	// read by mapping it to the bool `enabled`.
	if v, ok := g["effectivestate"]; ok && v != nil && utils.AnyToString(v) == "ENABLED" {
		data.Enabled = types.BoolValue(true)
	} else {
		data.Enabled = types.BoolValue(false)
	}

	// Singleton resource -> static ID (matches the resource/datasource contract).
	data.Id = types.StringValue("systemextramgmtcpu-config")

	// reboot / reachable_* are Terraform-only action knobs never returned by
	// NITRO -> Null on a datasource read.
	data.Reboot = types.BoolNull()
	data.ReachableTimeout = types.StringNull()
	data.ReachablePollDelay = types.StringNull()
	data.ReachablePollInterval = types.StringNull()
	data.ReachablePollTimeout = types.StringNull()

	// Read-only state.
	data.Configuredstate = utils.MapGetString(g, "configuredstate")
	data.Effectivestate = utils.MapGetString(g, "effectivestate")
}
