package lsnip6profile

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
var _ resource.Resource = &Lsnip6profileResource{}
var _ resource.ResourceWithConfigure = (*Lsnip6profileResource)(nil)
var _ resource.ResourceWithImportState = (*Lsnip6profileResource)(nil)

func NewLsnip6profileResource() resource.Resource {
	return &Lsnip6profileResource{}
}

// Lsnip6profileResource defines the resource implementation.
type Lsnip6profileResource struct {
	client *service.NitroClient
}

func (r *Lsnip6profileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Lsnip6profileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnip6profile"
}

func (r *Lsnip6profileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Lsnip6profileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Lsnip6profileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnip6profile resource")

	lsnip6profile := lsnip6profileGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	name := data.Name.ValueString()
	_, err := r.client.AddResource(service.Lsnip6profile.Type(), name, &lsnip6profile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnip6profile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lsnip6profile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readLsnip6profileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsnip6profile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Lsnip6profileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Lsnip6profileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnip6profile resource")

	found := r.readLsnip6profileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *Lsnip6profileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state Lsnip6profileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// lsnip6profile has no NITRO-updatable attributes; every attribute is ForceNew
	// (RequiresReplace / RequiresReplaceIfConfigured), so any meaningful change forces a
	// replace and Update is never invoked with real changes. Just refresh state from the ADC.
	tflog.Debug(ctx, "Updating lsnip6profile resource (refresh only - no updatable attributes)")

	if !r.readLsnip6profileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsnip6profile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Lsnip6profileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Lsnip6profileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnip6profile resource")

	// Named resource - delete using DeleteResource (ID is the plain name value)
	name := data.Id.ValueString()
	err := r.client.DeleteResource(service.Lsnip6profile.Type(), name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsnip6profile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsnip6profile resource")
}

// Helper function to read lsnip6profile data from API.
// Returns false when the resource no longer exists on the ADC.
func (r *Lsnip6profileResource) readLsnip6profileFromApi(ctx context.Context, data *Lsnip6profileResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain value (name)
	name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Lsnip6profile.Type(), name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnip6profile, got error: %s", err))
		return false
	}

	lsnip6profileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
