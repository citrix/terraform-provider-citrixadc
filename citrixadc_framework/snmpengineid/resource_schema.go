package snmpengineid

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SnmpengineidResourceModel describes the resource data model.
type SnmpengineidResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Engineid  types.String `tfsdk:"engineid"`
	Ownernode types.Int64  `tfsdk:"ownernode"`
}

func (r *SnmpengineidResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the snmpengineid resource.",
			},
			"engineid": schema.StringAttribute{
				Required:    true,
				Description: "A hexadecimal value of at least 10 characters, uniquely identifying the engineid",
			},
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(-1),
				Description: "ID of the cluster node for which you are setting the engineid",
			},
		},
	}
}

func snmpengineidGetThePayloadFromtheConfig(ctx context.Context, data *SnmpengineidResourceModel) snmp.Snmpengineid {
	tflog.Debug(ctx, "In snmpengineidGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	snmpengineid := snmp.Snmpengineid{}
	if !data.Engineid.IsNull() {
		snmpengineid.Engineid = data.Engineid.ValueString()
	}
	if !data.Ownernode.IsNull() {
		snmpengineid.Ownernode = utils.IntPtr(int(data.Ownernode.ValueInt64()))
	}

	return snmpengineid
}

func snmpengineidSetAttrFromGet(ctx context.Context, data *SnmpengineidResourceModel, getResponseData map[string]interface{}) *SnmpengineidResourceModel {
	tflog.Debug(ctx, "In snmpengineidSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["engineid"]; ok && val != nil {
		data.Engineid = types.StringValue(val.(string))
	} else {
		data.Engineid = types.StringNull()
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
