package nshmackey

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NshmackeyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comments associated with this encryption key.",
			},
			"digest": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Digest (hash) function to be used in the HMAC computation.",
			},
			"keyvalue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The hex-encoded key to be used in the HMAC computation. The key can be any length (up to a Citrix ADC-imposed maximum of 255 bytes). If the length is less than the digest block size, it will be zero padded up to the block size. If it is greater than the block size, it will be hashed using the digest function to the block size. The block size for each digest is:\n   MD2    - 16 bytes\n   MD4    - 16 bytes\n   MD5    - 16 bytes\n   SHA1   - 20 bytes\n   SHA224 - 28 bytes\n   SHA256 - 32 bytes\n   SHA384 - 48 bytes\n   SHA512 - 64 bytes\nNote that the key will be encrypted when it it is saved\n\nThere is a special key value AUTO which generates a new random key for the specified digest. This kind of key is\nintended for use cases where the NetScaler both generates and verifies an HMAC on  the same data.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Key name.  This follows the same syntax rules as other expression entity names:\n   It must begin with an alpha character (A-Z or a-z) or an underscore (_).\n   The rest of the characters must be alpha, numeric (0-9) or underscores.\n   It cannot be re or xp (reserved for regular and XPath expressions).\n   It cannot be an expression reserved word (e.g. SYS or HTTP).\n   It cannot be used for an existing expression object (HTTP callout, patset, dataset, stringmap, or named expression).",
			},
		},
	}
}
