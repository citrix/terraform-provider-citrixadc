package appflowpolicylabel_appflowpolicy_binding

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
var _ resource.Resource = &AppflowpolicylabelAppflowpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AppflowpolicylabelAppflowpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppflowpolicylabelAppflowpolicyBindingResource)(nil)

func NewAppflowpolicylabelAppflowpolicyBindingResource() resource.Resource {
	return &AppflowpolicylabelAppflowpolicyBindingResource{}
}

// AppflowpolicylabelAppflowpolicyBindingResource defines the resource implementation.
type AppflowpolicylabelAppflowpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AppflowpolicylabelAppflowpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppflowpolicylabelAppflowpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appflowpolicylabel_appflowpolicy_binding"
}

func (r *AppflowpolicylabelAppflowpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppflowpolicylabelAppflowpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppflowpolicylabelAppflowpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appflowpolicylabel_appflowpolicy_binding resource")
	appflowpolicylabel_appflowpolicy_binding := appflowpolicylabel_appflowpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appflowpolicylabel_appflowpolicy_binding.Type(), &appflowpolicylabel_appflowpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appflowpolicylabel_appflowpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appflowpolicylabel_appflowpolicy_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("labelname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Labelname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAppflowpolicylabelAppflowpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "appflowpolicylabel_appflowpolicy_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowpolicylabelAppflowpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppflowpolicylabelAppflowpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appflowpolicylabel_appflowpolicy_binding resource")

	r.readAppflowpolicylabelAppflowpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppflowpolicylabelAppflowpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppflowpolicylabelAppflowpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appflowpolicylabel_appflowpolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		appflowpolicylabel_appflowpolicy_binding := appflowpolicylabel_appflowpolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Appflowpolicylabel_appflowpolicy_binding.Type(), &appflowpolicylabel_appflowpolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appflowpolicylabel_appflowpolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appflowpolicylabel_appflowpolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appflowpolicylabel_appflowpolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAppflowpolicylabelAppflowpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "appflowpolicylabel_appflowpolicy_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowpolicylabelAppflowpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppflowpolicylabelAppflowpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appflowpolicylabel_appflowpolicy_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"labelname", "policyname"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	labelname_value, ok := idMap["labelname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'labelname' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["policyname"]; ok && val != "" {
		argsMap["policyname"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Appflowpolicylabel_appflowpolicy_binding.Type(), labelname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appflowpolicylabel_appflowpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appflowpolicylabel_appflowpolicy_binding binding")
}

// Helper function to read appflowpolicylabel_appflowpolicy_binding data from API
func (r *AppflowpolicylabelAppflowpolicyBindingResource) readAppflowpolicylabelAppflowpolicyBindingFromApi(ctx context.Context, data *AppflowpolicylabelAppflowpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"labelname", "policyname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	labelname_Name, ok := idMap["labelname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'labelname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Appflowpolicylabel_appflowpolicy_binding.Type(),
		ResourceName:             labelname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appflowpolicylabel_appflowpolicy_binding, got error: %s", err))
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

		// Check policyname
		if idVal, ok := idMap["policyname"]; ok {
			if val, ok := v["policyname"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["policyname"].(string); ok {
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

	appflowpolicylabel_appflowpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
