package lsnclient

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LsnclientDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"clientname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN client entity. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN client is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn client1\" or 'lsn client1').",
			},
		},
	}
}
