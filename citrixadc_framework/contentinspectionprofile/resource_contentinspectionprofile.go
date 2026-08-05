package contentinspectionprofile

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
var _ resource.Resource = &ContentinspectionprofileResource{}
var _ resource.ResourceWithConfigure = (*ContentinspectionprofileResource)(nil)
var _ resource.ResourceWithImportState = (*ContentinspectionprofileResource)(nil)

func NewContentinspectionprofileResource() resource.Resource {
	return &ContentinspectionprofileResource{}
}

// ContentinspectionprofileResource defines the resource implementation.
type ContentinspectionprofileResource struct {
	client *service.NitroClient
}

func (r *ContentinspectionprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ContentinspectionprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_contentinspectionprofile"
}

func (r *ContentinspectionprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ContentinspectionprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ContentinspectionprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating contentinspectionprofile resource")

	contentinspectionprofile := contentinspectionprofileGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Contentinspectionprofile.Type(), name_value, &contentinspectionprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create contentinspectionprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created contentinspectionprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readContentinspectionprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "contentinspectionprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ContentinspectionprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading contentinspectionprofile resource")

	found := r.readContentinspectionprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *ContentinspectionprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ContentinspectionprofileResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating contentinspectionprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Egressinterface.Equal(state.Egressinterface) {
		tflog.Debug(ctx, "egressinterface has changed for contentinspectionprofile")
		hasChange = true
	}
	if !data.Egressvlan.Equal(state.Egressvlan) {
		tflog.Debug(ctx, "egressvlan has changed for contentinspectionprofile")
		hasChange = true
	}
	if !data.Ingressinterface.Equal(state.Ingressinterface) {
		tflog.Debug(ctx, "ingressinterface has changed for contentinspectionprofile")
		hasChange = true
	}
	if !data.Ingressvlan.Equal(state.Ingressvlan) {
		tflog.Debug(ctx, "ingressvlan has changed for contentinspectionprofile")
		hasChange = true
	}
	if !data.Iptunnel.Equal(state.Iptunnel) {
		tflog.Debug(ctx, "iptunnel has changed for contentinspectionprofile")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model, restricted to updatable fields
		contentinspectionprofile := contentinspectionprofileGetTheUpdatablePayloadFromThePlan(ctx, &data)

		// The update URL carries no name (PUT /contentinspectionprofile); the name is in
		// the payload, so use UpdateUnnamedResource (matches SDK v2 behavior).
		err := r.client.UpdateUnnamedResource(service.Contentinspectionprofile.Type(), &contentinspectionprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update contentinspectionprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated contentinspectionprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for contentinspectionprofile resource, skipping update")
	}

	// Read the updated state back
	if !r.readContentinspectionprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "contentinspectionprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ContentinspectionprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting contentinspectionprofile resource")

	// Named resource - delete using DeleteResource keyed on the profile name (ID)
	name_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Contentinspectionprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete contentinspectionprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted contentinspectionprofile resource")
}

// Helper function to read contentinspectionprofile data from API
func (r *ContentinspectionprofileResource) readContentinspectionprofileFromApi(ctx context.Context, data *ContentinspectionprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain profile name
	name_value := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Contentinspectionprofile.Type(), name_value)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read contentinspectionprofile, got error: %s", err))
		return false
	}

	contentinspectionprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
