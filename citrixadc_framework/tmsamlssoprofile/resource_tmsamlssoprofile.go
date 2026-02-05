package tmsamlssoprofile

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
var _ resource.Resource = &TmsamlssoprofileResource{}
var _ resource.ResourceWithConfigure = (*TmsamlssoprofileResource)(nil)
var _ resource.ResourceWithImportState = (*TmsamlssoprofileResource)(nil)

func NewTmsamlssoprofileResource() resource.Resource {
	return &TmsamlssoprofileResource{}
}

// TmsamlssoprofileResource defines the resource implementation.
type TmsamlssoprofileResource struct {
	client *service.NitroClient
}

func (r *TmsamlssoprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TmsamlssoprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tmsamlssoprofile"
}

func (r *TmsamlssoprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TmsamlssoprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TmsamlssoprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating tmsamlssoprofile resource")

	// tmsamlssoprofile := tmsamlssoprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Tmsamlssoprofile.Type(), &tmsamlssoprofile)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create tmsamlssoprofile, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("tmsamlssoprofile-config")

	tflog.Trace(ctx, "Created tmsamlssoprofile resource")

	// Read the updated state back
	r.readTmsamlssoprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmsamlssoprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TmsamlssoprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading tmsamlssoprofile resource")

	r.readTmsamlssoprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmsamlssoprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data TmsamlssoprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating tmsamlssoprofile resource")

	// Create API request body from the model
	// tmsamlssoprofile := tmsamlssoprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Tmsamlssoprofile.Type(), &tmsamlssoprofile)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update tmsamlssoprofile, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated tmsamlssoprofile resource")

	// Read the updated state back
	r.readTmsamlssoprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmsamlssoprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TmsamlssoprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting tmsamlssoprofile resource")

	// For tmsamlssoprofile, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted tmsamlssoprofile resource from state")
}

// Helper function to read tmsamlssoprofile data from API
func (r *TmsamlssoprofileResource) readTmsamlssoprofileFromApi(ctx context.Context, data *TmsamlssoprofileResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Tmsamlssoprofile.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read tmsamlssoprofile, got error: %s", err))
		return
	}

	tmsamlssoprofileSetAttrFromGet(ctx, data, getResponseData)

}
