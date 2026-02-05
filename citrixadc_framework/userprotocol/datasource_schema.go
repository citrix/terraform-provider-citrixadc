package userprotocol

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func UserprotocolDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments associated with the protocol.",
			},
			"extension": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the extension to add parsing and runtime handling of the protocol packets.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Unique name for the user protocol. Not case sensitive. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters.",
			},
			"transport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Transport layer's protocol.",
			},
		},
	}
}
