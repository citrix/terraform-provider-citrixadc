package sslfips

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslfipsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"fipsfw": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Path to the FIPS firmware file.",
			},
			"hsmlabel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Label to identify the Hardware Security Module (HSM).",
			},
			"inithsm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "FIPS initialization level. The appliance currently supports Level-2 (FIPS 140-2).",
			},
			"oldsopassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Old password for the security officer.",
			},
			"oldsopassword_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Old password for the security officer.",
			},
			"oldsopassword_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a oldsopassword_wo update.",
			},
			"sopassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Security officer password that will be in effect after you have configured the HSM.",
			},
			"sopassword_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Security officer password that will be in effect after you have configured the HSM.",
			},
			"sopassword_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a sopassword_wo update.",
			},
			"userpassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The Hardware Security Module's (HSM) User password.",
			},
			"userpassword_wo": schema.StringAttribute{
				Optional:    true,
				Description: "The Hardware Security Module's (HSM) User password.",
			},
			"userpassword_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a userpassword_wo update.",
			},
		},
	}
}
