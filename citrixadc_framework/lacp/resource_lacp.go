package lacp

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
var _ resource.Resource = &LacpResource{}
var _ resource.ResourceWithConfigure = (*LacpResource)(nil)
var _ resource.ResourceWithImportState = (*LacpResource)(nil)

func NewLacpResource() resource.Resource {
	return &LacpResource{}
}

// LacpResource defines the resource implementation.
type LacpResource struct {
	client *service.NitroClient
}

func (r *LacpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LacpResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lacp"
}

func (r *LacpResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LacpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LacpResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lacp resource")

	// lacp := lacpGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lacp.Type(), &lacp)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lacp, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lacp-config")

	tflog.Trace(ctx, "Created lacp resource")

	// Read the updated state back
	r.readLacpFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LacpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LacpResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lacp resource")

	r.readLacpFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LacpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LacpResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lacp resource")

	// Create API request body from the model
	// lacp := lacpGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lacp.Type(), &lacp)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lacp, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lacp resource")

	// Read the updated state back
	r.readLacpFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LacpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LacpResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lacp resource")

	// For lacp, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lacp resource from state")
}

// Helper function to read lacp data from API
func (r *LacpResource) readLacpFromApi(ctx context.Context, data *LacpResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lacp.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lacp, got error: %s", err))
		return
	}

	lacpSetAttrFromGet(ctx, data, getResponseData)

}
