package sslzerotouchparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslzerotouchparamResourceModel describes the resource data model.
type SslzerotouchparamResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Ocspcachetimeout       types.Int64  `tfsdk:"ocspcachetimeout"`
	Ocspbatchingdepth      types.Int64  `tfsdk:"ocspbatchingdepth"`
	Ocspbatchingdelay      types.Int64  `tfsdk:"ocspbatchingdelay"`
	Ocspresptimeout        types.Int64  `tfsdk:"ocspresptimeout"`
	Ocspurlresolvetimeout  types.Int64  `tfsdk:"ocspurlresolvetimeout"`
	Ocsptrustresponder     types.String `tfsdk:"ocsptrustresponder"`
	Ocspproducedattimeskew types.Int64  `tfsdk:"ocspproducedattimeskew"`
	Ocspusenonce           types.String `tfsdk:"ocspusenonce"`
	Ocsphttpmethod         types.String `tfsdk:"ocsphttpmethod"`
}

func (r *SslzerotouchparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslzerotouchparam resource.",
			},
			"ocspcachetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout(in minutes) for caching the OCSP response.",
			},
			"ocspbatchingdepth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of certificates to batch together into one OCSP request. Batching avoids overloading the OCSP responder. A value of 1 signifies that each request is queried independently. For a value greater than 1, specify a timeout (batching delay) to avoid inordinately delaying the processing of a single certificate.",
			},
			"ocspbatchingdelay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum time, in milliseconds, to wait to accumulate OCSP requests to batch. Does not apply if the Batching Depth is 1.",
			},
			"ocspresptimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in milliseconds, to wait for an OCSP response. When this time elapses, an error message appears or the transaction is forwarded, depending on the settings on the virtual server. Includes Batching Delay time.",
			},
			"ocspurlresolvetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in milliseconds, to wait for an OCSP URL Resolution. When this time elapses, an error message appears or the transaction is forwarded, depending on the settings on the virtual server.",
			},
			// ocsptrustresponder/ocspusenonce/ocsphttpmethod each revert to a KNOWN
			// non-null appliance default on unset (ocsptrustresponder="NO",
			// ocspusenonce="ENABLED", ocsphttpmethod="POST") and NITRO always returns
			// them on GET. A matching schema Default therefore keeps the plan stable
			// (config-omit -> default == state) while the Update still issues
			// ?action=unset for a removed attribute (Update keys the unset off the raw
			// config being null, not the plan value), without the perpetual "known
			// after apply" churn an unknown-marking modifier would cause. See the
			// dpsparameter / systemautosaveparam resources for the same pattern.
			// NOTE: the Int64 attributes intentionally have NO Default: at least
			// ocspcachetimeout is OMITTED from GET, so a Default there would trigger a
			// "provider produced inconsistent result after apply".
			"ocsptrustresponder": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NO"),
				Description: "If trustResponder is set to YES, signature verification will be skipped for the OCSP response. Possible values = YES, NO",
			},
			"ocspproducedattimeskew": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, for which the Citrix ADC waits before considering the response as invalid. The response is considered invalid if the Produced At time stamp in the OCSP response exceeds or precedes the current Citrix ADC clock time by the amount of time specified.",
			},
			"ocspusenonce": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable the OCSP nonce extension, which is designed to prevent replay attacks. Possible values = ENABLED, DISABLED",
			},
			"ocsphttpmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("POST"),
				Description: "HTTP method used to send ocsp request. POST is the default httpmethod. If request length is > 255, POST wil be used even if GET is set as httpMethod. Possible values = GET, POST",
			},
		},
	}
}

