package nslicenseproxyserver

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
var _ resource.Resource = &NslicenseproxyserverResource{}
var _ resource.ResourceWithConfigure = (*NslicenseproxyserverResource)(nil)
var _ resource.ResourceWithImportState = (*NslicenseproxyserverResource)(nil)

func NewNslicenseproxyserverResource() resource.Resource {
	return &NslicenseproxyserverResource{}
}

// NslicenseproxyserverResource defines the resource implementation.
type NslicenseproxyserverResource struct {
	client *service.NitroClient
}

func (r *NslicenseproxyserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NslicenseproxyserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nslicenseproxyserver"
}

func (r *NslicenseproxyserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NslicenseproxyserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NslicenseproxyserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nslicenseproxyserver resource")

	// nslicenseproxyserver := nslicenseproxyserverGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nslicenseproxyserver.Type(), &nslicenseproxyserver)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nslicenseproxyserver, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nslicenseproxyserver-config")

	tflog.Trace(ctx, "Created nslicenseproxyserver resource")

	// Read the updated state back
	r.readNslicenseproxyserverFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslicenseproxyserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NslicenseproxyserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nslicenseproxyserver resource")

	r.readNslicenseproxyserverFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslicenseproxyserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NslicenseproxyserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nslicenseproxyserver resource")

	// Create API request body from the model
	// nslicenseproxyserver := nslicenseproxyserverGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nslicenseproxyserver.Type(), &nslicenseproxyserver)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nslicenseproxyserver, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nslicenseproxyserver resource")

	// Read the updated state back
	r.readNslicenseproxyserverFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslicenseproxyserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NslicenseproxyserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nslicenseproxyserver resource")

	// For nslicenseproxyserver, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nslicenseproxyserver resource from state")
}

// Helper function to read nslicenseproxyserver data from API
func (r *NslicenseproxyserverResource) readNslicenseproxyserverFromApi(ctx context.Context, data *NslicenseproxyserverResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nslicenseproxyserver.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nslicenseproxyserver, got error: %s", err))
		return
	}

	nslicenseproxyserverSetAttrFromGet(ctx, data, getResponseData)

}
