package nsaptlicense

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsaptlicenseDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "License ID",
			},
			"bindtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bind type",
			},
			"countavailable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The user can allocate one or more licenses. Ensure the value is less than (for partial allocation) or equal to the total number of available licenses",
			},
			"licensedir": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "License Directory",
			},
			"serialno": schema.StringAttribute{
				Required:    true,
				Description: "Hardware Serial Number/License Activation Code(LAC)",
			},
			"sessionid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Session ID",
			},
			"useproxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies whether to use the licenseproxyserver to reach the internet. Make sure to configure licenseproxyserver to use this option.",
			},
		},
	}
}
