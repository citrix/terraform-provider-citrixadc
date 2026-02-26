package nsservicepath_nsservicefunction_binding

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
var _ resource.Resource = &NsservicepathNsservicefunctionBindingResource{}
var _ resource.ResourceWithConfigure = (*NsservicepathNsservicefunctionBindingResource)(nil)
var _ resource.ResourceWithImportState = (*NsservicepathNsservicefunctionBindingResource)(nil)

func NewNsservicepathNsservicefunctionBindingResource() resource.Resource {
	return &NsservicepathNsservicefunctionBindingResource{}
}

// NsservicepathNsservicefunctionBindingResource defines the resource implementation.
type NsservicepathNsservicefunctionBindingResource struct {
	client *service.NitroClient
}

func (r *NsservicepathNsservicefunctionBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsservicepathNsservicefunctionBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsservicepath_nsservicefunction_binding"
}

func (r *NsservicepathNsservicefunctionBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsservicepathNsservicefunctionBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsservicepathNsservicefunctionBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsservicepath_nsservicefunction_binding resource")

	// nsservicepath_nsservicefunction_binding := nsservicepath_nsservicefunction_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nsservicepath_nsservicefunction_binding.Type(), &nsservicepath_nsservicefunction_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsservicepath_nsservicefunction_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nsservicepath_nsservicefunction_binding-config")

	tflog.Trace(ctx, "Created nsservicepath_nsservicefunction_binding resource")

	// Read the updated state back
	r.readNsservicepathNsservicefunctionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsservicepathNsservicefunctionBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsservicepathNsservicefunctionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsservicepath_nsservicefunction_binding resource")

	r.readNsservicepathNsservicefunctionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsservicepathNsservicefunctionBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NsservicepathNsservicefunctionBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nsservicepath_nsservicefunction_binding resource")

	// Create API request body from the model
	// nsservicepath_nsservicefunction_binding := nsservicepath_nsservicefunction_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nsservicepath_nsservicefunction_binding.Type(), &nsservicepath_nsservicefunction_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsservicepath_nsservicefunction_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nsservicepath_nsservicefunction_binding resource")

	// Read the updated state back
	r.readNsservicepathNsservicefunctionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsservicepathNsservicefunctionBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsservicepathNsservicefunctionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsservicepath_nsservicefunction_binding resource")

	// For nsservicepath_nsservicefunction_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nsservicepath_nsservicefunction_binding resource from state")
}

// Helper function to read nsservicepath_nsservicefunction_binding data from API
func (r *NsservicepathNsservicefunctionBindingResource) readNsservicepathNsservicefunctionBindingFromApi(ctx context.Context, data *NsservicepathNsservicefunctionBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nsservicepath_nsservicefunction_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsservicepath_nsservicefunction_binding, got error: %s", err))
		return
	}

	nsservicepath_nsservicefunction_bindingSetAttrFromGet(ctx, data, getResponseData)

}
