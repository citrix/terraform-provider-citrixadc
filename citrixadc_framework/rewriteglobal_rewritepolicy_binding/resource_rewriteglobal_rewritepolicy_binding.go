package rewriteglobal_rewritepolicy_binding

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
var _ resource.Resource = &RewriteglobalRewritepolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*RewriteglobalRewritepolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*RewriteglobalRewritepolicyBindingResource)(nil)

func NewRewriteglobalRewritepolicyBindingResource() resource.Resource {
	return &RewriteglobalRewritepolicyBindingResource{}
}

// RewriteglobalRewritepolicyBindingResource defines the resource implementation.
type RewriteglobalRewritepolicyBindingResource struct {
	client *service.NitroClient
}

func (r *RewriteglobalRewritepolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RewriteglobalRewritepolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rewriteglobal_rewritepolicy_binding"
}

func (r *RewriteglobalRewritepolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RewriteglobalRewritepolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RewriteglobalRewritepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rewriteglobal_rewritepolicy_binding resource")
	rewriteglobal_rewritepolicy_binding := rewriteglobal_rewritepolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - SDK v2 used AddResource (POST); NITRO add verb is POST
	_, err := r.client.AddResource(service.Rewriteglobal_rewritepolicy_binding.Type(), "", &rewriteglobal_rewritepolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create rewriteglobal_rewritepolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created rewriteglobal_rewritepolicy_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("priority:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Priority.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Type.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readRewriteglobalRewritepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "rewriteglobal_rewritepolicy_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RewriteglobalRewritepolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RewriteglobalRewritepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading rewriteglobal_rewritepolicy_binding resource")

	r.readRewriteglobalRewritepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *RewriteglobalRewritepolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state RewriteglobalRewritepolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating rewriteglobal_rewritepolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		rewriteglobal_rewritepolicy_binding := rewriteglobal_rewritepolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Rewriteglobal_rewritepolicy_binding.Type(), &rewriteglobal_rewritepolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update rewriteglobal_rewritepolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated rewriteglobal_rewritepolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for rewriteglobal_rewritepolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readRewriteglobalRewritepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "rewriteglobal_rewritepolicy_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RewriteglobalRewritepolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RewriteglobalRewritepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting rewriteglobal_rewritepolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Multiple unique attributes - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"policyname", "priority", "type"}, nil)
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
	if val, ok := idMap["priority"]; ok && val != "" {
		argsMap["priority"] = url.QueryEscape(val)
	}
	if val, ok := idMap["type"]; ok && val != "" {
		argsMap["type"] = url.QueryEscape(val)
	}

	err = r.client.DeleteResourceWithArgsMap(service.Rewriteglobal_rewritepolicy_binding.Type(), "", argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete rewriteglobal_rewritepolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted rewriteglobal_rewritepolicy_binding binding")
}

// Helper function to read rewriteglobal_rewritepolicy_binding data from API
func (r *RewriteglobalRewritepolicyBindingResource) readRewriteglobalRewritepolicyBindingFromApi(ctx context.Context, data *RewriteglobalRewritepolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 3: Array filter without parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"policyname", "priority", "type"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	var dataArr []map[string]interface{}
	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["type"]; ok && val != "" {
		argsMap["type"] = val
	}

	findParams := service.FindParams{
		ResourceType:             service.Rewriteglobal_rewritepolicy_binding.Type(),
		ArgsMap:                  argsMap,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read rewriteglobal_rewritepolicy_binding, got error: %s", err))
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

		// Check priority
		if idVal, ok := idMap["priority"]; ok {
			if val, ok := v["priority"]; ok {
				val, _ = utils.ConvertToInt64(val)
				idValInt64, _ := strconv.ParseInt(idVal, 10, 64)
				if val != idValInt64 {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["priority"]; ok {
			match = false
			continue
		}
		// Check type
		if val, ok := idMap["type"]; ok && val != "" {
			if v, ok := v["type"]; ok {
				if v.(string) != val {
					match = false
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
		data.Id = types.StringNull()
		return
	}

	rewriteglobal_rewritepolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
