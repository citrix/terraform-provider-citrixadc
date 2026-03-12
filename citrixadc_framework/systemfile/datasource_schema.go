package systemfile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SystemfileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"filecontent": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "file content in Base64 format.",
			},
			"fileencoding": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "encoding type of the file content.",
			},
			"filelocation": schema.StringAttribute{
				Required:    true,
				Description: "location of the file on Citrix ADC.",
			},
			"filename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the file. It should not include filepath.",
			},
		},
	}
}
