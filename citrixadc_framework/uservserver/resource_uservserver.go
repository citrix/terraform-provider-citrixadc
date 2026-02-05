package uservserver

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
var _ resource.Resource = &UservserverResource{}
var _ resource.ResourceWithConfigure = (*UservserverResource)(nil)
var _ resource.ResourceWithImportState = (*UservserverResource)(nil)

func NewUservserverResource() resource.Resource {
	return &UservserverResource{}
}

// UservserverResource defines the resource implementation.
type UservserverResource struct {
	client *service.NitroClient
}

func (r *UservserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *UservserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_uservserver"
}

func (r *UservserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *UservserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data UservserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating uservserver resource")

	// uservserver := uservserverGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Uservserver.Type(), &uservserver)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create uservserver, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("uservserver-config")

	tflog.Trace(ctx, "Created uservserver resource")

	// Read the updated state back
	r.readUservserverFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UservserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data UservserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading uservserver resource")

	r.readUservserverFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UservserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data UservserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating uservserver resource")

	// Create API request body from the model
	// uservserver := uservserverGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Uservserver.Type(), &uservserver)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update uservserver, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated uservserver resource")

	// Read the updated state back
	r.readUservserverFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UservserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data UservserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting uservserver resource")

	// For uservserver, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted uservserver resource from state")
}

// Helper function to read uservserver data from API
func (r *UservserverResource) readUservserverFromApi(ctx context.Context, data *UservserverResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Uservserver.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read uservserver, got error: %s", err))
		return
	}

	uservserverSetAttrFromGet(ctx, data, getResponseData)

}
