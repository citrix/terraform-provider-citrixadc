package policydataset

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func PolicydatasetDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this dataset or a data bound to this dataset.",
			},
			"dynamic": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is used to populate internal dataset information so that the dataset can also be used dynamically in an expression. Here dynamically means the dataset name can also be derived using an expression. For example for a given dataset name \"allow_test\" it can be used dynamically as client.ip.src.equals_any(\"allow_\" + http.req.url.path.get(1)). This cannot be used with default datasets.",
			},
			"dynamiconly": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Shows only dynamic datasets when set true.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the dataset. Must not exceed 127 characters.",
			},
			"patsetfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "File which contains list of patterns that needs to be bound to the dataset. A patsetfile cannot be associated with multiple datasets.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of value to bind to the dataset.",
			},
		},
	}
}
