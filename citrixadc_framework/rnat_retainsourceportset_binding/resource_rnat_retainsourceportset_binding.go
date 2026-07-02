package rnat_retainsourceportset_binding

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
var _ resource.Resource = &RnatRetainsourceportsetBindingResource{}
var _ resource.ResourceWithConfigure = (*RnatRetainsourceportsetBindingResource)(nil)
var _ resource.ResourceWithImportState = (*RnatRetainsourceportsetBindingResource)(nil)

func NewRnatRetainsourceportsetBindingResource() resource.Resource {
	return &RnatRetainsourceportsetBindingResource{}
}

// RnatRetainsourceportsetBindingResource defines the resource implementation.
type RnatRetainsourceportsetBindingResource struct {
	client *service.NitroClient
}

func (r *RnatRetainsourceportsetBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RnatRetainsourceportsetBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rnat_retainsourceportset_binding"
}

func (r *RnatRetainsourceportsetBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RnatRetainsourceportsetBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RnatRetainsourceportsetBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rnat_retainsourceportset_binding resource")
	rnat_retainsourceportset_binding := rnat_retainsourceportset_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Rnat_retainsourceportset_binding.Type(), &rnat_retainsourceportset_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create rnat_retainsourceportset_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created rnat_retainsourceportset_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("retainsourceportrange:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Retainsourceportrange.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readRnatRetainsourceportsetBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatRetainsourceportsetBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RnatRetainsourceportsetBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading rnat_retainsourceportset_binding resource")

	r.readRnatRetainsourceportsetBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatRetainsourceportsetBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state RnatRetainsourceportsetBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for rnat_retainsourceportset_binding; the binding is immutable
	// (NITRO exposes only add/delete/get, no set/change action) and both attributes are
	// RequiresReplace, so any change forces recreation and Update is never reached with a
	// real diff.
	tflog.Debug(ctx, "Update is a no-op for rnat_retainsourceportset_binding; binding is immutable (bind/unbind only)")

	// Read the updated state back
	r.readRnatRetainsourceportsetBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatRetainsourceportsetBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RnatRetainsourceportsetBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting rnat_retainsourceportset_binding resource")
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
	if val, ok := idMap["retainsourceportrange"]; ok && val != "" {
		argsMap["retainsourceportrange"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Rnat_retainsourceportset_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete rnat_retainsourceportset_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted rnat_retainsourceportset_binding binding")
}

// Helper function to read rnat_retainsourceportset_binding data from API
func (r *RnatRetainsourceportsetBindingResource) readRnatRetainsourceportsetBindingFromApi(ctx context.Context, data *RnatRetainsourceportsetBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Rnat_retainsourceportset_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read rnat_retainsourceportset_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "rnat_retainsourceportset_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check retainsourceportrange
		if idVal, ok := idMap["retainsourceportrange"]; ok {
			if val, ok := v["retainsourceportrange"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["retainsourceportrange"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("rnat_retainsourceportset_binding not found with the provided ID attributes"))
		return
	}

	rnat_retainsourceportset_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
