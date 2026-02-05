package authenticationepaaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthenticationepaactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"csecexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "it holds the ClientSecurityExpression to be sent to the client",
			},
			"defaultepagroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the EPA check succeeds.",
			},
			"deletefiles": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String specifying the path(s) and name(s) of the files to be deleted by the endpoint analysis (EPA) tool. Multiple files to be delimited by comma",
			},
			"deviceposture": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parameter to enable/disable device posture service scan",
			},
			"killprocess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String specifying the name of a process to be terminated by the endpoint analysis (EPA) tool. Multiple processes to be delimited by comma",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the epa action. Must begin with a\n	    letter, number, or the underscore character (_), and must consist\n	    only of letters, numbers, and the hyphen (-), period (.) pound\n	    (#), space ( ), at (@), equals (=), colon (:), and underscore\n		    characters. Cannot be changed after epa action is created.The following requirement applies only to the Citrix ADC CLI:If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my aaa action\" or 'my aaa action').",
			},
			"quarantinegroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the quarantine group that is chosen when the EPA check fails\nif configured.",
			},
		},
	}
}
