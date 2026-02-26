package sslservicegroup_ecccurve_binding

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
var _ resource.Resource = &SslservicegroupEcccurveBindingResource{}
var _ resource.ResourceWithConfigure = (*SslservicegroupEcccurveBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslservicegroupEcccurveBindingResource)(nil)

func NewSslservicegroupEcccurveBindingResource() resource.Resource {
	return &SslservicegroupEcccurveBindingResource{}
}

// SslservicegroupEcccurveBindingResource defines the resource implementation.
type SslservicegroupEcccurveBindingResource struct {
	client *service.NitroClient
}

func (r *SslservicegroupEcccurveBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslservicegroupEcccurveBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslservicegroup_ecccurve_binding"
}

func (r *SslservicegroupEcccurveBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslservicegroupEcccurveBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslservicegroupEcccurveBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslservicegroup_ecccurve_binding resource")

	// sslservicegroup_ecccurve_binding := sslservicegroup_ecccurve_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslservicegroup_ecccurve_binding.Type(), &sslservicegroup_ecccurve_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslservicegroup_ecccurve_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("sslservicegroup_ecccurve_binding-config")

	tflog.Trace(ctx, "Created sslservicegroup_ecccurve_binding resource")

	// Read the updated state back
	r.readSslservicegroupEcccurveBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupEcccurveBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslservicegroupEcccurveBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslservicegroup_ecccurve_binding resource")

	r.readSslservicegroupEcccurveBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupEcccurveBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SslservicegroupEcccurveBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating sslservicegroup_ecccurve_binding resource")

	// Create API request body from the model
	// sslservicegroup_ecccurve_binding := sslservicegroup_ecccurve_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslservicegroup_ecccurve_binding.Type(), &sslservicegroup_ecccurve_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslservicegroup_ecccurve_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated sslservicegroup_ecccurve_binding resource")

	// Read the updated state back
	r.readSslservicegroupEcccurveBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupEcccurveBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslservicegroupEcccurveBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslservicegroup_ecccurve_binding resource")

	// For sslservicegroup_ecccurve_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted sslservicegroup_ecccurve_binding resource from state")
}

// Helper function to read sslservicegroup_ecccurve_binding data from API
func (r *SslservicegroupEcccurveBindingResource) readSslservicegroupEcccurveBindingFromApi(ctx context.Context, data *SslservicegroupEcccurveBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Sslservicegroup_ecccurve_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslservicegroup_ecccurve_binding, got error: %s", err))
		return
	}

	sslservicegroup_ecccurve_bindingSetAttrFromGet(ctx, data, getResponseData)

}
