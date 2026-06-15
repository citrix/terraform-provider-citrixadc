package rewritepolicylabel_rewritepolicy_binding

import (
	"context"
	"fmt"
	"net/url"
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
var _ resource.Resource = &RewritepolicylabelRewritepolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*RewritepolicylabelRewritepolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*RewritepolicylabelRewritepolicyBindingResource)(nil)

func NewRewritepolicylabelRewritepolicyBindingResource() resource.Resource {
	return &RewritepolicylabelRewritepolicyBindingResource{}
}

// RewritepolicylabelRewritepolicyBindingResource defines the resource implementation.
type RewritepolicylabelRewritepolicyBindingResource struct {
	client *service.NitroClient
}

func (r *RewritepolicylabelRewritepolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RewritepolicylabelRewritepolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rewritepolicylabel_rewritepolicy_binding"
}

func (r *RewritepolicylabelRewritepolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RewritepolicylabelRewritepolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RewritepolicylabelRewritepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rewritepolicylabel_rewritepolicy_binding resource")
	rewritepolicylabel_rewritepolicy_binding := rewritepolicylabel_rewritepolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - NITRO `add` is HTTP POST (matches SDK v2 AddResource). Pattern 1.
	_, err := r.client.AddResource(service.Rewritepolicylabel_rewritepolicy_binding.Type(), data.Labelname.ValueString(), &rewritepolicylabel_rewritepolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create rewritepolicylabel_rewritepolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created rewritepolicylabel_rewritepolicy_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("labelname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Labelname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("priority:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Priority.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readRewritepolicylabelRewritepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RewritepolicylabelRewritepolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RewritepolicylabelRewritepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading rewritepolicylabel_rewritepolicy_binding resource")

	r.readRewritepolicylabelRewritepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RewritepolicylabelRewritepolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state RewritepolicylabelRewritepolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for this binding; every schema attribute is RequiresReplace,
	// so Terraform recreates the resource instead of ever reaching a real update.
	// There is no NITRO update endpoint for the binding. Pattern 5.
	tflog.Debug(ctx, "Update is a no-op for rewritepolicylabel_rewritepolicy_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readRewritepolicylabelRewritepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RewritepolicylabelRewritepolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RewritepolicylabelRewritepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting rewritepolicylabel_rewritepolicy_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"labelname", "policyname"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	labelname_value, ok := idMap["labelname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'labelname' not found in ID")
		return
	}

	// URL-encode the delete arg values: ParseIdString returns decoded values, and
	// DeleteResourceWithArgsMap does not escape them, so slashy/special values
	// (e.g. policy names with reserved chars) would break the NITRO query. Matches
	// the SDK v2 url.QueryEscape behavior. Pattern (b).
	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["policyname"]; ok && val != "" {
		argsMap["policyname"] = url.QueryEscape(val)
	}
	if val, ok := idMap["priority"]; ok && val != "" {
		argsMap["priority"] = url.QueryEscape(val)
	}

	err = r.client.DeleteResourceWithArgsMap(service.Rewritepolicylabel_rewritepolicy_binding.Type(), labelname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete rewritepolicylabel_rewritepolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted rewritepolicylabel_rewritepolicy_binding binding")
}

// Helper function to read rewritepolicylabel_rewritepolicy_binding data from API
func (r *RewritepolicylabelRewritepolicyBindingResource) readRewritepolicylabelRewritepolicyBindingFromApi(ctx context.Context, data *RewritepolicylabelRewritepolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"labelname", "policyname"}, nil)
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
		ResourceType:             service.Rewritepolicylabel_rewritepolicy_binding.Type(),
		ResourceName:             labelname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read rewritepolicylabel_rewritepolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "rewritepolicylabel_rewritepolicy_binding returned empty array.")
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

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("rewritepolicylabel_rewritepolicy_binding not found with the provided ID attributes"))
		return
	}

	rewritepolicylabel_rewritepolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
