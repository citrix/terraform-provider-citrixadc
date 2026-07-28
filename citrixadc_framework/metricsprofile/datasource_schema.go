package metricsprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func MetricsprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
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
				Optional:    true,
				Computed:    true,
				Description: "Token for authenticating with the endpoint. If the endpoint requires the Authorization header in a particular format, specify the complete format as the value to this parameter. For eg., in case of splunk, the Authorizaiton header is required to be of the form - Splunk <auth-token>.",
			},
			"metricsauthtoken_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Token for authenticating with the endpoint. If the endpoint requires the Authorization header in a particular format, specify the complete format as the value to this parameter. For eg., in case of splunk, the Authorizaiton header is required to be of the form - Splunk <auth-token>.",
			},
			"metricsauthtoken_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a metricsauthtoken_wo update.",
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
		},
	}
}
