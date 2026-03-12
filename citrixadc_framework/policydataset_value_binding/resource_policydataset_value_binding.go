package policydataset_value_binding

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
var _ resource.Resource = &PolicydatasetValueBindingResource{}
var _ resource.ResourceWithConfigure = (*PolicydatasetValueBindingResource)(nil)
var _ resource.ResourceWithImportState = (*PolicydatasetValueBindingResource)(nil)

func NewPolicydatasetValueBindingResource() resource.Resource {
	return &PolicydatasetValueBindingResource{}
}

// PolicydatasetValueBindingResource defines the resource implementation.
type PolicydatasetValueBindingResource struct {
	client *service.NitroClient
}

func (r *PolicydatasetValueBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *PolicydatasetValueBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policydataset_value_binding"
}

func (r *PolicydatasetValueBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PolicydatasetValueBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PolicydatasetValueBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating policydataset_value_binding resource")

	// policydataset_value_binding := policydataset_value_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Policydataset_value_binding.Type(), &policydataset_value_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create policydataset_value_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("policydataset_value_binding-config")

	tflog.Trace(ctx, "Created policydataset_value_binding resource")

	// Read the updated state back
	r.readPolicydatasetValueBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicydatasetValueBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data PolicydatasetValueBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading policydataset_value_binding resource")

	r.readPolicydatasetValueBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicydatasetValueBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data PolicydatasetValueBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating policydataset_value_binding resource")

	// Create API request body from the model
	// policydataset_value_binding := policydataset_value_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Policydataset_value_binding.Type(), &policydataset_value_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update policydataset_value_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated policydataset_value_binding resource")

	// Read the updated state back
	r.readPolicydatasetValueBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicydatasetValueBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data PolicydatasetValueBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting policydataset_value_binding resource")

	// For policydataset_value_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted policydataset_value_binding resource from state")
}

// Helper function to read policydataset_value_binding data from API
func (r *PolicydatasetValueBindingResource) readPolicydatasetValueBindingFromApi(ctx context.Context, data *PolicydatasetValueBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Policydataset_value_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read policydataset_value_binding, got error: %s", err))
		return
	}

	policydataset_value_bindingSetAttrFromGet(ctx, data, getResponseData)

}
