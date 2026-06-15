package systemsshkey

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SystemsshkeyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "URL \\(protocol, host, path, and file name\\) from where the location file will be imported.\n            NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Description: "URL \\(protocol, host, path, and file name\\) from where the location file will be imported.\n            NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
			"sshkeytype": schema.StringAttribute{
				Required:    true,
				Description: "The type of the ssh key whether public or private key",
			},
		},
	}
}
