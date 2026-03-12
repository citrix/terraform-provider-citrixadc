package nsxmlnamespace

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsxmlnamespaceDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"namespace": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expanded namespace for which the XML prefix is provided.",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Description for the prefix.",
			},
			"prefix": schema.StringAttribute{
				Required:    true,
				Description: "XML prefix.",
			},
		},
	}
}
