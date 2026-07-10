package lsngroup_lsnhttphdrlogprofile_binding

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
var _ resource.Resource = &LsngroupLsnhttphdrlogprofileBindingResource{}
var _ resource.ResourceWithConfigure = (*LsngroupLsnhttphdrlogprofileBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LsngroupLsnhttphdrlogprofileBindingResource)(nil)

func NewLsngroupLsnhttphdrlogprofileBindingResource() resource.Resource {
	return &LsngroupLsnhttphdrlogprofileBindingResource{}
}

// LsngroupLsnhttphdrlogprofileBindingResource defines the resource implementation.
type LsngroupLsnhttphdrlogprofileBindingResource struct {
	client *service.NitroClient
}

func (r *LsngroupLsnhttphdrlogprofileBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsngroupLsnhttphdrlogprofileBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsngroup_lsnhttphdrlogprofile_binding"
}

func (r *LsngroupLsnhttphdrlogprofileBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsngroupLsnhttphdrlogprofileBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsngroupLsnhttphdrlogprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsngroup_lsnhttphdrlogprofile_binding resource")
	lsngroup_lsnhttphdrlogprofile_binding := lsngroup_lsnhttphdrlogprofile_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Lsngroup_lsnhttphdrlogprofile_binding.Type(), &lsngroup_lsnhttphdrlogprofile_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsngroup_lsnhttphdrlogprofile_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lsngroup_lsnhttphdrlogprofile_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("httphdrlogprofilename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Httphdrlogprofilename.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readLsngroupLsnhttphdrlogprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsnhttphdrlogprofileBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsngroupLsnhttphdrlogprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsngroup_lsnhttphdrlogprofile_binding resource")

	r.readLsngroupLsnhttphdrlogprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsnhttphdrlogprofileBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsngroupLsnhttphdrlogprofileBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lsngroup_lsnhttphdrlogprofile_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		lsngroup_lsnhttphdrlogprofile_binding := lsngroup_lsnhttphdrlogprofile_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Lsngroup_lsnhttphdrlogprofile_binding.Type(), &lsngroup_lsnhttphdrlogprofile_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsngroup_lsnhttphdrlogprofile_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated lsngroup_lsnhttphdrlogprofile_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lsngroup_lsnhttphdrlogprofile_binding resource, skipping update")
	}

	// Read the updated state back
	r.readLsngroupLsnhttphdrlogprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsnhttphdrlogprofileBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsngroupLsnhttphdrlogprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsngroup_lsnhttphdrlogprofile_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"groupname", "httphdrlogprofilename"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	groupname_value, ok := idMap["groupname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'groupname' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["httphdrlogprofilename"]; ok && val != "" {
		argsMap["httphdrlogprofilename"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Lsngroup_lsnhttphdrlogprofile_binding.Type(), groupname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsngroup_lsnhttphdrlogprofile_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsngroup_lsnhttphdrlogprofile_binding binding")
}

// Helper function to read lsngroup_lsnhttphdrlogprofile_binding data from API
func (r *LsngroupLsnhttphdrlogprofileBindingResource) readLsngroupLsnhttphdrlogprofileBindingFromApi(ctx context.Context, data *LsngroupLsnhttphdrlogprofileBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"groupname", "httphdrlogprofilename"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	groupname_Name, ok := idMap["groupname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'groupname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Lsngroup_lsnhttphdrlogprofile_binding.Type(),
		ResourceName:             groupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsngroup_lsnhttphdrlogprofile_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "lsngroup_lsnhttphdrlogprofile_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check httphdrlogprofilename
		if idVal, ok := idMap["httphdrlogprofilename"]; ok {
			if val, ok := v["httphdrlogprofilename"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["httphdrlogprofilename"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("lsngroup_lsnhttphdrlogprofile_binding not found with the provided ID attributes"))
		return
	}

	lsngroup_lsnhttphdrlogprofile_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
