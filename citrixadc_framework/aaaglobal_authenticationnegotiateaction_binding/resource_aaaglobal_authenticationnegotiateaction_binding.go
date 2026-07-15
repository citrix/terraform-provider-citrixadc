package aaaglobal_authenticationnegotiateaction_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AaaglobalAuthenticationnegotiateactionBindingResource{}
var _ resource.ResourceWithConfigure = (*AaaglobalAuthenticationnegotiateactionBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AaaglobalAuthenticationnegotiateactionBindingResource)(nil)

func NewAaaglobalAuthenticationnegotiateactionBindingResource() resource.Resource {
	return &AaaglobalAuthenticationnegotiateactionBindingResource{}
}

// AaaglobalAuthenticationnegotiateactionBindingResource defines the resource implementation.
type AaaglobalAuthenticationnegotiateactionBindingResource struct {
	client *service.NitroClient
}

func (r *AaaglobalAuthenticationnegotiateactionBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaaglobalAuthenticationnegotiateactionBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaaglobal_authenticationnegotiateaction_binding"
}

func (r *AaaglobalAuthenticationnegotiateactionBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaaglobalAuthenticationnegotiateactionBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaaglobalAuthenticationnegotiateactionBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaaglobal_authenticationnegotiateaction_binding resource")
	aaaglobal_authenticationnegotiateaction_binding := aaaglobal_authenticationnegotiateaction_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Aaaglobal_authenticationnegotiateaction_binding.Type(), &aaaglobal_authenticationnegotiateaction_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaaglobal_authenticationnegotiateaction_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created aaaglobal_authenticationnegotiateaction_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Windowsprofile.ValueString()))

	// Read the updated state back
	r.readAaaglobalAuthenticationnegotiateactionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaaglobalAuthenticationnegotiateactionBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaaglobalAuthenticationnegotiateactionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaaglobal_authenticationnegotiateaction_binding resource")

	r.readAaaglobalAuthenticationnegotiateactionBindingFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Resource was deleted out-of-band - remove from state for self-healing
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaaglobalAuthenticationnegotiateactionBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AaaglobalAuthenticationnegotiateactionBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// No-op: aaaglobal_authenticationnegotiateaction_binding has no NITRO update endpoint
	// (only add/delete/get). The sole writable attribute windowsprofile is RequiresReplace,
	// so Terraform never calls Update for an actual change.
	tflog.Debug(ctx, "Update is a no-op for aaaglobal_authenticationnegotiateaction_binding; windowsprofile is RequiresReplace and there is no update endpoint")

	// Read the updated state back
	r.readAaaglobalAuthenticationnegotiateactionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaaglobalAuthenticationnegotiateactionBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaaglobalAuthenticationnegotiateactionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaaglobal_authenticationnegotiateaction_binding resource")
	// Keyless global binding - delete using DeleteResourceWithArgs with empty resource name
	// Single unique attribute - ID is the plain value; UrlEncode the value defensively
	windowsprofile_value := data.Id.ValueString()
	args := []string{
		fmt.Sprintf("windowsprofile:%s", utils.UrlEncode(windowsprofile_value)),
	}

	err := r.client.DeleteResourceWithArgs(service.Aaaglobal_authenticationnegotiateaction_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete aaaglobal_authenticationnegotiateaction_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted aaaglobal_authenticationnegotiateaction_binding binding")
}

// Helper function to read aaaglobal_authenticationnegotiateaction_binding data from API
func (r *AaaglobalAuthenticationnegotiateactionBindingResource) readAaaglobalAuthenticationnegotiateactionBindingFromApi(ctx context.Context, data *AaaglobalAuthenticationnegotiateactionBindingResourceModel, diags *diag.Diagnostics) {

	// Keyless global binding - single unique attribute, ID is the plain windowsprofile value
	windowsprofile_value := data.Id.ValueString()

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Aaaglobal_authenticationnegotiateaction_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaaglobal_authenticationnegotiateaction_binding, got error: %s", err))
		return
	}

	// Resource is missing - signal deletion for self-healing
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right windowsprofile
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["windowsprofile"].(string); ok && val == windowsprofile_value {
			foundIndex = i
			break
		}
	}

	// Resource is missing - signal deletion for self-healing
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	aaaglobal_authenticationnegotiateaction_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
