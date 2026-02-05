package vpnclientlessaccessprofile

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
var _ resource.Resource = &VpnclientlessaccessprofileResource{}
var _ resource.ResourceWithConfigure = (*VpnclientlessaccessprofileResource)(nil)
var _ resource.ResourceWithImportState = (*VpnclientlessaccessprofileResource)(nil)

func NewVpnclientlessaccessprofileResource() resource.Resource {
	return &VpnclientlessaccessprofileResource{}
}

// VpnclientlessaccessprofileResource defines the resource implementation.
type VpnclientlessaccessprofileResource struct {
	client *service.NitroClient
}

func (r *VpnclientlessaccessprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnclientlessaccessprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnclientlessaccessprofile"
}

func (r *VpnclientlessaccessprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnclientlessaccessprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnclientlessaccessprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnclientlessaccessprofile resource")

	// vpnclientlessaccessprofile := vpnclientlessaccessprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnclientlessaccessprofile.Type(), &vpnclientlessaccessprofile)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnclientlessaccessprofile, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnclientlessaccessprofile-config")

	tflog.Trace(ctx, "Created vpnclientlessaccessprofile resource")

	// Read the updated state back
	r.readVpnclientlessaccessprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnclientlessaccessprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnclientlessaccessprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnclientlessaccessprofile resource")

	r.readVpnclientlessaccessprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnclientlessaccessprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnclientlessaccessprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnclientlessaccessprofile resource")

	// Create API request body from the model
	// vpnclientlessaccessprofile := vpnclientlessaccessprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnclientlessaccessprofile.Type(), &vpnclientlessaccessprofile)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnclientlessaccessprofile, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnclientlessaccessprofile resource")

	// Read the updated state back
	r.readVpnclientlessaccessprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnclientlessaccessprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnclientlessaccessprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnclientlessaccessprofile resource")

	// For vpnclientlessaccessprofile, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnclientlessaccessprofile resource from state")
}

// Helper function to read vpnclientlessaccessprofile data from API
func (r *VpnclientlessaccessprofileResource) readVpnclientlessaccessprofileFromApi(ctx context.Context, data *VpnclientlessaccessprofileResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnclientlessaccessprofile.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnclientlessaccessprofile, got error: %s", err))
		return
	}

	vpnclientlessaccessprofileSetAttrFromGet(ctx, data, getResponseData)

}
