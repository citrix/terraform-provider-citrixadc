package lsngroup_lsntransportprofile_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LsngroupLsntransportprofileBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn group1\" or 'lsn group1').",
			},
			"transportprofilename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the LSN transport profile to bind to the specified LSN group. Bind a profile for each protocol for which you want to specify settings.\n\nBy default, one LSN transport profile with default settings for TCP, UDP, and ICMP protocols is bound to an LSN group during its creation. This profile is called a default transport.\n\nAn LSN transport profile that you bind to an LSN group overrides the default LSN transport profile for that protocol.",
			},
		},
	}
}
