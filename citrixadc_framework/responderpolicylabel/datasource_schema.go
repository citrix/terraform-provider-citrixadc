package responderpolicylabel

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ResponderpolicylabelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this responder policy label.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the responder policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the responder policy label is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my responder policy label\" or my responder policy label').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the responder policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"policylabeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of responses sent by the policies bound to this policy label. Types are:\n* HTTP - HTTP responses.\n* OTHERTCP - NON-HTTP TCP responses.\n* SIP_UDP - SIP responses.\n* RADIUS - RADIUS responses.\n* MYSQL - SQL responses in MySQL format.\n* MSSQL - SQL responses in Microsoft SQL format.\n* NAT - NAT response.\n* MQTT - Trigger policies bind with MQTT type.\n* MQTT_JUMBO - Trigger policies bind with MQTT Jumbo type.",
			},
		},
	}
}
