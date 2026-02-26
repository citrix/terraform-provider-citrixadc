package aaagroup_intranetip_binding

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
var _ resource.Resource = &AaagroupIntranetipBindingResource{}
var _ resource.ResourceWithConfigure = (*AaagroupIntranetipBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AaagroupIntranetipBindingResource)(nil)

func NewAaagroupIntranetipBindingResource() resource.Resource {
	return &AaagroupIntranetipBindingResource{}
}

// AaagroupIntranetipBindingResource defines the resource implementation.
type AaagroupIntranetipBindingResource struct {
	client *service.NitroClient
}

func (r *AaagroupIntranetipBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaagroupIntranetipBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaagroup_intranetip_binding"
}

func (r *AaagroupIntranetipBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaagroupIntranetipBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaagroupIntranetipBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaagroup_intranetip_binding resource")

	// aaagroup_intranetip_binding := aaagroup_intranetip_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaagroup_intranetip_binding.Type(), &aaagroup_intranetip_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaagroup_intranetip_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("aaagroup_intranetip_binding-config")

	tflog.Trace(ctx, "Created aaagroup_intranetip_binding resource")

	// Read the updated state back
	r.readAaagroupIntranetipBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupIntranetipBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaagroupIntranetipBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaagroup_intranetip_binding resource")

	r.readAaagroupIntranetipBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupIntranetipBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AaagroupIntranetipBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating aaagroup_intranetip_binding resource")

	// Create API request body from the model
	// aaagroup_intranetip_binding := aaagroup_intranetip_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaagroup_intranetip_binding.Type(), &aaagroup_intranetip_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaagroup_intranetip_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated aaagroup_intranetip_binding resource")

	// Read the updated state back
	r.readAaagroupIntranetipBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupIntranetipBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaagroupIntranetipBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaagroup_intranetip_binding resource")

	// For aaagroup_intranetip_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted aaagroup_intranetip_binding resource from state")
}

// Helper function to read aaagroup_intranetip_binding data from API
func (r *AaagroupIntranetipBindingResource) readAaagroupIntranetipBindingFromApi(ctx context.Context, data *AaagroupIntranetipBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Aaagroup_intranetip_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaagroup_intranetip_binding, got error: %s", err))
		return
	}

	aaagroup_intranetip_bindingSetAttrFromGet(ctx, data, getResponseData)

}
