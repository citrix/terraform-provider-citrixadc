package nd6ravariables_onlinkipv6prefix_binding

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
var _ resource.Resource = &Nd6ravariablesOnlinkipv6prefixBindingResource{}
var _ resource.ResourceWithConfigure = (*Nd6ravariablesOnlinkipv6prefixBindingResource)(nil)
var _ resource.ResourceWithImportState = (*Nd6ravariablesOnlinkipv6prefixBindingResource)(nil)

func NewNd6ravariablesOnlinkipv6prefixBindingResource() resource.Resource {
	return &Nd6ravariablesOnlinkipv6prefixBindingResource{}
}

// Nd6ravariablesOnlinkipv6prefixBindingResource defines the resource implementation.
type Nd6ravariablesOnlinkipv6prefixBindingResource struct {
	client *service.NitroClient
}

func (r *Nd6ravariablesOnlinkipv6prefixBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Nd6ravariablesOnlinkipv6prefixBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nd6ravariables_onlinkipv6prefix_binding"
}

func (r *Nd6ravariablesOnlinkipv6prefixBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Nd6ravariablesOnlinkipv6prefixBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Nd6ravariablesOnlinkipv6prefixBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nd6ravariables_onlinkipv6prefix_binding resource")
	nd6ravariables_onlinkipv6prefix_binding := nd6ravariables_onlinkipv6prefix_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Nd6ravariables_onlinkipv6prefix_binding.Type(), &nd6ravariables_onlinkipv6prefix_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nd6ravariables_onlinkipv6prefix_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nd6ravariables_onlinkipv6prefix_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ipv6prefix:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ipv6prefix.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vlan:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vlan.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readNd6ravariablesOnlinkipv6prefixBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "nd6ravariables_onlinkipv6prefix_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nd6ravariablesOnlinkipv6prefixBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Nd6ravariablesOnlinkipv6prefixBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nd6ravariables_onlinkipv6prefix_binding resource")

	r.readNd6ravariablesOnlinkipv6prefixBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *Nd6ravariablesOnlinkipv6prefixBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state Nd6ravariablesOnlinkipv6prefixBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nd6ravariables_onlinkipv6prefix_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		nd6ravariables_onlinkipv6prefix_binding := nd6ravariables_onlinkipv6prefix_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Nd6ravariables_onlinkipv6prefix_binding.Type(), &nd6ravariables_onlinkipv6prefix_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nd6ravariables_onlinkipv6prefix_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nd6ravariables_onlinkipv6prefix_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nd6ravariables_onlinkipv6prefix_binding resource, skipping update")
	}

	// Read the updated state back
	r.readNd6ravariablesOnlinkipv6prefixBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "nd6ravariables_onlinkipv6prefix_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nd6ravariablesOnlinkipv6prefixBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Nd6ravariablesOnlinkipv6prefixBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nd6ravariables_onlinkipv6prefix_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"vlan", "ipv6prefix"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	vlan_value, ok := idMap["vlan"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'vlan' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["ipv6prefix"]; ok && val != "" {
		// ipv6prefix contains ':' and '/' (e.g. "2003::/64"); the NITRO client joins
		// arg values into the DELETE URL raw, so escape the value here (matches the
		// SDK v2 resource which used url.PathEscape on the ipv6prefix delete arg).
		argsMap["ipv6prefix"] = url.PathEscape(val)
	}

	err = r.client.DeleteResourceWithArgsMap(service.Nd6ravariables_onlinkipv6prefix_binding.Type(), vlan_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nd6ravariables_onlinkipv6prefix_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nd6ravariables_onlinkipv6prefix_binding binding")
}

// Helper function to read nd6ravariables_onlinkipv6prefix_binding data from API
func (r *Nd6ravariablesOnlinkipv6prefixBindingResource) readNd6ravariablesOnlinkipv6prefixBindingFromApi(ctx context.Context, data *Nd6ravariablesOnlinkipv6prefixBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"vlan", "ipv6prefix"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	vlan_Name, ok := idMap["vlan"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'vlan' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Nd6ravariables_onlinkipv6prefix_binding.Type(),
		ResourceName:             vlan_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nd6ravariables_onlinkipv6prefix_binding, got error: %s", err))
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

		// Check ipv6prefix
		if idVal, ok := idMap["ipv6prefix"]; ok {
			if val, ok := v["ipv6prefix"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["ipv6prefix"].(string); ok {
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

	nd6ravariables_onlinkipv6prefix_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
