package appfwprofile_jsonxssurl_binding

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
var _ resource.Resource = &AppfwprofileJsonxssurlBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileJsonxssurlBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileJsonxssurlBindingResource)(nil)

func NewAppfwprofileJsonxssurlBindingResource() resource.Resource {
	return &AppfwprofileJsonxssurlBindingResource{}
}

// AppfwprofileJsonxssurlBindingResource defines the resource implementation.
type AppfwprofileJsonxssurlBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileJsonxssurlBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileJsonxssurlBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_jsonxssurl_binding"
}

func (r *AppfwprofileJsonxssurlBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileJsonxssurlBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileJsonxssurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_jsonxssurl_binding resource")
	appfwprofile_jsonxssurl_binding := appfwprofile_jsonxssurl_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwprofile_jsonxssurl_binding.Type(), &appfwprofile_jsonxssurl_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_jsonxssurl_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile_jsonxssurl_binding resource")

	// Set ID for the resource before reading state.
	// Order mirrors resource_id_mapping.json: name, jsonxssurl, then the optional
	// keys keyname_json_xss / as_value_type_json_xss / as_value_expr_json_xss. Only
	// compose optional keys that the user actually set (family pattern f) so the Read
	// match loop can distinguish "no keyname binding" from a keyname binding.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("jsonxssurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Jsonxssurl.ValueString()))))
	if !data.KeynameJsonXss.IsNull() && data.KeynameJsonXss.ValueString() != "" {
		idParts = append(idParts, fmt.Sprintf("keyname_json_xss:%s", utils.UrlEncode(fmt.Sprintf("%v", data.KeynameJsonXss.ValueString()))))
	}
	if !data.AsValueTypeJsonXss.IsNull() && data.AsValueTypeJsonXss.ValueString() != "" {
		idParts = append(idParts, fmt.Sprintf("as_value_type_json_xss:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueTypeJsonXss.ValueString()))))
	}
	if !data.AsValueExprJsonXss.IsNull() && data.AsValueExprJsonXss.ValueString() != "" {
		idParts = append(idParts, fmt.Sprintf("as_value_expr_json_xss:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueExprJsonXss.ValueString()))))
	}
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAppfwprofileJsonxssurlBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_jsonxssurl_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileJsonxssurlBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileJsonxssurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_jsonxssurl_binding resource")

	r.readAppfwprofileJsonxssurlBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwprofileJsonxssurlBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileJsonxssurlBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwprofile_jsonxssurl_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		appfwprofile_jsonxssurl_binding := appfwprofile_jsonxssurl_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Appfwprofile_jsonxssurl_binding.Type(), &appfwprofile_jsonxssurl_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_jsonxssurl_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appfwprofile_jsonxssurl_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appfwprofile_jsonxssurl_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAppfwprofileJsonxssurlBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_jsonxssurl_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileJsonxssurlBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileJsonxssurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_jsonxssurl_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// Legacy order matches resource_id_mapping.json so old SDK v2 comma IDs still parse.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "jsonxssurl", "keyname_json_xss", "as_value_type_json_xss", "as_value_expr_json_xss"}, []string{"keyname_json_xss", "as_value_type_json_xss", "as_value_expr_json_xss"})
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	// Build URL-encoded delete args mirroring the SDK v2 resource. ParseIdString returns
	// decoded values, so each value must be re-encoded for the args=key:value URL
	// (DeleteResourceWithArgs does NOT encode arg values - family pattern b).
	args := make([]string, 0)
	if val, ok := idMap["jsonxssurl"]; ok && val != "" {
		args = append(args, fmt.Sprintf("jsonxssurl:%s", url.QueryEscape(val)))
	}
	if val, ok := idMap["keyname_json_xss"]; ok && val != "" {
		args = append(args, fmt.Sprintf("keyname_json_xss:%s", url.QueryEscape(val)))
	}
	if val, ok := idMap["as_value_type_json_xss"]; ok && val != "" {
		args = append(args, fmt.Sprintf("as_value_type_json_xss:%s", url.QueryEscape(val)))
	}
	if val, ok := idMap["as_value_expr_json_xss"]; ok && val != "" {
		args = append(args, fmt.Sprintf("as_value_expr_json_xss:%s", url.QueryEscape(val)))
	}
	// ruletype is a real delete arg in SDK v2; include it when present in state.
	if !data.Ruletype.IsNull() && data.Ruletype.ValueString() != "" {
		args = append(args, fmt.Sprintf("ruletype:%s", url.QueryEscape(data.Ruletype.ValueString())))
	}

	err = r.client.DeleteResourceWithArgs(service.Appfwprofile_jsonxssurl_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile_jsonxssurl_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile_jsonxssurl_binding binding")
}

// Helper function to read appfwprofile_jsonxssurl_binding data from API
func (r *AppfwprofileJsonxssurlBindingResource) readAppfwprofileJsonxssurlBindingFromApi(ctx context.Context, data *AppfwprofileJsonxssurlBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "jsonxssurl", "keyname_json_xss", "as_value_type_json_xss", "as_value_expr_json_xss"}, []string{"keyname_json_xss", "as_value_type_json_xss", "as_value_expr_json_xss"})
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
		ResourceType:             service.Appfwprofile_jsonxssurl_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_jsonxssurl_binding, got error: %s", err))
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

		// Check as_value_expr_json_xss
		if idVal, ok := idMap["as_value_expr_json_xss"]; ok {
			if val, ok := v["as_value_expr_json_xss"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_value_expr_json_xss"].(string); ok {
			match = false
			continue
		}

		// Check as_value_type_json_xss
		if idVal, ok := idMap["as_value_type_json_xss"]; ok {
			if val, ok := v["as_value_type_json_xss"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_value_type_json_xss"].(string); ok {
			match = false
			continue
		}

		// Check jsonxssurl
		if idVal, ok := idMap["jsonxssurl"]; ok {
			if val, ok := v["jsonxssurl"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["jsonxssurl"].(string); ok {
			match = false
			continue
		}

		// Check keyname_json_xss
		if idVal, ok := idMap["keyname_json_xss"]; ok {
			if val, ok := v["keyname_json_xss"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["keyname_json_xss"].(string); ok {
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

	appfwprofile_jsonxssurl_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
