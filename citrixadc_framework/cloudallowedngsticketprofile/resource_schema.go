package cloudallowedngsticketprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cloud"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CloudallowedngsticketprofileResourceModel describes the resource data model.
type CloudallowedngsticketprofileResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Creator types.String `tfsdk:"creator"`
	Name    types.String `tfsdk:"name"`
}

func (r *CloudallowedngsticketprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cloudallowedngsticketprofile resource.",
			},
			"creator": schema.StringAttribute{
				Optional:    true,
				Description: "Created name for allowed tickets",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Profile name for allowed tickets",
			},
		},
	}
}

func cloudallowedngsticketprofileGetThePayloadFromthePlan(ctx context.Context, data *CloudallowedngsticketprofileResourceModel) cloud.Cloudallowedngsticketprofile {
	tflog.Debug(ctx, "In cloudallowedngsticketprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	cloudallowedngsticketprofile := cloud.Cloudallowedngsticketprofile{}
	if !data.Creator.IsNull() && !data.Creator.IsUnknown() {
		cloudallowedngsticketprofile.Creator = data.Creator.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		cloudallowedngsticketprofile.Name = data.Name.ValueString()
	}

	return cloudallowedngsticketprofile
}

func cloudallowedngsticketprofileSetAttrFromGet(ctx context.Context, data *CloudallowedngsticketprofileResourceModel, getResponseData map[string]interface{}) *CloudallowedngsticketprofileResourceModel {
	tflog.Debug(ctx, "In cloudallowedngsticketprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["creator"]; ok && val != nil {
		data.Creator = types.StringValue(val.(string))
	} else {
		data.Creator = types.StringNull()
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
