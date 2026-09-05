package lbvserver_servicegroup_binding

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
var _ resource.Resource = &LbvserverServicegroupBindingResource{}
var _ resource.ResourceWithConfigure = (*LbvserverServicegroupBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbvserverServicegroupBindingResource)(nil)

func NewLbvserverServicegroupBindingResource() resource.Resource {
	return &LbvserverServicegroupBindingResource{}
}

// LbvserverServicegroupBindingResource defines the resource implementation.
type LbvserverServicegroupBindingResource struct {
	client *service.NitroClient
}

func (r *LbvserverServicegroupBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbvserverServicegroupBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbvserver_servicegroup_binding"
}

func (r *LbvserverServicegroupBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbvserverServicegroupBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbvserverServicegroupBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbvserver_servicegroup_binding resource")
	lbvserver_servicegroup_binding := lbvserver_servicegroup_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Lbvserver_servicegroup_binding.Type(), &lbvserver_servicegroup_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbvserver_servicegroup_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lbvserver_servicegroup_binding resource")

	// Set ID for the resource before reading state.
	// Legacy SDK v2 ID order is "name,servicegroupname" (see resource_id_mapping.json).
	// servicename is server-derived/optional and is NOT part of the resource identity.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readLbvserverServicegroupBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "lbvserver_servicegroup_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverServicegroupBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbvserverServicegroupBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbvserver_servicegroup_binding resource")

	r.readLbvserverServicegroupBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LbvserverServicegroupBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LbvserverServicegroupBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lbvserver_servicegroup_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Order.Equal(state.Order) {
		tflog.Debug(ctx, fmt.Sprintf("order has changed for lbvserver_servicegroup_binding"))
		hasChange = true
	}
	if !data.Servicename.Equal(state.Servicename) {
		tflog.Debug(ctx, fmt.Sprintf("servicename has changed for lbvserver_servicegroup_binding"))
		hasChange = true
	}
	if !data.Weight.Equal(state.Weight) {
		tflog.Debug(ctx, fmt.Sprintf("weight has changed for lbvserver_servicegroup_binding"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		lbvserver_servicegroup_binding := lbvserver_servicegroup_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Lbvserver_servicegroup_binding.Type(), &lbvserver_servicegroup_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbvserver_servicegroup_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated lbvserver_servicegroup_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lbvserver_servicegroup_binding resource, skipping update")
	}

	// Read the updated state back
	r.readLbvserverServicegroupBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "lbvserver_servicegroup_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverServicegroupBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbvserverServicegroupBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbvserver_servicegroup_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "servicegroupname"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["servicegroupname"]; ok && val != "" {
		argsMap["servicegroupname"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Lbvserver_servicegroup_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lbvserver_servicegroup_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lbvserver_servicegroup_binding binding")
}

// Helper function to read lbvserver_servicegroup_binding data from API
func (r *LbvserverServicegroupBindingResource) readLbvserverServicegroupBindingFromApi(ctx context.Context, data *LbvserverServicegroupBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "servicegroupname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	name_Name, ok := idMap["name"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'name' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Lbvserver_servicegroup_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbvserver_servicegroup_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
		// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check servicegroupname (the only identity key besides the parent name)
		if idVal, ok := idMap["servicegroupname"]; ok {
			if val, ok := v["servicegroupname"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		// Binding not present in the returned set: signal removal via a null Id (see above).
		data.Id = types.StringNull()
		return
	}

	lbvserver_servicegroup_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
