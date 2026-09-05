package lsnclient_network6_binding

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
var _ resource.Resource = &LsnclientNetwork6BindingResource{}
var _ resource.ResourceWithConfigure = (*LsnclientNetwork6BindingResource)(nil)
var _ resource.ResourceWithImportState = (*LsnclientNetwork6BindingResource)(nil)

func NewLsnclientNetwork6BindingResource() resource.Resource {
	return &LsnclientNetwork6BindingResource{}
}

// LsnclientNetwork6BindingResource defines the resource implementation.
type LsnclientNetwork6BindingResource struct {
	client *service.NitroClient
}

func (r *LsnclientNetwork6BindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnclientNetwork6BindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnclient_network6_binding"
}

func (r *LsnclientNetwork6BindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnclientNetwork6BindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnclientNetwork6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnclient_network6_binding resource")
	lsnclient_network6_binding := lsnclient_network6_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Lsnclient_network6_binding.Type(), &lsnclient_network6_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnclient_network6_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lsnclient_network6_binding resource")

	// Set ID for the resource before reading state.
	// td is excluded: the NITRO GET response for this binding does not echo td,
	// so including it in the composite ID would break the read-back match loop.
	// This also matches the SDK v2 ID order (clientname,network6) in
	// resource_id_mapping.json.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("clientname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Clientname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("network6:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Network6.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readLsnclientNetwork6BindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "lsnclient_network6_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnclientNetwork6BindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnclientNetwork6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnclient_network6_binding resource")

	r.readLsnclientNetwork6BindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LsnclientNetwork6BindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsnclientNetwork6BindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lsnclient_network6_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		lsnclient_network6_binding := lsnclient_network6_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Lsnclient_network6_binding.Type(), &lsnclient_network6_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnclient_network6_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated lsnclient_network6_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lsnclient_network6_binding resource, skipping update")
	}

	// Read the updated state back
	r.readLsnclientNetwork6BindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "lsnclient_network6_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnclientNetwork6BindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnclientNetwork6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnclient_network6_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"clientname", "network6"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	clientname_value, ok := idMap["clientname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'clientname' not found in ID")
		return
	}

	// network6 values contain '/' and ':' (e.g. "2001:db8:5001::/96") which are
	// not URL-escaped by the NITRO client's args path; escape them explicitly
	// (matches the SDK v2 resource which used url.PathEscape on network6).
	args := make([]string, 0)
	if val, ok := idMap["network6"]; ok && val != "" {
		args = append(args, fmt.Sprintf("network6:%s", url.QueryEscape(val)))
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() && data.Td.ValueInt64() != 0 {
		args = append(args, fmt.Sprintf("td:%d", data.Td.ValueInt64()))
	}

	err = r.client.DeleteResourceWithArgs(service.Lsnclient_network6_binding.Type(), clientname_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsnclient_network6_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsnclient_network6_binding binding")
}

// Helper function to read lsnclient_network6_binding data from API
func (r *LsnclientNetwork6BindingResource) readLsnclientNetwork6BindingFromApi(ctx context.Context, data *LsnclientNetwork6BindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"clientname", "network6"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	clientname_Name, ok := idMap["clientname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'clientname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Lsnclient_network6_binding.Type(),
		ResourceName:             clientname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnclient_network6_binding, got error: %s", err))
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

		// Check network6
		if idVal, ok := idMap["network6"]; ok {
			if val, ok := v["network6"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["network6"].(string); ok {
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

	lsnclient_network6_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
