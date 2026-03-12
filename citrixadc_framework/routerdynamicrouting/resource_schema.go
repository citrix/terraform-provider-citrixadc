package routerdynamicrouting

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/router"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// RouterdynamicroutingResourceModel describes the resource data model.
type RouterdynamicroutingResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Commandstring types.String `tfsdk:"commandstring"`
	Nodeid        types.Int64  `tfsdk:"nodeid"`
}

func (r *RouterdynamicroutingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the routerdynamicrouting resource.",
			},
			"commandstring": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "command to be executed",
			},
			"nodeid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
		},
	}
}

func routerdynamicroutingGetThePayloadFromtheConfig(ctx context.Context, data *RouterdynamicroutingResourceModel) router.Routerdynamicrouting {
	tflog.Debug(ctx, "In routerdynamicroutingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	routerdynamicrouting := router.Routerdynamicrouting{}
	if !data.Commandstring.IsNull() {
		routerdynamicrouting.Commandstring = data.Commandstring.ValueString()
	}
	if !data.Nodeid.IsNull() {
		routerdynamicrouting.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}

	return routerdynamicrouting
}

func routerdynamicroutingSetAttrFromGet(ctx context.Context, data *RouterdynamicroutingResourceModel, getResponseData map[string]interface{}) *RouterdynamicroutingResourceModel {
	tflog.Debug(ctx, "In routerdynamicroutingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["commandstring"]; ok && val != nil {
		data.Commandstring = types.StringValue(val.(string))
	} else {
		data.Commandstring = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else {
		data.Nodeid = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Commandstring.ValueString())

	return data
}
