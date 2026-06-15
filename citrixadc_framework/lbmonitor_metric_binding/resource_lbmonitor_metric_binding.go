package lbmonitor_metric_binding

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
var _ resource.Resource = &LbmonitorMetricBindingResource{}
var _ resource.ResourceWithConfigure = (*LbmonitorMetricBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbmonitorMetricBindingResource)(nil)

func NewLbmonitorMetricBindingResource() resource.Resource {
	return &LbmonitorMetricBindingResource{}
}

// LbmonitorMetricBindingResource defines the resource implementation.
type LbmonitorMetricBindingResource struct {
	client *service.NitroClient
}

func (r *LbmonitorMetricBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbmonitorMetricBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbmonitor_metric_binding"
}

func (r *LbmonitorMetricBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbmonitorMetricBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbmonitorMetricBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbmonitor_metric_binding resource")
	lbmonitor_metric_binding := lbmonitor_metric_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Lbmonitor_metric_binding.Type(), &lbmonitor_metric_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbmonitor_metric_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lbmonitor_metric_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("metric:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Metric.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("monitorname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Monitorname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readLbmonitorMetricBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorMetricBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbmonitorMetricBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbmonitor_metric_binding resource")

	r.readLbmonitorMetricBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorMetricBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LbmonitorMetricBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lbmonitor_metric_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Metric.Equal(state.Metric) {
		tflog.Debug(ctx, fmt.Sprintf("metric has changed for lbmonitor_metric_binding"))
		hasChange = true
	}
	if !data.Metricthreshold.Equal(state.Metricthreshold) {
		tflog.Debug(ctx, fmt.Sprintf("metricthreshold has changed for lbmonitor_metric_binding"))
		hasChange = true
	}
	if !data.Metricweight.Equal(state.Metricweight) {
		tflog.Debug(ctx, fmt.Sprintf("metricweight has changed for lbmonitor_metric_binding"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		lbmonitor_metric_binding := lbmonitor_metric_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Lbmonitor_metric_binding.Type(), &lbmonitor_metric_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbmonitor_metric_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated lbmonitor_metric_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lbmonitor_metric_binding resource, skipping update")
	}

	// Read the updated state back
	r.readLbmonitorMetricBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorMetricBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbmonitorMetricBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbmonitor_metric_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"monitorname", "metric"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	monitorname_value, ok := idMap["monitorname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'monitorname' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["metric"]; ok && val != "" {
		argsMap["metric"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Lbmonitor_metric_binding.Type(), monitorname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lbmonitor_metric_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lbmonitor_metric_binding binding")
}

// Helper function to read lbmonitor_metric_binding data from API
func (r *LbmonitorMetricBindingResource) readLbmonitorMetricBindingFromApi(ctx context.Context, data *LbmonitorMetricBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"monitorname", "metric"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	monitorname_Name, ok := idMap["monitorname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'monitorname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Lbmonitor_metric_binding.Type(),
		ResourceName:             monitorname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbmonitor_metric_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "lbmonitor_metric_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check metric
		if idVal, ok := idMap["metric"]; ok {
			if val, ok := v["metric"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["metric"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("lbmonitor_metric_binding not found with the provided ID attributes"))
		return
	}

	lbmonitor_metric_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
