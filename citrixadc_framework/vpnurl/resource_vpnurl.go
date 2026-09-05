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
	var data, config, state VpnurlResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnurl resource")

	// Determine whether there is an actual update and/or attributes to unset.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Actualurl.Equal(state.Actualurl) {
		hasChange = true
	}
	if !data.Linkname.Equal(state.Linkname) {
		hasChange = true
	}
	if !data.Appjson.Equal(state.Appjson) {
		if config.Appjson.IsNull() {
			attributesToUnset = append(attributesToUnset, "appjson")
		} else {
			hasChange = true
		}
	}
	if !data.Applicationtype.Equal(state.Applicationtype) {
		if config.Applicationtype.IsNull() {
			attributesToUnset = append(attributesToUnset, "applicationtype")
		} else {
			hasChange = true
		}
	}
	if !data.Clientlessaccess.Equal(state.Clientlessaccess) {
		if config.Clientlessaccess.IsNull() {
			attributesToUnset = append(attributesToUnset, "clientlessaccess")
		} else {
			hasChange = true
		}
	}
	if !data.Comment.Equal(state.Comment) {
		if config.Comment.IsNull() {
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Iconurl.Equal(state.Iconurl) {
		if config.Iconurl.IsNull() {
			attributesToUnset = append(attributesToUnset, "iconurl")
		} else {
			hasChange = true
		}
	}
	if !data.Ssotype.Equal(state.Ssotype) {
		if config.Ssotype.IsNull() {
			attributesToUnset = append(attributesToUnset, "ssotype")
		} else {
			hasChange = true
		}
	}
	if !data.Vservername.Equal(state.Vservername) {
		if config.Vservername.IsNull() {
			attributesToUnset = append(attributesToUnset, "vservername")
		} else {
			hasChange = true
		}
	}

	urlname := data.Urlname.ValueString()

	if hasChange {
		vpnurl := vpnurlGetThePayloadFromtheConfig(ctx, &data)

		// Named resource - use UpdateResource keyed on urlname
		_, err := r.client.UpdateResource(service.Vpnurl.Type(), urlname, &vpnurl)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnurl, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vpnurl resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnurl resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"urlname": urlname,
	}
	if err := utils.ExecuteUnset(r.client, service.Vpnurl.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset vpnurl attributes, got error: %s", err))
		return
	}

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
