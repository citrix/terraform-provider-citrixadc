package nshttpprofile

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
var _ resource.Resource = &NshttpprofileResource{}
var _ resource.ResourceWithConfigure = (*NshttpprofileResource)(nil)
var _ resource.ResourceWithImportState = (*NshttpprofileResource)(nil)

func NewNshttpprofileResource() resource.Resource {
	return &NshttpprofileResource{}
}

// NshttpprofileResource defines the resource implementation.
type NshttpprofileResource struct {
	client *service.NitroClient
}

func (r *NshttpprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NshttpprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nshttpprofile"
}

func (r *NshttpprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NshttpprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NshttpprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nshttpprofile resource")

	// Create API request body from the plan
	nshttpprofile := nshttpprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	nshttpprofileName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Nshttpprofile.Type(), nshttpprofileName, &nshttpprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nshttpprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nshttpprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(nshttpprofileName)

	// Read the updated state back
	if !r.readNshttpprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nshttpprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NshttpprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NshttpprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nshttpprofile resource")

	found := r.readNshttpprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NshttpprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NshttpprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (name is ForceNew, so the ID never changes on update)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nshttpprofile resource")

	// Create API request body from the plan (only known, configured attributes are sent)
	nshttpprofile := nshttpprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use UpdateResource
	nshttpprofileName := data.Id.ValueString()
	_, err := r.client.UpdateResource(service.Nshttpprofile.Type(), nshttpprofileName, &nshttpprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nshttpprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated nshttpprofile resource")

	// Read the updated state back
	if !r.readNshttpprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nshttpprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NshttpprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NshttpprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nshttpprofile resource")

	// Named resource - delete using DeleteResource
	nshttpprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Nshttpprofile.Type(), nshttpprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nshttpprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nshttpprofile resource")
}

// Helper function to read nshttpprofile data from API.
// Returns false when the resource no longer exists on the ADC.
func (r *NshttpprofileResource) readNshttpprofileFromApi(ctx context.Context, data *NshttpprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain name value
	nshttpprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nshttpprofile.Type(), nshttpprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nshttpprofile, got error: %s", err))
		return false
	}

	nshttpprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
