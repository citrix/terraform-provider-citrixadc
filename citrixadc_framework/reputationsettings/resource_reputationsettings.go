package reputationsettings

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
var _ resource.Resource = &ReputationsettingsResource{}
var _ resource.ResourceWithConfigure = (*ReputationsettingsResource)(nil)
var _ resource.ResourceWithImportState = (*ReputationsettingsResource)(nil)

func NewReputationsettingsResource() resource.Resource {
	return &ReputationsettingsResource{}
}

// ReputationsettingsResource defines the resource implementation.
type ReputationsettingsResource struct {
	client *service.NitroClient
}

func (r *ReputationsettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ReputationsettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_reputationsettings"
}

func (r *ReputationsettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ReputationsettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ReputationsettingsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating reputationsettings resource")

	// reputationsettings := reputationsettingsGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Reputationsettings.Type(), &reputationsettings)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create reputationsettings, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("reputationsettings-config")

	tflog.Trace(ctx, "Created reputationsettings resource")

	// Read the updated state back
	r.readReputationsettingsFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ReputationsettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ReputationsettingsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading reputationsettings resource")

	r.readReputationsettingsFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ReputationsettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ReputationsettingsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating reputationsettings resource")

	// Create API request body from the model
	// reputationsettings := reputationsettingsGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Reputationsettings.Type(), &reputationsettings)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update reputationsettings, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated reputationsettings resource")

	// Read the updated state back
	r.readReputationsettingsFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ReputationsettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ReputationsettingsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting reputationsettings resource")

	// For reputationsettings, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted reputationsettings resource from state")
}

// Helper function to read reputationsettings data from API
func (r *ReputationsettingsResource) readReputationsettingsFromApi(ctx context.Context, data *ReputationsettingsResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Reputationsettings.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read reputationsettings, got error: %s", err))
		return
	}

	reputationsettingsSetAttrFromGet(ctx, data, getResponseData)

}
