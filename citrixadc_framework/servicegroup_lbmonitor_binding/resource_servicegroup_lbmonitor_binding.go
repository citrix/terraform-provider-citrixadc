package servicegroup_lbmonitor_binding

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
var _ resource.Resource = &ServicegroupLbmonitorBindingResource{}
var _ resource.ResourceWithConfigure = (*ServicegroupLbmonitorBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ServicegroupLbmonitorBindingResource)(nil)

func NewServicegroupLbmonitorBindingResource() resource.Resource {
	return &ServicegroupLbmonitorBindingResource{}
}

// ServicegroupLbmonitorBindingResource defines the resource implementation.
type ServicegroupLbmonitorBindingResource struct {
	client *service.NitroClient
}

func (r *ServicegroupLbmonitorBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ServicegroupLbmonitorBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_servicegroup_lbmonitor_binding"
}

func (r *ServicegroupLbmonitorBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ServicegroupLbmonitorBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ServicegroupLbmonitorBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating servicegroup_lbmonitor_binding resource")
	servicegroup_lbmonitor_binding := servicegroup_lbmonitor_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Servicegroup_lbmonitor_binding.Type(), &servicegroup_lbmonitor_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create servicegroup_lbmonitor_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created servicegroup_lbmonitor_binding resource")

	// Set ID for the resource before reading state
	// Composite ID order matches resource_id_mapping.json ("servicegroupname,monitorname")
	// so legacy SDK v2 comma IDs (servicegroupname,monitorname) remain importable.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("monitorname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.MonitorName.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readServicegroupLbmonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServicegroupLbmonitorBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ServicegroupLbmonitorBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading servicegroup_lbmonitor_binding resource")

	r.readServicegroupLbmonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServicegroupLbmonitorBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ServicegroupLbmonitorBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating servicegroup_lbmonitor_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		servicegroup_lbmonitor_binding := servicegroup_lbmonitor_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Servicegroup_lbmonitor_binding.Type(), &servicegroup_lbmonitor_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update servicegroup_lbmonitor_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated servicegroup_lbmonitor_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for servicegroup_lbmonitor_binding resource, skipping update")
	}

	// Read the updated state back
	r.readServicegroupLbmonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServicegroupLbmonitorBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ServicegroupLbmonitorBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting servicegroup_lbmonitor_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicegroupname", "monitorname"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	servicegroupname_value, ok := idMap["servicegroupname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'servicegroupname' not found in ID")
		return
	}

	// Build delete args. NITRO expects the monitor name under "monitor_name".
	// URL-encode values to handle slashy/special characters (matches SDK v2 args).
	args := make([]string, 0)
	if val, ok := idMap["monitorname"]; ok && val != "" {
		args = append(args, fmt.Sprintf("monitor_name:%s", url.QueryEscape(val)))
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		args = append(args, fmt.Sprintf("port:%v", data.Port.ValueInt64()))
	}

	err = r.client.DeleteResourceWithArgs(service.Servicegroup_lbmonitor_binding.Type(), servicegroupname_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete servicegroup_lbmonitor_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted servicegroup_lbmonitor_binding binding")
}

// Helper function to read servicegroup_lbmonitor_binding data from API
func (r *ServicegroupLbmonitorBindingResource) readServicegroupLbmonitorBindingFromApi(ctx context.Context, data *ServicegroupLbmonitorBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicegroupname", "monitorname"}, nil)
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
		ResourceType:             service.Servicegroup_lbmonitor_binding.Type(),
		ResourceName:             servicegroupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read servicegroup_lbmonitor_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "servicegroup_lbmonitor_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right monitor name.
	// NITRO returns the monitor name under the "monitor_name" key; the composite ID
	// carries it under the legacy "monitorname" key.
	monitorname_value, ok := idMap["monitorname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'monitorname' not found in ID string")
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
		diags.AddError("Client Error", fmt.Sprintf("servicegroup_lbmonitor_binding not found with the provided ID attributes"))
		return
	}

	servicegroup_lbmonitor_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
