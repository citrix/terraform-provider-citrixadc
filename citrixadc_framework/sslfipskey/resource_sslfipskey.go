package sslfipskey

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
var _ resource.Resource = &SslfipskeyResource{}
var _ resource.ResourceWithConfigure = (*SslfipskeyResource)(nil)
var _ resource.ResourceWithImportState = (*SslfipskeyResource)(nil)

func NewSslfipskeyResource() resource.Resource {
	return &SslfipskeyResource{}
}

// SslfipskeyResource defines the resource implementation.
type SslfipskeyResource struct {
	client *service.NitroClient
}

func (r *SslfipskeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslfipskeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslfipskey"
}

func (r *SslfipskeyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslfipskeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslfipskeyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslfipskey resource")

	// sslfipskey := sslfipskeyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslfipskey.Type(), &sslfipskey)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslfipskey, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("sslfipskey-config")

	tflog.Trace(ctx, "Created sslfipskey resource")

	// Read the updated state back
	r.readSslfipskeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipskeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslfipskeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslfipskey resource")

	r.readSslfipskeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipskeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SslfipskeyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating sslfipskey resource")

	// Create API request body from the model
	// sslfipskey := sslfipskeyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslfipskey.Type(), &sslfipskey)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslfipskey, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated sslfipskey resource")

	// Read the updated state back
	r.readSslfipskeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipskeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslfipskeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslfipskey resource")

	// For sslfipskey, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted sslfipskey resource from state")
}

// Helper function to read sslfipskey data from API
func (r *SslfipskeyResource) readSslfipskeyFromApi(ctx context.Context, data *SslfipskeyResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Sslfipskey.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslfipskey, got error: %s", err))
		return
	}

	sslfipskeySetAttrFromGet(ctx, data, getResponseData)

}
