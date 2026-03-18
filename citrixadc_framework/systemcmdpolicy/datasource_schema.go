package systemcmdpolicy

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SystemcmdpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform when a request matches the policy.",
			},
			"cmdspec": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Regular expression specifying the data that matches the policy.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name for a command policy. Must begin with a letter, number, or the underscore (_) character, and must contain only alphanumeric, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), and underscore characters. Cannot be changed after the policy is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy').",
			},
		},
	}
}
