package rewritepolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	sdkid "github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RewritepolicyResource{}
var _ resource.ResourceWithConfigure = (*RewritepolicyResource)(nil)
var _ resource.ResourceWithImportState = (*RewritepolicyResource)(nil)

func NewRewritepolicyResource() resource.Resource {
	return &RewritepolicyResource{}
}

// RewritepolicyResource defines the resource implementation.
type RewritepolicyResource struct {
	client *service.NitroClient
}

func (r *RewritepolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RewritepolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rewritepolicy"
}

func (r *RewritepolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RewritepolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RewritepolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rewritepolicy resource")

	// Backward-compatible with the SDK v2 resource: name is optional. When the user
	// does not supply a name, generate a unique one.
	rewritepolicyName := data.Name.ValueString()
	if data.Name.IsNull() || data.Name.IsUnknown() || rewritepolicyName == "" {
		rewritepolicyName = sdkid.PrefixedUniqueId("tf-rewritepolicy-")
		data.Name = types.StringValue(rewritepolicyName)
	}

	// Create API request body from the model
	rewritepolicy := rewritepolicyGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource (NITRO add is HTTP POST)
	_, err := r.client.AddResource(service.Rewritepolicy.Type(), rewritepolicyName, &rewritepolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create rewritepolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created rewritepolicy resource")

	// Set ID for the resource before applying bindings / reading state
	data.Id = types.StringValue(rewritepolicyName)

	// Apply the convenience-block bindings (global / lbvserver / csvserver).
	r.applyBindingsOnCreate(ctx, rewritepolicyName, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read the updated state back
	if !r.readRewritepolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "rewritepolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RewritepolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RewritepolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading rewritepolicy resource")

	found := r.readRewritepolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *RewritepolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state RewritepolicyResourceModel

	// Read Terraform prior state to preserve ID and diff the bindings.
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (candidates to unset).
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id
	rewritepolicyName := state.Id.ValueString()

	tflog.Debug(ctx, "Updating rewritepolicy resource")

	// Check for changes in the base (scalar) attributes.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for rewritepolicy")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for rewritepolicy")
		hasChange = true
	}
	if !data.Logaction.Equal(state.Logaction) {
		tflog.Debug(ctx, "logaction has changed for rewritepolicy")
		if config.Logaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "logaction")
		} else {
			hasChange = true
		}
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for rewritepolicy")
		hasChange = true
	}
	if !data.Undefaction.Equal(state.Undefaction) {
		tflog.Debug(ctx, "undefaction has changed for rewritepolicy")
		if config.Undefaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "undefaction")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		rewritepolicy := rewritepolicyGetThePayloadFromthePlan(ctx, &data)
		rewritepolicy.Name = rewritepolicyName
		_, err := r.client.UpdateResource(service.Rewritepolicy.Type(), rewritepolicyName, &rewritepolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update rewritepolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated rewritepolicy resource")
	} else {
		tflog.Debug(ctx, "No base attribute changes detected for rewritepolicy resource, skipping update")
	}

	// Unset attributes removed from config so the appliance reverts them to
	// their NITRO defaults. Done after the update so any value the update
	// payload carried for a removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": rewritepolicyName,
	}
	if err := utils.ExecuteUnset(r.client, service.Rewritepolicy.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset rewritepolicy attributes, got error: %s", err))
		return
	}

	// Reconcile the convenience-block bindings against prior state.
	r.applyBindingsOnUpdate(ctx, rewritepolicyName, &data, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read the updated state back
	if !r.readRewritepolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "rewritepolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RewritepolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RewritepolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting rewritepolicy resource")

	rewritepolicyName := data.Id.ValueString()

	// Delete all bindings prior to deleting the rewrite policy.
	r.deleteAllBindings(ctx, rewritepolicyName, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Named resource - delete using DeleteResource
	err := r.client.DeleteResource(service.Rewritepolicy.Type(), rewritepolicyName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete rewritepolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted rewritepolicy resource")
}

// Helper function to read rewritepolicy data from API. Returns false when the
// resource no longer exists on the appliance.
func (r *RewritepolicyResource) readRewritepolicyFromApi(ctx context.Context, data *RewritepolicyResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain value (the name).
	rewritepolicyName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Rewritepolicy.Type(), rewritepolicyName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read rewritepolicy, got error: %s", err))
		return false
	}

	rewritepolicySetAttrFromGet(ctx, data, getResponseData)

	// Refresh the managed convenience-block bindings.
	r.readBindings(ctx, rewritepolicyName, data)

	return true
}
