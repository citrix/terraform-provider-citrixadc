package responderparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ResponderparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum time in milliseconds to allow for processing all the policies and their selected actions without interruption. If the timeout is reached then the evaluation causes an UNDEF to be raised and no further processing is performed.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform when policy evaluation creates an UNDEF condition. Available settings function as follows:\n* NOOP - Send the request to the protected server.\n* RESET - Reset the request and notify the user's browser, so that the user can resend the request.\n* DROP - Drop the request without sending a response to the user.",
			},
		},
	}
}
