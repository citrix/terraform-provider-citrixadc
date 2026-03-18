package snmpgroup

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SnmpgroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the SNMPv3 group. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.  You should choose a name that helps identify the SNMPv3 group. \n            \nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose it in double or single quotation marks (for example, \"my name\" or 'my name').",
			},
			"readviewname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the configured SNMPv3 view that you want to bind to this SNMPv3 group. An SNMPv3 user bound to this group can access the subtrees that are bound to this SNMPv3 view as type INCLUDED, but cannot access the ones that are type EXCLUDED. If the Citrix ADC has multiple SNMPv3 view entries with the same name, all such entries are associated with the SNMPv3 group.",
			},
			"securitylevel": schema.StringAttribute{
				Required:    true,
				Description: "Security level required for communication between the Citrix ADC and the SNMPv3 users who belong to the group. Specify one of the following options:\nnoAuthNoPriv. Require neither authentication nor encryption.\nauthNoPriv. Require authentication but no encryption.\nauthPriv. Require authentication and encryption.\nNote: If you specify authentication, you must specify an encryption algorithm when you assign an SNMPv3 user to the group. If you also specify encryption, you must assign both an authentication and an encryption algorithm for each group member.",
			},
		},
	}
}
