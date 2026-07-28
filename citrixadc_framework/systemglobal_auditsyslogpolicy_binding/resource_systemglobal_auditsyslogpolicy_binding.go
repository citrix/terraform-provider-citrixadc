package systemglobal_auditsyslogpolicy_binding

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
var _ resource.Resource = &SystemglobalAuditsyslogpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*SystemglobalAuditsyslogpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SystemglobalAuditsyslogpolicyBindingResource)(nil)

func NewSystemglobalAuditsyslogpolicyBindingResource() resource.Resource {
	return &SystemglobalAuditsyslogpolicyBindingResource{}
}

// SystemglobalAuditsyslogpolicyBindingResource defines the resource implementation.
type SystemglobalAuditsyslogpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *SystemglobalAuditsyslogpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemglobalAuditsyslogpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemglobal_auditsyslogpolicy_binding"
}

func (r *SystemglobalAuditsyslogpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemglobalAuditsyslogpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemglobal_auditsyslogpolicy_binding resource")
	systemglobal_auditsyslogpolicy_binding := systemglobal_auditsyslogpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Systemglobal_auditsyslogpolicy_binding.Type(), &systemglobal_auditsyslogpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemglobal_auditsyslogpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created systemglobal_auditsyslogpolicy_binding resource")

	// Set ID for the resource before reading state (Pattern 6: set ID once here).
	// Single unique attribute - ID is the plain policyname value.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	// Read the updated state back
	found := r.readSystemglobalAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.Diagnostics.AddError("Client Error", "systemglobal_auditsyslogpolicy_binding not found after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuditsyslogpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemglobal_auditsyslogpolicy_binding resource")

	found := r.readSystemglobalAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		// Binding removed out-of-band - remove from state to trigger recreation
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuditsyslogpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Pattern 5: NITRO exposes no update/set endpoint for this binding and every
	// schema attribute is RequiresReplace, so Update is a no-op (it is never
	// reached for an attribute change - Terraform recreates instead).
	tflog.Debug(ctx, "Update is a no-op for systemglobal_auditsyslogpolicy_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readSystemglobalAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuditsyslogpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemglobal_auditsyslogpolicy_binding resource")
	// Global binding - keyless URL, delete via DeleteResourceWithArgs with an
	// empty resource name and policyname passed as a UrlEncoded arg (not a path key).
	policyname_value := data.Id.ValueString()
	args := []string{
		"policyname:" + utils.UrlEncode(policyname_value),
	}

	err := r.client.DeleteResourceWithArgs(service.Systemglobal_auditsyslogpolicy_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete systemglobal_auditsyslogpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted systemglobal_auditsyslogpolicy_binding binding")
}

// Helper function to read systemglobal_auditsyslogpolicy_binding data from API.
// Returns true if the binding was found on the appliance, false otherwise (drift).
func (r *SystemglobalAuditsyslogpolicyBindingResource) readSystemglobalAuditsyslogpolicyBindingFromApi(ctx context.Context, data *SystemglobalAuditsyslogpolicyBindingResourceModel, diags *diag.Diagnostics) bool {

	// Keyless aggregate read: the systemglobal_auditsyslogpolicy_binding URL has
	// no key segment. Fetch the full array and filter client-side by policyname.
	// Pattern 10: the ID is a single plain value (policyname), so use it directly
	// as the filter value rather than ParseIdString.
	policynameFilter := data.Id.ValueString()

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Systemglobal_auditsyslogpolicy_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemglobal_auditsyslogpolicy_binding, got error: %s", err))
		return false
	}

	// Iterate through results to find the one with the right policyname
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["policyname"].(string); ok && val == policynameFilter {
			foundIndex = i
			break
		}
	}

	// Resource is missing on the appliance - signal drift to the caller
	if foundIndex == -1 {
		tflog.Debug(ctx, "systemglobal_auditsyslogpolicy_binding not found on appliance")
		return false
	}

	systemglobal_auditsyslogpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
	return true
}
