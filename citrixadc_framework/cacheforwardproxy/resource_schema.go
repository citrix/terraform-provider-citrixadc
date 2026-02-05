package cacheforwardproxy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cache"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// CacheforwardproxyResourceModel describes the resource data model.
type CacheforwardproxyResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Ipaddress types.String `tfsdk:"ipaddress"`
	Port      types.Int64  `tfsdk:"port"`
}

func (r *CacheforwardproxyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cacheforwardproxy resource.",
			},
			"ipaddress": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP address of the Citrix ADC or a cache server for which the cache acts as a proxy. Requests coming to the Citrix ADC with the configured IP address are forwarded to the particular address, without involving the Integrated Cache in any way.",
			},
			"port": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Port on the Citrix ADC or a server for which the cache acts as a proxy",
			},
		},
	}
}

func cacheforwardproxyGetThePayloadFromtheConfig(ctx context.Context, data *CacheforwardproxyResourceModel) cache.Cacheforwardproxy {
	tflog.Debug(ctx, "In cacheforwardproxyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	cacheforwardproxy := cache.Cacheforwardproxy{}
	if !data.Ipaddress.IsNull() {
		cacheforwardproxy.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Port.IsNull() {
		cacheforwardproxy.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}

	return cacheforwardproxy
}

func cacheforwardproxySetAttrFromGet(ctx context.Context, data *CacheforwardproxyResourceModel, getResponseData map[string]interface{}) *CacheforwardproxyResourceModel {
	tflog.Debug(ctx, "In cacheforwardproxySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Ipaddress.ValueString())

	return data
}
