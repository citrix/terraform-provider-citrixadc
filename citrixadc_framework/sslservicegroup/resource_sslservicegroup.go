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
	var data, state SslservicegroupResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslservicegroup resource")

	// Check whether any updateable attribute changed (servicegroupname is
	// ForceNew and never reaches Update).
	hasChange := false
	if !data.Commonname.Equal(state.Commonname) {
		hasChange = true
	}
	if !data.Ocspstapling.Equal(state.Ocspstapling) {
		hasChange = true
	}
	if !data.Sendclosenotify.Equal(state.Sendclosenotify) {
		hasChange = true
	}
	if !data.Serverauth.Equal(state.Serverauth) {
		hasChange = true
	}
	if !data.Sessreuse.Equal(state.Sessreuse) {
		hasChange = true
	}
	if !data.Sesstimeout.Equal(state.Sesstimeout) {
		hasChange = true
	}
	if !data.Snienable.Equal(state.Snienable) {
		hasChange = true
	}
	if !data.Ssl3.Equal(state.Ssl3) {
		hasChange = true
	}
	if !data.Sslclientlogs.Equal(state.Sslclientlogs) {
		hasChange = true
	}
	if !data.Sslprofile.Equal(state.Sslprofile) {
		hasChange = true
	}
	if !data.Strictsigdigestcheck.Equal(state.Strictsigdigestcheck) {
		hasChange = true
	}
	if !data.Tls1.Equal(state.Tls1) {
		hasChange = true
	}
	if !data.Tls11.Equal(state.Tls11) {
		hasChange = true
	}
	if !data.Tls12.Equal(state.Tls12) {
		hasChange = true
	}
	if !data.Tls13.Equal(state.Tls13) {
		hasChange = true
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
