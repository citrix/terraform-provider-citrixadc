package lbpolicy

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LbpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of action to use if the request matches this LB policy.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any type of information about this LB policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the messagelog action to use for requests that match this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the LB policy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the LB policy is added.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my lb policy\" or 'my lb policy').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the LB policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my lb policy\" or 'my lb policy').",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression against which traffic is evaluated.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Available settings function as follows:\n* NOLBACTION - Does not consider LB actions in making LB decision.\n* RESET - Reset the request and notify the user, so that the user can resend the request.\n* DROP - Drop the request without sending a response to the user.",
			},
		},
	}
}
