package analyticsglobal_analyticsprofile_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AnalyticsglobalAnalyticsprofileBindingResource{}
var _ resource.ResourceWithConfigure = (*AnalyticsglobalAnalyticsprofileBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AnalyticsglobalAnalyticsprofileBindingResource)(nil)

func NewAnalyticsglobalAnalyticsprofileBindingResource() resource.Resource {
	return &AnalyticsglobalAnalyticsprofileBindingResource{}
}

// AnalyticsglobalAnalyticsprofileBindingResource defines the resource implementation.
type AnalyticsglobalAnalyticsprofileBindingResource struct {
	client *service.NitroClient
}

func (r *AnalyticsglobalAnalyticsprofileBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AnalyticsglobalAnalyticsprofileBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_analyticsglobal_analyticsprofile_binding"
}

func (r *AnalyticsglobalAnalyticsprofileBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AnalyticsglobalAnalyticsprofileBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AnalyticsglobalAnalyticsprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating analyticsglobal_analyticsprofile_binding resource")
	analyticsglobal_analyticsprofile_binding := analyticsglobal_analyticsprofile_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Analyticsglobal_analyticsprofile_binding.Type(), &analyticsglobal_analyticsprofile_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create analyticsglobal_analyticsprofile_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created analyticsglobal_analyticsprofile_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Analyticsprofile.ValueString()))

	// Read the updated state back
	r.readAnalyticsglobalAnalyticsprofileBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "analyticsglobal_analyticsprofile_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AnalyticsglobalAnalyticsprofileBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AnalyticsglobalAnalyticsprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading analyticsglobal_analyticsprofile_binding resource")

	r.readAnalyticsglobalAnalyticsprofileBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AnalyticsglobalAnalyticsprofileBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AnalyticsglobalAnalyticsprofileBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating analyticsglobal_analyticsprofile_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		analyticsglobal_analyticsprofile_binding := analyticsglobal_analyticsprofile_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Analyticsglobal_analyticsprofile_binding.Type(), &analyticsglobal_analyticsprofile_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update analyticsglobal_analyticsprofile_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated analyticsglobal_analyticsprofile_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for analyticsglobal_analyticsprofile_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAnalyticsglobalAnalyticsprofileBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "analyticsglobal_analyticsprofile_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AnalyticsglobalAnalyticsprofileBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AnalyticsglobalAnalyticsprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting analyticsglobal_analyticsprofile_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Single unique attribute - ID is the plain value
	analyticsprofile_value := data.Id.ValueString()
	args := []string{
		fmt.Sprintf("analyticsprofile:%s", analyticsprofile_value),
	}

	err := r.client.DeleteResourceWithArgs(service.Analyticsglobal_analyticsprofile_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete analyticsglobal_analyticsprofile_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted analyticsglobal_analyticsprofile_binding binding")
}

// Helper function to read analyticsglobal_analyticsprofile_binding data from API
func (r *AnalyticsglobalAnalyticsprofileBindingResource) readAnalyticsglobalAnalyticsprofileBindingFromApi(ctx context.Context, data *AnalyticsglobalAnalyticsprofileBindingResourceModel, diags *diag.Diagnostics) {

	// Case 3: Array filter without parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"analyticsprofile"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             "analyticsglobal",
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read analyticsglobal_analyticsprofile_binding, got error: %s", err))
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

		// Check analyticsprofile
		if idVal, ok := idMap["analyticsprofile"]; ok {
			if val, ok := v["analyticsprofile"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["analyticsprofile"].(string); ok {
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

	analyticsglobal_analyticsprofile_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
