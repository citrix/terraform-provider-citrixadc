package bridgegroup

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
var _ resource.Resource = &BridgegroupResource{}
var _ resource.ResourceWithConfigure = (*BridgegroupResource)(nil)
var _ resource.ResourceWithImportState = (*BridgegroupResource)(nil)

func NewBridgegroupResource() resource.Resource {
	return &BridgegroupResource{}
}

// BridgegroupResource defines the resource implementation.
type BridgegroupResource struct {
	client *service.NitroClient
}

func (r *BridgegroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BridgegroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bridgegroup"
}

func (r *BridgegroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BridgegroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BridgegroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating bridgegroup resource")

	// bridgegroup := bridgegroupGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Bridgegroup.Type(), &bridgegroup)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create bridgegroup, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("bridgegroup-config")

	tflog.Trace(ctx, "Created bridgegroup resource")

	// Read the updated state back
	r.readBridgegroupFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BridgegroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BridgegroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading bridgegroup resource")

	r.readBridgegroupFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BridgegroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data BridgegroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating bridgegroup resource")

	// Create API request body from the model
	// bridgegroup := bridgegroupGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Bridgegroup.Type(), &bridgegroup)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update bridgegroup, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated bridgegroup resource")

	// Read the updated state back
	r.readBridgegroupFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BridgegroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BridgegroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting bridgegroup resource")

	// For bridgegroup, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted bridgegroup resource from state")
}

// Helper function to read bridgegroup data from API
func (r *BridgegroupResource) readBridgegroupFromApi(ctx context.Context, data *BridgegroupResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Bridgegroup.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read bridgegroup, got error: %s", err))
		return
	}

	bridgegroupSetAttrFromGet(ctx, data, getResponseData)

}
