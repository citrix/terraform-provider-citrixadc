package csvserver_lbvserver_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CsvserverLbvserverBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"lbvserver": schema.StringAttribute{
				Required:    true,
				Description: "Name of the default lb vserver bound. Use this param for Default binding only. For Example: bind cs vserver cs1 -lbvserver lb1",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the content switching virtual server to which the content switching policy applies.",
			},
			"targetvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The virtual server name (created with the add lb vserver command) to which content will be switched.",
			},
		},
	}
}
