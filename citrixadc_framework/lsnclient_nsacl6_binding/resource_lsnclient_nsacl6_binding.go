package lsnclient_nsacl6_binding

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
var _ resource.Resource = &LsnclientNsacl6BindingResource{}
var _ resource.ResourceWithConfigure = (*LsnclientNsacl6BindingResource)(nil)
var _ resource.ResourceWithImportState = (*LsnclientNsacl6BindingResource)(nil)

func NewLsnclientNsacl6BindingResource() resource.Resource {
	return &LsnclientNsacl6BindingResource{}
}

// LsnclientNsacl6BindingResource defines the resource implementation.
type LsnclientNsacl6BindingResource struct {
	client *service.NitroClient
}

func (r *LsnclientNsacl6BindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnclientNsacl6BindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnclient_nsacl6_binding"
}

func (r *LsnclientNsacl6BindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnclientNsacl6BindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnclientNsacl6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnclient_nsacl6_binding resource")
	lsnclient_nsacl6_binding := lsnclient_nsacl6_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Lsnclient_nsacl6_binding.Type(), &lsnclient_nsacl6_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnclient_nsacl6_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lsnclient_nsacl6_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("acl6name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Acl6name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("clientname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Clientname.ValueString()))))
	// td is optional and only part of the binding identity when set.
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		idParts = append(idParts, fmt.Sprintf("td:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Td.ValueInt64()))))
	}
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readLsnclientNsacl6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnclientNsacl6BindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnclientNsacl6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnclient_nsacl6_binding resource")

	r.readLsnclientNsacl6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnclientNsacl6BindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsnclientNsacl6BindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lsnclient_nsacl6_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		lsnclient_nsacl6_binding := lsnclient_nsacl6_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Lsnclient_nsacl6_binding.Type(), &lsnclient_nsacl6_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnclient_nsacl6_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated lsnclient_nsacl6_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lsnclient_nsacl6_binding resource, skipping update")
	}

	// Read the updated state back
	r.readLsnclientNsacl6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnclientNsacl6BindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnclientNsacl6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnclient_nsacl6_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"clientname", "acl6name"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	clientname_value, ok := idMap["clientname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'clientname' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["acl6name"]; ok && val != "" {
		argsMap["acl6name"] = val
	}
	if val, ok := idMap["td"]; ok && val != "" {
		argsMap["td"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Lsnclient_nsacl6_binding.Type(), clientname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsnclient_nsacl6_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsnclient_nsacl6_binding binding")
}

// Helper function to read lsnclient_nsacl6_binding data from API
func (r *LsnclientNsacl6BindingResource) readLsnclientNsacl6BindingFromApi(ctx context.Context, data *LsnclientNsacl6BindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"clientname", "acl6name"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	clientname_Name, ok := idMap["clientname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'clientname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Lsnclient_nsacl6_binding.Type(),
		ResourceName:             clientname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnclient_nsacl6_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "lsnclient_nsacl6_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check acl6name
		if idVal, ok := idMap["acl6name"]; ok {
			if val, ok := v["acl6name"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["acl6name"].(string); ok {
			match = false
			continue
		}

		// Check td - the NITRO GET for this binding does not echo td, so only
		// compare when the record actually carries the field (matching SDK v2,
		// which keyed the binding solely on acl6name).
		if idVal, ok := idMap["td"]; ok {
			if val, ok := v["td"]; ok {
				val, _ = utils.ConvertToInt64(val)
				idValInt64, _ := strconv.ParseInt(idVal, 10, 64)
				if val != idValInt64 {
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

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("lsnclient_nsacl6_binding not found with the provided ID attributes"))
		return
	}

	lsnclient_nsacl6_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
