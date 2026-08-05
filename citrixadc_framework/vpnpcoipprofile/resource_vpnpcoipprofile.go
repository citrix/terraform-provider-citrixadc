package vpnpcoipprofile

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
var _ resource.Resource = &VpnpcoipprofileResource{}
var _ resource.ResourceWithConfigure = (*VpnpcoipprofileResource)(nil)
var _ resource.ResourceWithImportState = (*VpnpcoipprofileResource)(nil)

func NewVpnpcoipprofileResource() resource.Resource {
	return &VpnpcoipprofileResource{}
}

// VpnpcoipprofileResource defines the resource implementation.
type VpnpcoipprofileResource struct {
	client *service.NitroClient
}

func (r *VpnpcoipprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnpcoipprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnpcoipprofile"
}

func (r *VpnpcoipprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnpcoipprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnpcoipprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnpcoipprofile resource")

	// Create API request body from the model
	vpnpcoipprofile := vpnpcoipprofileGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Vpnpcoipprofile.Type(), name_value, &vpnpcoipprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnpcoipprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnpcoipprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(name_value)

	// Read the updated state back
	if !r.readVpnpcoipprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnpcoipprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnpcoipprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnpcoipprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnpcoipprofile resource")

	found := r.readVpnpcoipprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnpcoipprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnpcoipprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnpcoipprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Conserverurl.Equal(state.Conserverurl) {
		tflog.Debug(ctx, "conserverurl has changed for vpnpcoipprofile")
		hasChange = true
	}
	if !data.Icvverification.Equal(state.Icvverification) {
		tflog.Debug(ctx, "icvverification has changed for vpnpcoipprofile")
		hasChange = true
	}
	if !data.Sessionidletimeout.Equal(state.Sessionidletimeout) {
		tflog.Debug(ctx, "sessionidletimeout has changed for vpnpcoipprofile")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		vpnpcoipprofile := vpnpcoipprofileGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Vpnpcoipprofile.Type(), name_value, &vpnpcoipprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnpcoipprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vpnpcoipprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnpcoipprofile resource, skipping update")
	}

	// Read the updated state back
	if !r.readVpnpcoipprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnpcoipprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnpcoipprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnpcoipprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnpcoipprofile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Vpnpcoipprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnpcoipprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnpcoipprofile resource")
}

// Helper function to read vpnpcoipprofile data from API
func (r *VpnpcoipprofileResource) readVpnpcoipprofileFromApi(ctx context.Context, data *VpnpcoipprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Vpnpcoipprofile.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnpcoipprofile, got error: %s", err))
		return false
	}

	vpnpcoipprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
