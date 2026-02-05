package nat64param

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
var _ resource.Resource = &Nat64paramResource{}
var _ resource.ResourceWithConfigure = (*Nat64paramResource)(nil)
var _ resource.ResourceWithImportState = (*Nat64paramResource)(nil)

func NewNat64paramResource() resource.Resource {
	return &Nat64paramResource{}
}

// Nat64paramResource defines the resource implementation.
type Nat64paramResource struct {
	client *service.NitroClient
}

func (r *Nat64paramResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Nat64paramResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nat64param"
}

func (r *Nat64paramResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Nat64paramResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Nat64paramResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nat64param resource")

	// nat64param := nat64paramGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nat64param.Type(), &nat64param)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nat64param, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nat64param-config")

	tflog.Trace(ctx, "Created nat64param resource")

	// Read the updated state back
	r.readNat64paramFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nat64paramResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Nat64paramResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nat64param resource")

	r.readNat64paramFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nat64paramResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data Nat64paramResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nat64param resource")

	// Create API request body from the model
	// nat64param := nat64paramGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nat64param.Type(), &nat64param)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nat64param, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nat64param resource")

	// Read the updated state back
	r.readNat64paramFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nat64paramResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Nat64paramResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nat64param resource")

	// For nat64param, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nat64param resource from state")
}

// Helper function to read nat64param data from API
func (r *Nat64paramResource) readNat64paramFromApi(ctx context.Context, data *Nat64paramResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nat64param.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nat64param, got error: %s", err))
		return
	}

	nat64paramSetAttrFromGet(ctx, data, getResponseData)

}
