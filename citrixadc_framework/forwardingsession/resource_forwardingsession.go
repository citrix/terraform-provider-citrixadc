package forwardingsession

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
var _ resource.Resource = &ForwardingsessionResource{}
var _ resource.ResourceWithConfigure = (*ForwardingsessionResource)(nil)
var _ resource.ResourceWithImportState = (*ForwardingsessionResource)(nil)

func NewForwardingsessionResource() resource.Resource {
	return &ForwardingsessionResource{}
}

// ForwardingsessionResource defines the resource implementation.
type ForwardingsessionResource struct {
	client *service.NitroClient
}

func (r *ForwardingsessionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ForwardingsessionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_forwardingsession"
}

func (r *ForwardingsessionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ForwardingsessionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ForwardingsessionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating forwardingsession resource")

	// Create API request body from the model
	forwardingsession := forwardingsessionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	forwardingsessionName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Forwardingsession.Type(), forwardingsessionName, &forwardingsession)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create forwardingsession, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created forwardingsession resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", forwardingsessionName))

	// Read the updated state back
	if !r.readForwardingsessionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "forwardingsession not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ForwardingsessionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ForwardingsessionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading forwardingsession resource")

	found := r.readForwardingsessionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *ForwardingsessionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ForwardingsessionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating forwardingsession resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Acl6name.Equal(state.Acl6name) {
		tflog.Debug(ctx, "acl6name has changed for forwardingsession")
		hasChange = true
	}
	if !data.Aclname.Equal(state.Aclname) {
		tflog.Debug(ctx, "aclname has changed for forwardingsession")
		hasChange = true
	}
	if !data.Connfailover.Equal(state.Connfailover) {
		tflog.Debug(ctx, "connfailover has changed for forwardingsession")
		hasChange = true
	}
	if !data.Processlocal.Equal(state.Processlocal) {
		tflog.Debug(ctx, "processlocal has changed for forwardingsession")
		hasChange = true
	}
	if !data.Sourceroutecache.Equal(state.Sourceroutecache) {
		tflog.Debug(ctx, "sourceroutecache has changed for forwardingsession")
		hasChange = true
	}
	if !data.Td.Equal(state.Td) {
		tflog.Debug(ctx, "td has changed for forwardingsession")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model, restricted to NITRO-updatable fields
		// (excludes create-only params network, netmask, td).
		forwardingsession := forwardingsessionGetTheUpdatePayloadFromThePlan(ctx, &data)
		// Set the name key to the live id so the PUT addresses the existing resource.
		forwardingsession.Name = data.Id.ValueString()
		// Make API call
		// Named resource - use UpdateResource
		forwardingsessionName := data.Id.ValueString()
		_, err := r.client.UpdateResource(service.Forwardingsession.Type(), forwardingsessionName, &forwardingsession)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update forwardingsession, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated forwardingsession resource")
	} else {
		tflog.Debug(ctx, "No changes detected for forwardingsession resource, skipping update")
	}

	// Read the updated state back
	if !r.readForwardingsessionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "forwardingsession not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ForwardingsessionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ForwardingsessionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting forwardingsession resource")

	// Named resource - delete using DeleteResource keyed on the live ID (name)
	forwardingsessionName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Forwardingsession.Type(), forwardingsessionName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete forwardingsession, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted forwardingsession resource")
}

// Helper function to read forwardingsession data from API.
// Returns false (without an error diagnostic) when the resource no longer exists.
func (r *ForwardingsessionResource) readForwardingsessionFromApi(ctx context.Context, data *ForwardingsessionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (name)
	forwardingsessionName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Forwardingsession.Type(), forwardingsessionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read forwardingsession, got error: %s", err))
		return false
	}

	forwardingsessionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
