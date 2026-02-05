package botsignature

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
var _ resource.Resource = &BotsignatureResource{}
var _ resource.ResourceWithConfigure = (*BotsignatureResource)(nil)
var _ resource.ResourceWithImportState = (*BotsignatureResource)(nil)

func NewBotsignatureResource() resource.Resource {
	return &BotsignatureResource{}
}

// BotsignatureResource defines the resource implementation.
type BotsignatureResource struct {
	client *service.NitroClient
}

func (r *BotsignatureResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BotsignatureResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botsignature"
}

func (r *BotsignatureResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BotsignatureResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BotsignatureResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating botsignature resource")

	// botsignature := botsignatureGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Botsignature.Type(), &botsignature)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create botsignature, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("botsignature-config")

	tflog.Trace(ctx, "Created botsignature resource")

	// Read the updated state back
	r.readBotsignatureFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotsignatureResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BotsignatureResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading botsignature resource")

	r.readBotsignatureFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotsignatureResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data BotsignatureResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating botsignature resource")

	// Create API request body from the model
	// botsignature := botsignatureGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Botsignature.Type(), &botsignature)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update botsignature, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated botsignature resource")

	// Read the updated state back
	r.readBotsignatureFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotsignatureResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BotsignatureResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting botsignature resource")

	// For botsignature, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted botsignature resource from state")
}

// Helper function to read botsignature data from API
func (r *BotsignatureResource) readBotsignatureFromApi(ctx context.Context, data *BotsignatureResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Botsignature.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read botsignature, got error: %s", err))
		return
	}

	botsignatureSetAttrFromGet(ctx, data, getResponseData)

}
