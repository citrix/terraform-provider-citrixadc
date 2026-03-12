package nsdiameter

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

// NsdiameterResourceModel describes the resource data model.
type NsdiameterResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Identity               types.String `tfsdk:"identity"`
	Ownernode              types.Int64  `tfsdk:"ownernode"`
	Realm                  types.String `tfsdk:"realm"`
	Serverclosepropagation types.String `tfsdk:"serverclosepropagation"`
}

func (r *NsdiameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsdiameter resource.",
			},
			"identity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DiameterIdentity to be used by NS. DiameterIdentity is used to identify a Diameter node uniquely. Before setting up diameter configuration, Citrix ADC (as a Diameter node) MUST be assigned a unique DiameterIdentity.\nexample =>\nset ns diameter -identity netscaler.com\nNow whenever Citrix ADC needs to use identity in diameter messages. It will use 'netscaler.com' as Origin-Host AVP as defined in RFC3588",
			},
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(-1),
				Description: "ID of the cluster node for which the diameter id is set, can be configured only through CLIP",
			},
			"realm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Diameter Realm to be used by NS.\nexample =>\nset ns diameter -realm com\nNow whenever Citrix ADC system needs to use realm in diameter messages. It will use 'com' as Origin-Realm AVP as defined in RFC3588",
			},
			"serverclosepropagation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "when a Server connection goes down, whether to close the corresponding client connection if there were requests pending on the server.",
			},
		},
	}
}

func nsdiameterGetThePayloadFromtheConfig(ctx context.Context, data *NsdiameterResourceModel) ns.Nsdiameter {
	tflog.Debug(ctx, "In nsdiameterGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsdiameter := ns.Nsdiameter{}
	if !data.Identity.IsNull() {
		nsdiameter.Identity = data.Identity.ValueString()
	}
	if !data.Ownernode.IsNull() {
		nsdiameter.Ownernode = utils.IntPtr(int(data.Ownernode.ValueInt64()))
	}
	if !data.Realm.IsNull() {
		nsdiameter.Realm = data.Realm.ValueString()
	}
	if !data.Serverclosepropagation.IsNull() {
		nsdiameter.Serverclosepropagation = data.Serverclosepropagation.ValueString()
	}

	return nsdiameter
}

func nsdiameterSetAttrFromGet(ctx context.Context, data *NsdiameterResourceModel, getResponseData map[string]interface{}) *NsdiameterResourceModel {
	tflog.Debug(ctx, "In nsdiameterSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["identity"]; ok && val != nil {
		data.Identity = types.StringValue(val.(string))
	} else {
		data.Identity = types.StringNull()
	}
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		}
	} else {
		data.Ownernode = types.Int64Null()
	}
	if val, ok := getResponseData["realm"]; ok && val != nil {
		data.Realm = types.StringValue(val.(string))
	} else {
		data.Realm = types.StringNull()
	}
	if val, ok := getResponseData["serverclosepropagation"]; ok && val != nil {
		data.Serverclosepropagation = types.StringValue(val.(string))
	} else {
		data.Serverclosepropagation = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Ownernode.ValueInt64()))

	return data
}
