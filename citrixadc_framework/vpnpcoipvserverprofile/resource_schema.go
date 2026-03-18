package vpnpcoipvserverprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpnpcoipvserverprofileResourceModel describes the resource data model.
type VpnpcoipvserverprofileResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Logindomain types.String `tfsdk:"logindomain"`
	Name        types.String `tfsdk:"name"`
	Udpport     types.Int64  `tfsdk:"udpport"`
}

func (r *VpnpcoipvserverprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnpcoipvserverprofile resource.",
			},
			"logindomain": schema.StringAttribute{
				Required:    true,
				Description: "Login domain for PCoIP users",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "name of PCoIP vserver profile",
			},
			"udpport": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(4172),
				Description: "UDP port for PCoIP data traffic",
			},
		},
	}
}

func vpnpcoipvserverprofileGetThePayloadFromtheConfig(ctx context.Context, data *VpnpcoipvserverprofileResourceModel) vpn.Vpnpcoipvserverprofile {
	tflog.Debug(ctx, "In vpnpcoipvserverprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnpcoipvserverprofile := vpn.Vpnpcoipvserverprofile{}
	if !data.Logindomain.IsNull() {
		vpnpcoipvserverprofile.Logindomain = data.Logindomain.ValueString()
	}
	if !data.Name.IsNull() {
		vpnpcoipvserverprofile.Name = data.Name.ValueString()
	}
	if !data.Udpport.IsNull() {
		vpnpcoipvserverprofile.Udpport = utils.IntPtr(int(data.Udpport.ValueInt64()))
	}

	return vpnpcoipvserverprofile
}

func vpnpcoipvserverprofileSetAttrFromGet(ctx context.Context, data *VpnpcoipvserverprofileResourceModel, getResponseData map[string]interface{}) *VpnpcoipvserverprofileResourceModel {
	tflog.Debug(ctx, "In vpnpcoipvserverprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["logindomain"]; ok && val != nil {
		data.Logindomain = types.StringValue(val.(string))
	} else {
		data.Logindomain = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["udpport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Udpport = types.Int64Value(intVal)
		}
	} else {
		data.Udpport = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
