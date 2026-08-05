package feoaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/feo"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// FeoactionResourceModel describes the resource data model.
type FeoactionResourceModel struct {
	Id                     types.String `tfsdk:"id"`
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
	Name                   types.String `tfsdk:"name"`
	Pageextendcache        types.Bool   `tfsdk:"pageextendcache"`
}

func (r *FeoactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the feoaction resource.",
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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

func feoactionGetThePayloadFromthePlan(ctx context.Context, data *FeoactionResourceModel) feo.Feoaction {
	tflog.Debug(ctx, "In feoactionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	feoaction := feo.Feoaction{}
	if !data.Cachemaxage.IsNull() && !data.Cachemaxage.IsUnknown() {
		feoaction.Cachemaxage = utils.IntPtr(int(data.Cachemaxage.ValueInt64()))
	}
	if !data.Clientsidemeasurements.IsNull() && !data.Clientsidemeasurements.IsUnknown() {
		feoaction.Clientsidemeasurements = data.Clientsidemeasurements.ValueBool()
	}
	if !data.Convertimporttolink.IsNull() && !data.Convertimporttolink.IsUnknown() {
		feoaction.Convertimporttolink = data.Convertimporttolink.ValueBool()
	}
	if !data.Csscombine.IsNull() && !data.Csscombine.IsUnknown() {
		feoaction.Csscombine = data.Csscombine.ValueBool()
	}
	if !data.Cssimginline.IsNull() && !data.Cssimginline.IsUnknown() {
		feoaction.Cssimginline = data.Cssimginline.ValueBool()
	}
	if !data.Cssinline.IsNull() && !data.Cssinline.IsUnknown() {
		feoaction.Cssinline = data.Cssinline.ValueBool()
	}
	if !data.Cssminify.IsNull() && !data.Cssminify.IsUnknown() {
		feoaction.Cssminify = data.Cssminify.ValueBool()
	}
	if !data.Cssmovetohead.IsNull() && !data.Cssmovetohead.IsUnknown() {
		feoaction.Cssmovetohead = data.Cssmovetohead.ValueBool()
	}
	if !data.Dnsshards.IsNull() && !data.Dnsshards.IsUnknown() {
		var dnsshardsList []string
		data.Dnsshards.ElementsAs(ctx, &dnsshardsList, false)
		feoaction.Dnsshards = dnsshardsList
	}
	if !data.Domainsharding.IsNull() && !data.Domainsharding.IsUnknown() {
		feoaction.Domainsharding = data.Domainsharding.ValueString()
	}
	if !data.Htmlminify.IsNull() && !data.Htmlminify.IsUnknown() {
		feoaction.Htmlminify = data.Htmlminify.ValueBool()
	}
	if !data.Imggiftopng.IsNull() && !data.Imggiftopng.IsUnknown() {
		feoaction.Imggiftopng = data.Imggiftopng.ValueBool()
	}
	if !data.Imginline.IsNull() && !data.Imginline.IsUnknown() {
		feoaction.Imginline = data.Imginline.ValueBool()
	}
	if !data.Imglazyload.IsNull() && !data.Imglazyload.IsUnknown() {
		feoaction.Imglazyload = data.Imglazyload.ValueBool()
	}
	if !data.Imgshrinktoattrib.IsNull() && !data.Imgshrinktoattrib.IsUnknown() {
		feoaction.Imgshrinktoattrib = data.Imgshrinktoattrib.ValueBool()
	}
	if !data.Imgtojpegxr.IsNull() && !data.Imgtojpegxr.IsUnknown() {
		feoaction.Imgtojpegxr = data.Imgtojpegxr.ValueBool()
	}
	if !data.Imgtowebp.IsNull() && !data.Imgtowebp.IsUnknown() {
		feoaction.Imgtowebp = data.Imgtowebp.ValueBool()
	}
	if !data.Jpgoptimize.IsNull() && !data.Jpgoptimize.IsUnknown() {
		feoaction.Jpgoptimize = data.Jpgoptimize.ValueBool()
	}
	if !data.Jsinline.IsNull() && !data.Jsinline.IsUnknown() {
		feoaction.Jsinline = data.Jsinline.ValueBool()
	}
	if !data.Jsminify.IsNull() && !data.Jsminify.IsUnknown() {
		feoaction.Jsminify = data.Jsminify.ValueBool()
	}
	if !data.Jsmovetoend.IsNull() && !data.Jsmovetoend.IsUnknown() {
		feoaction.Jsmovetoend = data.Jsmovetoend.ValueBool()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		feoaction.Name = data.Name.ValueString()
	}
	if !data.Pageextendcache.IsNull() && !data.Pageextendcache.IsUnknown() {
		feoaction.Pageextendcache = data.Pageextendcache.ValueBool()
	}

	return feoaction
}

func feoactionSetAttrFromGet(ctx context.Context, data *FeoactionResourceModel, getResponseData map[string]interface{}) *FeoactionResourceModel {
	tflog.Debug(ctx, "In feoactionSetAttrFromGet Function")

	// Convert API response to model.
	// For Optional+Computed attributes: when NITRO omits the field (e.g. a
	// default 0/false/empty value that is not echoed back), preserve a value
	// the user explicitly configured (do not clobber it). Only reset to Null
	// when the model value is still Unknown (Computed, not configured) so the
	// state never carries an unknown value after apply.
	if val, ok := getResponseData["cachemaxage"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cachemaxage = types.Int64Value(intVal)
		}
	} else if data.Cachemaxage.IsUnknown() {
		data.Cachemaxage = types.Int64Null()
	}
	if val, ok := getResponseData["clientsidemeasurements"]; ok && val != nil {
		data.Clientsidemeasurements = types.BoolValue(val.(bool))
	} else if data.Clientsidemeasurements.IsUnknown() {
		data.Clientsidemeasurements = types.BoolNull()
	}
	if val, ok := getResponseData["convertimporttolink"]; ok && val != nil {
		data.Convertimporttolink = types.BoolValue(val.(bool))
	} else if data.Convertimporttolink.IsUnknown() {
		data.Convertimporttolink = types.BoolNull()
	}
	if val, ok := getResponseData["csscombine"]; ok && val != nil {
		data.Csscombine = types.BoolValue(val.(bool))
	} else if data.Csscombine.IsUnknown() {
		data.Csscombine = types.BoolNull()
	}
	if val, ok := getResponseData["cssimginline"]; ok && val != nil {
		data.Cssimginline = types.BoolValue(val.(bool))
	} else if data.Cssimginline.IsUnknown() {
		data.Cssimginline = types.BoolNull()
	}
	if val, ok := getResponseData["cssinline"]; ok && val != nil {
		data.Cssinline = types.BoolValue(val.(bool))
	} else if data.Cssinline.IsUnknown() {
		data.Cssinline = types.BoolNull()
	}
	if val, ok := getResponseData["cssminify"]; ok && val != nil {
		data.Cssminify = types.BoolValue(val.(bool))
	} else if data.Cssminify.IsUnknown() {
		data.Cssminify = types.BoolNull()
	}
	if val, ok := getResponseData["cssmovetohead"]; ok && val != nil {
		data.Cssmovetohead = types.BoolValue(val.(bool))
	} else if data.Cssmovetohead.IsUnknown() {
		data.Cssmovetohead = types.BoolNull()
	}
	if val, ok := getResponseData["dnsshards"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			stringList := utils.ToStringList(sliceVal)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Dnsshards = listValue
		} else if data.Dnsshards.IsUnknown() {
			data.Dnsshards = types.ListNull(types.StringType)
		}
	} else if data.Dnsshards.IsUnknown() {
		data.Dnsshards = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["domainsharding"]; ok && val != nil {
		data.Domainsharding = types.StringValue(val.(string))
	} else if data.Domainsharding.IsUnknown() {
		data.Domainsharding = types.StringNull()
	}
	if val, ok := getResponseData["htmlminify"]; ok && val != nil {
		data.Htmlminify = types.BoolValue(val.(bool))
	} else if data.Htmlminify.IsUnknown() {
		data.Htmlminify = types.BoolNull()
	}
	if val, ok := getResponseData["imggiftopng"]; ok && val != nil {
		data.Imggiftopng = types.BoolValue(val.(bool))
	} else if data.Imggiftopng.IsUnknown() {
		data.Imggiftopng = types.BoolNull()
	}
	if val, ok := getResponseData["imginline"]; ok && val != nil {
		data.Imginline = types.BoolValue(val.(bool))
	} else if data.Imginline.IsUnknown() {
		data.Imginline = types.BoolNull()
	}
	if val, ok := getResponseData["imglazyload"]; ok && val != nil {
		data.Imglazyload = types.BoolValue(val.(bool))
	} else if data.Imglazyload.IsUnknown() {
		data.Imglazyload = types.BoolNull()
	}
	if val, ok := getResponseData["imgshrinktoattrib"]; ok && val != nil {
		data.Imgshrinktoattrib = types.BoolValue(val.(bool))
	} else if data.Imgshrinktoattrib.IsUnknown() {
		data.Imgshrinktoattrib = types.BoolNull()
	}
	if val, ok := getResponseData["imgtojpegxr"]; ok && val != nil {
		data.Imgtojpegxr = types.BoolValue(val.(bool))
	} else if data.Imgtojpegxr.IsUnknown() {
		data.Imgtojpegxr = types.BoolNull()
	}
	if val, ok := getResponseData["imgtowebp"]; ok && val != nil {
		data.Imgtowebp = types.BoolValue(val.(bool))
	} else if data.Imgtowebp.IsUnknown() {
		data.Imgtowebp = types.BoolNull()
	}
	if val, ok := getResponseData["jpgoptimize"]; ok && val != nil {
		data.Jpgoptimize = types.BoolValue(val.(bool))
	} else if data.Jpgoptimize.IsUnknown() {
		data.Jpgoptimize = types.BoolNull()
	}
	if val, ok := getResponseData["jsinline"]; ok && val != nil {
		data.Jsinline = types.BoolValue(val.(bool))
	} else if data.Jsinline.IsUnknown() {
		data.Jsinline = types.BoolNull()
	}
	if val, ok := getResponseData["jsminify"]; ok && val != nil {
		data.Jsminify = types.BoolValue(val.(bool))
	} else if data.Jsminify.IsUnknown() {
		data.Jsminify = types.BoolNull()
	}
	if val, ok := getResponseData["jsmovetoend"]; ok && val != nil {
		data.Jsmovetoend = types.BoolValue(val.(bool))
	} else if data.Jsmovetoend.IsUnknown() {
		data.Jsmovetoend = types.BoolNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["pageextendcache"]; ok && val != nil {
		data.Pageextendcache = types.BoolValue(val.(bool))
	} else if data.Pageextendcache.IsUnknown() {
		data.Pageextendcache = types.BoolNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}

// feoactionSetAttrFromGetForDatasource copies every attribute returned by the
// NITRO GET into the datasource model (nulling absent fields) and sets the ID.
// The datasource never has configured/unknown values to preserve, so it uses a
// plain else->Null form rather than the resource's state-preserving guard.
func feoactionSetAttrFromGetForDatasource(ctx context.Context, data *FeoactionResourceModel, getResponseData map[string]interface{}) *FeoactionResourceModel {
	tflog.Debug(ctx, "In feoactionSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["cachemaxage"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cachemaxage = types.Int64Value(intVal)
		}
	} else {
		data.Cachemaxage = types.Int64Null()
	}
	if val, ok := getResponseData["clientsidemeasurements"]; ok && val != nil {
		data.Clientsidemeasurements = types.BoolValue(val.(bool))
	} else {
		data.Clientsidemeasurements = types.BoolNull()
	}
	if val, ok := getResponseData["convertimporttolink"]; ok && val != nil {
		data.Convertimporttolink = types.BoolValue(val.(bool))
	} else {
		data.Convertimporttolink = types.BoolNull()
	}
	if val, ok := getResponseData["csscombine"]; ok && val != nil {
		data.Csscombine = types.BoolValue(val.(bool))
	} else {
		data.Csscombine = types.BoolNull()
	}
	if val, ok := getResponseData["cssimginline"]; ok && val != nil {
		data.Cssimginline = types.BoolValue(val.(bool))
	} else {
		data.Cssimginline = types.BoolNull()
	}
	if val, ok := getResponseData["cssinline"]; ok && val != nil {
		data.Cssinline = types.BoolValue(val.(bool))
	} else {
		data.Cssinline = types.BoolNull()
	}
	if val, ok := getResponseData["cssminify"]; ok && val != nil {
		data.Cssminify = types.BoolValue(val.(bool))
	} else {
		data.Cssminify = types.BoolNull()
	}
	if val, ok := getResponseData["cssmovetohead"]; ok && val != nil {
		data.Cssmovetohead = types.BoolValue(val.(bool))
	} else {
		data.Cssmovetohead = types.BoolNull()
	}
	if val, ok := getResponseData["dnsshards"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			stringList := utils.ToStringList(sliceVal)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Dnsshards = listValue
		} else {
			data.Dnsshards = types.ListNull(types.StringType)
		}
	} else {
		data.Dnsshards = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["domainsharding"]; ok && val != nil {
		data.Domainsharding = types.StringValue(val.(string))
	} else {
		data.Domainsharding = types.StringNull()
	}
	if val, ok := getResponseData["htmlminify"]; ok && val != nil {
		data.Htmlminify = types.BoolValue(val.(bool))
	} else {
		data.Htmlminify = types.BoolNull()
	}
	if val, ok := getResponseData["imggiftopng"]; ok && val != nil {
		data.Imggiftopng = types.BoolValue(val.(bool))
	} else {
		data.Imggiftopng = types.BoolNull()
	}
	if val, ok := getResponseData["imginline"]; ok && val != nil {
		data.Imginline = types.BoolValue(val.(bool))
	} else {
		data.Imginline = types.BoolNull()
	}
	if val, ok := getResponseData["imglazyload"]; ok && val != nil {
		data.Imglazyload = types.BoolValue(val.(bool))
	} else {
		data.Imglazyload = types.BoolNull()
	}
	if val, ok := getResponseData["imgshrinktoattrib"]; ok && val != nil {
		data.Imgshrinktoattrib = types.BoolValue(val.(bool))
	} else {
		data.Imgshrinktoattrib = types.BoolNull()
	}
	if val, ok := getResponseData["imgtojpegxr"]; ok && val != nil {
		data.Imgtojpegxr = types.BoolValue(val.(bool))
	} else {
		data.Imgtojpegxr = types.BoolNull()
	}
	if val, ok := getResponseData["imgtowebp"]; ok && val != nil {
		data.Imgtowebp = types.BoolValue(val.(bool))
	} else {
		data.Imgtowebp = types.BoolNull()
	}
	if val, ok := getResponseData["jpgoptimize"]; ok && val != nil {
		data.Jpgoptimize = types.BoolValue(val.(bool))
	} else {
		data.Jpgoptimize = types.BoolNull()
	}
	if val, ok := getResponseData["jsinline"]; ok && val != nil {
		data.Jsinline = types.BoolValue(val.(bool))
	} else {
		data.Jsinline = types.BoolNull()
	}
	if val, ok := getResponseData["jsminify"]; ok && val != nil {
		data.Jsminify = types.BoolValue(val.(bool))
	} else {
		data.Jsminify = types.BoolNull()
	}
	if val, ok := getResponseData["jsmovetoend"]; ok && val != nil {
		data.Jsmovetoend = types.BoolValue(val.(bool))
	} else {
		data.Jsmovetoend = types.BoolNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["pageextendcache"]; ok && val != nil {
		data.Pageextendcache = types.BoolValue(val.(bool))
	} else {
		data.Pageextendcache = types.BoolNull()
	}

	// Set ID for the datasource
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
