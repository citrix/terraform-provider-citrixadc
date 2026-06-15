package systemglobal_auditnslogpolicy_binding

import (
	"context"
	"fmt"
	"net/url"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystemglobalAuditnslogpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*SystemglobalAuditnslogpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SystemglobalAuditnslogpolicyBindingResource)(nil)

func NewSystemglobalAuditnslogpolicyBindingResource() resource.Resource {
	return &SystemglobalAuditnslogpolicyBindingResource{}
}

// SystemglobalAuditnslogpolicyBindingResource defines the resource implementation.
type SystemglobalAuditnslogpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *SystemglobalAuditnslogpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemglobalAuditnslogpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemglobal_auditnslogpolicy_binding"
}

func (r *SystemglobalAuditnslogpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemglobalAuditnslogpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemglobal_auditnslogpolicy_binding resource")
	systemglobal_auditnslogpolicy_binding := systemglobal_auditnslogpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Systemglobal_auditnslogpolicy_binding.Type(), &systemglobal_auditnslogpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemglobal_auditnslogpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created systemglobal_auditnslogpolicy_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	// Read the updated state back
	r.readSystemglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuditnslogpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemglobal_auditnslogpolicy_binding resource")

	r.readSystemglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuditnslogpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating systemglobal_auditnslogpolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		systemglobal_auditnslogpolicy_binding := systemglobal_auditnslogpolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Systemglobal_auditnslogpolicy_binding.Type(), &systemglobal_auditnslogpolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemglobal_auditnslogpolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated systemglobal_auditnslogpolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for systemglobal_auditnslogpolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSystemglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuditnslogpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemglobal_auditnslogpolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Single unique attribute - ID is the plain value
	policyname_value := data.Id.ValueString()
	args := []string{
		fmt.Sprintf("policyname:%s", url.QueryEscape(policyname_value)),
	}

	err := r.client.DeleteResourceWithArgs(service.Systemglobal_auditnslogpolicy_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete systemglobal_auditnslogpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted systemglobal_auditnslogpolicy_binding binding")
}

// Helper function to read systemglobal_auditnslogpolicy_binding data from API
func (r *SystemglobalAuditnslogpolicyBindingResource) readSystemglobalAuditnslogpolicyBindingFromApi(ctx context.Context, data *SystemglobalAuditnslogpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 3: Array filter without parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"policyname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Systemglobal_auditnslogpolicy_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemglobal_auditnslogpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "systemglobal_auditnslogpolicy_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check policyname
		if idVal, ok := idMap["policyname"]; ok {
			if val, ok := v["policyname"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["policyname"].(string); ok {
			match = false
			continue
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("systemglobal_auditnslogpolicy_binding not found with the provided ID attributes"))
		return
	}

	systemglobal_auditnslogpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
