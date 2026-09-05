package streamselector

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
var _ resource.Resource = &StreamselectorResource{}
var _ resource.ResourceWithConfigure = (*StreamselectorResource)(nil)
var _ resource.ResourceWithImportState = (*StreamselectorResource)(nil)

func NewStreamselectorResource() resource.Resource {
	return &StreamselectorResource{}
}

// StreamselectorResource defines the resource implementation.
type StreamselectorResource struct {
	client *service.NitroClient
}

func (r *StreamselectorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *StreamselectorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_streamselector"
}

func (r *StreamselectorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *StreamselectorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data StreamselectorResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating streamselector resource")

	streamselector := streamselectorGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	streamselectorName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Streamselector.Type(), streamselectorName, &streamselector)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create streamselector, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created streamselector resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(streamselectorName)

	// Read the updated state back
	if !r.readStreamselectorFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "streamselector not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StreamselectorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data StreamselectorResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading streamselector resource")

	found := r.readStreamselectorFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *StreamselectorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state StreamselectorResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating streamselector resource")

	// name is ForceNew/RequiresReplace, so only rule can change here.
	hasChange := false
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for streamselector")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		streamselector := streamselectorGetThePayloadFromthePlan(ctx, &data)
		// Update uses the unnamed endpoint (PUT /nitro/v1/config/streamselector) with
		// name carried in the payload, matching the SDK v2 behavior.
		err := r.client.UpdateUnnamedResource(service.Streamselector.Type(), &streamselector)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update streamselector, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated streamselector resource")
	} else {
		tflog.Debug(ctx, "No changes detected for streamselector resource, skipping update")
	}

	// Read the updated state back
	if !r.readStreamselectorFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "streamselector not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StreamselectorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data StreamselectorResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting streamselector resource")

	// Named resource - delete using DeleteResource (keyed on the live ID)
	streamselectorName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Streamselector.Type(), streamselectorName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete streamselector, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted streamselector resource")
}

// Helper function to read streamselector data from API
func (r *StreamselectorResource) readStreamselectorFromApi(ctx context.Context, data *StreamselectorResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain name value
	streamselectorName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Streamselector.Type(), streamselectorName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read streamselector, got error: %s", err))
		return false
	}

	streamselectorSetAttrFromGet(ctx, data, getResponseData)

	return true
}
