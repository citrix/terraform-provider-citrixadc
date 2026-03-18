package vxlanvlanmap

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VxlanvlanmapResourceModel describes the resource data model.
type VxlanvlanmapResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (r *VxlanvlanmapResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vxlanvlanmap resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the mapping table.",
			},
		},
	}
}

func vxlanvlanmapGetThePayloadFromtheConfig(ctx context.Context, data *VxlanvlanmapResourceModel) network.Vxlanvlanmap {
	tflog.Debug(ctx, "In vxlanvlanmapGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vxlanvlanmap := network.Vxlanvlanmap{}
	if !data.Name.IsNull() {
		vxlanvlanmap.Name = data.Name.ValueString()
	}

	return vxlanvlanmap
}

func vxlanvlanmapSetAttrFromGet(ctx context.Context, data *VxlanvlanmapResourceModel, getResponseData map[string]interface{}) *VxlanvlanmapResourceModel {
	tflog.Debug(ctx, "In vxlanvlanmapSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
