package sslservice_ecccurve_binding

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
var _ resource.Resource = &SslserviceEcccurveBindingResource{}
var _ resource.ResourceWithConfigure = (*SslserviceEcccurveBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslserviceEcccurveBindingResource)(nil)

func NewSslserviceEcccurveBindingResource() resource.Resource {
	return &SslserviceEcccurveBindingResource{}
}

// SslserviceEcccurveBindingResource defines the resource implementation.
type SslserviceEcccurveBindingResource struct {
	client *service.NitroClient
}

func (r *SslserviceEcccurveBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslserviceEcccurveBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslservice_ecccurve_binding"
}

func (r *SslserviceEcccurveBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslserviceEcccurveBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslserviceEcccurveBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslservice_ecccurve_binding resource")

	// sslservice_ecccurve_binding := sslservice_ecccurve_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslservice_ecccurve_binding.Type(), &sslservice_ecccurve_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslservice_ecccurve_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("sslservice_ecccurve_binding-config")

	tflog.Trace(ctx, "Created sslservice_ecccurve_binding resource")

	// Read the updated state back
	r.readSslserviceEcccurveBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceEcccurveBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslserviceEcccurveBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslservice_ecccurve_binding resource")

	r.readSslserviceEcccurveBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceEcccurveBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SslserviceEcccurveBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating sslservice_ecccurve_binding resource")

	// Create API request body from the model
	// sslservice_ecccurve_binding := sslservice_ecccurve_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslservice_ecccurve_binding.Type(), &sslservice_ecccurve_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslservice_ecccurve_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated sslservice_ecccurve_binding resource")

	// Read the updated state back
	r.readSslserviceEcccurveBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceEcccurveBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslserviceEcccurveBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslservice_ecccurve_binding resource")

	// For sslservice_ecccurve_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted sslservice_ecccurve_binding resource from state")
}

// Helper function to read sslservice_ecccurve_binding data from API
func (r *SslserviceEcccurveBindingResource) readSslserviceEcccurveBindingFromApi(ctx context.Context, data *SslserviceEcccurveBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Sslservice_ecccurve_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslservice_ecccurve_binding, got error: %s", err))
		return
	}

	sslservice_ecccurve_bindingSetAttrFromGet(ctx, data, getResponseData)

}
