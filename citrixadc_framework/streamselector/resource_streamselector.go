package streamselector

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

	// streamselector := streamselectorGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Streamselector.Type(), &streamselector)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create streamselector, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("streamselector-config")

	tflog.Trace(ctx, "Created streamselector resource")

	// Read the updated state back
	r.readStreamselectorFromApi(ctx, &data, &resp.Diagnostics)

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

	r.readStreamselectorFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StreamselectorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data StreamselectorResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating streamselector resource")

	// Create API request body from the model
	// streamselector := streamselectorGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Streamselector.Type(), &streamselector)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update streamselector, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated streamselector resource")

	// Read the updated state back
	r.readStreamselectorFromApi(ctx, &data, &resp.Diagnostics)

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

	// For streamselector, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted streamselector resource from state")
}

// Helper function to read streamselector data from API
func (r *StreamselectorResource) readStreamselectorFromApi(ctx context.Context, data *StreamselectorResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Streamselector.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read streamselector, got error: %s", err))
		return
	}

	streamselectorSetAttrFromGet(ctx, data, getResponseData)

}
