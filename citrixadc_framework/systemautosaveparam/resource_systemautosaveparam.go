package systemautosaveparam

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
var _ resource.Resource = &SystemautosaveparamResource{}
var _ resource.ResourceWithConfigure = (*SystemautosaveparamResource)(nil)
var _ resource.ResourceWithImportState = (*SystemautosaveparamResource)(nil)

func NewSystemautosaveparamResource() resource.Resource {
	return &SystemautosaveparamResource{}
}

// SystemautosaveparamResource defines the resource implementation.
type SystemautosaveparamResource struct {
	client *service.NitroClient
}

func (r *SystemautosaveparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemautosaveparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemautosaveparam"
}

func (r *SystemautosaveparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemautosaveparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemautosaveparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemautosaveparam resource")

	// Create API request body from the model
	systemautosaveparam := systemautosaveparamGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Unnamed singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Systemautosaveparam.Type(), &systemautosaveparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemautosaveparam, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created systemautosaveparam resource")

	// Read the updated state back (also sets the ID)
	if !r.readSystemautosaveparamFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "systemautosaveparam not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemautosaveparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemautosaveparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemautosaveparam resource")

	found := r.readSystemautosaveparamFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SystemautosaveparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SystemautosaveparamResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating systemautosaveparam resource")

	// Determine which attributes changed. All writable attributes (status,
	// periodicsave, periodicsavefrequency) support the NITRO unset operation, so
	// an attribute removed from config is unset (reverted to its appliance
	// default) rather than pushed via update.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Status.Equal(state.Status) {
		tflog.Debug(ctx, "status has changed for systemautosaveparam")
		if config.Status.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "status")
		} else {
			hasChange = true
		}
	}
	if !data.Periodicsave.Equal(state.Periodicsave) {
		tflog.Debug(ctx, "periodicsave has changed for systemautosaveparam")
		if config.Periodicsave.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "periodicsave")
		} else {
			hasChange = true
		}
	}
	if !data.Periodicsavefrequency.Equal(state.Periodicsavefrequency) {
		tflog.Debug(ctx, "periodicsavefrequency has changed for systemautosaveparam")
		if config.Periodicsavefrequency.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "periodicsavefrequency")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		systemautosaveparam := systemautosaveparamGetThePayloadFromthePlan(ctx, &data)

		// Make API call
		// Unnamed singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Systemautosaveparam.Type(), &systemautosaveparam)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemautosaveparam, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated systemautosaveparam resource")
	} else {
		tflog.Debug(ctx, "No updatable changes detected for systemautosaveparam resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Singleton resource -> empty id payload.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Systemautosaveparam.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset systemautosaveparam attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readSystemautosaveparamFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "systemautosaveparam not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemautosaveparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemautosaveparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemautosaveparam resource")

	// systemautosaveparam is a global configuration singleton and does not
	// support a DELETE operation. We simply remove it from Terraform state.
	tflog.Trace(ctx, "Removed systemautosaveparam from Terraform state")
}

// Helper function to read systemautosaveparam data from API
func (r *SystemautosaveparamResource) readSystemautosaveparamFromApi(ctx context.Context, data *SystemautosaveparamResourceModel, diags *diag.Diagnostics) bool {
	getResponseData, err := r.client.FindResource(service.Systemautosaveparam.Type(), "")
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemautosaveparam, got error: %s", err))
		return false
	}

	systemautosaveparamSetAttrFromGet(ctx, data, getResponseData)

	return true
}
