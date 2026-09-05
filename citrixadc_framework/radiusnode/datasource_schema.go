package radiusnode

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func RadiusnodeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"nodeprefix": schema.StringAttribute{
				Required:    true,
				Description: "IP address/IP prefix of radius node in CIDR format",
			},
			"radkey": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "The key shared between the RADIUS server and clients.\n      Required for NetScaler to communicate with the RADIUS nodes.",
			},
		},
	}
}

type RadiusnodeDataSourceModel struct {
	Id         types.String `tfsdk:"id"`
	Nodeprefix types.String `tfsdk:"nodeprefix"`
	Radkey     types.String `tfsdk:"radkey"`
}

func radiusnodeDataSourceSetAttrFromGet(ctx context.Context, data *RadiusnodeDataSourceModel, getResponseData map[string]interface{}) *RadiusnodeDataSourceModel {
	tflog.Debug(ctx, "In radiusnodeDataSourceSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["nodeprefix"]; ok && val != nil {
		data.Nodeprefix = types.StringValue(val.(string))
	} else {
		data.Nodeprefix = types.StringNull()
	}
	// radkey is not returned by NITRO API (secret/ephemeral) - retain from config

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Nodeprefix.ValueString()))

	return data
}
