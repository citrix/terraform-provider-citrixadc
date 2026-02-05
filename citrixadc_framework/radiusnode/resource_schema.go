package radiusnode

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RadiusnodeResourceModel describes the resource data model.
type RadiusnodeResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Nodeprefix types.String `tfsdk:"nodeprefix"`
	Radkey     types.String `tfsdk:"radkey"`
}

func (r *RadiusnodeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the radiusnode resource.",
			},
			"nodeprefix": schema.StringAttribute{
				Required:    true,
				Description: "IP address/IP prefix of radius node in CIDR format",
			},
			"radkey": schema.StringAttribute{
				Required:    true,
				Description: "The key shared between the RADIUS server and clients.\n      Required for NetScaler to communicate with the RADIUS nodes.",
			},
		},
	}
}

func radiusnodeGetThePayloadFromtheConfig(ctx context.Context, data *RadiusnodeResourceModel) basic.Radiusnode {
	tflog.Debug(ctx, "In radiusnodeGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	radiusnode := basic.Radiusnode{}
	if !data.Nodeprefix.IsNull() {
		radiusnode.Nodeprefix = data.Nodeprefix.ValueString()
	}
	if !data.Radkey.IsNull() {
		radiusnode.Radkey = data.Radkey.ValueString()
	}

	return radiusnode
}

func radiusnodeSetAttrFromGet(ctx context.Context, data *RadiusnodeResourceModel, getResponseData map[string]interface{}) *RadiusnodeResourceModel {
	tflog.Debug(ctx, "In radiusnodeSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["nodeprefix"]; ok && val != nil {
		data.Nodeprefix = types.StringValue(val.(string))
	} else {
		data.Nodeprefix = types.StringNull()
	}
	if val, ok := getResponseData["radkey"]; ok && val != nil {
		data.Radkey = types.StringValue(val.(string))
	} else {
		data.Radkey = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Nodeprefix.ValueString())

	return data
}
