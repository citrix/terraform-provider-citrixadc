package iptunnelparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// IptunnelparamResourceModel describes the resource data model.
type IptunnelparamResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Dropfrag             types.String `tfsdk:"dropfrag"`
	Dropfragcputhreshold types.Int64  `tfsdk:"dropfragcputhreshold"`
	Enablestrictrx       types.String `tfsdk:"enablestrictrx"`
	Enablestricttx       types.String `tfsdk:"enablestricttx"`
	Mac                  types.String `tfsdk:"mac"`
	Srcip                types.String `tfsdk:"srcip"`
	Srciproundrobin      types.String `tfsdk:"srciproundrobin"`
	Useclientsourceip    types.String `tfsdk:"useclientsourceip"`
}

func (r *IptunnelparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the iptunnelparam resource.",
			},
			"dropfrag": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NO"),
				Description: "Drop any IP packet that requires fragmentation before it is sent through the tunnel.",
			},
			"dropfragcputhreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
				Description: "Threshold value, as a percentage of CPU usage, at which to drop packets that require fragmentation to use the IP tunnel. Applies only if dropFragparameter is set to NO. The default value, 0, specifies that this parameter is not set.",
			},
			"enablestrictrx": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NO"),
				Description: "Strict PBR check for IPSec packets received through tunnel",
			},
			"enablestricttx": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NO"),
				Description: "Strict PBR check for packets to be sent IPSec protected",
			},
			"mac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The shared MAC used for shared IP between cluster nodes/HA peers",
			},
			"srcip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Common source-IP address for all tunnels. For a specific tunnel, this global setting is overridden if you have specified another source IP address. Must be a MIP or SNIP address.",
			},
			"srciproundrobin": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NO"),
				Description: "Use a different source IP address for each new session through a particular IP tunnel, as determined by round robin selection of one of the SNIP addresses. This setting is ignored if a common global source IP address has been specified for all the IP tunnels. This setting does not apply to a tunnel for which a source IP address has been specified.",
			},
			"useclientsourceip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NO"),
				Description: "Use client source IP as source IP for outer tunnel IP header",
			},
		},
	}
}

func iptunnelparamGetThePayloadFromtheConfig(ctx context.Context, data *IptunnelparamResourceModel) network.Iptunnelparam {
	tflog.Debug(ctx, "In iptunnelparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	iptunnelparam := network.Iptunnelparam{}
	if !data.Dropfrag.IsNull() && !data.Dropfrag.IsUnknown() {
		iptunnelparam.Dropfrag = data.Dropfrag.ValueString()
	}
	if !data.Dropfragcputhreshold.IsNull() && !data.Dropfragcputhreshold.IsUnknown() {
		iptunnelparam.Dropfragcputhreshold = utils.IntPtr(int(data.Dropfragcputhreshold.ValueInt64()))
	}
	if !data.Enablestrictrx.IsNull() && !data.Enablestrictrx.IsUnknown() {
		iptunnelparam.Enablestrictrx = data.Enablestrictrx.ValueString()
	}
	if !data.Enablestricttx.IsNull() && !data.Enablestricttx.IsUnknown() {
		iptunnelparam.Enablestricttx = data.Enablestricttx.ValueString()
	}
	if !data.Mac.IsNull() && !data.Mac.IsUnknown() {
		iptunnelparam.Mac = data.Mac.ValueString()
	}
	if !data.Srcip.IsNull() && !data.Srcip.IsUnknown() {
		iptunnelparam.Srcip = data.Srcip.ValueString()
	}
	if !data.Srciproundrobin.IsNull() && !data.Srciproundrobin.IsUnknown() {
		iptunnelparam.Srciproundrobin = data.Srciproundrobin.ValueString()
	}
	if !data.Useclientsourceip.IsNull() && !data.Useclientsourceip.IsUnknown() {
		iptunnelparam.Useclientsourceip = data.Useclientsourceip.ValueString()
	}

	return iptunnelparam
}

func iptunnelparamSetAttrFromGet(ctx context.Context, data *IptunnelparamResourceModel, getResponseData map[string]interface{}) *IptunnelparamResourceModel {
	tflog.Debug(ctx, "In iptunnelparamSetAttrFromGet Function")

	// Convert API response to model.
	// NOTE (omit-on-default trap): NITRO may omit an attribute from GET when it holds
	// its default value (e.g. "NO"/0). Only null an omitted attribute when it is Unknown
	// (an unresolved Computed value); never clobber a known configured/state value, or
	// Terraform reports "inconsistent result after apply".
	if val, ok := getResponseData["dropfrag"]; ok && val != nil {
		data.Dropfrag = types.StringValue(val.(string))
	} else if data.Dropfrag.IsUnknown() {
		data.Dropfrag = types.StringNull()
	}
	if val, ok := getResponseData["dropfragcputhreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Dropfragcputhreshold = types.Int64Value(intVal)
		}
	} else if data.Dropfragcputhreshold.IsUnknown() {
		data.Dropfragcputhreshold = types.Int64Null()
	}
	if val, ok := getResponseData["enablestrictrx"]; ok && val != nil {
		data.Enablestrictrx = types.StringValue(val.(string))
	} else if data.Enablestrictrx.IsUnknown() {
		data.Enablestrictrx = types.StringNull()
	}
	if val, ok := getResponseData["enablestricttx"]; ok && val != nil {
		data.Enablestricttx = types.StringValue(val.(string))
	} else if data.Enablestricttx.IsUnknown() {
		data.Enablestricttx = types.StringNull()
	}
	if val, ok := getResponseData["mac"]; ok && val != nil {
		data.Mac = types.StringValue(val.(string))
	} else if data.Mac.IsUnknown() {
		data.Mac = types.StringNull()
	}
	if val, ok := getResponseData["srcip"]; ok && val != nil {
		data.Srcip = types.StringValue(val.(string))
	} else if data.Srcip.IsUnknown() {
		data.Srcip = types.StringNull()
	}
	if val, ok := getResponseData["srciproundrobin"]; ok && val != nil {
		data.Srciproundrobin = types.StringValue(val.(string))
	} else if data.Srciproundrobin.IsUnknown() {
		data.Srciproundrobin = types.StringNull()
	}
	if val, ok := getResponseData["useclientsourceip"]; ok && val != nil {
		data.Useclientsourceip = types.StringValue(val.(string))
	} else if data.Useclientsourceip.IsUnknown() {
		data.Useclientsourceip = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("iptunnelparam-config")

	return data
}
