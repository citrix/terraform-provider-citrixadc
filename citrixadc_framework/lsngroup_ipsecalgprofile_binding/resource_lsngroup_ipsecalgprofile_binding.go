package lsngroup_ipsecalgprofile_binding

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
var _ resource.Resource = &LsngroupIpsecalgprofileBindingResource{}
var _ resource.ResourceWithConfigure = (*LsngroupIpsecalgprofileBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LsngroupIpsecalgprofileBindingResource)(nil)

func NewLsngroupIpsecalgprofileBindingResource() resource.Resource {
	return &LsngroupIpsecalgprofileBindingResource{}
}

// LsngroupIpsecalgprofileBindingResource defines the resource implementation.
type LsngroupIpsecalgprofileBindingResource struct {
	client *service.NitroClient
}

func (r *LsngroupIpsecalgprofileBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsngroupIpsecalgprofileBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsngroup_ipsecalgprofile_binding"
}

func (r *LsngroupIpsecalgprofileBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsngroupIpsecalgprofileBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsngroupIpsecalgprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsngroup_ipsecalgprofile_binding resource")
	lsngroup_ipsecalgprofile_binding := lsngroup_ipsecalgprofile_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Lsngroup_ipsecalgprofile_binding.Type(), &lsngroup_ipsecalgprofile_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsngroup_ipsecalgprofile_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lsngroup_ipsecalgprofile_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("ipsecalgprofile:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ipsecalgprofile.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readLsngroupIpsecalgprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupIpsecalgprofileBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsngroupIpsecalgprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsngroup_ipsecalgprofile_binding resource")

	r.readLsngroupIpsecalgprofileBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Binding genuinely absent on the appliance: treat as drift and clear state.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupIpsecalgprofileBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsngroupIpsecalgprofileBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for lsngroup_ipsecalgprofile_binding: NITRO exposes only add (PUT)
	// and delete (no update/change endpoint), and all schema attributes are RequiresReplace, so
	// Terraform recreates the resource on any change rather than calling Update.
	tflog.Debug(ctx, "Update is a no-op for lsngroup_ipsecalgprofile_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readLsngroupIpsecalgprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupIpsecalgprofileBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsngroupIpsecalgprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsngroup_ipsecalgprofile_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs with the parent (groupname) as the
	// resource name and the bound ipsecalgprofile passed as an arg. This matches the NITRO delete
	// URL: DELETE .../lsngroup_ipsecalgprofile_binding/<groupname>?args=ipsecalgprofile:<value>
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"groupname", "ipsecalgprofile"}, nil)
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
	if val, ok := idMap["ipsecalgprofile"]; ok && val != "" {
		args = append(args, fmt.Sprintf("ipsecalgprofile:%s", utils.UrlEncode(val)))
	}

	err = r.client.DeleteResourceWithArgs(service.Lsngroup_ipsecalgprofile_binding.Type(), groupname_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsngroup_ipsecalgprofile_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsngroup_ipsecalgprofile_binding binding")
}

// Helper function to read lsngroup_ipsecalgprofile_binding data from API
func (r *LsngroupIpsecalgprofileBindingResource) readLsngroupIpsecalgprofileBindingFromApi(ctx context.Context, data *LsngroupIpsecalgprofileBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"groupname", "ipsecalgprofile"}, nil)
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
		ResourceType:             service.Lsngroup_ipsecalgprofile_binding.Type(),
		ResourceName:             groupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsngroup_ipsecalgprofile_binding, got error: %s", err))
		return
	}

	// Resource is missing: signal drift by nulling the Id so Read removes it from state.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ipsecalgprofile
		if idVal, ok := idMap["ipsecalgprofile"]; ok {
			if val, ok := v["ipsecalgprofile"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["ipsecalgprofile"].(string); ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing: signal drift by nulling the Id so Read removes it from state.
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	lsngroup_ipsecalgprofile_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
