package smppuser

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
var _ resource.Resource = &SmppuserResource{}
var _ resource.ResourceWithConfigure = (*SmppuserResource)(nil)
var _ resource.ResourceWithImportState = (*SmppuserResource)(nil)

func NewSmppuserResource() resource.Resource {
	return &SmppuserResource{}
}

// SmppuserResource defines the resource implementation.
type SmppuserResource struct {
	client *service.NitroClient
}

func (r *SmppuserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SmppuserResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_smppuser"
}

func (r *SmppuserResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SmppuserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config SmppuserResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating smppuser resource")
	// Get payload from plan (regular attributes)
	smppuser := smppuserGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	smppuserGetThePayloadFromtheConfig(ctx, &config, &smppuser)

	// Make API call
	// Named resource - use AddResource
	username_value := data.Username.ValueString()
	_, err := r.client.AddResource(service.Smppuser.Type(), username_value, &smppuser)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create smppuser, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created smppuser resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Username.ValueString()))

	// Read the updated state back
	r.readSmppuserFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SmppuserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SmppuserResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading smppuser resource")

	r.readSmppuserFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SmppuserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SmppuserResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating smppuser resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	// Check secret attribute password or its version tracker
	if !data.Password.Equal(state.Password) {
		tflog.Debug(ctx, fmt.Sprintf("password has changed for smppuser"))
		hasChange = true
	} else if !data.PasswordWoVersion.Equal(state.PasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("password_wo_version has changed for smppuser"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		smppuser := smppuserGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		smppuserGetThePayloadFromtheConfig(ctx, &config, &smppuser)
		// Make API call
		// Named resource - use UpdateResource
		username_value := data.Username.ValueString()
		_, err := r.client.UpdateResource(service.Smppuser.Type(), username_value, &smppuser)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update smppuser, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated smppuser resource")
	} else {
		tflog.Debug(ctx, "No changes detected for smppuser resource, skipping update")
	}

	// Read the updated state back
	r.readSmppuserFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SmppuserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SmppuserResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting smppuser resource")
	// Named resource - delete using DeleteResource
	username_value := data.Username.ValueString()
	err := r.client.DeleteResource(service.Smppuser.Type(), username_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete smppuser, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted smppuser resource")
}

// Helper function to read smppuser data from API
func (r *SmppuserResource) readSmppuserFromApi(ctx context.Context, data *SmppuserResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	username_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Smppuser.Type(), username_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read smppuser, got error: %s", err))
		return
	}

	smppuserSetAttrFromGet(ctx, data, getResponseData)

}
