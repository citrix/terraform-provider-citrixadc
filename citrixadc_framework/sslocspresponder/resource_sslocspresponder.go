package sslocspresponder

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
var _ resource.Resource = &SslocspresponderResource{}
var _ resource.ResourceWithConfigure = (*SslocspresponderResource)(nil)
var _ resource.ResourceWithImportState = (*SslocspresponderResource)(nil)

func NewSslocspresponderResource() resource.Resource {
	return &SslocspresponderResource{}
}

// SslocspresponderResource defines the resource implementation.
type SslocspresponderResource struct {
	client *service.NitroClient
}

func (r *SslocspresponderResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslocspresponderResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslocspresponder"
}

func (r *SslocspresponderResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslocspresponderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslocspresponderResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslocspresponder resource")

	sslocspresponder := sslocspresponderGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	sslocspresponderName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Sslocspresponder.Type(), sslocspresponderName, &sslocspresponder)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslocspresponder, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslocspresponder resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readSslocspresponderFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslocspresponder not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslocspresponderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslocspresponderResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslocspresponder resource")

	found := r.readSslocspresponderFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SslocspresponderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SslocspresponderResourceModel

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

	tflog.Debug(ctx, "Updating sslocspresponder resource")

	// Check if there are any changes in updateable attributes.
	// name, respondercert, signingcert are RequiresReplace and never reach Update.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Batchingdelay.Equal(state.Batchingdelay) {
		tflog.Debug(ctx, "batchingdelay has changed for sslocspresponder")
		hasChange = true
	}
	if !data.Batchingdepth.Equal(state.Batchingdepth) {
		tflog.Debug(ctx, "batchingdepth has changed for sslocspresponder")
		hasChange = true
	}
	if !data.Cache.Equal(state.Cache) {
		tflog.Debug(ctx, "cache has changed for sslocspresponder")
		if config.Cache.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "cache")
		} else {
			hasChange = true
		}
	}
	if !data.Cachetimeout.Equal(state.Cachetimeout) {
		tflog.Debug(ctx, "cachetimeout has changed for sslocspresponder")
		if config.Cachetimeout.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "cachetimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Httpmethod.Equal(state.Httpmethod) {
		tflog.Debug(ctx, "httpmethod has changed for sslocspresponder")
		if config.Httpmethod.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "httpmethod")
		} else {
			hasChange = true
		}
	}
	if !data.Insertclientcert.Equal(state.Insertclientcert) {
		tflog.Debug(ctx, "insertclientcert has changed for sslocspresponder")
		hasChange = true
	}
	if !data.Ocspurlresolvetimeout.Equal(state.Ocspurlresolvetimeout) {
		tflog.Debug(ctx, "ocspurlresolvetimeout has changed for sslocspresponder")
		hasChange = true
	}
	if !data.Producedattimeskew.Equal(state.Producedattimeskew) {
		tflog.Debug(ctx, "producedattimeskew has changed for sslocspresponder")
		if config.Producedattimeskew.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "producedattimeskew")
		} else {
			hasChange = true
		}
	}
	if !data.Resptimeout.Equal(state.Resptimeout) {
		tflog.Debug(ctx, "resptimeout has changed for sslocspresponder")
		hasChange = true
	}
	if !data.Trustresponder.Equal(state.Trustresponder) {
		tflog.Debug(ctx, "trustresponder has changed for sslocspresponder")
		if config.Trustresponder.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "trustresponder")
		} else {
			hasChange = true
		}
	}
	if !data.Url.Equal(state.Url) {
		tflog.Debug(ctx, "url has changed for sslocspresponder")
		hasChange = true
	}
	if !data.Usenonce.Equal(state.Usenonce) {
		tflog.Debug(ctx, "usenonce has changed for sslocspresponder")
		hasChange = true
	}

	if hasChange {
		sslocspresponder := sslocspresponderGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Named resource - use UpdateResource
		sslocspresponderName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Sslocspresponder.Type(), sslocspresponderName, &sslocspresponder)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslocspresponder, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslocspresponder resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslocspresponder resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Sslocspresponder.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset sslocspresponder attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readSslocspresponderFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslocspresponder not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslocspresponderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslocspresponderResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslocspresponder resource")

	// Named resource - delete using DeleteResource
	sslocspresponderName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Sslocspresponder.Type(), sslocspresponderName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslocspresponder, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslocspresponder resource")
}

// Helper function to read sslocspresponder data from API
func (r *SslocspresponderResource) readSslocspresponderFromApi(ctx context.Context, data *SslocspresponderResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (the name)
	sslocspresponderName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Sslocspresponder.Type(), sslocspresponderName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslocspresponder, got error: %s", err))
		return false
	}

	sslocspresponderSetAttrFromGet(ctx, data, getResponseData)

	return true
}
