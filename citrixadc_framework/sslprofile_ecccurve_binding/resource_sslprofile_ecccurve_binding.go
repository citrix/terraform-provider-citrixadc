package sslprofile_ecccurve_binding

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
var _ resource.Resource = &SslprofileEcccurveBindingResource{}
var _ resource.ResourceWithConfigure = (*SslprofileEcccurveBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslprofileEcccurveBindingResource)(nil)

func NewSslprofileEcccurveBindingResource() resource.Resource {
	return &SslprofileEcccurveBindingResource{}
}

// SslprofileEcccurveBindingResource defines the resource implementation.
type SslprofileEcccurveBindingResource struct {
	client *service.NitroClient
}

func (r *SslprofileEcccurveBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslprofileEcccurveBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslprofile_ecccurve_binding"
}

func (r *SslprofileEcccurveBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslprofileEcccurveBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslprofileEcccurveBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslprofile_ecccurve_binding resource")

	// sslprofile_ecccurve_binding := sslprofile_ecccurve_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslprofile_ecccurve_binding.Type(), &sslprofile_ecccurve_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslprofile_ecccurve_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("sslprofile_ecccurve_binding-config")

	tflog.Trace(ctx, "Created sslprofile_ecccurve_binding resource")

	// Read the updated state back
	r.readSslprofileEcccurveBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileEcccurveBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslprofileEcccurveBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslprofile_ecccurve_binding resource")

	r.readSslprofileEcccurveBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileEcccurveBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SslprofileEcccurveBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating sslprofile_ecccurve_binding resource")

	// Create API request body from the model
	// sslprofile_ecccurve_binding := sslprofile_ecccurve_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Sslprofile_ecccurve_binding.Type(), &sslprofile_ecccurve_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslprofile_ecccurve_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated sslprofile_ecccurve_binding resource")

	// Read the updated state back
	r.readSslprofileEcccurveBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileEcccurveBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslprofileEcccurveBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslprofile_ecccurve_binding resource")

	// For sslprofile_ecccurve_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted sslprofile_ecccurve_binding resource from state")
}

// Helper function to read sslprofile_ecccurve_binding data from API
func (r *SslprofileEcccurveBindingResource) readSslprofileEcccurveBindingFromApi(ctx context.Context, data *SslprofileEcccurveBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Sslprofile_ecccurve_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslprofile_ecccurve_binding, got error: %s", err))
		return
	}

	sslprofile_ecccurve_bindingSetAttrFromGet(ctx, data, getResponseData)

}
