package locationfile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LocationfileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"locationfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the location file, with or without absolute path. If the path is not included, the default path (/var/netscaler/locdb) is assumed. In a high availability setup, the static database must be stored in the same location on both NetScalers.",
			},
			"format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of the location file. Required for the NetScaler to identify how to read the location file.",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL \\(protocol, host, path, and file name\\) from where the location file will be imported.\n            NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
		},
	}
}
