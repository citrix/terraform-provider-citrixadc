package lbvserver_tmtrafficpolicy_binding

import (
	"context"
	"fmt"
	"net/url"
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
var _ resource.Resource = &LbvserverTmtrafficpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*LbvserverTmtrafficpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbvserverTmtrafficpolicyBindingResource)(nil)

func NewLbvserverTmtrafficpolicyBindingResource() resource.Resource {
	return &LbvserverTmtrafficpolicyBindingResource{}
}

// LbvserverTmtrafficpolicyBindingResource defines the resource implementation.
type LbvserverTmtrafficpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *LbvserverTmtrafficpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbvserverTmtrafficpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbvserver_tmtrafficpolicy_binding"
}

func (r *LbvserverTmtrafficpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbvserverTmtrafficpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbvserverTmtrafficpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbvserver_tmtrafficpolicy_binding resource")
	lbvserver_tmtrafficpolicy_binding := lbvserver_tmtrafficpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - SDK v2 used AddResource (POST), matching NITRO's add verb
	_, err := r.client.AddResource(service.Lbvserver_tmtrafficpolicy_binding.Type(), data.Name.ValueString(), &lbvserver_tmtrafficpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbvserver_tmtrafficpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lbvserver_tmtrafficpolicy_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readLbvserverTmtrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverTmtrafficpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbvserverTmtrafficpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbvserver_tmtrafficpolicy_binding resource")

	r.readLbvserverTmtrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverTmtrafficpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LbvserverTmtrafficpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lbvserver_tmtrafficpolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		lbvserver_tmtrafficpolicy_binding := lbvserver_tmtrafficpolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Lbvserver_tmtrafficpolicy_binding.Type(), &lbvserver_tmtrafficpolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbvserver_tmtrafficpolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated lbvserver_tmtrafficpolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lbvserver_tmtrafficpolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readLbvserverTmtrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverTmtrafficpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbvserverTmtrafficpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbvserver_tmtrafficpolicy_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "policyname"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["policyname"]; ok && val != "" {
		// URL-encode the delete arg value (SDK v2 used url.QueryEscape) so slashy/special
		// values (e.g. bindpoint REQUEST/RESPONSE, expression policynames) are handled.
		argsMap["policyname"] = url.QueryEscape(val)
	}
	if !data.Bindpoint.IsNull() && data.Bindpoint.ValueString() != "" {
		argsMap["bindpoint"] = url.QueryEscape(data.Bindpoint.ValueString())
	}
	if !data.Priority.IsNull() {
		argsMap["priority"] = url.QueryEscape(fmt.Sprintf("%v", data.Priority.ValueInt64()))
	}

	err = r.client.DeleteResourceWithArgsMap(service.Lbvserver_tmtrafficpolicy_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lbvserver_tmtrafficpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lbvserver_tmtrafficpolicy_binding binding")
}

// Helper function to read lbvserver_tmtrafficpolicy_binding data from API
func (r *LbvserverTmtrafficpolicyBindingResource) readLbvserverTmtrafficpolicyBindingFromApi(ctx context.Context, data *LbvserverTmtrafficpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "policyname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	name_Name, ok := idMap["name"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'name' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Lbvserver_tmtrafficpolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbvserver_tmtrafficpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "lbvserver_tmtrafficpolicy_binding returned empty array.")
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

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("lbvserver_tmtrafficpolicy_binding not found with the provided ID attributes"))
		return
	}

	lbvserver_tmtrafficpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
