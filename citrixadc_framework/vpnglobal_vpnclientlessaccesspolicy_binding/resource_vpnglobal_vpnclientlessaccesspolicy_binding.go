package vpnglobal_vpnclientlessaccesspolicy_binding

import (
	"context"
	"fmt"
	"net/url"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VpnglobalVpnclientlessaccesspolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalVpnclientlessaccesspolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalVpnclientlessaccesspolicyBindingResource)(nil)

func NewVpnglobalVpnclientlessaccesspolicyBindingResource() resource.Resource {
	return &VpnglobalVpnclientlessaccesspolicyBindingResource{}
}

// VpnglobalVpnclientlessaccesspolicyBindingResource defines the resource implementation.
type VpnglobalVpnclientlessaccesspolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalVpnclientlessaccesspolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalVpnclientlessaccesspolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_vpnclientlessaccesspolicy_binding"
}

func (r *VpnglobalVpnclientlessaccesspolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalVpnclientlessaccesspolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalVpnclientlessaccesspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_vpnclientlessaccesspolicy_binding resource")
	vpnglobal_vpnclientlessaccesspolicy_binding := vpnglobal_vpnclientlessaccesspolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vpnglobal_vpnclientlessaccesspolicy_binding.Type(), &vpnglobal_vpnclientlessaccesspolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_vpnclientlessaccesspolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnglobal_vpnclientlessaccesspolicy_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	// Read the updated state back
	r.readVpnglobalVpnclientlessaccesspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalVpnclientlessaccesspolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalVpnclientlessaccesspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_vpnclientlessaccesspolicy_binding resource")

	r.readVpnglobalVpnclientlessaccesspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalVpnclientlessaccesspolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnglobalVpnclientlessaccesspolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnglobal_vpnclientlessaccesspolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		vpnglobal_vpnclientlessaccesspolicy_binding := vpnglobal_vpnclientlessaccesspolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Vpnglobal_vpnclientlessaccesspolicy_binding.Type(), &vpnglobal_vpnclientlessaccesspolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnglobal_vpnclientlessaccesspolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vpnglobal_vpnclientlessaccesspolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnglobal_vpnclientlessaccesspolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readVpnglobalVpnclientlessaccesspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalVpnclientlessaccesspolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalVpnclientlessaccesspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_vpnclientlessaccesspolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Single unique attribute - ID is the plain value
	// URL-encode the arg value so policy names with slashes/special chars survive (Pattern b).
	policyname_value := data.Id.ValueString()
	args := []string{
		fmt.Sprintf("policyname:%s", url.QueryEscape(policyname_value)),
	}

	err := r.client.DeleteResourceWithArgs(service.Vpnglobal_vpnclientlessaccesspolicy_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnglobal_vpnclientlessaccesspolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnglobal_vpnclientlessaccesspolicy_binding binding")
}

// Helper function to read vpnglobal_vpnclientlessaccesspolicy_binding data from API
func (r *VpnglobalVpnclientlessaccesspolicyBindingResource) readVpnglobalVpnclientlessaccesspolicyBindingFromApi(ctx context.Context, data *VpnglobalVpnclientlessaccesspolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 3: Array filter without parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"policyname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Vpnglobal_vpnclientlessaccesspolicy_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_vpnclientlessaccesspolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "vpnglobal_vpnclientlessaccesspolicy_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check policyname
		if idVal, ok := idMap["policyname"]; ok {
			if val, ok := v["policyname"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["policyname"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("vpnglobal_vpnclientlessaccesspolicy_binding not found with the provided ID attributes"))
		return
	}

	vpnglobal_vpnclientlessaccesspolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
