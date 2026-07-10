package appfwprofile_cmdinjection_binding

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
var _ resource.Resource = &AppfwprofileCmdinjectionBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileCmdinjectionBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileCmdinjectionBindingResource)(nil)

func NewAppfwprofileCmdinjectionBindingResource() resource.Resource {
	return &AppfwprofileCmdinjectionBindingResource{}
}

// AppfwprofileCmdinjectionBindingResource defines the resource implementation.
type AppfwprofileCmdinjectionBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileCmdinjectionBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileCmdinjectionBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_cmdinjection_binding"
}

func (r *AppfwprofileCmdinjectionBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileCmdinjectionBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileCmdinjectionBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_cmdinjection_binding resource")
	appfwprofile_cmdinjection_binding := appfwprofile_cmdinjection_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwprofile_cmdinjection_binding.Type(), &appfwprofile_cmdinjection_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_cmdinjection_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile_cmdinjection_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("as_scan_location_cmd:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsScanLocationCmd.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_value_expr_cmd:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueExprCmd.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_value_type_cmd:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueTypeCmd.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("cmdinjection:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Cmdinjection.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("formactionurl_cmd:%s", utils.UrlEncode(fmt.Sprintf("%v", data.FormactionurlCmd.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAppfwprofileCmdinjectionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileCmdinjectionBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileCmdinjectionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_cmdinjection_binding resource")

	r.readAppfwprofileCmdinjectionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileCmdinjectionBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileCmdinjectionBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwprofile_cmdinjection_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		appfwprofile_cmdinjection_binding := appfwprofile_cmdinjection_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Appfwprofile_cmdinjection_binding.Type(), &appfwprofile_cmdinjection_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_cmdinjection_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appfwprofile_cmdinjection_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appfwprofile_cmdinjection_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAppfwprofileCmdinjectionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileCmdinjectionBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileCmdinjectionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_cmdinjection_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "cmdinjection", "formactionurl_cmd", "as_scan_location_cmd", "as_value_type_cmd", "as_value_expr_cmd"}, []string{"as_value_type_cmd", "as_value_expr_cmd"})
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	// The NITRO client places these arg values directly into the request URL
	// (?args=key:value) without URL-encoding, so values containing special
	// characters (e.g. formactionurl_cmd, the expr/type values) must be
	// pre-escaped here, mirroring the SDK v2 implementation. Without this the
	// list filter fails to match and Delete silently no-ops.
	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["as_scan_location_cmd"]; ok && val != "" {
		argsMap["as_scan_location_cmd"] = val
	}
	if val, ok := idMap["as_value_expr_cmd"]; ok && val != "" {
		argsMap["as_value_expr_cmd"] = url.QueryEscape(val)
	}
	if val, ok := idMap["as_value_type_cmd"]; ok && val != "" {
		argsMap["as_value_type_cmd"] = url.QueryEscape(val)
	}
	if val, ok := idMap["cmdinjection"]; ok && val != "" {
		argsMap["cmdinjection"] = val
	}
	if val, ok := idMap["formactionurl_cmd"]; ok && val != "" {
		argsMap["formactionurl_cmd"] = url.QueryEscape(val)
	}
	if !data.Ruletype.IsNull() && data.Ruletype.ValueString() != "" {
		argsMap["ruletype"] = url.QueryEscape(data.Ruletype.ValueString())
	}

	err = r.client.DeleteResourceWithArgsMap(service.Appfwprofile_cmdinjection_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile_cmdinjection_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile_cmdinjection_binding binding")
}

// Helper function to read appfwprofile_cmdinjection_binding data from API
func (r *AppfwprofileCmdinjectionBindingResource) readAppfwprofileCmdinjectionBindingFromApi(ctx context.Context, data *AppfwprofileCmdinjectionBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "cmdinjection", "formactionurl_cmd", "as_scan_location_cmd", "as_value_type_cmd", "as_value_expr_cmd"}, []string{"as_value_type_cmd", "as_value_expr_cmd"})
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
		ResourceType:             service.Appfwprofile_cmdinjection_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_cmdinjection_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "appfwprofile_cmdinjection_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check as_scan_location_cmd
		if idVal, ok := idMap["as_scan_location_cmd"]; ok {
			if val, ok := v["as_scan_location_cmd"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_scan_location_cmd"].(string); ok {
			match = false
			continue
		}

		// Check as_value_expr_cmd
		if idVal, ok := idMap["as_value_expr_cmd"]; ok {
			if val, ok := v["as_value_expr_cmd"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_value_expr_cmd"].(string); ok {
			match = false
			continue
		}

		// Check as_value_type_cmd
		if idVal, ok := idMap["as_value_type_cmd"]; ok {
			if val, ok := v["as_value_type_cmd"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_value_type_cmd"].(string); ok {
			match = false
			continue
		}

		// Check cmdinjection
		if idVal, ok := idMap["cmdinjection"]; ok {
			if val, ok := v["cmdinjection"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["cmdinjection"].(string); ok {
			match = false
			continue
		}

		// Check formactionurl_cmd
		if idVal, ok := idMap["formactionurl_cmd"]; ok {
			if val, ok := v["formactionurl_cmd"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["formactionurl_cmd"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("appfwprofile_cmdinjection_binding not found with the provided ID attributes"))
		return
	}

	appfwprofile_cmdinjection_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
