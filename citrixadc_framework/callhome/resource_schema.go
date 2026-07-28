package callhome

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/utility"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// CallhomeResourceModel describes the resource data model.
type CallhomeResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Emailaddress     types.String `tfsdk:"emailaddress"`
	Hbcustominterval types.Int64  `tfsdk:"hbcustominterval"`
	Ipaddress        types.String `tfsdk:"ipaddress"`
	Mode             types.String `tfsdk:"mode"`
	Nodeid           types.Int64  `tfsdk:"nodeid"`
	Port             types.Int64  `tfsdk:"port"`
	Proxyauthservice types.String `tfsdk:"proxyauthservice"`
	Proxymode        types.String `tfsdk:"proxymode"`
}

func (r *CallhomeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the callhome resource.",
			},
			"emailaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Email address of the contact administrator.",
			},
			"hbcustominterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(7),
				Description: "Interval (in days) between CallHome heartbeats",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the proxy server.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("Default"),
				Description: "CallHome mode of operation",
			},
			"nodeid": schema.Int64Attribute{
				// nodeid is a GET-only cluster-node filter argument, not a settable
				// property. It is Computed-only and excluded from the write payload.
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP port on the Proxy server. This is a mandatory parameter for both IP address and service name based configuration.",
			},
			"proxyauthservice": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the service that represents the proxy server.",
			},
			"proxymode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NO"),
				Description: "Enables or disables the proxy mode. The proxy server can be set by either specifying the IP address of the server or the name of the service representing the proxy server.",
			},
		},
	}
}

func callhomeGetThePayloadFromthePlan(ctx context.Context, data *CallhomeResourceModel) utility.Callhome {
	tflog.Debug(ctx, "In callhomeGetThePayloadFromthePlan Function")

	// Create API request body from the model
	callhome := utility.Callhome{}
	if !data.Emailaddress.IsNull() && !data.Emailaddress.IsUnknown() {
		callhome.Emailaddress = data.Emailaddress.ValueString()
	}
	if !data.Hbcustominterval.IsNull() && !data.Hbcustominterval.IsUnknown() {
		callhome.Hbcustominterval = utils.IntPtr(int(data.Hbcustominterval.ValueInt64()))
	}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		callhome.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Mode.IsNull() && !data.Mode.IsUnknown() {
		callhome.Mode = data.Mode.ValueString()
	}
	// nodeid is a GET-only cluster-node filter argument and is intentionally
	// excluded from the write/set payload (NITRO rejects it with errorcode 278).
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		callhome.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Proxyauthservice.IsNull() && !data.Proxyauthservice.IsUnknown() {
		callhome.Proxyauthservice = data.Proxyauthservice.ValueString()
	}
	if !data.Proxymode.IsNull() && !data.Proxymode.IsUnknown() {
		callhome.Proxymode = data.Proxymode.ValueString()
	}

	return callhome
}

func callhomeSetAttrFromGet(ctx context.Context, data *CallhomeResourceModel, getResponseData map[string]interface{}) *CallhomeResourceModel {
	tflog.Debug(ctx, "In callhomeSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["emailaddress"]; ok && val != nil {
		data.Emailaddress = types.StringValue(val.(string))
	} else {
		data.Emailaddress = types.StringNull()
	}
	if val, ok := getResponseData["hbcustominterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Hbcustominterval = types.Int64Value(intVal)
		}
	} else {
		data.Hbcustominterval = types.Int64Null()
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["mode"]; ok && val != nil {
		data.Mode = types.StringValue(val.(string))
	} else {
		data.Mode = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["proxyauthservice"]; ok && val != nil {
		data.Proxyauthservice = types.StringValue(val.(string))
	} else {
		data.Proxyauthservice = types.StringNull()
	}
	if val, ok := getResponseData["proxymode"]; ok && val != nil {
		data.Proxymode = types.StringValue(val.(string))
	} else {
		data.Proxymode = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID (singleton)
	data.Id = types.StringValue("callhome")

	return data
}
