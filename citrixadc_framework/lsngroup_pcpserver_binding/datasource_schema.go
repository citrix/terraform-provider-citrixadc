package lsngroup_pcpserver_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LsngroupPcpserverBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn group1\" or 'lsn group1').",
			},
			"pcpserver": schema.StringAttribute{
				Required:    true,
				Description: "Name of the PCP server to be associated with lsn group.",
			},
		},
	}
}
