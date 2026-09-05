package clusternodegroup_streamidentifier_binding

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
var _ resource.Resource = &ClusternodegroupStreamidentifierBindingResource{}
var _ resource.ResourceWithConfigure = (*ClusternodegroupStreamidentifierBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ClusternodegroupStreamidentifierBindingResource)(nil)

func NewClusternodegroupStreamidentifierBindingResource() resource.Resource {
	return &ClusternodegroupStreamidentifierBindingResource{}
}

// ClusternodegroupStreamidentifierBindingResource defines the resource implementation.
type ClusternodegroupStreamidentifierBindingResource struct {
	client *service.NitroClient
}

func (r *ClusternodegroupStreamidentifierBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusternodegroupStreamidentifierBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusternodegroup_streamidentifier_binding"
}

func (r *ClusternodegroupStreamidentifierBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusternodegroupStreamidentifierBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusternodegroupStreamidentifierBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating clusternodegroup_streamidentifier_binding resource")
	clusternodegroup_streamidentifier_binding := clusternodegroup_streamidentifier_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Clusternodegroup_streamidentifier_binding.Type(), &clusternodegroup_streamidentifier_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create clusternodegroup_streamidentifier_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created clusternodegroup_streamidentifier_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("identifiername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Identifiername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readClusternodegroupStreamidentifierBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "clusternodegroup_streamidentifier_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupStreamidentifierBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusternodegroupStreamidentifierBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading clusternodegroup_streamidentifier_binding resource")

	r.readClusternodegroupStreamidentifierBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *ClusternodegroupStreamidentifierBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ClusternodegroupStreamidentifierBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating clusternodegroup_streamidentifier_binding resource")

	// SDK v2 resource had no Update: all attributes are ForceNew. No-op update.
	tflog.Debug(ctx, "No updateable attributes for clusternodegroup_streamidentifier_binding resource, skipping update")

	// Read the updated state back
	r.readClusternodegroupStreamidentifierBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "clusternodegroup_streamidentifier_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupStreamidentifierBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusternodegroupStreamidentifierBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting clusternodegroup_streamidentifier_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgsMap
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "identifiername"}, nil)
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
	if val, ok := idMap["identifiername"]; ok && val != "" {
		argsMap["identifiername"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Clusternodegroup_streamidentifier_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete clusternodegroup_streamidentifier_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted clusternodegroup_streamidentifier_binding binding")
}

// Helper function to read clusternodegroup_streamidentifier_binding data from API
func (r *ClusternodegroupStreamidentifierBindingResource) readClusternodegroupStreamidentifierBindingFromApi(ctx context.Context, data *ClusternodegroupStreamidentifierBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "identifiername"}, nil)
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
		ResourceType:             service.Clusternodegroup_streamidentifier_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read clusternodegroup_streamidentifier_binding, got error: %s", err))
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

		// Check identifiername
		if idVal, ok := idMap["identifiername"]; ok {
			if val, ok := v["identifiername"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["identifiername"].(string); ok {
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
		// Binding not present in the returned set: signal removal via a null Id (see above).
		data.Id = types.StringNull()
		return
	}

	clusternodegroup_streamidentifier_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
