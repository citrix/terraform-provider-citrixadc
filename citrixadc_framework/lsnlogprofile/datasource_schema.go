package lsnlogprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LsnlogprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"analyticsprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Analytics Profile attached to this lsn profile.",
			},
			"logcompact": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Logs in Compact Logging format if option is enabled.",
			},
			"logipfix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Logs in IPFIX  format if option is enabled.",
			},
			"logprofilename": schema.StringAttribute{
				Required:    true,
				Description: "The name of the logging Profile.",
			},
			"logsessdeletion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "LSN Session deletion will not be logged if disabled.",
			},
			"logsubscrinfo": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subscriber ID information is logged if option is enabled.",
			},
		},
	}
}
