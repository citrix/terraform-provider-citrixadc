package vpnglobal_secureprivateaccessurl_binding

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
var _ resource.Resource = &VpnglobalSecureprivateaccessurlBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalSecureprivateaccessurlBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalSecureprivateaccessurlBindingResource)(nil)

func NewVpnglobalSecureprivateaccessurlBindingResource() resource.Resource {
	return &VpnglobalSecureprivateaccessurlBindingResource{}
}

// VpnglobalSecureprivateaccessurlBindingResource defines the resource implementation.
type VpnglobalSecureprivateaccessurlBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalSecureprivateaccessurlBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalSecureprivateaccessurlBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_secureprivateaccessurl_binding"
}

func (r *VpnglobalSecureprivateaccessurlBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalSecureprivateaccessurlBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalSecureprivateaccessurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_secureprivateaccessurl_binding resource")
	vpnglobal_secureprivateaccessurl_binding := vpnglobal_secureprivateaccessurl_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vpnglobal_secureprivateaccessurl_binding.Type(), &vpnglobal_secureprivateaccessurl_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_secureprivateaccessurl_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnglobal_secureprivateaccessurl_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Secureprivateaccessurl.ValueString()))

	// Read the updated state back
	r.readVpnglobalSecureprivateaccessurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalSecureprivateaccessurlBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalSecureprivateaccessurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_secureprivateaccessurl_binding resource")

	r.readVpnglobalSecureprivateaccessurlBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnglobalSecureprivateaccessurlBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnglobalSecureprivateaccessurlBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for vpnglobal_secureprivateaccessurl_binding; NITRO exposes no
	// update endpoint and all attributes are RequiresReplace.
	tflog.Debug(ctx, "Update is a no-op for vpnglobal_secureprivateaccessurl_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readVpnglobalSecureprivateaccessurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalSecureprivateaccessurlBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalSecureprivateaccessurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_secureprivateaccessurl_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Single unique attribute - ID is the plain value
	secureprivateaccessurl_value := data.Id.ValueString()
	// The value is a URL containing reserved characters (':' and '/').
	// nitro-go appends the arg value to the query string without encoding,
	// so encode the value (not the "secureprivateaccessurl:" separator) to avoid a 400 from NITRO.
	args := []string{
		fmt.Sprintf("secureprivateaccessurl:%s", utils.UrlEncode(secureprivateaccessurl_value)),
	}

	err := r.client.DeleteResourceWithArgs(service.Vpnglobal_secureprivateaccessurl_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnglobal_secureprivateaccessurl_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnglobal_secureprivateaccessurl_binding binding")
}

// Helper function to read vpnglobal_secureprivateaccessurl_binding data from API
func (r *VpnglobalSecureprivateaccessurlBindingResource) readVpnglobalSecureprivateaccessurlBindingFromApi(ctx context.Context, data *VpnglobalSecureprivateaccessurlBindingResourceModel, diags *diag.Diagnostics) {

	// Single unique attribute - ID is the plain secureprivateaccessurl value
	secureprivateaccessurl_value := data.Id.ValueString()

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Vpnglobal_secureprivateaccessurl_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_secureprivateaccessurl_binding, got error: %s", err))
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
		// Check secureprivateaccessurl
		if val, ok := v["secureprivateaccessurl"].(string); ok {
			if val == secureprivateaccessurl_value {
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

	vpnglobal_secureprivateaccessurl_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
