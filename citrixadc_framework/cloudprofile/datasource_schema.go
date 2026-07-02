package cloudprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CloudprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"azurepollperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Azure polling period (in seconds)",
			},
			"azuretagname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Azure tag name",
			},
			"azuretagvalue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Azure tag value",
			},
			"boundservicegroupsvctype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The type of bound service",
			},
			"delay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which all the services configured on the server are disabled.",
			},
			"graceful": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates graceful shutdown of the service. System will wait for all outstanding connections to this service to be closed before disabling the service.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 or IPv6 address to assign to the virtual server.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Cloud profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the profile is created.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number for the virtual server.",
			},
			"servicegroupname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "servicegroups bind to this server",
			},
			"servicetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol used by the service (also called the service type).",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of cloud profile that you want to create, Vserver or based on Azure Tags",
			},
			"vservername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vserver\" or 'my vserver').",
			},
			"vsvrbindsvcport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The port number to be used for the bound service.",
			},
		},
	}
}