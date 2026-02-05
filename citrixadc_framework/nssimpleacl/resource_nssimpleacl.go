package nssimpleacl

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
var _ resource.Resource = &NssimpleaclResource{}
var _ resource.ResourceWithConfigure = (*NssimpleaclResource)(nil)
var _ resource.ResourceWithImportState = (*NssimpleaclResource)(nil)

func NewNssimpleaclResource() resource.Resource {
	return &NssimpleaclResource{}
}

// NssimpleaclResource defines the resource implementation.
type NssimpleaclResource struct {
	client *service.NitroClient
}

func (r *NssimpleaclResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NssimpleaclResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nssimpleacl"
}

func (r *NssimpleaclResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NssimpleaclResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NssimpleaclResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nssimpleacl resource")

	// nssimpleacl := nssimpleaclGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nssimpleacl.Type(), &nssimpleacl)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nssimpleacl, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nssimpleacl-config")

	tflog.Trace(ctx, "Created nssimpleacl resource")

	// Read the updated state back
	r.readNssimpleaclFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NssimpleaclResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NssimpleaclResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nssimpleacl resource")

	r.readNssimpleaclFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NssimpleaclResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NssimpleaclResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nssimpleacl resource")

	// Create API request body from the model
	// nssimpleacl := nssimpleaclGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nssimpleacl.Type(), &nssimpleacl)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nssimpleacl, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nssimpleacl resource")

	// Read the updated state back
	r.readNssimpleaclFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NssimpleaclResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NssimpleaclResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nssimpleacl resource")

	// For nssimpleacl, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nssimpleacl resource from state")
}

// Helper function to read nssimpleacl data from API
func (r *NssimpleaclResource) readNssimpleaclFromApi(ctx context.Context, data *NssimpleaclResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nssimpleacl.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nssimpleacl, got error: %s", err))
		return
	}

	nssimpleaclSetAttrFromGet(ctx, data, getResponseData)

}
