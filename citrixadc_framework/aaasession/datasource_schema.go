package aaasession

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// AaasessionDataSourceSchema is a read-only, get(all)-backed filter datasource.
// All AAA-session selectors are optional filters; the datasource returns the
// first session matching the supplied filters. nodeid is the GET-only cluster
// filter and is exposed here (it is valid for get, unlike the kill payload).
func AaasessionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"all": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Terminate all active AAA-TM/VPN sessions.",
			},
			"groupname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the AAA group.",
			},
			"iip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address or the first address in the intranet IP range.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet mask for the intranet IP range.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"sessionkey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Show aaa session associated with given session key",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the AAA user.",
			},
		},
	}
}
