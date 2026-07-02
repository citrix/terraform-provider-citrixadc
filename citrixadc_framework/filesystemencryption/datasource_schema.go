package filesystemencryption

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func FilesystemencryptionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"ntimes0flash": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of times /flash directory has to be written with 0s.",
			},
			"ntimes0var": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of times /var directory has to be written with 0s.",
			},
			"passphrase": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Encryption Passphrase.",
			},
			"passphrase_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Encryption Passphrase.",
			},
			"passphrase_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a passphrase_wo update.",
			},
			"supportedstate": schema.StringAttribute{
				Computed:    true,
				Description: "Get the supported state of File System Encryption. Possible values = DISABLED, ENABLED, UNKNOWN.",
			},
			"effectivestate": schema.StringAttribute{
				Computed:    true,
				Description: "Get the current encrypted state of the File System. Possible values = ENABLED, DISABLED.",
			},
		},
	}
}
