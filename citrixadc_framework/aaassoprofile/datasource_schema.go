package aaassoprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AaassoprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the SSO Profile. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after a SSO Profile is created.\n\nThe following requirement applies only to the NetScaler CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password with which the user logs on. Required for Single sign on to  external server.",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name for the user. Must begin with a letter, number, or the underscore (_) character, and must contain only alphanumeric, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), and underscore characters.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my group\" or 'my group').",
			},
		},
	}
}
