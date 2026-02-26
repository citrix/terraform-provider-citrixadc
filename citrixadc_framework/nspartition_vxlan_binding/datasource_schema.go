package nspartition_vxlan_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NspartitionVxlanBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"partitionname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Partition. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"vxlan": schema.Int64Attribute{
				Required:    true,
				Description: "Identifier of the vxlan that is assigned to this partition.",
			},
		},
	}
}
