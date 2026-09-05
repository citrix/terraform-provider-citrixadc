package policyparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func PolicyparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"maxeventsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum event size in kilobytes that the policy engine will process. When event data exceeds this limit, the action specified by maxEventSizeExceedAction is taken. This parameter helps prevent resource exhaustion from processing extremely large events.",
			},
			"maxeventsizeexceedaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to take when event data exceeds maxEventSize:",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum time in milliseconds to allow for processing expressions and policies without interruption. If the timeout is reached then the evaluation causes an UNDEF to be raised and no further processing is performed.",
			},
		},
	}
}
