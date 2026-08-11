package tmsamlssoprofile

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
var _ resource.Resource = &TmsamlssoprofileResource{}
var _ resource.ResourceWithConfigure = (*TmsamlssoprofileResource)(nil)
var _ resource.ResourceWithImportState = (*TmsamlssoprofileResource)(nil)

func NewTmsamlssoprofileResource() resource.Resource {
	return &TmsamlssoprofileResource{}
}

// TmsamlssoprofileResource defines the resource implementation.
type TmsamlssoprofileResource struct {
	client *service.NitroClient
}

func (r *TmsamlssoprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TmsamlssoprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tmsamlssoprofile"
}

func (r *TmsamlssoprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TmsamlssoprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TmsamlssoprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating tmsamlssoprofile resource")

	// Create API request body from the model
	tmsamlssoprofile := tmsamlssoprofileGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	tmsamlssoprofileName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Tmsamlssoprofile.Type(), tmsamlssoprofileName, &tmsamlssoprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create tmsamlssoprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created tmsamlssoprofile resource")

	// Set ID for the resource before reading state (Case 2: single unique attribute)
	data.Id = types.StringValue(tmsamlssoprofileName)

	// Read the updated state back
	if !r.readTmsamlssoprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "tmsamlssoprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmsamlssoprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TmsamlssoprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading tmsamlssoprofile resource")

	found := r.readTmsamlssoprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *TmsamlssoprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state TmsamlssoprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating tmsamlssoprofile resource")

	// Determine attributes removed from config so they can be unset (reverted to
	// their NITRO defaults) after the update. The update call below runs
	// unconditionally (mirrors SDK v2), so only the unset list is tracked here.
	attributesToUnset := []string{}
	if !data.Digestmethod.Equal(state.Digestmethod) && config.Digestmethod.IsNull() {
		attributesToUnset = append(attributesToUnset, "digestmethod")
	}
	if !data.Nameidformat.Equal(state.Nameidformat) && config.Nameidformat.IsNull() {
		attributesToUnset = append(attributesToUnset, "nameidformat")
	}
	if !data.Signassertion.Equal(state.Signassertion) && config.Signassertion.IsNull() {
		attributesToUnset = append(attributesToUnset, "signassertion")
	}
	if !data.Signaturealg.Equal(state.Signaturealg) && config.Signaturealg.IsNull() {
		attributesToUnset = append(attributesToUnset, "signaturealg")
	}
	if !data.Skewtime.Equal(state.Skewtime) && config.Skewtime.IsNull() {
		attributesToUnset = append(attributesToUnset, "skewtime")
	}

	// Create API request body from the model (Name is included in the payload)
	tmsamlssoprofile := tmsamlssoprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call - mirrors SDK v2 which used UpdateUnnamedResource (name in body)
	err := r.client.UpdateUnnamedResource(service.Tmsamlssoprofile.Type(), &tmsamlssoprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update tmsamlssoprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated tmsamlssoprofile resource")

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Done after the update so any default value the
	// update payload carried for a removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Tmsamlssoprofile.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset tmsamlssoprofile attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readTmsamlssoprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "tmsamlssoprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmsamlssoprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TmsamlssoprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting tmsamlssoprofile resource")

	// Named resource - delete using DeleteResource
	tmsamlssoprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Tmsamlssoprofile.Type(), tmsamlssoprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete tmsamlssoprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted tmsamlssoprofile resource")
}

// Helper function to read tmsamlssoprofile data from API
func (r *TmsamlssoprofileResource) readTmsamlssoprofileFromApi(ctx context.Context, data *TmsamlssoprofileResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain value (the profile name)
	tmsamlssoprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Tmsamlssoprofile.Type(), tmsamlssoprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read tmsamlssoprofile, got error: %s", err))
		return false
	}

	tmsamlssoprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
