package nsicapprofile

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
var _ resource.Resource = &NsicapprofileResource{}
var _ resource.ResourceWithConfigure = (*NsicapprofileResource)(nil)
var _ resource.ResourceWithImportState = (*NsicapprofileResource)(nil)

func NewNsicapprofileResource() resource.Resource {
	return &NsicapprofileResource{}
}

// NsicapprofileResource defines the resource implementation.
type NsicapprofileResource struct {
	client *service.NitroClient
}

func (r *NsicapprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsicapprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsicapprofile"
}

func (r *NsicapprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsicapprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsicapprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsicapprofile resource")

	// Create API request body from the model
	nsicapprofile := nsicapprofileGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	nsicapprofileName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Nsicapprofile.Type(), nsicapprofileName, &nsicapprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsicapprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nsicapprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(nsicapprofileName)

	// Read the updated state back
	if !r.readNsicapprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsicapprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsicapprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsicapprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsicapprofile resource")

	found := r.readNsicapprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NsicapprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NsicapprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nsicapprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Allow204.Equal(state.Allow204) {
		tflog.Debug(ctx, "allow204 has changed for nsicapprofile")
		if config.Allow204.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "allow204")
		} else {
			hasChange = true
		}
	}
	if !data.Connectionkeepalive.Equal(state.Connectionkeepalive) {
		tflog.Debug(ctx, "connectionkeepalive has changed for nsicapprofile")
		if config.Connectionkeepalive.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "connectionkeepalive")
		} else {
			hasChange = true
		}
	}
	if !data.Hostheader.Equal(state.Hostheader) {
		tflog.Debug(ctx, "hostheader has changed for nsicapprofile")
		hasChange = true
	}
	if !data.Inserthttprequest.Equal(state.Inserthttprequest) {
		tflog.Debug(ctx, "inserthttprequest has changed for nsicapprofile")
		hasChange = true
	}
	if !data.Inserticapheaders.Equal(state.Inserticapheaders) {
		tflog.Debug(ctx, "inserticapheaders has changed for nsicapprofile")
		hasChange = true
	}
	if !data.Logaction.Equal(state.Logaction) {
		tflog.Debug(ctx, "logaction has changed for nsicapprofile")
		hasChange = true
	}
	if !data.Mode.Equal(state.Mode) {
		tflog.Debug(ctx, "mode has changed for nsicapprofile")
		hasChange = true
	}
	if !data.Preview.Equal(state.Preview) {
		tflog.Debug(ctx, "preview has changed for nsicapprofile")
		if config.Preview.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "preview")
		} else {
			hasChange = true
		}
	}
	if !data.Previewlength.Equal(state.Previewlength) {
		tflog.Debug(ctx, "previewlength has changed for nsicapprofile")
		if config.Previewlength.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "previewlength")
		} else {
			hasChange = true
		}
	}
	if !data.Queryparams.Equal(state.Queryparams) {
		tflog.Debug(ctx, "queryparams has changed for nsicapprofile")
		hasChange = true
	}
	if !data.Reqtimeout.Equal(state.Reqtimeout) {
		tflog.Debug(ctx, "reqtimeout has changed for nsicapprofile")
		if config.Reqtimeout.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "reqtimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Reqtimeoutaction.Equal(state.Reqtimeoutaction) {
		tflog.Debug(ctx, "reqtimeoutaction has changed for nsicapprofile")
		if config.Reqtimeoutaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "reqtimeoutaction")
		} else {
			hasChange = true
		}
	}
	if !data.Uri.Equal(state.Uri) {
		tflog.Debug(ctx, "uri has changed for nsicapprofile")
		hasChange = true
	}
	if !data.Useragent.Equal(state.Useragent) {
		tflog.Debug(ctx, "useragent has changed for nsicapprofile")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		nsicapprofile := nsicapprofileGetThePayloadFromtheConfig(ctx, &data)
		// Named resource - use UpdateResource
		nsicapprofileName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Nsicapprofile.Type(), nsicapprofileName, &nsicapprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsicapprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nsicapprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nsicapprofile resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Nsicapprofile.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset nsicapprofile attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readNsicapprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsicapprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsicapprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsicapprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsicapprofile resource")
	// Named resource - delete using DeleteResource
	nsicapprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Nsicapprofile.Type(), nsicapprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nsicapprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nsicapprofile resource")
}

// Helper function to read nsicapprofile data from API
func (r *NsicapprofileResource) readNsicapprofileFromApi(ctx context.Context, data *NsicapprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	nsicapprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nsicapprofile.Type(), nsicapprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsicapprofile, got error: %s", err))
		return false
	}

	nsicapprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
