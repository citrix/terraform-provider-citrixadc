package authenticationsamlidpprofile

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
var _ resource.Resource = &AuthenticationsamlidpprofileResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationsamlidpprofileResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationsamlidpprofileResource)(nil)

func NewAuthenticationsamlidpprofileResource() resource.Resource {
	return &AuthenticationsamlidpprofileResource{}
}

// AuthenticationsamlidpprofileResource defines the resource implementation.
type AuthenticationsamlidpprofileResource struct {
	client *service.NitroClient
}

func (r *AuthenticationsamlidpprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationsamlidpprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationsamlidpprofile"
}

func (r *AuthenticationsamlidpprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationsamlidpprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationsamlidpprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationsamlidpprofile resource")

	// Build the payload from the plan
	authenticationsamlidpprofile := authenticationsamlidpprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	authenticationsamlidpprofileName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationsamlidpprofile.Type(), authenticationsamlidpprofileName, &authenticationsamlidpprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationsamlidpprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationsamlidpprofile resource")

	// Set ID for the resource before reading state (id == name)
	data.Id = types.StringValue(authenticationsamlidpprofileName)

	// Read the updated state back
	if !r.readAuthenticationsamlidpprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationsamlidpprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationsamlidpprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationsamlidpprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationsamlidpprofile resource")

	found := r.readAuthenticationsamlidpprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationsamlidpprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationsamlidpprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (name is RequiresReplace, so it never changes here)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationsamlidpprofile resource")

	// Build the payload from the plan and update the resource
	authenticationsamlidpprofile := authenticationsamlidpprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use UpdateResource
	authenticationsamlidpprofileName := data.Name.ValueString()
	_, err := r.client.UpdateResource(service.Authenticationsamlidpprofile.Type(), authenticationsamlidpprofileName, &authenticationsamlidpprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationsamlidpprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated authenticationsamlidpprofile resource")

	// Read the updated state back
	if !r.readAuthenticationsamlidpprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationsamlidpprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationsamlidpprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationsamlidpprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationsamlidpprofile resource")

	// Named resource - delete using DeleteResource keyed on the ID (name)
	authenticationsamlidpprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Authenticationsamlidpprofile.Type(), authenticationsamlidpprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationsamlidpprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationsamlidpprofile resource")
}

// Helper function to read authenticationsamlidpprofile data from API
func (r *AuthenticationsamlidpprofileResource) readAuthenticationsamlidpprofileFromApi(ctx context.Context, data *AuthenticationsamlidpprofileResourceModel, diags *diag.Diagnostics) bool {

	// Find with single ID attribute - ID is the plain value (the profile name)
	authenticationsamlidpprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Authenticationsamlidpprofile.Type(), authenticationsamlidpprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationsamlidpprofile, got error: %s", err))
		return false
	}

	authenticationsamlidpprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
