package rsskeytype

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
var _ resource.Resource = &RsskeytypeResource{}
var _ resource.ResourceWithConfigure = (*RsskeytypeResource)(nil)
var _ resource.ResourceWithImportState = (*RsskeytypeResource)(nil)

func NewRsskeytypeResource() resource.Resource {
	return &RsskeytypeResource{}
}

// RsskeytypeResource defines the resource implementation.
type RsskeytypeResource struct {
	client *service.NitroClient
}

func (r *RsskeytypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RsskeytypeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rsskeytype"
}

func (r *RsskeytypeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RsskeytypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RsskeytypeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rsskeytype resource")

	// rsskeytype := rsskeytypeGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Rsskeytype.Type(), &rsskeytype)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create rsskeytype, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("rsskeytype-config")

	tflog.Trace(ctx, "Created rsskeytype resource")

	// Read the updated state back
	r.readRsskeytypeFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RsskeytypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RsskeytypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading rsskeytype resource")

	r.readRsskeytypeFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RsskeytypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data RsskeytypeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating rsskeytype resource")

	// Create API request body from the model
	// rsskeytype := rsskeytypeGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Rsskeytype.Type(), &rsskeytype)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update rsskeytype, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated rsskeytype resource")

	// Read the updated state back
	r.readRsskeytypeFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RsskeytypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RsskeytypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting rsskeytype resource")

	// For rsskeytype, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted rsskeytype resource from state")
}

// Helper function to read rsskeytype data from API
func (r *RsskeytypeResource) readRsskeytypeFromApi(ctx context.Context, data *RsskeytypeResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Rsskeytype.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read rsskeytype, got error: %s", err))
		return
	}

	rsskeytypeSetAttrFromGet(ctx, data, getResponseData)

}
