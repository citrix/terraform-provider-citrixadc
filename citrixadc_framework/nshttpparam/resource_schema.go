package nshttpparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NshttpparamResourceModel describes the resource data model.
type NshttpparamResourceModel struct {
	Id                        types.String `tfsdk:"id"`
	Conmultiplex              types.String `tfsdk:"conmultiplex"`
	Dropinvalreqs             types.String `tfsdk:"dropinvalreqs"`
	Http2serverside           types.String `tfsdk:"http2serverside"`
	Ignoreconnectcodingscheme types.String `tfsdk:"ignoreconnectcodingscheme"`
	Insnssrvrhdr              types.String `tfsdk:"insnssrvrhdr"`
	Logerrresp                types.String `tfsdk:"logerrresp"`
	Markconnreqinval          types.String `tfsdk:"markconnreqinval"`
	Markhttp09inval           types.String `tfsdk:"markhttp09inval"`
	Maxreusepool              types.Int64  `tfsdk:"maxreusepool"`
	Nssrvrhdr                 types.String `tfsdk:"nssrvrhdr"`
}

func (r *NshttpparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nshttpparam resource.",
			},
			"conmultiplex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Reuse server connections for requests from more than one client connections.",
			},
			"dropinvalreqs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("OFF"),
				Description: "Drop invalid HTTP requests or responses.",
			},
			"http2serverside": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("OFF"),
				Description: "Enable/Disable HTTP/2 on server side",
			},
			"ignoreconnectcodingscheme": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Ignore Coding scheme in CONNECT request.",
			},
			"insnssrvrhdr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("OFF"),
				Description: "Enable or disable Citrix ADC server header insertion for Citrix ADC generated HTTP responses.",
			},
			"logerrresp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ON"),
				Description: "Server header value to be inserted.",
			},
			"markconnreqinval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("OFF"),
				Description: "Mark CONNECT requests as invalid.",
			},
			"markhttp09inval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("OFF"),
				Description: "Mark HTTP/0.9 requests as invalid.",
			},
			"maxreusepool": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
				Description: "Maximum limit on the number of connections, from the Citrix ADC to a particular server that are kept in the reuse pool. This setting is helpful for optimal memory utilization and for reducing the idle connections to the server just after the peak time.",
			},
			"nssrvrhdr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The server header value to be inserted. If no explicit header is specified then NSBUILD.RELEASE is used as default server header.",
			},
		},
	}
}

func nshttpparamGetThePayloadFromtheConfig(ctx context.Context, data *NshttpparamResourceModel) ns.Nshttpparam {
	tflog.Debug(ctx, "In nshttpparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model.
	// Only send explicitly-configured (known, non-null) values so that
	// unconfigured Optional+Computed attributes retain their current ADC value.
	nshttpparam := ns.Nshttpparam{}
	if !data.Conmultiplex.IsNull() && !data.Conmultiplex.IsUnknown() {
		nshttpparam.Conmultiplex = data.Conmultiplex.ValueString()
	}
	if !data.Dropinvalreqs.IsNull() && !data.Dropinvalreqs.IsUnknown() {
		nshttpparam.Dropinvalreqs = data.Dropinvalreqs.ValueString()
	}
	if !data.Http2serverside.IsNull() && !data.Http2serverside.IsUnknown() {
		nshttpparam.Http2serverside = data.Http2serverside.ValueString()
	}
	if !data.Ignoreconnectcodingscheme.IsNull() && !data.Ignoreconnectcodingscheme.IsUnknown() {
		nshttpparam.Ignoreconnectcodingscheme = data.Ignoreconnectcodingscheme.ValueString()
	}
	if !data.Insnssrvrhdr.IsNull() && !data.Insnssrvrhdr.IsUnknown() {
		nshttpparam.Insnssrvrhdr = data.Insnssrvrhdr.ValueString()
	}
	if !data.Logerrresp.IsNull() && !data.Logerrresp.IsUnknown() {
		nshttpparam.Logerrresp = data.Logerrresp.ValueString()
	}
	if !data.Markconnreqinval.IsNull() && !data.Markconnreqinval.IsUnknown() {
		nshttpparam.Markconnreqinval = data.Markconnreqinval.ValueString()
	}
	if !data.Markhttp09inval.IsNull() && !data.Markhttp09inval.IsUnknown() {
		nshttpparam.Markhttp09inval = data.Markhttp09inval.ValueString()
	}
	if !data.Maxreusepool.IsNull() && !data.Maxreusepool.IsUnknown() {
		nshttpparam.Maxreusepool = utils.IntPtr(int(data.Maxreusepool.ValueInt64()))
	}
	if !data.Nssrvrhdr.IsNull() && !data.Nssrvrhdr.IsUnknown() {
		nshttpparam.Nssrvrhdr = data.Nssrvrhdr.ValueString()
	}

	return nshttpparam
}

func nshttpparamSetAttrFromGet(ctx context.Context, data *NshttpparamResourceModel, getResponseData map[string]interface{}) *NshttpparamResourceModel {
	tflog.Debug(ctx, "In nshttpparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["conmultiplex"]; ok && val != nil {
		data.Conmultiplex = types.StringValue(val.(string))
	} else if data.Conmultiplex.IsUnknown() {
		data.Conmultiplex = types.StringNull()
	}
	if val, ok := getResponseData["dropinvalreqs"]; ok && val != nil {
		data.Dropinvalreqs = types.StringValue(val.(string))
	} else if data.Dropinvalreqs.IsUnknown() {
		data.Dropinvalreqs = types.StringNull()
	}
	if val, ok := getResponseData["http2serverside"]; ok && val != nil {
		data.Http2serverside = types.StringValue(val.(string))
	} else if data.Http2serverside.IsUnknown() {
		data.Http2serverside = types.StringNull()
	}
	if val, ok := getResponseData["ignoreconnectcodingscheme"]; ok && val != nil {
		data.Ignoreconnectcodingscheme = types.StringValue(val.(string))
	} else if data.Ignoreconnectcodingscheme.IsUnknown() {
		data.Ignoreconnectcodingscheme = types.StringNull()
	}
	if val, ok := getResponseData["insnssrvrhdr"]; ok && val != nil {
		data.Insnssrvrhdr = types.StringValue(val.(string))
	} else if data.Insnssrvrhdr.IsUnknown() {
		data.Insnssrvrhdr = types.StringNull()
	}
	if val, ok := getResponseData["logerrresp"]; ok && val != nil {
		data.Logerrresp = types.StringValue(val.(string))
	} else if data.Logerrresp.IsUnknown() {
		data.Logerrresp = types.StringNull()
	}
	if val, ok := getResponseData["markconnreqinval"]; ok && val != nil {
		data.Markconnreqinval = types.StringValue(val.(string))
	} else if data.Markconnreqinval.IsUnknown() {
		data.Markconnreqinval = types.StringNull()
	}
	if val, ok := getResponseData["markhttp09inval"]; ok && val != nil {
		data.Markhttp09inval = types.StringValue(val.(string))
	} else if data.Markhttp09inval.IsUnknown() {
		data.Markhttp09inval = types.StringNull()
	}
	if val, ok := getResponseData["maxreusepool"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxreusepool = types.Int64Value(intVal)
		}
	} else if data.Maxreusepool.IsUnknown() {
		data.Maxreusepool = types.Int64Null()
	}
	if val, ok := getResponseData["nssrvrhdr"]; ok && val != nil {
		data.Nssrvrhdr = types.StringValue(val.(string))
	} else if data.Nssrvrhdr.IsUnknown() {
		data.Nssrvrhdr = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("nshttpparam-config")

	return data
}
