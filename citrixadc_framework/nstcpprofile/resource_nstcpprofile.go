package nstcpprofile

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
var _ resource.Resource = &NstcpprofileResource{}
var _ resource.ResourceWithConfigure = (*NstcpprofileResource)(nil)
var _ resource.ResourceWithImportState = (*NstcpprofileResource)(nil)

func NewNstcpprofileResource() resource.Resource {
	return &NstcpprofileResource{}
}

// NstcpprofileResource defines the resource implementation.
type NstcpprofileResource struct {
	client *service.NitroClient
}

func (r *NstcpprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NstcpprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nstcpprofile"
}

func (r *NstcpprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NstcpprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NstcpprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nstcpprofile resource")

	// Create API request body from the model
	nstcpprofile := nstcpprofileGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	nstcpprofileName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Nstcpprofile.Type(), nstcpprofileName, &nstcpprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nstcpprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nstcpprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(nstcpprofileName)

	// Read the updated state back
	if !r.readNstcpprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nstcpprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstcpprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NstcpprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nstcpprofile resource")

	found := r.readNstcpprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NstcpprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NstcpprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nstcpprofile resource")

	// Build the payload from the plan (name identifies the resource)
	nstcpprofile := nstcpprofileGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use UpdateResource
	nstcpprofileName := data.Name.ValueString()
	_, err := r.client.UpdateResource(service.Nstcpprofile.Type(), nstcpprofileName, &nstcpprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nstcpprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated nstcpprofile resource")

	// Read the updated state back
	if !r.readNstcpprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nstcpprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstcpprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NstcpprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nstcpprofile resource")
	// Named resource - delete using DeleteResource
	nstcpprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Nstcpprofile.Type(), nstcpprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nstcpprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nstcpprofile resource")
}

// Helper function to read nstcpprofile data from API
func (r *NstcpprofileResource) readNstcpprofileFromApi(ctx context.Context, data *NstcpprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	nstcpprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nstcpprofile.Type(), nstcpprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nstcpprofile, got error: %s", err))
		return false
	}

	nstcpprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
