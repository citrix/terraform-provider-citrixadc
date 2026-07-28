package systemsession

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SystemsessionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"sid": schema.Int64Attribute{
				Required:    true,
				Description: "ID of the system session about which to display information.",
			},
			// Read-only session fields returned by GET.
			"username": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the user who is logged in.",
			},
			"logintime": schema.StringAttribute{
				Computed:    true,
				Description: "Time when the user logged in.",
			},
			"logintimelocal": schema.StringAttribute{
				Computed:    true,
				Description: "Time (local) when the user logged in.",
			},
			"lastactivitytime": schema.StringAttribute{
				Computed:    true,
				Description: "Time of last activity in the session.",
			},
			"lastactivitytimelocal": schema.StringAttribute{
				Computed:    true,
				Description: "Time (local) of last activity in the session.",
			},
			"expirytime": schema.StringAttribute{
				Computed:    true,
				Description: "Time when the session expires.",
			},
			"numofconnections": schema.StringAttribute{
				Computed:    true,
				Description: "Number of connections in the session.",
			},
			"currentconn": schema.StringAttribute{
				Computed:    true,
				Description: "Indicates the current connection.",
			},
			"clienttype": schema.StringAttribute{
				Computed:    true,
				Description: "Type of client used for the session.",
			},
			"partitionname": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the partition for the session.",
			},
			"clientipaddress": schema.StringAttribute{
				Computed:    true,
				Description: "Client IP address for the session.",
			},
		},
	}
}
