package filesystemencryption

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

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &FilesystemencryptionResource{}
var _ resource.ResourceWithConfigure = (*FilesystemencryptionResource)(nil)
var _ resource.ResourceWithImportState = (*FilesystemencryptionResource)(nil)
var _ resource.ResourceWithValidateConfig = (*FilesystemencryptionResource)(nil)

func NewFilesystemencryptionResource() resource.Resource {
	return &FilesystemencryptionResource{}
}

// FilesystemencryptionResource defines the resource implementation.
type FilesystemencryptionResource struct {
	client *service.NitroClient
}

func (r *FilesystemencryptionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *FilesystemencryptionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_filesystemencryption"
}

func (r *FilesystemencryptionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// ValidateConfig enforces the CLI-mandatory passphrase secret: at least one of
// passphrase / passphrase_wo must be set (Pattern 17).
func (r *FilesystemencryptionResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data FilesystemencryptionResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Passphrase.IsNull() && data.PassphraseWo.IsNull() {
		resp.Diagnostics.AddError(
			"Missing required attribute",
			"At least one of \"passphrase\" or \"passphrase_wo\" must be set for filesystemencryption.",
		)
	}
}

func (r *FilesystemencryptionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config FilesystemencryptionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating filesystemencryption resource (action=enable)")
	// Get payload from plan (regular attributes)
	filesystemencryption := filesystemencryptionGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	filesystemencryptionGetThePayloadFromtheConfig(ctx, &config, &filesystemencryption)

	// Action-only resource: fire ?action=enable (POST)
	err := r.client.ActOnResource(service.Filesystemencryption.Type(), &filesystemencryption, "enable")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to enable filesystemencryption, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Enabled filesystemencryption resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("filesystemencryption-config")

	// Read the read-only state back
	r.readFilesystemencryptionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FilesystemencryptionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data FilesystemencryptionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading filesystemencryption resource")

	r.readFilesystemencryptionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op: all action inputs are RequiresReplace, so any change re-runs
// enable via a destroy/create cycle. NITRO exposes no set/update endpoint.
func (r *FilesystemencryptionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state FilesystemencryptionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for filesystemencryption; all attributes are RequiresReplace")

	r.readFilesystemencryptionFromApi(ctx, &data, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FilesystemencryptionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FilesystemencryptionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting filesystemencryption resource (action=disable)")

	// disable requires the same three args as enable, from state.
	payload := filesystemencryptionGetThePayloadFromthePlan(ctx, &data)
	// passphrase is write-only: on delete it is only available if it was persisted
	// in state (legacy passphrase attribute). The _wo path is not in state.
	err := r.client.ActOnResource(service.Filesystemencryption.Type(), &payload, "disable")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to disable filesystemencryption, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Disabled filesystemencryption resource")
}

// Helper function to read filesystemencryption read-only state from the nameless singleton GET.
func (r *FilesystemencryptionResource) readFilesystemencryptionFromApi(ctx context.Context, data *FilesystemencryptionResourceModel, diags *diag.Diagnostics) {

	getResponseData, err := r.client.FindResource(service.Filesystemencryption.Type(), "")
	if err != nil {
		// Nameless singleton GET may return nothing meaningful; preserve state rather
		// than fail hard (Pattern 13-ish, action-only resource).
		tflog.Debug(ctx, fmt.Sprintf("filesystemencryption GET returned no data, preserving state: %s", err))
		return
	}

	filesystemencryptionSetAttrFromGet(ctx, data, getResponseData)
}
