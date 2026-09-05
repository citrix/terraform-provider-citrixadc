package systemautosaveparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SystemautosaveparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configure autosave feature. Available options are: DEFAULT - NetScaler decides default option for autosave feature. DISABLED - Autosave feature is disabled. ENABLED - Autosave feature is enabled. Possible values = DEFAULT, DISABLED, ENABLED",
			},
			"periodicsave": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable periodic save of autosave configuration. If enabled, saveconfig will be done periodically for all partitions including default. Possible values = ENABLED, DISABLED",
			},
			"periodicsavefrequency": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Frequency in multiple of 60 minutes for periodic save of autosave configuration. Default value is 720 minutes. Minimum value = 60, Maximum value = 7200",
			},
		},
	}
}
