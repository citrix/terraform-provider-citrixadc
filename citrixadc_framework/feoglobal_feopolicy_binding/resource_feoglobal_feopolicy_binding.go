package feoglobal_feopolicy_binding

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
var _ resource.Resource = &FeoglobalFeopolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*FeoglobalFeopolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*FeoglobalFeopolicyBindingResource)(nil)

func NewFeoglobalFeopolicyBindingResource() resource.Resource {
	return &FeoglobalFeopolicyBindingResource{}
}

// FeoglobalFeopolicyBindingResource defines the resource implementation.
type FeoglobalFeopolicyBindingResource struct {
	client *service.NitroClient
}

func (r *FeoglobalFeopolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *FeoglobalFeopolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_feoglobal_feopolicy_binding"
}

func (r *FeoglobalFeopolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *FeoglobalFeopolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data FeoglobalFeopolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating feoglobal_feopolicy_binding resource")

	// feoglobal_feopolicy_binding := feoglobal_feopolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Feoglobal_feopolicy_binding.Type(), &feoglobal_feopolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create feoglobal_feopolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("feoglobal_feopolicy_binding-config")

	tflog.Trace(ctx, "Created feoglobal_feopolicy_binding resource")

	// Read the updated state back
	r.readFeoglobalFeopolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FeoglobalFeopolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data FeoglobalFeopolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading feoglobal_feopolicy_binding resource")

	r.readFeoglobalFeopolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FeoglobalFeopolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data FeoglobalFeopolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating feoglobal_feopolicy_binding resource")

	// Create API request body from the model
	// feoglobal_feopolicy_binding := feoglobal_feopolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Feoglobal_feopolicy_binding.Type(), &feoglobal_feopolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update feoglobal_feopolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated feoglobal_feopolicy_binding resource")

	// Read the updated state back
	r.readFeoglobalFeopolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FeoglobalFeopolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FeoglobalFeopolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting feoglobal_feopolicy_binding resource")

	// For feoglobal_feopolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted feoglobal_feopolicy_binding resource from state")
}

// Helper function to read feoglobal_feopolicy_binding data from API
func (r *FeoglobalFeopolicyBindingResource) readFeoglobalFeopolicyBindingFromApi(ctx context.Context, data *FeoglobalFeopolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Feoglobal_feopolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read feoglobal_feopolicy_binding, got error: %s", err))
		return
	}

	feoglobal_feopolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
