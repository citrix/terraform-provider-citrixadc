package cloudtrafficroutes

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
var _ resource.Resource = &CloudtrafficroutesResource{}
var _ resource.ResourceWithConfigure = (*CloudtrafficroutesResource)(nil)
var _ resource.ResourceWithImportState = (*CloudtrafficroutesResource)(nil)

func NewCloudtrafficroutesResource() resource.Resource {
	return &CloudtrafficroutesResource{}
}

// CloudtrafficroutesResource defines the resource implementation.
type CloudtrafficroutesResource struct {
	client *service.NitroClient
}

func (r *CloudtrafficroutesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CloudtrafficroutesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloudtrafficroutes"
}

func (r *CloudtrafficroutesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CloudtrafficroutesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CloudtrafficroutesResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cloudtrafficroutes resource")

	cloudtrafficroutes := cloudtrafficroutesGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource. The name is the resource name.
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Cloudtrafficroutes.Type(), name_value, &cloudtrafficroutes)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cloudtrafficroutes, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created cloudtrafficroutes resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the created state back
	if !r.readCloudtrafficroutesFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "cloudtrafficroutes not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudtrafficroutesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CloudtrafficroutesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cloudtrafficroutes resource")

	found := r.readCloudtrafficroutesFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *CloudtrafficroutesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state CloudtrafficroutesResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating cloudtrafficroutes resource")

	// Check if there are any changes in updateable attributes.
	// All writable attributes (targetvpcnetwork, destrange, nexthopip, ownernode)
	// support the NITRO unset operation, so removal from config reverts them.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Targetvpcnetwork.Equal(state.Targetvpcnetwork) {
		tflog.Debug(ctx, "targetvpcnetwork has changed for cloudtrafficroutes")
		if config.Targetvpcnetwork.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "targetvpcnetwork")
		} else {
			hasChange = true
		}
	}
	if !data.Destrange.Equal(state.Destrange) {
		tflog.Debug(ctx, "destrange has changed for cloudtrafficroutes")
		if config.Destrange.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "destrange")
		} else {
			hasChange = true
		}
	}
	if !data.Nexthopip.Equal(state.Nexthopip) {
		tflog.Debug(ctx, "nexthopip has changed for cloudtrafficroutes")
		if config.Nexthopip.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "nexthopip")
		} else {
			hasChange = true
		}
	}
	if !data.Ownernode.Equal(state.Ownernode) {
		tflog.Debug(ctx, "ownernode has changed for cloudtrafficroutes")
		if config.Ownernode.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "ownernode")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		cloudtrafficroutes := cloudtrafficroutesGetThePayloadFromthePlan(ctx, &data)
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Cloudtrafficroutes.Type(), name_value, &cloudtrafficroutes)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cloudtrafficroutes, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated cloudtrafficroutes resource")
	} else {
		tflog.Debug(ctx, "No changes detected for cloudtrafficroutes resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts them
	// to their defaults. Update-then-unset ordering ensures any default carried in
	// the update payload is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Cloudtrafficroutes.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset cloudtrafficroutes attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readCloudtrafficroutesFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "cloudtrafficroutes not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudtrafficroutesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CloudtrafficroutesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cloudtrafficroutes resource")
	// Named resource - delete using DeleteResource
	name_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Cloudtrafficroutes.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete cloudtrafficroutes, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted cloudtrafficroutes resource")
}

// Helper function to read cloudtrafficroutes data from API. Returns false when
// the resource no longer exists on the ADC.
func (r *CloudtrafficroutesResource) readCloudtrafficroutesFromApi(ctx context.Context, data *CloudtrafficroutesResourceModel, diags *diag.Diagnostics) bool {
	// Named resource - find by ID (the name value).
	name_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Cloudtrafficroutes.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cloudtrafficroutes, got error: %s", err))
		return false
	}

	cloudtrafficroutesSetAttrFromGet(ctx, data, getResponseData)

	return true
}
