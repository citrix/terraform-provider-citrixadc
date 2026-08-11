package nstcpparam

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
var _ resource.Resource = &NstcpparamResource{}
var _ resource.ResourceWithConfigure = (*NstcpparamResource)(nil)
var _ resource.ResourceWithImportState = (*NstcpparamResource)(nil)

func NewNstcpparamResource() resource.Resource {
	return &NstcpparamResource{}
}

// NstcpparamResource defines the resource implementation.
type NstcpparamResource struct {
	client *service.NitroClient
}

func (r *NstcpparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NstcpparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nstcpparam"
}

func (r *NstcpparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NstcpparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NstcpparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nstcpparam resource")

	nstcpparam := nstcpparamGetThePayloadFromtheConfig(ctx, &data)

	// nstcpparam is a singleton/unnamed configuration resource (SDK v2 parity:
	// createNstcpparamFunc called client.UpdateUnnamedResource).
	err := r.client.UpdateUnnamedResource(service.Nstcpparam.Type(), &nstcpparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nstcpparam, got error: %s", err))
		return
	}

	// Static singleton ID (no unique attributes).
	data.Id = types.StringValue("nstcpparam-config")

	tflog.Trace(ctx, "Created nstcpparam resource")

	// Read the updated state back
	r.readNstcpparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstcpparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NstcpparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nstcpparam resource")

	r.readNstcpparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstcpparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NstcpparamResourceModel

	// Read Terraform prior state, plan, and config into the models
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (candidates for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nstcpparam resource")

	// Most configurable attributes are RequiresReplaceIfConfigured (SDK v2
	// ForceNew), so a configured change is handled by recreate rather than
	// Update. This branch still pushes the config to keep computed-value
	// resolution consistent.
	//
	// A handful of attributes carry a schema Default so that removing them from
	// config produces a plan diff (Default value vs. prior non-default state),
	// routing through Update. For those, if the attribute is now absent from
	// config, we unset it on the appliance so it reverts to its NITRO default.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Mptcpmaxpendingsf.Equal(state.Mptcpmaxpendingsf) {
		if config.Mptcpmaxpendingsf.IsNull() {
			attributesToUnset = append(attributesToUnset, "mptcpmaxpendingsf")
		} else {
			hasChange = true
		}
	}
	if !data.Mptcppendingjointhreshold.Equal(state.Mptcppendingjointhreshold) {
		if config.Mptcppendingjointhreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "mptcppendingjointhreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Mptcpsfreplacetimeout.Equal(state.Mptcpsfreplacetimeout) {
		if config.Mptcpsfreplacetimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "mptcpsfreplacetimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Mptcpsftimeout.Equal(state.Mptcpsftimeout) {
		if config.Mptcpsftimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "mptcpsftimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Oooqsize.Equal(state.Oooqsize) {
		if config.Oooqsize.IsNull() {
			attributesToUnset = append(attributesToUnset, "oooqsize")
		} else {
			hasChange = true
		}
	}
	if !data.Rfc5961chlgacklimit.Equal(state.Rfc5961chlgacklimit) {
		if config.Rfc5961chlgacklimit.IsNull() {
			attributesToUnset = append(attributesToUnset, "rfc5961chlgacklimit")
		} else {
			hasChange = true
		}
	}
	if !data.Tcpfastopencookietimeout.Equal(state.Tcpfastopencookietimeout) {
		if config.Tcpfastopencookietimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "tcpfastopencookietimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Wsval.Equal(state.Wsval) {
		if config.Wsval.IsNull() {
			attributesToUnset = append(attributesToUnset, "wsval")
		} else {
			hasChange = true
		}
	}

	_ = hasChange

	nstcpparam := nstcpparamGetThePayloadFromtheConfig(ctx, &data)

	err := r.client.UpdateUnnamedResource(service.Nstcpparam.Type(), &nstcpparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nstcpparam, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated nstcpparam resource")

	// Unset attributes removed from config so the appliance reverts them to
	// their NITRO defaults. nstcpparam is a singleton (unnamed) resource, so the
	// unset id payload is empty. Done after the update so any default value the
	// update payload carried is superseded by the unset.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Nstcpparam.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset nstcpparam attributes, got error: %s", err))
		return
	}

	// Read the updated state back (data.Id carried over from plan/prior state)
	r.readNstcpparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstcpparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NstcpparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nstcpparam resource")

	// nstcpparam is a global configuration resource: it cannot actually be
	// deleted, only reset. Matching SDK v2 deleteNstcpparamFunc, Delete just
	// removes the reference from Terraform state.
	tflog.Trace(ctx, "Deleted nstcpparam resource from state")
}

// Helper function to read nstcpparam data from API
func (r *NstcpparamResource) readNstcpparamFromApi(ctx context.Context, data *NstcpparamResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nstcpparam.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nstcpparam, got error: %s", err))
		return
	}

	nstcpparamSetAttrFromGet(ctx, data, getResponseData)
}
