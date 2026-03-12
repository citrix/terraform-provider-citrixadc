package nsxmlnamespace

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
var _ resource.Resource = &NsxmlnamespaceResource{}
var _ resource.ResourceWithConfigure = (*NsxmlnamespaceResource)(nil)
var _ resource.ResourceWithImportState = (*NsxmlnamespaceResource)(nil)

func NewNsxmlnamespaceResource() resource.Resource {
	return &NsxmlnamespaceResource{}
}

// NsxmlnamespaceResource defines the resource implementation.
type NsxmlnamespaceResource struct {
	client *service.NitroClient
}

func (r *NsxmlnamespaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsxmlnamespaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsxmlnamespace"
}

func (r *NsxmlnamespaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsxmlnamespaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsxmlnamespaceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsxmlnamespace resource")

	// nsxmlnamespace := nsxmlnamespaceGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nsxmlnamespace.Type(), &nsxmlnamespace)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsxmlnamespace, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nsxmlnamespace-config")

	tflog.Trace(ctx, "Created nsxmlnamespace resource")

	// Read the updated state back
	r.readNsxmlnamespaceFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsxmlnamespaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsxmlnamespaceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsxmlnamespace resource")

	r.readNsxmlnamespaceFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsxmlnamespaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NsxmlnamespaceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nsxmlnamespace resource")

	// Create API request body from the model
	// nsxmlnamespace := nsxmlnamespaceGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nsxmlnamespace.Type(), &nsxmlnamespace)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsxmlnamespace, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nsxmlnamespace resource")

	// Read the updated state back
	r.readNsxmlnamespaceFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsxmlnamespaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsxmlnamespaceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsxmlnamespace resource")

	// For nsxmlnamespace, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nsxmlnamespace resource from state")
}

// Helper function to read nsxmlnamespace data from API
func (r *NsxmlnamespaceResource) readNsxmlnamespaceFromApi(ctx context.Context, data *NsxmlnamespaceResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nsxmlnamespace.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsxmlnamespace, got error: %s", err))
		return
	}

	nsxmlnamespaceSetAttrFromGet(ctx, data, getResponseData)

}
