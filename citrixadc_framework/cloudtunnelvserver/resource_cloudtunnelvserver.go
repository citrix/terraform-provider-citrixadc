package cloudtunnelvserver

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CloudtunnelvserverResource{}
var _ resource.ResourceWithConfigure = (*CloudtunnelvserverResource)(nil)
var _ resource.ResourceWithImportState = (*CloudtunnelvserverResource)(nil)

func NewCloudtunnelvserverResource() resource.Resource {
	return &CloudtunnelvserverResource{}
}

// CloudtunnelvserverResource defines the resource implementation.
type CloudtunnelvserverResource struct {
	client *service.NitroClient
}

func (r *CloudtunnelvserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CloudtunnelvserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloudtunnelvserver"
}

func (r *CloudtunnelvserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CloudtunnelvserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CloudtunnelvserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cloudtunnelvserver resource")
	cloudtunnelvserver := cloudtunnelvserverGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Cloudtunnelvserver.Type(), name_value, &cloudtunnelvserver)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cloudtunnelvserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created cloudtunnelvserver resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readCloudtunnelvserverFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudtunnelvserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CloudtunnelvserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cloudtunnelvserver resource")

	r.readCloudtunnelvserverFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// If the object was deleted out-of-band, remove it from state so a
	// subsequent apply re-creates it instead of erroring.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudtunnelvserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CloudtunnelvserverResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating cloudtunnelvserver resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Listenpolicy.Equal(state.Listenpolicy) {
		tflog.Debug(ctx, fmt.Sprintf("listenpolicy has changed for cloudtunnelvserver"))
		hasChange = true
	}
	if !data.Listenpriority.Equal(state.Listenpriority) {
		tflog.Debug(ctx, fmt.Sprintf("listenpriority has changed for cloudtunnelvserver"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		cloudtunnelvserver := cloudtunnelvserverGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Cloudtunnelvserver.Type(), name_value, &cloudtunnelvserver)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cloudtunnelvserver, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated cloudtunnelvserver resource")
	} else {
		tflog.Debug(ctx, "No changes detected for cloudtunnelvserver resource, skipping update")
	}

	// Read the updated state back
	r.readCloudtunnelvserverFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudtunnelvserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CloudtunnelvserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cloudtunnelvserver resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Cloudtunnelvserver.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete cloudtunnelvserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted cloudtunnelvserver resource")
}

// Helper function to read cloudtunnelvserver data from API
func (r *CloudtunnelvserverResource) readCloudtunnelvserverFromApi(ctx context.Context, data *CloudtunnelvserverResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Cloudtunnelvserver.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			// Object deleted out-of-band: signal "gone" so Read removes it from state.
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cloudtunnelvserver, got error: %s", err))
		return
	}

	cloudtunnelvserverSetAttrFromGet(ctx, data, getResponseData)

}
