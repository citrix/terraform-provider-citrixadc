package nsappflowcollector

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
var _ resource.Resource = &NsappflowcollectorResource{}
var _ resource.ResourceWithConfigure = (*NsappflowcollectorResource)(nil)
var _ resource.ResourceWithImportState = (*NsappflowcollectorResource)(nil)

func NewNsappflowcollectorResource() resource.Resource {
	return &NsappflowcollectorResource{}
}

// NsappflowcollectorResource defines the resource implementation.
type NsappflowcollectorResource struct {
	client *service.NitroClient
}

func (r *NsappflowcollectorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsappflowcollectorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsappflowcollector"
}

func (r *NsappflowcollectorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsappflowcollectorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsappflowcollectorResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsappflowcollector resource")

	// nsappflowcollector := nsappflowcollectorGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nsappflowcollector.Type(), &nsappflowcollector)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsappflowcollector, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nsappflowcollector-config")

	tflog.Trace(ctx, "Created nsappflowcollector resource")

	// Read the updated state back
	r.readNsappflowcollectorFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsappflowcollectorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsappflowcollectorResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsappflowcollector resource")

	r.readNsappflowcollectorFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsappflowcollectorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NsappflowcollectorResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nsappflowcollector resource")

	// Create API request body from the model
	// nsappflowcollector := nsappflowcollectorGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nsappflowcollector.Type(), &nsappflowcollector)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsappflowcollector, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nsappflowcollector resource")

	// Read the updated state back
	r.readNsappflowcollectorFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsappflowcollectorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsappflowcollectorResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsappflowcollector resource")

	// For nsappflowcollector, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nsappflowcollector resource from state")
}

// Helper function to read nsappflowcollector data from API
func (r *NsappflowcollectorResource) readNsappflowcollectorFromApi(ctx context.Context, data *NsappflowcollectorResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nsappflowcollector.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsappflowcollector, got error: %s", err))
		return
	}

	nsappflowcollectorSetAttrFromGet(ctx, data, getResponseData)

}
