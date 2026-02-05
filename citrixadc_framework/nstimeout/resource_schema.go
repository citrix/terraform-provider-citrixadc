package nstimeout

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NstimeoutResourceModel describes the resource data model.
type NstimeoutResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Anyclient          types.Int64  `tfsdk:"anyclient"`
	Anyserver          types.Int64  `tfsdk:"anyserver"`
	Anytcpclient       types.Int64  `tfsdk:"anytcpclient"`
	Anytcpserver       types.Int64  `tfsdk:"anytcpserver"`
	Client             types.Int64  `tfsdk:"client"`
	Halfclose          types.Int64  `tfsdk:"halfclose"`
	Httpclient         types.Int64  `tfsdk:"httpclient"`
	Httpserver         types.Int64  `tfsdk:"httpserver"`
	Newconnidletimeout types.Int64  `tfsdk:"newconnidletimeout"`
	Nontcpzombie       types.Int64  `tfsdk:"nontcpzombie"`
	Reducedfintimeout  types.Int64  `tfsdk:"reducedfintimeout"`
	Reducedrsttimeout  types.Int64  `tfsdk:"reducedrsttimeout"`
	Server             types.Int64  `tfsdk:"server"`
	Tcpclient          types.Int64  `tfsdk:"tcpclient"`
	Tcpserver          types.Int64  `tfsdk:"tcpserver"`
	Zombie             types.Int64  `tfsdk:"zombie"`
}

func (r *NstimeoutResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nstimeout resource.",
			},
			"anyclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Global idle timeout, in seconds, for non-TCP client connections. This value is over ridden by the client timeout that is configured on individual entities.",
			},
			"anyserver": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Global idle timeout, in seconds, for non TCP server connections. This value is over ridden by the server timeout that is configured on individual entities.",
			},
			"anytcpclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Global idle timeout, in seconds, for TCP client connections. This value takes precedence over  entity level timeout settings (vserver/service). This is applicable only to transport protocol TCP.",
			},
			"anytcpserver": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Global idle timeout, in seconds, for TCP server connections. This value takes precedence over entity level timeout settings ( vserver/service). This is applicable only to transport protocol TCP.",
			},
			"client": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Client idle timeout (in seconds). If zero, the service-type default value is taken when service is created.",
			},
			"halfclose": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10),
				Description: "Idle timeout, in seconds, for connections that are in TCP half-closed state.",
			},
			"httpclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Global idle timeout, in seconds, for client connections of HTTP service type. This value is over ridden by the client timeout that is configured on individual entities.",
			},
			"httpserver": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Global idle timeout, in seconds, for server connections of HTTP service type. This value is over ridden by the server timeout that is configured on individual entities.",
			},
			"newconnidletimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(4),
				Description: "Timer interval, in seconds, for new TCP NATPCB connections on which no data was received.",
			},
			"nontcpzombie": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(60),
				Description: "Interval at which the zombie clean-up process for non-TCP connections should run. Inactive IP NAT connections will be cleaned up.",
			},
			"reducedfintimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(30),
				Description: "Alternative idle timeout, in seconds, for closed TCP NATPCB connections.",
			},
			"reducedrsttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timer interval, in seconds, for abruptly terminated TCP NATPCB connections.",
			},
			"server": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Server idle timeout (in seconds).  If zero, the service-type default value is taken when service is created.",
			},
			"tcpclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Global idle timeout, in seconds, for non-HTTP client connections of TCP service type. This value is over ridden by the client timeout that is configured on individual entities.",
			},
			"tcpserver": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Global idle timeout, in seconds, for non-HTTP server connections of TCP service type. This value is over ridden by the server timeout that is configured on entities.",
			},
			"zombie": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(120),
				Description: "Interval, in seconds, at which the Citrix ADC zombie cleanup process must run. This process cleans up inactive TCP connections.",
			},
		},
	}
}

