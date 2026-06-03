package videooptimizationglobalpacing_videooptimizationpacingpolicy_binding

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
var _ resource.Resource = &VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResource)(nil)

func NewVideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResource() resource.Resource {
	return &VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResource{}
}

// VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResource defines the resource implementation.
type VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_videooptimizationglobalpacing_videooptimizationpacingpolicy_binding"
}

func (r *VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating videooptimizationglobalpacing_videooptimizationpacingpolicy_binding resource")
	videooptimizationglobalpacing_videooptimizationpacingpolicy_binding := videooptimizationglobalpacing_videooptimizationpacingpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Videooptimizationglobalpacing_videooptimizationpacingpolicy_binding.Type(), &videooptimizationglobalpacing_videooptimizationpacingpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create videooptimizationglobalpacing_videooptimizationpacingpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created videooptimizationglobalpacing_videooptimizationpacingpolicy_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("priority:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Priority.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Type.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readVideooptimizationglobalpacingVideooptimizationpacingpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading videooptimizationglobalpacing_videooptimizationpacingpolicy_binding resource")

	r.readVideooptimizationglobalpacingVideooptimizationpacingpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResourceModel

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
	// NOTE: video pacing is marked deprecated by the NITRO/CLI, but the binding is still
	// wired correctly here for backward compatibility.
	tflog.Debug(ctx, "Update is a no-op for videooptimizationglobalpacing_videooptimizationpacingpolicy_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readVideooptimizationglobalpacingVideooptimizationpacingpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting videooptimizationglobalpacing_videooptimizationpacingpolicy_binding resource")
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

	err = r.client.DeleteResourceWithArgsMap(service.Videooptimizationglobalpacing_videooptimizationpacingpolicy_binding.Type(), "", argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete videooptimizationglobalpacing_videooptimizationpacingpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted videooptimizationglobalpacing_videooptimizationpacingpolicy_binding binding")
}

// Helper function to read videooptimizationglobalpacing_videooptimizationpacingpolicy_binding data from API
func (r *VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResource) readVideooptimizationglobalpacingVideooptimizationpacingpolicyBindingFromApi(ctx context.Context, data *VideooptimizationglobalpacingVideooptimizationpacingpolicyBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Videooptimizationglobalpacing_videooptimizationpacingpolicy_binding.Type(),
		ArgsMap:                  argsMap,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read videooptimizationglobalpacing_videooptimizationpacingpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "videooptimizationglobalpacing_videooptimizationpacingpolicy_binding returned empty array")
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

	// Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("videooptimizationglobalpacing_videooptimizationpacingpolicy_binding not found with the provided ID attributes"))
		return
	}

	videooptimizationglobalpacing_videooptimizationpacingpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
