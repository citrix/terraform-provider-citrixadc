package vpnsamlssoprofile

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
var _ resource.Resource = &VpnsamlssoprofileResource{}
var _ resource.ResourceWithConfigure = (*VpnsamlssoprofileResource)(nil)
var _ resource.ResourceWithImportState = (*VpnsamlssoprofileResource)(nil)

func NewVpnsamlssoprofileResource() resource.Resource {
	return &VpnsamlssoprofileResource{}
}

// VpnsamlssoprofileResource defines the resource implementation.
type VpnsamlssoprofileResource struct {
	client *service.NitroClient
}

func (r *VpnsamlssoprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnsamlssoprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnsamlssoprofile"
}

func (r *VpnsamlssoprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnsamlssoprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnsamlssoprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnsamlssoprofile resource")

	// Create API request body from the plan
	vpnsamlssoprofile := vpnsamlssoprofileGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	vpnsamlssoprofileName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Vpnsamlssoprofile.Type(), vpnsamlssoprofileName, &vpnsamlssoprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnsamlssoprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnsamlssoprofile resource")

	// Set ID for the resource before reading state back
	data.Id = types.StringValue(vpnsamlssoprofileName)

	// Read the updated state back
	if !r.readVpnsamlssoprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnsamlssoprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnsamlssoprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnsamlssoprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnsamlssoprofile resource")

	found := r.readVpnsamlssoprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnsamlssoprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnsamlssoprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnsamlssoprofile resource")

	// name is ForceNew/RequiresReplace, so any Update invocation is a change to an
	// updateable attribute. Build the payload from the plan and push it.
	vpnsamlssoprofile := vpnsamlssoprofileGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use UpdateResource
	vpnsamlssoprofileName := data.Name.ValueString()
	_, err := r.client.UpdateResource(service.Vpnsamlssoprofile.Type(), vpnsamlssoprofileName, &vpnsamlssoprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnsamlssoprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated vpnsamlssoprofile resource")

	// Read the updated state back
	if !r.readVpnsamlssoprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnsamlssoprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnsamlssoprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnsamlssoprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnsamlssoprofile resource")

	// Named resource - delete using DeleteResource
	vpnsamlssoprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Vpnsamlssoprofile.Type(), vpnsamlssoprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnsamlssoprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnsamlssoprofile resource")
}

// Helper function to read vpnsamlssoprofile data from API
func (r *VpnsamlssoprofileResource) readVpnsamlssoprofileFromApi(ctx context.Context, data *VpnsamlssoprofileResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain name value
	vpnsamlssoprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Vpnsamlssoprofile.Type(), vpnsamlssoprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnsamlssoprofile, got error: %s", err))
		return false
	}

	vpnsamlssoprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
