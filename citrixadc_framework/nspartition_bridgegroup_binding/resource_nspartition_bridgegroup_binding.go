package nspartition_bridgegroup_binding

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
var _ resource.Resource = &NspartitionBridgegroupBindingResource{}
var _ resource.ResourceWithConfigure = (*NspartitionBridgegroupBindingResource)(nil)
var _ resource.ResourceWithImportState = (*NspartitionBridgegroupBindingResource)(nil)

func NewNspartitionBridgegroupBindingResource() resource.Resource {
	return &NspartitionBridgegroupBindingResource{}
}

// NspartitionBridgegroupBindingResource defines the resource implementation.
type NspartitionBridgegroupBindingResource struct {
	client *service.NitroClient
}

func (r *NspartitionBridgegroupBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NspartitionBridgegroupBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nspartition_bridgegroup_binding"
}

func (r *NspartitionBridgegroupBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NspartitionBridgegroupBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NspartitionBridgegroupBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nspartition_bridgegroup_binding resource")

	// nspartition_bridgegroup_binding := nspartition_bridgegroup_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nspartition_bridgegroup_binding.Type(), &nspartition_bridgegroup_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nspartition_bridgegroup_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nspartition_bridgegroup_binding-config")

	tflog.Trace(ctx, "Created nspartition_bridgegroup_binding resource")

	// Read the updated state back
	r.readNspartitionBridgegroupBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NspartitionBridgegroupBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NspartitionBridgegroupBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nspartition_bridgegroup_binding resource")

	r.readNspartitionBridgegroupBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NspartitionBridgegroupBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NspartitionBridgegroupBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nspartition_bridgegroup_binding resource")

	// Create API request body from the model
	// nspartition_bridgegroup_binding := nspartition_bridgegroup_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nspartition_bridgegroup_binding.Type(), &nspartition_bridgegroup_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nspartition_bridgegroup_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nspartition_bridgegroup_binding resource")

	// Read the updated state back
	r.readNspartitionBridgegroupBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NspartitionBridgegroupBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NspartitionBridgegroupBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nspartition_bridgegroup_binding resource")

	// For nspartition_bridgegroup_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nspartition_bridgegroup_binding resource from state")
}

// Helper function to read nspartition_bridgegroup_binding data from API
func (r *NspartitionBridgegroupBindingResource) readNspartitionBridgegroupBindingFromApi(ctx context.Context, data *NspartitionBridgegroupBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nspartition_bridgegroup_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nspartition_bridgegroup_binding, got error: %s", err))
		return
	}

	nspartition_bridgegroup_bindingSetAttrFromGet(ctx, data, getResponseData)

}
