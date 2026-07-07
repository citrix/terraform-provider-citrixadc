package ssldefaultprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ssldefaultprofile is an ACTION-ONLY, ZERO-ATTRIBUTE resource.
//
//   - NITRO exposes only the convert action:
//     POST /nitro/v1/config/ssldefaultprofile?action=convert, which converts the
//     appliance to the SSL default profile mode.
//   - There is NO add/set/get/delete endpoint, so:
//     Create performs the convert action, Read/Update are no-ops (there is
//     nothing to reconcile), and Delete is a state-only removal.
//   - Because there is no GET endpoint, there is NO datasource for
//     ssldefaultprofile.
var _ resource.Resource = &SsldefaultprofileResource{}
var _ resource.ResourceWithConfigure = (*SsldefaultprofileResource)(nil)
var _ resource.ResourceWithImportState = (*SsldefaultprofileResource)(nil)

func NewSsldefaultprofileResource() resource.Resource {
	return &SsldefaultprofileResource{}
}

// SsldefaultprofileResource defines the resource implementation.
type SsldefaultprofileResource struct {
	client *service.NitroClient
}

func (r *SsldefaultprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SsldefaultprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssldefaultprofile"
}

func (r *SsldefaultprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SsldefaultprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SsldefaultprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ssldefaultprofile resource (convert action)")
	ssldefaultprofile := ssldefaultprofileGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - NITRO exposes only POST ?action=convert
	err := r.client.ActOnResource(ssldefaultprofileResourceType, ssldefaultprofile, "convert")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to convert ssldefaultprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Converted ssldefaultprofile")

	// Synthetic ID - no GET endpoint exists to derive it from
	data.Id = types.StringValue("ssldefaultprofile-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. ssldefaultprofile has no GET endpoint; nothing to reconcile.
func (r *SsldefaultprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SsldefaultprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for ssldefaultprofile; NITRO exposes no GET endpoint (action=convert only)")

	// Preserve prior state unchanged - no GET endpoint to reconcile against
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. ssldefaultprofile has no attributes and no set endpoint.
func (r *SsldefaultprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SsldefaultprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for ssldefaultprofile; it has no attributes and no set endpoint")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. ssldefaultprofile has no delete endpoint; the action is not
// reversible and there is no persistent object to remove.
func (r *SsldefaultprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for ssldefaultprofile; NITRO has no delete endpoint")
}
