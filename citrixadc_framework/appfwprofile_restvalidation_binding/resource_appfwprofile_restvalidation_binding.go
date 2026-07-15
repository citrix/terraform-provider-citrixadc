package appfwprofile_restvalidation_binding

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
var _ resource.Resource = &AppfwprofileRestvalidationBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileRestvalidationBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileRestvalidationBindingResource)(nil)

func NewAppfwprofileRestvalidationBindingResource() resource.Resource {
	return &AppfwprofileRestvalidationBindingResource{}
}

// AppfwprofileRestvalidationBindingResource defines the resource implementation.
type AppfwprofileRestvalidationBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileRestvalidationBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileRestvalidationBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_restvalidation_binding"
}

func (r *AppfwprofileRestvalidationBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileRestvalidationBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileRestvalidationBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_restvalidation_binding resource")
	appfwprofile_restvalidation_binding := appfwprofile_restvalidation_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwprofile_restvalidation_binding.Type(), &appfwprofile_restvalidation_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_restvalidation_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile_restvalidation_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("rest_validation_action:%s", utils.UrlEncode(fmt.Sprintf("%v", data.RestValidationAction.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("restvalidation:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Restvalidation.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAppfwprofileRestvalidationBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileRestvalidationBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileRestvalidationBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_restvalidation_binding resource")

	r.readAppfwprofileRestvalidationBindingFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Self-heal: object was deleted out-of-band, remove from state so a
	// subsequent apply re-creates it.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileRestvalidationBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileRestvalidationBindingResourceModel

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
	// endpoint for appfwprofile_restvalidation_binding (only add/delete/get), and
	// every schema attribute is RequiresReplace, so Terraform never invokes Update
	// with an actual changed value. Just re-read and persist state.
	tflog.Debug(ctx, "Update is a no-op for appfwprofile_restvalidation_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readAppfwprofileRestvalidationBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileRestvalidationBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileRestvalidationBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_restvalidation_binding resource")
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
	if val, ok := idMap["rest_validation_action"]; ok && val != "" {
		argsMap["rest_validation_action"] = val
	}
	if val, ok := idMap["restvalidation"]; ok && val != "" {
		// restvalidation is a pattern like "GET:/v1/bookstore/viewbooks" containing
		// reserved characters (':' and '/'). nitro-go does not encode ?args= values,
		// so encode it to avoid a 400 from NITRO.
		argsMap["restvalidation"] = utils.UrlEncode(val)
	}

	err = r.client.DeleteResourceWithArgsMap(service.Appfwprofile_restvalidation_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile_restvalidation_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile_restvalidation_binding binding")
}

// Helper function to read appfwprofile_restvalidation_binding data from API
func (r *AppfwprofileRestvalidationBindingResource) readAppfwprofileRestvalidationBindingFromApi(ctx context.Context, data *AppfwprofileRestvalidationBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Appfwprofile_restvalidation_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_restvalidation_binding, got error: %s", err))
		return
	}

	// Resource is missing - signal removal for self-heal.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check rest_validation_action
		if idVal, ok := idMap["rest_validation_action"]; ok {
			if val, ok := v["rest_validation_action"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["rest_validation_action"].(string); ok {
			match = false
			continue
		}

		// Check restvalidation
		if idVal, ok := idMap["restvalidation"]; ok {
			if val, ok := v["restvalidation"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["restvalidation"].(string); ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing - signal removal for self-heal.
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	appfwprofile_restvalidation_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])

	// Backfill identity / ID-component attributes from the parsed composite ID so
	// that `terraform import` (which has no prior plan/state) fully round-trips.
	// These are always encoded in the ID and are authoritative. Done after the
	// found/len self-heal checks so a not-found still yields a null Id.
	data.Name = types.StringValue(name_Name)
	if val, ok := idMap["restvalidation"]; ok && val != "" {
		data.Restvalidation = types.StringValue(val)
	}
	if val, ok := idMap["rest_validation_action"]; ok && val != "" {
		data.RestValidationAction = types.StringValue(val)
	} else {
		data.RestValidationAction = types.StringNull()
	}
}
