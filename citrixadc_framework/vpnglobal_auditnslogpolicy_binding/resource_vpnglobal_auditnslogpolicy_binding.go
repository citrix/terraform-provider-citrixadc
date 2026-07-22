package vpnglobal_auditnslogpolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VpnglobalAuditnslogpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalAuditnslogpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalAuditnslogpolicyBindingResource)(nil)

func NewVpnglobalAuditnslogpolicyBindingResource() resource.Resource {
	return &VpnglobalAuditnslogpolicyBindingResource{}
}

// VpnglobalAuditnslogpolicyBindingResource defines the resource implementation.
type VpnglobalAuditnslogpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalAuditnslogpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalAuditnslogpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_auditnslogpolicy_binding"
}

func (r *VpnglobalAuditnslogpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalAuditnslogpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_auditnslogpolicy_binding resource")
	vpnglobal_auditnslogpolicy_binding := vpnglobal_auditnslogpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vpnglobal_auditnslogpolicy_binding.Type(), &vpnglobal_auditnslogpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_auditnslogpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnglobal_auditnslogpolicy_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	// Read the updated state back
	r.readVpnglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuditnslogpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_auditnslogpolicy_binding resource")

	r.readVpnglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// If the object was deleted out-of-band, remove it from state so a subsequent apply re-creates it
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuditnslogpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// NITRO exposes no update endpoint for vpnglobal_auditnslogpolicy_binding
	// (only add/delete/get/count); every schema attribute is RequiresReplace,
	// so Terraform never reaches Update with a changed value. This is a no-op.
	tflog.Debug(ctx, "Update is a no-op for vpnglobal_auditnslogpolicy_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readVpnglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuditnslogpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_auditnslogpolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Single unique attribute - ID is the plain value
	policyname_value := data.Id.ValueString()
	// NITRO delete args = policyname,secondary,groupextraction (per NITRO doc)
	args := []string{
		fmt.Sprintf("policyname:%s", policyname_value),
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		args = append(args, fmt.Sprintf("secondary:%t", data.Secondary.ValueBool()))
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		args = append(args, fmt.Sprintf("groupextraction:%t", data.Groupextraction.ValueBool()))
	}

	err := r.client.DeleteResourceWithArgs(service.Vpnglobal_auditnslogpolicy_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnglobal_auditnslogpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnglobal_auditnslogpolicy_binding binding")
}

// Helper function to read vpnglobal_auditnslogpolicy_binding data from API
func (r *VpnglobalAuditnslogpolicyBindingResource) readVpnglobalAuditnslogpolicyBindingFromApi(ctx context.Context, data *VpnglobalAuditnslogpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Single unique attribute - ID is the plain policyname value (Pattern 10:
	// do not use ParseIdString on a plain-value ID, it returns an empty map).
	policyname_value := data.Id.ValueString()

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Vpnglobal_auditnslogpolicy_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_auditnslogpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing (deleted out-of-band) - signal removal by nulling the Id
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right policyname
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["policyname"].(string); ok && val == policyname_value {
			foundIndex = i
			break
		}
	}

	// Resource is missing (deleted out-of-band) - signal removal by nulling the Id
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	vpnglobal_auditnslogpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
