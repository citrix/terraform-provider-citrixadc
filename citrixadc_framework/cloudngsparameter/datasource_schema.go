package cloudngsparameter

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CloudngsparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"allowdtls12": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables DTLS1.2 for client connections on CGS",
			},
			"allowedudtversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables the required UDT version to EDT connections in the CGS deployment",
			},
			"blockonallowedngstktprof": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables blocking connections authenticated with a ticket createdby by an entity not whitelisted in allowedngstktprofile",
			},
			"csvserverticketingdecouple": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables Decoupling CSVSERVER state from Ticketing Service state in the CGS deployment",
			},
		},
	}
}
