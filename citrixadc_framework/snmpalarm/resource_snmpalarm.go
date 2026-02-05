package snmpalarm

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
var _ resource.Resource = &SnmpalarmResource{}
var _ resource.ResourceWithConfigure = (*SnmpalarmResource)(nil)
var _ resource.ResourceWithImportState = (*SnmpalarmResource)(nil)

func NewSnmpalarmResource() resource.Resource {
	return &SnmpalarmResource{}
}

// SnmpalarmResource defines the resource implementation.
type SnmpalarmResource struct {
	client *service.NitroClient
}

func (r *SnmpalarmResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SnmpalarmResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmpalarm"
}

func (r *SnmpalarmResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SnmpalarmResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SnmpalarmResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating snmpalarm resource")

	// snmpalarm := snmpalarmGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Snmpalarm.Type(), &snmpalarm)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create snmpalarm, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("snmpalarm-config")

	tflog.Trace(ctx, "Created snmpalarm resource")

	// Read the updated state back
	r.readSnmpalarmFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpalarmResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SnmpalarmResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading snmpalarm resource")

	r.readSnmpalarmFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpalarmResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SnmpalarmResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating snmpalarm resource")

	// Create API request body from the model
	// snmpalarm := snmpalarmGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Snmpalarm.Type(), &snmpalarm)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update snmpalarm, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated snmpalarm resource")

	// Read the updated state back
	r.readSnmpalarmFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpalarmResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SnmpalarmResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting snmpalarm resource")

	// For snmpalarm, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted snmpalarm resource from state")
}

// Helper function to read snmpalarm data from API
func (r *SnmpalarmResource) readSnmpalarmFromApi(ctx context.Context, data *SnmpalarmResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Snmpalarm.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read snmpalarm, got error: %s", err))
		return
	}

	snmpalarmSetAttrFromGet(ctx, data, getResponseData)

}
