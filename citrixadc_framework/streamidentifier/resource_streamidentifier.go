package streamidentifier

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
var _ resource.Resource = &StreamidentifierResource{}
var _ resource.ResourceWithConfigure = (*StreamidentifierResource)(nil)
var _ resource.ResourceWithImportState = (*StreamidentifierResource)(nil)

func NewStreamidentifierResource() resource.Resource {
	return &StreamidentifierResource{}
}

// StreamidentifierResource defines the resource implementation.
type StreamidentifierResource struct {
	client *service.NitroClient
}

func (r *StreamidentifierResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *StreamidentifierResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_streamidentifier"
}

func (r *StreamidentifierResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *StreamidentifierResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data StreamidentifierResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating streamidentifier resource")

	// streamidentifier := streamidentifierGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Streamidentifier.Type(), &streamidentifier)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create streamidentifier, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("streamidentifier-config")

	tflog.Trace(ctx, "Created streamidentifier resource")

	// Read the updated state back
	r.readStreamidentifierFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StreamidentifierResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data StreamidentifierResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading streamidentifier resource")

	r.readStreamidentifierFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StreamidentifierResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data StreamidentifierResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating streamidentifier resource")

	// Create API request body from the model
	// streamidentifier := streamidentifierGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Streamidentifier.Type(), &streamidentifier)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update streamidentifier, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated streamidentifier resource")

	// Read the updated state back
	r.readStreamidentifierFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StreamidentifierResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data StreamidentifierResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting streamidentifier resource")

	// For streamidentifier, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted streamidentifier resource from state")
}

// Helper function to read streamidentifier data from API
func (r *StreamidentifierResource) readStreamidentifierFromApi(ctx context.Context, data *StreamidentifierResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Streamidentifier.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read streamidentifier, got error: %s", err))
		return
	}

	streamidentifierSetAttrFromGet(ctx, data, getResponseData)

}
