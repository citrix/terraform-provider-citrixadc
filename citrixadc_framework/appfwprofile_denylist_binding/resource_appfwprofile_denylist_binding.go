package appfwprofile_denylist_binding

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
var _ resource.Resource = &AppfwprofileDenylistBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileDenylistBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileDenylistBindingResource)(nil)

func NewAppfwprofileDenylistBindingResource() resource.Resource {
	return &AppfwprofileDenylistBindingResource{}
}

// AppfwprofileDenylistBindingResource defines the resource implementation.
type AppfwprofileDenylistBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileDenylistBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileDenylistBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_denylist_binding"
}

func (r *AppfwprofileDenylistBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileDenylistBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileDenylistBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_denylist_binding resource")
	appfwprofile_denylist_binding := appfwprofile_denylist_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwprofile_denylist_binding.Type(), &appfwprofile_denylist_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_denylist_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile_denylist_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("as_deny_list:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsDenyList.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_deny_list_location:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsDenyListLocation.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_deny_list_value_type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsDenyListValueType.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAppfwprofileDenylistBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileDenylistBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileDenylistBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_denylist_binding resource")

	r.readAppfwprofileDenylistBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileDenylistBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileDenylistBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Pattern 5: Update is a no-op for this binding. NITRO exposes no update
	// endpoint for appfwprofile_denylist_binding (only add/delete/get), and
	// every schema attribute is RequiresReplace, so Terraform never invokes Update
	// with an actual changed value. Just re-read and persist state.
	tflog.Debug(ctx, "Update is a no-op for appfwprofile_denylist_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readAppfwprofileDenylistBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileDenylistBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileDenylistBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_denylist_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
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
	if val, ok := idMap["as_deny_list"]; ok && val != "" {
		// as_deny_list is a value that may be a URL, PCRE, or Expression containing
		// reserved characters. nitro-go does not encode ?args= values, so encode it to
		// avoid a 400 from NITRO.
		argsMap["as_deny_list"] = utils.UrlEncode(val)
	}
	if val, ok := idMap["as_deny_list_location"]; ok && val != "" {
		argsMap["as_deny_list_location"] = val
	}
	if val, ok := idMap["as_deny_list_value_type"]; ok && val != "" {
		argsMap["as_deny_list_value_type"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Appfwprofile_denylist_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile_denylist_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile_denylist_binding binding")
}

// Helper function to read appfwprofile_denylist_binding data from API
func (r *AppfwprofileDenylistBindingResource) readAppfwprofileDenylistBindingFromApi(ctx context.Context, data *AppfwprofileDenylistBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
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
		ResourceType:             service.Appfwprofile_denylist_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_denylist_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "appfwprofile_denylist_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check as_deny_list
		if idVal, ok := idMap["as_deny_list"]; ok {
			if val, ok := v["as_deny_list"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_deny_list"].(string); ok {
			match = false
			continue
		}

		// Check as_deny_list_location
		if idVal, ok := idMap["as_deny_list_location"]; ok {
			if val, ok := v["as_deny_list_location"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_deny_list_location"].(string); ok {
			match = false
			continue
		}

		// Check as_deny_list_value_type
		if idVal, ok := idMap["as_deny_list_value_type"]; ok {
			if val, ok := v["as_deny_list_value_type"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_deny_list_value_type"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("appfwprofile_denylist_binding not found with the provided ID attributes"))
		return
	}

	appfwprofile_denylist_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
