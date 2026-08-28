package vpnglobal_vpnsecureprivateaccessprofile_binding

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
var _ resource.Resource = &VpnglobalVpnsecureprivateaccessprofileBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalVpnsecureprivateaccessprofileBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalVpnsecureprivateaccessprofileBindingResource)(nil)

func NewVpnglobalVpnsecureprivateaccessprofileBindingResource() resource.Resource {
	return &VpnglobalVpnsecureprivateaccessprofileBindingResource{}
}

// VpnglobalVpnsecureprivateaccessprofileBindingResource defines the resource implementation.
type VpnglobalVpnsecureprivateaccessprofileBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalVpnsecureprivateaccessprofileBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalVpnsecureprivateaccessprofileBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_vpnsecureprivateaccessprofile_binding"
}

func (r *VpnglobalVpnsecureprivateaccessprofileBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalVpnsecureprivateaccessprofileBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalVpnsecureprivateaccessprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_vpnsecureprivateaccessprofile_binding resource")
	vpnglobal_vpnsecureprivateaccessprofile_binding := vpnglobal_vpnsecureprivateaccessprofile_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vpnglobal_vpnsecureprivateaccessprofile_binding.Type(), &vpnglobal_vpnsecureprivateaccessprofile_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_vpnsecureprivateaccessprofile_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnglobal_vpnsecureprivateaccessprofile_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Secureprivateaccessprofile.ValueString()))

	// Read the updated state back
	r.readVpnglobalVpnsecureprivateaccessprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalVpnsecureprivateaccessprofileBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalVpnsecureprivateaccessprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_vpnsecureprivateaccessprofile_binding resource")

	r.readVpnglobalVpnsecureprivateaccessprofileBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Object deleted out-of-band: remove from state so a subsequent apply re-creates it.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalVpnsecureprivateaccessprofileBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnglobalVpnsecureprivateaccessprofileBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for vpnglobal_vpnsecureprivateaccessprofile_binding; NITRO exposes no
	// update endpoint and all attributes are RequiresReplace.
	tflog.Debug(ctx, "Update is a no-op for vpnglobal_vpnsecureprivateaccessprofile_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readVpnglobalVpnsecureprivateaccessprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalVpnsecureprivateaccessprofileBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalVpnsecureprivateaccessprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_vpnsecureprivateaccessprofile_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Single unique attribute - ID is the plain value
	secureprivateaccessprofile_value := data.Id.ValueString()
	args := []string{
		fmt.Sprintf("secureprivateaccessprofile:%s", secureprivateaccessprofile_value),
	}

	err := r.client.DeleteResourceWithArgs(service.Vpnglobal_vpnsecureprivateaccessprofile_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnglobal_vpnsecureprivateaccessprofile_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnglobal_vpnsecureprivateaccessprofile_binding binding")
}

// Helper function to read vpnglobal_vpnsecureprivateaccessprofile_binding data from API
func (r *VpnglobalVpnsecureprivateaccessprofileBindingResource) readVpnglobalVpnsecureprivateaccessprofileBindingFromApi(ctx context.Context, data *VpnglobalVpnsecureprivateaccessprofileBindingResourceModel, diags *diag.Diagnostics) {

	// Single unique attribute - ID is the plain secureprivateaccessprofile value
	secureprivateaccessprofile_value := data.Id.ValueString()

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Vpnglobal_vpnsecureprivateaccessprofile_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_vpnsecureprivateaccessprofile_binding, got error: %s", err))
		return
	}

	// Resource is missing (deleted out-of-band): signal removal via null Id.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		// Check secureprivateaccessprofile
		if val, ok := v["secureprivateaccessprofile"].(string); ok {
			if val == secureprivateaccessprofile_value {
				foundIndex = i
				break
			}
		}
	}

	// Resource is missing (deleted out-of-band): signal removal via null Id.
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	vpnglobal_vpnsecureprivateaccessprofile_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
