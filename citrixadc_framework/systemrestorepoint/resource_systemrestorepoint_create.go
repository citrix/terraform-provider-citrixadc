package systemrestorepoint

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// systemrestorepoint creates a restore point: a snapshot of the appliance
// configuration plus a tech-support bundle. NITRO exposes `create`
// (POST ?action=create), plain DELETE /<filename>, and get/get-byname/count.
// There is NO `add` and NO update verb.
//
// NOTE: the appliance enforces a MAXIMUM of 3 restore points. Creating a 4th
// will fail on the NITRO side until an existing restore point is deleted.
//
// This is a FULL managed resource (lifecycle mode): Create/Read/Update/Delete
// are all preserved and the ID is the real object name (the filename), so
// Delete targets the actual restore point on the appliance.

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystemrestorepointCreateResource{}
var _ resource.ResourceWithConfigure = (*SystemrestorepointCreateResource)(nil)
var _ resource.ResourceWithImportState = (*SystemrestorepointCreateResource)(nil)

func NewSystemrestorepointCreateResource() resource.Resource {
	return &SystemrestorepointCreateResource{}
}

// SystemrestorepointCreateResource defines the resource implementation.
type SystemrestorepointCreateResource struct {
	client *service.NitroClient
}

// SystemrestorepointCreateResourceModel describes the resource data model.
type SystemrestorepointCreateResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Filename types.String `tfsdk:"filename"`
}

func (r *SystemrestorepointCreateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemrestorepointCreateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemrestorepoint_create"
}

func (r *SystemrestorepointCreateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemrestorepointCreateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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

func (r *SystemrestorepointCreateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemrestorepointCreateResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemrestorepoint resource")
	systemrestorepoint := systemrestorepointCreateGetThePayloadFromthePlan(ctx, &data)

	// Action-only create: NITRO has no `add` verb. Create is POST ?action=create
	// (lowercase "create"). This snapshots the config + tech-support bundle.
	err := r.client.ActOnResource(service.Systemrestorepoint.Type(), &systemrestorepoint, "create")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemrestorepoint, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created systemrestorepoint resource")

	// Set ID for the resource before reading state (Pattern 6: ID set once in Create).
	data.Id = types.StringValue(data.Filename.ValueString())

	// Read the created state back
	r.readSystemrestorepointFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemrestorepointCreateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemrestorepointCreateResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemrestorepoint resource")

	r.readSystemrestorepointFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Resource has been deleted out-of-band - remove from state
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemrestorepointCreateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemrestorepointCreateResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op for systemrestorepoint; NITRO has no update verb and
	// every attribute is RequiresReplace, so changes force recreation.
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for systemrestorepoint; all attributes are RequiresReplace")

	// Read the current state back
	r.readSystemrestorepointFromApi(ctx, &data, &resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemrestorepointCreateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemrestorepointCreateResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemrestorepoint resource")
	// Plain DELETE /systemrestorepoint/<filename>, no extra args.
	filename_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Systemrestorepoint.Type(), filename_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete systemrestorepoint, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted systemrestorepoint resource")
}

// Helper function to read systemrestorepoint data from API
func (r *SystemrestorepointCreateResource) readSystemrestorepointFromApi(ctx context.Context, data *SystemrestorepointCreateResourceModel, diags *diag.Diagnostics) {

	// get-byname endpoint exists: GET /systemrestorepoint/<filename>
	filename_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Systemrestorepoint.Type(), filename_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemrestorepoint, got error: %s", err))
		return
	}

	systemrestorepointCreateSetAttrFromGet(ctx, data, getResponseData)

}

func systemrestorepointCreateGetThePayloadFromthePlan(ctx context.Context, data *SystemrestorepointCreateResourceModel) system.Systemrestorepoint {
	tflog.Debug(ctx, "In systemrestorepointCreateGetThePayloadFromthePlan Function")

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

// systemrestorepointCreateSetAttrFromGet is used by the resource Read path. It preserves
// the resource ID (set once in Create, Pattern 6) and only refreshes the filename
// when present in the GET response.
func systemrestorepointCreateSetAttrFromGet(ctx context.Context, data *SystemrestorepointCreateResourceModel, getResponseData map[string]interface{}) *SystemrestorepointCreateResourceModel {
	tflog.Debug(ctx, "In systemrestorepointCreateSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["filename"]; ok && val != nil {
		data.Filename = types.StringValue(val.(string))
	}

	// ID is set once in Create (Pattern 6); do not recompute it here.

	return data
}
