package lsnappsprofile_port_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LsnappsprofilePortBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"appsprofilename": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN application profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN application profile is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn application profile1\" or 'lsn application profile1').",
			},
			"lsnport": schema.StringAttribute{
				Required:    true,
				Description: "Port numbers or range of port numbers to match against the destination port of the incoming packet from a subscriber. When the destination port is matched, the LSN application profile is applied for the LSN session. Separate a range of ports with a hyphen. For example, 40-90.",
			},
		},
	}
}
