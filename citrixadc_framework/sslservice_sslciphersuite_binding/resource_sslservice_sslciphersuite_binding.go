package sslservice_sslciphersuite_binding

import (
	"context"
	"fmt"
	"net/url"
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
var _ resource.Resource = &SslserviceSslciphersuiteBindingResource{}
var _ resource.ResourceWithConfigure = (*SslserviceSslciphersuiteBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslserviceSslciphersuiteBindingResource)(nil)

func NewSslserviceSslciphersuiteBindingResource() resource.Resource {
	return &SslserviceSslciphersuiteBindingResource{}
}

// SslserviceSslciphersuiteBindingResource defines the resource implementation.
type SslserviceSslciphersuiteBindingResource struct {
	client *service.NitroClient
}

func (r *SslserviceSslciphersuiteBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslserviceSslciphersuiteBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslservice_sslciphersuite_binding"
}

func (r *SslserviceSslciphersuiteBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslserviceSslciphersuiteBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslserviceSslciphersuiteBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslservice_sslciphersuite_binding resource")
	sslservice_sslciphersuite_binding := sslservice_sslciphersuite_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - NITRO add is POST (matches SDK v2 AddResource)
	_, err := r.client.AddResource(service.Sslservice_sslciphersuite_binding.Type(), "", &sslservice_sslciphersuite_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslservice_sslciphersuite_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslservice_sslciphersuite_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ciphername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ciphername.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicename.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSslserviceSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "sslservice_sslciphersuite_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceSslciphersuiteBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslserviceSslciphersuiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslservice_sslciphersuite_binding resource")

	r.readSslserviceSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SslserviceSslciphersuiteBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslserviceSslciphersuiteBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslservice_sslciphersuite_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		sslservice_sslciphersuite_binding := sslservice_sslciphersuite_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Sslservice_sslciphersuite_binding.Type(), &sslservice_sslciphersuite_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslservice_sslciphersuite_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslservice_sslciphersuite_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslservice_sslciphersuite_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSslserviceSslciphersuiteBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "sslservice_sslciphersuite_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceSslciphersuiteBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslserviceSslciphersuiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslservice_sslciphersuite_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicename", "ciphername"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	servicename_value, ok := idMap["servicename"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'servicename' not found in ID")
		return
	}

	// DeleteResourceWithArgs joins args verbatim into the URL query string and does NOT
	// URL-encode the value, so encode the ciphername here (handles slashy/special cipher names).
	args := make([]string, 0)
	if val, ok := idMap["ciphername"]; ok && val != "" {
		args = append(args, fmt.Sprintf("ciphername:%s", url.QueryEscape(val)))
	}

	err = r.client.DeleteResourceWithArgs(service.Sslservice_sslciphersuite_binding.Type(), servicename_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslservice_sslciphersuite_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslservice_sslciphersuite_binding binding")
}

// Helper function to read sslservice_sslciphersuite_binding data from API
func (r *SslserviceSslciphersuiteBindingResource) readSslserviceSslciphersuiteBindingFromApi(ctx context.Context, data *SslserviceSslciphersuiteBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicename", "ciphername"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	servicename_Name, ok := idMap["servicename"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'servicename' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Sslservice_sslciphersuite_binding.Type(),
		ResourceName:             servicename_Name,
		ResourceMissingErrorCode: 463,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslservice_sslciphersuite_binding, got error: %s", err))
		return
	}

	// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
	// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ciphername
		if idVal, ok := idMap["ciphername"]; ok {
			if val, ok := v["ciphername"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["ciphername"].(string); ok {
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

	sslservice_sslciphersuite_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
