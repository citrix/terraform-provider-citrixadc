package appalgparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppalgparamResourceModel describes the resource data model.
type AppalgparamResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Pptpgreidletimeout types.Int64  `tfsdk:"pptpgreidletimeout"`
}

func (r *AppalgparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appalgparam resource.",
			},
			"pptpgreidletimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(9000),
				Description: "Interval in sec, after which data sessions of PPTP GRE is cleared.",
			},
		},
	}
}

func appalgparamGetThePayloadFromtheConfig(ctx context.Context, data *AppalgparamResourceModel) network.Appalgparam {
	tflog.Debug(ctx, "In appalgparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appalgparam := network.Appalgparam{}
	if !data.Pptpgreidletimeout.IsNull() {
		appalgparam.Pptpgreidletimeout = utils.IntPtr(int(data.Pptpgreidletimeout.ValueInt64()))
	}

	return appalgparam
}

func appalgparamSetAttrFromGet(ctx context.Context, data *AppalgparamResourceModel, getResponseData map[string]interface{}) *AppalgparamResourceModel {
	tflog.Debug(ctx, "In appalgparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["pptpgreidletimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Pptpgreidletimeout = types.Int64Value(intVal)
		}
	} else {
		data.Pptpgreidletimeout = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("appalgparam-config")

	return data
}
