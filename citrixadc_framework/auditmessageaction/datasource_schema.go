package auditmessageaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuditmessageactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bypasssafetycheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bypass the safety check and allow unsafe expressions.",
			},
			"loglevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Audit log level, which specifies the severity level of the log message being generated..\nThe following loglevels are valid:\n* EMERGENCY - Events that indicate an immediate crisis on the server.\n* ALERT - Events that might require action.\n* CRITICAL - Events that indicate an imminent server crisis.\n* ERROR - Events that indicate some type of error.\n* WARNING - Events that require action in the near future.\n* NOTICE - Events that the administrator should know about.\n* INFORMATIONAL - All but low-level events.\n* DEBUG - All events, in extreme detail.",
			},
			"logtonewnslog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send the message to the new nslog.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the audit message action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the message action is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my message action\" or 'my message action').",
			},
			"stringbuilderexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Default-syntax expression that defines the format and content of the log message.",
			},
		},
	}
}
