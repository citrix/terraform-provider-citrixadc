package systemgroup

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SystemgroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"allowedmanagementinterface": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Allowed Management interfaces of the system users in the group. By default allowed from both API and CLI interfaces. If management interface for a group is set to API, then all users under this group will not allowed to access NS through CLI. GUI interface will come under API interface",
			},
			"daystoexpire": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Password days to expire for system groups. The daystoexpire value ranges from 30 to 255.",
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the group. Must begin with a letter, number, hash(#) or the underscore (_) character, and must contain only alphanumeric, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), and underscore characters. Cannot be changed after the group is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my group\" or 'my group').",
			},
			"promptstring": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String to display at the command-line prompt. Can consist of letters, numbers, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), underscore (_), and the following variables:\n* %u - Will be replaced by the user name.\n* %h - Will be replaced by the hostname of the Citrix ADC.\n* %t - Will be replaced by the current time in 12-hour format.\n* %T - Will be replaced by the current time in 24-hour format.\n* %d - Will be replaced by the current date.\n* %s - Will be replaced by the state of the Citrix ADC.\n\nNote: The 63-character limit for the length of the string does not apply to the characters that replace the variables.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "CLI session inactivity timeout, in seconds. If Restrictedtimeout argument of system parameter is enabled, Timeout can have values in the range [300-86400] seconds.If Restrictedtimeout argument of system parameter is disabled, Timeout can have values in the range [0, 10-100000000] seconds. Default value is 900 seconds.",
			},
			"warnpriorndays": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of days before which password expiration warning would be thrown with respect to daystoexpire. The warnpriorndays value ranges from 5 to 40.",
			},
		},
	}
}
