package appfwarchive

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwarchiveDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comments associated with this archive.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of tar archive",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates the source of the tar archive file as a URL\nof the form\n\n    <protocol>://<host>[:<port>][/<path>]\n\n<protocol> is http or https.\n<host> is the DNS name or IP address of the http or https server.\n<port> is the port number of the server. If omitted, the\ndefault port for http or https will be used.\n<path> is the path of the file on the server.\n\nImport will fail if an https server requires client\ncertificate authentication.",
			},
			"target": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Path to the file to be exported",
			},
		},
	}
}
