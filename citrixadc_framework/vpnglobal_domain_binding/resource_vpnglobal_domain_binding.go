package vpnglobal_domain_binding

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
var _ resource.Resource = &VpnglobalDomainBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalDomainBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalDomainBindingResource)(nil)

func NewVpnglobalDomainBindingResource() resource.Resource {
	return &VpnglobalDomainBindingResource{}
}

// VpnglobalDomainBindingResource defines the resource implementation.
type VpnglobalDomainBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalDomainBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalDomainBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_domain_binding"
}

func (r *VpnglobalDomainBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalDomainBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalDomainBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_domain_binding resource")
	vpnglobal_domain_binding := vpnglobal_domain_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vpnglobal_domain_binding.Type(), &vpnglobal_domain_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_domain_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnglobal_domain_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Intranetdomain.ValueString()))

	// Read the updated state back
	r.readVpnglobalDomainBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "vpnglobal_domain_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalDomainBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalDomainBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_domain_binding resource")

	r.readVpnglobalDomainBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnglobalDomainBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnglobalDomainBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnglobal_domain_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		vpnglobal_domain_binding := vpnglobal_domain_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Vpnglobal_domain_binding.Type(), &vpnglobal_domain_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnglobal_domain_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vpnglobal_domain_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnglobal_domain_binding resource, skipping update")
	}

	// Read the updated state back
	r.readVpnglobalDomainBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "vpnglobal_domain_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalDomainBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalDomainBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_domain_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Single unique attribute - ID is the plain value.
	// URL-encode the value: intranetdomain can be a URL with '/' and ':'
	// (matches the SDK v2 resource which used url.QueryEscape).
	intranetdomain_value := data.Id.ValueString()
	args := []string{
		fmt.Sprintf("intranetdomain:%s", url.QueryEscape(intranetdomain_value)),
	}

	err := r.client.DeleteResourceWithArgs(service.Vpnglobal_domain_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnglobal_domain_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnglobal_domain_binding binding")
}

// Helper function to read vpnglobal_domain_binding data from API
func (r *VpnglobalDomainBindingResource) readVpnglobalDomainBindingFromApi(ctx context.Context, data *VpnglobalDomainBindingResourceModel, diags *diag.Diagnostics) {

	// Single unique attribute - the ID is the plain intranetdomain value.
	// intranetdomain values can contain ':' (e.g. URLs like
	// "http://www.example.com/"), so do NOT run it through ParseIdString
	// (Pattern 10) - that would misparse "http:" as a key. Use the ID directly.
	intranetdomain_value := data.Id.ValueString()

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Vpnglobal_domain_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_domain_binding, got error: %s", err))
		return
	}

	// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
	// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["intranetdomain"].(string); ok && val == intranetdomain_value {
			foundIndex = i
			break
		}
	}

	// Binding not present in the returned set: signal removal via a null Id (see above).
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	vpnglobal_domain_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
