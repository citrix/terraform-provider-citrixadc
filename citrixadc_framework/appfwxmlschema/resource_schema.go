package appfwxmlschema

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

// AppfwxmlschemaResourceModel describes the resource data model.
type AppfwxmlschemaResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Name      types.String `tfsdk:"name"`
	Overwrite types.Bool   `tfsdk:"overwrite"`
	Src       types.String `tfsdk:"src"`
}

func (r *AppfwxmlschemaResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwxmlschema resource.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about the XML Schema object.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the XML Schema object to remove.",
			},
			"overwrite": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Overwrite any existing XML Schema object of the same name.",
			},
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL (protocol, host, path, and file name) for the location at which to store the imported XML Schema.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
		},
	}
}

func appfwxmlschemaGetThePayloadFromtheConfig(ctx context.Context, data *AppfwxmlschemaResourceModel) appfw.Appfwxmlschema {
	tflog.Debug(ctx, "In appfwxmlschemaGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwxmlschema := appfw.Appfwxmlschema{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwxmlschema.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwxmlschema.Name = data.Name.ValueString()
	}
	if !data.Overwrite.IsNull() && !data.Overwrite.IsUnknown() {
		appfwxmlschema.Overwrite = data.Overwrite.ValueBool()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		appfwxmlschema.Src = data.Src.ValueString()
	}

	return appfwxmlschema
}

func appfwxmlschemaSetAttrFromGet(ctx context.Context, data *AppfwxmlschemaResourceModel, getResponseData map[string]interface{}) *AppfwxmlschemaResourceModel {
	tflog.Debug(ctx, "In appfwxmlschemaSetAttrFromGet Function")

	// Convert API response to model.
	// NOTE: the appfwxmlschema GET API only returns "name" (plus read-only
	// "response"/"_nextgenapiresource"). The "comment", "src" and "overwrite"
	// attributes are import-only and are NOT echoed back, so they must be
	// retained from the plan/state rather than overwritten with null.
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}

	// Resolve any unresolved (unknown) Optional+Computed values so the saved
	// state never contains an unknown value (avoids "inconsistent result after
	// apply"). These fields are not returned by the API.
	if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	if data.Src.IsUnknown() {
		data.Src = types.StringNull()
	}
	if data.Overwrite.IsUnknown() {
		data.Overwrite = types.BoolNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
