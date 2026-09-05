package aaagroup_vpnsecureprivateaccessprofile_binding

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
var _ resource.Resource = &AaagroupVpnsecureprivateaccessprofileBindingResource{}
var _ resource.ResourceWithConfigure = (*AaagroupVpnsecureprivateaccessprofileBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AaagroupVpnsecureprivateaccessprofileBindingResource)(nil)

func NewAaagroupVpnsecureprivateaccessprofileBindingResource() resource.Resource {
	return &AaagroupVpnsecureprivateaccessprofileBindingResource{}
}

// AaagroupVpnsecureprivateaccessprofileBindingResource defines the resource implementation.
type AaagroupVpnsecureprivateaccessprofileBindingResource struct {
	client *service.NitroClient
}

func (r *AaagroupVpnsecureprivateaccessprofileBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaagroupVpnsecureprivateaccessprofileBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaagroup_vpnsecureprivateaccessprofile_binding"
}

func (r *AaagroupVpnsecureprivateaccessprofileBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaagroupVpnsecureprivateaccessprofileBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaagroupVpnsecureprivateaccessprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaagroup_vpnsecureprivateaccessprofile_binding resource")
	aaagroup_vpnsecureprivateaccessprofile_binding := aaagroup_vpnsecureprivateaccessprofile_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Aaagroup_vpnsecureprivateaccessprofile_binding.Type(), &aaagroup_vpnsecureprivateaccessprofile_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaagroup_vpnsecureprivateaccessprofile_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created aaagroup_vpnsecureprivateaccessprofile_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("secureprivateaccessprofile:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Secureprivateaccessprofile.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAaagroupVpnsecureprivateaccessprofileBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "aaagroup_vpnsecureprivateaccessprofile_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupVpnsecureprivateaccessprofileBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaagroupVpnsecureprivateaccessprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaagroup_vpnsecureprivateaccessprofile_binding resource")

	r.readAaagroupVpnsecureprivateaccessprofileBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AaagroupVpnsecureprivateaccessprofileBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AaagroupVpnsecureprivateaccessprofileBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating aaagroup_vpnsecureprivateaccessprofile_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		aaagroup_vpnsecureprivateaccessprofile_binding := aaagroup_vpnsecureprivateaccessprofile_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Aaagroup_vpnsecureprivateaccessprofile_binding.Type(), &aaagroup_vpnsecureprivateaccessprofile_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaagroup_vpnsecureprivateaccessprofile_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated aaagroup_vpnsecureprivateaccessprofile_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for aaagroup_vpnsecureprivateaccessprofile_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAaagroupVpnsecureprivateaccessprofileBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "aaagroup_vpnsecureprivateaccessprofile_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupVpnsecureprivateaccessprofileBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaagroupVpnsecureprivateaccessprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaagroup_vpnsecureprivateaccessprofile_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgsMap
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"groupname", "secureprivateaccessprofile"}, nil)
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
	if val, ok := idMap["secureprivateaccessprofile"]; ok && val != "" {
		argsMap["secureprivateaccessprofile"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Aaagroup_vpnsecureprivateaccessprofile_binding.Type(), groupname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete aaagroup_vpnsecureprivateaccessprofile_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted aaagroup_vpnsecureprivateaccessprofile_binding binding")
}

// Helper function to read aaagroup_vpnsecureprivateaccessprofile_binding data from API
func (r *AaagroupVpnsecureprivateaccessprofileBindingResource) readAaagroupVpnsecureprivateaccessprofileBindingFromApi(ctx context.Context, data *AaagroupVpnsecureprivateaccessprofileBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"groupname", "secureprivateaccessprofile"}, nil)
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
		ResourceType:             service.Aaagroup_vpnsecureprivateaccessprofile_binding.Type(),
		ResourceName:             groupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaagroup_vpnsecureprivateaccessprofile_binding, got error: %s", err))
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

		// Check secureprivateaccessprofile
		if idVal, ok := idMap["secureprivateaccessprofile"]; ok {
			if val, ok := v["secureprivateaccessprofile"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["secureprivateaccessprofile"].(string); ok {
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

	aaagroup_vpnsecureprivateaccessprofile_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
