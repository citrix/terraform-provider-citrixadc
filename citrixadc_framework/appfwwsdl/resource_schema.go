package appfwwsdl

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwwsdlResourceModel describes the resource data model.
type AppfwwsdlResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Name      types.String `tfsdk:"name"`
	Overwrite types.Bool   `tfsdk:"overwrite"`
	Src       types.String `tfsdk:"src"`
}

func (r *AppfwwsdlResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwwsdl resource.",
			},
			// SDK v2: Optional + Computed + ForceNew. NITRO never echoes `comment`
			// back on GET, so a Default keeps the planned value known and avoids an
			// "inconsistent result after apply" when the attribute is unset.
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about the WSDL.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the WSDL file to remove.",
			},
			// SDK v2: Optional + Computed + ForceNew. NITRO never echoes `overwrite`
			// back on GET, so a Default keeps the planned value known and avoids an
			// "inconsistent result after apply" when the attribute is unset.
			"overwrite": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Overwrite any existing WSDL of the same name.",
			},
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL (protocol, host, path, and name) of the WSDL file to be imported is stored.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
		},
	}
}

func appfwwsdlGetThePayloadFromthePlan(ctx context.Context, data *AppfwwsdlResourceModel) appfw.Appfwwsdl {
	tflog.Debug(ctx, "In appfwwsdlGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwwsdl := appfw.Appfwwsdl{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwwsdl.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwwsdl.Name = data.Name.ValueString()
	}
	if !data.Overwrite.IsNull() && !data.Overwrite.IsUnknown() {
		appfwwsdl.Overwrite = data.Overwrite.ValueBool()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		appfwwsdl.Src = data.Src.ValueString()
	}

	return appfwwsdl
}

func appfwwsdlSetAttrFromGet(ctx context.Context, data *AppfwwsdlResourceModel, getResponseData map[string]interface{}) *AppfwwsdlResourceModel {
	tflog.Debug(ctx, "In appfwwsdlSetAttrFromGet Function")

	// NITRO's appfwwsdl `get` response payload only carries `name`, `response`,
	// and `_nextgenapiresource`. The user-supplied Import inputs `comment`,
	// `overwrite`, and `src` are NEVER echoed back. Touching them here would null
	// them on every Read and cause a perpetual diff / "inconsistent result after
	// apply", so only update `name` (the response-side field) from the API and
	// preserve the existing plan/state values for the rest. This mirrors the
	// SDK v2 read, which only did d.Set("name", ...).
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID (matches SDK v2 d.SetId(name))
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
