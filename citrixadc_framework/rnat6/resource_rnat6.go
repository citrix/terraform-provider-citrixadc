package rnat6

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
var _ resource.Resource = &Rnat6Resource{}
var _ resource.ResourceWithConfigure = (*Rnat6Resource)(nil)
var _ resource.ResourceWithImportState = (*Rnat6Resource)(nil)

func NewRnat6Resource() resource.Resource {
	return &Rnat6Resource{}
}

// Rnat6Resource defines the resource implementation.
type Rnat6Resource struct {
	client *service.NitroClient
}

func (r *Rnat6Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Rnat6Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rnat6"
}

func (r *Rnat6Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Rnat6Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Rnat6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rnat6 resource")

	// Build payload from the plan.
	rnat6 := rnat6GetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource with the rule name.
	rnat6Name := data.Name.ValueString()
	_, err := r.client.AddResource(service.Rnat6.Type(), rnat6Name, &rnat6)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create rnat6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created rnat6 resource")

	// Set ID for the resource before reading state (Case 2: single unique attr).
	data.Id = types.StringValue(rnat6Name)

	// Read the updated state back
	if !r.readRnat6FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "rnat6 not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Rnat6Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Rnat6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading rnat6 resource")

	found := r.readRnat6FromApi(ctx, &data, &resp.Diagnostics)
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

func (r *Rnat6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state Rnat6ResourceModel

	// Read Terraform prior state to detect changes and preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (-> unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating rnat6 resource")

	// Build payload restricted to NITRO-updatable fields (matches SDK v2).
	rnat6, hasChange := rnat6GetTheUpdatablePayloadFromThePlan(ctx, &data, &state)

	// srcippersistency: when removed from config, unset it (revert to NITRO
	// default); otherwise include the new value in the update payload.
	attributesToUnset := []string{}
	if !data.Srcippersistency.Equal(state.Srcippersistency) {
		tflog.Debug(ctx, "srcippersistency has changed for rnat6")
		if config.Srcippersistency.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "srcippersistency")
		} else if !data.Srcippersistency.IsUnknown() {
			rnat6.Srcippersistency = data.Srcippersistency.ValueString()
			hasChange = true
		}
	}

	if hasChange {
		err := r.client.UpdateUnnamedResource(service.Rnat6.Type(), &rnat6)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update rnat6, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated rnat6 resource")
	} else {
		tflog.Debug(ctx, "No changes detected for rnat6 resource, skipping update")
	}

	// Unset attributes removed from config so the appliance reverts them to
	// their NITRO defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Rnat6.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset rnat6 attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readRnat6FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "rnat6 not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Rnat6Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Rnat6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting rnat6 resource")
	// Backward-compat: the SDK v2 resource does NOT issue a NITRO delete for
	// rnat6 ("rnat6 does not support DELETE operation"); it simply drops the
	// resource from state. Preserve that behaviour here.
	tflog.Trace(ctx, "Deleted rnat6 resource from state")
}

// Helper function to read rnat6 data from API. Returns false when the resource
// no longer exists on the ADC so callers can remove it from state.
func (r *Rnat6Resource) readRnat6FromApi(ctx context.Context, data *Rnat6ResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain rule name.
	rnat6Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Rnat6.Type(), rnat6Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read rnat6, got error: %s", err))
		return false
	}
	if getResponseData == nil {
		return false
	}

	rnat6SetAttrFromGet(ctx, data, getResponseData)

	return true
}
