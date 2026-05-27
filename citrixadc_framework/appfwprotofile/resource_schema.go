package appfwprotofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwprotofileResourceModel describes the resource data model.
type AppfwprotofileResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Name      types.String `tfsdk:"name"`
	Overwrite types.Bool   `tfsdk:"overwrite"`
	Src       types.String `tfsdk:"src"`
}

func (r *AppfwprotofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprotofile resource.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Comments associated with this gRPC schema file.",
			},
			// NITRO Import payload marks `name` as mandatory (red/bold).
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the gRPC schema object.",
			},
			"overwrite": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Overwrite any existing gRPC schema object of the same name.",
			},
			// NITRO Import payload marks `src` as mandatory (red/bold).
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Indicates source path of the gRPC schema file.",
			},
		},
	}
}

func appfwprotofileGetThePayloadFromthePlan(ctx context.Context, data *AppfwprotofileResourceModel) appfw.Appfwprotofile {
	tflog.Debug(ctx, "In appfwprotofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwprotofile := appfw.Appfwprotofile{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprotofile.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprotofile.Name = data.Name.ValueString()
	}
	if !data.Overwrite.IsNull() && !data.Overwrite.IsUnknown() {
		appfwprotofile.Overwrite = data.Overwrite.ValueBool()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		appfwprotofile.Src = data.Src.ValueString()
	}

	return appfwprotofile
}

func appfwprotofileSetAttrFromGet(ctx context.Context, data *AppfwprotofileResourceModel, getResponseData map[string]interface{}) *AppfwprotofileResourceModel {
	tflog.Debug(ctx, "In appfwprotofileSetAttrFromGet Function")

	// NITRO's appfwprotofile `get` response payload only carries `name`,
	// `response`, `src`, and `_nextgenapiresource`. The user-supplied write-only
	// Import inputs `comment` and `overwrite` are NEVER echoed back. Touching
	// them here would null them on every Read and cause a perpetual diff, so
	// only update `name` and `src` (the response-side fields) from the API.
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	// `src` is the URL supplied at Import time; the NITRO server normalizes it
	// (e.g., strips the "local:" protocol prefix) and returns the bare filename.
	// Preserve the user-configured value to avoid an "inconsistent result after
	// apply" error from Terraform.

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
