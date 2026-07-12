package gslbservicegroup_lbmonitor_binding

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
var _ resource.Resource = &GslbservicegroupLbmonitorBindingResource{}
var _ resource.ResourceWithConfigure = (*GslbservicegroupLbmonitorBindingResource)(nil)
var _ resource.ResourceWithImportState = (*GslbservicegroupLbmonitorBindingResource)(nil)

func NewGslbservicegroupLbmonitorBindingResource() resource.Resource {
	return &GslbservicegroupLbmonitorBindingResource{}
}

// GslbservicegroupLbmonitorBindingResource defines the resource implementation.
type GslbservicegroupLbmonitorBindingResource struct {
	client *service.NitroClient
}

func (r *GslbservicegroupLbmonitorBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbservicegroupLbmonitorBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbservicegroup_lbmonitor_binding"
}

func (r *GslbservicegroupLbmonitorBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbservicegroupLbmonitorBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbservicegroupLbmonitorBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbservicegroup_lbmonitor_binding resource")
	gslbservicegroup_lbmonitor_binding := gslbservicegroup_lbmonitor_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Gslbservicegroup_lbmonitor_binding.Type(), &gslbservicegroup_lbmonitor_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create gslbservicegroup_lbmonitor_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created gslbservicegroup_lbmonitor_binding resource")

	// Set ID for the resource before reading state
	// Composite ID order matches the legacy SDK v2 comma ID ("servicegroupname,monitor_name")
	// so existing configurations remain importable.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("monitor_name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.MonitorName.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readGslbservicegroupLbmonitorBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "gslbservicegroup_lbmonitor_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbservicegroupLbmonitorBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbservicegroupLbmonitorBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading gslbservicegroup_lbmonitor_binding resource")

	r.readGslbservicegroupLbmonitorBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *GslbservicegroupLbmonitorBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state GslbservicegroupLbmonitorBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating gslbservicegroup_lbmonitor_binding resource")

	// All attributes force replacement, so there are no in-place updates.
	hasChange := false

	if hasChange {
		// Create API request body from the model
		gslbservicegroup_lbmonitor_binding := gslbservicegroup_lbmonitor_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Gslbservicegroup_lbmonitor_binding.Type(), &gslbservicegroup_lbmonitor_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update gslbservicegroup_lbmonitor_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated gslbservicegroup_lbmonitor_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for gslbservicegroup_lbmonitor_binding resource, skipping update")
	}

	// Read the updated state back
	r.readGslbservicegroupLbmonitorBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "gslbservicegroup_lbmonitor_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbservicegroupLbmonitorBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbservicegroupLbmonitorBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting gslbservicegroup_lbmonitor_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgsMap
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicegroupname", "monitor_name"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	servicegroupname_value, ok := idMap["servicegroupname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'servicegroupname' not found in ID")
		return
	}

	// Build delete args. NITRO expects the monitor name under "monitor_name" (SDK v2 parity).
	argsMap := make(map[string]string)
	if val, ok := idMap["monitor_name"]; ok && val != "" {
		argsMap["monitor_name"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Gslbservicegroup_lbmonitor_binding.Type(), servicegroupname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete gslbservicegroup_lbmonitor_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted gslbservicegroup_lbmonitor_binding binding")
}

// Helper function to read gslbservicegroup_lbmonitor_binding data from API
func (r *GslbservicegroupLbmonitorBindingResource) readGslbservicegroupLbmonitorBindingFromApi(ctx context.Context, data *GslbservicegroupLbmonitorBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicegroupname", "monitor_name"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	servicegroupname_Name, ok := idMap["servicegroupname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'servicegroupname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Gslbservicegroup_lbmonitor_binding.Type(),
		ResourceName:             servicegroupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbservicegroup_lbmonitor_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
		// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right monitor name.
	// NITRO returns the monitor name under the "monitor_name" key.
	monitorname_value, ok := idMap["monitor_name"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'monitor_name' not found in ID string")
		return
	}
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["monitor_name"].(string); ok && val == monitorname_value {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	gslbservicegroup_lbmonitor_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
