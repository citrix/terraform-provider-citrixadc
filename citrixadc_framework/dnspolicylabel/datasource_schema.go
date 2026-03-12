package dnspolicylabel

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DnspolicylabelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the dns policy label.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new name of the dns policylabel.",
			},
			"transform": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The type of transformations allowed by the policies bound to the label.",
			},
		},
	}
}
