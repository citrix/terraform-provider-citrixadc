package sslechconfig

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslechconfigDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"echcipher": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The supported cipher suite that encrypts the client Hello Message.",
			},
			"echconfigid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The config id of the ech config.",
			},
			"echconfigname": schema.StringAttribute{
				Required:    true,
				Description: "The ECH config name configured.",
			},
			"echpublicname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The public name of ech config means FQDN or any string",
			},
			"hpkekeyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the configured HPKE key",
			},
			"version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The version of ECH for which this configuration is used.",
			},
		},
	}
}
