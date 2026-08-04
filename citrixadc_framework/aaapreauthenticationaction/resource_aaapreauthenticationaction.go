package aaapreauthenticationaction

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
var _ resource.Resource = &AaapreauthenticationactionResource{}
var _ resource.ResourceWithConfigure = (*AaapreauthenticationactionResource)(nil)
var _ resource.ResourceWithImportState = (*AaapreauthenticationactionResource)(nil)

func NewAaapreauthenticationactionResource() resource.Resource {
	return &AaapreauthenticationactionResource{}
}

// AaapreauthenticationactionResource defines the resource implementation.
type AaapreauthenticationactionResource struct {
	client *service.NitroClient
}

func (r *AaapreauthenticationactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaapreauthenticationactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaapreauthenticationaction"
}

func (r *AaapreauthenticationactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaapreauthenticationactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaapreauthenticationactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaapreauthenticationaction resource")
	// Get payload from plan
	aaapreauthenticationaction := aaapreauthenticationactionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Aaapreauthenticationaction.Type(), name_value, &aaapreauthenticationaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaapreauthenticationaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created aaapreauthenticationaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAaapreauthenticationactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "aaapreauthenticationaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaapreauthenticationactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaapreauthenticationactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaapreauthenticationaction resource")

	found := r.readAaapreauthenticationactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AaapreauthenticationactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AaapreauthenticationactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating aaapreauthenticationaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Defaultepagroup.Equal(state.Defaultepagroup) {
		tflog.Debug(ctx, "defaultepagroup has changed for aaapreauthenticationaction")
		hasChange = true
	}
	if !data.Deletefiles.Equal(state.Deletefiles) {
		tflog.Debug(ctx, "deletefiles has changed for aaapreauthenticationaction")
		hasChange = true
	}
	if !data.Killprocess.Equal(state.Killprocess) {
		tflog.Debug(ctx, "killprocess has changed for aaapreauthenticationaction")
		hasChange = true
	}
	if !data.Preauthenticationaction.Equal(state.Preauthenticationaction) {
		tflog.Debug(ctx, "preauthenticationaction has changed for aaapreauthenticationaction")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model, restricted to NITRO-updatable fields
		aaapreauthenticationaction := aaapreauthenticationactionGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Aaapreauthenticationaction.Type(), name_value, &aaapreauthenticationaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaapreauthenticationaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated aaapreauthenticationaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for aaapreauthenticationaction resource, skipping update")
	}

	// Read the updated state back
	if !r.readAaapreauthenticationactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "aaapreauthenticationaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaapreauthenticationactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaapreauthenticationactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaapreauthenticationaction resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Aaapreauthenticationaction.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete aaapreauthenticationaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted aaapreauthenticationaction resource")
}

// Helper function to read aaapreauthenticationaction data from API
func (r *AaapreauthenticationactionResource) readAaapreauthenticationactionFromApi(ctx context.Context, data *AaapreauthenticationactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Aaapreauthenticationaction.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaapreauthenticationaction, got error: %s", err))
		return false
	}

	aaapreauthenticationactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
