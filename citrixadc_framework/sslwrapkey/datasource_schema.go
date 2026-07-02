package sslwrapkey

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslwrapkeyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password string for the wrap key.",
			},
			"password_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Password string for the wrap key.",
			},
			"password_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a password_wo update.",
			},
			"salt": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Salt string for the wrap key.",
			},
			"salt_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Salt string for the wrap key.",
			},
			"salt_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a salt_wo update.",
			},
			"wrapkeyname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the wrap key. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the wrap key is created.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my key\" or 'my key').",
			},
		},
	}
}
