package snmpcommunity

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
var _ resource.Resource = &SnmpcommunityResource{}
var _ resource.ResourceWithConfigure = (*SnmpcommunityResource)(nil)
var _ resource.ResourceWithImportState = (*SnmpcommunityResource)(nil)

func NewSnmpcommunityResource() resource.Resource {
	return &SnmpcommunityResource{}
}

// SnmpcommunityResource defines the resource implementation.
type SnmpcommunityResource struct {
	client *service.NitroClient
}

func (r *SnmpcommunityResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SnmpcommunityResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmpcommunity"
}

func (r *SnmpcommunityResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SnmpcommunityResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SnmpcommunityResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating snmpcommunity resource")

	// snmpcommunity := snmpcommunityGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Snmpcommunity.Type(), &snmpcommunity)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create snmpcommunity, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("snmpcommunity-config")

	tflog.Trace(ctx, "Created snmpcommunity resource")

	// Read the updated state back
	r.readSnmpcommunityFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpcommunityResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SnmpcommunityResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading snmpcommunity resource")

	r.readSnmpcommunityFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpcommunityResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SnmpcommunityResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating snmpcommunity resource")

	// Create API request body from the model
	// snmpcommunity := snmpcommunityGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Snmpcommunity.Type(), &snmpcommunity)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update snmpcommunity, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated snmpcommunity resource")

	// Read the updated state back
	r.readSnmpcommunityFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpcommunityResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SnmpcommunityResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting snmpcommunity resource")

	// For snmpcommunity, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted snmpcommunity resource from state")
}

// Helper function to read snmpcommunity data from API
func (r *SnmpcommunityResource) readSnmpcommunityFromApi(ctx context.Context, data *SnmpcommunityResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Snmpcommunity.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read snmpcommunity, got error: %s", err))
		return
	}

	snmpcommunitySetAttrFromGet(ctx, data, getResponseData)

}
