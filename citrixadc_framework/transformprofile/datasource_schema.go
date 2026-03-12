package transformprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func TransformprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this URL Transformation profile.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the URL transformation profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the URL transformation profile is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, my transform profile or my transform profile).",
			},
			"onlytransformabsurlinbody": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "In the HTTP body, transform only absolute URLs. Relative URLs are ignored.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of transformation. Always URL for URL Transformation profiles.",
			},
		},
	}
}
