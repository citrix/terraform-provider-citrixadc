package nshttpparam

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
var _ resource.Resource = &NshttpparamResource{}
var _ resource.ResourceWithConfigure = (*NshttpparamResource)(nil)
var _ resource.ResourceWithImportState = (*NshttpparamResource)(nil)

func NewNshttpparamResource() resource.Resource {
	return &NshttpparamResource{}
}

// NshttpparamResource defines the resource implementation.
type NshttpparamResource struct {
	client *service.NitroClient
}

func (r *NshttpparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NshttpparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nshttpparam"
}

func (r *NshttpparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NshttpparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NshttpparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nshttpparam resource")

	nshttpparam := nshttpparamGetThePayloadFromtheConfig(ctx, &data)

	// nshttpparam is a singleton configuration resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Nshttpparam.Type(), &nshttpparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nshttpparam, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nshttpparam-config")

	tflog.Trace(ctx, "Created nshttpparam resource")

	// Read the updated state back
	r.readNshttpparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NshttpparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NshttpparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nshttpparam resource")

	r.readNshttpparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NshttpparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NshttpparamResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nshttpparam resource")

	// Determine attributes removed from config so they can be unset (reverted
	// to their NITRO defaults) after the update.
	attributesToUnset := []string{}
	if !data.Conmultiplex.Equal(state.Conmultiplex) && config.Conmultiplex.IsNull() {
		attributesToUnset = append(attributesToUnset, "conmultiplex")
	}
	if !data.Dropinvalreqs.Equal(state.Dropinvalreqs) && config.Dropinvalreqs.IsNull() {
		attributesToUnset = append(attributesToUnset, "dropinvalreqs")
	}
	if !data.Http2serverside.Equal(state.Http2serverside) && config.Http2serverside.IsNull() {
		attributesToUnset = append(attributesToUnset, "http2serverside")
	}
	if !data.Ignoreconnectcodingscheme.Equal(state.Ignoreconnectcodingscheme) && config.Ignoreconnectcodingscheme.IsNull() {
		attributesToUnset = append(attributesToUnset, "ignoreconnectcodingscheme")
	}
	if !data.Insnssrvrhdr.Equal(state.Insnssrvrhdr) && config.Insnssrvrhdr.IsNull() {
		attributesToUnset = append(attributesToUnset, "insnssrvrhdr")
	}
	if !data.Logerrresp.Equal(state.Logerrresp) && config.Logerrresp.IsNull() {
		attributesToUnset = append(attributesToUnset, "logerrresp")
	}
	if !data.Markconnreqinval.Equal(state.Markconnreqinval) && config.Markconnreqinval.IsNull() {
		attributesToUnset = append(attributesToUnset, "markconnreqinval")
	}
	if !data.Markhttp09inval.Equal(state.Markhttp09inval) && config.Markhttp09inval.IsNull() {
		attributesToUnset = append(attributesToUnset, "markhttp09inval")
	}
	if !data.Maxreusepool.Equal(state.Maxreusepool) && config.Maxreusepool.IsNull() {
		attributesToUnset = append(attributesToUnset, "maxreusepool")
	}

	// Create API request body from the model
	nshttpparam := nshttpparamGetThePayloadFromtheConfig(ctx, &data)

	// nshttpparam is a singleton configuration resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Nshttpparam.Type(), &nshttpparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nshttpparam, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated nshttpparam resource")

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Nshttpparam.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset nshttpparam attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readNshttpparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NshttpparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NshttpparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nshttpparam resource")

	// For nshttpparam, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nshttpparam resource from state")
}

// Helper function to read nshttpparam data from API
func (r *NshttpparamResource) readNshttpparamFromApi(ctx context.Context, data *NshttpparamResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nshttpparam.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nshttpparam, got error: %s", err))
		return
	}

	nshttpparamSetAttrFromGet(ctx, data, getResponseData)

}
