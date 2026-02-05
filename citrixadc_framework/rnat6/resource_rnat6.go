package rnat6

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
var _ resource.Resource = &Rnat6Resource{}
var _ resource.ResourceWithConfigure = (*Rnat6Resource)(nil)
var _ resource.ResourceWithImportState = (*Rnat6Resource)(nil)

func NewRnat6Resource() resource.Resource {
	return &Rnat6Resource{}
}

// Rnat6Resource defines the resource implementation.
type Rnat6Resource struct {
	client *service.NitroClient
}

func (r *Rnat6Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Rnat6Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rnat6"
}

func (r *Rnat6Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Rnat6Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Rnat6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rnat6 resource")

	// rnat6 := rnat6GetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Rnat6.Type(), &rnat6)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create rnat6, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("rnat6-config")

	tflog.Trace(ctx, "Created rnat6 resource")

	// Read the updated state back
	r.readRnat6FromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Rnat6Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Rnat6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading rnat6 resource")

	r.readRnat6FromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Rnat6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data Rnat6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating rnat6 resource")

	// Create API request body from the model
	// rnat6 := rnat6GetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Rnat6.Type(), &rnat6)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update rnat6, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated rnat6 resource")

	// Read the updated state back
	r.readRnat6FromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Rnat6Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Rnat6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting rnat6 resource")

	// For rnat6, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted rnat6 resource from state")
}

// Helper function to read rnat6 data from API
func (r *Rnat6Resource) readRnat6FromApi(ctx context.Context, data *Rnat6ResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Rnat6.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read rnat6, got error: %s", err))
		return
	}

	rnat6SetAttrFromGet(ctx, data, getResponseData)

}
