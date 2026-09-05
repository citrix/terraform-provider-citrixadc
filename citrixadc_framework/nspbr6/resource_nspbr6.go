package nspbr6

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Nspbr6Resource{}
var _ resource.ResourceWithConfigure = (*Nspbr6Resource)(nil)
var _ resource.ResourceWithImportState = (*Nspbr6Resource)(nil)

func NewNspbr6Resource() resource.Resource {
	return &Nspbr6Resource{}
}

// Nspbr6Resource defines the resource implementation.
type Nspbr6Resource struct {
	client *service.NitroClient
}

func (r *Nspbr6Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Nspbr6Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nspbr6"
}

func (r *Nspbr6Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Nspbr6Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Nspbr6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nspbr6 resource")

	// Named resource - build the full payload (including state, matching SDK v2 Create)
	nspbr6 := nspbr6GetThePayloadFromthePlan(ctx, &data)

	nspbr6Name := data.Name.ValueString()
	_, err := r.client.AddResource(service.Nspbr6.Type(), nspbr6Name, &nspbr6)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nspbr6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nspbr6 resource")

	// ID is the plain name value (matches SDK v2 d.SetId(name))
	data.Id = types.StringValue(nspbr6Name)

	// Read the updated state back
	if !r.readNspbr6FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nspbr6 not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nspbr6Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Nspbr6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nspbr6 resource")

	found := r.readNspbr6FromApi(ctx, &data, &resp.Diagnostics)
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

func (r *Nspbr6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state Nspbr6ResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (candidates for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nspbr6 resource")

	nspbr6Name := data.Name.ValueString()

	// Detect attributes that were removed from config so they can be unset on the
	// appliance (reverting to NITRO defaults). Only spec-unsettable mutable attrs
	// are considered here.
	attributesToUnset := []string{}
	if !data.Msr.Equal(state.Msr) && config.Msr.IsNull() {
		attributesToUnset = append(attributesToUnset, "msr")
	}

	// State (ENABLED/DISABLED) is handled via enable/disable action, not the update body (matches SDK v2)
	stateChanged := !data.State.Equal(state.State) && !data.State.IsUnknown() && !data.State.IsNull()
	if stateChanged {
		if err := r.doNspbr6StateChange(ctx, &data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error enabling/disabling nspbr6 %s: %s", nspbr6Name, err))
			return
		}
	}

	// Any non-state attribute change is pushed via UpdateResource. The payload contains ONLY
	// the changed attributes (matching SDK v2's d.HasChange gating): NITRO rejects a PBR6 SET
	// that carries certain unchanged fields such as "iptunnel" (errorcode 383), so sending the
	// full plan would fail even when those fields did not change.
	if nspbr6HasNonStateChange(&data, &state) {
		nspbr6 := nspbr6GetTheChangedPayloadFromThePlan(ctx, &data, &state)
		_, err := r.client.UpdateResource(service.Nspbr6.Type(), nspbr6Name, &nspbr6)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nspbr6 %s, got error: %s", nspbr6Name, err))
			return
		}
		tflog.Trace(ctx, "Updated nspbr6 resource")
	} else {
		tflog.Debug(ctx, "No non-state changes detected for nspbr6 resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts them
	// to their defaults. Done after the update so any default value the update
	// payload carried for a removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Nspbr6.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset nspbr6 attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readNspbr6FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nspbr6 not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nspbr6Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Nspbr6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nspbr6 resource")

	// Named resource - delete using DeleteResource keyed off the ID (plain name)
	nspbr6Name := data.Id.ValueString()
	err := r.client.DeleteResource(service.Nspbr6.Type(), nspbr6Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nspbr6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nspbr6 resource")
}

// doNspbr6StateChange enables or disables the PBR6 via the NITRO enable/disable action.
func (r *Nspbr6Resource) doNspbr6StateChange(ctx context.Context, data *Nspbr6ResourceModel) error {
	tflog.Debug(ctx, "In doNspbr6StateChange Function")

	// A minimal struct is required; ActOnResource fails on superfluous attributes.
	nspbr6 := ns.Nspbr6{
		Name: data.Name.ValueString(),
	}

	newstate := data.State.ValueString()
	switch newstate {
	case "ENABLED":
		return r.client.ActOnResource(service.Nspbr6.Type(), nspbr6, "enable")
	case "DISABLED":
		return r.client.ActOnResource(service.Nspbr6.Type(), nspbr6, "disable")
	default:
		return fmt.Errorf("%q is not a valid state. Use (\"ENABLED\", \"DISABLED\")", newstate)
	}
}

// Helper function to read nspbr6 data from API. Returns false if the resource no longer exists.
func (r *Nspbr6Resource) readNspbr6FromApi(ctx context.Context, data *Nspbr6ResourceModel, diags *diag.Diagnostics) bool {
	nspbr6Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nspbr6.Type(), nspbr6Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nspbr6, got error: %s", err))
		return false
	}

	nspbr6SetAttrFromGet(ctx, data, getResponseData)

	return true
}
