package lbmonitor_sslcertkey_binding

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
var _ resource.Resource = &LbmonitorSslcertkeyBindingResource{}
var _ resource.ResourceWithConfigure = (*LbmonitorSslcertkeyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbmonitorSslcertkeyBindingResource)(nil)

func NewLbmonitorSslcertkeyBindingResource() resource.Resource {
	return &LbmonitorSslcertkeyBindingResource{}
}

// LbmonitorSslcertkeyBindingResource defines the resource implementation.
type LbmonitorSslcertkeyBindingResource struct {
	client *service.NitroClient
}

func (r *LbmonitorSslcertkeyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbmonitorSslcertkeyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbmonitor_sslcertkey_binding"
}

func (r *LbmonitorSslcertkeyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbmonitorSslcertkeyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbmonitorSslcertkeyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbmonitor_sslcertkey_binding resource")
	lbmonitor_sslcertkey_binding := lbmonitor_sslcertkey_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Lbmonitor_sslcertkey_binding.Type(), &lbmonitor_sslcertkey_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbmonitor_sslcertkey_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lbmonitor_sslcertkey_binding resource")

	// Set ID for the resource before reading state.
	// Legacy SDK v2 ID order: monitorname,certkeyname (see resource_id_mapping.json).
	// 'ca' is a binding property, not part of the identity, so it is excluded from the ID.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("monitorname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Monitorname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("certkeyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Certkeyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readLbmonitorSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorSslcertkeyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbmonitorSslcertkeyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbmonitor_sslcertkey_binding resource")

	r.readLbmonitorSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorSslcertkeyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LbmonitorSslcertkeyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lbmonitor_sslcertkey_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		lbmonitor_sslcertkey_binding := lbmonitor_sslcertkey_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Lbmonitor_sslcertkey_binding.Type(), &lbmonitor_sslcertkey_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbmonitor_sslcertkey_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated lbmonitor_sslcertkey_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lbmonitor_sslcertkey_binding resource, skipping update")
	}

	// Read the updated state back
	r.readLbmonitorSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorSslcertkeyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbmonitorSslcertkeyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbmonitor_sslcertkey_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgsMap.
	// Legacy SDK v2 ID order: monitorname,certkeyname (see resource_id_mapping.json).
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"monitorname", "certkeyname"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	monitorname_value, ok := idMap["monitorname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'monitorname' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["certkeyname"]; ok && val != "" {
		// ParseIdString already URL-decodes; re-encode for the delete query arg.
		argsMap["certkeyname"] = url.QueryEscape(val)
	}
	// 'ca' distinguishes CA-cert bindings from server-cert bindings on the same
	// monitor; pass it only when set true (matches SDK v2 delete behavior).
	if !data.Ca.IsNull() && !data.Ca.IsUnknown() && data.Ca.ValueBool() {
		argsMap["ca"] = url.QueryEscape(fmt.Sprintf("%v", data.Ca.ValueBool()))
	}

	err = r.client.DeleteResourceWithArgsMap(service.Lbmonitor_sslcertkey_binding.Type(), monitorname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lbmonitor_sslcertkey_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lbmonitor_sslcertkey_binding binding")
}

// Helper function to read lbmonitor_sslcertkey_binding data from API
func (r *LbmonitorSslcertkeyBindingResource) readLbmonitorSslcertkeyBindingFromApi(ctx context.Context, data *LbmonitorSslcertkeyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"monitorname", "certkeyname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	monitorname_Name, ok := idMap["monitorname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'monitorname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Lbmonitor_sslcertkey_binding.Type(),
		ResourceName:             monitorname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbmonitor_sslcertkey_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "lbmonitor_sslcertkey_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id.
	// Identity is monitorname (parent) + certkeyname; match on certkeyname here.
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check certkeyname
		if idVal, ok := idMap["certkeyname"]; ok {
			if val, ok := v["certkeyname"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["certkeyname"].(string); ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("lbmonitor_sslcertkey_binding not found with the provided ID attributes"))
		return
	}

	lbmonitor_sslcertkey_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
