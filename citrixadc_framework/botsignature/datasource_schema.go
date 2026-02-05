package botsignature

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func BotsignatureDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about the signature file object.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name to assign to the bot signature file object on the Citrix ADC.",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Overwrites the existing file",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Local path to and name of, or URL (protocol, host, path, and file name) for, the file in which to store the imported signature file.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
		},
	}
}
