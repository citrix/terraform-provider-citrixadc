package tmformssoaction

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
var _ resource.Resource = &TmformssoactionResource{}
var _ resource.ResourceWithConfigure = (*TmformssoactionResource)(nil)
var _ resource.ResourceWithImportState = (*TmformssoactionResource)(nil)

func NewTmformssoactionResource() resource.Resource {
	return &TmformssoactionResource{}
}

// TmformssoactionResource defines the resource implementation.
type TmformssoactionResource struct {
	client *service.NitroClient
}

func (r *TmformssoactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TmformssoactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tmformssoaction"
}

func (r *TmformssoactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TmformssoactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TmformssoactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating tmformssoaction resource")

	// Create API request body from the model
	tmformssoaction := tmformssoactionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	tmformssoactionName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Tmformssoaction.Type(), tmformssoactionName, &tmformssoaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create tmformssoaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created tmformssoaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", tmformssoactionName))

	// Read the updated state back
	if !r.readTmformssoactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "tmformssoaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmformssoactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TmformssoactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading tmformssoaction resource")

	found := r.readTmformssoactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *TmformssoactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state TmformssoactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating tmformssoaction resource")

	// Create API request body from the model (name is included in the body).
	// SDK v2 used UpdateUnnamedResource (PUT with name in the body) for this
	// resource; preserve that backward-compatible behavior.
	tmformssoaction := tmformssoactionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	err := r.client.UpdateUnnamedResource(service.Tmformssoaction.Type(), &tmformssoaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update tmformssoaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated tmformssoaction resource")

	// Read the updated state back
	if !r.readTmformssoactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "tmformssoaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmformssoactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TmformssoactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting tmformssoaction resource")

	// Named resource - delete using DeleteResource
	tmformssoactionName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Tmformssoaction.Type(), tmformssoactionName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete tmformssoaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted tmformssoaction resource")
}

// Helper function to read tmformssoaction data from API.
// Returns false (without an error) when the resource no longer exists.
func (r *TmformssoactionResource) readTmformssoactionFromApi(ctx context.Context, data *TmformssoactionResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain name value
	tmformssoactionName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Tmformssoaction.Type(), tmformssoactionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read tmformssoaction, got error: %s", err))
		return false
	}

	tmformssoactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
