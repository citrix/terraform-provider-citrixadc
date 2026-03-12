package lacp

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LacpDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ownernode": schema.Int64Attribute{
				Required:    true,
				Description: "The owner node in a cluster for which we want to set the lacp priority. Owner node can vary from 0 to 31. Ownernode value of 254 is used for Cluster.",
			},
			"syspriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority number that determines which peer device of an LACP LA channel can have control over the LA channel. This parameter is globally applied to all LACP channels on the Citrix ADC. The lower the number, the higher the priority.",
			},
		},
	}
}
