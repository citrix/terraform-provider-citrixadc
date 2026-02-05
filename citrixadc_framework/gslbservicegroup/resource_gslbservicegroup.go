package gslbservicegroup

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
var _ resource.Resource = &GslbservicegroupResource{}
var _ resource.ResourceWithConfigure = (*GslbservicegroupResource)(nil)
var _ resource.ResourceWithImportState = (*GslbservicegroupResource)(nil)

func NewGslbservicegroupResource() resource.Resource {
	return &GslbservicegroupResource{}
}

// GslbservicegroupResource defines the resource implementation.
type GslbservicegroupResource struct {
	client *service.NitroClient
}

func (r *GslbservicegroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbservicegroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbservicegroup"
}

func (r *GslbservicegroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbservicegroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbservicegroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbservicegroup resource")

	// gslbservicegroup := gslbservicegroupGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Gslbservicegroup.Type(), &gslbservicegroup)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create gslbservicegroup, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("gslbservicegroup-config")

	tflog.Trace(ctx, "Created gslbservicegroup resource")

	// Read the updated state back
	r.readGslbservicegroupFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbservicegroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbservicegroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading gslbservicegroup resource")

	r.readGslbservicegroupFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbservicegroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data GslbservicegroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating gslbservicegroup resource")

	// Create API request body from the model
	// gslbservicegroup := gslbservicegroupGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Gslbservicegroup.Type(), &gslbservicegroup)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update gslbservicegroup, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated gslbservicegroup resource")

	// Read the updated state back
	r.readGslbservicegroupFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbservicegroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbservicegroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting gslbservicegroup resource")

	// For gslbservicegroup, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted gslbservicegroup resource from state")
}

// Helper function to read gslbservicegroup data from API
func (r *GslbservicegroupResource) readGslbservicegroupFromApi(ctx context.Context, data *GslbservicegroupResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Gslbservicegroup.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbservicegroup, got error: %s", err))
		return
	}

	gslbservicegroupSetAttrFromGet(ctx, data, getResponseData)

}
