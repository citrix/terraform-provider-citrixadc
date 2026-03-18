package responderhtmlpage

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ResponderhtmlpageDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cacertfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CA certificate file name which will be used to verify the peer's certificate. The certificate should be imported using \"import ssl certfile\" CLI command or equivalent in API or GUI. If certificate name is not configured, then default root CA certificates are used for peer's certificate verification.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about the HTML page object.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name to assign to the HTML page object on the Citrix ADC.",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Overwrites the existing file",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Local path or URL (protocol, host, path, and file name) for the file from which to retrieve the imported HTML page.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
		},
	}
}
