package onlinkipv6prefix

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
var _ resource.Resource = &Onlinkipv6prefixResource{}
var _ resource.ResourceWithConfigure = (*Onlinkipv6prefixResource)(nil)
var _ resource.ResourceWithImportState = (*Onlinkipv6prefixResource)(nil)

func NewOnlinkipv6prefixResource() resource.Resource {
	return &Onlinkipv6prefixResource{}
}

// Onlinkipv6prefixResource defines the resource implementation.
type Onlinkipv6prefixResource struct {
	client *service.NitroClient
}

func (r *Onlinkipv6prefixResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Onlinkipv6prefixResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_onlinkipv6prefix"
}

func (r *Onlinkipv6prefixResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Onlinkipv6prefixResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Onlinkipv6prefixResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating onlinkipv6prefix resource")

	// onlinkipv6prefix := onlinkipv6prefixGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Onlinkipv6prefix.Type(), &onlinkipv6prefix)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create onlinkipv6prefix, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("onlinkipv6prefix-config")

	tflog.Trace(ctx, "Created onlinkipv6prefix resource")

	// Read the updated state back
	r.readOnlinkipv6prefixFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Onlinkipv6prefixResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Onlinkipv6prefixResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading onlinkipv6prefix resource")

	r.readOnlinkipv6prefixFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Onlinkipv6prefixResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data Onlinkipv6prefixResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating onlinkipv6prefix resource")

	// Create API request body from the model
	// onlinkipv6prefix := onlinkipv6prefixGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Onlinkipv6prefix.Type(), &onlinkipv6prefix)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update onlinkipv6prefix, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated onlinkipv6prefix resource")

	// Read the updated state back
	r.readOnlinkipv6prefixFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Onlinkipv6prefixResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Onlinkipv6prefixResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting onlinkipv6prefix resource")

	// For onlinkipv6prefix, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted onlinkipv6prefix resource from state")
}

// Helper function to read onlinkipv6prefix data from API
func (r *Onlinkipv6prefixResource) readOnlinkipv6prefixFromApi(ctx context.Context, data *Onlinkipv6prefixResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Onlinkipv6prefix.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read onlinkipv6prefix, got error: %s", err))
		return
	}

	onlinkipv6prefixSetAttrFromGet(ctx, data, getResponseData)

}
