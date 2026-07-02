package sslkeyfile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslkeyfileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name to assign to the imported key file. Must begin with an ASCII alphanumeric or underscore(_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@),equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my file\" or 'my file').",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"password_wo": schema.StringAttribute{
				Optional:    true,
				Description: "0",
			},
			"password_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a password_wo update.",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL specifying the protocol, host, and path, including file name, to the key file to be imported. For example, http://www.example.com/key_file.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access, and the issuer certificate of the HTTPS server is not present in the specific path on NetScaler to authenticate the HTTPS server.",
			},
		},
	}
}
