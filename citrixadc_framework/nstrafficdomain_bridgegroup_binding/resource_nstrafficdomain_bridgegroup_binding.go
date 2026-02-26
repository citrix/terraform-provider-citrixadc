package nstrafficdomain_bridgegroup_binding

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
var _ resource.Resource = &NstrafficdomainBridgegroupBindingResource{}
var _ resource.ResourceWithConfigure = (*NstrafficdomainBridgegroupBindingResource)(nil)
var _ resource.ResourceWithImportState = (*NstrafficdomainBridgegroupBindingResource)(nil)

func NewNstrafficdomainBridgegroupBindingResource() resource.Resource {
	return &NstrafficdomainBridgegroupBindingResource{}
}

// NstrafficdomainBridgegroupBindingResource defines the resource implementation.
type NstrafficdomainBridgegroupBindingResource struct {
	client *service.NitroClient
}

func (r *NstrafficdomainBridgegroupBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NstrafficdomainBridgegroupBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nstrafficdomain_bridgegroup_binding"
}

func (r *NstrafficdomainBridgegroupBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NstrafficdomainBridgegroupBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NstrafficdomainBridgegroupBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nstrafficdomain_bridgegroup_binding resource")

	// nstrafficdomain_bridgegroup_binding := nstrafficdomain_bridgegroup_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nstrafficdomain_bridgegroup_binding.Type(), &nstrafficdomain_bridgegroup_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nstrafficdomain_bridgegroup_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nstrafficdomain_bridgegroup_binding-config")

	tflog.Trace(ctx, "Created nstrafficdomain_bridgegroup_binding resource")

	// Read the updated state back
	r.readNstrafficdomainBridgegroupBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstrafficdomainBridgegroupBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NstrafficdomainBridgegroupBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nstrafficdomain_bridgegroup_binding resource")

	r.readNstrafficdomainBridgegroupBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstrafficdomainBridgegroupBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NstrafficdomainBridgegroupBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nstrafficdomain_bridgegroup_binding resource")

	// Create API request body from the model
	// nstrafficdomain_bridgegroup_binding := nstrafficdomain_bridgegroup_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nstrafficdomain_bridgegroup_binding.Type(), &nstrafficdomain_bridgegroup_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nstrafficdomain_bridgegroup_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nstrafficdomain_bridgegroup_binding resource")

	// Read the updated state back
	r.readNstrafficdomainBridgegroupBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstrafficdomainBridgegroupBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NstrafficdomainBridgegroupBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nstrafficdomain_bridgegroup_binding resource")

	// For nstrafficdomain_bridgegroup_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nstrafficdomain_bridgegroup_binding resource from state")
}

// Helper function to read nstrafficdomain_bridgegroup_binding data from API
func (r *NstrafficdomainBridgegroupBindingResource) readNstrafficdomainBridgegroupBindingFromApi(ctx context.Context, data *NstrafficdomainBridgegroupBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nstrafficdomain_bridgegroup_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nstrafficdomain_bridgegroup_binding, got error: %s", err))
		return
	}

	nstrafficdomain_bridgegroup_bindingSetAttrFromGet(ctx, data, getResponseData)

}
