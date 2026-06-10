package lsngroup_lsnsipalgprofile_binding

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
var _ resource.Resource = &LsngroupLsnsipalgprofileBindingResource{}
var _ resource.ResourceWithConfigure = (*LsngroupLsnsipalgprofileBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LsngroupLsnsipalgprofileBindingResource)(nil)

func NewLsngroupLsnsipalgprofileBindingResource() resource.Resource {
	return &LsngroupLsnsipalgprofileBindingResource{}
}

// LsngroupLsnsipalgprofileBindingResource defines the resource implementation.
type LsngroupLsnsipalgprofileBindingResource struct {
	client *service.NitroClient
}

func (r *LsngroupLsnsipalgprofileBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsngroupLsnsipalgprofileBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsngroup_lsnsipalgprofile_binding"
}

func (r *LsngroupLsnsipalgprofileBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsngroupLsnsipalgprofileBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsngroupLsnsipalgprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsngroup_lsnsipalgprofile_binding resource")
	lsngroup_lsnsipalgprofile_binding := lsngroup_lsnsipalgprofile_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Lsngroup_lsnsipalgprofile_binding.Type(), &lsngroup_lsnsipalgprofile_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsngroup_lsnsipalgprofile_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lsngroup_lsnsipalgprofile_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("sipalgprofilename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Sipalgprofilename.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readLsngroupLsnsipalgprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsnsipalgprofileBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsngroupLsnsipalgprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsngroup_lsnsipalgprofile_binding resource")

	r.readLsngroupLsnsipalgprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsnsipalgprofileBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsngroupLsnsipalgprofileBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for lsngroup_lsnsipalgprofile_binding: NITRO exposes only add (PUT)
	// and delete (no update/change endpoint), and all schema attributes are RequiresReplace, so
	// Terraform recreates the resource on any change rather than calling Update.
	tflog.Debug(ctx, "Update is a no-op for lsngroup_lsnsipalgprofile_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readLsngroupLsnsipalgprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsnsipalgprofileBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsngroupLsnsipalgprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsngroup_lsnsipalgprofile_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs with the parent (groupname) as the
	// resource name and the bound sipalgprofilename passed as an arg. This matches the NITRO delete
	// URL: DELETE .../lsngroup_lsnsipalgprofile_binding/<groupname>?args=sipalgprofilename:<value>
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"groupname", "sipalgprofilename"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	groupname_value, ok := idMap["groupname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'groupname' not found in ID")
		return
	}

	args := make([]string, 0)
	if val, ok := idMap["sipalgprofilename"]; ok && val != "" {
		args = append(args, fmt.Sprintf("sipalgprofilename:%s", utils.UrlEncode(val)))
	}

	err = r.client.DeleteResourceWithArgs(service.Lsngroup_lsnsipalgprofile_binding.Type(), groupname_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsngroup_lsnsipalgprofile_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsngroup_lsnsipalgprofile_binding binding")
}

// Helper function to read lsngroup_lsnsipalgprofile_binding data from API
func (r *LsngroupLsnsipalgprofileBindingResource) readLsngroupLsnsipalgprofileBindingFromApi(ctx context.Context, data *LsngroupLsnsipalgprofileBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"groupname", "sipalgprofilename"}, nil)
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
		ResourceType:             service.Lsngroup_lsnsipalgprofile_binding.Type(),
		ResourceName:             groupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsngroup_lsnsipalgprofile_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "lsngroup_lsnsipalgprofile_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check sipalgprofilename
		if idVal, ok := idMap["sipalgprofilename"]; ok {
			if val, ok := v["sipalgprofilename"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["sipalgprofilename"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("lsngroup_lsnsipalgprofile_binding not found with the provided ID attributes"))
		return
	}

	lsngroup_lsnsipalgprofile_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
