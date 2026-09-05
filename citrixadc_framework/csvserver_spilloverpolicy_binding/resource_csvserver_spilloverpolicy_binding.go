package csvserver_spilloverpolicy_binding

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
var _ resource.Resource = &CsvserverSpilloverpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*CsvserverSpilloverpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CsvserverSpilloverpolicyBindingResource)(nil)

func NewCsvserverSpilloverpolicyBindingResource() resource.Resource {
	return &CsvserverSpilloverpolicyBindingResource{}
}

// CsvserverSpilloverpolicyBindingResource defines the resource implementation.
type CsvserverSpilloverpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *CsvserverSpilloverpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CsvserverSpilloverpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_csvserver_spilloverpolicy_binding"
}

func (r *CsvserverSpilloverpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CsvserverSpilloverpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CsvserverSpilloverpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating csvserver_spilloverpolicy_binding resource")
	csvserver_spilloverpolicy_binding := csvserver_spilloverpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - NITRO add is POST (matches SDK v2 AddResource), use AddResource
	_, err := r.client.AddResource(service.Csvserver_spilloverpolicy_binding.Type(), data.Name.ValueString(), &csvserver_spilloverpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create csvserver_spilloverpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created csvserver_spilloverpolicy_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readCsvserverSpilloverpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "csvserver_spilloverpolicy_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverSpilloverpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CsvserverSpilloverpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading csvserver_spilloverpolicy_binding resource")

	r.readCsvserverSpilloverpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *CsvserverSpilloverpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CsvserverSpilloverpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating csvserver_spilloverpolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		csvserver_spilloverpolicy_binding := csvserver_spilloverpolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Csvserver_spilloverpolicy_binding.Type(), &csvserver_spilloverpolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update csvserver_spilloverpolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated csvserver_spilloverpolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for csvserver_spilloverpolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readCsvserverSpilloverpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "csvserver_spilloverpolicy_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverSpilloverpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CsvserverSpilloverpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting csvserver_spilloverpolicy_binding resource")
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
	if val, ok := idMap["policyname"]; ok && val != "" {
		argsMap["policyname"] = url.QueryEscape(val)
	}
	// Match SDK v2 delete-arg set: include bindpoint and priority when present so the
	// correct binding is disambiguated. URL-encode slashy/special values.
	if !data.Bindpoint.IsNull() && data.Bindpoint.ValueString() != "" {
		argsMap["bindpoint"] = url.QueryEscape(data.Bindpoint.ValueString())
	}
	if !data.Priority.IsNull() {
		argsMap["priority"] = url.QueryEscape(fmt.Sprintf("%v", data.Priority.ValueInt64()))
	}

	err = r.client.DeleteResourceWithArgsMap(service.Csvserver_spilloverpolicy_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete csvserver_spilloverpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted csvserver_spilloverpolicy_binding binding")
}

// Helper function to read csvserver_spilloverpolicy_binding data from API
func (r *CsvserverSpilloverpolicyBindingResource) readCsvserverSpilloverpolicyBindingFromApi(ctx context.Context, data *CsvserverSpilloverpolicyBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Csvserver_spilloverpolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read csvserver_spilloverpolicy_binding, got error: %s", err))
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

	// Binding not present in the returned set: signal removal via a null Id (see above).
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	csvserver_spilloverpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
