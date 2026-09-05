package policydataset

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
var _ resource.Resource = &PolicydatasetResource{}
var _ resource.ResourceWithConfigure = (*PolicydatasetResource)(nil)
var _ resource.ResourceWithImportState = (*PolicydatasetResource)(nil)

func NewPolicydatasetResource() resource.Resource {
	return &PolicydatasetResource{}
}

// PolicydatasetResource defines the resource implementation.
type PolicydatasetResource struct {
	client *service.NitroClient
}

func (r *PolicydatasetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *PolicydatasetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policydataset"
}

func (r *PolicydatasetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PolicydatasetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PolicydatasetResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating policydataset resource")

	policydataset := policydatasetGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource (NITRO add is HTTP POST)
	policydatasetName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Policydataset.Type(), policydatasetName, &policydataset)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create policydataset, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created policydataset resource")

	// Set ID for the resource before reading state (plain name value, matching SDK v2)
	data.Id = types.StringValue(policydatasetName)

	// Read the updated state back
	if !r.readPolicydatasetFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "policydataset not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicydatasetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data PolicydatasetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading policydataset resource")

	found := r.readPolicydatasetFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *PolicydatasetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state PolicydatasetResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (name is RequiresReplace, so it never changes here)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating policydataset resource")

	// Only "dynamic" is updateable in SDK v2; every other attribute is
	// RequiresReplace and never reaches Update.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Dynamic.Equal(state.Dynamic) {
		tflog.Debug(ctx, "dynamic has changed for policydataset, starting update")
		if config.Dynamic.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "dynamic")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		policydataset := policydatasetGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Named resource - use UpdateResource (NITRO update endpoint)
		policydatasetName := data.Id.ValueString()
		_, err := r.client.UpdateResource(service.Policydataset.Type(), policydatasetName, &policydataset)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update policydataset, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated policydataset resource")
	} else {
		tflog.Debug(ctx, "No changes detected for policydataset resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Policydataset.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset policydataset attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readPolicydatasetFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "policydataset not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicydatasetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data PolicydatasetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting policydataset resource")

	// Named resource - delete using DeleteResource (NITRO DELETE /policydataset/{name})
	policydatasetName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Policydataset.Type(), policydatasetName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete policydataset, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted policydataset resource")
}

// Helper function to read policydataset data from API.
// Returns false (without error) when the resource no longer exists on the ADC.
func (r *PolicydatasetResource) readPolicydatasetFromApi(ctx context.Context, data *PolicydatasetResourceModel, diags *diag.Diagnostics) bool {
	// Single unique attribute - ID is the plain name value
	policydatasetName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Policydataset.Type(), policydatasetName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read policydataset, got error: %s", err))
		return false
	}

	policydatasetSetAttrFromGet(ctx, data, getResponseData)

	return true
}
