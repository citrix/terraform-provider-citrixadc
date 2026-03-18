package subscriberradiusinterface

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SubscriberradiusinterfaceDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"listeningservice": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of RADIUS LISTENING service that will process RADIUS accounting requests.",
			},
			"radiusinterimasstart": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Treat radius interim message as start radius messages.",
			},
		},
	}
}
