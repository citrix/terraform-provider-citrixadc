package videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding

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
var _ resource.Resource = &VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResource)(nil)

func NewVideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResource() resource.Resource {
	return &VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResource{}
}

// VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResource defines the resource implementation.
type VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding"
}

func (r *VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding resource")
	videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding := videooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.Type(), &videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("priority:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Priority.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Type.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readVideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding resource")

	r.readVideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for this binding: NITRO exposes no update endpoint and every
	// schema attribute uses RequiresReplace, so Terraform never reaches Update with an
	// in-place change. Re-read current state from the API to refresh computed fields.
	tflog.Debug(ctx, "Update is a no-op for videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readVideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Multiple unique attributes - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["policyname"]; ok && val != "" {
		argsMap["policyname"] = val
	}
	if val, ok := idMap["priority"]; ok && val != "" {
		argsMap["priority"] = val
	}
	if val, ok := idMap["type"]; ok && val != "" {
		argsMap["type"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.Type(), "", argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding binding")
}

// Helper function to read videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding data from API
func (r *VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResource) readVideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingFromApi(ctx context.Context, data *VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 3: Array filter without parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	var dataArr []map[string]interface{}
	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["type"]; ok && val != "" {
		argsMap["type"] = val
	}

	findParams := service.FindParams{
		ResourceType:             service.Videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.Type(),
		ArgsMap:                  argsMap,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding, got error: %s", err))
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
		// Check type
		if val, ok := idMap["type"]; ok && val != "" {
			if v, ok := v["type"]; ok {
				if v.(string) != val {
					match = false
				}
			}
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

	videooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
