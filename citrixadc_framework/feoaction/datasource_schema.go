package feoaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func FeoactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cachemaxage": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maxage for cache extension.",
			},
			"clientsidemeasurements": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send AppFlow records about the web pages optimized by this action. The records provide FEO statistics, such as the number of HTTP requests that have been reduced for this page. You must enable the Appflow feature before enabling this parameter.",
			},
			"convertimporttolink": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Convert CSS import statements to HTML link tags.",
			},
			"csscombine": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Combine one or more CSS files into one file.",
			},
			"cssimginline": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Inline small images (less than 2KB) referred within CSS files as background-URLs",
			},
			"cssinline": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Inline CSS files, whose size is less than 2KB, within the main page.",
			},
			"cssminify": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Remove comments and whitespaces from CSSs.",
			},
			"cssmovetohead": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Move any CSS file present within the body tag of an HTML page to the head tag.",
			},
			"dnsshards": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Set of domain names that replaces the parent domain.",
			},
			"domainsharding": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain name of the server",
			},
			"htmlminify": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Remove comments and whitespaces from an HTML page.",
			},
			"imggiftopng": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Convert GIF image formats to PNG formats.",
			},
			"imginline": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Inline images whose size is less than 2KB.",
			},
			"imglazyload": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Download images, only when the user scrolls the page to view them.",
			},
			"imgshrinktoattrib": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Shrink image dimensions as per the height and width attributes specified in the <img> tag.",
			},
			"imgtojpegxr": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Convert JPEG, GIF, PNG image formats to JXR format.",
			},
			"imgtowebp": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Convert JPEG, GIF, PNG image formats to WEBP format.",
			},
			"jpgoptimize": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Remove non-image data such as comments from JPEG images.",
			},
			"jsinline": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Convert linked JavaScript files (less than 2KB) to inline JavaScript files.",
			},
			"jsminify": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Remove comments and whitespaces from JavaScript.",
			},
			"jsmovetoend": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Move any JavaScript present in the body tag to the end of the body tag.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the front end optimization action.",
			},
			"pageextendcache": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Extend the time period during which the browser can use the cached resource.",
			},
		},
	}
}
