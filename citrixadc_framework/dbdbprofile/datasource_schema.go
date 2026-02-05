package dbdbprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DbdbprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"conmultiplex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the same server-side connection for multiple client-side requests. Default is enabled.",
			},
			"enablecachingconmuxoff": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable caching when connection multiplexing is OFF.",
			},
			"interpretquery": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If ENABLED, inspect the query and update the connection information, if required. If DISABLED, forward the query to the server.",
			},
			"kcdaccount": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the KCD account that is used for Windows authentication.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the database profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.\n	    CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my profile\" or 'my profile').",
			},
			"stickiness": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If the queries are related to each other, forward to the same backend server.",
			},
		},
	}
}
