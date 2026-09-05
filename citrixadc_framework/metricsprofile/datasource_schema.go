package metricsprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MetricsprofileDataSourceModel is the data-source-specific model, decoupled
// from MetricsprofileResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the configurable attributes (as
// Computed outputs) AND the read-only attributes the resource deliberately
// omits. Every non-key attribute is Computed.
type MetricsprofileDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Collector              types.String `tfsdk:"collector"`
	Metrics                types.String `tfsdk:"metrics"`
	Metricsauthtoken       types.String `tfsdk:"metricsauthtoken"`
	Metricsendpointurl     types.String `tfsdk:"metricsendpointurl"`
	Metricsexportfrequency types.Int64  `tfsdk:"metricsexportfrequency"`
	Name                   types.String `tfsdk:"name"` // Required lookup key
	Outputmode             types.String `tfsdk:"outputmode"`
	Schemafile             types.String `tfsdk:"schemafile"`
	Servemode              types.String `tfsdk:"servemode"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/metricsprofile.json). Never settable; populated from GET.
	Refcnt types.Int64 `tfsdk:"refcnt"`
}

func MetricsprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to read a metrics profile configuration.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"collector": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The collector should be a HTTP/HTTPS service.",
			},
			"metrics": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is used enable or disable metrics",
			},
			"metricsauthtoken": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "Token for authenticating with the endpoint. If the endpoint requires the Authorization header in a particular format, specify the complete format as the value to this parameter. For eg., in case of splunk, the Authorizaiton header is required to be of the form - Splunk <auth-token>.",
			},
			"metricsendpointurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The URL at which to upload the metrics data on the endpoint",
			},
			"metricsexportfrequency": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is for configuring the metrics export frequency in seconds, frequency value must be in [30,300] seconds range",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the metrics profile. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters.!",
			},
			"outputmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option indicates the format in which metrics data is generated",
			},
			"schemafile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is for configuring json schema file containing a list of counters to be exported by metricscollector. Schema file should be present under /var/metrics_conf path",
			},
			"servemode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is to configure metrics pull or push mode. In push mode metricscollector exports metrics to configured collector. In pull mode, metricscollector only generates the metrics which will be pulled by external agent. No collector configuration is required in pull mode and it is applicable only for output mode Prometheus",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"refcnt": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of references to the profile.",
			},
		},
	}
}

// metricsprofileDataSourceSetAttrFromGet projects a NITRO metricsprofile GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func metricsprofileDataSourceSetAttrFromGet(ctx context.Context, data *MetricsprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In metricsprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Collector = utils.MapGetString(g, "collector")
	data.Metrics = utils.MapGetString(g, "metrics")
	data.Metricsendpointurl = utils.MapGetString(g, "metricsendpointurl")
	data.Metricsexportfrequency = utils.MapGetInt64(g, "metricsexportfrequency")
	data.Outputmode = utils.MapGetString(g, "outputmode")
	data.Schemafile = utils.MapGetString(g, "schemafile")
	data.Servemode = utils.MapGetString(g, "servemode")

	// metricsauthtoken is a secret input the GET never returns -> Null.
	data.Metricsauthtoken = types.StringNull()

	// Read-only attributes.
	data.Refcnt = utils.MapGetInt64(g, "refcnt")
}
