package systembackup

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SystembackupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comment specified at the time of creation of the backup file(*.tgz).",
			},
			"filename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the backup file(*.tgz) to be restored.",
			},
			"includekernel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this option to add kernel in the backup file",
			},
			"level": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Level of data to be backed up.",
			},
			"skipbackup": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this option to skip taking backup during restore operation",
			},
			"uselocaltimezone": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option will create backup file with local timezone timestamp",
			},
		},
	}
}
