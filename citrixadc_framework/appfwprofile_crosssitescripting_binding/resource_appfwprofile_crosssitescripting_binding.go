package appfwprofile_crosssitescripting_binding

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
var _ resource.Resource = &AppfwprofileCrosssitescriptingBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileCrosssitescriptingBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileCrosssitescriptingBindingResource)(nil)

func NewAppfwprofileCrosssitescriptingBindingResource() resource.Resource {
	return &AppfwprofileCrosssitescriptingBindingResource{}
}

// AppfwprofileCrosssitescriptingBindingResource defines the resource implementation.
type AppfwprofileCrosssitescriptingBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileCrosssitescriptingBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileCrosssitescriptingBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_crosssitescripting_binding"
}

func (r *AppfwprofileCrosssitescriptingBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileCrosssitescriptingBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileCrosssitescriptingBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_crosssitescripting_binding resource")
	appfwprofile_crosssitescripting_binding := appfwprofile_crosssitescripting_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwprofile_crosssitescripting_binding.Type(), &appfwprofile_crosssitescripting_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_crosssitescripting_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile_crosssitescripting_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("as_scan_location_xss:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsScanLocationXss.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_value_expr_xss:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueExprXss.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_value_type_xss:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueTypeXss.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("crosssitescripting:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Crosssitescripting.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("formactionurl_xss:%s", utils.UrlEncode(fmt.Sprintf("%v", data.FormactionurlXss.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAppfwprofileCrosssitescriptingBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_crosssitescripting_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileCrosssitescriptingBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileCrosssitescriptingBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_crosssitescripting_binding resource")

	r.readAppfwprofileCrosssitescriptingBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwprofileCrosssitescriptingBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileCrosssitescriptingBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwprofile_crosssitescripting_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		appfwprofile_crosssitescripting_binding := appfwprofile_crosssitescripting_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Appfwprofile_crosssitescripting_binding.Type(), &appfwprofile_crosssitescripting_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_crosssitescripting_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appfwprofile_crosssitescripting_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appfwprofile_crosssitescripting_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAppfwprofileCrosssitescriptingBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "appfwprofile_crosssitescripting_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileCrosssitescriptingBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileCrosssitescriptingBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_crosssitescripting_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "crosssitescripting", "formactionurl_xss", "as_scan_location_xss", "as_value_type_xss", "as_value_expr_xss"}, []string{"as_value_type_xss", "as_value_expr_xss"})
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	// DeleteResourceWithArgsMap does NOT URL-encode the arg values, so values that
	// contain special characters (regex form-action URLs, value expressions) must be
	// encoded here, mirroring the SDK v2 resource (which url.QueryEscape'd them).
	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["as_scan_location_xss"]; ok && val != "" {
		argsMap["as_scan_location_xss"] = url.QueryEscape(val)
	}
	if val, ok := idMap["as_value_expr_xss"]; ok && val != "" {
		argsMap["as_value_expr_xss"] = url.QueryEscape(val)
	}
	if val, ok := idMap["as_value_type_xss"]; ok && val != "" {
		argsMap["as_value_type_xss"] = url.QueryEscape(val)
	}
	if val, ok := idMap["crosssitescripting"]; ok && val != "" {
		argsMap["crosssitescripting"] = url.QueryEscape(val)
	}
	if val, ok := idMap["formactionurl_xss"]; ok && val != "" {
		argsMap["formactionurl_xss"] = url.QueryEscape(val)
	}
	// ruletype is not part of the composite ID; pull it from state and include it in
	// the delete args (mirrors SDK v2 + the NITRO delete args list).
	if !data.Ruletype.IsNull() && data.Ruletype.ValueString() != "" {
		argsMap["ruletype"] = url.QueryEscape(data.Ruletype.ValueString())
	}

	err = r.client.DeleteResourceWithArgsMap(service.Appfwprofile_crosssitescripting_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile_crosssitescripting_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile_crosssitescripting_binding binding")
}

// Helper function to read appfwprofile_crosssitescripting_binding data from API
func (r *AppfwprofileCrosssitescriptingBindingResource) readAppfwprofileCrosssitescriptingBindingFromApi(ctx context.Context, data *AppfwprofileCrosssitescriptingBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "crosssitescripting", "formactionurl_xss", "as_scan_location_xss", "as_value_type_xss", "as_value_expr_xss"}, []string{"as_value_type_xss", "as_value_expr_xss"})
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
		ResourceType:             service.Appfwprofile_crosssitescripting_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_crosssitescripting_binding, got error: %s", err))
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

		// Check as_scan_location_xss
		if idVal, ok := idMap["as_scan_location_xss"]; ok {
			if val, ok := v["as_scan_location_xss"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_scan_location_xss"].(string); ok {
			match = false
			continue
		}

		// Check as_value_expr_xss
		if idVal, ok := idMap["as_value_expr_xss"]; ok {
			if val, ok := v["as_value_expr_xss"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_value_expr_xss"].(string); ok {
			match = false
			continue
		}

		// Check as_value_type_xss
		if idVal, ok := idMap["as_value_type_xss"]; ok {
			if val, ok := v["as_value_type_xss"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_value_type_xss"].(string); ok {
			match = false
			continue
		}

		// Check crosssitescripting
		if idVal, ok := idMap["crosssitescripting"]; ok {
			if val, ok := v["crosssitescripting"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["crosssitescripting"].(string); ok {
			match = false
			continue
		}

		// Check formactionurl_xss
		if idVal, ok := idMap["formactionurl_xss"]; ok {
			if val, ok := v["formactionurl_xss"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["formactionurl_xss"].(string); ok {
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

	appfwprofile_crosssitescripting_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
