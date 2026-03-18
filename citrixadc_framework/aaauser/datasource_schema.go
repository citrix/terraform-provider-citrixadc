package aaauser

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AaauserDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"loggedin": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Show whether the user is logged in or not.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password with which the user logs on. Required for any user account that does not exist on an external authentication server.\nIf you are not using an external authentication server, all user accounts must have a password. If you are using an external authentication server, you must provide a password for local user accounts that do not exist on the authentication server.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Name for the user. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the user is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or\nsingle quotation marks (for example, \"my aaa user\" or \"my aaa user\").",
			},
		},
	}
}
