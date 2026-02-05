package appfwconfidfield

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwconfidfieldDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about the form field designation.",
			},
			"fieldname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the form field to designate as confidential.",
			},
			"isregex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Method of specifying the form field name. Available settings function as follows:\n* REGEX. Form field is a regular expression.\n* NOTREGEX. Form field is a literal string.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the confidential field designation.",
			},
			"url": schema.StringAttribute{
				Required:    true,
				Description: "URL of the web page that contains the web form.",
			},
		},
	}
}
