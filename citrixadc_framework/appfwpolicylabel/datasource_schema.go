package appfwpolicylabel

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwpolicylabelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the policy label is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy label\" or 'my policy label').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new name of the application firewall policylabel.",
			},
			"policylabeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of transformations allowed by the policies bound to the label. Always http_req for application firewall policy labels.",
			},
		},
	}
}
