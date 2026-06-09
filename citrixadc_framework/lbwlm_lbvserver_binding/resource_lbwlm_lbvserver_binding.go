package lbwlm_lbvserver_binding

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
var _ resource.Resource = &LbwlmLbvserverBindingResource{}
var _ resource.ResourceWithConfigure = (*LbwlmLbvserverBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbwlmLbvserverBindingResource)(nil)

func NewLbwlmLbvserverBindingResource() resource.Resource {
	return &LbwlmLbvserverBindingResource{}
}

// LbwlmLbvserverBindingResource defines the resource implementation.
type LbwlmLbvserverBindingResource struct {
	client *service.NitroClient
}

func (r *LbwlmLbvserverBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbwlmLbvserverBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbwlm_lbvserver_binding"
}

func (r *LbwlmLbvserverBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbwlmLbvserverBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbwlmLbvserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbwlm_lbvserver_binding resource")
	lbwlm_lbvserver_binding := lbwlm_lbvserver_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Lbwlm_lbvserver_binding.Type(), &lbwlm_lbvserver_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbwlm_lbvserver_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lbwlm_lbvserver_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("vservername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vservername.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("wlmname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Wlmname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readLbwlmLbvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbwlmLbvserverBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbwlmLbvserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbwlm_lbvserver_binding resource")

	r.readLbwlmLbvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbwlmLbvserverBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LbwlmLbvserverBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for lbwlm_lbvserver_binding: NITRO exposes no update endpoint
	// (bind/unbind only) and all attributes are RequiresReplace.
	tflog.Debug(ctx, "Update is a no-op for lbwlm_lbvserver_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readLbwlmLbvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbwlmLbvserverBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbwlmLbvserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbwlm_lbvserver_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	wlmname_value, ok := idMap["wlmname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'wlmname' not found in ID")
		return
	}

	args := []string{}
	if val, ok := idMap["vservername"]; ok && val != "" {
		args = append(args, "vservername:"+utils.UrlEncode(val))
	}

	err = r.client.DeleteResourceWithArgs(service.Lbwlm_lbvserver_binding.Type(), wlmname_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lbwlm_lbvserver_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lbwlm_lbvserver_binding binding")
}

// Helper function to read lbwlm_lbvserver_binding data from API
func (r *LbwlmLbvserverBindingResource) readLbwlmLbvserverBindingFromApi(ctx context.Context, data *LbwlmLbvserverBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	wlmname_Name, ok := idMap["wlmname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'wlmname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Lbwlm_lbvserver_binding.Type(),
		ResourceName:             wlmname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbwlm_lbvserver_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "lbwlm_lbvserver_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check vservername
		if idVal, ok := idMap["vservername"]; ok {
			if val, ok := v["vservername"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["vservername"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("lbwlm_lbvserver_binding not found with the provided ID attributes"))
		return
	}

	lbwlm_lbvserver_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
