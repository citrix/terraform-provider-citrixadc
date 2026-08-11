package autoscaleaction

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
var _ resource.Resource = &AutoscaleactionResource{}
var _ resource.ResourceWithConfigure = (*AutoscaleactionResource)(nil)
var _ resource.ResourceWithImportState = (*AutoscaleactionResource)(nil)

func NewAutoscaleactionResource() resource.Resource {
	return &AutoscaleactionResource{}
}

// AutoscaleactionResource defines the resource implementation.
type AutoscaleactionResource struct {
	client *service.NitroClient
}

func (r *AutoscaleactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AutoscaleactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_autoscaleaction"
}

func (r *AutoscaleactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AutoscaleactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AutoscaleactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating autoscaleaction resource")

	autoscaleaction := autoscaleactionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource (NITRO add is POST /config/autoscaleaction)
	autoscaleactionName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Autoscaleaction.Type(), autoscaleactionName, &autoscaleaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create autoscaleaction, got error: %s", err))
		return
	}

	// Generate ID for this resource (single unique attribute -> plain value)
	data.Id = types.StringValue(autoscaleactionName)

	tflog.Trace(ctx, "Created autoscaleaction resource")

	// Read the updated state back
	if !r.readAutoscaleactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "autoscaleaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AutoscaleactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AutoscaleactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading autoscaleaction resource")

	found := r.readAutoscaleactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AutoscaleactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AutoscaleactionResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating autoscaleaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Parameters.Equal(state.Parameters) {
		tflog.Debug(ctx, "parameters has changed for autoscaleaction")
		hasChange = true
	}
	if !data.Profilename.Equal(state.Profilename) {
		tflog.Debug(ctx, "profilename has changed for autoscaleaction")
		hasChange = true
	}
	if !data.Quiettime.Equal(state.Quiettime) {
		tflog.Debug(ctx, "quiettime has changed for autoscaleaction")
		if config.Quiettime.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "quiettime")
		} else {
			hasChange = true
		}
	}
	if !data.Vmdestroygraceperiod.Equal(state.Vmdestroygraceperiod) {
		tflog.Debug(ctx, "vmdestroygraceperiod has changed for autoscaleaction")
		if config.Vmdestroygraceperiod.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "vmdestroygraceperiod")
		} else {
			hasChange = true
		}
	}
	if !data.Vserver.Equal(state.Vserver) {
		tflog.Debug(ctx, "vserver has changed for autoscaleaction")
		hasChange = true
	}

	if hasChange {
		// NITRO update is a PUT to /config/autoscaleaction (unnamed URL, name in body)
		autoscaleaction := autoscaleactionGetTheUpdatablePayloadFromThePlan(ctx, &data)
		err := r.client.UpdateUnnamedResource(service.Autoscaleaction.Type(), &autoscaleaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update autoscaleaction, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated autoscaleaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for autoscaleaction resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their NITRO defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Autoscaleaction.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset autoscaleaction attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readAutoscaleactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "autoscaleaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AutoscaleactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AutoscaleactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting autoscaleaction resource")
	// Named resource - delete using DeleteResource keyed on the name (the ID)
	autoscaleactionName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Autoscaleaction.Type(), autoscaleactionName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete autoscaleaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted autoscaleaction resource")
}

// Helper function to read autoscaleaction data from API.
// Returns false when the resource no longer exists on the ADC.
func (r *AutoscaleactionResource) readAutoscaleactionFromApi(ctx context.Context, data *AutoscaleactionResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain name value
	autoscaleactionName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Autoscaleaction.Type(), autoscaleactionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read autoscaleaction, got error: %s", err))
		return false
	}

	autoscaleactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
