package feoaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// FeoactionDataSourceModel is the data-source-specific model, decoupled from
// FeoactionResourceModel. A data source is a pure read surface, so it can expose
// the FULL GET projection: the read/write attributes (as Computed outputs) AND
// the read-only attributes the resource deliberately omits (image/HTML/JS
// optimization flags, hit counters, builtin, feature).
type FeoactionDataSourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"` // Required lookup key

	// Read/write attributes, surfaced here as Computed outputs.
	Cachemaxage            types.Int64  `tfsdk:"cachemaxage"`
	Clientsidemeasurements types.Bool   `tfsdk:"clientsidemeasurements"`
	Convertimporttolink    types.Bool   `tfsdk:"convertimporttolink"`
	Csscombine             types.Bool   `tfsdk:"csscombine"`
	Cssimginline           types.Bool   `tfsdk:"cssimginline"`
	Cssinline              types.Bool   `tfsdk:"cssinline"`
	Cssminify              types.Bool   `tfsdk:"cssminify"`
	Cssmovetohead          types.Bool   `tfsdk:"cssmovetohead"`
	Dnsshards              types.List   `tfsdk:"dnsshards"`
	Domainsharding         types.String `tfsdk:"domainsharding"`
	Htmlminify             types.Bool   `tfsdk:"htmlminify"`
	Imggiftopng            types.Bool   `tfsdk:"imggiftopng"`
	Imginline              types.Bool   `tfsdk:"imginline"`
	Imglazyload            types.Bool   `tfsdk:"imglazyload"`
	Imgshrinktoattrib      types.Bool   `tfsdk:"imgshrinktoattrib"`
	Imgtojpegxr            types.Bool   `tfsdk:"imgtojpegxr"`
	Imgtowebp              types.Bool   `tfsdk:"imgtowebp"`
	Jpgoptimize            types.Bool   `tfsdk:"jpgoptimize"`
	Jsinline               types.Bool   `tfsdk:"jsinline"`
	Jsminify               types.Bool   `tfsdk:"jsminify"`
	Jsmovetoend            types.Bool   `tfsdk:"jsmovetoend"`
	Pageextendcache        types.Bool   `tfsdk:"pageextendcache"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/feoaction.json). Never settable; populated from GET.
	Imgadddimensions     types.Bool   `tfsdk:"imgadddimensions"`
	Imgshrinkformobile   types.Bool   `tfsdk:"imgshrinkformobile"`
	Imgweaken            types.Bool   `tfsdk:"imgweaken"`
	Jpgprogressive       types.Bool   `tfsdk:"jpgprogressive"`
	Cssflattenimports    types.Bool   `tfsdk:"cssflattenimports"`
	Jscombine            types.Bool   `tfsdk:"jscombine"`
	Htmlrmdefaultattribs types.Bool   `tfsdk:"htmlrmdefaultattribs"`
	Htmlrmattribquotes   types.Bool   `tfsdk:"htmlrmattribquotes"`
	Htmltrimurls         types.Bool   `tfsdk:"htmltrimurls"`
	Hits                 types.Int64  `tfsdk:"hits"`
	Undefhits            types.Int64  `tfsdk:"undefhits"`
	Builtin              types.List   `tfsdk:"builtin"`
	Feature              types.String `tfsdk:"feature"`
}

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

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"imgadddimensions": schema.BoolAttribute{
				Computed:    true,
				Description: "Add dimension attributes to images, if not specified within the <img> tag.",
			},
			"imgshrinkformobile": schema.BoolAttribute{
				Computed:    true,
				Description: "Serve smaller images for mobile users.",
			},
			"imgweaken": schema.BoolAttribute{
				Computed:    true,
				Description: "Reduce the image quality.",
			},
			"jpgprogressive": schema.BoolAttribute{
				Computed:    true,
				Description: "Convert JPEG image formats to progressive formats.",
			},
			"cssflattenimports": schema.BoolAttribute{
				Computed:    true,
				Description: "Replace CSS import statements with the file content.",
			},
			"jscombine": schema.BoolAttribute{
				Computed:    true,
				Description: "Combine one or more JavaScript files into one file.",
			},
			"htmlrmdefaultattribs": schema.BoolAttribute{
				Computed:    true,
				Description: "Remove default redundant attributes from an HTML file.",
			},
			"htmlrmattribquotes": schema.BoolAttribute{
				Computed:    true,
				Description: "Remove unnecessary quotes present within the HTML attributes.",
			},
			"htmltrimurls": schema.BoolAttribute{
				Computed:    true,
				Description: "Trim URLs.",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of times the action has been taken.",
			},
			"undefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "Total number of undefined policy hits.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if front end optimization action is built-in or not. Possible values: [ MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL ].",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// feoactionDataSourceSetAttrFromGet projects a NITRO feoaction GET response onto
// the data-source model. Attributes are simply filled from the GET (or left Null
// when the GET omits them) via the shared utils.MapGet* helpers.
func feoactionDataSourceSetAttrFromGet(ctx context.Context, data *FeoactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In feoactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Cachemaxage = utils.MapGetInt64(g, "cachemaxage")
	data.Clientsidemeasurements = utils.MapGetBool(g, "clientsidemeasurements")
	data.Convertimporttolink = utils.MapGetBool(g, "convertimporttolink")
	data.Csscombine = utils.MapGetBool(g, "csscombine")
	data.Cssimginline = utils.MapGetBool(g, "cssimginline")
	data.Cssinline = utils.MapGetBool(g, "cssinline")
	data.Cssminify = utils.MapGetBool(g, "cssminify")
	data.Cssmovetohead = utils.MapGetBool(g, "cssmovetohead")
	data.Dnsshards = utils.MapGetStringList(g, "dnsshards")
	data.Domainsharding = utils.MapGetString(g, "domainsharding")
	data.Htmlminify = utils.MapGetBool(g, "htmlminify")
	data.Imggiftopng = utils.MapGetBool(g, "imggiftopng")
	data.Imginline = utils.MapGetBool(g, "imginline")
	data.Imglazyload = utils.MapGetBool(g, "imglazyload")
	data.Imgshrinktoattrib = utils.MapGetBool(g, "imgshrinktoattrib")
	data.Imgtojpegxr = utils.MapGetBool(g, "imgtojpegxr")
	data.Imgtowebp = utils.MapGetBool(g, "imgtowebp")
	data.Jpgoptimize = utils.MapGetBool(g, "jpgoptimize")
	data.Jsinline = utils.MapGetBool(g, "jsinline")
	data.Jsminify = utils.MapGetBool(g, "jsminify")
	data.Jsmovetoend = utils.MapGetBool(g, "jsmovetoend")
	data.Pageextendcache = utils.MapGetBool(g, "pageextendcache")

	// Read-only metadata.
	data.Imgadddimensions = utils.MapGetBool(g, "imgadddimensions")
	data.Imgshrinkformobile = utils.MapGetBool(g, "imgshrinkformobile")
	data.Imgweaken = utils.MapGetBool(g, "imgweaken")
	data.Jpgprogressive = utils.MapGetBool(g, "jpgprogressive")
	data.Cssflattenimports = utils.MapGetBool(g, "cssflattenimports")
	data.Jscombine = utils.MapGetBool(g, "jscombine")
	data.Htmlrmdefaultattribs = utils.MapGetBool(g, "htmlrmdefaultattribs")
	data.Htmlrmattribquotes = utils.MapGetBool(g, "htmlrmattribquotes")
	data.Htmltrimurls = utils.MapGetBool(g, "htmltrimurls")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
