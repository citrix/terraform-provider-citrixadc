package systemextramgmtcpu

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
var _ resource.Resource = &SystemextramgmtcpuResource{}
var _ resource.ResourceWithConfigure = (*SystemextramgmtcpuResource)(nil)
var _ resource.ResourceWithImportState = (*SystemextramgmtcpuResource)(nil)

func NewSystemextramgmtcpuResource() resource.Resource {
	return &SystemextramgmtcpuResource{}
}

// SystemextramgmtcpuResource defines the resource implementation.
type SystemextramgmtcpuResource struct {
	client *service.NitroClient
}

func (r *SystemextramgmtcpuResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemextramgmtcpuResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemextramgmtcpu"
}

func (r *SystemextramgmtcpuResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemextramgmtcpuResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemextramgmtcpuResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemextramgmtcpu resource")

	// systemextramgmtcpu := systemextramgmtcpuGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Systemextramgmtcpu.Type(), &systemextramgmtcpu)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemextramgmtcpu, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("systemextramgmtcpu-config")

	tflog.Trace(ctx, "Created systemextramgmtcpu resource")

	// Read the updated state back
	r.readSystemextramgmtcpuFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemextramgmtcpuResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemextramgmtcpuResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemextramgmtcpu resource")

	r.readSystemextramgmtcpuFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemextramgmtcpuResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SystemextramgmtcpuResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating systemextramgmtcpu resource")

	// Create API request body from the model
	// systemextramgmtcpu := systemextramgmtcpuGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Systemextramgmtcpu.Type(), &systemextramgmtcpu)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemextramgmtcpu, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated systemextramgmtcpu resource")

	// Read the updated state back
	r.readSystemextramgmtcpuFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemextramgmtcpuResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemextramgmtcpuResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemextramgmtcpu resource")

	// For systemextramgmtcpu, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted systemextramgmtcpu resource from state")
}

// Helper function to read systemextramgmtcpu data from API
func (r *SystemextramgmtcpuResource) readSystemextramgmtcpuFromApi(ctx context.Context, data *SystemextramgmtcpuResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Systemextramgmtcpu.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemextramgmtcpu, got error: %s", err))
		return
	}

	systemextramgmtcpuSetAttrFromGet(ctx, data, getResponseData)

}
