package hasecureheartbeats

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &HasecureheartbeatsResource{}
var _ resource.ResourceWithConfigure = (*HasecureheartbeatsResource)(nil)
var _ resource.ResourceWithImportState = (*HasecureheartbeatsResource)(nil)

func NewHasecureheartbeatsResource() resource.Resource {
	return &HasecureheartbeatsResource{}
}

// HasecureheartbeatsResource defines the resource implementation.
type HasecureheartbeatsResource struct {
	client *service.NitroClient
}

func (r *HasecureheartbeatsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *HasecureheartbeatsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hasecureheartbeats"
}

func (r *HasecureheartbeatsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *HasecureheartbeatsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config HasecureheartbeatsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating hasecureheartbeats resource")

	// Create API request body from the model (regular attributes)
	hasecureheartbeats := hasecureheartbeatsGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	hasecureheartbeatsGetThePayloadFromtheConfig(ctx, &config, &hasecureheartbeats)

	// Make API call
	// Unnamed singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Hasecureheartbeats.Type(), &hasecureheartbeats)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create hasecureheartbeats, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created hasecureheartbeats resource")

	// Read the updated state back (also sets the ID)
	if !r.readHasecureheartbeatsFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "hasecureheartbeats not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HasecureheartbeatsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data HasecureheartbeatsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading hasecureheartbeats resource")

	found := r.readHasecureheartbeatsFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HasecureheartbeatsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state HasecureheartbeatsResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config for write-only attribute handling
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating hasecureheartbeats resource")

	// Determine whether any attribute changed. The NITRO singleton rejects an
	// empty PUT, so we only issue the update when something actually changed.
	hasChange := false
	if !data.State.Equal(state.State) {
		tflog.Debug(ctx, "state has changed for hasecureheartbeats")
		hasChange = true
	}
	// Check secret attribute hapsk or its version tracker. Gated on the secret
	// still being configured (plain or write-only) so removing it does not fire
	// an empty singleton PUT.
	if !config.Hapsk.IsNull() || !config.HapskWo.IsNull() {
		if !data.Hapsk.Equal(state.Hapsk) {
			tflog.Debug(ctx, "hapsk has changed for hasecureheartbeats")
			hasChange = true
		} else if !data.HapskWoVersion.Equal(state.HapskWoVersion) {
			tflog.Debug(ctx, "hapsk_wo_version has changed for hasecureheartbeats")
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model (regular attributes)
		hasecureheartbeats := hasecureheartbeatsGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		hasecureheartbeatsGetThePayloadFromtheConfig(ctx, &config, &hasecureheartbeats)

		// Make API call
		// Unnamed singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Hasecureheartbeats.Type(), &hasecureheartbeats)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update hasecureheartbeats, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated hasecureheartbeats resource")
	} else {
		tflog.Debug(ctx, "No updatable changes detected for hasecureheartbeats resource, skipping update")
	}

	// Read the updated state back
	if !r.readHasecureheartbeatsFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "hasecureheartbeats not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HasecureheartbeatsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data HasecureheartbeatsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting hasecureheartbeats resource")

	// hasecureheartbeats is a global configuration singleton and does not support
	// a DELETE operation. We simply remove it from Terraform state.
	tflog.Trace(ctx, "Removed hasecureheartbeats from Terraform state")
}

// Helper function to read hasecureheartbeats data from API
func (r *HasecureheartbeatsResource) readHasecureheartbeatsFromApi(ctx context.Context, data *HasecureheartbeatsResourceModel, diags *diag.Diagnostics) bool {
	getResponseData, err := r.client.FindResource(service.Hasecureheartbeats.Type(), "")
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read hasecureheartbeats, got error: %s", err))
		return false
	}

	hasecureheartbeatsSetAttrFromGet(ctx, data, getResponseData)

	return true
}

// UpgradeState migrates pre-write-only state (GH #1441): it seeds the
// "hapsk_wo_version" tracker attribute to 1 when the stored state has no value
// for it, so the schema Default does not plan a spurious "null -> 1" update
// after upgrading the provider. Paired with the schema Version bump so the
// upgrade path actually runs. See utils.WoVersionUpgradeState.
