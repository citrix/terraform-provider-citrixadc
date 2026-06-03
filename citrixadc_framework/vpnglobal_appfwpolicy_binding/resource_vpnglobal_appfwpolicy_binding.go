package vpnglobal_appfwpolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VpnglobalAppfwpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalAppfwpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalAppfwpolicyBindingResource)(nil)

func NewVpnglobalAppfwpolicyBindingResource() resource.Resource {
	return &VpnglobalAppfwpolicyBindingResource{}
}

// VpnglobalAppfwpolicyBindingResource defines the resource implementation.
type VpnglobalAppfwpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalAppfwpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalAppfwpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_appfwpolicy_binding"
}

func (r *VpnglobalAppfwpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalAppfwpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalAppfwpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_appfwpolicy_binding resource")
	vpnglobal_appfwpolicy_binding := vpnglobal_appfwpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vpnglobal_appfwpolicy_binding.Type(), &vpnglobal_appfwpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_appfwpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnglobal_appfwpolicy_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	// Read the updated state back
	r.readVpnglobalAppfwpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAppfwpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalAppfwpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_appfwpolicy_binding resource")

	r.readVpnglobalAppfwpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAppfwpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnglobalAppfwpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for vpnglobal_appfwpolicy_binding; NITRO exposes no update endpoint
	// (only add/delete/get) and all attributes are RequiresReplace.
	tflog.Debug(ctx, "Update is a no-op for vpnglobal_appfwpolicy_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readVpnglobalAppfwpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAppfwpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalAppfwpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_appfwpolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// NITRO delete args = policyname, secondary, groupextraction (per NITRO doc)
	policyname_value := data.Id.ValueString()
	args := []string{
		fmt.Sprintf("policyname:%s", policyname_value),
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		args = append(args, fmt.Sprintf("secondary:%t", data.Secondary.ValueBool()))
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		args = append(args, fmt.Sprintf("groupextraction:%t", data.Groupextraction.ValueBool()))
	}

	err := r.client.DeleteResourceWithArgs(service.Vpnglobal_appfwpolicy_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnglobal_appfwpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnglobal_appfwpolicy_binding binding")
}

// Helper function to read vpnglobal_appfwpolicy_binding data from API
func (r *VpnglobalAppfwpolicyBindingResource) readVpnglobalAppfwpolicyBindingFromApi(ctx context.Context, data *VpnglobalAppfwpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// ID is a plain value (single unique attr: policyname) - use it directly as the filter value
	policyname_value := data.Id.ValueString()

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Vpnglobal_appfwpolicy_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_appfwpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "vpnglobal_appfwpolicy_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check policyname
		if val, ok := v["policyname"].(string); ok {
			if val != policyname_value {
				match = false
				continue
			}
		} else {
			match = false
			continue
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("vpnglobal_appfwpolicy_binding not found with the provided ID attributes"))
		return
	}

	vpnglobal_appfwpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
