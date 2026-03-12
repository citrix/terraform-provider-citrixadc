package rewriteparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func RewriteparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum time in milliseconds to allow for processing all the policies and their selected actions without interruption. If the timeout is reached then the evaluation causes an UNDEF to be raised and no further processing is performed. Note that some rewrites may have already been performed.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an error condition in evaluating the expression.\nAvailable settings function as follows:\n* NOREWRITE - Do not modify the message.\n* RESET - Reset the connection and notify the user's browser, so that the user can resend the request.\n* DROP - Drop the message without sending a response to the user.",
			},
		},
	}
}
