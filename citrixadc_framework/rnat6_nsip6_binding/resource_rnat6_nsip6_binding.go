package rnat6_nsip6_binding

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
var _ resource.Resource = &Rnat6Nsip6BindingResource{}
var _ resource.ResourceWithConfigure = (*Rnat6Nsip6BindingResource)(nil)
var _ resource.ResourceWithImportState = (*Rnat6Nsip6BindingResource)(nil)

func NewRnat6Nsip6BindingResource() resource.Resource {
	return &Rnat6Nsip6BindingResource{}
}

// Rnat6Nsip6BindingResource defines the resource implementation.
type Rnat6Nsip6BindingResource struct {
	client *service.NitroClient
}

func (r *Rnat6Nsip6BindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Rnat6Nsip6BindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rnat6_nsip6_binding"
}

func (r *Rnat6Nsip6BindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Rnat6Nsip6BindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Rnat6Nsip6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rnat6_nsip6_binding resource")

	// rnat6_nsip6_binding := rnat6_nsip6_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Rnat6_nsip6_binding.Type(), &rnat6_nsip6_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create rnat6_nsip6_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("rnat6_nsip6_binding-config")

	tflog.Trace(ctx, "Created rnat6_nsip6_binding resource")

	// Read the updated state back
	r.readRnat6Nsip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Rnat6Nsip6BindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Rnat6Nsip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading rnat6_nsip6_binding resource")

	r.readRnat6Nsip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Rnat6Nsip6BindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data Rnat6Nsip6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating rnat6_nsip6_binding resource")

	// Create API request body from the model
	// rnat6_nsip6_binding := rnat6_nsip6_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Rnat6_nsip6_binding.Type(), &rnat6_nsip6_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update rnat6_nsip6_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated rnat6_nsip6_binding resource")

	// Read the updated state back
	r.readRnat6Nsip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Rnat6Nsip6BindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Rnat6Nsip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting rnat6_nsip6_binding resource")

	// For rnat6_nsip6_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted rnat6_nsip6_binding resource from state")
}

// Helper function to read rnat6_nsip6_binding data from API
func (r *Rnat6Nsip6BindingResource) readRnat6Nsip6BindingFromApi(ctx context.Context, data *Rnat6Nsip6BindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Rnat6_nsip6_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read rnat6_nsip6_binding, got error: %s", err))
		return
	}

	rnat6_nsip6_bindingSetAttrFromGet(ctx, data, getResponseData)

}
