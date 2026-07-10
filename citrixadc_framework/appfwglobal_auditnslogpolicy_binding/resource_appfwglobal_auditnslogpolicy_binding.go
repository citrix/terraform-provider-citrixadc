package appfwglobal_auditnslogpolicy_binding

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
var _ resource.Resource = &AppfwglobalAuditnslogpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwglobalAuditnslogpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwglobalAuditnslogpolicyBindingResource)(nil)

func NewAppfwglobalAuditnslogpolicyBindingResource() resource.Resource {
	return &AppfwglobalAuditnslogpolicyBindingResource{}
}

// AppfwglobalAuditnslogpolicyBindingResource defines the resource implementation.
type AppfwglobalAuditnslogpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwglobalAuditnslogpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwglobalAuditnslogpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwglobal_auditnslogpolicy_binding"
}

func (r *AppfwglobalAuditnslogpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwglobalAuditnslogpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwglobal_auditnslogpolicy_binding resource")
	appfwglobal_auditnslogpolicy_binding := appfwglobal_auditnslogpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwglobal_auditnslogpolicy_binding.Type(), &appfwglobal_auditnslogpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwglobal_auditnslogpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwglobal_auditnslogpolicy_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Type.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAppfwglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwglobalAuditnslogpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwglobal_auditnslogpolicy_binding resource")

	r.readAppfwglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwglobalAuditnslogpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwglobal_auditnslogpolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		appfwglobal_auditnslogpolicy_binding := appfwglobal_auditnslogpolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Appfwglobal_auditnslogpolicy_binding.Type(), &appfwglobal_auditnslogpolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwglobal_auditnslogpolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appfwglobal_auditnslogpolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appfwglobal_auditnslogpolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAppfwglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwglobalAuditnslogpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwglobal_auditnslogpolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Multiple unique attributes - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"policyname"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["policyname"]; ok && val != "" {
		argsMap["policyname"] = val
	}
	// NOTE: 'type' is intentionally NOT passed as a delete arg. NITRO rejects any
	// explicit 'type' value for this binding with errorcode 1097.
	// Include priority from state for disambiguation, mirroring the SDK v2 delete.
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		argsMap["priority"] = fmt.Sprintf("%d", data.Priority.ValueInt64())
	}

	err = r.client.DeleteResourceWithArgsMap(service.Appfwglobal_auditnslogpolicy_binding.Type(), "", argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwglobal_auditnslogpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwglobal_auditnslogpolicy_binding binding")
}

// Helper function to read appfwglobal_auditnslogpolicy_binding data from API
func (r *AppfwglobalAuditnslogpolicyBindingResource) readAppfwglobalAuditnslogpolicyBindingFromApi(ctx context.Context, data *AppfwglobalAuditnslogpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 3: Array filter without parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"policyname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	var dataArr []map[string]interface{}
	// NOTE: 'type' is intentionally NOT used as a GET filter arg. NITRO rejects any
	// explicit 'type' value for this binding with errorcode 1097.
	var argsMap map[string]string = make(map[string]string)

	findParams := service.FindParams{
		ResourceType:             service.Appfwglobal_auditnslogpolicy_binding.Type(),
		ArgsMap:                  argsMap,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwglobal_auditnslogpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "appfwglobal_auditnslogpolicy_binding returned empty array")
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
		// NOTE: 'type' is not matched here — the GET response does not echo it for
		// this binding (it surfaces 'bindpolicytype' instead).

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("appfwglobal_auditnslogpolicy_binding not found with the provided ID attributes"))
		return
	}

	appfwglobal_auditnslogpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
