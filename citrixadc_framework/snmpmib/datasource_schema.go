package snmpmib

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SnmpmibDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"contact": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the administrator for this Citrix ADC. Along with the name, you can include information on how to contact this person, such as a phone number or an email address. Can consist of 1 to 127 characters that include uppercase and  lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the information includes one or more spaces, enclose it in double or single quotation marks (for example, \"my contact\" or 'my contact').",
			},
			"customid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom identification number for the Citrix ADC. Can consist of 1 to 127 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters. You should choose a custom identification that helps identify the Citrix ADC appliance.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the ID includes one or more spaces, enclose it in double or single quotation marks (for example, \"my ID\" or 'my ID').",
			},
			"location": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Physical location of the Citrix ADC. For example, you can specify building name, lab number, and rack number. Can consist of 1 to 127 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the location includes one or more spaces, enclose it in double or single quotation marks (for example, \"my location\" or 'my location').",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name for this Citrix ADC. Can consist of 1 to 127 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.  You should choose a name that helps identify the Citrix ADC appliance.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose it in double or single quotation marks (for example, \"my name\" or 'my name').",
			},
			"ownernode": schema.Int64Attribute{
				Required:    true,
				Description: "ID of the cluster node for which we are setting the mib. This is a mandatory argument to set snmp mib on CLIP.",
			},
		},
	}
}
