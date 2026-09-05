package snmpalarm

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
var _ resource.Resource = &SnmpalarmResource{}
var _ resource.ResourceWithConfigure = (*SnmpalarmResource)(nil)
var _ resource.ResourceWithImportState = (*SnmpalarmResource)(nil)

func NewSnmpalarmResource() resource.Resource {
	return &SnmpalarmResource{}
}

// SnmpalarmResource defines the resource implementation.
type SnmpalarmResource struct {
	client *service.NitroClient
}

func (r *SnmpalarmResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SnmpalarmResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmpalarm"
}

func (r *SnmpalarmResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SnmpalarmResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SnmpalarmResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating snmpalarm resource")

	// snmpalarm has no add/delete operation - it is a predefined alarm that is
	// configured via the update (PUT) operation. Mirror the SDK v2 create, which
	// pushed the full configuration (including state) via UpdateUnnamedResource.
	snmpalarm := snmpalarmGetThePayloadFromthePlan(ctx, &data)

	err := r.client.UpdateUnnamedResource(service.Snmpalarm.Type(), &snmpalarm)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create snmpalarm, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created snmpalarm resource")

	// Set ID for the resource before reading state back (single unique attribute).
	data.Id = types.StringValue(data.Trapname.ValueString())

	// Read the updated state back
	if !r.readSnmpalarmFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "snmpalarm not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpalarmResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SnmpalarmResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading snmpalarm resource")

	found := r.readSnmpalarmFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SnmpalarmResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SnmpalarmResourceModel

	// Read Terraform prior state to preserve ID and for change detection
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (to be unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating snmpalarm resource")

	// Detect changes to the non-state updateable attributes (pushed via PUT).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Holdtime.Equal(state.Holdtime) {
		tflog.Debug(ctx, "holdtime has changed for snmpalarm")
		hasChange = true
	}
	if !data.Logging.Equal(state.Logging) {
		tflog.Debug(ctx, "logging has changed for snmpalarm")
		if config.Logging.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "logging")
		} else {
			hasChange = true
		}
	}
	if !data.Normalvalue.Equal(state.Normalvalue) {
		tflog.Debug(ctx, "normalvalue has changed for snmpalarm")
		hasChange = true
	}
	if !data.Severity.Equal(state.Severity) {
		tflog.Debug(ctx, "severity has changed for snmpalarm")
		if config.Severity.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "severity")
		} else {
			hasChange = true
		}
	}
	if !data.Thresholdvalue.Equal(state.Thresholdvalue) {
		tflog.Debug(ctx, "thresholdvalue has changed for snmpalarm")
		hasChange = true
	}
	if !data.Time.Equal(state.Time) {
		tflog.Debug(ctx, "time has changed for snmpalarm")
		hasChange = true
	}

	// State transitions are applied through the enable/disable actions, mirroring
	// the SDK v2 doSnmpalarmStateChange behaviour.
	if !data.State.Equal(state.State) && !data.State.IsNull() && !data.State.IsUnknown() {
		tflog.Debug(ctx, "state has changed for snmpalarm")
		if err := r.doSnmpalarmStateChange(ctx, data.Trapname.ValueString(), data.State.ValueString()); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update snmpalarm state, got error: %s", err))
			return
		}
	}

	if hasChange {
		snmpalarm := snmpalarmGetTheUpdatablePayloadFromThePlan(ctx, &data)
		err := r.client.UpdateUnnamedResource(service.Snmpalarm.Type(), &snmpalarm)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update snmpalarm, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated snmpalarm resource")
	} else {
		tflog.Debug(ctx, "No PUT-updateable changes detected for snmpalarm resource")
	}

	// Unset attributes removed from config so the appliance reverts them to
	// their defaults. Done after the update so any default value the update
	// payload carried for a removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"trapname": data.Trapname.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Snmpalarm.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset snmpalarm attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readSnmpalarmFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "snmpalarm not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpalarmResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SnmpalarmResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting snmpalarm resource")
	// snmpalarm is a predefined resource with no DELETE operation. Mirror the
	// SDK v2 behaviour: removing it from state only (Terraform handles the state
	// removal automatically once Delete returns without error).
	tflog.Trace(ctx, "Deleted snmpalarm resource from state")
}

// doSnmpalarmStateChange enables/disables the alarm via the dedicated NITRO
// actions (mirrors SDK v2 doSnmpalarmStateChange).
func (r *SnmpalarmResource) doSnmpalarmStateChange(ctx context.Context, trapname string, newstate string) error {
	tflog.Debug(ctx, "In doSnmpalarmStateChange")

	// A fresh payload with only the identifier is required; ActOnResource fails
	// if superfluous attributes are supplied.
	payload := map[string]interface{}{
		"trapname": trapname,
	}

	switch newstate {
	case "ENABLED":
		return r.client.ActOnResource(service.Snmpalarm.Type(), payload, "enable")
	case "DISABLED":
		return r.client.ActOnResource(service.Snmpalarm.Type(), payload, "disable")
	default:
		return fmt.Errorf("%q is not a valid state. Use (\"ENABLED\", \"DISABLED\")", newstate)
	}
}

// Helper function to read snmpalarm data from API
func (r *SnmpalarmResource) readSnmpalarmFromApi(ctx context.Context, data *SnmpalarmResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain trapname value.
	trapname := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Snmpalarm.Type(), trapname)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read snmpalarm, got error: %s", err))
		return false
	}

	snmpalarmSetAttrFromGet(ctx, data, getResponseData)

	return true
}
