package systemrestorepoint

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SystemrestorepointResourceModel describes the resource data model.
type SystemrestorepointResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Filename types.String `tfsdk:"filename"`
}

func (r *SystemrestorepointResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemrestorepoint resource.",
			},
			"filename": schema.StringAttribute{
				// CLI + NITRO mandatory (tfdata wrongly had is_required:false) -> Required (Pattern 8).
				// Not Computed: it is a pure user input, the resource key.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the restore point.",
			},
		},
	}
}

func systemrestorepointGetThePayloadFromthePlan(ctx context.Context, data *SystemrestorepointResourceModel) system.Systemrestorepoint {
	tflog.Debug(ctx, "In systemrestorepointGetThePayloadFromthePlan Function")

	// Create API request body from the model.
	// Read-only properties (backupfilename, techsuprtname, creationtime, version,
	// createdby, ipaddress, _nextgenapiresource, __count) are NOT in the model and
	// are never sent in the create payload (Pattern 15).
	systemrestorepoint := system.Systemrestorepoint{}
	if !data.Filename.IsNull() && !data.Filename.IsUnknown() {
		systemrestorepoint.Filename = data.Filename.ValueString()
	}

	return systemrestorepoint
}

// systemrestorepointSetAttrFromGet is used by the resource Read path. It preserves
// the resource ID (set once in Create, Pattern 6) and only refreshes the filename
// when present in the GET response.
func systemrestorepointSetAttrFromGet(ctx context.Context, data *SystemrestorepointResourceModel, getResponseData map[string]interface{}) *SystemrestorepointResourceModel {
	tflog.Debug(ctx, "In systemrestorepointSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["filename"]; ok && val != nil {
		data.Filename = types.StringValue(val.(string))
	}

	// ID is set once in Create (Pattern 6); do not recompute it here.

	return data
}

// systemrestorepointSetAttrFromGetForDatasource faithfully copies the GET response
// into the model and sets the ID, since the datasource has no Create step (Pattern 7).
func systemrestorepointSetAttrFromGetForDatasource(ctx context.Context, data *SystemrestorepointResourceModel, getResponseData map[string]interface{}) *SystemrestorepointResourceModel {
	tflog.Debug(ctx, "In systemrestorepointSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["filename"]; ok && val != nil {
		data.Filename = types.StringValue(val.(string))
	} else {
		data.Filename = types.StringNull()
	}

	data.Id = types.StringValue(data.Filename.ValueString())

	return data
}
