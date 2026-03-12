package pcpprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func PcpprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"announcemulticount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that identify the number announce message to be send.",
			},
			"mapping": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This argument is for enabling/disabling the MAP opcode  of current PCP Profile",
			},
			"maxmaplife": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that identify the maximum mapping lifetime (in seconds) for a pcp profile. default(86400s = 24Hours).",
			},
			"minmaplife": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that identify the minimum mapping lifetime (in seconds) for a pcp profile. default(120s)",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the PCP Profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my pcpProfile\" or my pcpProfile).",
			},
			"peer": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This argument is for enabling/disabling the PEER opcode of current PCP Profile",
			},
			"thirdparty": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This argument is for enabling/disabling the THIRD PARTY opcode of current PCP Profile",
			},
		},
	}
}
