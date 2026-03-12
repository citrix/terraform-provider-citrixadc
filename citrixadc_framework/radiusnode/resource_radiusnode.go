package radiusnode

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
var _ resource.Resource = &RadiusnodeResource{}
var _ resource.ResourceWithConfigure = (*RadiusnodeResource)(nil)
var _ resource.ResourceWithImportState = (*RadiusnodeResource)(nil)

func NewRadiusnodeResource() resource.Resource {
	return &RadiusnodeResource{}
}

// RadiusnodeResource defines the resource implementation.
type RadiusnodeResource struct {
	client *service.NitroClient
}

func (r *RadiusnodeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RadiusnodeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_radiusnode"
}

func (r *RadiusnodeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RadiusnodeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RadiusnodeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating radiusnode resource")

	// radiusnode := radiusnodeGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Radiusnode.Type(), &radiusnode)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create radiusnode, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("radiusnode-config")

	tflog.Trace(ctx, "Created radiusnode resource")

	// Read the updated state back
	r.readRadiusnodeFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RadiusnodeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RadiusnodeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading radiusnode resource")

	r.readRadiusnodeFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RadiusnodeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data RadiusnodeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating radiusnode resource")

	// Create API request body from the model
	// radiusnode := radiusnodeGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Radiusnode.Type(), &radiusnode)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update radiusnode, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated radiusnode resource")

	// Read the updated state back
	r.readRadiusnodeFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RadiusnodeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RadiusnodeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting radiusnode resource")

	// For radiusnode, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted radiusnode resource from state")
}

// Helper function to read radiusnode data from API
func (r *RadiusnodeResource) readRadiusnodeFromApi(ctx context.Context, data *RadiusnodeResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Radiusnode.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read radiusnode, got error: %s", err))
		return
	}

	radiusnodeSetAttrFromGet(ctx, data, getResponseData)

}
