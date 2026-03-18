package cmpparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cmp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// CmpparameterResourceModel describes the resource data model.
type CmpparameterResourceModel struct {
	Id                          types.String `tfsdk:"id"`
	Addvaryheader               types.String `tfsdk:"addvaryheader"`
	Cmpbypasspct                types.Int64  `tfsdk:"cmpbypasspct"`
	Cmplevel                    types.String `tfsdk:"cmplevel"`
	Cmponpush                   types.String `tfsdk:"cmponpush"`
	Externalcache               types.String `tfsdk:"externalcache"`
	Heurexpiry                  types.String `tfsdk:"heurexpiry"`
	Heurexpiryhistwt            types.Int64  `tfsdk:"heurexpiryhistwt"`
	Heurexpirythres             types.Int64  `tfsdk:"heurexpirythres"`
	Minressize                  types.Int64  `tfsdk:"minressize"`
	Policytype                  types.String `tfsdk:"policytype"`
	Quantumsize                 types.Int64  `tfsdk:"quantumsize"`
	Randomgzipfilename          types.String `tfsdk:"randomgzipfilename"`
	Randomgzipfilenamemaxlength types.Int64  `tfsdk:"randomgzipfilenamemaxlength"`
	Randomgzipfilenameminlength types.Int64  `tfsdk:"randomgzipfilenameminlength"`
	Servercmp                   types.String `tfsdk:"servercmp"`
	Varyheadervalue             types.String `tfsdk:"varyheadervalue"`
}

func (r *CmpparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cmpparameter resource.",
			},
			"addvaryheader": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Control insertion of the Vary header in HTTP responses compressed by Citrix ADC. Intermediate caches store different versions of the response for different values of the headers present in the Vary response header.",
			},
			"cmpbypasspct": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(100),
				Description: "Citrix ADC CPU threshold after which compression is not performed. Range: 0 - 100",
			},
			"cmplevel": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("optimal"),
				Description: "Specify a compression level. Available settings function as follows:\n * Optimal - Corresponds to a gzip GZIP level of 5-7.\n * Best speed - Corresponds to a gzip level of 1.\n * Best compression - Corresponds to a gzip level of 9.",
			},
			"cmponpush": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Citrix ADC does not wait for the quantum to be filled before starting to compress data. Upon receipt of a packet with a PUSH flag, the appliance immediately begins compression of the accumulated packets.",
			},
			"externalcache": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable insertion of  Cache-Control: private response directive to indicate response message is intended for a single user and must not be cached by a shared or proxy cache.",
			},
			"heurexpiry": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Heuristic basefile expiry.",
			},
			"heurexpiryhistwt": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(50),
				Description: "For heuristic basefile expiry, weightage to be given to historical delta compression ratio, specified as percentage.  For example, to give 25% weightage to historical ratio (and therefore 75% weightage to the ratio for current delta compression transaction), specify 25.",
			},
			"heurexpirythres": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(100),
				Description: "Threshold compression ratio for heuristic basefile expiry, multiplied by 100. For example, to set the threshold ratio to 1.25, specify 125.",
			},
			"minressize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Smallest response size, in bytes, to be compressed.",
			},
			"policytype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ADVANCED"),
				Description: "Type of the policy. The only possible value is ADVANCED",
			},
			"quantumsize": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(57344),
				Description: "Minimum quantum of data to be filled before compression begins.",
			},
			"randomgzipfilename": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Control the addition of a random filename of random length in the GZIP header to apply the Heal-the-BREACH mitigation for the BREACH attack.",
			},
			"randomgzipfilenamemaxlength": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(63),
				Description: "Maximum length of the random filename to be added in the GZIP header to apply the Heal-the-BREACH mitigation for the BREACH attack.",
			},
			"randomgzipfilenameminlength": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(8),
				Description: "Minimum length of the random filename to be added in the GZIP header to apply the Heal-the-BREACH mitigation for the BREACH attack.",
			},
			"servercmp": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Allow the server to send compressed data to the Citrix ADC. With the default setting, the Citrix ADC appliance handles all compression.",
			},
			"varyheadervalue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The value of the HTTP Vary header for compressed responses. If this argument is not specified, a default value of \"Accept-Encoding\" will be used.",
			},
		},
	}
}

