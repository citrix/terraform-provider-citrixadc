package systemrestorepoint

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
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

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystemrestorepointResource{}
var _ resource.ResourceWithConfigure = (*SystemrestorepointResource)(nil)
var _ resource.ResourceWithImportState = (*SystemrestorepointResource)(nil)

func NewSystemrestorepointResource() resource.Resource {
	return &SystemrestorepointResource{}
}

// SystemrestorepointResource defines the resource implementation.
type SystemrestorepointResource struct {
	client *service.NitroClient
}

func (r *SystemrestorepointResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemrestorepointResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemrestorepoint"
}

func (r *SystemrestorepointResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemrestorepointResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemrestorepointResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemrestorepoint resource")
	systemrestorepoint := systemrestorepointGetThePayloadFromthePlan(ctx, &data)

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

func (r *SystemrestorepointResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemrestorepointResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemrestorepoint resource")

	r.readSystemrestorepointFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemrestorepointResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemrestorepointResourceModel

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

func (r *SystemrestorepointResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemrestorepointResourceModel

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
func (r *SystemrestorepointResource) readSystemrestorepointFromApi(ctx context.Context, data *SystemrestorepointResourceModel, diags *diag.Diagnostics) {

	// get-byname endpoint exists: GET /systemrestorepoint/<filename>
	filename_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Systemrestorepoint.Type(), filename_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemrestorepoint, got error: %s", err))
		return
	}

	systemrestorepointSetAttrFromGet(ctx, data, getResponseData)

}
