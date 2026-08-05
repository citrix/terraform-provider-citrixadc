package systembackup

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	sdkresource "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystembackupResource{}
var _ resource.ResourceWithConfigure = (*SystembackupResource)(nil)
var _ resource.ResourceWithImportState = (*SystembackupResource)(nil)

func NewSystembackupResource() resource.Resource {
	return &SystembackupResource{}
}

// SystembackupResource defines the resource implementation.
//
// This mirrors the SDK v2 citrixadc_systembackup resource exactly: a named
// resource keyed on filename. Create issues a NITRO `add` (POST
// /config/systembackup), Read is a no-op (SDK v2 used schema.Noop — there is no
// reliable GET-backed object for the base add), there is no Update endpoint
// (all attributes are ForceNew/RequiresReplace), and Delete removes the backup
// file by filename.
type SystembackupResource struct {
	client *service.NitroClient
}

func (r *SystembackupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystembackupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systembackup"
}

func (r *SystembackupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystembackupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystembackupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systembackup resource")

	// Mirror SDK v2 id scheme: PrefixedUniqueId(filename + "-").
	systembackupName := sdkresource.PrefixedUniqueId(data.Filename.ValueString() + "-")

	systembackup := systembackupGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - SDK v2 created via AddResource (POST /config/systembackup).
	_, err := r.client.AddResource(service.Systembackup.Type(), systembackupName, &systembackup)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systembackup, got error: %s", err))
		return
	}

	// Mirror SDK v2 d.SetId(systembackupName).
	data.Id = types.StringValue(systembackupName)

	tflog.Trace(ctx, "Created systembackup resource")

	// SDK v2 Read was schema.Noop, so Create does not read state back; persist the
	// planned configuration as-is.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystembackupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// SDK v2 Read was schema.Noop. There is no reliable GET-backed object for the
	// base systembackup add, so Read is a pure preserve-state no-op.
	var data SystembackupResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for systembackup (mirrors SDK v2 schema.Noop)")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystembackupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// SDK v2 defined no Update function; every schema attribute is ForceNew
	// (RequiresReplace here), so Terraform never invokes Update for a real change.
	var data, state SystembackupResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for systembackup; all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystembackupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystembackupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systembackup resource")

	// Named resource - SDK v2 deleted the backup file by filename.
	err := r.client.DeleteResource(service.Systembackup.Type(), data.Filename.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete systembackup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted systembackup resource")
}
