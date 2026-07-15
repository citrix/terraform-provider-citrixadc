package videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding

import (
	"context"
	"fmt"
	"strconv"
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
var _ resource.Resource = &VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResource)(nil)

func NewVideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResource() resource.Resource {
	return &VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResource{}
}

// VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResource defines the resource implementation.
type VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding"
}

func (r *VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding resource")
	videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding := videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding.Type(), &videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("labelname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Labelname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("priority:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Priority.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readVideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding resource")

	r.readVideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Object was deleted out-of-band; remove it from state so a subsequent apply re-creates it.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for this binding: NITRO exposes no update endpoint and all
	// attributes use RequiresReplace, so Terraform never reaches Update for a real change (Pattern 5).
	tflog.Debug(ctx, "Update is a no-op for videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readVideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	labelname_value, ok := idMap["labelname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'labelname' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["policyname"]; ok && val != "" {
		argsMap["policyname"] = val
	}
	if val, ok := idMap["priority"]; ok && val != "" {
		argsMap["priority"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding.Type(), labelname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding binding")
}

// Helper function to read videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding data from API
func (r *VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResource) readVideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingFromApi(ctx context.Context, data *VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	labelname_Name, ok := idMap["labelname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'labelname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding.Type(),
		ResourceName:             labelname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing (deleted out-of-band); signal removal via null Id.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
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

		// Check priority
		if idVal, ok := idMap["priority"]; ok {
			if val, ok := v["priority"]; ok {
				val, _ = utils.ConvertToInt64(val)
				idValInt64, _ := strconv.ParseInt(idVal, 10, 64)
				if val != idValInt64 {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["priority"]; ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	// Matching item not found (deleted out-of-band); signal removal via null Id.
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
