package feoglobal_feopolicy_binding

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
var _ resource.Resource = &FeoglobalFeopolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*FeoglobalFeopolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*FeoglobalFeopolicyBindingResource)(nil)

func NewFeoglobalFeopolicyBindingResource() resource.Resource {
	return &FeoglobalFeopolicyBindingResource{}
}

// FeoglobalFeopolicyBindingResource defines the resource implementation.
type FeoglobalFeopolicyBindingResource struct {
	client *service.NitroClient
}

func (r *FeoglobalFeopolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *FeoglobalFeopolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_feoglobal_feopolicy_binding"
}

func (r *FeoglobalFeopolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *FeoglobalFeopolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data FeoglobalFeopolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating feoglobal_feopolicy_binding resource")
	feoglobal_feopolicy_binding := feoglobal_feopolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Feoglobal_feopolicy_binding.Type(), &feoglobal_feopolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create feoglobal_feopolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created feoglobal_feopolicy_binding resource")

	// Set ID for the resource before reading state.
	// Backward-compatible with SDK v2, which used the plain policyname as the ID
	// (d.SetId(policyname)); resource_id_mapping.json legacy order is "policyname".
	data.Id = types.StringValue(data.Policyname.ValueString())

	// Read the updated state back
	r.readFeoglobalFeopolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "feoglobal_feopolicy_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FeoglobalFeopolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data FeoglobalFeopolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading feoglobal_feopolicy_binding resource")

	r.readFeoglobalFeopolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *FeoglobalFeopolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state FeoglobalFeopolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating feoglobal_feopolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		feoglobal_feopolicy_binding := feoglobal_feopolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Feoglobal_feopolicy_binding.Type(), &feoglobal_feopolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update feoglobal_feopolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated feoglobal_feopolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for feoglobal_feopolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readFeoglobalFeopolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "feoglobal_feopolicy_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FeoglobalFeopolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FeoglobalFeopolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting feoglobal_feopolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name.
	// ID is the plain policyname (single legacy key); read the remaining identity
	// attributes from prior state. SDK v2 passed policyname, priority and type as
	// delete args. URL-encode values for slashy/special policyname expressions.
	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", url.QueryEscape(data.Policyname.ValueString())))
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		args = append(args, fmt.Sprintf("priority:%d", data.Priority.ValueInt64()))
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() && data.Type.ValueString() != "" {
		args = append(args, fmt.Sprintf("type:%s", url.QueryEscape(data.Type.ValueString())))
	} else {
		args = append(args, "type:REQ_DEFAULT")
	}

	err := r.client.DeleteResourceWithArgs(service.Feoglobal_feopolicy_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete feoglobal_feopolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted feoglobal_feopolicy_binding binding")
}

// Helper function to read feoglobal_feopolicy_binding data from API
func (r *FeoglobalFeopolicyBindingResource) readFeoglobalFeopolicyBindingFromApi(ctx context.Context, data *FeoglobalFeopolicyBindingResourceModel, diags *diag.Diagnostics) {

	// ID is the plain policyname (single legacy key). The type bindpoint is read
	// from state and used as the GET filter argument (SDK v2 defaulted to REQ_DEFAULT).
	policyname := data.Id.ValueString()
	typeVal := data.Type.ValueString()
	if typeVal == "" {
		typeVal = "REQ_DEFAULT"
	}

	var dataArr []map[string]interface{}
	var argsMap map[string]string = make(map[string]string)
	argsMap["type"] = typeVal

	findParams := service.FindParams{
		ResourceType:             service.Feoglobal_feopolicy_binding.Type(),
		ArgsMap:                  argsMap,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read feoglobal_feopolicy_binding, got error: %s", err))
		return
	}

	// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
	// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the matching policyname
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["policyname"].(string); ok && val == policyname {
			foundIndex = i
			break
		}
	}

	// Binding not present in the returned set: signal removal via a null Id (see above).
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	feoglobal_feopolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
