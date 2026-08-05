package vpnurl

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
var _ resource.Resource = &VpnurlResource{}
var _ resource.ResourceWithConfigure = (*VpnurlResource)(nil)
var _ resource.ResourceWithImportState = (*VpnurlResource)(nil)

func NewVpnurlResource() resource.Resource {
	return &VpnurlResource{}
}

// VpnurlResource defines the resource implementation.
type VpnurlResource struct {
	client *service.NitroClient
}

func (r *VpnurlResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnurlResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnurl"
}

func (r *VpnurlResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnurlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnurlResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnurl resource")

	vpnurl := vpnurlGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource keyed on urlname
	urlname := data.Urlname.ValueString()
	_, err := r.client.AddResource(service.Vpnurl.Type(), urlname, &vpnurl)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnurl, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnurl resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(urlname)

	// Read the updated state back
	if !r.readVpnurlFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnurl not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnurlResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnurlResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnurl resource")

	found := r.readVpnurlFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnurlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnurlResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnurl resource")

	vpnurl := vpnurlGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use UpdateResource keyed on urlname
	urlname := data.Urlname.ValueString()
	_, err := r.client.UpdateResource(service.Vpnurl.Type(), urlname, &vpnurl)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnurl, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated vpnurl resource")

	// Read the updated state back
	if !r.readVpnurlFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnurl not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnurlResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnurlResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnurl resource")

	// Named resource - delete using DeleteResource keyed on urlname
	urlname := data.Urlname.ValueString()
	err := r.client.DeleteResource(service.Vpnurl.Type(), urlname)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnurl, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnurl resource")
}

// Helper function to read vpnurl data from API. Returns false when the
// resource no longer exists on the ADC (so callers can drop it from state).
func (r *VpnurlResource) readVpnurlFromApi(ctx context.Context, data *VpnurlResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain urlname value
	urlname := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Vpnurl.Type(), urlname)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnurl, got error: %s", err))
		return false
	}

	vpnurlSetAttrFromGet(ctx, data, getResponseData)

	return true
}
