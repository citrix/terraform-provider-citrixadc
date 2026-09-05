package nsaigwprofile

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
var _ resource.Resource = &NsaigwprofileResource{}
var _ resource.ResourceWithConfigure = (*NsaigwprofileResource)(nil)
var _ resource.ResourceWithImportState = (*NsaigwprofileResource)(nil)

func NewNsaigwprofileResource() resource.Resource {
	return &NsaigwprofileResource{}
}

// NsaigwprofileResource defines the resource implementation.
type NsaigwprofileResource struct {
	client *service.NitroClient
}

func (r *NsaigwprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsaigwprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsaigwprofile"
}

func (r *NsaigwprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsaigwprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config NsaigwprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsaigwprofile resource")
	// Get payload from plan (regular attributes)
	nsaigwprofile := nsaigwprofileGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	nsaigwprofileGetThePayloadFromtheConfig(ctx, &config, &nsaigwprofile)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Nsaigwprofile.Type(), name_value, &nsaigwprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsaigwprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nsaigwprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readNsaigwprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsaigwprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsaigwprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsaigwprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsaigwprofile resource")

	found := r.readNsaigwprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NsaigwprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NsaigwprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nsaigwprofile resource")

	// Only tokenquota and quotarefreshfrequency are mutable in place. endpointtype,
	// profiletype, authtoken and authtoken_wo_version are create-only (NITRO rejects
	// them on `set` with errorcode 278) and are RequiresReplace, so Terraform never
	// invokes Update for a change to them.
	hasChange := false
	// Collect eligible attributes that were removed from config so they can be unset on the appliance
	attributesToUnset := []string{}
	if !data.Tokenquota.Equal(state.Tokenquota) {
		tflog.Debug(ctx, "tokenquota has changed for nsaigwprofile")
		if config.Tokenquota.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "tokenquota")
		} else {
			hasChange = true
		}
	}
	if !data.Quotarefreshfrequency.Equal(state.Quotarefreshfrequency) {
		tflog.Debug(ctx, "quotarefreshfrequency has changed for nsaigwprofile")
		if config.Quotarefreshfrequency.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "quotarefreshfrequency")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Build the UPDATE payload (only the in-place-mutable attributes; create-only
		// attributes are excluded so the `set` is not rejected).
		nsaigwprofile := nsaigwprofileGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Nsaigwprofile.Type(), name_value, &nsaigwprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsaigwprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nsaigwprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nsaigwprofile resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts them to their defaults.
	// Update-then-unset ordering ensures any default carried in the update payload is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Nsaigwprofile.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset nsaigwprofile attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readNsaigwprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsaigwprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsaigwprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsaigwprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsaigwprofile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Nsaigwprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nsaigwprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nsaigwprofile resource")
}

// Helper function to read nsaigwprofile data from API
func (r *NsaigwprofileResource) readNsaigwprofileFromApi(ctx context.Context, data *NsaigwprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Nsaigwprofile.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsaigwprofile, got error: %s", err))
		return false
	}

	nsaigwprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
