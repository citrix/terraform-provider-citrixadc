package sslcertkey_sslocspresponder_binding

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslcertkeySslocspresponderBindingResource{}
var _ resource.ResourceWithConfigure = (*SslcertkeySslocspresponderBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslcertkeySslocspresponderBindingResource)(nil)

func NewSslcertkeySslocspresponderBindingResource() resource.Resource {
	return &SslcertkeySslocspresponderBindingResource{}
}

// SslcertkeySslocspresponderBindingResource defines the resource implementation.
type SslcertkeySslocspresponderBindingResource struct {
	client *service.NitroClient
}

func (r *SslcertkeySslocspresponderBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslcertkeySslocspresponderBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcertkey_sslocspresponder_binding"
}

func (r *SslcertkeySslocspresponderBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslcertkeySslocspresponderBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslcertkeySslocspresponderBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslcertkey_sslocspresponder_binding resource")
	sslcertkey_sslocspresponder_binding := sslcertkey_sslocspresponder_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - NITRO "add" for this binding is a POST (matches the SDK v2
	// implementation which used AddResource). Pattern 1.
	_, err := r.client.AddResource(service.Sslcertkey_sslocspresponder_binding.Type(), "", &sslcertkey_sslocspresponder_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslcertkey_sslocspresponder_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslcertkey_sslocspresponder_binding resource")

	// Set ID for the resource before reading state.
	// Legacy-compatible composite key order: certkey,ocspresponder (see
	// resource_id_mapping.json). "ca" is a delete-disambiguation arg, not an identity
	// attribute, so it is NOT part of the ID.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("certkey:%s", utils.UrlEncode(data.Certkey.ValueString())))
	idParts = append(idParts, fmt.Sprintf("ocspresponder:%s", utils.UrlEncode(data.Ocspresponder.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSslcertkeySslocspresponderBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "sslcertkey_sslocspresponder_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertkeySslocspresponderBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslcertkeySslocspresponderBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslcertkey_sslocspresponder_binding resource")

	r.readSslcertkeySslocspresponderBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	// Binding is gone on the ADC (readFromApi nulled the Id): drop it from state so a
	// subsequent apply recreates it, matching the SDK v2 provider's behaviour.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertkeySslocspresponderBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslcertkeySslocspresponderBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslcertkey_sslocspresponder_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		sslcertkey_sslocspresponder_binding := sslcertkey_sslocspresponder_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Sslcertkey_sslocspresponder_binding.Type(), &sslcertkey_sslocspresponder_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslcertkey_sslocspresponder_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslcertkey_sslocspresponder_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslcertkey_sslocspresponder_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSslcertkeySslocspresponderBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "sslcertkey_sslocspresponder_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertkeySslocspresponderBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslcertkeySslocspresponderBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslcertkey_sslocspresponder_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// The composite ID is certkey,ocspresponder (legacy order). "ca" comes from state.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"certkey", "ocspresponder"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	certkey_value, ok := idMap["certkey"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'certkey' not found in ID")
		return
	}

	args := make([]string, 0)
	if val, ok := idMap["ocspresponder"]; ok && val != "" {
		// URL-encode the arg value to handle slashy/special characters.
		args = append(args, fmt.Sprintf("ocspresponder:%s", url.QueryEscape(val)))
	}
	if !data.Ca.IsNull() && !data.Ca.IsUnknown() {
		args = append(args, fmt.Sprintf("ca:%s", strconv.FormatBool(data.Ca.ValueBool())))
	}

	err = r.client.DeleteResourceWithArgs(service.Sslcertkey_sslocspresponder_binding.Type(), certkey_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslcertkey_sslocspresponder_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslcertkey_sslocspresponder_binding binding")
}

// Helper function to read sslcertkey_sslocspresponder_binding data from API
func (r *SslcertkeySslocspresponderBindingResource) readSslcertkeySslocspresponderBindingFromApi(ctx context.Context, data *SslcertkeySslocspresponderBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"certkey", "ocspresponder"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	certkey_Name, ok := idMap["certkey"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'certkey' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Sslcertkey_sslocspresponder_binding.Type(),
		ResourceName:             certkey_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslcertkey_sslocspresponder_binding, got error: %s", err))
		return
	}

	// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
	// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id.
	// Match on ocspresponder only - "ca" is never returned by the NITRO GET response
	// for this binding, so it cannot be used as a filter (it is a delete-only arg).
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ocspresponder
		if idVal, ok := idMap["ocspresponder"]; ok {
			if val, ok := v["ocspresponder"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["ocspresponder"].(string); ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	// Binding not present in the returned set: signal removal via a null Id (see above).
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	sslcertkey_sslocspresponder_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
