package contentinspectionparameter

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ContentinspectionparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an error condition in evaluating the expression.\nAvailable settings function as follows:\n* NOINSPECTION - Do not Inspect the traffic.\n* RESET - Reset the connection and notify the user's browser, so that the user can resend the request.\n* DROP - Drop the message without sending a response to the user.",
			},
		},
	}
}
