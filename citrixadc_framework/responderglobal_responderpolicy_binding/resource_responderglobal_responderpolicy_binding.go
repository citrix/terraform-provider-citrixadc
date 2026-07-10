package responderglobal_responderpolicy_binding

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
var _ resource.Resource = &ResponderglobalResponderpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*ResponderglobalResponderpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ResponderglobalResponderpolicyBindingResource)(nil)

func NewResponderglobalResponderpolicyBindingResource() resource.Resource {
	return &ResponderglobalResponderpolicyBindingResource{}
}

// ResponderglobalResponderpolicyBindingResource defines the resource implementation.
type ResponderglobalResponderpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *ResponderglobalResponderpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ResponderglobalResponderpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_responderglobal_responderpolicy_binding"
}

func (r *ResponderglobalResponderpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ResponderglobalResponderpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ResponderglobalResponderpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating responderglobal_responderpolicy_binding resource")
	responderglobal_responderpolicy_binding := responderglobal_responderpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Responderglobal_responderpolicy_binding.Type(), &responderglobal_responderpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create responderglobal_responderpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created responderglobal_responderpolicy_binding resource")

	// Set ID for the resource before reading state.
	// Backward-compatible with SDK v2 (d.SetId(policyname)) and resource_id_mapping.json
	// which maps this binding to a single "policyname" key.
	data.Id = types.StringValue(data.Policyname.ValueString())

	// Read the updated state back
	r.readResponderglobalResponderpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderglobalResponderpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ResponderglobalResponderpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading responderglobal_responderpolicy_binding resource")

	r.readResponderglobalResponderpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderglobalResponderpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ResponderglobalResponderpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating responderglobal_responderpolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		responderglobal_responderpolicy_binding := responderglobal_responderpolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Responderglobal_responderpolicy_binding.Type(), &responderglobal_responderpolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update responderglobal_responderpolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated responderglobal_responderpolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for responderglobal_responderpolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readResponderglobalResponderpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderglobalResponderpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ResponderglobalResponderpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting responderglobal_responderpolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name.
	// The ID is the plain policyname (SDK v2 contract). priority and type are read from
	// state (they are part of the binding's unique identity). Values are URL-encoded so
	// slashy/special characters survive (matches SDK v2 deleteResponderglobal_..._bindingFunc).
	policyname := data.Id.ValueString()
	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", url.QueryEscape(policyname)))
	if !data.Type.IsNull() && data.Type.ValueString() != "" {
		args = append(args, fmt.Sprintf("type:%s", url.QueryEscape(data.Type.ValueString())))
	}
	if !data.Priority.IsNull() {
		args = append(args, fmt.Sprintf("priority:%v", data.Priority.ValueInt64()))
	}

	err := r.client.DeleteResourceWithArgs(service.Responderglobal_responderpolicy_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete responderglobal_responderpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted responderglobal_responderpolicy_binding binding")
}

// Helper function to read responderglobal_responderpolicy_binding data from API
func (r *ResponderglobalResponderpolicyBindingResource) readResponderglobalResponderpolicyBindingFromApi(ctx context.Context, data *ResponderglobalResponderpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// ID is the plain policyname (SDK v2 contract). priority and type are read from
	// state since they form the binding's unique identity (Pattern 10 - no ParseIdString).
	policyname := data.Id.ValueString()

	var dataArr []map[string]interface{}
	var argsMap map[string]string = make(map[string]string)
	// Match SDK v2 read: the GET requires a bind-point type filter; default to
	// REQ_DEFAULT when the user did not set type explicitly.
	if !data.Type.IsNull() && !data.Type.IsUnknown() && data.Type.ValueString() != "" {
		argsMap["type"] = data.Type.ValueString()
	} else {
		argsMap["type"] = "REQ_DEFAULT"
	}

	findParams := service.FindParams{
		ResourceType:             service.Responderglobal_responderpolicy_binding.Type(),
		ArgsMap:                  argsMap,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read responderglobal_responderpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "responderglobal_responderpolicy_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the matching policyname (and priority
	// if it is known in state).
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check policyname
		if val, ok := v["policyname"].(string); ok {
			if val != policyname {
				match = false
				continue
			}
		} else {
			match = false
			continue
		}

		// Check priority when known in state
		if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
			if val, ok := v["priority"]; ok {
				vi, _ := utils.ConvertToInt64(val)
				if vi != data.Priority.ValueInt64() {
					match = false
					continue
				}
			}
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("responderglobal_responderpolicy_binding not found with the provided ID attributes"))
		return
	}

	responderglobal_responderpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
