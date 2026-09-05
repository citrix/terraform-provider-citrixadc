package icaparameter

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// IcaparameterDataSourceModel is the data-source-specific model, decoupled from
// IcaparameterResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (builtin). Every non-key attribute is Computed; the Framework's per-attribute
// model <-> schema reflection requires this model to have exactly the attributes
// the data-source schema declares, which is why it cannot reuse the resource
// model. icaparameter is a singleton, so the ID is static.
type IcaparameterDataSourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Dfpersistence         types.String `tfsdk:"dfpersistence"`
	Edtlosstolerant       types.String `tfsdk:"edtlosstolerant"`
	Edtpmtuddf            types.String `tfsdk:"edtpmtuddf"`
	Edtpmtuddftimeout     types.Int64  `tfsdk:"edtpmtuddftimeout"`
	Edtpmtudrediscovery   types.String `tfsdk:"edtpmtudrediscovery"`
	Enablesronhafailover  types.String `tfsdk:"enablesronhafailover"`
	Hdxinsightnonnsap     types.String `tfsdk:"hdxinsightnonnsap"`
	Insightonlytodirector types.String `tfsdk:"insightonlytodirector"`
	L7latencyfrequency    types.Int64  `tfsdk:"l7latencyfrequency"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/icaparameter.json). Never settable; populated from GET.
	Builtin types.List `tfsdk:"builtin"`
}

func IcaparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"dfpersistence": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable DF Persistence",
			},
			"edtlosstolerant": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable EDT Loss Tolerant feature",
			},
			"edtpmtuddf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable DF enforcement for EDT PMTUD Control Blocks",
			},
			"edtpmtuddftimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "DF enforcement timeout for EDTPMTUDDF",
			},
			"edtpmtudrediscovery": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable EDT PMTUD Rediscovery",
			},
			"enablesronhafailover": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable Session Reliability on HA failover. The default value is No",
			},
			"hdxinsightnonnsap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable HDXInsight for Non NSAP ICA Sessions. The default value is Yes",
			},
			"insightonlytodirector": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable HDX Insight data to Director even if HDX Insight policy is not configured on Gateway and Network Telemtry policy is enabled on VDA. Default value: ENABLED Possible values = ENABLED, DISABLED",
			},
			"l7latencyfrequency": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the time interval/period for which L7 Client Latency value is to be calculated. By default, L7 Client Latency is calculated for every packet. The default value is 0",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that the ICA parameter is a built-in (SYSTEM INTERNAL) type. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
		},
	}
}

// icaparameterDataSourceSetAttrFromGet projects a NITRO icaparameter GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection. icaparameter is a singleton, so the ID is static.
func icaparameterDataSourceSetAttrFromGet(ctx context.Context, data *IcaparameterDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In icaparameterDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Dfpersistence = utils.MapGetString(g, "dfpersistence")
	data.Edtlosstolerant = utils.MapGetString(g, "edtlosstolerant")
	data.Edtpmtuddf = utils.MapGetString(g, "edtpmtuddf")
	data.Edtpmtuddftimeout = utils.MapGetInt64(g, "edtpmtuddftimeout")
	data.Edtpmtudrediscovery = utils.MapGetString(g, "edtpmtudrediscovery")
	data.Enablesronhafailover = utils.MapGetString(g, "enablesronhafailover")
	data.Hdxinsightnonnsap = utils.MapGetString(g, "hdxinsightnonnsap")
	data.Insightonlytodirector = utils.MapGetString(g, "insightonlytodirector")
	data.L7latencyfrequency = utils.MapGetInt64(g, "l7latencyfrequency")

	// Read-only attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")

	// Singleton (no unique attributes) - static ID.
	data.Id = types.StringValue("icaparameter-config")
}
