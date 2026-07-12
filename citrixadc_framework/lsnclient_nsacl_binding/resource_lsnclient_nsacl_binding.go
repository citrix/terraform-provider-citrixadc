package lsnclient_nsacl_binding

import (
	"context"
	"fmt"
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
var _ resource.Resource = &LsnclientNsaclBindingResource{}
var _ resource.ResourceWithConfigure = (*LsnclientNsaclBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LsnclientNsaclBindingResource)(nil)

func NewLsnclientNsaclBindingResource() resource.Resource {
	return &LsnclientNsaclBindingResource{}
}

// LsnclientNsaclBindingResource defines the resource implementation.
type LsnclientNsaclBindingResource struct {
	client *service.NitroClient
}

func (r *LsnclientNsaclBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnclientNsaclBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnclient_nsacl_binding"
}

func (r *LsnclientNsaclBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnclientNsaclBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnclientNsaclBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnclient_nsacl_binding resource")
	lsnclient_nsacl_binding := lsnclient_nsacl_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Lsnclient_nsacl_binding.Type(), &lsnclient_nsacl_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnclient_nsacl_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lsnclient_nsacl_binding resource")

	// Set ID for the resource before reading state.
	// Identity is clientname + aclname only (mirrors SDK v2 d.SetId("clientname,aclname")
	// and resource_id_mapping.json). td is excluded - it is a default-able traffic domain
	// that NITRO omits from the GET response, so it must not participate in the identity.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("clientname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Clientname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("aclname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Aclname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readLsnclientNsaclBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "lsnclient_nsacl_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnclientNsaclBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnclientNsaclBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnclient_nsacl_binding resource")

	r.readLsnclientNsaclBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LsnclientNsaclBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsnclientNsaclBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lsnclient_nsacl_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		lsnclient_nsacl_binding := lsnclient_nsacl_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Lsnclient_nsacl_binding.Type(), &lsnclient_nsacl_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnclient_nsacl_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated lsnclient_nsacl_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lsnclient_nsacl_binding resource, skipping update")
	}

	// Read the updated state back
	r.readLsnclientNsaclBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "lsnclient_nsacl_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnclientNsaclBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnclientNsaclBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnclient_nsacl_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"clientname", "aclname"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	clientname_value, ok := idMap["clientname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'clientname' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["aclname"]; ok && val != "" {
		argsMap["aclname"] = val
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() && data.Td.ValueInt64() != 0 {
		argsMap["td"] = fmt.Sprintf("%d", data.Td.ValueInt64())
	}

	err = r.client.DeleteResourceWithArgsMap(service.Lsnclient_nsacl_binding.Type(), clientname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsnclient_nsacl_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsnclient_nsacl_binding binding")
}

// Helper function to read lsnclient_nsacl_binding data from API
func (r *LsnclientNsaclBindingResource) readLsnclientNsaclBindingFromApi(ctx context.Context, data *LsnclientNsaclBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"clientname", "aclname"}, nil)
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
		ResourceType:             service.Lsnclient_nsacl_binding.Type(),
		ResourceName:             clientname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnclient_nsacl_binding, got error: %s", err))
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

		// Check aclname
		if idVal, ok := idMap["aclname"]; ok {
			if val, ok := v["aclname"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["aclname"].(string); ok {
			match = false
			continue
		}

		// td is not part of the identity (default-able, omitted from GET) and is
		// resolved from the matched record via SetAttrFromGet, so it is not filtered.
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

	lsnclient_nsacl_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
