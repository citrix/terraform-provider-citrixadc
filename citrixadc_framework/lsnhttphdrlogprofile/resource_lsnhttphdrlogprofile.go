package lsnhttphdrlogprofile

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
var _ resource.Resource = &LsnhttphdrlogprofileResource{}
var _ resource.ResourceWithConfigure = (*LsnhttphdrlogprofileResource)(nil)
var _ resource.ResourceWithImportState = (*LsnhttphdrlogprofileResource)(nil)

func NewLsnhttphdrlogprofileResource() resource.Resource {
	return &LsnhttphdrlogprofileResource{}
}

// LsnhttphdrlogprofileResource defines the resource implementation.
type LsnhttphdrlogprofileResource struct {
	client *service.NitroClient
}

func (r *LsnhttphdrlogprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnhttphdrlogprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnhttphdrlogprofile"
}

func (r *LsnhttphdrlogprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnhttphdrlogprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnhttphdrlogprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnhttphdrlogprofile resource")

	lsnhttphdrlogprofile := lsnhttphdrlogprofileGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	httphdrlogprofilename_value := data.Httphdrlogprofilename.ValueString()
	_, err := r.client.AddResource(service.Lsnhttphdrlogprofile.Type(), httphdrlogprofilename_value, &lsnhttphdrlogprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnhttphdrlogprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lsnhttphdrlogprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", httphdrlogprofilename_value))

	// Read the updated state back
	if !r.readLsnhttphdrlogprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsnhttphdrlogprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnhttphdrlogprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnhttphdrlogprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnhttphdrlogprofile resource")

	found := r.readLsnhttphdrlogprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LsnhttphdrlogprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state LsnhttphdrlogprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (to unset them)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lsnhttphdrlogprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Loghost.Equal(state.Loghost) {
		tflog.Debug(ctx, "loghost has changed for lsnhttphdrlogprofile")
		if config.Loghost.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "loghost")
		} else {
			hasChange = true
		}
	}
	if !data.Logmethod.Equal(state.Logmethod) {
		tflog.Debug(ctx, "logmethod has changed for lsnhttphdrlogprofile")
		if config.Logmethod.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "logmethod")
		} else {
			hasChange = true
		}
	}
	if !data.Logurl.Equal(state.Logurl) {
		tflog.Debug(ctx, "logurl has changed for lsnhttphdrlogprofile")
		if config.Logurl.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "logurl")
		} else {
			hasChange = true
		}
	}
	if !data.Logversion.Equal(state.Logversion) {
		tflog.Debug(ctx, "logversion has changed for lsnhttphdrlogprofile")
		if config.Logversion.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "logversion")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		lsnhttphdrlogprofile := lsnhttphdrlogprofileGetThePayloadFromthePlan(ctx, &data)
		// Named resource - use UpdateResource
		httphdrlogprofilename_value := data.Httphdrlogprofilename.ValueString()
		_, err := r.client.UpdateResource(service.Lsnhttphdrlogprofile.Type(), httphdrlogprofilename_value, &lsnhttphdrlogprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnhttphdrlogprofile, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated lsnhttphdrlogprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lsnhttphdrlogprofile resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"httphdrlogprofilename": data.Httphdrlogprofilename.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Lsnhttphdrlogprofile.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset lsnhttphdrlogprofile attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readLsnhttphdrlogprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsnhttphdrlogprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnhttphdrlogprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnhttphdrlogprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnhttphdrlogprofile resource")

	// Named resource - delete using DeleteResource (keyed by the live ID)
	httphdrlogprofilename_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Lsnhttphdrlogprofile.Type(), httphdrlogprofilename_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsnhttphdrlogprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsnhttphdrlogprofile resource")
}

// Helper function to read lsnhttphdrlogprofile data from API
func (r *LsnhttphdrlogprofileResource) readLsnhttphdrlogprofileFromApi(ctx context.Context, data *LsnhttphdrlogprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	httphdrlogprofilename_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Lsnhttphdrlogprofile.Type(), httphdrlogprofilename_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnhttphdrlogprofile, got error: %s", err))
		return false
	}

	lsnhttphdrlogprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
