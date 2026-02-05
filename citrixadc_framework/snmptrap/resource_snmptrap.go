package snmptrap

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
var _ resource.Resource = &SnmptrapResource{}
var _ resource.ResourceWithConfigure = (*SnmptrapResource)(nil)
var _ resource.ResourceWithImportState = (*SnmptrapResource)(nil)

func NewSnmptrapResource() resource.Resource {
	return &SnmptrapResource{}
}

// SnmptrapResource defines the resource implementation.
type SnmptrapResource struct {
	client *service.NitroClient
}

func (r *SnmptrapResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SnmptrapResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmptrap"
}

func (r *SnmptrapResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SnmptrapResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SnmptrapResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating snmptrap resource")

	// snmptrap := snmptrapGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Snmptrap.Type(), &snmptrap)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create snmptrap, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("snmptrap-config")

	tflog.Trace(ctx, "Created snmptrap resource")

	// Read the updated state back
	r.readSnmptrapFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmptrapResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SnmptrapResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading snmptrap resource")

	r.readSnmptrapFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmptrapResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SnmptrapResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating snmptrap resource")

	// Create API request body from the model
	// snmptrap := snmptrapGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Snmptrap.Type(), &snmptrap)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update snmptrap, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated snmptrap resource")

	// Read the updated state back
	r.readSnmptrapFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmptrapResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SnmptrapResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting snmptrap resource")

	// For snmptrap, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted snmptrap resource from state")
}

// Helper function to read snmptrap data from API
func (r *SnmptrapResource) readSnmptrapFromApi(ctx context.Context, data *SnmptrapResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Snmptrap.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read snmptrap, got error: %s", err))
		return
	}

	snmptrapSetAttrFromGet(ctx, data, getResponseData)

}
