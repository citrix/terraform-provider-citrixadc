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
			// SDK v2: Optional + Computed + ForceNew
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					// GH #1436: create-only (no NITRO update op)
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Any comments to preserve information about the JSON error object.",
			},
			// SDK v2: Required + ForceNew (primary key)
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Indicates name of the imported json error page to be removed.",
			},
			// SDK v2: Optional + Computed + ForceNew
			"overwrite": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					// GH #1436: create-only (no NITRO update op)
					boolplanmodifier.UseStateForUnknown(),
					boolplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Overwrite any existing JSON error object of the same name.",
			},
			// SDK v2: Required + ForceNew
			"src": schema.StringAttribute{
				Required: true,
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
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwjsonerrorpage.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwjsonerrorpage.Name = data.Name.ValueString()
	}
	if !data.Overwrite.IsNull() && !data.Overwrite.IsUnknown() {
		appfwjsonerrorpage.Overwrite = data.Overwrite.ValueBool()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		appfwjsonerrorpage.Src = data.Src.ValueString()
	}

	return appfwjsonerrorpage
}

// appfwjsonerrorpageSetAttrFromGet is used by the RESOURCE Create/Read/Update
// paths. NITRO's appfwjsonerrorpage get response returns only name/response/src
// (never `comment` or `overwrite`), and the SDK v2 resource intentionally only
// echoed `name` back into state. Preserve every user-supplied input here and
// only resolve Computed unknowns, so there is no perpetual diff and no
// "inconsistent result after apply".
func appfwjsonerrorpageSetAttrFromGet(ctx context.Context, data *AppfwjsonerrorpageResourceModel, getResponseData map[string]interface{}) *AppfwjsonerrorpageResourceModel {
	tflog.Debug(ctx, "In appfwjsonerrorpageSetAttrFromGet Function")

	// name is the key and is echoed back; use it (also populates state on import).
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}

	// comment / overwrite are never returned by NITRO. Preserve the plan/state
	// value; only resolve a still-unknown Computed value to a known null.
	if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	if data.Overwrite.IsUnknown() {
		data.Overwrite = types.BoolNull()
	}
	// src is Required (always known from config for the resource); preserve it
	// as-is to match SDK v2 (which never read src back).
	if data.Src.IsUnknown() {
		data.Src = types.StringNull()
	}

	// Set ID for the resource (SDK v2: d.SetId(name))
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}

// appfwjsonerrorpageSetAttrFromGetForDatasource is used by the DATASOURCE Read
// path. Unlike the resource, the datasource has no user-supplied plan to
// preserve, so it copies every attribute NITRO returns and sets the ID.
func appfwjsonerrorpageSetAttrFromGetForDatasource(ctx context.Context, data *AppfwjsonerrorpageResourceModel, getResponseData map[string]interface{}) *AppfwjsonerrorpageResourceModel {
	tflog.Debug(ctx, "In appfwjsonerrorpageSetAttrFromGetForDatasource Function")

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
