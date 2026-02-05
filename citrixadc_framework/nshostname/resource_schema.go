package nshostname

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
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
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(255),
				Description: "ID of the cluster node for which you are setting the hostname. Can be configured only through the cluster IP address.",
			},
		},
	}
}

func nshostnameGetThePayloadFromtheConfig(ctx context.Context, data *NshostnameResourceModel) ns.Nshostname {
	tflog.Debug(ctx, "In nshostnameGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nshostname := ns.Nshostname{}
	if !data.Hostname.IsNull() {
		nshostname.Hostname = data.Hostname.ValueString()
	}
	if !data.Ownernode.IsNull() {
		nshostname.Ownernode = utils.IntPtr(int(data.Ownernode.ValueInt64()))
	}

	return nshostname
}

func nshostnameSetAttrFromGet(ctx context.Context, data *NshostnameResourceModel, getResponseData map[string]interface{}) *NshostnameResourceModel {
	tflog.Debug(ctx, "In nshostnameSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["hostname"]; ok && val != nil {
		data.Hostname = types.StringValue(val.(string))
	} else {
		data.Hostname = types.StringNull()
	}
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		}
	} else {
		data.Ownernode = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Ownernode.ValueInt64()))

	return data
}
