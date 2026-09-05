package lsngroup_lsntransportprofile_binding

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
var _ resource.Resource = &LsngroupLsntransportprofileBindingResource{}
var _ resource.ResourceWithConfigure = (*LsngroupLsntransportprofileBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LsngroupLsntransportprofileBindingResource)(nil)

func NewLsngroupLsntransportprofileBindingResource() resource.Resource {
	return &LsngroupLsntransportprofileBindingResource{}
}

// LsngroupLsntransportprofileBindingResource defines the resource implementation.
type LsngroupLsntransportprofileBindingResource struct {
	client *service.NitroClient
}

func (r *LsngroupLsntransportprofileBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsngroupLsntransportprofileBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsngroup_lsntransportprofile_binding"
}

func (r *LsngroupLsntransportprofileBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsngroupLsntransportprofileBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsngroupLsntransportprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsngroup_lsntransportprofile_binding resource")
	lsngroup_lsntransportprofile_binding := lsngroup_lsntransportprofile_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Lsngroup_lsntransportprofile_binding.Type(), &lsngroup_lsntransportprofile_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsngroup_lsntransportprofile_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lsngroup_lsntransportprofile_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("transportprofilename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Transportprofilename.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readLsngroupLsntransportprofileBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "lsngroup_lsntransportprofile_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsntransportprofileBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsngroupLsntransportprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsngroup_lsntransportprofile_binding resource")

	r.readLsngroupLsntransportprofileBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	// Binding is gone on the ADC (readFromApi nulled the Id): drop it from state so a
	// subsequent apply recreates it, matching the SDK v2 provider's behaviour.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsntransportprofileBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsngroupLsntransportprofileBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lsngroup_lsntransportprofile_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		lsngroup_lsntransportprofile_binding := lsngroup_lsntransportprofile_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Lsngroup_lsntransportprofile_binding.Type(), &lsngroup_lsntransportprofile_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsngroup_lsntransportprofile_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated lsngroup_lsntransportprofile_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lsngroup_lsntransportprofile_binding resource, skipping update")
	}

	// Read the updated state back
	r.readLsngroupLsntransportprofileBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "lsngroup_lsntransportprofile_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsntransportprofileBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsngroupLsntransportprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsngroup_lsntransportprofile_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"groupname", "transportprofilename"}, nil)
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
	if val, ok := idMap["transportprofilename"]; ok && val != "" {
		argsMap["transportprofilename"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Lsngroup_lsntransportprofile_binding.Type(), groupname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsngroup_lsntransportprofile_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsngroup_lsntransportprofile_binding binding")
}

// Helper function to read lsngroup_lsntransportprofile_binding data from API
func (r *LsngroupLsntransportprofileBindingResource) readLsngroupLsntransportprofileBindingFromApi(ctx context.Context, data *LsngroupLsntransportprofileBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"groupname", "transportprofilename"}, nil)
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
		ResourceType:             service.Lsngroup_lsntransportprofile_binding.Type(),
		ResourceName:             groupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsngroup_lsntransportprofile_binding, got error: %s", err))
		return
	}

	// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
	// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check transportprofilename
		if idVal, ok := idMap["transportprofilename"]; ok {
			if val, ok := v["transportprofilename"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["transportprofilename"].(string); ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	// Binding not present in the returned set: signal removal via a null Id (see above).
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	lsngroup_lsntransportprofile_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
