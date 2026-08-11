package rewriteparam

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
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RewriteparamResource{}
var _ resource.ResourceWithConfigure = (*RewriteparamResource)(nil)
var _ resource.ResourceWithImportState = (*RewriteparamResource)(nil)

func NewRewriteparamResource() resource.Resource {
	return &RewriteparamResource{}
}

// RewriteparamResource defines the resource implementation.
type RewriteparamResource struct {
	client *service.NitroClient
}

func (r *RewriteparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RewriteparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rewriteparam"
}

func (r *RewriteparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RewriteparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RewriteparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rewriteparam resource")

	rewriteparam := rewriteparamGetThePayloadFromthePlan(ctx, &data)

	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Rewriteparam.Type(), &rewriteparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create rewriteparam, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("rewriteparam-config")

	tflog.Trace(ctx, "Created rewriteparam resource")

	// Read the updated state back
	r.readRewriteparamFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RewriteparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RewriteparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading rewriteparam resource")

	r.readRewriteparamFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RewriteparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state RewriteparamResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read Terraform config to detect attributes removed from configuration
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating rewriteparam resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Timeout.Equal(state.Timeout) {
		tflog.Debug(ctx, "timeout has changed for rewriteparam")
		if config.Timeout.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "timeout")
		} else {
			hasChange = true
		}
	}
	if !data.Undefaction.Equal(state.Undefaction) {
		tflog.Debug(ctx, "undefaction has changed for rewriteparam")
		if config.Undefaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "undefaction")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		rewriteparam := rewriteparamGetThePayloadFromthePlan(ctx, &data)

		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Rewriteparam.Type(), &rewriteparam)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update rewriteparam, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated rewriteparam resource")
	} else {
		tflog.Debug(ctx, "No changes detected for rewriteparam resource, skipping update")
	}

	// Issue a single batched unset for attributes removed from configuration.
	// Update-then-unset ordering ensures any default carried in the update
	// payload is superseded by the unset. rewriteparam is a singleton, so the
	// unset carries no identity fields.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Rewriteparam.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset rewriteparam attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readRewriteparamFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RewriteparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RewriteparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting rewriteparam resource")

	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed rewriteparam from Terraform state")
}

// Helper function to read rewriteparam data from API
func (r *RewriteparamResource) readRewriteparamFromApi(ctx context.Context, data *RewriteparamResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Rewriteparam.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read rewriteparam, got error: %s", err))
		return
	}

	rewriteparamSetAttrFromGet(ctx, data, getResponseData)
}
