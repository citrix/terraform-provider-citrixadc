package rnatglobal_auditsyslogpolicy_binding

import (
	"context"
	"fmt"
	"net/url"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RnatglobalAuditsyslogpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*RnatglobalAuditsyslogpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*RnatglobalAuditsyslogpolicyBindingResource)(nil)

func NewRnatglobalAuditsyslogpolicyBindingResource() resource.Resource {
	return &RnatglobalAuditsyslogpolicyBindingResource{}
}

// RnatglobalAuditsyslogpolicyBindingResource defines the resource implementation.
type RnatglobalAuditsyslogpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *RnatglobalAuditsyslogpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RnatglobalAuditsyslogpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rnatglobal_auditsyslogpolicy_binding"
}

func (r *RnatglobalAuditsyslogpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RnatglobalAuditsyslogpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RnatglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rnatglobal_auditsyslogpolicy_binding resource")
	rnatglobal_auditsyslogpolicy_binding := rnatglobal_auditsyslogpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Rnatglobal_auditsyslogpolicy_binding.Type(), &rnatglobal_auditsyslogpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create rnatglobal_auditsyslogpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created rnatglobal_auditsyslogpolicy_binding resource")

	// Set ID for the resource before reading state
	// Single unique key (policy) - plain value ID (matches SDK v2 d.SetId(policy))
	data.Id = types.StringValue(data.Policy.ValueString())

	// Read the updated state back
	r.readRnatglobalAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "rnatglobal_auditsyslogpolicy_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatglobalAuditsyslogpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RnatglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading rnatglobal_auditsyslogpolicy_binding resource")

	r.readRnatglobalAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *RnatglobalAuditsyslogpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state RnatglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating rnatglobal_auditsyslogpolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		rnatglobal_auditsyslogpolicy_binding := rnatglobal_auditsyslogpolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Rnatglobal_auditsyslogpolicy_binding.Type(), &rnatglobal_auditsyslogpolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update rnatglobal_auditsyslogpolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated rnatglobal_auditsyslogpolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for rnatglobal_auditsyslogpolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readRnatglobalAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "rnatglobal_auditsyslogpolicy_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatglobalAuditsyslogpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RnatglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting rnatglobal_auditsyslogpolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name.
	// Single unique key (policy) is the plain-value ID; "all" is an optional delete arg
	// (matches SDK v2 deleteRnatglobal_auditsyslogpolicy_bindingFunc).
	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policy:%s", url.QueryEscape(data.Policy.ValueString())))
	if !data.All.IsNull() && !data.All.IsUnknown() && data.All.ValueBool() {
		args = append(args, fmt.Sprintf("all:%v", data.All.ValueBool()))
	}

	err := r.client.DeleteResourceWithArgs(service.Rnatglobal_auditsyslogpolicy_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete rnatglobal_auditsyslogpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted rnatglobal_auditsyslogpolicy_binding binding")
}

// Helper function to read rnatglobal_auditsyslogpolicy_binding data from API
func (r *RnatglobalAuditsyslogpolicyBindingResource) readRnatglobalAuditsyslogpolicyBindingFromApi(ctx context.Context, data *RnatglobalAuditsyslogpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Single unique key (policy): the ID is the plain policy value (matches SDK v2).
	policy := data.Id.ValueString()

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Rnatglobal_auditsyslogpolicy_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read rnatglobal_auditsyslogpolicy_binding, got error: %s", err))
		return
	}

	// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
	// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the matching policy
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["policy"].(string); ok && val == policy {
			foundIndex = i
			break
		}
	}

	// Binding not present in the returned set: signal removal via a null Id (see above).
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	rnatglobal_auditsyslogpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
