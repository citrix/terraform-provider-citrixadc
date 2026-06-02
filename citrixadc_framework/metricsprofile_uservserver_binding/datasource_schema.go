package metricsprofile_uservserver_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func MetricsprofileUservserverBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"entityname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the entity bound to the metrics profile.",
			},
			"entitytype": schema.StringAttribute{
				Required:    true,
				Description: "Type of the entity bound to the metrics profile.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the metrics profile. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters.!",
			},
		},
	}
}