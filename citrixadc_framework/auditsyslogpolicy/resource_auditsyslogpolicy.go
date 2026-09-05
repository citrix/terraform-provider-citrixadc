package auditsyslogpolicy

import (
	"context"
	"fmt"
	"net/url"

	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AuditsyslogpolicyResource{}
var _ resource.ResourceWithConfigure = (*AuditsyslogpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AuditsyslogpolicyResource)(nil)

func NewAuditsyslogpolicyResource() resource.Resource {
	return &AuditsyslogpolicyResource{}
}

// AuditsyslogpolicyResource defines the resource implementation.
type AuditsyslogpolicyResource struct {
	client *service.NitroClient
}

func (r *AuditsyslogpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuditsyslogpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auditsyslogpolicy"
}

func (r *AuditsyslogpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuditsyslogpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuditsyslogpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating auditsyslogpolicy resource")

	// Get payload from plan
	auditsyslogpolicy := auditsyslogpolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Auditsyslogpolicy.Type(), name_value, &auditsyslogpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create auditsyslogpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created auditsyslogpolicy resource")

	// Handle inline globalbinding if configured
	if !data.Globalbinding.IsNull() {
		if err := r.syncGlobalbinding(ctx, name_value, nil, &data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create globalbinding for auditsyslogpolicy, got error: %s", err))
			return
		}
	}

	// Set ID for the resource before reading state
	data.Id = types.StringValue(name_value)

	// Read the updated state back
	if !r.readAuditsyslogpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "auditsyslogpolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditsyslogpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuditsyslogpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading auditsyslogpolicy resource")

	found := r.readAuditsyslogpolicyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditsyslogpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuditsyslogpolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating auditsyslogpolicy resource")

	name_value := data.Name.ValueString()

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for auditsyslogpolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for auditsyslogpolicy")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		auditsyslogpolicy := auditsyslogpolicyGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		_, err := r.client.UpdateResource(service.Auditsyslogpolicy.Type(), name_value, &auditsyslogpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update auditsyslogpolicy, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated auditsyslogpolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for auditsyslogpolicy resource, skipping update")
	}

	// Handle inline globalbinding changes
	if !data.Globalbinding.Equal(state.Globalbinding) {
		if err := r.syncGlobalbinding(ctx, name_value, &state, &data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update globalbinding for auditsyslogpolicy, got error: %s", err))
			return
		}
	}

	// Read the updated state back
	if !r.readAuditsyslogpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "auditsyslogpolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditsyslogpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuditsyslogpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting auditsyslogpolicy resource")
	name_value := data.Name.ValueString()

	// Unbind from systemglobal if bound, before deleting the policy.
	if !data.Globalbinding.IsNull() {
		if err := r.deleteAllGlobalbinding(name_value); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete globalbinding for auditsyslogpolicy, got error: %s", err))
			return
		}
	}

	// Named resource - delete using DeleteResource
	err := r.client.DeleteResource(service.Auditsyslogpolicy.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete auditsyslogpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted auditsyslogpolicy resource")
}

// Helper function to read auditsyslogpolicy data from API
func (r *AuditsyslogpolicyResource) readAuditsyslogpolicyFromApi(ctx context.Context, data *AuditsyslogpolicyResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Auditsyslogpolicy.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read auditsyslogpolicy, got error: %s", err))
		return false
	}

	auditsyslogpolicySetAttrFromGet(ctx, data, getResponseData)

	// Read globalbinding only if it was configured, to avoid a perpetual diff.
	if !data.Globalbinding.IsNull() {
		r.readGlobalbinding(ctx, data, diags)
	}

	return true
}

// auditsyslogpolicyGlobalbindingAttrTypes returns the attribute types for the globalbinding set.
func auditsyslogpolicyGlobalbindingAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"feature":                types.StringType,
		"globalbindtype":         types.StringType,
		"gotopriorityexpression": types.StringType,
		"nextfactor":             types.StringType,
		"priority":               types.Int64Type,
	}
}

