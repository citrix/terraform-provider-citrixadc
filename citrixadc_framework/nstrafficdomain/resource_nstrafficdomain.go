package nstrafficdomain

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
var _ resource.Resource = &NstrafficdomainResource{}
var _ resource.ResourceWithConfigure = (*NstrafficdomainResource)(nil)
var _ resource.ResourceWithImportState = (*NstrafficdomainResource)(nil)

func NewNstrafficdomainResource() resource.Resource {
	return &NstrafficdomainResource{}
}

// NstrafficdomainResource defines the resource implementation.
type NstrafficdomainResource struct {
	client *service.NitroClient
}

func (r *NstrafficdomainResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NstrafficdomainResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nstrafficdomain"
}

func (r *NstrafficdomainResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NstrafficdomainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NstrafficdomainResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nstrafficdomain resource")

	// nstrafficdomain := nstrafficdomainGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nstrafficdomain.Type(), &nstrafficdomain)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nstrafficdomain, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nstrafficdomain-config")

	tflog.Trace(ctx, "Created nstrafficdomain resource")

	// Read the updated state back
	r.readNstrafficdomainFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstrafficdomainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NstrafficdomainResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nstrafficdomain resource")

	r.readNstrafficdomainFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstrafficdomainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NstrafficdomainResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nstrafficdomain resource")

	// Create API request body from the model
	// nstrafficdomain := nstrafficdomainGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nstrafficdomain.Type(), &nstrafficdomain)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nstrafficdomain, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nstrafficdomain resource")

	// Read the updated state back
	r.readNstrafficdomainFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstrafficdomainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NstrafficdomainResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nstrafficdomain resource")

	// For nstrafficdomain, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nstrafficdomain resource from state")
}

// Helper function to read nstrafficdomain data from API
func (r *NstrafficdomainResource) readNstrafficdomainFromApi(ctx context.Context, data *NstrafficdomainResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nstrafficdomain.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nstrafficdomain, got error: %s", err))
		return
	}

	nstrafficdomainSetAttrFromGet(ctx, data, getResponseData)

}
