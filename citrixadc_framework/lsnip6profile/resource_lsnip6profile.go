package lsnip6profile

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

	// lsnip6profile := lsnip6profileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsnip6profile.Type(), &lsnip6profile)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnip6profile, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lsnip6profile-config")

	tflog.Trace(ctx, "Created lsnip6profile resource")

	// Read the updated state back
	r.readLsnip6profileFromApi(ctx, &data, &resp.Diagnostics)

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

	r.readLsnip6profileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Lsnip6profileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data Lsnip6profileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lsnip6profile resource")

	// Create API request body from the model
	// lsnip6profile := lsnip6profileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsnip6profile.Type(), &lsnip6profile)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnip6profile, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lsnip6profile resource")

	// Read the updated state back
	r.readLsnip6profileFromApi(ctx, &data, &resp.Diagnostics)

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

	// For lsnip6profile, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lsnip6profile resource from state")
}

// Helper function to read lsnip6profile data from API
func (r *Lsnip6profileResource) readLsnip6profileFromApi(ctx context.Context, data *Lsnip6profileResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lsnip6profile.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnip6profile, got error: %s", err))
		return
	}

	lsnip6profileSetAttrFromGet(ctx, data, getResponseData)

}
