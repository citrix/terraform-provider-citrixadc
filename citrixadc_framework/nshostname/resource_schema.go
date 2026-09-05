package nshostname

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NshostnameResourceModel describes the resource data model.
type NshostnameResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Hostname  types.String `tfsdk:"hostname"`
	Ownernode types.Int64  `tfsdk:"ownernode"`
}

func (r *NshostnameResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nshostname resource.",
			},
			"hostname": schema.StringAttribute{
				Required:    true,
				Description: "Host name for the Citrix ADC.",
			},
			// ownernode was Optional+Computed (no Default, no ForceNew) in SDK v2.
			// A Default is invalid without Computed, and SDK v2 had no default, so we
			// read the effective value back from the ADC.
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the cluster node for which you are setting the hostname. Can be configured only through the cluster IP address.",
			},
		},
	}
}

func nshostnameGetThePayloadFromtheConfig(ctx context.Context, data *NshostnameResourceModel) ns.Nshostname {
	tflog.Debug(ctx, "In nshostnameGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nshostname := ns.Nshostname{}
	if !data.Hostname.IsNull() && !data.Hostname.IsUnknown() {
		nshostname.Hostname = data.Hostname.ValueString()
	}
	if !data.Ownernode.IsNull() && !data.Ownernode.IsUnknown() {
		nshostname.Ownernode = utils.IntPtr(int(data.Ownernode.ValueInt64()))
	}

	return nshostname
}

func nshostnameSetAttrFromGet(ctx context.Context, data *NshostnameResourceModel, getResponseData map[string]interface{}) *NshostnameResourceModel {
	tflog.Debug(ctx, "In nshostnameSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["hostname"]; ok && val != nil {
		data.Hostname = types.StringValue(val.(string))
	} else if data.Hostname.IsUnknown() {
		data.Hostname = types.StringNull()
	}
	// Do not clobber a known/configured ownernode when NITRO omits it from GET
	// (omit-on-default trap). Only null it when the value is unknown.
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		}
	} else if data.Ownernode.IsUnknown() {
		data.Ownernode = types.Int64Null()
	}

	// Set ID for the resource
	// nshostname is a singleton resource - use a static ID
	data.Id = types.StringValue("nshostname-config")

	return data
}
