package clusterinstance

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
var _ resource.Resource = &ClusterinstanceResource{}
var _ resource.ResourceWithConfigure = (*ClusterinstanceResource)(nil)
var _ resource.ResourceWithImportState = (*ClusterinstanceResource)(nil)

func NewClusterinstanceResource() resource.Resource {
	return &ClusterinstanceResource{}
}

// ClusterinstanceResource defines the resource implementation.
type ClusterinstanceResource struct {
	client *service.NitroClient
}

func (r *ClusterinstanceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusterinstanceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusterinstance"
}

func (r *ClusterinstanceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusterinstanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusterinstanceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating clusterinstance resource")

	// clusterinstance := clusterinstanceGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Clusterinstance.Type(), &clusterinstance)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create clusterinstance, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("clusterinstance-config")

	tflog.Trace(ctx, "Created clusterinstance resource")

	// Read the updated state back
	r.readClusterinstanceFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterinstanceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusterinstanceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading clusterinstance resource")

	r.readClusterinstanceFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterinstanceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ClusterinstanceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating clusterinstance resource")

	// Create API request body from the model
	// clusterinstance := clusterinstanceGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Clusterinstance.Type(), &clusterinstance)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update clusterinstance, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated clusterinstance resource")

	// Read the updated state back
	r.readClusterinstanceFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterinstanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusterinstanceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting clusterinstance resource")

	// For clusterinstance, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted clusterinstance resource from state")
}

// Helper function to read clusterinstance data from API
func (r *ClusterinstanceResource) readClusterinstanceFromApi(ctx context.Context, data *ClusterinstanceResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Clusterinstance.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read clusterinstance, got error: %s", err))
		return
	}

	clusterinstanceSetAttrFromGet(ctx, data, getResponseData)

}
