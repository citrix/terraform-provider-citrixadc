package quicbridgeprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/quicbridge"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// QuicbridgeprofileResourceModel describes the resource data model.
type QuicbridgeprofileResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Routingalgorithm types.String `tfsdk:"routingalgorithm"`
	Serveridlength   types.Int64  `tfsdk:"serveridlength"`
}

func (r *QuicbridgeprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the quicbridgeprofile resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the QUIC profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@),equals sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.",
			},
			"routingalgorithm": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("PLAINTEXT"),
				Description: "Routing algorithm to generate routable connection IDs.",
			},
			"serveridlength": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(4),
				Description: "Length of serverid to encode/decode server information",
			},
		},
	}
}

func quicbridgeprofileGetThePayloadFromtheConfig(ctx context.Context, data *QuicbridgeprofileResourceModel) quicbridge.Quicbridgeprofile {
	tflog.Debug(ctx, "In quicbridgeprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	quicbridgeprofile := quicbridge.Quicbridgeprofile{}
	if !data.Name.IsNull() {
		quicbridgeprofile.Name = data.Name.ValueString()
	}
	if !data.Routingalgorithm.IsNull() {
		quicbridgeprofile.Routingalgorithm = data.Routingalgorithm.ValueString()
	}
	if !data.Serveridlength.IsNull() {
		quicbridgeprofile.Serveridlength = utils.IntPtr(int(data.Serveridlength.ValueInt64()))
	}

	return quicbridgeprofile
}

func quicbridgeprofileSetAttrFromGet(ctx context.Context, data *QuicbridgeprofileResourceModel, getResponseData map[string]interface{}) *QuicbridgeprofileResourceModel {
	tflog.Debug(ctx, "In quicbridgeprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["routingalgorithm"]; ok && val != nil {
		data.Routingalgorithm = types.StringValue(val.(string))
	} else {
		data.Routingalgorithm = types.StringNull()
	}
	if val, ok := getResponseData["serveridlength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Serveridlength = types.Int64Value(intVal)
		}
	} else {
		data.Serveridlength = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
