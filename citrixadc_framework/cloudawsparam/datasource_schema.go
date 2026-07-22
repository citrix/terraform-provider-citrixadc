package cloudawsparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CloudawsparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"rolearn": schema.StringAttribute{
				Computed:    true,
				Description: "IAM Role ARN",
			},
		},
	}
}
