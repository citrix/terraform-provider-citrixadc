package lsnpool_lsnip_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LsnpoolLsnipBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"lsnip": schema.StringAttribute{
				Required:    true,
				Description: "IPv4 address or a range of IPv4 addresses to be used as NAT IP address(es) for LSN.\nAfter the pool is created, these IPv4 addresses are added to the Citrix ADC as Citrix ADC owned IP address of type LSN. A maximum of 4096 IP addresses can be bound to an LSN pool. An LSN IP address associated with an LSN pool cannot be shared with other LSN pools. IP addresses specified for this parameter must not already exist on the Citrix ADC as any Citrix ADC owned IP addresses. In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189. You can later remove some or all the LSN IP addresses from the pool, and add IP addresses to the LSN pool.\nBy default , arp is enabled on LSN IP address but, you can disable it using command - \"set ns ip\"",
			},
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Description: "ID(s) of cluster node(s) on which command is to be executed",
			},
			"poolname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN pool. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN pool is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn pool1\" or 'lsn pool1').",
			},
		},
	}
}
