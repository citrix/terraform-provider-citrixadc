package vpnvserver

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VpnvserverResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverResource)(nil)

func NewVpnvserverResource() resource.Resource {
	return &VpnvserverResource{}
}

// VpnvserverResource defines the resource implementation.
type VpnvserverResource struct {
	client *service.NitroClient
}

func (r *VpnvserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver"
}

func (r *VpnvserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver resource")

	// Create API request body from the model
	vpnvserver := vpnvserverGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	vpnvserverName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Vpnvserver.Type(), vpnvserverName, &vpnvserver)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnvserver resource")

	// Set ID (the live vserver name) before reading state back
	data.Id = types.StringValue(vpnvserverName)

	// Read the updated state back
	if !r.readVpnvserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnvserver not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver resource")

	found := r.readVpnvserverFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnvserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnvserverResourceModel

	// Read Terraform prior state to preserve the live name/ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID (live name) from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnvserver resource")

	// Handle in-place rename (NITRO ?action=rename) when newname changes.
	if !data.Newname.IsNull() && !data.Newname.IsUnknown() && data.Newname.ValueString() != "" &&
		!data.Newname.Equal(state.Newname) {
		newName := data.Newname.ValueString()
		// The rename source must be the CURRENT LIVE name, tracked in state.Id.
		renamePayload := vpn.Vpnvserver{
			Name:    state.Id.ValueString(),
			Newname: newName,
		}
		tflog.Debug(ctx, fmt.Sprintf("Renaming vpnvserver %s to %s", state.Id.ValueString(), newName))
		err := r.client.ActOnResource(service.Vpnvserver.Type(), &renamePayload, "rename")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename vpnvserver, got error: %s", err))
			return
		}
		// The ID now tracks the live (renamed) object.
		data.Id = types.StringValue(newName)
	}

	// Build a set payload containing ONLY the changed updateable attributes,
	// mirroring the SDK v2 update contract. name/servicetype are create-only
	// (RequiresReplace) and never reach the set payload — NITRO rejects
	// servicetype on set (errorcode 278); newname is handled above via rename.
	vpnvserver, hasChange := vpnvserverGetTheUpdatablePayloadFromThePlan(ctx, &data, &state)

	if hasChange {
		// Key the update on the live name (data.Id), which reflects any rename above.
		vpnvserver.Name = data.Id.ValueString()
		_, err := r.client.UpdateResource(service.Vpnvserver.Type(), data.Id.ValueString(), &vpnvserver)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated vpnvserver resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnvserver resource, skipping update")
	}

	// Read the updated state back
	if !r.readVpnvserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnvserver not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver resource")

	// Named resource - delete by live name (data.Id), robust to a prior rename.
	err := r.client.DeleteResource(service.Vpnvserver.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnvserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnvserver resource")
}

// Helper function to read vpnvserver data from API. Returns false if not found.
func (r *VpnvserverResource) readVpnvserverFromApi(ctx context.Context, data *VpnvserverResourceModel, diags *diag.Diagnostics) bool {
	vpnvserverName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Vpnvserver.Type(), vpnvserverName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver, got error: %s", err))
		return false
	}

	vpnvserverSetAttrFromGet(ctx, data, getResponseData)

	return true
}
