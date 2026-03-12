package interfacepair

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func InterfacepairDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"interface_id": schema.Int64Attribute{
				Required:    true,
				Description: "The Interface pair id",
			},
			"ifnum": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "The constituent interfaces in the interface pair",
			},
		},
	}
}
