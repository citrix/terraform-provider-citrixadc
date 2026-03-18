package gslbvserver_gslbservicegroup_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func GslbvserverGslbservicegroupBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server on which to perform the binding operation.",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order number to be assigned to the service when it is bound to the lb vserver.",
			},
			"servicegroupname": schema.StringAttribute{
				Required:    true,
				Description: "The GSLB service group name bound to the selected GSLB virtual server.",
			},
		},
	}
}
