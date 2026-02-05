package forwardingsession

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

	// forwardingsession := forwardingsessionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Forwardingsession.Type(), &forwardingsession)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create forwardingsession, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("forwardingsession-config")

	tflog.Trace(ctx, "Created forwardingsession resource")

	// Read the updated state back
	r.readForwardingsessionFromApi(ctx, &data, &resp.Diagnostics)

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

	r.readForwardingsessionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ForwardingsessionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ForwardingsessionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating forwardingsession resource")

	// Create API request body from the model
	// forwardingsession := forwardingsessionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Forwardingsession.Type(), &forwardingsession)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update forwardingsession, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated forwardingsession resource")

	// Read the updated state back
	r.readForwardingsessionFromApi(ctx, &data, &resp.Diagnostics)

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

	// For forwardingsession, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted forwardingsession resource from state")
}

// Helper function to read forwardingsession data from API
func (r *ForwardingsessionResource) readForwardingsessionFromApi(ctx context.Context, data *ForwardingsessionResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Forwardingsession.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read forwardingsession, got error: %s", err))
		return
	}

	forwardingsessionSetAttrFromGet(ctx, data, getResponseData)

}
