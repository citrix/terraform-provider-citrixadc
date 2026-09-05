package nstrafficdomain_vlan_binding

import (
	"context"
	"fmt"
	"strconv"
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
var _ resource.Resource = &NstrafficdomainVlanBindingResource{}
var _ resource.ResourceWithConfigure = (*NstrafficdomainVlanBindingResource)(nil)
var _ resource.ResourceWithImportState = (*NstrafficdomainVlanBindingResource)(nil)

func NewNstrafficdomainVlanBindingResource() resource.Resource {
	return &NstrafficdomainVlanBindingResource{}
}

// NstrafficdomainVlanBindingResource defines the resource implementation.
type NstrafficdomainVlanBindingResource struct {
	client *service.NitroClient
}

func (r *NstrafficdomainVlanBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NstrafficdomainVlanBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nstrafficdomain_vlan_binding"
}

func (r *NstrafficdomainVlanBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NstrafficdomainVlanBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NstrafficdomainVlanBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nstrafficdomain_vlan_binding resource")
	nstrafficdomain_vlan_binding := nstrafficdomain_vlan_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Nstrafficdomain_vlan_binding.Type(), &nstrafficdomain_vlan_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nstrafficdomain_vlan_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nstrafficdomain_vlan_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("td:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Td.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("vlan:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vlan.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readNstrafficdomainVlanBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "nstrafficdomain_vlan_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstrafficdomainVlanBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NstrafficdomainVlanBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nstrafficdomain_vlan_binding resource")

	r.readNstrafficdomainVlanBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NstrafficdomainVlanBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NstrafficdomainVlanBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nstrafficdomain_vlan_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		nstrafficdomain_vlan_binding := nstrafficdomain_vlan_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Nstrafficdomain_vlan_binding.Type(), &nstrafficdomain_vlan_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nstrafficdomain_vlan_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nstrafficdomain_vlan_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nstrafficdomain_vlan_binding resource, skipping update")
	}

	// Read the updated state back
	r.readNstrafficdomainVlanBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "nstrafficdomain_vlan_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstrafficdomainVlanBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NstrafficdomainVlanBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nstrafficdomain_vlan_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"td", "vlan"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	td_value, ok := idMap["td"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'td' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["vlan"]; ok && val != "" {
		argsMap["vlan"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Nstrafficdomain_vlan_binding.Type(), td_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nstrafficdomain_vlan_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nstrafficdomain_vlan_binding binding")
}

// Helper function to read nstrafficdomain_vlan_binding data from API
func (r *NstrafficdomainVlanBindingResource) readNstrafficdomainVlanBindingFromApi(ctx context.Context, data *NstrafficdomainVlanBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"td", "vlan"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	td_Name, ok := idMap["td"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'td' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Nstrafficdomain_vlan_binding.Type(),
		ResourceName:             td_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nstrafficdomain_vlan_binding, got error: %s", err))
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

		// Check vlan
		if idVal, ok := idMap["vlan"]; ok {
			if val, ok := v["vlan"]; ok {
				val, _ = utils.ConvertToInt64(val)
				idValInt64, _ := strconv.ParseInt(idVal, 10, 64)
				if val != idValInt64 {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["vlan"]; ok {
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

	nstrafficdomain_vlan_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
