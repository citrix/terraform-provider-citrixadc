package appflowglobal_appflowpolicy_binding

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
var _ resource.Resource = &AppflowglobalAppflowpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AppflowglobalAppflowpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppflowglobalAppflowpolicyBindingResource)(nil)

func NewAppflowglobalAppflowpolicyBindingResource() resource.Resource {
	return &AppflowglobalAppflowpolicyBindingResource{}
}

// AppflowglobalAppflowpolicyBindingResource defines the resource implementation.
type AppflowglobalAppflowpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AppflowglobalAppflowpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppflowglobalAppflowpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appflowglobal_appflowpolicy_binding"
}

func (r *AppflowglobalAppflowpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppflowglobalAppflowpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppflowglobalAppflowpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appflowglobal_appflowpolicy_binding resource")
	appflowglobal_appflowpolicy_binding := appflowglobal_appflowpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appflowglobal_appflowpolicy_binding.Type(), &appflowglobal_appflowpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appflowglobal_appflowpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appflowglobal_appflowpolicy_binding resource")

	// Set ID for the resource before reading state
	// Backward-compatible with SDK v2: ID is the plain policyname value.
	data.Id = types.StringValue(data.Policyname.ValueString())

	// Read the updated state back
	r.readAppflowglobalAppflowpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "appflowglobal_appflowpolicy_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowglobalAppflowpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppflowglobalAppflowpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appflowglobal_appflowpolicy_binding resource")

	r.readAppflowglobalAppflowpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppflowglobalAppflowpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppflowglobalAppflowpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appflowglobal_appflowpolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		appflowglobal_appflowpolicy_binding := appflowglobal_appflowpolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Appflowglobal_appflowpolicy_binding.Type(), &appflowglobal_appflowpolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appflowglobal_appflowpolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appflowglobal_appflowpolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appflowglobal_appflowpolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAppflowglobalAppflowpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "appflowglobal_appflowpolicy_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowglobalAppflowpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppflowglobalAppflowpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appflowglobal_appflowpolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgsMap with empty resource name.
	// ID is the plain policyname; type and priority come from the state model
	// (mirrors SDK v2 delete args).
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"policyname"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["policyname"]; ok && val != "" {
		argsMap["policyname"] = val
	}
	if !data.Type.IsNull() && data.Type.ValueString() != "" {
		argsMap["type"] = data.Type.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		argsMap["priority"] = fmt.Sprintf("%d", data.Priority.ValueInt64())
	}

	err = r.client.DeleteResourceWithArgsMap(service.Appflowglobal_appflowpolicy_binding.Type(), "", argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appflowglobal_appflowpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appflowglobal_appflowpolicy_binding binding")
}

// Helper function to read appflowglobal_appflowpolicy_binding data from API
func (r *AppflowglobalAppflowpolicyBindingResource) readAppflowglobalAppflowpolicyBindingFromApi(ctx context.Context, data *AppflowglobalAppflowpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// ID is the plain policyname value (SDK v2 backward-compatible). The bind point
	// "type" is required as a GET filter for the binding GET to echo back the
	// per-policy records (policyname, priority); without it NITRO returns only an
	// aggregate summary that lacks policyname. Source "type" from the model field.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"policyname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}
	policynameId := idMap["policyname"]

	var dataArr []map[string]interface{}
	var argsMap map[string]string = make(map[string]string)
	if !data.Type.IsNull() && data.Type.ValueString() != "" {
		argsMap["type"] = data.Type.ValueString()
	}

	findParams := service.FindParams{
		ResourceType:             service.Appflowglobal_appflowpolicy_binding.Type(),
		ArgsMap:                  argsMap,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appflowglobal_appflowpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
		// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right policyname (and priority if set)
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check policyname
		if val, ok := v["policyname"].(string); ok {
			if val != policynameId {
				match = false
				continue
			}
		} else {
			match = false
			continue
		}

		// Check priority (if known in state)
		if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
			if val, ok := v["priority"]; ok {
				valInt64, _ := utils.ConvertToInt64(val)
				if valInt64 != data.Priority.ValueInt64() {
					match = false
					continue
				}
			}
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		// Binding not present in the returned set: signal removal via a null Id (see above).
		data.Id = types.StringNull()
		return
	}

	appflowglobal_appflowpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
