package appfwjsonerrorpage

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

// AppfwjsonerrorpageResourceModel describes the resource data model.
type AppfwjsonerrorpageResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Name      types.String `tfsdk:"name"`
	Overwrite types.Bool   `tfsdk:"overwrite"`
	Src       types.String `tfsdk:"src"`
}

func (r *AppfwjsonerrorpageResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwjsonerrorpage resource.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about the JSON error object.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Indicates name of the imported json error page to be removed.",
			},
			"overwrite": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Overwrite any existing JSON error object of the same name.",
			},
			"src": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL (protocol, host, path, and name) for the location at which to store the imported JSON error object.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
		},
	}
}

func appfwjsonerrorpageGetThePayloadFromtheConfig(ctx context.Context, data *AppfwjsonerrorpageResourceModel) appfw.Appfwjsonerrorpage {
	tflog.Debug(ctx, "In appfwjsonerrorpageGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwjsonerrorpage := appfw.Appfwjsonerrorpage{}
	if !data.Comment.IsNull() {
		appfwjsonerrorpage.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() {
		appfwjsonerrorpage.Name = data.Name.ValueString()
	}
	if !data.Overwrite.IsNull() {
		appfwjsonerrorpage.Overwrite = data.Overwrite.ValueBool()
	}
	if !data.Src.IsNull() {
		appfwjsonerrorpage.Src = data.Src.ValueString()
	}

	return appfwjsonerrorpage
}

func appfwjsonerrorpageSetAttrFromGet(ctx context.Context, data *AppfwjsonerrorpageResourceModel, getResponseData map[string]interface{}) *AppfwjsonerrorpageResourceModel {
	tflog.Debug(ctx, "In appfwjsonerrorpageSetAttrFromGet Function")

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
