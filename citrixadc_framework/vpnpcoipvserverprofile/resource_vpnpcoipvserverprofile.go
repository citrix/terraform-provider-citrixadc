package vpnpcoipvserverprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
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

	vpnpcoipvserverprofile := vpnpcoipvserverprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	vpnpcoipvserverprofileName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Vpnpcoipvserverprofile.Type(), vpnpcoipvserverprofileName, &vpnpcoipvserverprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnpcoipvserverprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnpcoipvserverprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readVpnpcoipvserverprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnpcoipvserverprofile not found immediately after create")
		}
		return
	}

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

	found := r.readVpnpcoipvserverprofileFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnpcoipvserverprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnpcoipvserverprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnpcoipvserverprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Logindomain.Equal(state.Logindomain) {
		tflog.Debug(ctx, "logindomain has changed for vpnpcoipvserverprofile")
		hasChange = true
	}
	if !data.Udpport.Equal(state.Udpport) {
		tflog.Debug(ctx, "udpport has changed for vpnpcoipvserverprofile")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		vpnpcoipvserverprofile := vpnpcoipvserverprofileGetThePayloadFromtheConfig(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		vpnpcoipvserverprofileName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Vpnpcoipvserverprofile.Type(), vpnpcoipvserverprofileName, &vpnpcoipvserverprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnpcoipvserverprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vpnpcoipvserverprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnpcoipvserverprofile resource, skipping update")
	}

	// Read the updated state back
	if !r.readVpnpcoipvserverprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnpcoipvserverprofile not found immediately after update")
		}
		return
	}

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

	// Named resource - delete using DeleteResource
	vpnpcoipvserverprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Vpnpcoipvserverprofile.Type(), vpnpcoipvserverprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnpcoipvserverprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnpcoipvserverprofile resource")
}

// Helper function to read vpnpcoipvserverprofile data from API
func (r *VpnpcoipvserverprofileResource) readVpnpcoipvserverprofileFromApi(ctx context.Context, data *VpnpcoipvserverprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	vpnpcoipvserverprofileName := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Vpnpcoipvserverprofile.Type(), vpnpcoipvserverprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnpcoipvserverprofile, got error: %s", err))
		return false
	}

	vpnpcoipvserverprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
