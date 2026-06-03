package vpnepaprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnepaprofileResourceModel describes the resource data model.
type VpnepaprofileResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Data     types.String `tfsdk:"data"`
	Filename types.String `tfsdk:"filename"`
	Name     types.String `tfsdk:"name"`
}

func (r *VpnepaprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnepaprofile resource.",
			},
			"data": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "deviceprofile data xml",
			},
			"filename": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "filename of the deviceprofile data xml",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "name of device profile",
			},
		},
	}
}

func vpnepaprofileGetThePayloadFromthePlan(ctx context.Context, data *VpnepaprofileResourceModel) vpn.Vpnepaprofile {
	tflog.Debug(ctx, "In vpnepaprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnepaprofile := vpn.Vpnepaprofile{}
	if !data.Data.IsNull() && !data.Data.IsUnknown() {
		vpnepaprofile.Data = data.Data.ValueString()
	}
	if !data.Filename.IsNull() && !data.Filename.IsUnknown() {
		vpnepaprofile.Filename = data.Filename.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnepaprofile.Name = data.Name.ValueString()
	}

	return vpnepaprofile
}

func vpnepaprofileSetAttrFromGet(ctx context.Context, data *VpnepaprofileResourceModel, getResponseData map[string]interface{}) *VpnepaprofileResourceModel {
	tflog.Debug(ctx, "In vpnepaprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["data"]; ok && val != nil {
		data.Data = types.StringValue(val.(string))
	} else {
		data.Data = types.StringNull()
	}
	if val, ok := getResponseData["filename"]; ok && val != nil {
		data.Filename = types.StringValue(val.(string))
	} else {
		data.Filename = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
