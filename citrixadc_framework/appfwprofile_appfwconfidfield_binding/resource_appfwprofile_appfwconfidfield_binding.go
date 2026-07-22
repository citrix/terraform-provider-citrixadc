package appfwprofile_appfwconfidfield_binding

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
var _ resource.Resource = &AppfwprofileAppfwconfidfieldBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileAppfwconfidfieldBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileAppfwconfidfieldBindingResource)(nil)

func NewAppfwprofileAppfwconfidfieldBindingResource() resource.Resource {
	return &AppfwprofileAppfwconfidfieldBindingResource{}
}

// AppfwprofileAppfwconfidfieldBindingResource defines the resource implementation.
type AppfwprofileAppfwconfidfieldBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileAppfwconfidfieldBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileAppfwconfidfieldBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_appfwconfidfield_binding"
}

func (r *AppfwprofileAppfwconfidfieldBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileAppfwconfidfieldBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileAppfwconfidfieldBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_appfwconfidfield_binding resource")
	appfwprofile_appfwconfidfield_binding := appfwprofile_appfwconfidfield_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwprofile_appfwconfidfield_binding.Type(), &appfwprofile_appfwconfidfield_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_appfwconfidfield_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile_appfwconfidfield_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("cffield_url:%s", utils.UrlEncode(fmt.Sprintf("%v", data.CffieldUrl.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("confidfield:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Confidfield.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAppfwprofileAppfwconfidfieldBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileAppfwconfidfieldBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileAppfwconfidfieldBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_appfwconfidfield_binding resource")

	r.readAppfwprofileAppfwconfidfieldBindingFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Resource was deleted out-of-band - remove from state for self-healing
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileAppfwconfidfieldBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileAppfwconfidfieldBindingResourceModel

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
	// endpoint for appfwprofile_appfwconfidfield_binding (only add/delete/get),
	// and every schema attribute is RequiresReplace, so Terraform never invokes
	// Update with an actual changed value. Just re-read and persist state.
	tflog.Debug(ctx, "Update is a no-op for appfwprofile_appfwconfidfield_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readAppfwprofileAppfwconfidfieldBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileAppfwconfidfieldBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileAppfwconfidfieldBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_appfwconfidfield_binding resource")
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
	if val, ok := idMap["cffield_url"]; ok && val != "" {
		// cffield_url is a URL containing reserved characters (':' and '/').
		// nitro-go does not encode ?args= values, so encode it to avoid a 400 from NITRO.
		argsMap["cffield_url"] = utils.UrlEncode(val)
	}
	if val, ok := idMap["confidfield"]; ok && val != "" {
		argsMap["confidfield"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Appfwprofile_appfwconfidfield_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile_appfwconfidfield_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile_appfwconfidfield_binding binding")
}

// Helper function to read appfwprofile_appfwconfidfield_binding data from API
func (r *AppfwprofileAppfwconfidfieldBindingResource) readAppfwprofileAppfwconfidfieldBindingFromApi(ctx context.Context, data *AppfwprofileAppfwconfidfieldBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Appfwprofile_appfwconfidfield_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_appfwconfidfield_binding, got error: %s", err))
		return
	}

	// Resource is missing - signal deletion for self-healing
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check cffield_url
		if idVal, ok := idMap["cffield_url"]; ok {
			if val, ok := v["cffield_url"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["cffield_url"].(string); ok {
			match = false
			continue
		}

		// Check confidfield
		if idVal, ok := idMap["confidfield"]; ok {
			if val, ok := v["confidfield"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["confidfield"].(string); ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing - signal deletion for self-healing
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	appfwprofile_appfwconfidfield_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])

	// Backfill the identity attributes from the parsed composite ID so that
	// `terraform import` (which has no prior plan/state) repopulates them. These
	// are the exact values the ID was composed from, so the ID stays identical.
	data.Name = types.StringValue(name_Name)
	if val, ok := idMap["confidfield"]; ok && val != "" {
		data.Confidfield = types.StringValue(val)
	}
	if val, ok := idMap["cffield_url"]; ok && val != "" {
		data.CffieldUrl = types.StringValue(val)
	}
}
