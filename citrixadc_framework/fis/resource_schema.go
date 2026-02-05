package fis

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// FisResourceModel describes the resource data model.
type FisResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Ownernode types.Int64  `tfsdk:"ownernode"`
}

func (r *FisResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the fis resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the FIS to be created. Leading character must be a number or letter. Other characters allowed, after the first character, are @ _ - . (period) : (colon) # and space ( ). Note: In a cluster setup, the FIS name on each node must be unique.",
			},
			"ownernode": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "ID of the cluster node for which you are creating the FIS. Can be configured only through the cluster IP address.",
			},
		},
	}
}

func fisGetThePayloadFromtheConfig(ctx context.Context, data *FisResourceModel) network.Fis {
	tflog.Debug(ctx, "In fisGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	fis := network.Fis{}
	if !data.Name.IsNull() {
		fis.Name = data.Name.ValueString()
	}
	if !data.Ownernode.IsNull() {
		fis.Ownernode = utils.IntPtr(int(data.Ownernode.ValueInt64()))
	}

	return fis
}

func fisSetAttrFromGet(ctx context.Context, data *FisResourceModel, getResponseData map[string]interface{}) *FisResourceModel {
	tflog.Debug(ctx, "In fisSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
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
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
