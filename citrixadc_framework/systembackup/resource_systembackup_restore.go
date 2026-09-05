package systembackup

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	sdkid "github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystembackupRestoreResource{}
var _ resource.ResourceWithConfigure = (*SystembackupRestoreResource)(nil)

func NewSystembackupRestoreResource() resource.Resource {
	return &SystembackupRestoreResource{}
}

// SystembackupRestoreResource defines the resource implementation.
//
// This resource models the NITRO systembackup `?action=restore` action. restore
// is a one-shot side-effect action: it restores the appliance from a previously
// created backup file (*.tgz). NITRO exposes no restore-state GET endpoint and
// there is no inverse API, so Read/Update/Delete are no-ops (mirrors SDK v2,
// which set Read and Delete to schema.Noop and defined no Update). Both
// arguments (filename, skipbackup) are ForceNew in SDK v2 and therefore carry a
// RequiresReplace plan modifier here.
type SystembackupRestoreResource struct {
	client *service.NitroClient
}

// SystembackupRestoreResourceModel describes the resource data model.
type SystembackupRestoreResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Filename   types.String `tfsdk:"filename"`
	Skipbackup types.Bool   `tfsdk:"skipbackup"`
}

func (r *SystembackupRestoreResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systembackup_restore"
}

func (r *SystembackupRestoreResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystembackupRestoreResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systembackup_restore resource.",
			},
			// filename is Required + ForceNew in SDK v2. restore is an action,
			// Read is a no-op, so this attribute must NOT be Computed.
			"filename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the backup file(*.tgz) to be restored.",
			},
			// skipbackup is Optional + ForceNew in SDK v2.
			"skipbackup": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Use this option to skip taking backup during restore operation.",
			},
		},
	}
}

func (r *SystembackupRestoreResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystembackupRestoreResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Restoring systembackup (action-only resource)")

	// Mirror SDK v2 id scheme: PrefixedUniqueId(filename + "-").
	systembackupName := sdkid.PrefixedUniqueId(data.Filename.ValueString() + "-")

	payload := systembackup_restoreGetThePayloadFromthePlan(ctx, &data)

	// restore is a POST ?action=restore action (Pattern 1). There is no add
	// endpoint; the verb casing is lower-case per the NITRO URL.
	err := r.client.ActOnResource(service.Systembackup.Type(), payload, "restore")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to restore systembackup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Restored systembackup")

	data.Id = types.StringValue(systembackupName)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystembackupRestoreResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// restore is a one-shot action. NITRO has no GET endpoint that reports
	// restore-state, so Read is a pure preserve-state no-op (SDK v2 schema.Noop).
	var data SystembackupRestoreResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for systembackup_restore; restore has no stable GET-backed object")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystembackupRestoreResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for restore; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state SystembackupRestoreResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for systembackup_restore; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystembackupRestoreResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// restore is a one-shot side-effect action. There is no inverse NITRO API.
	// Delete simply removes the resource from Terraform state (SDK v2 schema.Noop).
	tflog.Debug(ctx, "Delete is a no-op for systembackup_restore; NITRO has no inverse of the restore action")
}

// systembackup_restoreGetThePayloadFromthePlan builds the body for the restore
// action. It includes ONLY the arguments the restore action accepts: filename
// (required) and skipbackup (optional). All read-only fields on the NITRO struct
// (size, creationtime, version, ...) are excluded.
func systembackup_restoreGetThePayloadFromthePlan(ctx context.Context, data *SystembackupRestoreResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In systembackup_restoreGetThePayloadFromthePlan Function")

	systembackup := map[string]interface{}{}
	if !data.Filename.IsNull() && !data.Filename.IsUnknown() {
		systembackup["filename"] = data.Filename.ValueString()
	}
	if !data.Skipbackup.IsNull() && !data.Skipbackup.IsUnknown() {
		systembackup["skipbackup"] = data.Skipbackup.ValueBool()
	}

	return systembackup
}
