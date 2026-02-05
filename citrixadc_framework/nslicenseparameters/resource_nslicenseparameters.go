package nslicenseparameters

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
var _ resource.Resource = &NslicenseparametersResource{}
var _ resource.ResourceWithConfigure = (*NslicenseparametersResource)(nil)
var _ resource.ResourceWithImportState = (*NslicenseparametersResource)(nil)

func NewNslicenseparametersResource() resource.Resource {
	return &NslicenseparametersResource{}
}

// NslicenseparametersResource defines the resource implementation.
type NslicenseparametersResource struct {
	client *service.NitroClient
}

func (r *NslicenseparametersResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NslicenseparametersResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nslicenseparameters"
}

func (r *NslicenseparametersResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NslicenseparametersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NslicenseparametersResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nslicenseparameters resource")

	// nslicenseparameters := nslicenseparametersGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nslicenseparameters.Type(), &nslicenseparameters)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nslicenseparameters, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nslicenseparameters-config")

	tflog.Trace(ctx, "Created nslicenseparameters resource")

	// Read the updated state back
	r.readNslicenseparametersFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslicenseparametersResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NslicenseparametersResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nslicenseparameters resource")

	r.readNslicenseparametersFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslicenseparametersResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NslicenseparametersResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nslicenseparameters resource")

	// Create API request body from the model
	// nslicenseparameters := nslicenseparametersGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nslicenseparameters.Type(), &nslicenseparameters)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nslicenseparameters, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nslicenseparameters resource")

	// Read the updated state back
	r.readNslicenseparametersFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslicenseparametersResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NslicenseparametersResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nslicenseparameters resource")

	// For nslicenseparameters, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nslicenseparameters resource from state")
}

// Helper function to read nslicenseparameters data from API
func (r *NslicenseparametersResource) readNslicenseparametersFromApi(ctx context.Context, data *NslicenseparametersResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nslicenseparameters.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nslicenseparameters, got error: %s", err))
		return
	}

	nslicenseparametersSetAttrFromGet(ctx, data, getResponseData)

}
