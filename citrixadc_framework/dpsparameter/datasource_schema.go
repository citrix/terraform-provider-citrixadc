package dpsparameter

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DpsparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"customerid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Customer ID of the Citrix Cloud customer.",
			},
			"deployment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Describes if the customer is connecting to Commerical/JapanCloud/Gov Citrix Cloud customer. Possible values = COMMERCIAL, GOV, JAPANCLOUD",
			},
			"serviceurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Service URL of the Citrix Cloud customer.",
			},
		},
	}
}
