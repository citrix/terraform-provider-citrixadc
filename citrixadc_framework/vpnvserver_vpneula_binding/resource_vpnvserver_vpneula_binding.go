package vpnvserver_vpneula_binding

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
var _ resource.Resource = &VpnvserverVpneulaBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverVpneulaBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverVpneulaBindingResource)(nil)

func NewVpnvserverVpneulaBindingResource() resource.Resource {
	return &VpnvserverVpneulaBindingResource{}
}

// VpnvserverVpneulaBindingResource defines the resource implementation.
type VpnvserverVpneulaBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverVpneulaBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverVpneulaBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_vpneula_binding"
}

func (r *VpnvserverVpneulaBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverVpneulaBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverVpneulaBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_vpneula_binding resource")
	vpnvserver_vpneula_binding := vpnvserver_vpneula_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vpnvserver_vpneula_binding.Type(), &vpnvserver_vpneula_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_vpneula_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnvserver_vpneula_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("eula:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Eula.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readVpnvserverVpneulaBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "vpnvserver_vpneula_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverVpneulaBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverVpneulaBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_vpneula_binding resource")

	r.readVpnvserverVpneulaBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnvserverVpneulaBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnvserverVpneulaBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnvserver_vpneula_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		vpnvserver_vpneula_binding := vpnvserver_vpneula_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Vpnvserver_vpneula_binding.Type(), &vpnvserver_vpneula_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_vpneula_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vpnvserver_vpneula_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnvserver_vpneula_binding resource, skipping update")
	}

	// Read the updated state back
	r.readVpnvserverVpneulaBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "vpnvserver_vpneula_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverVpneulaBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverVpneulaBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_vpneula_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "eula"}, nil)
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
	if val, ok := idMap["eula"]; ok && val != "" {
		// URL-encode the delete arg value to safely handle slashy/special characters
		argsMap["eula"] = url.QueryEscape(val)
	}

	err = r.client.DeleteResourceWithArgsMap(service.Vpnvserver_vpneula_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnvserver_vpneula_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnvserver_vpneula_binding binding")
}

// Helper function to read vpnvserver_vpneula_binding data from API
func (r *VpnvserverVpneulaBindingResource) readVpnvserverVpneulaBindingFromApi(ctx context.Context, data *VpnvserverVpneulaBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "eula"}, nil)
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
		ResourceType:             service.Vpnvserver_vpneula_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_vpneula_binding, got error: %s", err))
		return
	}

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

		// Check eula
		if idVal, ok := idMap["eula"]; ok {
			if val, ok := v["eula"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["eula"].(string); ok {
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

	vpnvserver_vpneula_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
