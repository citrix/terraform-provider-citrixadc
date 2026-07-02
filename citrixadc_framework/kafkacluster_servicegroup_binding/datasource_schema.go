package kafkacluster_servicegroup_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func KafkaclusterServicegroupBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Kafka cluster",
			},
			"servicegroupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the bound servicegroup.",
			},
		},
	}
}
