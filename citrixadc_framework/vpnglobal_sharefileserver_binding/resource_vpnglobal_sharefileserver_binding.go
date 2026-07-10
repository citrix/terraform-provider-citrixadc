package vpnglobal_sharefileserver_binding

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
var _ resource.Resource = &VpnglobalSharefileserverBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalSharefileserverBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalSharefileserverBindingResource)(nil)

func NewVpnglobalSharefileserverBindingResource() resource.Resource {
	return &VpnglobalSharefileserverBindingResource{}
}

// VpnglobalSharefileserverBindingResource defines the resource implementation.
type VpnglobalSharefileserverBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalSharefileserverBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalSharefileserverBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_sharefileserver_binding"
}

func (r *VpnglobalSharefileserverBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalSharefileserverBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalSharefileserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_sharefileserver_binding resource")
	vpnglobal_sharefileserver_binding := vpnglobal_sharefileserver_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vpnglobal_sharefileserver_binding.Type(), &vpnglobal_sharefileserver_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_sharefileserver_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnglobal_sharefileserver_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Sharefile.ValueString()))

	// Read the updated state back
	r.readVpnglobalSharefileserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalSharefileserverBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalSharefileserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_sharefileserver_binding resource")

	r.readVpnglobalSharefileserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalSharefileserverBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnglobalSharefileserverBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnglobal_sharefileserver_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		vpnglobal_sharefileserver_binding := vpnglobal_sharefileserver_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Vpnglobal_sharefileserver_binding.Type(), &vpnglobal_sharefileserver_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnglobal_sharefileserver_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vpnglobal_sharefileserver_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnglobal_sharefileserver_binding resource, skipping update")
	}

	// Read the updated state back
	r.readVpnglobalSharefileserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalSharefileserverBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalSharefileserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_sharefileserver_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Single unique attribute - ID is the plain value.
	// URL-encode the value: sharefile is "IP:PORT / FQDN:PORT", so it contains ':'
	// (and possibly '/') which collide with the NITRO "key:value,key:value" arg
	// syntax. Encoding only the value keeps the "sharefile:" key prefix intact.
	sharefile_value := data.Id.ValueString()
	args := []string{
		fmt.Sprintf("sharefile:%s", url.QueryEscape(sharefile_value)),
	}

	err := r.client.DeleteResourceWithArgs(service.Vpnglobal_sharefileserver_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnglobal_sharefileserver_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnglobal_sharefileserver_binding binding")
}

// Helper function to read vpnglobal_sharefileserver_binding data from API
func (r *VpnglobalSharefileserverBindingResource) readVpnglobalSharefileserverBindingFromApi(ctx context.Context, data *VpnglobalSharefileserverBindingResourceModel, diags *diag.Diagnostics) {

	// Pattern 10: single unique attribute => ID is the plain "sharefile" value.
	// The value itself can contain ':' (IP:PORT), which would confuse
	// ParseIdString's key:value detection, so use the ID directly as the filter
	// value instead of parsing it.
	sharefileId := data.Id.ValueString()

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Vpnglobal_sharefileserver_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_sharefileserver_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "vpnglobal_sharefileserver_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["sharefile"].(string); ok && val == sharefileId {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("vpnglobal_sharefileserver_binding not found with the provided ID attributes"))
		return
	}

	vpnglobal_sharefileserver_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