// readGlobalbinding reads the current systemglobal binding from the API and sets it in state.
func (r *AuditsyslogpolicyResource) readGlobalbinding(ctx context.Context, data *AuditsyslogpolicyResourceModel, diags *diag.Diagnostics) {
	policyName := data.Name.ValueString()

	findParams := service.FindParams{
		ResourceType: "systemglobal_auditsyslogpolicy_binding",
		FilterMap:    map[string]string{"policyname": url.QueryEscape(policyName)},
	}

	bindings, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read globalbinding for auditsyslogpolicy, got error: %s", err))
		return
	}

	if len(bindings) == 0 {
		data.Globalbinding = types.SetNull(types.ObjectType{
			AttrTypes: auditsyslogpolicyGlobalbindingAttrTypes(),
		})
		return
	}

	b := bindings[0]
	model := AuditsyslogpolicyGlobalbindingModel{}
	if v, ok := b["feature"].(string); ok {
		model.Feature = types.StringValue(v)
	}
	if v, ok := b["globalbindtype"].(string); ok {
		model.Globalbindtype = types.StringValue(v)
	}
	if v, ok := b["gotopriorityexpression"].(string); ok {
		model.Gotopriorityexpression = types.StringValue(v)
	}
	if v, ok := b["nextfactor"].(string); ok {
		model.Nextfactor = types.StringValue(v)
	}
	if v, ok := b["priority"]; ok && v != nil {
		if intVal, err := utils.ConvertToInt64(v); err == nil {
			model.Priority = types.Int64Value(intVal)
		}
	}

	bindingModels := []AuditsyslogpolicyGlobalbindingModel{model}
	setValue, setDiags := types.SetValueFrom(ctx, types.ObjectType{
		AttrTypes: auditsyslogpolicyGlobalbindingAttrTypes(),
	}, bindingModels)
	diags.Append(setDiags...)
	data.Globalbinding = setValue
}

// syncGlobalbinding reconciles the systemglobal binding: it removes any prior binding
// and (re)adds the configured binding(s). Mirrors the SDK v2 delete-then-add sync.
func (r *AuditsyslogpolicyResource) syncGlobalbinding(ctx context.Context, policyName string, oldState *AuditsyslogpolicyResourceModel, newData *AuditsyslogpolicyResourceModel) error {
	// Delete any previously configured binding(s).
	if oldState != nil && !oldState.Globalbinding.IsNull() {
		if err := r.deleteAllGlobalbinding(policyName); err != nil {
			return err
		}
	}

	// Add the newly configured binding(s).
	if newData != nil && !newData.Globalbinding.IsNull() {
		var newBindings []AuditsyslogpolicyGlobalbindingModel
		newData.Globalbinding.ElementsAs(ctx, &newBindings, false)
		for _, b := range newBindings {
			if err := r.addSingleGlobalbinding(policyName, b); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *AuditsyslogpolicyResource) addSingleGlobalbinding(policyName string, binding AuditsyslogpolicyGlobalbindingModel) error {
	bindingStruct := system.Systemglobalauditsyslogpolicybinding{}
	bindingStruct.Policyname = policyName
	if !binding.Feature.IsNull() && !binding.Feature.IsUnknown() {
		bindingStruct.Feature = binding.Feature.ValueString()
	}
	if !binding.Globalbindtype.IsNull() && !binding.Globalbindtype.IsUnknown() {
		bindingStruct.Globalbindtype = binding.Globalbindtype.ValueString()
	}
	if !binding.Gotopriorityexpression.IsNull() && !binding.Gotopriorityexpression.IsUnknown() {
		bindingStruct.Gotopriorityexpression = binding.Gotopriorityexpression.ValueString()
	}
	if !binding.Nextfactor.IsNull() && !binding.Nextfactor.IsUnknown() {
		bindingStruct.Nextfactor = binding.Nextfactor.ValueString()
	}
	if !binding.Priority.IsNull() && !binding.Priority.IsUnknown() {
		bindingStruct.Priority = utils.IntPtr(int(binding.Priority.ValueInt64()))
	}

	return r.client.UpdateUnnamedResource("systemglobal_auditsyslogpolicy_binding", bindingStruct)
}

func (r *AuditsyslogpolicyResource) deleteAllGlobalbinding(policyName string) error {
	args := []string{fmt.Sprintf("policyname:%s", policyName)}
	return r.client.DeleteResourceWithArgs("systemglobal_auditsyslogpolicy_binding", "", args)
}
