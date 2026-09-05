package cloudroutes

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
var _ resource.Resource = &CloudroutesResource{}
var _ resource.ResourceWithConfigure = (*CloudroutesResource)(nil)
var _ resource.ResourceWithImportState = (*CloudroutesResource)(nil)

func NewCloudroutesResource() resource.Resource {
	return &CloudroutesResource{}
}

// CloudroutesResource defines the resource implementation.
type CloudroutesResource struct {
	client *service.NitroClient
}

func (r *CloudroutesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CloudroutesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloudroutes"
}

func (r *CloudroutesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CloudroutesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CloudroutesResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cloudroutes resource")
	// Get payload from plan (regular attributes)
	cloudroutes := cloudroutesGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Cloudroutes.Type(), name_value, &cloudroutes)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cloudroutes, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created cloudroutes resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readCloudroutesFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "cloudroutes not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudroutesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CloudroutesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cloudroutes resource")

	found := r.readCloudroutesFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *CloudroutesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state CloudroutesResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect config-removal (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating cloudroutes resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	// Collect eligible attributes that were removed from config so they can be unset on the appliance
	attributesToUnset := []string{}
	if !data.Routesvpcnetwork.Equal(state.Routesvpcnetwork) {
		tflog.Debug(ctx, "routesvpcnetwork has changed for cloudroutes")
		hasChange = true
	}
	if !data.Vipsubnet.Equal(state.Vipsubnet) {
		tflog.Debug(ctx, "vipsubnet has changed for cloudroutes")
		if config.Vipsubnet.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "vipsubnet")
		} else {
			hasChange = true
		}
	}
	if !data.Vipvpcnetwork.Equal(state.Vipvpcnetwork) {
		tflog.Debug(ctx, "vipvpcnetwork has changed for cloudroutes")
		if config.Vipvpcnetwork.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "vipvpcnetwork")
		} else {
			hasChange = true
		}
	}
	if !data.Clientipaddress.Equal(state.Clientipaddress) {
		tflog.Debug(ctx, "clientipaddress has changed for cloudroutes")
		if config.Clientipaddress.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "clientipaddress")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		cloudroutes := cloudroutesGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Cloudroutes.Type(), name_value, &cloudroutes)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cloudroutes, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated cloudroutes resource")
	} else {
		tflog.Debug(ctx, "No changes detected for cloudroutes resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts them to their defaults.
	// Update-then-unset ordering ensures any default carried in the update payload is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Cloudroutes.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset cloudroutes attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readCloudroutesFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "cloudroutes not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudroutesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CloudroutesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cloudroutes resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Cloudroutes.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete cloudroutes, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted cloudroutes resource")
}

// Helper function to read cloudroutes data from API
func (r *CloudroutesResource) readCloudroutesFromApi(ctx context.Context, data *CloudroutesResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Cloudroutes.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cloudroutes, got error: %s", err))
		return false
	}

	cloudroutesSetAttrFromGet(ctx, data, getResponseData)

	return true
}
