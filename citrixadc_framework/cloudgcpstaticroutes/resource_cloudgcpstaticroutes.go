package cloudgcpstaticroutes

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CloudgcpstaticroutesResource{}
var _ resource.ResourceWithConfigure = (*CloudgcpstaticroutesResource)(nil)
var _ resource.ResourceWithImportState = (*CloudgcpstaticroutesResource)(nil)

func NewCloudgcpstaticroutesResource() resource.Resource {
	return &CloudgcpstaticroutesResource{}
}

// CloudgcpstaticroutesResource defines the resource implementation.
type CloudgcpstaticroutesResource struct {
	client *service.NitroClient
}

func (r *CloudgcpstaticroutesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CloudgcpstaticroutesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloudgcpstaticroutes"
}

func (r *CloudgcpstaticroutesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CloudgcpstaticroutesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CloudgcpstaticroutesResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cloudgcpstaticroutes resource")
	cloudgcpstaticroutes := cloudgcpstaticroutesGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Unnamed singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Cloudgcpstaticroutes.Type(), &cloudgcpstaticroutes)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cloudgcpstaticroutes, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created cloudgcpstaticroutes resource")

	// Read the updated state back (also sets the ID)
	if !r.readCloudgcpstaticroutesFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "cloudgcpstaticroutes not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudgcpstaticroutesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CloudgcpstaticroutesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cloudgcpstaticroutes resource")

	found := r.readCloudgcpstaticroutesFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *CloudgcpstaticroutesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state CloudgcpstaticroutesResourceModel

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

	tflog.Debug(ctx, "Updating cloudgcpstaticroutes resource")

	// Determine which attributes changed. Attributes removed from config are
	// unset (reverted to their NITRO default); others are pushed via update.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Status.Equal(state.Status) {
		tflog.Debug(ctx, "status has changed for cloudgcpstaticroutes")
		if config.Status.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "status")
		} else {
			hasChange = true
		}
	}
	if !data.Project.Equal(state.Project) {
		tflog.Debug(ctx, "project has changed for cloudgcpstaticroutes")
		if config.Project.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "project")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		cloudgcpstaticroutes := cloudgcpstaticroutesGetThePayloadFromthePlan(ctx, &data)

		// Make API call
		// Unnamed singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Cloudgcpstaticroutes.Type(), &cloudgcpstaticroutes)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cloudgcpstaticroutes, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated cloudgcpstaticroutes resource")
	} else {
		tflog.Debug(ctx, "No updatable changes detected for cloudgcpstaticroutes resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Singleton resource -> empty id payload.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Cloudgcpstaticroutes.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset cloudgcpstaticroutes attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readCloudgcpstaticroutesFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "cloudgcpstaticroutes not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudgcpstaticroutesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CloudgcpstaticroutesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cloudgcpstaticroutes resource")

	// cloudgcpstaticroutes is a global configuration singleton and does not
	// support a DELETE operation. We simply remove it from Terraform state.
	tflog.Trace(ctx, "Removed cloudgcpstaticroutes from Terraform state")
}

// Helper function to read cloudgcpstaticroutes data from API
func (r *CloudgcpstaticroutesResource) readCloudgcpstaticroutesFromApi(ctx context.Context, data *CloudgcpstaticroutesResourceModel, diags *diag.Diagnostics) bool {
	getResponseData, err := r.client.FindResource(service.Cloudgcpstaticroutes.Type(), "")
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cloudgcpstaticroutes, got error: %s", err))
		return false
	}

	cloudgcpstaticroutesSetAttrFromGet(ctx, data, getResponseData)

	return true
}
