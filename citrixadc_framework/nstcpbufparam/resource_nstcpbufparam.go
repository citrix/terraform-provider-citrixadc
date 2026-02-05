package nstcpbufparam

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
var _ resource.Resource = &NstcpbufparamResource{}
var _ resource.ResourceWithConfigure = (*NstcpbufparamResource)(nil)
var _ resource.ResourceWithImportState = (*NstcpbufparamResource)(nil)

func NewNstcpbufparamResource() resource.Resource {
	return &NstcpbufparamResource{}
}

// NstcpbufparamResource defines the resource implementation.
type NstcpbufparamResource struct {
	client *service.NitroClient
}

func (r *NstcpbufparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NstcpbufparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nstcpbufparam"
}

func (r *NstcpbufparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NstcpbufparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NstcpbufparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nstcpbufparam resource")

	// nstcpbufparam := nstcpbufparamGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nstcpbufparam.Type(), &nstcpbufparam)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nstcpbufparam, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nstcpbufparam-config")

	tflog.Trace(ctx, "Created nstcpbufparam resource")

	// Read the updated state back
	r.readNstcpbufparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstcpbufparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NstcpbufparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nstcpbufparam resource")

	r.readNstcpbufparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstcpbufparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NstcpbufparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nstcpbufparam resource")

	// Create API request body from the model
	// nstcpbufparam := nstcpbufparamGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nstcpbufparam.Type(), &nstcpbufparam)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nstcpbufparam, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nstcpbufparam resource")

	// Read the updated state back
	r.readNstcpbufparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstcpbufparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NstcpbufparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nstcpbufparam resource")

	// For nstcpbufparam, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nstcpbufparam resource from state")
}

// Helper function to read nstcpbufparam data from API
func (r *NstcpbufparamResource) readNstcpbufparamFromApi(ctx context.Context, data *NstcpbufparamResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nstcpbufparam.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nstcpbufparam, got error: %s", err))
		return
	}

	nstcpbufparamSetAttrFromGet(ctx, data, getResponseData)

}
