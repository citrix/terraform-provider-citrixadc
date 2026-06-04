package appfwprofile_bypasslist_binding

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
var _ resource.Resource = &AppfwprofileBypasslistBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileBypasslistBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileBypasslistBindingResource)(nil)

func NewAppfwprofileBypasslistBindingResource() resource.Resource {
	return &AppfwprofileBypasslistBindingResource{}
}

// AppfwprofileBypasslistBindingResource defines the resource implementation.
type AppfwprofileBypasslistBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileBypasslistBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileBypasslistBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_bypasslist_binding"
}

func (r *AppfwprofileBypasslistBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileBypasslistBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileBypasslistBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_bypasslist_binding resource")
	appfwprofile_bypasslist_binding := appfwprofile_bypasslist_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwprofile_bypasslist_binding.Type(), &appfwprofile_bypasslist_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_bypasslist_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile_bypasslist_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("as_bypass_list:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsBypassList.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_bypass_list_location:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsBypassListLocation.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("as_bypass_list_value_type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsBypassListValueType.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAppfwprofileBypasslistBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileBypasslistBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileBypasslistBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_bypasslist_binding resource")

	r.readAppfwprofileBypasslistBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileBypasslistBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileBypasslistBindingResourceModel

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
	// endpoint for appfwprofile_bypasslist_binding (only add/delete/get), and
	// every schema attribute is RequiresReplace, so Terraform never invokes Update
	// with an actual changed value. Just re-read and persist state.
	tflog.Debug(ctx, "Update is a no-op for appfwprofile_bypasslist_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readAppfwprofileBypasslistBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileBypasslistBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileBypasslistBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_bypasslist_binding resource")
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
	if val, ok := idMap["as_bypass_list"]; ok && val != "" {
		// as_bypass_list is a value that may be a URL, PCRE, or Expression containing
		// reserved characters. nitro-go does not encode ?args= values, so encode it to
		// avoid a 400 from NITRO.
		argsMap["as_bypass_list"] = utils.UrlEncode(val)
	}
	if val, ok := idMap["as_bypass_list_location"]; ok && val != "" {
		argsMap["as_bypass_list_location"] = val
	}
	if val, ok := idMap["as_bypass_list_value_type"]; ok && val != "" {
		argsMap["as_bypass_list_value_type"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Appfwprofile_bypasslist_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile_bypasslist_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile_bypasslist_binding binding")
}

// Helper function to read appfwprofile_bypasslist_binding data from API
func (r *AppfwprofileBypasslistBindingResource) readAppfwprofileBypasslistBindingFromApi(ctx context.Context, data *AppfwprofileBypasslistBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Appfwprofile_bypasslist_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_bypasslist_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "appfwprofile_bypasslist_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check as_bypass_list
		if idVal, ok := idMap["as_bypass_list"]; ok {
			if val, ok := v["as_bypass_list"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_bypass_list"].(string); ok {
			match = false
			continue
		}

		// Check as_bypass_list_location
		if idVal, ok := idMap["as_bypass_list_location"]; ok {
			if val, ok := v["as_bypass_list_location"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_bypass_list_location"].(string); ok {
			match = false
			continue
		}

		// Check as_bypass_list_value_type
		if idVal, ok := idMap["as_bypass_list_value_type"]; ok {
			if val, ok := v["as_bypass_list_value_type"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["as_bypass_list_value_type"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("appfwprofile_bypasslist_binding not found with the provided ID attributes"))
		return
	}

	appfwprofile_bypasslist_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
