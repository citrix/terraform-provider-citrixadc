package nschannelparam

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NschannelparamResource{}
var _ resource.ResourceWithConfigure = (*NschannelparamResource)(nil)
var _ resource.ResourceWithImportState = (*NschannelparamResource)(nil)

func NewNschannelparamResource() resource.Resource {
	return &NschannelparamResource{}
}

// NschannelparamResource defines the resource implementation.
type NschannelparamResource struct {
	client *service.NitroClient
}

func (r *NschannelparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NschannelparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nschannelparam"
}

func (r *NschannelparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NschannelparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NschannelparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nschannelparam resource")
	nschannelparam := nschannelparamGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Parameter singleton - NITRO has no "add" verb; create is the "set" verb (PUT).
	_, err := r.client.UpdateResource(service.Nschannelparam.Type(), "", &nschannelparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nschannelparam, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nschannelparam resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("nschannelparam-config")

	// Read the updated state back
	r.readNschannelparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NschannelparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NschannelparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nschannelparam resource")

	r.readNschannelparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NschannelparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NschannelparamResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nschannelparam resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Vfautorecover.Equal(state.Vfautorecover) {
		tflog.Debug(ctx, fmt.Sprintf("vfautorecover has changed for nschannelparam"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		nschannelparam := nschannelparamGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Parameter singleton - update uses the same "set" verb (PUT).
		_, err := r.client.UpdateResource(service.Nschannelparam.Type(), "", &nschannelparam)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nschannelparam, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nschannelparam resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nschannelparam resource, skipping update")
	}

	// Read the updated state back
	r.readNschannelparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NschannelparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NschannelparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nschannelparam resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed nschannelparam from Terraform state")
}

// Helper function to read nschannelparam data from API
func (r *NschannelparamResource) readNschannelparamFromApi(ctx context.Context, data *NschannelparamResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Nschannelparam.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nschannelparam, got error: %s", err))
		return
	}

	nschannelparamSetAttrFromGet(ctx, data, getResponseData)

}
