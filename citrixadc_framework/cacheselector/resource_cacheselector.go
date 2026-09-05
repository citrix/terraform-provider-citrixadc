package cacheselector

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
var _ resource.Resource = &CacheselectorResource{}
var _ resource.ResourceWithConfigure = (*CacheselectorResource)(nil)
var _ resource.ResourceWithImportState = (*CacheselectorResource)(nil)

func NewCacheselectorResource() resource.Resource {
	return &CacheselectorResource{}
}

// CacheselectorResource defines the resource implementation.
type CacheselectorResource struct {
	client *service.NitroClient
}

func (r *CacheselectorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CacheselectorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cacheselector"
}

func (r *CacheselectorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CacheselectorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CacheselectorResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cacheselector resource")

	cacheselector := cacheselectorGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	selectorname := data.Selectorname.ValueString()
	_, err := r.client.AddResource(service.Cacheselector.Type(), selectorname, &cacheselector)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cacheselector, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created cacheselector resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(selectorname)

	// Read the updated state back
	if !r.readCacheselectorFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "cacheselector not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheselectorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CacheselectorResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cacheselector resource")

	found := r.readCacheselectorFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *CacheselectorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CacheselectorResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating cacheselector resource")

	// selectorname is ForceNew/RequiresReplace, so only rule can change here.
	hasChange := false
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for cacheselector")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		cacheselector := cacheselectorGetThePayloadFromthePlan(ctx, &data)
		// Update uses the unnamed endpoint (PUT /nitro/v1/config/cacheselector) with
		// selectorname carried in the payload, matching the SDK v2 behavior.
		err := r.client.UpdateUnnamedResource(service.Cacheselector.Type(), &cacheselector)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cacheselector, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated cacheselector resource")
	} else {
		tflog.Debug(ctx, "No changes detected for cacheselector resource, skipping update")
	}

	// Read the updated state back
	if !r.readCacheselectorFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "cacheselector not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheselectorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CacheselectorResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cacheselector resource")

	// Named resource - delete using DeleteResource (keyed on the live ID)
	selectorname := data.Id.ValueString()
	err := r.client.DeleteResource(service.Cacheselector.Type(), selectorname)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete cacheselector, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted cacheselector resource")
}

// Helper function to read cacheselector data from API
func (r *CacheselectorResource) readCacheselectorFromApi(ctx context.Context, data *CacheselectorResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain selectorname value
	selectorname := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Cacheselector.Type(), selectorname)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cacheselector, got error: %s", err))
		return false
	}

	cacheselectorSetAttrFromGet(ctx, data, getResponseData)

	return true
}
