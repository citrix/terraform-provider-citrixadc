package nshttpprofile

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
var _ resource.Resource = &NshttpprofileResource{}
var _ resource.ResourceWithConfigure = (*NshttpprofileResource)(nil)
var _ resource.ResourceWithImportState = (*NshttpprofileResource)(nil)

func NewNshttpprofileResource() resource.Resource {
	return &NshttpprofileResource{}
}

// NshttpprofileResource defines the resource implementation.
type NshttpprofileResource struct {
	client *service.NitroClient
}

func (r *NshttpprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NshttpprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nshttpprofile"
}

func (r *NshttpprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NshttpprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NshttpprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nshttpprofile resource")

	// Create API request body from the plan
	nshttpprofile := nshttpprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	nshttpprofileName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Nshttpprofile.Type(), nshttpprofileName, &nshttpprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nshttpprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nshttpprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(nshttpprofileName)

	// Read the updated state back
	if !r.readNshttpprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nshttpprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NshttpprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NshttpprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nshttpprofile resource")

	found := r.readNshttpprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NshttpprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NshttpprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (name is ForceNew, so the ID never changes on update)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nshttpprofile resource")

	// Determine which unset-eligible attributes were removed from config so the
	// appliance reverts them to their NITRO defaults. An attribute is unset only
	// when it changed relative to prior state AND is absent from config.
	attributesToUnset := []string{}
	if !data.Adpttimeout.Equal(state.Adpttimeout) && config.Adpttimeout.IsNull() {
		attributesToUnset = append(attributesToUnset, "adpttimeout")
	}
	if !data.Allowonlywordcharactersandhyphen.Equal(state.Allowonlywordcharactersandhyphen) && config.Allowonlywordcharactersandhyphen.IsNull() {
		attributesToUnset = append(attributesToUnset, "allowonlywordcharactersandhyphen")
	}
	if !data.Altsvc.Equal(state.Altsvc) && config.Altsvc.IsNull() {
		attributesToUnset = append(attributesToUnset, "altsvc")
	}
	if !data.Cmponpush.Equal(state.Cmponpush) && config.Cmponpush.IsNull() {
		attributesToUnset = append(attributesToUnset, "cmponpush")
	}
	if !data.Conmultiplex.Equal(state.Conmultiplex) && config.Conmultiplex.IsNull() {
		attributesToUnset = append(attributesToUnset, "conmultiplex")
	}
	if !data.Dropextracrlf.Equal(state.Dropextracrlf) && config.Dropextracrlf.IsNull() {
		attributesToUnset = append(attributesToUnset, "dropextracrlf")
	}
	if !data.Dropextradata.Equal(state.Dropextradata) && config.Dropextradata.IsNull() {
		attributesToUnset = append(attributesToUnset, "dropextradata")
	}
	if !data.Dropinvalreqs.Equal(state.Dropinvalreqs) && config.Dropinvalreqs.IsNull() {
		attributesToUnset = append(attributesToUnset, "dropinvalreqs")
	}
	if !data.Grpclengthdelimitation.Equal(state.Grpclengthdelimitation) && config.Grpclengthdelimitation.IsNull() {
		attributesToUnset = append(attributesToUnset, "grpclengthdelimitation")
	}
	if !data.Hostheadervalidation.Equal(state.Hostheadervalidation) && config.Hostheadervalidation.IsNull() {
		attributesToUnset = append(attributesToUnset, "hostheadervalidation")
	}
	if !data.Http2.Equal(state.Http2) && config.Http2.IsNull() {
		attributesToUnset = append(attributesToUnset, "http2")
	}
	if !data.Http2altsvcframe.Equal(state.Http2altsvcframe) && config.Http2altsvcframe.IsNull() {
		attributesToUnset = append(attributesToUnset, "http2altsvcframe")
	}
	if !data.Http2direct.Equal(state.Http2direct) && config.Http2direct.IsNull() {
		attributesToUnset = append(attributesToUnset, "http2direct")
	}
	if !data.Http2extendedconnect.Equal(state.Http2extendedconnect) && config.Http2extendedconnect.IsNull() {
		attributesToUnset = append(attributesToUnset, "http2extendedconnect")
	}
	if !data.Http2strictcipher.Equal(state.Http2strictcipher) && config.Http2strictcipher.IsNull() {
		attributesToUnset = append(attributesToUnset, "http2strictcipher")
	}
	if !data.Http3.Equal(state.Http3) && config.Http3.IsNull() {
		attributesToUnset = append(attributesToUnset, "http3")
	}
	if !data.Http3webtransport.Equal(state.Http3webtransport) && config.Http3webtransport.IsNull() {
		attributesToUnset = append(attributesToUnset, "http3webtransport")
	}
	if !data.Markconnreqinval.Equal(state.Markconnreqinval) && config.Markconnreqinval.IsNull() {
		attributesToUnset = append(attributesToUnset, "markconnreqinval")
	}
	if !data.Markhttp09inval.Equal(state.Markhttp09inval) && config.Markhttp09inval.IsNull() {
		attributesToUnset = append(attributesToUnset, "markhttp09inval")
	}
	if !data.Markhttpheaderextrawserror.Equal(state.Markhttpheaderextrawserror) && config.Markhttpheaderextrawserror.IsNull() {
		attributesToUnset = append(attributesToUnset, "markhttpheaderextrawserror")
	}
	if !data.Markrfc7230noncompliantinval.Equal(state.Markrfc7230noncompliantinval) && config.Markrfc7230noncompliantinval.IsNull() {
		attributesToUnset = append(attributesToUnset, "markrfc7230noncompliantinval")
	}
	if !data.Marktracereqinval.Equal(state.Marktracereqinval) && config.Marktracereqinval.IsNull() {
		attributesToUnset = append(attributesToUnset, "marktracereqinval")
	}
	if !data.Passprotocolupgrade.Equal(state.Passprotocolupgrade) && config.Passprotocolupgrade.IsNull() {
		attributesToUnset = append(attributesToUnset, "passprotocolupgrade")
	}
	if !data.Persistentetag.Equal(state.Persistentetag) && config.Persistentetag.IsNull() {
		attributesToUnset = append(attributesToUnset, "persistentetag")
	}
	if !data.Rtsptunnel.Equal(state.Rtsptunnel) && config.Rtsptunnel.IsNull() {
		attributesToUnset = append(attributesToUnset, "rtsptunnel")
	}
	if !data.Weblog.Equal(state.Weblog) && config.Weblog.IsNull() {
		attributesToUnset = append(attributesToUnset, "weblog")
	}
	if !data.Websocket.Equal(state.Websocket) && config.Websocket.IsNull() {
		attributesToUnset = append(attributesToUnset, "websocket")
	}

	// Create API request body from the plan (only known, configured attributes are sent)
	nshttpprofile := nshttpprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use UpdateResource
	nshttpprofileName := data.Id.ValueString()
	_, err := r.client.UpdateResource(service.Nshttpprofile.Type(), nshttpprofileName, &nshttpprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nshttpprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated nshttpprofile resource")

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Done after the update so any default value the
	// update payload carried for a removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Nshttpprofile.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset nshttpprofile attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readNshttpprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nshttpprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NshttpprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NshttpprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nshttpprofile resource")

	// Named resource - delete using DeleteResource
	nshttpprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Nshttpprofile.Type(), nshttpprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nshttpprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nshttpprofile resource")
}

// Helper function to read nshttpprofile data from API.
// Returns false when the resource no longer exists on the ADC.
func (r *NshttpprofileResource) readNshttpprofileFromApi(ctx context.Context, data *NshttpprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain name value
	nshttpprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nshttpprofile.Type(), nshttpprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nshttpprofile, got error: %s", err))
		return false
	}

	nshttpprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
