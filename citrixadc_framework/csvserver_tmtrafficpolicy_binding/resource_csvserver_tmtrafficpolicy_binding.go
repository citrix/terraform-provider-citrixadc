package csvserver_tmtrafficpolicy_binding

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
var _ resource.Resource = &CsvserverTmtrafficpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*CsvserverTmtrafficpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CsvserverTmtrafficpolicyBindingResource)(nil)

func NewCsvserverTmtrafficpolicyBindingResource() resource.Resource {
	return &CsvserverTmtrafficpolicyBindingResource{}
}

// CsvserverTmtrafficpolicyBindingResource defines the resource implementation.
type CsvserverTmtrafficpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *CsvserverTmtrafficpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CsvserverTmtrafficpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_csvserver_tmtrafficpolicy_binding"
}

func (r *CsvserverTmtrafficpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CsvserverTmtrafficpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CsvserverTmtrafficpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating csvserver_tmtrafficpolicy_binding resource")
	csvserver_tmtrafficpolicy_binding := csvserver_tmtrafficpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - NITRO `add` is POST; mirrors SDK v2 AddResource (Pattern 1, POST-where-PUT-emitted)
	_, err := r.client.AddResource(service.Csvserver_tmtrafficpolicy_binding.Type(), "", &csvserver_tmtrafficpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create csvserver_tmtrafficpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created csvserver_tmtrafficpolicy_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readCsvserverTmtrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverTmtrafficpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CsvserverTmtrafficpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading csvserver_tmtrafficpolicy_binding resource")

	r.readCsvserverTmtrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverTmtrafficpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CsvserverTmtrafficpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating csvserver_tmtrafficpolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		csvserver_tmtrafficpolicy_binding := csvserver_tmtrafficpolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Csvserver_tmtrafficpolicy_binding.Type(), &csvserver_tmtrafficpolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update csvserver_tmtrafficpolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated csvserver_tmtrafficpolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for csvserver_tmtrafficpolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readCsvserverTmtrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverTmtrafficpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CsvserverTmtrafficpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting csvserver_tmtrafficpolicy_binding resource")
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

	// URL-encode delete arg values (slashy/special policy names, bindpoints) — the NITRO
	// client puts these directly into the `args=key:value` query string without encoding,
	// mirroring SDK v2's url.QueryEscape on each arg value (Pattern (b)).
	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["policyname"]; ok && val != "" {
		argsMap["policyname"] = url.QueryEscape(val)
	}
	if !data.Bindpoint.IsNull() && data.Bindpoint.ValueString() != "" {
		argsMap["bindpoint"] = url.QueryEscape(data.Bindpoint.ValueString())
	}
	if !data.Priority.IsNull() {
		argsMap["priority"] = url.QueryEscape(fmt.Sprintf("%v", data.Priority.ValueInt64()))
	}

	err = r.client.DeleteResourceWithArgsMap(service.Csvserver_tmtrafficpolicy_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete csvserver_tmtrafficpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted csvserver_tmtrafficpolicy_binding binding")
}

// Helper function to read csvserver_tmtrafficpolicy_binding data from API
func (r *CsvserverTmtrafficpolicyBindingResource) readCsvserverTmtrafficpolicyBindingFromApi(ctx context.Context, data *CsvserverTmtrafficpolicyBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Csvserver_tmtrafficpolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read csvserver_tmtrafficpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "csvserver_tmtrafficpolicy_binding returned empty array.")
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
		diags.AddError("Client Error", fmt.Sprintf("csvserver_tmtrafficpolicy_binding not found with the provided ID attributes"))
		return
	}

	csvserver_tmtrafficpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
