package sslvserver

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
var _ resource.Resource = &SslvserverResource{}
var _ resource.ResourceWithConfigure = (*SslvserverResource)(nil)
var _ resource.ResourceWithImportState = (*SslvserverResource)(nil)

func NewSslvserverResource() resource.Resource {
	return &SslvserverResource{}
}

// SslvserverResource defines the resource implementation.
type SslvserverResource struct {
	client *service.NitroClient
}

func (r *SslvserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslvserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslvserver"
}

func (r *SslvserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslvserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslvserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslvserver resource")

	// An sslvserver is not created directly; it represents advanced SSL settings on an
	// already-existing SSL virtual server. Mirror SDK v2 behaviour and configure it via
	// UpdateResource keyed on vservername (there is no AddResource for sslvserver).
	sslvserver := sslvserverGetThePayloadFromtheConfig(ctx, &data)
	sslvserverName := data.Vservername.ValueString()

	_, err := r.client.UpdateResource(service.Sslvserver.Type(), sslvserverName, &sslvserver)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslvserver, got error: %s", err))
		return
	}

	// ID is the vservername (single unique attribute).
	data.Id = types.StringValue(sslvserverName)

	tflog.Trace(ctx, "Created sslvserver resource")

	// Read the updated state back
	if !r.readSslvserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslvserver not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslvserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslvserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslvserver resource")

	found := r.readSslvserverFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SslvserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslvserverResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslvserver resource")

	// All attributes except vservername (RequiresReplace) are updateable. Push the full
	// configured payload via UpdateResource, matching SDK v2 semantics.
	sslvserver := sslvserverGetThePayloadFromtheConfig(ctx, &data)
	sslvserverName := data.Vservername.ValueString()

	_, err := r.client.UpdateResource(service.Sslvserver.Type(), sslvserverName, &sslvserver)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslvserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated sslvserver resource")

	// Read the updated state back
	if !r.readSslvserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslvserver not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslvserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslvserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslvserver resource")

	// sslvserver has no DELETE operation. Mirror SDK v2 behaviour: if the sslprofile was
	// not managed by this resource, unset it so a bound sslprofile can subsequently be
	// deleted. The resource is then simply removed from Terraform state.
	if data.Sslprofile.IsNull() || data.Sslprofile.IsUnknown() || data.Sslprofile.ValueString() == "" {
		sslvserverName := data.Vservername.ValueString()
		if sslvserverName == "" {
			sslvserverName = data.Id.ValueString()
		}
		unsetPayload := map[string]interface{}{
			"vservername": sslvserverName,
			"sslprofile":  true,
		}
		if err := r.client.ActOnResource(service.Sslvserver.Type(), unsetPayload, "unset"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset sslprofile for sslvserver, got error: %s", err))
			return
		}
	}

	tflog.Trace(ctx, "Deleted sslvserver resource from state")
}

// Helper function to read sslvserver data from API. Returns false if the sslvserver no
// longer exists on the ADC (so the caller can remove it from state).
func (r *SslvserverResource) readSslvserverFromApi(ctx context.Context, data *SslvserverResourceModel, diags *diag.Diagnostics) bool {
	// ID is the plain vservername value.
	sslvserverName := data.Vservername.ValueString()
	if sslvserverName == "" {
		sslvserverName = data.Id.ValueString()
	}

	getResponseData, err := r.client.FindResource(service.Sslvserver.Type(), sslvserverName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslvserver, got error: %s", err))
		return false
	}

	sslvserverSetAttrFromGet(ctx, data, getResponseData)

	return true
}
