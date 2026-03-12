package vpnpcoipvserverprofile

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
var _ resource.Resource = &VpnpcoipvserverprofileResource{}
var _ resource.ResourceWithConfigure = (*VpnpcoipvserverprofileResource)(nil)
var _ resource.ResourceWithImportState = (*VpnpcoipvserverprofileResource)(nil)

func NewVpnpcoipvserverprofileResource() resource.Resource {
	return &VpnpcoipvserverprofileResource{}
}

// VpnpcoipvserverprofileResource defines the resource implementation.
type VpnpcoipvserverprofileResource struct {
	client *service.NitroClient
}

func (r *VpnpcoipvserverprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnpcoipvserverprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnpcoipvserverprofile"
}

func (r *VpnpcoipvserverprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnpcoipvserverprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnpcoipvserverprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnpcoipvserverprofile resource")

	// vpnpcoipvserverprofile := vpnpcoipvserverprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnpcoipvserverprofile.Type(), &vpnpcoipvserverprofile)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnpcoipvserverprofile, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnpcoipvserverprofile-config")

	tflog.Trace(ctx, "Created vpnpcoipvserverprofile resource")

	// Read the updated state back
	r.readVpnpcoipvserverprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnpcoipvserverprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnpcoipvserverprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnpcoipvserverprofile resource")

	r.readVpnpcoipvserverprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnpcoipvserverprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnpcoipvserverprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnpcoipvserverprofile resource")

	// Create API request body from the model
	// vpnpcoipvserverprofile := vpnpcoipvserverprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnpcoipvserverprofile.Type(), &vpnpcoipvserverprofile)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnpcoipvserverprofile, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnpcoipvserverprofile resource")

	// Read the updated state back
	r.readVpnpcoipvserverprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnpcoipvserverprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnpcoipvserverprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnpcoipvserverprofile resource")

	// For vpnpcoipvserverprofile, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnpcoipvserverprofile resource from state")
}

// Helper function to read vpnpcoipvserverprofile data from API
func (r *VpnpcoipvserverprofileResource) readVpnpcoipvserverprofileFromApi(ctx context.Context, data *VpnpcoipvserverprofileResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnpcoipvserverprofile.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnpcoipvserverprofile, got error: %s", err))
		return
	}

	vpnpcoipvserverprofileSetAttrFromGet(ctx, data, getResponseData)

}
