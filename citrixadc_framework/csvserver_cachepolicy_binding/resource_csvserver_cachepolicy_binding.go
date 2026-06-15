package csvserver_cachepolicy_binding

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
var _ resource.Resource = &CsvserverCachepolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*CsvserverCachepolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CsvserverCachepolicyBindingResource)(nil)

func NewCsvserverCachepolicyBindingResource() resource.Resource {
	return &CsvserverCachepolicyBindingResource{}
}

// CsvserverCachepolicyBindingResource defines the resource implementation.
type CsvserverCachepolicyBindingResource struct {
	client *service.NitroClient
}

func (r *CsvserverCachepolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CsvserverCachepolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_csvserver_cachepolicy_binding"
}

func (r *CsvserverCachepolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CsvserverCachepolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CsvserverCachepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating csvserver_cachepolicy_binding resource")
	csvserver_cachepolicy_binding := csvserver_cachepolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - SDK v2 used AddResource (POST). NITRO add for this binding is POST. (Pattern 1)
	_, err := r.client.AddResource(service.Csvserver_cachepolicy_binding.Type(), data.Name.ValueString(), &csvserver_cachepolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create csvserver_cachepolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created csvserver_cachepolicy_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bindpoint:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Bindpoint.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readCsvserverCachepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverCachepolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CsvserverCachepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading csvserver_cachepolicy_binding resource")

	r.readCsvserverCachepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverCachepolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CsvserverCachepolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating csvserver_cachepolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		csvserver_cachepolicy_binding := csvserver_cachepolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Csvserver_cachepolicy_binding.Type(), &csvserver_cachepolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update csvserver_cachepolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated csvserver_cachepolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for csvserver_cachepolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readCsvserverCachepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverCachepolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CsvserverCachepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting csvserver_cachepolicy_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "policyname"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["bindpoint"]; ok && val != "" {
		argsMap["bindpoint"] = val
	}
	if val, ok := idMap["policyname"]; ok && val != "" {
		argsMap["policyname"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Csvserver_cachepolicy_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete csvserver_cachepolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted csvserver_cachepolicy_binding binding")
}

// Helper function to read csvserver_cachepolicy_binding data from API
func (r *CsvserverCachepolicyBindingResource) readCsvserverCachepolicyBindingFromApi(ctx context.Context, data *CsvserverCachepolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "policyname"}, nil)
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
		ResourceType:             service.Csvserver_cachepolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read csvserver_cachepolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "csvserver_cachepolicy_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check bindpoint
		if idVal, ok := idMap["bindpoint"]; ok {
			if val, ok := v["bindpoint"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["bindpoint"].(string); ok {
			match = false
			continue
		}

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
		diags.AddError("Client Error", fmt.Sprintf("csvserver_cachepolicy_binding not found with the provided ID attributes"))
		return
	}

	csvserver_cachepolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
