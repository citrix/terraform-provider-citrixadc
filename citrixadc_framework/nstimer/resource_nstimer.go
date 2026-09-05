package nstimer

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NstimerResource{}
var _ resource.ResourceWithConfigure = (*NstimerResource)(nil)
var _ resource.ResourceWithImportState = (*NstimerResource)(nil)

func NewNstimerResource() resource.Resource {
	return &NstimerResource{}
}

// NstimerResource defines the resource implementation.
type NstimerResource struct {
	client *service.NitroClient
}

func (r *NstimerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NstimerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nstimer"
}

func (r *NstimerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NstimerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NstimerResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nstimer resource")

	nstimer := nstimerGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource (NITRO add is POST)
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Nstimer.Type(), name_value, &nstimer)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nstimer, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nstimer resource")

	// Set ID for the resource (single unique attribute -> plain value) before reading back
	data.Id = types.StringValue(data.Name.ValueString())

	// Read the updated state back
	r.readNstimerFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstimerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NstimerResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nstimer resource")

	r.readNstimerFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Resource was deleted out-of-band
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstimerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NstimerResourceModel

	// Read Terraform prior state to preserve ID / detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (tracks the current live name)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nstimer resource")

	// Rename support: nstimer exposes a NITRO ?action=rename. A newname change drives
	// an in-place rename instead of a destroy/recreate. Every other attribute except
	// the RequiresReplace key (name) reaches Update as a normal change.
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, tracked by the ID (== name at
		// create, == the prior newname after a rename) - NOT data.Name, which stays
		// pinned to the originally configured value.
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming nstimer from %q to %q", oldName, newName))

		renamePayload := ns.Nstimer{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Nstimer.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename nstimer, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so the update /
		// read below (and all future reads) address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// Regular in-place update of the updatable attributes (comment, interval, unit).
	hasChange := false
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for nstimer")
		hasChange = true
	}
	if !data.Interval.Equal(state.Interval) {
		tflog.Debug(ctx, "interval has changed for nstimer")
		hasChange = true
	}
	if !data.Unit.Equal(state.Unit) {
		tflog.Debug(ctx, "unit has changed for nstimer")
		hasChange = true
	}

	if hasChange {
		nstimer := nstimerGetThePayloadFromthePlan(ctx, &data)
		// Target the CURRENT LIVE name (== data.Id, which reflects any rename above).
		nstimer.Name = data.Id.ValueString()
		_, err := r.client.UpdateResource(service.Nstimer.Type(), data.Id.ValueString(), &nstimer)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nstimer, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated nstimer resource")
	} else {
		tflog.Debug(ctx, "No non-rename changes detected for nstimer resource, skipping update")
	}

	// Read the current state back. Preserve the plan's name/newname so a rename does
	// not let the GET response (which now reports the live/new name) clobber the
	// user-facing configured values and cause an inconsistent-result / perpetual diff.
	planName := data.Name
	planNewname := data.Newname
	r.readNstimerFromApi(ctx, &data, &resp.Diagnostics)
	data.Name = planName
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstimerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NstimerResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nstimer resource")

	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE name
	// (== name at create, == newname after a rename), so delete by data.Id, NOT
	// data.Name (which stays at the originally configured value and would target a
	// non-existent name after a rename, leaving the object dangling).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Nstimer.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nstimer, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nstimer resource")
}

// Helper function to read nstimer data from API
func (r *NstimerResource) readNstimerFromApi(ctx context.Context, data *NstimerResourceModel, diags *diag.Diagnostics) {
	// Case 2: Find with single ID attribute - ID is the plain value (current live name)
	name_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nstimer.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nstimer, got error: %s", err))
		return
	}

	nstimerSetAttrFromGet(ctx, data, getResponseData)
}
