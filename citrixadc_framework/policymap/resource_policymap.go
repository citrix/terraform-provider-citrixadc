package policymap

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
var _ resource.Resource = &PolicymapResource{}
var _ resource.ResourceWithConfigure = (*PolicymapResource)(nil)
var _ resource.ResourceWithImportState = (*PolicymapResource)(nil)

func NewPolicymapResource() resource.Resource {
	return &PolicymapResource{}
}

// PolicymapResource defines the resource implementation.
type PolicymapResource struct {
	client *service.NitroClient
}

func (r *PolicymapResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *PolicymapResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policymap"
}

func (r *PolicymapResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PolicymapResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PolicymapResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating policymap resource")

	policymap := policymapGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	mappolicyname_value := data.Mappolicyname.ValueString()
	_, err := r.client.AddResource(service.Policymap.Type(), mappolicyname_value, &policymap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create policymap, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created policymap resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Mappolicyname.ValueString()))

	// Read the updated state back
	if !r.readPolicymapFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "policymap not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicymapResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data PolicymapResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading policymap resource")

	found := r.readPolicymapFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *PolicymapResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state PolicymapResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// policymap exposes no NITRO "set"/update operation; every attribute is
	// create-only (ForceNew in SDK v2 -> RequiresReplace/RequiresReplaceIfConfigured
	// here), so any real change triggers a destroy/recreate and this function is
	// effectively never invoked. Re-read to keep state coherent if it is.
	tflog.Debug(ctx, "Updating policymap resource (no updateable attributes; re-reading)")

	if !r.readPolicymapFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "policymap not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicymapResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data PolicymapResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting policymap resource")

	// Named resource - delete using DeleteResource keyed off the live ID
	mappolicyname_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Policymap.Type(), mappolicyname_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete policymap, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted policymap resource")
}

// Helper function to read policymap data from API
func (r *PolicymapResource) readPolicymapFromApi(ctx context.Context, data *PolicymapResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	mappolicyname_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Policymap.Type(), mappolicyname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read policymap, got error: %s", err))
		return false
	}

	policymapSetAttrFromGet(ctx, data, getResponseData)

	return true
}
