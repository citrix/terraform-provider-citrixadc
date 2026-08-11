package policystringmap

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
var _ resource.Resource = &PolicystringmapResource{}
var _ resource.ResourceWithConfigure = (*PolicystringmapResource)(nil)
var _ resource.ResourceWithImportState = (*PolicystringmapResource)(nil)

func NewPolicystringmapResource() resource.Resource {
	return &PolicystringmapResource{}
}

// PolicystringmapResource defines the resource implementation.
type PolicystringmapResource struct {
	client *service.NitroClient
}

func (r *PolicystringmapResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *PolicystringmapResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policystringmap"
}

func (r *PolicystringmapResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PolicystringmapResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PolicystringmapResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating policystringmap resource")

	// Create API request body from the plan
	policystringmap := policystringmapGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	policystringmapName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Policystringmap.Type(), policystringmapName, &policystringmap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create policystringmap, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created policystringmap resource")

	// Set ID for the resource before reading state (single unique attribute -> plain value)
	data.Id = types.StringValue(policystringmapName)

	// Read the updated state back
	if !r.readPolicystringmapFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "policystringmap not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicystringmapResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data PolicystringmapResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading policystringmap resource")

	found := r.readPolicystringmapFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *PolicystringmapResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state PolicystringmapResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (candidates to unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating policystringmap resource")

	// Check if there are any changes in updateable attributes.
	// Mirrors SDK v2 updatePolicystringmapFunc which reacted to changes in
	// both comment and name (name has no ForceNew in SDK v2).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for policystringmap")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Name.Equal(state.Name) {
		tflog.Debug(ctx, "name has changed for policystringmap")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the plan
		policystringmap := policystringmapGetThePayloadFromtheConfig(ctx, &data)
		// Named resource - use UpdateResource
		policystringmapName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Policystringmap.Type(), policystringmapName, &policystringmap)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update policystringmap, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated policystringmap resource")
	} else {
		tflog.Debug(ctx, "No changes detected for policystringmap resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Policystringmap.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset policystringmap attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readPolicystringmapFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "policystringmap not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicystringmapResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data PolicystringmapResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting policystringmap resource")

	// Named resource - delete using DeleteResource keyed by the live ID
	policystringmapName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Policystringmap.Type(), policystringmapName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete policystringmap, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted policystringmap resource")
}

// Helper function to read policystringmap data from API
func (r *PolicystringmapResource) readPolicystringmapFromApi(ctx context.Context, data *PolicystringmapResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	policystringmapName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Policystringmap.Type(), policystringmapName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read policystringmap, got error: %s", err))
		return false
	}

	policystringmapSetAttrFromGet(ctx, data, getResponseData)

	return true
}
