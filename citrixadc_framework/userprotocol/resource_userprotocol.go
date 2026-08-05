package userprotocol

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
var _ resource.Resource = &UserprotocolResource{}
var _ resource.ResourceWithConfigure = (*UserprotocolResource)(nil)
var _ resource.ResourceWithImportState = (*UserprotocolResource)(nil)

func NewUserprotocolResource() resource.Resource {
	return &UserprotocolResource{}
}

// UserprotocolResource defines the resource implementation.
type UserprotocolResource struct {
	client *service.NitroClient
}

func (r *UserprotocolResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *UserprotocolResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_userprotocol"
}

func (r *UserprotocolResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *UserprotocolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data UserprotocolResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating userprotocol resource")

	userprotocol := userprotocolGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	userprotocolName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Userprotocol.Type(), userprotocolName, &userprotocol)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create userprotocol, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created userprotocol resource")

	// Set ID for the resource before reading state back (single unique attr - plain value)
	data.Id = types.StringValue(userprotocolName)

	// Read the updated state back
	if !r.readUserprotocolFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "userprotocol not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UserprotocolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data UserprotocolResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading userprotocol resource")

	found := r.readUserprotocolFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *UserprotocolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state UserprotocolResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating userprotocol resource")

	// Only `comment` is updatable in SDK v2; name/extension/transport are ForceNew.
	hasChange := false
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for userprotocol")
		hasChange = true
	}

	if hasChange {
		userprotocol := userprotocolGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Matches SDK v2 update path: PUT to the userprotocol resource with name+comment.
		err := r.client.UpdateUnnamedResource(service.Userprotocol.Type(), &userprotocol)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update userprotocol, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated userprotocol resource")
	} else {
		tflog.Debug(ctx, "No changes detected for userprotocol resource, skipping update")
	}

	// Read the updated state back
	if !r.readUserprotocolFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "userprotocol not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UserprotocolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data UserprotocolResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting userprotocol resource")

	// Named resource - delete using DeleteResource (ID is the plain name)
	userprotocolName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Userprotocol.Type(), userprotocolName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete userprotocol, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted userprotocol resource")
}

// Helper function to read userprotocol data from API. Returns false when the
// resource no longer exists on the ADC (so the caller can drop it from state).
func (r *UserprotocolResource) readUserprotocolFromApi(ctx context.Context, data *UserprotocolResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain value (name)
	userprotocolName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Userprotocol.Type(), userprotocolName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read userprotocol, got error: %s", err))
		return false
	}

	userprotocolSetAttrFromGet(ctx, data, getResponseData)

	return true
}
