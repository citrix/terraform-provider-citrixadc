package lsnrtspalgprofile

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
var _ resource.Resource = &LsnrtspalgprofileResource{}
var _ resource.ResourceWithConfigure = (*LsnrtspalgprofileResource)(nil)
var _ resource.ResourceWithImportState = (*LsnrtspalgprofileResource)(nil)

func NewLsnrtspalgprofileResource() resource.Resource {
	return &LsnrtspalgprofileResource{}
}

// LsnrtspalgprofileResource defines the resource implementation.
type LsnrtspalgprofileResource struct {
	client *service.NitroClient
}

func (r *LsnrtspalgprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnrtspalgprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnrtspalgprofile"
}

func (r *LsnrtspalgprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnrtspalgprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnrtspalgprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnrtspalgprofile resource")

	// lsnrtspalgprofile := lsnrtspalgprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsnrtspalgprofile.Type(), &lsnrtspalgprofile)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnrtspalgprofile, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lsnrtspalgprofile-config")

	tflog.Trace(ctx, "Created lsnrtspalgprofile resource")

	// Read the updated state back
	r.readLsnrtspalgprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnrtspalgprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnrtspalgprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnrtspalgprofile resource")

	r.readLsnrtspalgprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnrtspalgprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LsnrtspalgprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lsnrtspalgprofile resource")

	// Create API request body from the model
	// lsnrtspalgprofile := lsnrtspalgprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsnrtspalgprofile.Type(), &lsnrtspalgprofile)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnrtspalgprofile, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lsnrtspalgprofile resource")

	// Read the updated state back
	r.readLsnrtspalgprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnrtspalgprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnrtspalgprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnrtspalgprofile resource")

	// For lsnrtspalgprofile, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lsnrtspalgprofile resource from state")
}

// Helper function to read lsnrtspalgprofile data from API
func (r *LsnrtspalgprofileResource) readLsnrtspalgprofileFromApi(ctx context.Context, data *LsnrtspalgprofileResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lsnrtspalgprofile.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnrtspalgprofile, got error: %s", err))
		return
	}

	lsnrtspalgprofileSetAttrFromGet(ctx, data, getResponseData)

}
