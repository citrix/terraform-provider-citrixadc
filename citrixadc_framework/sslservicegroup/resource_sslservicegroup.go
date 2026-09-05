package sslservicegroup

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
var _ resource.Resource = &SslservicegroupResource{}
var _ resource.ResourceWithConfigure = (*SslservicegroupResource)(nil)
var _ resource.ResourceWithImportState = (*SslservicegroupResource)(nil)

func NewSslservicegroupResource() resource.Resource {
	return &SslservicegroupResource{}
}

// SslservicegroupResource defines the resource implementation.
type SslservicegroupResource struct {
	client *service.NitroClient
}

func (r *SslservicegroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslservicegroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslservicegroup"
}

func (r *SslservicegroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslservicegroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslservicegroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslservicegroup resource")

	// sslservicegroup does not have an ADD operation; its SSL configuration is
	// applied to an already-existing service group via an UPDATE (matches SDK v2).
	sslservicegroupName := data.Servicegroupname.ValueString()
	sslservicegroup := sslservicegroupGetThePayloadFromthePlan(ctx, &data)

	_, err := r.client.UpdateResource(service.Sslservicegroup.Type(), sslservicegroupName, &sslservicegroup)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslservicegroup, got error: %s", err))
		return
	}

	// ID is the service group name (single unique attribute), matching SDK v2.
	data.Id = types.StringValue(sslservicegroupName)

	tflog.Trace(ctx, "Created sslservicegroup resource")

	// Read the updated state back
	if !r.readSslservicegroupFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslservicegroup not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslservicegroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslservicegroup resource")

	found := r.readSslservicegroupFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SslservicegroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SslservicegroupResourceModel

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

	tflog.Debug(ctx, "Updating sslservicegroup resource")

	// Check whether any updateable attribute changed (servicegroupname is
	// ForceNew and never reaches Update).
	hasChange := false
	attributesToUnset := []string{}
	// commonname has no documented NITRO default, so it is not unset here.
	if !data.Commonname.Equal(state.Commonname) {
		hasChange = true
	}
	if !data.Ocspstapling.Equal(state.Ocspstapling) {
		if config.Ocspstapling.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "ocspstapling")
		} else {
			hasChange = true
		}
	}
	if !data.Sendclosenotify.Equal(state.Sendclosenotify) {
		if config.Sendclosenotify.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sendclosenotify")
		} else {
			hasChange = true
		}
	}
	if !data.Serverauth.Equal(state.Serverauth) {
		if config.Serverauth.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "serverauth")
		} else {
			hasChange = true
		}
	}
	if !data.Sessreuse.Equal(state.Sessreuse) {
		if config.Sessreuse.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sessreuse")
		} else {
			hasChange = true
		}
	}
	if !data.Sesstimeout.Equal(state.Sesstimeout) {
		if config.Sesstimeout.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sesstimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Snienable.Equal(state.Snienable) {
		if config.Snienable.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "snienable")
		} else {
			hasChange = true
		}
	}
	if !data.Ssl3.Equal(state.Ssl3) {
		if config.Ssl3.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "ssl3")
		} else {
			hasChange = true
		}
	}
	if !data.Sslclientlogs.Equal(state.Sslclientlogs) {
		if config.Sslclientlogs.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sslclientlogs")
		} else {
			hasChange = true
		}
	}
	// sslprofile has no documented NITRO default, so it is not unset here.
	if !data.Sslprofile.Equal(state.Sslprofile) {
		hasChange = true
	}
	if !data.Strictsigdigestcheck.Equal(state.Strictsigdigestcheck) {
		if config.Strictsigdigestcheck.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "strictsigdigestcheck")
		} else {
			hasChange = true
		}
	}
	if !data.Tls1.Equal(state.Tls1) {
		if config.Tls1.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "tls1")
		} else {
			hasChange = true
		}
	}
	if !data.Tls11.Equal(state.Tls11) {
		if config.Tls11.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "tls11")
		} else {
			hasChange = true
		}
	}
	if !data.Tls12.Equal(state.Tls12) {
		if config.Tls12.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "tls12")
		} else {
			hasChange = true
		}
	}
	if !data.Tls13.Equal(state.Tls13) {
		if config.Tls13.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "tls13")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		sslservicegroupName := data.Servicegroupname.ValueString()
		sslservicegroup := sslservicegroupGetThePayloadFromthePlan(ctx, &data)
		_, err := r.client.UpdateResource(service.Sslservicegroup.Type(), sslservicegroupName, &sslservicegroup)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslservicegroup, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated sslservicegroup resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslservicegroup resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"servicegroupname": data.Servicegroupname.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Sslservicegroup.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset sslservicegroup attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readSslservicegroupFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslservicegroup not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslservicegroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslservicegroup resource")

	// sslservicegroup has no DELETE operation on the ADC (its SSL config lives
	// with the underlying service group). Removing it from state only, matching
	// SDK v2 behavior.
	tflog.Trace(ctx, "Deleted sslservicegroup resource from state")
}

// Helper function to read sslservicegroup data from API. Returns false when the
// service group no longer exists so the caller can drop it from state.
func (r *SslservicegroupResource) readSslservicegroupFromApi(ctx context.Context, data *SslservicegroupResourceModel, diags *diag.Diagnostics) bool {
	sslservicegroupName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Sslservicegroup.Type(), sslservicegroupName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslservicegroup, got error: %s", err))
		return false
	}

	sslservicegroupSetAttrFromGet(ctx, data, getResponseData)

	return true
}
