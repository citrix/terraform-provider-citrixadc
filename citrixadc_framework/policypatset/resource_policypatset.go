package policypatset

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
var _ resource.Resource = &PolicypatsetResource{}
var _ resource.ResourceWithConfigure = (*PolicypatsetResource)(nil)
var _ resource.ResourceWithImportState = (*PolicypatsetResource)(nil)

func NewPolicypatsetResource() resource.Resource {
	return &PolicypatsetResource{}
}

// PolicypatsetResource defines the resource implementation.
type PolicypatsetResource struct {
	client *service.NitroClient
}

func (r *PolicypatsetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *PolicypatsetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policypatset"
}

func (r *PolicypatsetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PolicypatsetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PolicypatsetResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating policypatset resource")

	policypatset := policypatsetGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource (NITRO add: POST /policypatset)
	policypatsetName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Policypatset.Type(), policypatsetName, &policypatset)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create policypatset, got error: %s", err))
		return
	}

	// Set ID for the resource (plain value = name) before reading state back
	data.Id = types.StringValue(policypatsetName)

	tflog.Trace(ctx, "Created policypatset resource")

	// Read the updated state back
	if !r.readPolicypatsetFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "policypatset not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicypatsetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data PolicypatsetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading policypatset resource")

	found := r.readPolicypatsetFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *PolicypatsetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state PolicypatsetResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config (to detect attributes removed from config -> unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating policypatset resource")

	// Only "dynamic" is NITRO-updatable (all other attributes are RequiresReplace).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Dynamic.Equal(state.Dynamic) {
		tflog.Debug(ctx, "dynamic has changed for policypatset, starting update")
		if config.Dynamic.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "dynamic")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		policypatset := policypatsetGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// NITRO update: PUT /policypatset  {name, dynamic}
		policypatsetName := data.Id.ValueString()
		policypatset.Name = policypatsetName
		_, err := r.client.UpdateResource(service.Policypatset.Type(), policypatsetName, &policypatset)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update policypatset, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated policypatset resource")
	} else {
		tflog.Debug(ctx, "No changes detected for policypatset resource, skipping update")
	}

	// Unset attributes removed from config so the appliance reverts them to
	// their NITRO defaults. Done after the update so any default value the
	// update payload carried for a removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Policypatset.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset policypatset attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readPolicypatsetFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "policypatset not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicypatsetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data PolicypatsetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting policypatset resource")

	// Named resource - delete using DeleteResource (NITRO delete: DELETE /policypatset/<name>)
	policypatsetName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Policypatset.Type(), policypatsetName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete policypatset, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted policypatset resource")
}

// Helper function to read policypatset data from API.
// Returns false when the resource no longer exists on the ADC.
func (r *PolicypatsetResource) readPolicypatsetFromApi(ctx context.Context, data *PolicypatsetResourceModel, diags *diag.Diagnostics) bool {
	// Single unique attribute - ID is the plain value (name)
	policypatsetName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Policypatset.Type(), policypatsetName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read policypatset, got error: %s", err))
		return false
	}

	policypatsetSetAttrFromGet(ctx, data, getResponseData)

	return true
}
