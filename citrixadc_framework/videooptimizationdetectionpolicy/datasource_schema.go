package videooptimizationdetectionpolicy

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VideooptimizationdetectionpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the videooptimization detection action to perform if the request matches this videooptimization detection policy. Built-in actions should be used. These are:\n* DETECT_CLEARTEXT_PD - Cleartext PD is detected and increment related counters.\n* DETECT_CLEARTEXT_ABR - Cleartext ABR is detected and increment related counters.\n* DETECT_ENCRYPTED_ABR - Encrypted ABR is detected and increment related counters.\n* TRIGGER_ENC_ABR_DETECTION - This is potentially encrypted ABR. Internal traffic heuristics algorithms will further process traffic to confirm detection.\n* TRIGGER_CT_ABR_BODY_DETECTION -  This is potentially cleartext ABR. Internal traffic heuristics algorithms will further process traffic to confirm detection.\n* RESET - Reset the client connection by closing it.\n* DROP - Drop the connection without sending a response.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any type of information about this videooptimization detection policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the messagelog action to use for requests that match this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the videooptimization detection policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.Can be modified, removed or renamed.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the videooptimization detection policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that determines which request or response match the video optimization detection policy.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.",
			},
		},
	}
}
