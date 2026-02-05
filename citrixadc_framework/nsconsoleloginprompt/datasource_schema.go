package nsconsoleloginprompt

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsconsoleloginpromptDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"promptstring": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Console login prompt string",
			},
		},
	}
}
