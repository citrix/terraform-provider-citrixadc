package appfwxmlerrorpage

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

// AppfwxmlerrorpageResourceModel describes the resource data model.
type AppfwxmlerrorpageResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Name      types.String `tfsdk:"name"`
	Overwrite types.Bool   `tfsdk:"overwrite"`
	Src       types.String `tfsdk:"src"`
}

func (r *AppfwxmlerrorpageResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwxmlerrorpage resource.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about the XML error object.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Indicates name of the imported xml error page to be removed.",
			},
			"overwrite": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Overwrite any existing XML error object of the same name.",
			},
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL (protocol, host, path, and name) for the location at which to store the imported XML error object.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
		},
	}
}

func appfwxmlerrorpageGetThePayloadFromtheConfig(ctx context.Context, data *AppfwxmlerrorpageResourceModel) appfw.Appfwxmlerrorpage {
	tflog.Debug(ctx, "In appfwxmlerrorpageGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwxmlerrorpage := appfw.Appfwxmlerrorpage{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwxmlerrorpage.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwxmlerrorpage.Name = data.Name.ValueString()
	}
	if !data.Overwrite.IsNull() && !data.Overwrite.IsUnknown() {
		appfwxmlerrorpage.Overwrite = data.Overwrite.ValueBool()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		appfwxmlerrorpage.Src = data.Src.ValueString()
	}

	return appfwxmlerrorpage
}

// appfwxmlerrorpageSetAttrFromGet is used by the resource. It mirrors the SDK v2
// read semantics, which only refresh the name from the API. The NITRO GET does
// not return comment/overwrite and returns src in a normalized (basename) form,
// so the configured values are preserved to remain backward compatible and to
// avoid "inconsistent result after apply" errors.
func appfwxmlerrorpageSetAttrFromGet(ctx context.Context, data *AppfwxmlerrorpageResourceModel, getResponseData map[string]interface{}) *AppfwxmlerrorpageResourceModel {
	tflog.Debug(ctx, "In appfwxmlerrorpageSetAttrFromGet Function")

	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}

	// Resolve unknown values for Optional+Computed attributes that are not
	// returned by the NITRO GET so that a known value is always written to state.
	if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	if data.Overwrite.IsUnknown() {
		data.Overwrite = types.BoolNull()
	}
	if data.Src.IsUnknown() {
		data.Src = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}

// appfwxmlerrorpageSetAttrFromGetForDatasource is used by the datasource. It
// copies every attribute returned by the NITRO GET into the model (Pattern 7
// datasource split).
func appfwxmlerrorpageSetAttrFromGetForDatasource(ctx context.Context, data *AppfwxmlerrorpageResourceModel, getResponseData map[string]interface{}) *AppfwxmlerrorpageResourceModel {
	tflog.Debug(ctx, "In appfwxmlerrorpageSetAttrFromGetForDatasource Function")

	// Convert API response to model
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
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
