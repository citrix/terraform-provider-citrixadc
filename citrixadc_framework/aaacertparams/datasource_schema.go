package aaacertparams

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AaacertparamsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"groupnamefield": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Client certificate field that specifies the group, in the format <field>:<subfield>.",
			},
			"usernamefield": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Client certificate field that contains the username, in the format <field>:<subfield>.",
			},
		},
	}
}
