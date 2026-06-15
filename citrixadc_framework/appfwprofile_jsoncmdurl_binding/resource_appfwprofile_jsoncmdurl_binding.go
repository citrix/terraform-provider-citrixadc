package appfwprofile_jsoncmdurl_binding

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
var _ resource.Resource = &AppfwprofileJsoncmdurlBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileJsoncmdurlBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileJsoncmdurlBindingResource)(nil)

func NewAppfwprofileJsoncmdurlBindingResource() resource.Resource {
	return &AppfwprofileJsoncmdurlBindingResource{}
}

// AppfwprofileJsoncmdurlBindingResource defines the resource implementation.
type AppfwprofileJsoncmdurlBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileJsoncmdurlBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileJsoncmdurlBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_jsoncmdurl_binding"
}

func (r *AppfwprofileJsoncmdurlBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileJsoncmdurlBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileJsoncmdurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_jsoncmdurl_binding resource")
	appfwprofile_jsoncmdurl_binding := appfwprofile_jsoncmdurl_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwprofile_jsoncmdurl_binding.Type(), &appfwprofile_jsoncmdurl_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_jsoncmdurl_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile_jsoncmdurl_binding resource")

	// Set ID for the resource before reading state. Optional unique attrs
	// (keyname_json_cmd, as_value_type_json_cmd, as_value_expr_json_cmd) are
	// included ONLY when non-empty, so the Read match loop and Delete treat an
	// absent key as "the bound record has no such field" (mirrors SDK v2).
	data.Id = types.StringValue(buildAppfwprofileJsoncmdurlBindingId(&data))

	// Read the updated state back
	r.readAppfwprofileJsoncmdurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileJsoncmdurlBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileJsoncmdurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_jsoncmdurl_binding resource")

	r.readAppfwprofileJsoncmdurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileJsoncmdurlBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileJsoncmdurlBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwprofile_jsoncmdurl_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		appfwprofile_jsoncmdurl_binding := appfwprofile_jsoncmdurl_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Appfwprofile_jsoncmdurl_binding.Type(), &appfwprofile_jsoncmdurl_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_jsoncmdurl_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appfwprofile_jsoncmdurl_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appfwprofile_jsoncmdurl_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAppfwprofileJsoncmdurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileJsoncmdurlBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileJsoncmdurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_jsoncmdurl_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "jsoncmdurl", "keyname_json_cmd", "as_value_type_json_cmd", "as_value_expr_json_cmd"}, []string{"keyname_json_cmd", "as_value_type_json_cmd", "as_value_expr_json_cmd"})
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	// NITRO does NOT URL-encode the arg values itself (DeleteResourceWithArgs[Map]
	// joins raw key:value into the ?args= query string), so we must url.QueryEscape
	// each value here — mirroring the SDK v2 resource. jsoncmdurl and the JSON CMD
	// expressions contain regex/URL special characters (^ $ \ / : ?) that would
	// otherwise corrupt the delete URL.
	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["as_value_expr_json_cmd"]; ok && val != "" {
		argsMap["as_value_expr_json_cmd"] = url.QueryEscape(val)
	}
	if val, ok := idMap["as_value_type_json_cmd"]; ok && val != "" {
		argsMap["as_value_type_json_cmd"] = url.QueryEscape(val)
	}
	if val, ok := idMap["jsoncmdurl"]; ok && val != "" {
		argsMap["jsoncmdurl"] = url.QueryEscape(val)
	}
	if val, ok := idMap["keyname_json_cmd"]; ok && val != "" {
		argsMap["keyname_json_cmd"] = url.QueryEscape(val)
	}
	// ruletype is not part of the ID; carry it from state (mirrors SDK v2 delete args).
	if !data.Ruletype.IsNull() && data.Ruletype.ValueString() != "" {
		argsMap["ruletype"] = url.QueryEscape(data.Ruletype.ValueString())
	}

	err = r.client.DeleteResourceWithArgsMap(service.Appfwprofile_jsoncmdurl_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile_jsoncmdurl_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile_jsoncmdurl_binding binding")
}

// buildAppfwprofileJsoncmdurlBindingId composes the composite key:UrlEncode(value) ID.
// The mandatory keys (name, jsoncmdurl) are always present; the optional unique keys
// (keyname_json_cmd, as_value_type_json_cmd, as_value_expr_json_cmd) are included only
// when set, so an absent key signals "the bound record has no such field" to the Read
// match loop and Delete.
func buildAppfwprofileJsoncmdurlBindingId(data *AppfwprofileJsoncmdurlBindingResourceModel) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(data.Name.ValueString())))
	idParts = append(idParts, fmt.Sprintf("jsoncmdurl:%s", utils.UrlEncode(data.Jsoncmdurl.ValueString())))
	if !data.KeynameJsonCmd.IsNull() && !data.KeynameJsonCmd.IsUnknown() && data.KeynameJsonCmd.ValueString() != "" {
		idParts = append(idParts, fmt.Sprintf("keyname_json_cmd:%s", utils.UrlEncode(data.KeynameJsonCmd.ValueString())))
	}
	if !data.AsValueTypeJsonCmd.IsNull() && !data.AsValueTypeJsonCmd.IsUnknown() && data.AsValueTypeJsonCmd.ValueString() != "" {
		idParts = append(idParts, fmt.Sprintf("as_value_type_json_cmd:%s", utils.UrlEncode(data.AsValueTypeJsonCmd.ValueString())))
	}
	if !data.AsValueExprJsonCmd.IsNull() && !data.AsValueExprJsonCmd.IsUnknown() && data.AsValueExprJsonCmd.ValueString() != "" {
		idParts = append(idParts, fmt.Sprintf("as_value_expr_json_cmd:%s", utils.UrlEncode(data.AsValueExprJsonCmd.ValueString())))
	}
	return strings.Join(idParts, ",")
}

// Helper function to read appfwprofile_jsoncmdurl_binding data from API
func (r *AppfwprofileJsoncmdurlBindingResource) readAppfwprofileJsoncmdurlBindingFromApi(ctx context.Context, data *AppfwprofileJsoncmdurlBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "jsoncmdurl", "keyname_json_cmd", "as_value_type_json_cmd", "as_value_expr_json_cmd"}, []string{"keyname_json_cmd", "as_value_type_json_cmd", "as_value_expr_json_cmd"})
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
		ResourceType:             service.Appfwprofile_jsoncmdurl_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_jsoncmdurl_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "appfwprofile_jsoncmdurl_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check as_value_expr_json_cmd
		if idVal, ok := idMap["as_value_expr_json_cmd"]; ok {
			if val, ok := v["as_value_expr_json_cmd"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_value_expr_json_cmd"].(string); ok {
			match = false
			continue
		}

		// Check as_value_type_json_cmd
		if idVal, ok := idMap["as_value_type_json_cmd"]; ok {
			if val, ok := v["as_value_type_json_cmd"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_value_type_json_cmd"].(string); ok {
			match = false
			continue
		}

		// Check jsoncmdurl
		if idVal, ok := idMap["jsoncmdurl"]; ok {
			if val, ok := v["jsoncmdurl"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["jsoncmdurl"].(string); ok {
			match = false
			continue
		}

		// Check keyname_json_cmd
		if idVal, ok := idMap["keyname_json_cmd"]; ok {
			if val, ok := v["keyname_json_cmd"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["keyname_json_cmd"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("appfwprofile_jsoncmdurl_binding not found with the provided ID attributes"))
		return
	}

	appfwprofile_jsoncmdurl_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
