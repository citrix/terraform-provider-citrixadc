package sslocspresponder

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslocspresponderResourceModel describes the resource data model.
type SslocspresponderResourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Batchingdelay         types.Int64  `tfsdk:"batchingdelay"`
	Batchingdepth         types.Int64  `tfsdk:"batchingdepth"`
	Cache                 types.String `tfsdk:"cache"`
	Cachetimeout          types.Int64  `tfsdk:"cachetimeout"`
	Httpmethod            types.String `tfsdk:"httpmethod"`
	Insertclientcert      types.String `tfsdk:"insertclientcert"`
	Name                  types.String `tfsdk:"name"`
	Ocspurlresolvetimeout types.Int64  `tfsdk:"ocspurlresolvetimeout"`
	Producedattimeskew    types.Int64  `tfsdk:"producedattimeskew"`
	Respondercert         types.String `tfsdk:"respondercert"`
	Resptimeout           types.Int64  `tfsdk:"resptimeout"`
	Signingcert           types.String `tfsdk:"signingcert"`
	Trustresponder        types.Bool   `tfsdk:"trustresponder"`
	Url                   types.String `tfsdk:"url"`
	Usenonce              types.String `tfsdk:"usenonce"`
}

func (r *SslocspresponderResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslocspresponder resource.",
			},
			"batchingdelay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum time, in milliseconds, to wait to accumulate OCSP requests to batch.  Does not apply if the Batching Depth is 1.",
			},
			"batchingdepth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of client certificates to batch together into one OCSP request. Batching avoids overloading the OCSP responder. A value of 1 signifies that each request is queried independently. For a value greater than 1, specify a timeout (batching delay) to avoid inordinately delaying the processing of a single certificate.",
			},
			"cache": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable caching of responses. Caching of responses received from the OCSP responder enables faster responses to the clients and reduces the load on the OCSP responder.",
			},
			"cachetimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Timeout for caching the OCSP response. After the timeout, the Citrix ADC sends a fresh request to the OCSP responder for the certificate status. If a timeout is not specified, the timeout provided in the OCSP response applies.",
			},
			"httpmethod": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("POST"),
				Description: "HTTP method used to send ocsp request. POST is the default httpmethod. If request length is > 255, POST wil be used even if GET is set as httpMethod",
			},
			"insertclientcert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include the complete client certificate in the OCSP request.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the OCSP responder. Cannot begin with a hash (#) or space character and must contain only ASCII alphanumeric, underscore (_), hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the responder is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my responder\" or 'my responder').",
			},
			"ocspurlresolvetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in milliseconds, to wait for an OCSP URL Resolution. When this time elapses, an error message appears or the transaction is forwarded, depending on the settings on the virtual server.",
			},
			"producedattimeskew": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(300),
				Description: "Time, in seconds, for which the Citrix ADC waits before considering the response as invalid. The response is considered invalid if the Produced At time stamp in the OCSP response exceeds or precedes the current Citrix ADC clock time by the amount of time specified.",
			},
			"respondercert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"resptimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in milliseconds, to wait for an OCSP response. When this time elapses, an error message appears or the transaction is forwarded, depending on the settings on the virtual server. Includes Batching Delay time.",
			},
			"signingcert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Certificate-key pair that is used to sign OCSP requests. If this parameter is not set, the requests are not signed.",
			},
			"trustresponder": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A certificate to use to validate OCSP responses.  Alternatively, if -trustResponder is specified, no verification will be done on the reponse.  If both are omitted, only the response times (producedAt, lastUpdate, nextUpdate) will be verified.",
			},
			"url": schema.StringAttribute{
				Required:    true,
				Description: "URL of the OCSP responder.",
			},
			"usenonce": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the OCSP nonce extension, which is designed to prevent replay attacks.",
			},
		},
	}
}

