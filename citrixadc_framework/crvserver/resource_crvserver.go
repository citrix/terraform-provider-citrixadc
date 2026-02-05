package crvserver

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
var _ resource.Resource = &CrvserverResource{}
var _ resource.ResourceWithConfigure = (*CrvserverResource)(nil)
var _ resource.ResourceWithImportState = (*CrvserverResource)(nil)

func NewCrvserverResource() resource.Resource {
	return &CrvserverResource{}
}

// CrvserverResource defines the resource implementation.
type CrvserverResource struct {
	client *service.NitroClient
}

func (r *CrvserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CrvserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crvserver"
}

func (r *CrvserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CrvserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CrvserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating crvserver resource")

	// crvserver := crvserverGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Crvserver.Type(), &crvserver)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create crvserver, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("crvserver-config")

	tflog.Trace(ctx, "Created crvserver resource")

	// Read the updated state back
	r.readCrvserverFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrvserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CrvserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading crvserver resource")

	r.readCrvserverFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrvserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CrvserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating crvserver resource")

	// Create API request body from the model
	// crvserver := crvserverGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Crvserver.Type(), &crvserver)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update crvserver, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated crvserver resource")

	// Read the updated state back
	r.readCrvserverFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrvserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CrvserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting crvserver resource")

	// For crvserver, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted crvserver resource from state")
}

// Helper function to read crvserver data from API
func (r *CrvserverResource) readCrvserverFromApi(ctx context.Context, data *CrvserverResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Crvserver.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read crvserver, got error: %s", err))
		return
	}

	crvserverSetAttrFromGet(ctx, data, getResponseData)

}
