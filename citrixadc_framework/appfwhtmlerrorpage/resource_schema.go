package appfwhtmlerrorpage

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

// AppfwhtmlerrorpageResourceModel describes the resource data model.
type AppfwhtmlerrorpageResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Name      types.String `tfsdk:"name"`
	Overwrite types.Bool   `tfsdk:"overwrite"`
	Src       types.String `tfsdk:"src"`
}

func (r *AppfwhtmlerrorpageResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwhtmlerrorpage resource.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about the HTML error object.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the XML error object to remove.",
			},
			"overwrite": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Overwrite any existing HTML error object of the same name.",
			},
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL (protocol, host, path, and name) for the location at which to store the imported HTML error object.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
		},
	}
}

func appfwhtmlerrorpageGetThePayloadFromtheConfig(ctx context.Context, data *AppfwhtmlerrorpageResourceModel) appfw.Appfwhtmlerrorpage {
	tflog.Debug(ctx, "In appfwhtmlerrorpageGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwhtmlerrorpage := appfw.Appfwhtmlerrorpage{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwhtmlerrorpage.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwhtmlerrorpage.Name = data.Name.ValueString()
	}
	if !data.Overwrite.IsNull() && !data.Overwrite.IsUnknown() {
		appfwhtmlerrorpage.Overwrite = data.Overwrite.ValueBool()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		appfwhtmlerrorpage.Src = data.Src.ValueString()
	}

	return appfwhtmlerrorpage
}

// appfwhtmlerrorpageSetAttrFromGet updates the resource state after a create/read.
// The NITRO GET for appfwhtmlerrorpage only returns "name" (and read-only fields);
// it does NOT return comment, overwrite, or src. Mirroring the SDK v2 resource
// (which only refreshes "name"), we refresh name from the API response and preserve
// the plan/state values for the remaining attributes to avoid spurious diffs and
// "inconsistent result after apply" errors.
func appfwhtmlerrorpageSetAttrFromGet(ctx context.Context, data *AppfwhtmlerrorpageResourceModel, getResponseData map[string]interface{}) *AppfwhtmlerrorpageResourceModel {
	tflog.Debug(ctx, "In appfwhtmlerrorpageSetAttrFromGet Function")

	// name: refresh from the API response (primary key)
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}

	// comment: not returned by the NITRO GET - preserve plan/state value,
	// resolving an unknown (Optional+Computed, omitted in config) to null.
	if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}

	// overwrite: not returned by the NITRO GET - preserve plan/state value,
	// resolving an unknown (Optional+Computed, omitted in config) to null.
	if data.Overwrite.IsUnknown() {
		data.Overwrite = types.BoolNull()
	}

	// src: Required + ForceNew - not refreshed by SDK v2; preserve plan/state value.

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}

// appfwhtmlerrorpageSetAttrFromGetForDatasource populates the datasource state from
// the NITRO GET response, copying every attribute the API returns (and nulling the
// ones it does not) and setting the ID.
func appfwhtmlerrorpageSetAttrFromGetForDatasource(ctx context.Context, data *AppfwhtmlerrorpageResourceModel, getResponseData map[string]interface{}) *AppfwhtmlerrorpageResourceModel {
	tflog.Debug(ctx, "In appfwhtmlerrorpageSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["overwrite"]; ok && val != nil {
		data.Overwrite = types.BoolValue(val.(bool))
	} else {
		data.Overwrite = types.BoolNull()
	}
	if val, ok := getResponseData["src"]; ok && val != nil {
		data.Src = types.StringValue(val.(string))
	} else {
		data.Src = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
