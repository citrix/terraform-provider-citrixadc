package lbmetrictable_metric_binding

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
var _ resource.Resource = &LbmetrictableMetricBindingResource{}
var _ resource.ResourceWithConfigure = (*LbmetrictableMetricBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbmetrictableMetricBindingResource)(nil)

func NewLbmetrictableMetricBindingResource() resource.Resource {
	return &LbmetrictableMetricBindingResource{}
}

// LbmetrictableMetricBindingResource defines the resource implementation.
type LbmetrictableMetricBindingResource struct {
	client *service.NitroClient
}

func (r *LbmetrictableMetricBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbmetrictableMetricBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbmetrictable_metric_binding"
}

func (r *LbmetrictableMetricBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbmetrictableMetricBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbmetrictableMetricBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbmetrictable_metric_binding resource")
	lbmetrictable_metric_binding := lbmetrictable_metric_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Lbmetrictable_metric_binding.Type(), &lbmetrictable_metric_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbmetrictable_metric_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lbmetrictable_metric_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("metric:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Metric.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("metrictable:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Metrictable.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readLbmetrictableMetricBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "lbmetrictable_metric_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmetrictableMetricBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbmetrictableMetricBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbmetrictable_metric_binding resource")

	r.readLbmetrictableMetricBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LbmetrictableMetricBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LbmetrictableMetricBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lbmetrictable_metric_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Metric.Equal(state.Metric) {
		tflog.Debug(ctx, fmt.Sprintf("metric has changed for lbmetrictable_metric_binding"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		lbmetrictable_metric_binding := lbmetrictable_metric_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Lbmetrictable_metric_binding.Type(), &lbmetrictable_metric_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbmetrictable_metric_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated lbmetrictable_metric_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lbmetrictable_metric_binding resource, skipping update")
	}

	// Read the updated state back
	r.readLbmetrictableMetricBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "lbmetrictable_metric_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmetrictableMetricBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbmetrictableMetricBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbmetrictable_metric_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"metrictable", "metric"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	metrictable_value, ok := idMap["metrictable"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'metrictable' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["metric"]; ok && val != "" {
		argsMap["metric"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Lbmetrictable_metric_binding.Type(), metrictable_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lbmetrictable_metric_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lbmetrictable_metric_binding binding")
}

// Helper function to read lbmetrictable_metric_binding data from API
func (r *LbmetrictableMetricBindingResource) readLbmetrictableMetricBindingFromApi(ctx context.Context, data *LbmetrictableMetricBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"metrictable", "metric"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	metrictable_Name, ok := idMap["metrictable"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'metrictable' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Lbmetrictable_metric_binding.Type(),
		ResourceName:             metrictable_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbmetrictable_metric_binding, got error: %s", err))
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
		data.Id = types.StringNull()
		return
	}

	lbmetrictable_metric_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