func sslzerotouchparamGetThePayloadFromtheConfig(ctx context.Context, data *SslzerotouchparamResourceModel) ssl.Sslzerotouchparam {
	tflog.Debug(ctx, "In sslzerotouchparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslzerotouchparam := ssl.Sslzerotouchparam{}
	if !data.Ocspcachetimeout.IsNull() && !data.Ocspcachetimeout.IsUnknown() {
		sslzerotouchparam.Ocspcachetimeout = utils.IntPtr(int(data.Ocspcachetimeout.ValueInt64()))
	}
	if !data.Ocspbatchingdepth.IsNull() && !data.Ocspbatchingdepth.IsUnknown() {
		sslzerotouchparam.Ocspbatchingdepth = utils.IntPtr(int(data.Ocspbatchingdepth.ValueInt64()))
	}
	if !data.Ocspbatchingdelay.IsNull() && !data.Ocspbatchingdelay.IsUnknown() {
		sslzerotouchparam.Ocspbatchingdelay = utils.IntPtr(int(data.Ocspbatchingdelay.ValueInt64()))
	}
	if !data.Ocspresptimeout.IsNull() && !data.Ocspresptimeout.IsUnknown() {
		sslzerotouchparam.Ocspresptimeout = utils.IntPtr(int(data.Ocspresptimeout.ValueInt64()))
	}
	if !data.Ocspurlresolvetimeout.IsNull() && !data.Ocspurlresolvetimeout.IsUnknown() {
		sslzerotouchparam.Ocspurlresolvetimeout = utils.IntPtr(int(data.Ocspurlresolvetimeout.ValueInt64()))
	}
	if !data.Ocsptrustresponder.IsNull() && !data.Ocsptrustresponder.IsUnknown() {
		sslzerotouchparam.Ocsptrustresponder = data.Ocsptrustresponder.ValueString()
	}
	if !data.Ocspproducedattimeskew.IsNull() && !data.Ocspproducedattimeskew.IsUnknown() {
		sslzerotouchparam.Ocspproducedattimeskew = utils.IntPtr(int(data.Ocspproducedattimeskew.ValueInt64()))
	}
	if !data.Ocspusenonce.IsNull() && !data.Ocspusenonce.IsUnknown() {
		sslzerotouchparam.Ocspusenonce = data.Ocspusenonce.ValueString()
	}
	if !data.Ocsphttpmethod.IsNull() && !data.Ocsphttpmethod.IsUnknown() {
		sslzerotouchparam.Ocsphttpmethod = data.Ocsphttpmethod.ValueString()
	}

	return sslzerotouchparam
}

func sslzerotouchparamSetAttrFromGet(ctx context.Context, data *SslzerotouchparamResourceModel, getResponseData map[string]interface{}) *SslzerotouchparamResourceModel {
	tflog.Debug(ctx, "In sslzerotouchparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ocspcachetimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ocspcachetimeout = types.Int64Value(intVal)
		}
	} else {
		data.Ocspcachetimeout = types.Int64Null()
	}
	if val, ok := getResponseData["ocspbatchingdepth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ocspbatchingdepth = types.Int64Value(intVal)
		}
	} else {
		data.Ocspbatchingdepth = types.Int64Null()
	}
	if val, ok := getResponseData["ocspbatchingdelay"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ocspbatchingdelay = types.Int64Value(intVal)
		}
	} else {
		data.Ocspbatchingdelay = types.Int64Null()
	}
	if val, ok := getResponseData["ocspresptimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ocspresptimeout = types.Int64Value(intVal)
		}
	} else {
		data.Ocspresptimeout = types.Int64Null()
	}
	if val, ok := getResponseData["ocspurlresolvetimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ocspurlresolvetimeout = types.Int64Value(intVal)
		}
	} else {
		data.Ocspurlresolvetimeout = types.Int64Null()
	}
	if val, ok := getResponseData["ocsptrustresponder"]; ok && val != nil {
		data.Ocsptrustresponder = types.StringValue(val.(string))
	} else {
		data.Ocsptrustresponder = types.StringNull()
	}
	if val, ok := getResponseData["ocspproducedattimeskew"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ocspproducedattimeskew = types.Int64Value(intVal)
		}
	} else {
		data.Ocspproducedattimeskew = types.Int64Null()
	}
	if val, ok := getResponseData["ocspusenonce"]; ok && val != nil {
		data.Ocspusenonce = types.StringValue(val.(string))
	} else {
		data.Ocspusenonce = types.StringNull()
	}
	if val, ok := getResponseData["ocsphttpmethod"]; ok && val != nil {
		data.Ocsphttpmethod = types.StringValue(val.(string))
	} else {
		data.Ocsphttpmethod = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID (singleton)
	data.Id = types.StringValue("sslzerotouchparam-config")

	return data
}
