package lsnhttphdrlogprofile

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
var _ resource.Resource = &LsnhttphdrlogprofileResource{}
var _ resource.ResourceWithConfigure = (*LsnhttphdrlogprofileResource)(nil)
var _ resource.ResourceWithImportState = (*LsnhttphdrlogprofileResource)(nil)

func NewLsnhttphdrlogprofileResource() resource.Resource {
	return &LsnhttphdrlogprofileResource{}
}

// LsnhttphdrlogprofileResource defines the resource implementation.
type LsnhttphdrlogprofileResource struct {
	client *service.NitroClient
}

func (r *LsnhttphdrlogprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnhttphdrlogprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnhttphdrlogprofile"
}

func (r *LsnhttphdrlogprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnhttphdrlogprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnhttphdrlogprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnhttphdrlogprofile resource")

	// lsnhttphdrlogprofile := lsnhttphdrlogprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsnhttphdrlogprofile.Type(), &lsnhttphdrlogprofile)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnhttphdrlogprofile, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lsnhttphdrlogprofile-config")

	tflog.Trace(ctx, "Created lsnhttphdrlogprofile resource")

	// Read the updated state back
	r.readLsnhttphdrlogprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnhttphdrlogprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnhttphdrlogprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnhttphdrlogprofile resource")

	r.readLsnhttphdrlogprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnhttphdrlogprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LsnhttphdrlogprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lsnhttphdrlogprofile resource")

	// Create API request body from the model
	// lsnhttphdrlogprofile := lsnhttphdrlogprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsnhttphdrlogprofile.Type(), &lsnhttphdrlogprofile)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnhttphdrlogprofile, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lsnhttphdrlogprofile resource")

	// Read the updated state back
	r.readLsnhttphdrlogprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnhttphdrlogprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnhttphdrlogprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnhttphdrlogprofile resource")

	// For lsnhttphdrlogprofile, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lsnhttphdrlogprofile resource from state")
}

// Helper function to read lsnhttphdrlogprofile data from API
func (r *LsnhttphdrlogprofileResource) readLsnhttphdrlogprofileFromApi(ctx context.Context, data *LsnhttphdrlogprofileResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lsnhttphdrlogprofile.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnhttphdrlogprofile, got error: %s", err))
		return
	}

	lsnhttphdrlogprofileSetAttrFromGet(ctx, data, getResponseData)

}
