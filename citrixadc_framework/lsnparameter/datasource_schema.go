package lsnparameter

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LsnparameterDataSourceModel is the data-source-specific model, decoupled from
// LsnparameterResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the configurable attributes (as
// Computed outputs) AND the read-only attributes the resource deliberately
// omits. Every non-key attribute is Computed.
type LsnparameterDataSourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Memlimit             types.Int64  `tfsdk:"memlimit"`
	Sessionsync          types.String `tfsdk:"sessionsync"`
	Subscrsessionremoval types.String `tfsdk:"subscrsessionremoval"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/lsnparameter.json). Never settable; populated from GET.
	Memlimitactive types.Int64 `tfsdk:"memlimitactive"`
	Maxmemlimit    types.Int64 `tfsdk:"maxmemlimit"`
}

func LsnparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to read the LSN (Large Scale NAT) global parameters.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"memlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Amount of Citrix ADC memory to reserve for the LSN feature, in multiples of 2MB.\n\nNote: If you later reduce the value of this parameter, the amount of active memory is not reduced. Changing the configured memory limit can only increase the amount of active memory.\nThis command is deprecated, use 'set extendedmemoryparam -memlimit' instead.",
			},
			"sessionsync": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Synchronize all LSN sessions with the secondary node in a high availability (HA) deployment (global synchronization). After a failover, established TCP connections and UDP packet flows are kept active and resumed on the secondary node (new primary).\n\nThe global session synchronization parameter and session synchronization parameters (group level) of all LSN groups are enabled by default.\n\nFor a group, when both the global level and the group level LSN session synchronization parameters are enabled, the primary node synchronizes information of all LSN sessions related to this LSN group with the secondary node.",
			},
			"subscrsessionremoval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "LSN global setting for controlling subscriber aware session removal, when this is enabled, when ever the subscriber info is deleted from subscriber database, sessions corresponding to that subscriber will be removed. if this setting is disabled, subscriber sessions will be timed out as per the idle time out settings.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"memlimitactive": schema.Int64Attribute{
				Computed:    true,
				Description: "Amount of actual memory reserved for the LSN feature. The amount of active memory for the LSN feature might be less than the configured memory, because the available memory is shared across features.",
			},
			"maxmemlimit": schema.Int64Attribute{
				Computed:    true,
				Description: "Maximum amount of Citrix ADC memory that can be reserved for the LSN feature.",
			},
		},
	}
}

// lsnparameterDataSourceSetAttrFromGet projects a NITRO lsnparameter GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func lsnparameterDataSourceSetAttrFromGet(ctx context.Context, data *LsnparameterDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In lsnparameterDataSourceSetAttrFromGet Function")

	// lsnparameter is a singleton with no lookup key; use a static ID.
	data.Id = types.StringValue("lsnparameter-config")

	data.Memlimit = utils.MapGetInt64(g, "memlimit")
	data.Sessionsync = utils.MapGetString(g, "sessionsync")
	data.Subscrsessionremoval = utils.MapGetString(g, "subscrsessionremoval")

	// Read-only attributes.
	data.Memlimitactive = utils.MapGetInt64(g, "memlimitactive")
	data.Maxmemlimit = utils.MapGetInt64(g, "maxmemlimit")
}
