package aaauser_vpnintranetapplication_binding

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
var _ resource.Resource = &AaauserVpnintranetapplicationBindingResource{}
var _ resource.ResourceWithConfigure = (*AaauserVpnintranetapplicationBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AaauserVpnintranetapplicationBindingResource)(nil)

func NewAaauserVpnintranetapplicationBindingResource() resource.Resource {
	return &AaauserVpnintranetapplicationBindingResource{}
}

// AaauserVpnintranetapplicationBindingResource defines the resource implementation.
type AaauserVpnintranetapplicationBindingResource struct {
	client *service.NitroClient
}

func (r *AaauserVpnintranetapplicationBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaauserVpnintranetapplicationBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaauser_vpnintranetapplication_binding"
}

func (r *AaauserVpnintranetapplicationBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaauserVpnintranetapplicationBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaauserVpnintranetapplicationBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaauser_vpnintranetapplication_binding resource")

	// aaauser_vpnintranetapplication_binding := aaauser_vpnintranetapplication_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaauser_vpnintranetapplication_binding.Type(), &aaauser_vpnintranetapplication_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaauser_vpnintranetapplication_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("aaauser_vpnintranetapplication_binding-config")

	tflog.Trace(ctx, "Created aaauser_vpnintranetapplication_binding resource")

	// Read the updated state back
	r.readAaauserVpnintranetapplicationBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaauserVpnintranetapplicationBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaauserVpnintranetapplicationBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaauser_vpnintranetapplication_binding resource")

	r.readAaauserVpnintranetapplicationBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaauserVpnintranetapplicationBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AaauserVpnintranetapplicationBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating aaauser_vpnintranetapplication_binding resource")

	// Create API request body from the model
	// aaauser_vpnintranetapplication_binding := aaauser_vpnintranetapplication_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaauser_vpnintranetapplication_binding.Type(), &aaauser_vpnintranetapplication_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaauser_vpnintranetapplication_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated aaauser_vpnintranetapplication_binding resource")

	// Read the updated state back
	r.readAaauserVpnintranetapplicationBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaauserVpnintranetapplicationBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaauserVpnintranetapplicationBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaauser_vpnintranetapplication_binding resource")

	// For aaauser_vpnintranetapplication_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted aaauser_vpnintranetapplication_binding resource from state")
}

// Helper function to read aaauser_vpnintranetapplication_binding data from API
func (r *AaauserVpnintranetapplicationBindingResource) readAaauserVpnintranetapplicationBindingFromApi(ctx context.Context, data *AaauserVpnintranetapplicationBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Aaauser_vpnintranetapplication_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaauser_vpnintranetapplication_binding, got error: %s", err))
		return
	}

	aaauser_vpnintranetapplication_bindingSetAttrFromGet(ctx, data, getResponseData)

}
