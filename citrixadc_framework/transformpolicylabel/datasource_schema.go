package transformpolicylabel

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func TransformpolicylabelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the URL Transformation policy label is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, my transform policylabel or my transform policylabel).",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the policy label.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, my transform policylabel or my transform policylabel).",
			},
			"policylabeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Types of transformations allowed by the policies bound to the label. For URL transformation, always http_req (HTTP Request).",
			},
		},
	}
}
