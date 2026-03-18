package nsrpcnode

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsrpcnodeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ipaddress": schema.StringAttribute{
				Required:    true,
				Description: "IP address of the node. This has to be in the same subnet as the NSIP address.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password to be used in authentication with the peer system node.",
			},
			"secure": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of the channel when talking to the node.",
			},
			"srcip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source IP address to be used to communicate with the peer system node. The default value is 0, which means that the appliance uses the NSIP address as the source IP address.",
			},
			"validatecert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "validate the server certificate for secure SSL connections",
			},
		},
	}
}
