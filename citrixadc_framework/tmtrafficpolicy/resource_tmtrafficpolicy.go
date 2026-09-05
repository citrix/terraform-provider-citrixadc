package tmtrafficpolicy

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
var _ resource.Resource = &TmtrafficpolicyResource{}
var _ resource.ResourceWithConfigure = (*TmtrafficpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*TmtrafficpolicyResource)(nil)

func NewTmtrafficpolicyResource() resource.Resource {
	return &TmtrafficpolicyResource{}
}

// TmtrafficpolicyResource defines the resource implementation.
type TmtrafficpolicyResource struct {
	client *service.NitroClient
}

func (r *TmtrafficpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TmtrafficpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tmtrafficpolicy"
}

func (r *TmtrafficpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TmtrafficpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TmtrafficpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating tmtrafficpolicy resource")

	tmtrafficpolicy := tmtrafficpolicyGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	tmtrafficpolicyName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Tmtrafficpolicy.Type(), tmtrafficpolicyName, &tmtrafficpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create tmtrafficpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created tmtrafficpolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(tmtrafficpolicyName)

	// Read the updated state back
	if !r.readTmtrafficpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "tmtrafficpolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmtrafficpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TmtrafficpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading tmtrafficpolicy resource")

	found := r.readTmtrafficpolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *TmtrafficpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state TmtrafficpolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating tmtrafficpolicy resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for tmtrafficpolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for tmtrafficpolicy")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		tmtrafficpolicy := tmtrafficpolicyGetThePayloadFromtheConfig(ctx, &data)
		// Named resource - use UpdateResource
		tmtrafficpolicyName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Tmtrafficpolicy.Type(), tmtrafficpolicyName, &tmtrafficpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update tmtrafficpolicy, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated tmtrafficpolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for tmtrafficpolicy resource, skipping update")
	}

	// Read the updated state back
	if !r.readTmtrafficpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "tmtrafficpolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmtrafficpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TmtrafficpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting tmtrafficpolicy resource")
	// Named resource - delete using DeleteResource
	tmtrafficpolicyName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Tmtrafficpolicy.Type(), tmtrafficpolicyName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete tmtrafficpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted tmtrafficpolicy resource")
}

// Helper function to read tmtrafficpolicy data from API
func (r *TmtrafficpolicyResource) readTmtrafficpolicyFromApi(ctx context.Context, data *TmtrafficpolicyResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	tmtrafficpolicyName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Tmtrafficpolicy.Type(), tmtrafficpolicyName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read tmtrafficpolicy, got error: %s", err))
		return false
	}

	tmtrafficpolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
