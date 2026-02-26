package policystringmap_pattern_binding

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
var _ resource.Resource = &PolicystringmapPatternBindingResource{}
var _ resource.ResourceWithConfigure = (*PolicystringmapPatternBindingResource)(nil)
var _ resource.ResourceWithImportState = (*PolicystringmapPatternBindingResource)(nil)

func NewPolicystringmapPatternBindingResource() resource.Resource {
	return &PolicystringmapPatternBindingResource{}
}

// PolicystringmapPatternBindingResource defines the resource implementation.
type PolicystringmapPatternBindingResource struct {
	client *service.NitroClient
}

func (r *PolicystringmapPatternBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *PolicystringmapPatternBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policystringmap_pattern_binding"
}

func (r *PolicystringmapPatternBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PolicystringmapPatternBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PolicystringmapPatternBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating policystringmap_pattern_binding resource")

	// policystringmap_pattern_binding := policystringmap_pattern_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Policystringmap_pattern_binding.Type(), &policystringmap_pattern_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create policystringmap_pattern_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("policystringmap_pattern_binding-config")

	tflog.Trace(ctx, "Created policystringmap_pattern_binding resource")

	// Read the updated state back
	r.readPolicystringmapPatternBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicystringmapPatternBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data PolicystringmapPatternBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading policystringmap_pattern_binding resource")

	r.readPolicystringmapPatternBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicystringmapPatternBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data PolicystringmapPatternBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating policystringmap_pattern_binding resource")

	// Create API request body from the model
	// policystringmap_pattern_binding := policystringmap_pattern_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Policystringmap_pattern_binding.Type(), &policystringmap_pattern_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update policystringmap_pattern_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated policystringmap_pattern_binding resource")

	// Read the updated state back
	r.readPolicystringmapPatternBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicystringmapPatternBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data PolicystringmapPatternBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting policystringmap_pattern_binding resource")

	// For policystringmap_pattern_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted policystringmap_pattern_binding resource from state")
}

// Helper function to read policystringmap_pattern_binding data from API
func (r *PolicystringmapPatternBindingResource) readPolicystringmapPatternBindingFromApi(ctx context.Context, data *PolicystringmapPatternBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Policystringmap_pattern_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read policystringmap_pattern_binding, got error: %s", err))
		return
	}

	policystringmap_pattern_bindingSetAttrFromGet(ctx, data, getResponseData)

}
