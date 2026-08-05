package lsnappsattributes

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
var _ resource.Resource = &LsnappsattributesResource{}
var _ resource.ResourceWithConfigure = (*LsnappsattributesResource)(nil)
var _ resource.ResourceWithImportState = (*LsnappsattributesResource)(nil)

func NewLsnappsattributesResource() resource.Resource {
	return &LsnappsattributesResource{}
}

// LsnappsattributesResource defines the resource implementation.
type LsnappsattributesResource struct {
	client *service.NitroClient
}

func (r *LsnappsattributesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnappsattributesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnappsattributes"
}

func (r *LsnappsattributesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnappsattributesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnappsattributesResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnappsattributes resource")

	lsnappsattributes := lsnappsattributesGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	lsnappsattributesName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Lsnappsattributes.Type(), lsnappsattributesName, &lsnappsattributes)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnappsattributes, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lsnappsattributes resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(lsnappsattributesName)

	// Read the updated state back
	if !r.readLsnappsattributesFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsnappsattributes not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnappsattributesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnappsattributesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnappsattributes resource")

	found := r.readLsnappsattributesFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LsnappsattributesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsnappsattributesResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lsnappsattributes resource")

	// Only sessiontimeout is updateable in place; name/port/transportprotocol are
	// ForceNew and force recreation instead of reaching Update.
	hasChange := false
	if !data.Sessiontimeout.Equal(state.Sessiontimeout) {
		tflog.Debug(ctx, "sessiontimeout has changed for lsnappsattributes")
		hasChange = true
	}

	if hasChange {
		lsnappsattributes := lsnappsattributesGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Matches SDK v2: sessiontimeout is pushed via UpdateUnnamedResource.
		err := r.client.UpdateUnnamedResource(service.Lsnappsattributes.Type(), &lsnappsattributes)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnappsattributes, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated lsnappsattributes resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lsnappsattributes resource, skipping update")
	}

	// Read the updated state back
	if !r.readLsnappsattributesFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsnappsattributes not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnappsattributesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnappsattributesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnappsattributes resource")

	// Named resource - delete using DeleteResource
	lsnappsattributesName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Lsnappsattributes.Type(), lsnappsattributesName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsnappsattributes, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsnappsattributes resource")
}

// Helper function to read lsnappsattributes data from API
func (r *LsnappsattributesResource) readLsnappsattributesFromApi(ctx context.Context, data *LsnappsattributesResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain value (name)
	lsnappsattributesName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Lsnappsattributes.Type(), lsnappsattributesName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnappsattributes, got error: %s", err))
		return false
	}

	lsnappsattributesSetAttrFromGet(ctx, data, getResponseData)

	return true
}
