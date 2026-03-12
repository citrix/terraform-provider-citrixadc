package nspartition_vlan_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NspartitionVlanBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"partitionname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Partition. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"vlan": schema.Int64Attribute{
				Required:    true,
				Description: "Identifier of the vlan that is assigned to this partition.",
			},
		},
	}
}
