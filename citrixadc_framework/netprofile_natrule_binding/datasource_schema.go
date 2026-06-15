package netprofile_natrule_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NetprofileNatruleBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the netprofile to which to bind port ranges.",
			},
			"natrule": schema.StringAttribute{
				Required:    true,
				Description: "IPv4 network address on whose traffic you want the Citrix ADC to do rewrite ip prefix.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"rewriteip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
		},
	}
}
