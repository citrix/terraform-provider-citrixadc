package aaagroup_aaauser_binding

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
var _ resource.Resource = &AaagroupAaauserBindingResource{}
var _ resource.ResourceWithConfigure = (*AaagroupAaauserBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AaagroupAaauserBindingResource)(nil)

func NewAaagroupAaauserBindingResource() resource.Resource {
	return &AaagroupAaauserBindingResource{}
}

// AaagroupAaauserBindingResource defines the resource implementation.
type AaagroupAaauserBindingResource struct {
	client *service.NitroClient
}

func (r *AaagroupAaauserBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaagroupAaauserBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaagroup_aaauser_binding"
}

func (r *AaagroupAaauserBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaagroupAaauserBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaagroupAaauserBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaagroup_aaauser_binding resource")

	// aaagroup_aaauser_binding := aaagroup_aaauser_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.AaagroupAaauserBinding.Type(), &aaagroup_aaauser_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaagroup_aaauser_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("aaagroup_aaauser_binding-config")

	tflog.Trace(ctx, "Created aaagroup_aaauser_binding resource")

	// Read the updated state back
	r.readAaagroupAaauserBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupAaauserBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaagroupAaauserBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaagroup_aaauser_binding resource")

	r.readAaagroupAaauserBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupAaauserBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AaagroupAaauserBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating aaagroup_aaauser_binding resource")

	// Create API request body from the model
	// aaagroup_aaauser_binding := aaagroup_aaauser_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.AaagroupAaauserBinding.Type(), &aaagroup_aaauser_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaagroup_aaauser_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated aaagroup_aaauser_binding resource")

	// Read the updated state back
	r.readAaagroupAaauserBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupAaauserBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaagroupAaauserBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaagroup_aaauser_binding resource")

	// For aaagroup_aaauser_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted aaagroup_aaauser_binding resource from state")
}

// Helper function to read aaagroup_aaauser_binding data from API
func (r *AaagroupAaauserBindingResource) readAaagroupAaauserBindingFromApi(ctx context.Context, data *AaagroupAaauserBindingResourceModel, diags *diag.Diagnostics) {

	aaagroup_Name := data.Groupname.ValueString()

	getResponseData, err := r.client.FindResource(service.Aaagroup_aaauser_binding.Type(), aaagroup_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaagroup_aaauser_binding, got error: %s", err))
		return
	}

	aaagroup_aaauser_bindingSetAttrFromGet(ctx, data, getResponseData)

}
