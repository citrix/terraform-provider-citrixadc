package ip6tunnelparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ip6tunnelparamResourceModel describes the resource data model.
type Ip6tunnelparamResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Dropfrag             types.String `tfsdk:"dropfrag"`
	Dropfragcputhreshold types.Int64  `tfsdk:"dropfragcputhreshold"`
	Srcip                types.String `tfsdk:"srcip"`
	Srciproundrobin      types.String `tfsdk:"srciproundrobin"`
	Useclientsourceipv6  types.String `tfsdk:"useclientsourceipv6"`
}

func (r *Ip6tunnelparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ip6tunnelparam resource.",
			},
			"dropfrag": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Drop any packet that requires fragmentation.",
			},
			"dropfragcputhreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Threshold value, as a percentage of CPU usage, at which to drop packets that require fragmentation. Applies only if dropFragparameter is set to NO.",
			},
			"srcip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Common source IPv6 address for all IPv6 tunnels. Must be a SNIP6 or VIP6 address.",
			},
			"srciproundrobin": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use a different source IPv6 address for each new session through a particular IPv6 tunnel, as determined by round robin selection of one of the SNIP6 addresses. This setting is ignored if a common global source IPv6 address has been specified for all the IPv6 tunnels. This setting does not apply to a tunnel for which a source IPv6 address has been specified.",
			},
			"useclientsourceipv6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use client source IPv6 address as source IPv6 address for outer tunnel IPv6 header",
			},
		},
	}
}

func ip6tunnelparamGetThePayloadFromtheConfig(ctx context.Context, data *Ip6tunnelparamResourceModel) network.Ip6tunnelparam {
	tflog.Debug(ctx, "In ip6tunnelparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	ip6tunnelparam := network.Ip6tunnelparam{}
	if !data.Dropfrag.IsNull() {
		ip6tunnelparam.Dropfrag = data.Dropfrag.ValueString()
	}
	if !data.Dropfragcputhreshold.IsNull() {
		ip6tunnelparam.Dropfragcputhreshold = utils.IntPtr(int(data.Dropfragcputhreshold.ValueInt64()))
	}
	if !data.Srcip.IsNull() {
		ip6tunnelparam.Srcip = data.Srcip.ValueString()
	}
	if !data.Srciproundrobin.IsNull() {
		ip6tunnelparam.Srciproundrobin = data.Srciproundrobin.ValueString()
	}
	if !data.Useclientsourceipv6.IsNull() {
		ip6tunnelparam.Useclientsourceipv6 = data.Useclientsourceipv6.ValueString()
	}

	return ip6tunnelparam
}

func ip6tunnelparamSetAttrFromGet(ctx context.Context, data *Ip6tunnelparamResourceModel, getResponseData map[string]interface{}) *Ip6tunnelparamResourceModel {
	tflog.Debug(ctx, "In ip6tunnelparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["dropfrag"]; ok && val != nil {
		data.Dropfrag = types.StringValue(val.(string))
	} else {
		data.Dropfrag = types.StringNull()
	}
	if val, ok := getResponseData["dropfragcputhreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Dropfragcputhreshold = types.Int64Value(intVal)
		}
	} else {
		data.Dropfragcputhreshold = types.Int64Null()
	}
	if val, ok := getResponseData["srcip"]; ok && val != nil {
		data.Srcip = types.StringValue(val.(string))
	} else {
		data.Srcip = types.StringNull()
	}
	if val, ok := getResponseData["srciproundrobin"]; ok && val != nil {
		data.Srciproundrobin = types.StringValue(val.(string))
	} else {
		data.Srciproundrobin = types.StringNull()
	}
	if val, ok := getResponseData["useclientsourceipv6"]; ok && val != nil {
		data.Useclientsourceipv6 = types.StringValue(val.(string))
	} else {
		data.Useclientsourceipv6 = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("ip6tunnelparam-config")

	return data
}