func nstimeoutGetThePayloadFromtheConfig(ctx context.Context, data *NstimeoutResourceModel) ns.Nstimeout {
	tflog.Debug(ctx, "In nstimeoutGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nstimeout := ns.Nstimeout{}
	if !data.Anyclient.IsNull() {
		nstimeout.Anyclient = utils.IntPtr(int(data.Anyclient.ValueInt64()))
	}
	if !data.Anyserver.IsNull() {
		nstimeout.Anyserver = utils.IntPtr(int(data.Anyserver.ValueInt64()))
	}
	if !data.Anytcpclient.IsNull() {
		nstimeout.Anytcpclient = utils.IntPtr(int(data.Anytcpclient.ValueInt64()))
	}
	if !data.Anytcpserver.IsNull() {
		nstimeout.Anytcpserver = utils.IntPtr(int(data.Anytcpserver.ValueInt64()))
	}
	if !data.Client.IsNull() {
		nstimeout.Client = utils.IntPtr(int(data.Client.ValueInt64()))
	}
	if !data.Halfclose.IsNull() {
		nstimeout.Halfclose = utils.IntPtr(int(data.Halfclose.ValueInt64()))
	}
	if !data.Httpclient.IsNull() {
		nstimeout.Httpclient = utils.IntPtr(int(data.Httpclient.ValueInt64()))
	}
	if !data.Httpserver.IsNull() {
		nstimeout.Httpserver = utils.IntPtr(int(data.Httpserver.ValueInt64()))
	}
	if !data.Newconnidletimeout.IsNull() {
		nstimeout.Newconnidletimeout = utils.IntPtr(int(data.Newconnidletimeout.ValueInt64()))
	}
	if !data.Nontcpzombie.IsNull() {
		nstimeout.Nontcpzombie = utils.IntPtr(int(data.Nontcpzombie.ValueInt64()))
	}
	if !data.Reducedfintimeout.IsNull() {
		nstimeout.Reducedfintimeout = utils.IntPtr(int(data.Reducedfintimeout.ValueInt64()))
	}
	if !data.Reducedrsttimeout.IsNull() {
		nstimeout.Reducedrsttimeout = utils.IntPtr(int(data.Reducedrsttimeout.ValueInt64()))
	}
	if !data.Server.IsNull() {
		nstimeout.Server = utils.IntPtr(int(data.Server.ValueInt64()))
	}
	if !data.Tcpclient.IsNull() {
		nstimeout.Tcpclient = utils.IntPtr(int(data.Tcpclient.ValueInt64()))
	}
	if !data.Tcpserver.IsNull() {
		nstimeout.Tcpserver = utils.IntPtr(int(data.Tcpserver.ValueInt64()))
	}
	if !data.Zombie.IsNull() {
		nstimeout.Zombie = utils.IntPtr(int(data.Zombie.ValueInt64()))
	}

	return nstimeout
}

func nstimeoutSetAttrFromGet(ctx context.Context, data *NstimeoutResourceModel, getResponseData map[string]interface{}) *NstimeoutResourceModel {
	tflog.Debug(ctx, "In nstimeoutSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["anyclient"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Anyclient = types.Int64Value(intVal)
		}
	} else {
		data.Anyclient = types.Int64Null()
	}
	if val, ok := getResponseData["anyserver"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Anyserver = types.Int64Value(intVal)
		}
	} else {
		data.Anyserver = types.Int64Null()
	}
	if val, ok := getResponseData["anytcpclient"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Anytcpclient = types.Int64Value(intVal)
		}
	} else {
		data.Anytcpclient = types.Int64Null()
	}
	if val, ok := getResponseData["anytcpserver"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Anytcpserver = types.Int64Value(intVal)
		}
	} else {
		data.Anytcpserver = types.Int64Null()
	}
	if val, ok := getResponseData["client"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Client = types.Int64Value(intVal)
		}
	} else {
		data.Client = types.Int64Null()
	}
	if val, ok := getResponseData["halfclose"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Halfclose = types.Int64Value(intVal)
		}
	} else {
		data.Halfclose = types.Int64Null()
	}
	if val, ok := getResponseData["httpclient"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Httpclient = types.Int64Value(intVal)
		}
	} else {
		data.Httpclient = types.Int64Null()
	}
	if val, ok := getResponseData["httpserver"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Httpserver = types.Int64Value(intVal)
		}
	} else {
		data.Httpserver = types.Int64Null()
	}
	if val, ok := getResponseData["newconnidletimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Newconnidletimeout = types.Int64Value(intVal)
		}
	} else {
		data.Newconnidletimeout = types.Int64Null()
	}
	if val, ok := getResponseData["nontcpzombie"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nontcpzombie = types.Int64Value(intVal)
		}
	} else {
		data.Nontcpzombie = types.Int64Null()
	}
	if val, ok := getResponseData["reducedfintimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Reducedfintimeout = types.Int64Value(intVal)
		}
	} else {
		data.Reducedfintimeout = types.Int64Null()
	}
	if val, ok := getResponseData["reducedrsttimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Reducedrsttimeout = types.Int64Value(intVal)
		}
	} else {
		data.Reducedrsttimeout = types.Int64Null()
	}
	if val, ok := getResponseData["server"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Server = types.Int64Value(intVal)
		}
	} else {
		data.Server = types.Int64Null()
	}
	if val, ok := getResponseData["tcpclient"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tcpclient = types.Int64Value(intVal)
		}
	} else {
		data.Tcpclient = types.Int64Null()
	}
	if val, ok := getResponseData["tcpserver"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tcpserver = types.Int64Value(intVal)
		}
	} else {
		data.Tcpserver = types.Int64Null()
	}
	if val, ok := getResponseData["zombie"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Zombie = types.Int64Value(intVal)
		}
	} else {
		data.Zombie = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("nstimeout-config")

	return data
}
