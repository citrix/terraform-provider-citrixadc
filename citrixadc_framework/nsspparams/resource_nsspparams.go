package nsspparams

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
var _ resource.Resource = &NsspparamsResource{}
var _ resource.ResourceWithConfigure = (*NsspparamsResource)(nil)
var _ resource.ResourceWithImportState = (*NsspparamsResource)(nil)

func NewNsspparamsResource() resource.Resource {
	return &NsspparamsResource{}
}

// NsspparamsResource defines the resource implementation.
type NsspparamsResource struct {
	client *service.NitroClient
}

func (r *NsspparamsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsspparamsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsspparams"
}

func (r *NsspparamsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsspparamsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsspparamsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsspparams resource")

	// nsspparams := nsspparamsGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nsspparams.Type(), &nsspparams)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsspparams, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nsspparams-config")

	tflog.Trace(ctx, "Created nsspparams resource")

	// Read the updated state back
	r.readNsspparamsFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsspparamsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsspparamsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsspparams resource")

	r.readNsspparamsFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsspparamsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NsspparamsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nsspparams resource")

	// Create API request body from the model
	// nsspparams := nsspparamsGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nsspparams.Type(), &nsspparams)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsspparams, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nsspparams resource")

	// Read the updated state back
	r.readNsspparamsFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsspparamsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsspparamsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsspparams resource")

	// For nsspparams, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nsspparams resource from state")
}

// Helper function to read nsspparams data from API
func (r *NsspparamsResource) readNsspparamsFromApi(ctx context.Context, data *NsspparamsResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nsspparams.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsspparams, got error: %s", err))
		return
	}

	nsspparamsSetAttrFromGet(ctx, data, getResponseData)

}
