package appfwprofile_sqlinjection_binding

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
var _ resource.Resource = &AppfwprofileSqlinjectionBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileSqlinjectionBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileSqlinjectionBindingResource)(nil)

func NewAppfwprofileSqlinjectionBindingResource() resource.Resource {
	return &AppfwprofileSqlinjectionBindingResource{}
}

// AppfwprofileSqlinjectionBindingResource defines the resource implementation.
type AppfwprofileSqlinjectionBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileSqlinjectionBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileSqlinjectionBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_sqlinjection_binding"
}

func (r *AppfwprofileSqlinjectionBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileSqlinjectionBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileSqlinjectionBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_sqlinjection_binding resource")
	appfwprofile_sqlinjection_binding := appfwprofile_sqlinjection_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwprofile_sqlinjection_binding.Type(), &appfwprofile_sqlinjection_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_sqlinjection_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile_sqlinjection_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("as_scan_location_sql:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsScanLocationSql.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_value_expr_sql:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueExprSql.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_value_type_sql:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsValueTypeSql.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("formactionurl_sql:%s", utils.UrlEncode(fmt.Sprintf("%v", data.FormactionurlSql.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("sqlinjection:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Sqlinjection.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAppfwprofileSqlinjectionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileSqlinjectionBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileSqlinjectionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_sqlinjection_binding resource")

	r.readAppfwprofileSqlinjectionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileSqlinjectionBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileSqlinjectionBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwprofile_sqlinjection_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		appfwprofile_sqlinjection_binding := appfwprofile_sqlinjection_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Appfwprofile_sqlinjection_binding.Type(), &appfwprofile_sqlinjection_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_sqlinjection_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appfwprofile_sqlinjection_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appfwprofile_sqlinjection_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAppfwprofileSqlinjectionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileSqlinjectionBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileSqlinjectionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_sqlinjection_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "sqlinjection", "formactionurl_sql", "as_scan_location_sql", "as_value_type_sql", "as_value_expr_sql", "ruletype"}, []string{"as_value_type_sql", "as_value_expr_sql", "ruletype"})
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
	if val, ok := idMap["as_scan_location_sql"]; ok && val != "" {
		argsMap["as_scan_location_sql"] = url.QueryEscape(val)
	}
	if val, ok := idMap["as_value_expr_sql"]; ok && val != "" {
		argsMap["as_value_expr_sql"] = url.QueryEscape(val)
	}
	if val, ok := idMap["as_value_type_sql"]; ok && val != "" {
		argsMap["as_value_type_sql"] = url.QueryEscape(val)
	}
	if val, ok := idMap["formactionurl_sql"]; ok && val != "" {
		argsMap["formactionurl_sql"] = url.QueryEscape(val)
	}
	if val, ok := idMap["sqlinjection"]; ok && val != "" {
		argsMap["sqlinjection"] = url.QueryEscape(val)
	}
	// ruletype is not part of the composite ID; pull it from state and include it in
	// the delete args (mirrors SDK v2 + the NITRO delete args list).
	if !data.Ruletype.IsNull() && data.Ruletype.ValueString() != "" {
		argsMap["ruletype"] = url.QueryEscape(data.Ruletype.ValueString())
	}

	err = r.client.DeleteResourceWithArgsMap(service.Appfwprofile_sqlinjection_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile_sqlinjection_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile_sqlinjection_binding binding")
}

// Helper function to read appfwprofile_sqlinjection_binding data from API
func (r *AppfwprofileSqlinjectionBindingResource) readAppfwprofileSqlinjectionBindingFromApi(ctx context.Context, data *AppfwprofileSqlinjectionBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "sqlinjection", "formactionurl_sql", "as_scan_location_sql", "as_value_type_sql", "as_value_expr_sql", "ruletype"}, []string{"as_value_type_sql", "as_value_expr_sql", "ruletype"})
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
		ResourceType:             service.Appfwprofile_sqlinjection_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_sqlinjection_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "appfwprofile_sqlinjection_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check as_scan_location_sql
		if idVal, ok := idMap["as_scan_location_sql"]; ok {
			if val, ok := v["as_scan_location_sql"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_scan_location_sql"].(string); ok {
			match = false
			continue
		}

		// Check as_value_expr_sql
		if idVal, ok := idMap["as_value_expr_sql"]; ok {
			if val, ok := v["as_value_expr_sql"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_value_expr_sql"].(string); ok {
			match = false
			continue
		}

		// Check as_value_type_sql
		if idVal, ok := idMap["as_value_type_sql"]; ok {
			if val, ok := v["as_value_type_sql"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_value_type_sql"].(string); ok {
			match = false
			continue
		}

		// Check formactionurl_sql
		if idVal, ok := idMap["formactionurl_sql"]; ok {
			if val, ok := v["formactionurl_sql"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["formactionurl_sql"].(string); ok {
			match = false
			continue
		}

		// Check sqlinjection
		if idVal, ok := idMap["sqlinjection"]; ok {
			if val, ok := v["sqlinjection"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["sqlinjection"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("appfwprofile_sqlinjection_binding not found with the provided ID attributes"))
		return
	}

	appfwprofile_sqlinjection_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
