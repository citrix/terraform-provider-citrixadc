package aaagroup

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AaagroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the group. Must begin with a letter, number, or the underscore character (_), and must consist only of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore  characters. Cannot be changed after the group is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or\nsingle quotation marks (for example, \"my aaa group\" or 'my aaa group').",
			},
			"loggedin": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Display only the group members who are currently logged in. If there are large number of sessions, this command may provide partial details.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight of this group with respect to other configured aaa groups (lower the number higher the weight)",
			},
		},
	}
}
