package vpnglobal_appcontroller_binding

import (
	"context"
	"fmt"
	"net/url"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VpnglobalAppcontrollerBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalAppcontrollerBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalAppcontrollerBindingResource)(nil)

func NewVpnglobalAppcontrollerBindingResource() resource.Resource {
	return &VpnglobalAppcontrollerBindingResource{}
}

// VpnglobalAppcontrollerBindingResource defines the resource implementation.
type VpnglobalAppcontrollerBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalAppcontrollerBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalAppcontrollerBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_appcontroller_binding"
}

func (r *VpnglobalAppcontrollerBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalAppcontrollerBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalAppcontrollerBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_appcontroller_binding resource")
	vpnglobal_appcontroller_binding := vpnglobal_appcontroller_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - NITRO 'add' verb is POST (matches SDK v2 AddResource)
	_, err := r.client.AddResource(service.Vpnglobal_appcontroller_binding.Type(), "", &vpnglobal_appcontroller_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_appcontroller_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnglobal_appcontroller_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Appcontroller.ValueString()))

	// Read the updated state back
	r.readVpnglobalAppcontrollerBindingFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "vpnglobal_appcontroller_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAppcontrollerBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalAppcontrollerBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_appcontroller_binding resource")

	r.readVpnglobalAppcontrollerBindingFromApi(ctx, &data, &resp.Diagnostics)

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

func (r *VpnglobalAppcontrollerBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnglobalAppcontrollerBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnglobal_appcontroller_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		vpnglobal_appcontroller_binding := vpnglobal_appcontroller_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Vpnglobal_appcontroller_binding.Type(), &vpnglobal_appcontroller_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnglobal_appcontroller_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vpnglobal_appcontroller_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnglobal_appcontroller_binding resource, skipping update")
	}

	// Read the updated state back
	r.readVpnglobalAppcontrollerBindingFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "vpnglobal_appcontroller_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAppcontrollerBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalAppcontrollerBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_appcontroller_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Single unique attribute - ID is the plain value.
	// URL-encode the value: appcontroller is a URL with special chars (':' '/'),
	// which the NITRO delete args= query string rejects unless escaped (matches
	// SDK v2 url.QueryEscape). The NITRO client joins args verbatim into the URL.
	appcontroller_value := data.Id.ValueString()
	args := []string{
		fmt.Sprintf("appcontroller:%s", url.QueryEscape(appcontroller_value)),
	}

	err := r.client.DeleteResourceWithArgs(service.Vpnglobal_appcontroller_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnglobal_appcontroller_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnglobal_appcontroller_binding binding")
}

// Helper function to read vpnglobal_appcontroller_binding data from API
func (r *VpnglobalAppcontrollerBindingResource) readVpnglobalAppcontrollerBindingFromApi(ctx context.Context, data *VpnglobalAppcontrollerBindingResourceModel, diags *diag.Diagnostics) {

	// Single unique attribute - the ID is the plain appcontroller value.
	// Do NOT use ParseIdString here: appcontroller values are URLs (e.g.
	// "http://www.citrix.com") whose embedded colons make ParseIdString
	// mis-detect the new key:value format (Pattern 10).
	appcontrollerId := data.Id.ValueString()

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Vpnglobal_appcontroller_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_appcontroller_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
		// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["appcontroller"].(string); ok && val == appcontrollerId {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	vpnglobal_appcontroller_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