func cmpparameterGetThePayloadFromtheConfig(ctx context.Context, data *CmpparameterResourceModel) cmp.Cmpparameter {
	tflog.Debug(ctx, "In cmpparameterGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	cmpparameter := cmp.Cmpparameter{}
	if !data.Addvaryheader.IsNull() {
		cmpparameter.Addvaryheader = data.Addvaryheader.ValueString()
	}
	if !data.Cmpbypasspct.IsNull() {
		cmpparameter.Cmpbypasspct = utils.IntPtr(int(data.Cmpbypasspct.ValueInt64()))
	}
	if !data.Cmplevel.IsNull() {
		cmpparameter.Cmplevel = data.Cmplevel.ValueString()
	}
	if !data.Cmponpush.IsNull() {
		cmpparameter.Cmponpush = data.Cmponpush.ValueString()
	}
	if !data.Externalcache.IsNull() {
		cmpparameter.Externalcache = data.Externalcache.ValueString()
	}
	if !data.Heurexpiry.IsNull() {
		cmpparameter.Heurexpiry = data.Heurexpiry.ValueString()
	}
	if !data.Heurexpiryhistwt.IsNull() {
		cmpparameter.Heurexpiryhistwt = utils.IntPtr(int(data.Heurexpiryhistwt.ValueInt64()))
	}
	if !data.Heurexpirythres.IsNull() {
		cmpparameter.Heurexpirythres = utils.IntPtr(int(data.Heurexpirythres.ValueInt64()))
	}
	if !data.Minressize.IsNull() {
		cmpparameter.Minressize = utils.IntPtr(int(data.Minressize.ValueInt64()))
	}
	if !data.Policytype.IsNull() {
		cmpparameter.Policytype = data.Policytype.ValueString()
	}
	if !data.Quantumsize.IsNull() {
		cmpparameter.Quantumsize = utils.IntPtr(int(data.Quantumsize.ValueInt64()))
	}
	if !data.Randomgzipfilename.IsNull() {
		cmpparameter.Randomgzipfilename = data.Randomgzipfilename.ValueString()
	}
	if !data.Randomgzipfilenamemaxlength.IsNull() {
		cmpparameter.Randomgzipfilenamemaxlength = utils.IntPtr(int(data.Randomgzipfilenamemaxlength.ValueInt64()))
	}
	if !data.Randomgzipfilenameminlength.IsNull() {
		cmpparameter.Randomgzipfilenameminlength = utils.IntPtr(int(data.Randomgzipfilenameminlength.ValueInt64()))
	}
	if !data.Servercmp.IsNull() {
		cmpparameter.Servercmp = data.Servercmp.ValueString()
	}
	if !data.Varyheadervalue.IsNull() {
		cmpparameter.Varyheadervalue = data.Varyheadervalue.ValueString()
	}

	return cmpparameter
}

func cmpparameterSetAttrFromGet(ctx context.Context, data *CmpparameterResourceModel, getResponseData map[string]interface{}) *CmpparameterResourceModel {
	tflog.Debug(ctx, "In cmpparameterSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["addvaryheader"]; ok && val != nil {
		data.Addvaryheader = types.StringValue(val.(string))
	} else {
		data.Addvaryheader = types.StringNull()
	}
	if val, ok := getResponseData["cmpbypasspct"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cmpbypasspct = types.Int64Value(intVal)
		}
	} else {
		data.Cmpbypasspct = types.Int64Null()
	}
	if val, ok := getResponseData["cmplevel"]; ok && val != nil {
		data.Cmplevel = types.StringValue(val.(string))
	} else {
		data.Cmplevel = types.StringNull()
	}
	if val, ok := getResponseData["cmponpush"]; ok && val != nil {
		data.Cmponpush = types.StringValue(val.(string))
	} else {
		data.Cmponpush = types.StringNull()
	}
	if val, ok := getResponseData["externalcache"]; ok && val != nil {
		data.Externalcache = types.StringValue(val.(string))
	} else {
		data.Externalcache = types.StringNull()
	}
	if val, ok := getResponseData["heurexpiry"]; ok && val != nil {
		data.Heurexpiry = types.StringValue(val.(string))
	} else {
		data.Heurexpiry = types.StringNull()
	}
	if val, ok := getResponseData["heurexpiryhistwt"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Heurexpiryhistwt = types.Int64Value(intVal)
		}
	} else {
		data.Heurexpiryhistwt = types.Int64Null()
	}
	if val, ok := getResponseData["heurexpirythres"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Heurexpirythres = types.Int64Value(intVal)
		}
	} else {
		data.Heurexpirythres = types.Int64Null()
	}
	if val, ok := getResponseData["minressize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minressize = types.Int64Value(intVal)
		}
	} else {
		data.Minressize = types.Int64Null()
	}
	if val, ok := getResponseData["policytype"]; ok && val != nil {
		data.Policytype = types.StringValue(val.(string))
	} else {
		data.Policytype = types.StringNull()
	}
	if val, ok := getResponseData["quantumsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Quantumsize = types.Int64Value(intVal)
		}
	} else {
		data.Quantumsize = types.Int64Null()
	}
	if val, ok := getResponseData["randomgzipfilename"]; ok && val != nil {
		data.Randomgzipfilename = types.StringValue(val.(string))
	} else {
		data.Randomgzipfilename = types.StringNull()
	}
	if val, ok := getResponseData["randomgzipfilenamemaxlength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Randomgzipfilenamemaxlength = types.Int64Value(intVal)
		}
	} else {
		data.Randomgzipfilenamemaxlength = types.Int64Null()
	}
	if val, ok := getResponseData["randomgzipfilenameminlength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Randomgzipfilenameminlength = types.Int64Value(intVal)
		}
	} else {
		data.Randomgzipfilenameminlength = types.Int64Null()
	}
	if val, ok := getResponseData["servercmp"]; ok && val != nil {
		data.Servercmp = types.StringValue(val.(string))
	} else {
		data.Servercmp = types.StringNull()
	}
	if val, ok := getResponseData["varyheadervalue"]; ok && val != nil {
		data.Varyheadervalue = types.StringValue(val.(string))
	} else {
		data.Varyheadervalue = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("cmpparameter-config")

	return data
}
