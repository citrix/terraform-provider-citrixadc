package cmpaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CmpactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"addvaryheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Control insertion of the Vary header in HTTP responses compressed by Citrix ADC. Intermediate caches store different versions of the response for different values of the headers present in the Vary response header.",
			},
			"cmptype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of compression performed by this action.\nAvailable settings function as follows:\n* COMPRESS - Apply GZIP or DEFLATE compression to the response, depending on the request header. Prefer GZIP.\n* GZIP - Apply GZIP compression.\n* DEFLATE - Apply DEFLATE compression.\n* NOCOMPRESS - Do not compress the response if the request matches a policy that uses this action.",
			},
			"deltatype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The type of delta action (if delta type compression action is defined).",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the compression action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the action is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cmp action\" or 'my cmp action').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the compression action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters.\nChoose a name that can be correlated with the function that the action performs.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cmp action\" or 'my cmp action').",
			},
			"varyheadervalue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The value of the HTTP Vary header for compressed responses.",
			},
		},
	}
}
