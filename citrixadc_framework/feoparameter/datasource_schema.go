package feoparameter

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func FeoparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cssinlinethressize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Threshold value of the file size (in bytes) for converting external CSS files to inline CSS files.",
			},
			"imginlinethressize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum file size of an image (in bytes), for coverting linked images to inline images.",
			},
			"jpegqualitypercent": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The percentage value of a JPEG image quality to be reduced. Range: 0 - 100",
			},
			"jsinlinethressize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Threshold value of the file size (in bytes), for converting external JavaScript files to inline JavaScript files.",
			},
		},
	}
}
