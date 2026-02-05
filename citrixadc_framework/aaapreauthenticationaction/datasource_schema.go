package aaapreauthenticationaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AaapreauthenticationactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"defaultepagroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the EPA check succeeds.",
			},
			"deletefiles": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String specifying the path(s) and name(s) of the files to be deleted by the endpoint analysis (EPA) tool.",
			},
			"killprocess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String specifying the name of a process to be terminated by the endpoint analysis (EPA) tool.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the preauthentication action. Must begin with a letter, number, or the underscore character (_), and must consist only of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after preauthentication action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my aaa action\" or 'my aaa action').",
			},
			"preauthenticationaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow or deny logon after endpoint analysis (EPA) results.",
			},
		},
	}
}
