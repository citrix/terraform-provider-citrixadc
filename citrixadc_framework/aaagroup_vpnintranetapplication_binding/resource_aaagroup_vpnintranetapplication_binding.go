package aaagroup_vpnintranetapplication_binding

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
var _ resource.Resource = &AaagroupVpnintranetapplicationBindingResource{}
var _ resource.ResourceWithConfigure = (*AaagroupVpnintranetapplicationBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AaagroupVpnintranetapplicationBindingResource)(nil)

func NewAaagroupVpnintranetapplicationBindingResource() resource.Resource {
	return &AaagroupVpnintranetapplicationBindingResource{}
}

// AaagroupVpnintranetapplicationBindingResource defines the resource implementation.
type AaagroupVpnintranetapplicationBindingResource struct {
	client *service.NitroClient
}

func (r *AaagroupVpnintranetapplicationBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaagroupVpnintranetapplicationBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaagroup_vpnintranetapplication_binding"
}

func (r *AaagroupVpnintranetapplicationBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaagroupVpnintranetapplicationBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaagroupVpnintranetapplicationBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaagroup_vpnintranetapplication_binding resource")
	aaagroup_vpnintranetapplication_binding := aaagroup_vpnintranetapplication_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Aaagroup_vpnintranetapplication_binding.Type(), &aaagroup_vpnintranetapplication_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaagroup_vpnintranetapplication_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created aaagroup_vpnintranetapplication_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("intranetapplication:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Intranetapplication.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAaagroupVpnintranetapplicationBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "aaagroup_vpnintranetapplication_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupVpnintranetapplicationBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaagroupVpnintranetapplicationBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaagroup_vpnintranetapplication_binding resource")

	r.readAaagroupVpnintranetapplicationBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AaagroupVpnintranetapplicationBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AaagroupVpnintranetapplicationBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating aaagroup_vpnintranetapplication_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		aaagroup_vpnintranetapplication_binding := aaagroup_vpnintranetapplication_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Aaagroup_vpnintranetapplication_binding.Type(), &aaagroup_vpnintranetapplication_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaagroup_vpnintranetapplication_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated aaagroup_vpnintranetapplication_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for aaagroup_vpnintranetapplication_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAaagroupVpnintranetapplicationBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "aaagroup_vpnintranetapplication_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupVpnintranetapplicationBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaagroupVpnintranetapplicationBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaagroup_vpnintranetapplication_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"groupname", "intranetapplication"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	groupname_value, ok := idMap["groupname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'groupname' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["intranetapplication"]; ok && val != "" {
		argsMap["intranetapplication"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Aaagroup_vpnintranetapplication_binding.Type(), groupname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete aaagroup_vpnintranetapplication_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted aaagroup_vpnintranetapplication_binding binding")
}

// Helper function to read aaagroup_vpnintranetapplication_binding data from API
func (r *AaagroupVpnintranetapplicationBindingResource) readAaagroupVpnintranetapplicationBindingFromApi(ctx context.Context, data *AaagroupVpnintranetapplicationBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"groupname", "intranetapplication"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	groupname_Name, ok := idMap["groupname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'groupname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Aaagroup_vpnintranetapplication_binding.Type(),
		ResourceName:             groupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaagroup_vpnintranetapplication_binding, got error: %s", err))
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

		// Check intranetapplication
		if idVal, ok := idMap["intranetapplication"]; ok {
			if val, ok := v["intranetapplication"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["intranetapplication"].(string); ok {
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

	aaagroup_vpnintranetapplication_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
