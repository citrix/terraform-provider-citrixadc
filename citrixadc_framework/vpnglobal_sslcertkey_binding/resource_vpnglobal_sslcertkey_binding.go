package vpnglobal_sslcertkey_binding

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
var _ resource.Resource = &VpnglobalSslcertkeyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalSslcertkeyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalSslcertkeyBindingResource)(nil)

func NewVpnglobalSslcertkeyBindingResource() resource.Resource {
	return &VpnglobalSslcertkeyBindingResource{}
}

// VpnglobalSslcertkeyBindingResource defines the resource implementation.
type VpnglobalSslcertkeyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalSslcertkeyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalSslcertkeyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_sslcertkey_binding"
}

func (r *VpnglobalSslcertkeyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalSslcertkeyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalSslcertkeyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_sslcertkey_binding resource")
	vpnglobal_sslcertkey_binding := vpnglobal_sslcertkey_bindingGetThePayloadFromthePlan(ctx, &data)

	// Binding resource - SDK v2 used UpdateUnnamedResource (PUT). Keep that verb.
	err := r.client.UpdateUnnamedResource(service.Vpnglobal_sslcertkey_binding.Type(), &vpnglobal_sslcertkey_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_sslcertkey_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnglobal_sslcertkey_binding resource")

	// SDK v2 contract: ID is the plain certkeyname value (d.SetId(certkeyname)).
	data.Id = types.StringValue(data.Certkeyname.ValueString())

	// Read the updated state back
	r.readVpnglobalSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalSslcertkeyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalSslcertkeyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_sslcertkey_binding resource")

	r.readVpnglobalSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalSslcertkeyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// All attributes are RequiresReplace; this binding has no NITRO update endpoint.
	// Update is a documented no-op (Pattern 5).
	var data, state VpnglobalSslcertkeyBindingResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for vpnglobal_sslcertkey_binding; all attributes are RequiresReplace")

	r.readVpnglobalSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalSslcertkeyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalSslcertkeyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_sslcertkey_binding resource")

	// Global binding - delete with args (empty resource name). The ID is the plain
	// certkeyname. Mirror the SDK v2 delete args: always certkeyname, plus
	// userdataencryptionkey / cacert when set. URL-encode the values because the NITRO
	// client does NOT encode arg values (only the resourceName).
	certkeyname := data.Id.ValueString()
	args := []string{"certkeyname:" + url.QueryEscape(certkeyname)}
	if !data.Userdataencryptionkey.IsNull() && data.Userdataencryptionkey.ValueString() != "" {
		args = append(args, "userdataencryptionkey:"+url.QueryEscape(data.Userdataencryptionkey.ValueString()))
	}
	if !data.Cacert.IsNull() && data.Cacert.ValueString() != "" {
		args = append(args, "cacert:"+url.QueryEscape(data.Cacert.ValueString()))
	}

	err := r.client.DeleteResourceWithArgs(service.Vpnglobal_sslcertkey_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnglobal_sslcertkey_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnglobal_sslcertkey_binding binding")
}

// Helper function to read vpnglobal_sslcertkey_binding data from API
func (r *VpnglobalSslcertkeyBindingResource) readVpnglobalSslcertkeyBindingFromApi(ctx context.Context, data *VpnglobalSslcertkeyBindingResourceModel, diags *diag.Diagnostics) {

	// Single-key resource: the ID is the plain certkeyname value (Pattern 10).
	certkeyname := data.Id.ValueString()

	findParams := service.FindParams{
		ResourceType:             service.Vpnglobal_sslcertkey_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_sslcertkey_binding, got error: %s", err))
		return
	}

	// Iterate through results to find the one with the matching certkeyname
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["certkeyname"].(string); ok && val == certkeyname {
			foundIndex = i
			break
		}
	}

	// Resource is missing - clear ID so Terraform plans a recreate
	if foundIndex == -1 {
		tflog.Warn(ctx, fmt.Sprintf("vpnglobal_sslcertkey_binding %s not found, clearing state", certkeyname))
		data.Id = types.StringNull()
		return
	}

	vpnglobal_sslcertkey_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
