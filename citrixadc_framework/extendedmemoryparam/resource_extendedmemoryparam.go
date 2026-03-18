package extendedmemoryparam

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
var _ resource.Resource = &ExtendedmemoryparamResource{}
var _ resource.ResourceWithConfigure = (*ExtendedmemoryparamResource)(nil)
var _ resource.ResourceWithImportState = (*ExtendedmemoryparamResource)(nil)

func NewExtendedmemoryparamResource() resource.Resource {
	return &ExtendedmemoryparamResource{}
}

// ExtendedmemoryparamResource defines the resource implementation.
type ExtendedmemoryparamResource struct {
	client *service.NitroClient
}

func (r *ExtendedmemoryparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ExtendedmemoryparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_extendedmemoryparam"
}

func (r *ExtendedmemoryparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ExtendedmemoryparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ExtendedmemoryparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating extendedmemoryparam resource")

	// extendedmemoryparam := extendedmemoryparamGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Extendedmemoryparam.Type(), &extendedmemoryparam)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create extendedmemoryparam, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("extendedmemoryparam-config")

	tflog.Trace(ctx, "Created extendedmemoryparam resource")

	// Read the updated state back
	r.readExtendedmemoryparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ExtendedmemoryparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ExtendedmemoryparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading extendedmemoryparam resource")

	r.readExtendedmemoryparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ExtendedmemoryparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ExtendedmemoryparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating extendedmemoryparam resource")

	// Create API request body from the model
	// extendedmemoryparam := extendedmemoryparamGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Extendedmemoryparam.Type(), &extendedmemoryparam)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update extendedmemoryparam, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated extendedmemoryparam resource")

	// Read the updated state back
	r.readExtendedmemoryparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ExtendedmemoryparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ExtendedmemoryparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting extendedmemoryparam resource")

	// For extendedmemoryparam, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted extendedmemoryparam resource from state")
}

// Helper function to read extendedmemoryparam data from API
func (r *ExtendedmemoryparamResource) readExtendedmemoryparamFromApi(ctx context.Context, data *ExtendedmemoryparamResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Extendedmemoryparam.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read extendedmemoryparam, got error: %s", err))
		return
	}

	extendedmemoryparamSetAttrFromGet(ctx, data, getResponseData)

}
