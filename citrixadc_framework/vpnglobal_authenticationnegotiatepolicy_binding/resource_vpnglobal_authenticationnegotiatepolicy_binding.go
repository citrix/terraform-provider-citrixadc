package vpnglobal_authenticationnegotiatepolicy_binding

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
var _ resource.Resource = &VpnglobalAuthenticationnegotiatepolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalAuthenticationnegotiatepolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalAuthenticationnegotiatepolicyBindingResource)(nil)

func NewVpnglobalAuthenticationnegotiatepolicyBindingResource() resource.Resource {
	return &VpnglobalAuthenticationnegotiatepolicyBindingResource{}
}

// VpnglobalAuthenticationnegotiatepolicyBindingResource defines the resource implementation.
type VpnglobalAuthenticationnegotiatepolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalAuthenticationnegotiatepolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalAuthenticationnegotiatepolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_authenticationnegotiatepolicy_binding"
}

func (r *VpnglobalAuthenticationnegotiatepolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalAuthenticationnegotiatepolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalAuthenticationnegotiatepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_authenticationnegotiatepolicy_binding resource")
	vpnglobal_authenticationnegotiatepolicy_binding := vpnglobal_authenticationnegotiatepolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vpnglobal_authenticationnegotiatepolicy_binding.Type(), &vpnglobal_authenticationnegotiatepolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_authenticationnegotiatepolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnglobal_authenticationnegotiatepolicy_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	// Read the updated state back
	r.readVpnglobalAuthenticationnegotiatepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuthenticationnegotiatepolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalAuthenticationnegotiatepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_authenticationnegotiatepolicy_binding resource")

	r.readVpnglobalAuthenticationnegotiatepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuthenticationnegotiatepolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnglobalAuthenticationnegotiatepolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnglobal_authenticationnegotiatepolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		vpnglobal_authenticationnegotiatepolicy_binding := vpnglobal_authenticationnegotiatepolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Vpnglobal_authenticationnegotiatepolicy_binding.Type(), &vpnglobal_authenticationnegotiatepolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnglobal_authenticationnegotiatepolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vpnglobal_authenticationnegotiatepolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnglobal_authenticationnegotiatepolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readVpnglobalAuthenticationnegotiatepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuthenticationnegotiatepolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalAuthenticationnegotiatepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_authenticationnegotiatepolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Single unique attribute - ID is the plain value. URL-encode the value so a
	// policyname containing slashes/special characters is passed safely (Pattern b).
	policyname_value := data.Id.ValueString()
	args := []string{
		fmt.Sprintf("policyname:%s", url.QueryEscape(policyname_value)),
	}

	err := r.client.DeleteResourceWithArgs(service.Vpnglobal_authenticationnegotiatepolicy_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnglobal_authenticationnegotiatepolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnglobal_authenticationnegotiatepolicy_binding binding")
}

// Helper function to read vpnglobal_authenticationnegotiatepolicy_binding data from API
func (r *VpnglobalAuthenticationnegotiatepolicyBindingResource) readVpnglobalAuthenticationnegotiatepolicyBindingFromApi(ctx context.Context, data *VpnglobalAuthenticationnegotiatepolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Single unique attribute - the ID is the plain policyname value (Pattern 10:
	// do not use ParseIdString on a plain-value ID; it returns an empty map).
	// ParseIdString still resolves a legacy SDK v2 import ID (plain "policyname") correctly.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"policyname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}
	policynameValue := data.Id.ValueString()
	if v, ok := idMap["policyname"]; ok && v != "" {
		policynameValue = v
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Vpnglobal_authenticationnegotiatepolicy_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_authenticationnegotiatepolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "vpnglobal_authenticationnegotiatepolicy_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right policyname
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["policyname"].(string); ok && val == policynameValue {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("vpnglobal_authenticationnegotiatepolicy_binding not found with the provided ID attributes"))
		return
	}

	vpnglobal_authenticationnegotiatepolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
