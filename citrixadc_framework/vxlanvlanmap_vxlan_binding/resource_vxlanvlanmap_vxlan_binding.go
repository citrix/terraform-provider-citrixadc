package vxlanvlanmap_vxlan_binding

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
var _ resource.Resource = &VxlanvlanmapVxlanBindingResource{}
var _ resource.ResourceWithConfigure = (*VxlanvlanmapVxlanBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VxlanvlanmapVxlanBindingResource)(nil)

func NewVxlanvlanmapVxlanBindingResource() resource.Resource {
	return &VxlanvlanmapVxlanBindingResource{}
}

// VxlanvlanmapVxlanBindingResource defines the resource implementation.
type VxlanvlanmapVxlanBindingResource struct {
	client *service.NitroClient
}

func (r *VxlanvlanmapVxlanBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VxlanvlanmapVxlanBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vxlanvlanmap_vxlan_binding"
}

func (r *VxlanvlanmapVxlanBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VxlanvlanmapVxlanBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VxlanvlanmapVxlanBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vxlanvlanmap_vxlan_binding resource")
	vxlanvlanmap_vxlan_binding := vxlanvlanmap_vxlan_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vxlanvlanmap_vxlan_binding.Type(), &vxlanvlanmap_vxlan_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vxlanvlanmap_vxlan_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vxlanvlanmap_vxlan_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vxlan:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vxlan.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readVxlanvlanmapVxlanBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanvlanmapVxlanBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VxlanvlanmapVxlanBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vxlanvlanmap_vxlan_binding resource")

	r.readVxlanvlanmapVxlanBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanvlanmapVxlanBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VxlanvlanmapVxlanBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vxlanvlanmap_vxlan_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		vxlanvlanmap_vxlan_binding := vxlanvlanmap_vxlan_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Vxlanvlanmap_vxlan_binding.Type(), &vxlanvlanmap_vxlan_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vxlanvlanmap_vxlan_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vxlanvlanmap_vxlan_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vxlanvlanmap_vxlan_binding resource, skipping update")
	}

	// Read the updated state back
	r.readVxlanvlanmapVxlanBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanvlanmapVxlanBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VxlanvlanmapVxlanBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vxlanvlanmap_vxlan_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "vxlan"}, nil)
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
	if val, ok := idMap["vxlan"]; ok && val != "" {
		argsMap["vxlan"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Vxlanvlanmap_vxlan_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vxlanvlanmap_vxlan_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vxlanvlanmap_vxlan_binding binding")
}

// Helper function to read vxlanvlanmap_vxlan_binding data from API
func (r *VxlanvlanmapVxlanBindingResource) readVxlanvlanmapVxlanBindingFromApi(ctx context.Context, data *VxlanvlanmapVxlanBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "vxlan"}, nil)
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
		ResourceType:             service.Vxlanvlanmap_vxlan_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vxlanvlanmap_vxlan_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "vxlanvlanmap_vxlan_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check vxlan
		if idVal, ok := idMap["vxlan"]; ok {
			if val, ok := v["vxlan"]; ok {
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
		} else if _, ok := v["vxlan"]; ok {
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
		diags.AddError("Client Error", fmt.Sprintf("vxlanvlanmap_vxlan_binding not found with the provided ID attributes"))
		return
	}

	vxlanvlanmap_vxlan_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
