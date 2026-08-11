package nsparam

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
var _ resource.Resource = &NsparamResource{}
var _ resource.ResourceWithConfigure = (*NsparamResource)(nil)
var _ resource.ResourceWithImportState = (*NsparamResource)(nil)

func NewNsparamResource() resource.Resource {
	return &NsparamResource{}
}

// NsparamResource defines the resource implementation.
type NsparamResource struct {
	client *service.NitroClient
}

func (r *NsparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsparam"
}

func (r *NsparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsparam resource")

	// Build the payload from configured attributes only (mirrors SDK v2 GetOk).
	nsparam := nsparamGetThePayloadFromtheConfig(ctx, &data)

	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Nsparam.Type(), &nsparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsparam, got error: %s", err))
		return
	}

	// Singleton static ID (SDK v2 used a random per-apply ID; a stable static ID
	// is the Framework singleton convention and keeps refresh/import stable).
	data.Id = types.StringValue("nsparam-config")

	tflog.Trace(ctx, "Created nsparam resource")

	// Read the updated state back
	if !r.readNsparamFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsparam not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsparam resource")

	found := r.readNsparamFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NsparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NsparamResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (unset).
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nsparam resource")

	// Detect attributes that were removed from config (config null but the value
	// changed from prior state). Each such attribute is reverted to its NITRO
	// default via ?action=unset. Attributes carry a schema Default equal to the
	// server default, so removal produces a plan diff that reaches this Update.
	attributesToUnset := []string{}
	if !data.Advancedanalyticsstats.Equal(state.Advancedanalyticsstats) && config.Advancedanalyticsstats.IsNull() {
		attributesToUnset = append(attributesToUnset, "advancedanalyticsstats")
	}
	if !data.Aftpallowrandomsourceport.Equal(state.Aftpallowrandomsourceport) && config.Aftpallowrandomsourceport.IsNull() {
		attributesToUnset = append(attributesToUnset, "aftpallowrandomsourceport")
	}
	if !data.Securecookie.Equal(state.Securecookie) && config.Securecookie.IsNull() {
		attributesToUnset = append(attributesToUnset, "securecookie")
	}
	if !data.Tcpcip.Equal(state.Tcpcip) && config.Tcpcip.IsNull() {
		attributesToUnset = append(attributesToUnset, "tcpcip")
	}
	if !data.Proxyprotocol.Equal(state.Proxyprotocol) && config.Proxyprotocol.IsNull() {
		attributesToUnset = append(attributesToUnset, "proxyprotocol")
	}
	if !data.Pmtumin.Equal(state.Pmtumin) && config.Pmtumin.IsNull() {
		attributesToUnset = append(attributesToUnset, "pmtumin")
	}
	if !data.Pmtutimeout.Equal(state.Pmtutimeout) && config.Pmtutimeout.IsNull() {
		attributesToUnset = append(attributesToUnset, "pmtutimeout")
	}
	if !data.Grantquotamaxclient.Equal(state.Grantquotamaxclient) && config.Grantquotamaxclient.IsNull() {
		attributesToUnset = append(attributesToUnset, "grantquotamaxclient")
	}
	if !data.Exclusivequotamaxclient.Equal(state.Exclusivequotamaxclient) && config.Exclusivequotamaxclient.IsNull() {
		attributesToUnset = append(attributesToUnset, "exclusivequotamaxclient")
	}
	if !data.Grantquotaspillover.Equal(state.Grantquotaspillover) && config.Grantquotaspillover.IsNull() {
		attributesToUnset = append(attributesToUnset, "grantquotaspillover")
	}
	if !data.Exclusivequotaspillover.Equal(state.Exclusivequotaspillover) && config.Exclusivequotaspillover.IsNull() {
		attributesToUnset = append(attributesToUnset, "exclusivequotaspillover")
	}

	// NOTE: In SDK v2 every attribute was ForceNew, so this Update path is only
	// reached for non-force-new changes. It re-pushes the configured attributes to
	// keep behaviour correct regardless.
	nsparam := nsparamGetThePayloadFromtheConfig(ctx, &data)

	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Nsparam.Type(), &nsparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsparam, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated nsparam resource")

	// Unset attributes that were removed from config so the appliance reverts them
	// to their defaults. Done after the update so any default value the update
	// payload carried is superseded by the unset. nsparam is a singleton, so the
	// id payload is empty.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Nsparam.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset nsparam attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readNsparamFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsparam not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// nsparam is a global singleton configuration on the ADC. There is nothing to
	// delete (this mirrors the SDK v2 no-op delete); the framework removes the
	// resource from Terraform state after this returns without error.
	tflog.Debug(ctx, "Deleting nsparam resource (no-op for singleton, removing from state only)")
}

// Helper function to read nsparam data from API
func (r *NsparamResource) readNsparamFromApi(ctx context.Context, data *NsparamResourceModel, diags *diag.Diagnostics) bool {
	// Singleton - simple find without ID
	getResponseData, err := r.client.FindResource(service.Nsparam.Type(), "")
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsparam, got error: %s", err))
		return false
	}

	nsparamSetAttrFromGet(ctx, data, getResponseData)

	return true
}
