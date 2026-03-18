package uservserver

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func UservserverDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"params": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments associated with the protocol.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments that you might want to associate with the virtual server.",
			},
			"defaultlb": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the default Load Balancing virtual server used for load balancing of services. The protocol type of default Load Balancing virtual server should be a user type.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 or IPv6 address to assign to the virtual server.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vserver\" or 'my vserver').",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number for the virtual server.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial state of the user vserver.",
			},
			"userprotocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User protocol uesd by the service.",
			},
		},
	}
}
