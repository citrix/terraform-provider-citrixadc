package nscapacity

import (
	"context"
	"fmt"
	"time"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NscapacityResource{}
var _ resource.ResourceWithConfigure = (*NscapacityResource)(nil)
var _ resource.ResourceWithImportState = (*NscapacityResource)(nil)

func NewNscapacityResource() resource.Resource {
	return &NscapacityResource{}
}

// NscapacityResource defines the resource implementation.
type NscapacityResource struct {
	client *service.NitroClient
}

func (r *NscapacityResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NscapacityResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nscapacity"
}

func (r *NscapacityResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NscapacityResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NscapacityResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nscapacity resource")

	// Build the NITRO payload from the plan (Optional+Computed attributes that were
	// not configured are unknown and are skipped by the payload builder).
	nscapacity := nscapacityGetThePayloadFromtheConfig(ctx, &data)

	// Singleton resource - use UpdateUnnamedResource to push the capacity config.
	err := r.client.UpdateUnnamedResource(service.Nscapacity.Type(), &nscapacity)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nscapacity, got error: %s", err))
		return
	}

	// Applying capacity/licensing requires a warm reboot for it to take effect
	// (matches the SDK v2 createNscapacityFunc behaviour).
	if err := r.warmRebootNetScaler(ctx); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error warm rebooting ADC after applying nscapacity, got error: %s", err))
		return
	}

	// Singleton resource - static ID
	data.Id = types.StringValue("nscapacity-config")

	tflog.Trace(ctx, "Created nscapacity resource")

	// Read the updated state back
	r.readNscapacityFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NscapacityResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NscapacityResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nscapacity resource")

	r.readNscapacityFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NscapacityResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NscapacityResourceModel

	// Read Terraform prior state to preserve the ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (singleton)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nscapacity resource")

	// Build the NITRO payload from the plan.
	nscapacity := nscapacityGetThePayloadFromtheConfig(ctx, &data)

	// Singleton resource - use UpdateUnnamedResource. SDK v2 shared its create func
	// for updates, so the update path mirrors create (push config + warm reboot).
	err := r.client.UpdateUnnamedResource(service.Nscapacity.Type(), &nscapacity)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nscapacity, got error: %s", err))
		return
	}

	// Applying capacity/licensing requires a warm reboot for it to take effect.
	if err := r.warmRebootNetScaler(ctx); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error warm rebooting ADC after applying nscapacity, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated nscapacity resource")

	// Read the updated state back
	r.readNscapacityFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NscapacityResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NscapacityResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nscapacity resource")

	// nscapacity is a singleton/action-only configuration; there is no DELETE verb.
	// SDK v2 (deleteNscapacityFunc) resets the licensing knobs via the NITRO "unset"
	// action, flagging the fields that were set in state. Replicate that exactly.
	type nscapacityRemove struct {
		Bandwidth bool `json:"bandwidth,omitempty"`
		Platform  bool `json:"platform,omitempty"`
		Vcpu      bool `json:"vcpu,omitempty"`
	}
	nscapacity := nscapacityRemove{}

	if !data.Bandwidth.IsNull() && data.Bandwidth.ValueInt64() != 0 {
		nscapacity.Bandwidth = true
	}
	if !data.Platform.IsNull() && data.Platform.ValueString() != "" {
		nscapacity.Platform = true
	}
	if !data.Vcpu.IsNull() && data.Vcpu.ValueBool() {
		nscapacity.Vcpu = true
	}

	if err := r.client.ActOnResource(service.Nscapacity.Type(), &nscapacity, "unset"); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete (unset) nscapacity, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nscapacity resource (unset licensing knobs)")
}

// warmRebootNetScaler issues a warm reboot and waits for the appliance to come
// back, mirroring the SDK v2 rebootNetScaler(d, meta, warm=true) helper. Applying
// an nscapacity/licensing change only takes effect after a reboot.
func (r *NscapacityResource) warmRebootNetScaler(ctx context.Context) error {
	tflog.Debug(ctx, "Warm rebooting NetScaler after nscapacity change")
	reboot := ns.Reboot{
		Warm: true,
	}
	if err := r.client.ActOnResource("reboot", &reboot, ""); err != nil {
		return err
	}
	// Wait for the NetScaler to come back after a warm reboot (SDK v2 sleeps 120s).
	time.Sleep(time.Second * 120)
	return nil
}

// Helper function to read nscapacity data from API
func (r *NscapacityResource) readNscapacityFromApi(ctx context.Context, data *NscapacityResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nscapacity.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nscapacity, got error: %s", err))
		return
	}

	nscapacitySetAttrFromGet(ctx, data, getResponseData)

}
