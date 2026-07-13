package dbuser

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
var _ resource.Resource = &DbuserResource{}
var _ resource.ResourceWithConfigure = (*DbuserResource)(nil)
var _ resource.ResourceWithImportState = (*DbuserResource)(nil)

func NewDbuserResource() resource.Resource {
	return &DbuserResource{}
}

// DbuserResource defines the resource implementation.
type DbuserResource struct {
	client *service.NitroClient
}

func (r *DbuserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DbuserResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dbuser"
}

func (r *DbuserResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DbuserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config DbuserResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dbuser resource")
	// Get payload from plan (regular attributes)
	dbuser := dbuserGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	dbuserGetThePayloadFromtheConfig(ctx, &config, &dbuser)

	// Make API call
	// Named resource - use AddResource
	username_value := data.Username.ValueString()
	_, err := r.client.AddResource(service.Dbuser.Type(), username_value, &dbuser)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dbuser, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dbuser resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Username.ValueString()))

	// Read the updated state back
	if !r.readDbuserFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dbuser not found immediately after create/update")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DbuserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DbuserResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dbuser resource")

	found := r.readDbuserFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *DbuserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state DbuserResourceModel

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

	tflog.Debug(ctx, "Updating dbuser resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	// Check secret attribute password or its version tracker
	if !data.Password.Equal(state.Password) {
		tflog.Debug(ctx, fmt.Sprintf("password has changed for dbuser"))
		hasChange = true
	} else if !data.PasswordWoVersion.Equal(state.PasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("password_wo_version has changed for dbuser"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		dbuser := dbuserGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		dbuserGetThePayloadFromtheConfig(ctx, &config, &dbuser)
		// Make API call
		// Named resource - use UpdateResource
		username_value := data.Username.ValueString()
		_, err := r.client.UpdateResource(service.Dbuser.Type(), username_value, &dbuser)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dbuser, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated dbuser resource")
	} else {
		tflog.Debug(ctx, "No changes detected for dbuser resource, skipping update")
	}

	// Read the updated state back
	if !r.readDbuserFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dbuser not found immediately after create/update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DbuserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DbuserResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dbuser resource")
	// Named resource - delete using DeleteResource
	username_value := data.Username.ValueString()
	err := r.client.DeleteResource(service.Dbuser.Type(), username_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dbuser, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dbuser resource")
}

// Helper function to read dbuser data from API
func (r *DbuserResource) readDbuserFromApi(ctx context.Context, data *DbuserResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	username_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Dbuser.Type(), username_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dbuser, got error: %s", err))
		return false
	}

	dbuserSetAttrFromGet(ctx, data, getResponseData)

	return true
}
