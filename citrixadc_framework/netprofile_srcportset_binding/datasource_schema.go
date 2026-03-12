package netprofile_srcportset_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NetprofileSrcportsetBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the netprofile to which to bind port ranges.",
			},
			"srcportrange": schema.StringAttribute{
				Required:    true,
				Description: "When the source port range is configured and associated with the netprofile bound to a service group, Citrix ADC will choose a port from the range configured for connection establishment at the backend servers.",
			},
		},
	}
}
