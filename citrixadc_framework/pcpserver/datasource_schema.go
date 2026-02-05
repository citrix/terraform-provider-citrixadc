package pcpserver

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func PcpserverDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The IP address of the PCP server.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the PCP server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my pcpServer\" or my pcpServer).",
			},
			"pcpprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "pcp profile name",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number for the PCP server.",
			},
		},
	}
}
