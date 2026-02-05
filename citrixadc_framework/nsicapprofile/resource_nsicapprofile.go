package nsicapprofile

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
var _ resource.Resource = &NsicapprofileResource{}
var _ resource.ResourceWithConfigure = (*NsicapprofileResource)(nil)
var _ resource.ResourceWithImportState = (*NsicapprofileResource)(nil)

func NewNsicapprofileResource() resource.Resource {
	return &NsicapprofileResource{}
}

// NsicapprofileResource defines the resource implementation.
type NsicapprofileResource struct {
	client *service.NitroClient
}

func (r *NsicapprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsicapprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsicapprofile"
}

func (r *NsicapprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsicapprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsicapprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsicapprofile resource")

	// nsicapprofile := nsicapprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nsicapprofile.Type(), &nsicapprofile)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsicapprofile, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nsicapprofile-config")

	tflog.Trace(ctx, "Created nsicapprofile resource")

	// Read the updated state back
	r.readNsicapprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsicapprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsicapprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsicapprofile resource")

	r.readNsicapprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsicapprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NsicapprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nsicapprofile resource")

	// Create API request body from the model
	// nsicapprofile := nsicapprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nsicapprofile.Type(), &nsicapprofile)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsicapprofile, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nsicapprofile resource")

	// Read the updated state back
	r.readNsicapprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsicapprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsicapprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsicapprofile resource")

	// For nsicapprofile, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nsicapprofile resource from state")
}

// Helper function to read nsicapprofile data from API
func (r *NsicapprofileResource) readNsicapprofileFromApi(ctx context.Context, data *NsicapprofileResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nsicapprofile.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsicapprofile, got error: %s", err))
		return
	}

	nsicapprofileSetAttrFromGet(ctx, data, getResponseData)

}
