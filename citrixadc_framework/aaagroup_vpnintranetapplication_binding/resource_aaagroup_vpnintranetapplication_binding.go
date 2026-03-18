package aaagroup_vpnintranetapplication_binding

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
var _ resource.Resource = &AaagroupVpnintranetapplicationBindingResource{}
var _ resource.ResourceWithConfigure = (*AaagroupVpnintranetapplicationBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AaagroupVpnintranetapplicationBindingResource)(nil)

func NewAaagroupVpnintranetapplicationBindingResource() resource.Resource {
	return &AaagroupVpnintranetapplicationBindingResource{}
}

// AaagroupVpnintranetapplicationBindingResource defines the resource implementation.
type AaagroupVpnintranetapplicationBindingResource struct {
	client *service.NitroClient
}

func (r *AaagroupVpnintranetapplicationBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaagroupVpnintranetapplicationBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaagroup_vpnintranetapplication_binding"
}

func (r *AaagroupVpnintranetapplicationBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaagroupVpnintranetapplicationBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaagroupVpnintranetapplicationBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaagroup_vpnintranetapplication_binding resource")

	// aaagroup_vpnintranetapplication_binding := aaagroup_vpnintranetapplication_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaagroup_vpnintranetapplication_binding.Type(), &aaagroup_vpnintranetapplication_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaagroup_vpnintranetapplication_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("aaagroup_vpnintranetapplication_binding-config")

	tflog.Trace(ctx, "Created aaagroup_vpnintranetapplication_binding resource")

	// Read the updated state back
	r.readAaagroupVpnintranetapplicationBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupVpnintranetapplicationBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaagroupVpnintranetapplicationBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaagroup_vpnintranetapplication_binding resource")

	r.readAaagroupVpnintranetapplicationBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupVpnintranetapplicationBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AaagroupVpnintranetapplicationBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating aaagroup_vpnintranetapplication_binding resource")

	// Create API request body from the model
	// aaagroup_vpnintranetapplication_binding := aaagroup_vpnintranetapplication_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaagroup_vpnintranetapplication_binding.Type(), &aaagroup_vpnintranetapplication_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaagroup_vpnintranetapplication_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated aaagroup_vpnintranetapplication_binding resource")

	// Read the updated state back
	r.readAaagroupVpnintranetapplicationBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupVpnintranetapplicationBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaagroupVpnintranetapplicationBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaagroup_vpnintranetapplication_binding resource")

	// For aaagroup_vpnintranetapplication_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted aaagroup_vpnintranetapplication_binding resource from state")
}

// Helper function to read aaagroup_vpnintranetapplication_binding data from API
func (r *AaagroupVpnintranetapplicationBindingResource) readAaagroupVpnintranetapplicationBindingFromApi(ctx context.Context, data *AaagroupVpnintranetapplicationBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Aaagroup_vpnintranetapplication_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaagroup_vpnintranetapplication_binding, got error: %s", err))
		return
	}

	aaagroup_vpnintranetapplication_bindingSetAttrFromGet(ctx, data, getResponseData)

}
