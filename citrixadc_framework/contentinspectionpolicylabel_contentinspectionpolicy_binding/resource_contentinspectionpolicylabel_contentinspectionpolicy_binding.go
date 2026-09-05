package contentinspectionpolicylabel_contentinspectionpolicy_binding

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
var _ resource.Resource = &ContentinspectionpolicylabelContentinspectionpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*ContentinspectionpolicylabelContentinspectionpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ContentinspectionpolicylabelContentinspectionpolicyBindingResource)(nil)

func NewContentinspectionpolicylabelContentinspectionpolicyBindingResource() resource.Resource {
	return &ContentinspectionpolicylabelContentinspectionpolicyBindingResource{}
}

// ContentinspectionpolicylabelContentinspectionpolicyBindingResource defines the resource implementation.
type ContentinspectionpolicylabelContentinspectionpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *ContentinspectionpolicylabelContentinspectionpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ContentinspectionpolicylabelContentinspectionpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_contentinspectionpolicylabel_contentinspectionpolicy_binding"
}

func (r *ContentinspectionpolicylabelContentinspectionpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ContentinspectionpolicylabelContentinspectionpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ContentinspectionpolicylabelContentinspectionpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating contentinspectionpolicylabel_contentinspectionpolicy_binding resource")
	contentinspectionpolicylabel_contentinspectionpolicy_binding := contentinspectionpolicylabel_contentinspectionpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Contentinspectionpolicylabel_contentinspectionpolicy_binding.Type(), &contentinspectionpolicylabel_contentinspectionpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create contentinspectionpolicylabel_contentinspectionpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created contentinspectionpolicylabel_contentinspectionpolicy_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("labelname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Labelname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readContentinspectionpolicylabelContentinspectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "contentinspectionpolicylabel_contentinspectionpolicy_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionpolicylabelContentinspectionpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ContentinspectionpolicylabelContentinspectionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading contentinspectionpolicylabel_contentinspectionpolicy_binding resource")

	r.readContentinspectionpolicylabelContentinspectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *ContentinspectionpolicylabelContentinspectionpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ContentinspectionpolicylabelContentinspectionpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating contentinspectionpolicylabel_contentinspectionpolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		contentinspectionpolicylabel_contentinspectionpolicy_binding := contentinspectionpolicylabel_contentinspectionpolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Contentinspectionpolicylabel_contentinspectionpolicy_binding.Type(), &contentinspectionpolicylabel_contentinspectionpolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update contentinspectionpolicylabel_contentinspectionpolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated contentinspectionpolicylabel_contentinspectionpolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for contentinspectionpolicylabel_contentinspectionpolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readContentinspectionpolicylabelContentinspectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "contentinspectionpolicylabel_contentinspectionpolicy_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionpolicylabelContentinspectionpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ContentinspectionpolicylabelContentinspectionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting contentinspectionpolicylabel_contentinspectionpolicy_binding resource")
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
	// SDK v2 contract: delete args include priority to disambiguate the binding.
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		argsMap["priority"] = fmt.Sprintf("%d", data.Priority.ValueInt64())
	}

	err = r.client.DeleteResourceWithArgsMap(service.Contentinspectionpolicylabel_contentinspectionpolicy_binding.Type(), labelname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete contentinspectionpolicylabel_contentinspectionpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted contentinspectionpolicylabel_contentinspectionpolicy_binding binding")
}

// Helper function to read contentinspectionpolicylabel_contentinspectionpolicy_binding data from API
func (r *ContentinspectionpolicylabelContentinspectionpolicyBindingResource) readContentinspectionpolicylabelContentinspectionpolicyBindingFromApi(ctx context.Context, data *ContentinspectionpolicylabelContentinspectionpolicyBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Contentinspectionpolicylabel_contentinspectionpolicy_binding.Type(),
		ResourceName:             labelname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read contentinspectionpolicylabel_contentinspectionpolicy_binding, got error: %s", err))
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

	//  Resource is missing
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	contentinspectionpolicylabel_contentinspectionpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
