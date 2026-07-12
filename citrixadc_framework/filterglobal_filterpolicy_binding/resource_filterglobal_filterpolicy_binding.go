package filterglobal_filterpolicy_binding

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
var _ resource.Resource = &FilterglobalFilterpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*FilterglobalFilterpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*FilterglobalFilterpolicyBindingResource)(nil)

func NewFilterglobalFilterpolicyBindingResource() resource.Resource {
	return &FilterglobalFilterpolicyBindingResource{}
}

// FilterglobalFilterpolicyBindingResource defines the resource implementation.
type FilterglobalFilterpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *FilterglobalFilterpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *FilterglobalFilterpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_filterglobal_filterpolicy_binding"
}

func (r *FilterglobalFilterpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *FilterglobalFilterpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data FilterglobalFilterpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating filterglobal_filterpolicy_binding resource")
	filterglobal_filterpolicy_binding := filterglobal_filterpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Global binding resource - SDK v2 used UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Filterglobal_filterpolicy_binding.Type(), &filterglobal_filterpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create filterglobal_filterpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created filterglobal_filterpolicy_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readFilterglobalFilterpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "filterglobal_filterpolicy_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FilterglobalFilterpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data FilterglobalFilterpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading filterglobal_filterpolicy_binding resource")

	r.readFilterglobalFilterpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *FilterglobalFilterpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state FilterglobalFilterpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating filterglobal_filterpolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		filterglobal_filterpolicy_binding := filterglobal_filterpolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Global binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Filterglobal_filterpolicy_binding.Type(), &filterglobal_filterpolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update filterglobal_filterpolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated filterglobal_filterpolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for filterglobal_filterpolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readFilterglobalFilterpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "filterglobal_filterpolicy_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FilterglobalFilterpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FilterglobalFilterpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting filterglobal_filterpolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgsMap with empty resource name
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"policyname"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	// URL-encode arg values (matches SDK v2 behavior) so slashy/special values are
	// transmitted safely in the delete query string.
	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["policyname"]; ok && val != "" {
		argsMap["policyname"] = url.QueryEscape(val)
	}

	err = r.client.DeleteResourceWithArgsMap(service.Filterglobal_filterpolicy_binding.Type(), "", argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete filterglobal_filterpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted filterglobal_filterpolicy_binding binding")
}

// Helper function to read filterglobal_filterpolicy_binding data from API
func (r *FilterglobalFilterpolicyBindingResource) readFilterglobalFilterpolicyBindingFromApi(ctx context.Context, data *FilterglobalFilterpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Array filter without parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"policyname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	var dataArr []map[string]interface{}
	findParams := service.FindParams{
		ResourceType:             service.Filterglobal_filterpolicy_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read filterglobal_filterpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		// Binding no longer exists on the ADC. Signal removal via a null Id
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
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	filterglobal_filterpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
