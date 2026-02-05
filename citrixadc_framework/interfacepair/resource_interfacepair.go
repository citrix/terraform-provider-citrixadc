package interfacepair

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
var _ resource.Resource = &InterfacepairResource{}
var _ resource.ResourceWithConfigure = (*InterfacepairResource)(nil)
var _ resource.ResourceWithImportState = (*InterfacepairResource)(nil)

func NewInterfacepairResource() resource.Resource {
	return &InterfacepairResource{}
}

// InterfacepairResource defines the resource implementation.
type InterfacepairResource struct {
	client *service.NitroClient
}

func (r *InterfacepairResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *InterfacepairResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interfacepair"
}

func (r *InterfacepairResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *InterfacepairResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data InterfacepairResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating interfacepair resource")

	// interfacepair := interfacepairGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Interfacepair.Type(), &interfacepair)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create interfacepair, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("interfacepair-config")

	tflog.Trace(ctx, "Created interfacepair resource")

	// Read the updated state back
	r.readInterfacepairFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InterfacepairResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data InterfacepairResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading interfacepair resource")

	r.readInterfacepairFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InterfacepairResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data InterfacepairResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating interfacepair resource")

	// Create API request body from the model
	// interfacepair := interfacepairGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Interfacepair.Type(), &interfacepair)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update interfacepair, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated interfacepair resource")

	// Read the updated state back
	r.readInterfacepairFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InterfacepairResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data InterfacepairResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting interfacepair resource")

	// For interfacepair, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted interfacepair resource from state")
}

// Helper function to read interfacepair data from API
func (r *InterfacepairResource) readInterfacepairFromApi(ctx context.Context, data *InterfacepairResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Interfacepair.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read interfacepair, got error: %s", err))
		return
	}

	interfacepairSetAttrFromGet(ctx, data, getResponseData)

}