func sslocspresponderGetThePayloadFromtheConfig(ctx context.Context, data *SslocspresponderResourceModel) ssl.Sslocspresponder {
	tflog.Debug(ctx, "In sslocspresponderGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslocspresponder := ssl.Sslocspresponder{}
	if !data.Batchingdelay.IsNull() {
		sslocspresponder.Batchingdelay = utils.IntPtr(int(data.Batchingdelay.ValueInt64()))
	}
	if !data.Batchingdepth.IsNull() {
		sslocspresponder.Batchingdepth = utils.IntPtr(int(data.Batchingdepth.ValueInt64()))
	}
	if !data.Cache.IsNull() {
		sslocspresponder.Cache = data.Cache.ValueString()
	}
	if !data.Cachetimeout.IsNull() {
		sslocspresponder.Cachetimeout = utils.IntPtr(int(data.Cachetimeout.ValueInt64()))
	}
	if !data.Httpmethod.IsNull() {
		sslocspresponder.Httpmethod = data.Httpmethod.ValueString()
	}
	if !data.Insertclientcert.IsNull() {
		sslocspresponder.Insertclientcert = data.Insertclientcert.ValueString()
	}
	if !data.Name.IsNull() {
		sslocspresponder.Name = data.Name.ValueString()
	}
	if !data.Ocspurlresolvetimeout.IsNull() {
		sslocspresponder.Ocspurlresolvetimeout = utils.IntPtr(int(data.Ocspurlresolvetimeout.ValueInt64()))
	}
	if !data.Producedattimeskew.IsNull() {
		sslocspresponder.Producedattimeskew = utils.IntPtr(int(data.Producedattimeskew.ValueInt64()))
	}
	if !data.Respondercert.IsNull() {
		sslocspresponder.Respondercert = data.Respondercert.ValueString()
	}
	if !data.Resptimeout.IsNull() {
		sslocspresponder.Resptimeout = utils.IntPtr(int(data.Resptimeout.ValueInt64()))
	}
	if !data.Signingcert.IsNull() {
		sslocspresponder.Signingcert = data.Signingcert.ValueString()
	}
	if !data.Trustresponder.IsNull() {
		sslocspresponder.Trustresponder = data.Trustresponder.ValueBool()
	}
	if !data.Url.IsNull() {
		sslocspresponder.Url = data.Url.ValueString()
	}
	if !data.Usenonce.IsNull() {
		sslocspresponder.Usenonce = data.Usenonce.ValueString()
	}

	return sslocspresponder
}

func sslocspresponderSetAttrFromGet(ctx context.Context, data *SslocspresponderResourceModel, getResponseData map[string]interface{}) *SslocspresponderResourceModel {
	tflog.Debug(ctx, "In sslocspresponderSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["batchingdelay"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Batchingdelay = types.Int64Value(intVal)
		}
	} else {
		data.Batchingdelay = types.Int64Null()
	}
	if val, ok := getResponseData["batchingdepth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Batchingdepth = types.Int64Value(intVal)
		}
	} else {
		data.Batchingdepth = types.Int64Null()
	}
	if val, ok := getResponseData["cache"]; ok && val != nil {
		data.Cache = types.StringValue(val.(string))
	} else {
		data.Cache = types.StringNull()
	}
	if val, ok := getResponseData["cachetimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cachetimeout = types.Int64Value(intVal)
		}
	} else {
		data.Cachetimeout = types.Int64Null()
	}
	if val, ok := getResponseData["httpmethod"]; ok && val != nil {
		data.Httpmethod = types.StringValue(val.(string))
	} else {
		data.Httpmethod = types.StringNull()
	}
	if val, ok := getResponseData["insertclientcert"]; ok && val != nil {
		data.Insertclientcert = types.StringValue(val.(string))
	} else {
		data.Insertclientcert = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["ocspurlresolvetimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ocspurlresolvetimeout = types.Int64Value(intVal)
		}
	} else {
		data.Ocspurlresolvetimeout = types.Int64Null()
	}
	if val, ok := getResponseData["producedattimeskew"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Producedattimeskew = types.Int64Value(intVal)
		}
	} else {
		data.Producedattimeskew = types.Int64Null()
	}
	if val, ok := getResponseData["respondercert"]; ok && val != nil {
		data.Respondercert = types.StringValue(val.(string))
	} else {
		data.Respondercert = types.StringNull()
	}
	if val, ok := getResponseData["resptimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Resptimeout = types.Int64Value(intVal)
		}
	} else {
		data.Resptimeout = types.Int64Null()
	}
	if val, ok := getResponseData["signingcert"]; ok && val != nil {
		data.Signingcert = types.StringValue(val.(string))
	} else {
		data.Signingcert = types.StringNull()
	}
	if val, ok := getResponseData["trustresponder"]; ok && val != nil {
		data.Trustresponder = types.BoolValue(val.(bool))
	} else {
		data.Trustresponder = types.BoolNull()
	}
	if val, ok := getResponseData["url"]; ok && val != nil {
		data.Url = types.StringValue(val.(string))
	} else {
		data.Url = types.StringNull()
	}
	if val, ok := getResponseData["usenonce"]; ok && val != nil {
		data.Usenonce = types.StringValue(val.(string))
	} else {
		data.Usenonce = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
