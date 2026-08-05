package lsnstatic

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
var _ resource.Resource = &LsnstaticResource{}
var _ resource.ResourceWithConfigure = (*LsnstaticResource)(nil)
var _ resource.ResourceWithImportState = (*LsnstaticResource)(nil)

func NewLsnstaticResource() resource.Resource {
	return &LsnstaticResource{}
}

// LsnstaticResource defines the resource implementation.
type LsnstaticResource struct {
	client *service.NitroClient
}

func (r *LsnstaticResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnstaticResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnstatic"
}

func (r *LsnstaticResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnstaticResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnstaticResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnstatic resource")

	// Build the payload from the plan
	lsnstatic := lsnstaticGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	lsnstaticName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Lsnstatic.Type(), lsnstaticName, &lsnstatic)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnstatic, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lsnstatic resource")

	// Set ID for the resource (matches SDK v2 d.SetId(name)) before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", lsnstaticName))

	// Read the updated state back
	if !r.readLsnstaticFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsnstatic not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnstaticResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnstaticResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnstatic resource")

	found := r.readLsnstaticFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LsnstaticResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// All lsnstatic attributes are ForceNew in the SDK v2 contract, so any
	// user-driven change triggers a replacement (Delete + Create) rather than
	// an in-place Update. This method exists only to satisfy the framework
	// interface; it performs no NITRO write (SDK v2 had no update path) and
	// simply refreshes computed values from the ADC.
	var data, state LsnstaticResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lsnstatic resource")

	// Read the updated state back
	if !r.readLsnstaticFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsnstatic not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnstaticResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnstaticResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnstatic resource")

	// Named resource - delete using DeleteResource
	lsnstaticName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Lsnstatic.Type(), lsnstaticName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsnstatic, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsnstatic resource")
}

// Helper function to read lsnstatic data from API
func (r *LsnstaticResource) readLsnstaticFromApi(ctx context.Context, data *LsnstaticResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain name value
	lsnstaticName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Lsnstatic.Type(), lsnstaticName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnstatic, got error: %s", err))
		return false
	}

	lsnstaticSetAttrFromGet(ctx, data, getResponseData)

	return true
}
