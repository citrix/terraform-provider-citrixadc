package cmpglobal_cmppolicy_binding

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
var _ resource.Resource = &CmpglobalCmppolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*CmpglobalCmppolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CmpglobalCmppolicyBindingResource)(nil)

func NewCmpglobalCmppolicyBindingResource() resource.Resource {
	return &CmpglobalCmppolicyBindingResource{}
}

// CmpglobalCmppolicyBindingResource defines the resource implementation.
type CmpglobalCmppolicyBindingResource struct {
	client *service.NitroClient
}

func (r *CmpglobalCmppolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CmpglobalCmppolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cmpglobal_cmppolicy_binding"
}

func (r *CmpglobalCmppolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CmpglobalCmppolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CmpglobalCmppolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cmpglobal_cmppolicy_binding resource")
	cmpglobal_cmppolicy_binding := cmpglobal_cmppolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Cmpglobal_cmppolicy_binding.Type(), &cmpglobal_cmppolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cmpglobal_cmppolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created cmpglobal_cmppolicy_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Type.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readCmpglobalCmppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CmpglobalCmppolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CmpglobalCmppolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cmpglobal_cmppolicy_binding resource")

	r.readCmpglobalCmppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CmpglobalCmppolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CmpglobalCmppolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating cmpglobal_cmppolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		cmpglobal_cmppolicy_binding := cmpglobal_cmppolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Cmpglobal_cmppolicy_binding.Type(), &cmpglobal_cmppolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cmpglobal_cmppolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated cmpglobal_cmppolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for cmpglobal_cmppolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readCmpglobalCmppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CmpglobalCmppolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CmpglobalCmppolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cmpglobal_cmppolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Multiple unique attributes - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"policyname"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	// Build delete args, URL-encoding values for slashy/special characters.
	// deleteResourceWithArgs joins key:value directly into the URL without
	// encoding, so we must QueryEscape the values ourselves (matches SDK v2).
	args := make([]string, 0)
	if val, ok := idMap["policyname"]; ok && val != "" {
		args = append(args, fmt.Sprintf("policyname:%s", url.QueryEscape(val)))
	}
	if val, ok := idMap["type"]; ok && val != "" {
		args = append(args, fmt.Sprintf("type:%s", url.QueryEscape(val)))
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		args = append(args, fmt.Sprintf("priority:%d", data.Priority.ValueInt64()))
	}

	err = r.client.DeleteResourceWithArgs(service.Cmpglobal_cmppolicy_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete cmpglobal_cmppolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted cmpglobal_cmppolicy_binding binding")
}

// Helper function to read cmpglobal_cmppolicy_binding data from API
func (r *CmpglobalCmppolicyBindingResource) readCmpglobalCmppolicyBindingFromApi(ctx context.Context, data *CmpglobalCmppolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 3: Array filter without parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"policyname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	var dataArr []map[string]interface{}
	var argsMap map[string]string = make(map[string]string)
	// The unfiltered GET only returns aggregate (numpol) entries with no
	// policyname; individual bindings are only returned when the request is
	// filtered by the bindpoint "type". When type is not part of the ID, fall
	// back to the NITRO default bindpoint "RES_DEFAULT" (matches SDK v2 read).
	if val, ok := idMap["type"]; ok && val != "" {
		argsMap["type"] = url.QueryEscape(val)
	} else {
		argsMap["type"] = url.QueryEscape("RES_DEFAULT")
	}

	findParams := service.FindParams{
		ResourceType:             service.Cmpglobal_cmppolicy_binding.Type(),
		ArgsMap:                  argsMap,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cmpglobal_cmppolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "cmpglobal_cmppolicy_binding returned empty array")
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
		diags.AddError("Client Error", fmt.Sprintf("cmpglobal_cmppolicy_binding not found with the provided ID attributes"))
		return
	}

	cmpglobal_cmppolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
