package snmpuser

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SnmpuserDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"authpasswd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Plain-text pass phrase to be used by the authentication algorithm specified by the authType (Authentication Type) parameter. Can consist of 8 to 63 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the pass phrase includes one or more spaces, enclose it in double or single quotation marks (for example, \"my phrase\" or 'my phrase').",
			},
			"authtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Authentication algorithm used by the Citrix ADC and the SNMPv3 user for authenticating the communication between them. You must specify the same authentication algorithm when you configure the SNMPv3 user in the SNMP manager.",
			},
			"group": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the configured SNMPv3 group to which to bind this SNMPv3 user. The access rights (bound SNMPv3 views) and security level set for this group are assigned to this user.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the SNMPv3 user. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose it in double or single quotation marks (for example, \"my user\" or 'my user').",
			},
			"privpasswd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Encryption key to be used by the encryption algorithm specified by the privType (Encryption Type) parameter. Can consist of 8 to 63 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the key includes one or more spaces, enclose it in double or single quotation marks (for example, \"my key\" or 'my key').",
			},
			"privtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Encryption algorithm used by the Citrix ADC and the SNMPv3 user for encrypting the communication between them. You must specify the same encryption algorithm when you configure the SNMPv3 user in the SNMP manager.",
			},
		},
	}
}
