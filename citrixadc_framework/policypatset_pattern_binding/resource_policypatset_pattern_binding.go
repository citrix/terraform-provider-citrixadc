package policypatset_pattern_binding

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
var _ resource.Resource = &PolicypatsetPatternBindingResource{}
var _ resource.ResourceWithConfigure = (*PolicypatsetPatternBindingResource)(nil)
var _ resource.ResourceWithImportState = (*PolicypatsetPatternBindingResource)(nil)

func NewPolicypatsetPatternBindingResource() resource.Resource {
	return &PolicypatsetPatternBindingResource{}
}

// PolicypatsetPatternBindingResource defines the resource implementation.
type PolicypatsetPatternBindingResource struct {
	client *service.NitroClient
}

func (r *PolicypatsetPatternBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *PolicypatsetPatternBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policypatset_pattern_binding"
}

func (r *PolicypatsetPatternBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PolicypatsetPatternBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PolicypatsetPatternBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating policypatset_pattern_binding resource")

	// policypatset_pattern_binding := policypatset_pattern_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Policypatset_pattern_binding.Type(), &policypatset_pattern_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create policypatset_pattern_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("policypatset_pattern_binding-config")

	tflog.Trace(ctx, "Created policypatset_pattern_binding resource")

	// Read the updated state back
	r.readPolicypatsetPatternBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicypatsetPatternBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data PolicypatsetPatternBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading policypatset_pattern_binding resource")

	r.readPolicypatsetPatternBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicypatsetPatternBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data PolicypatsetPatternBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating policypatset_pattern_binding resource")

	// Create API request body from the model
	// policypatset_pattern_binding := policypatset_pattern_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Policypatset_pattern_binding.Type(), &policypatset_pattern_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update policypatset_pattern_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated policypatset_pattern_binding resource")

	// Read the updated state back
	r.readPolicypatsetPatternBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicypatsetPatternBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data PolicypatsetPatternBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting policypatset_pattern_binding resource")

	// For policypatset_pattern_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted policypatset_pattern_binding resource from state")
}

// Helper function to read policypatset_pattern_binding data from API
func (r *PolicypatsetPatternBindingResource) readPolicypatsetPatternBindingFromApi(ctx context.Context, data *PolicypatsetPatternBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Policypatset_pattern_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read policypatset_pattern_binding, got error: %s", err))
		return
	}

	policypatset_pattern_bindingSetAttrFromGet(ctx, data, getResponseData)

}
