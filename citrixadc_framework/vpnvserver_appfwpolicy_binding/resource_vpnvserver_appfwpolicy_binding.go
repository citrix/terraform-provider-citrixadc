package vpnvserver_appfwpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VpnvserverAppfwpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverAppfwpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverAppfwpolicyBindingResource)(nil)

func NewVpnvserverAppfwpolicyBindingResource() resource.Resource {
	return &VpnvserverAppfwpolicyBindingResource{}
}

// VpnvserverAppfwpolicyBindingResource defines the resource implementation.
type VpnvserverAppfwpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverAppfwpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverAppfwpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_appfwpolicy_binding"
}

func (r *VpnvserverAppfwpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverAppfwpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverAppfwpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_appfwpolicy_binding resource")

	vpnvserverAppfwpolicyBinding := vpnvserverAppfwpolicyBindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	err := r.client.UpdateUnnamedResource("vpnvserver_appfwpolicy_binding", &vpnvserverAppfwpolicyBinding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_appfwpolicy_binding, got error: %s", err))
		return
	}

	// Generate ID from name and policy
	bindingId := fmt.Sprintf("%s,%s", data.Name.ValueString(), data.Policy.ValueString())
	data.Id = types.StringValue(bindingId)

	tflog.Trace(ctx, "Created vpnvserver_appfwpolicy_binding resource")

	// Read the updated state back
	r.readVpnvserverAppfwpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAppfwpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverAppfwpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_appfwpolicy_binding resource")

	r.readVpnvserverAppfwpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAppfwpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnvserverAppfwpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnvserver_appfwpolicy_binding resource")

	// For binding resources, updates typically require delete and recreate
	// This should not be called as all fields are ForceNew
	resp.Diagnostics.AddError("Update Not Supported", "vpnvserver_appfwpolicy_binding does not support updates. All fields are ForceNew.")
}

func (r *VpnvserverAppfwpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverAppfwpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_appfwpolicy_binding resource")

	bindingId := data.Id.ValueString()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	policy := idSlice[1]

	// Build args for delete
	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policy:%s", policy))

	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		args = append(args, fmt.Sprintf("secondary:%t", data.Secondary.ValueBool()))
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		args = append(args, fmt.Sprintf("groupextraction:%t", data.Groupextraction.ValueBool()))
	}

	err := r.client.DeleteResourceWithArgs("vpnvserver_appfwpolicy_binding", name, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnvserver_appfwpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnvserver_appfwpolicy_binding resource")
}

// Helper function to read vpnvserver_appfwpolicy_binding data from API
func (r *VpnvserverAppfwpolicyBindingResource) readVpnvserverAppfwpolicyBindingFromApi(ctx context.Context, data *VpnvserverAppfwpolicyBindingResourceModel, diags *diag.Diagnostics) {
	bindingId := data.Id.ValueString()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	policy := idSlice[1]

	findParams := service.FindParams{
		ResourceType:             "vpnvserver_appfwpolicy_binding",
		ResourceName:             name,
		ResourceMissingErrorCode: 258,
	}

	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_appfwpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		tflog.Warn(ctx, fmt.Sprintf("Clearing vpnvserver_appfwpolicy_binding state %s - not found", bindingId))
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right policy
	foundIndex := -1
	for i, v := range dataArr {
		if v["policy"].(string) == policy {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		tflog.Warn(ctx, fmt.Sprintf("Clearing vpnvserver_appfwpolicy_binding state %s - policy not found", bindingId))
		data.Id = types.StringNull()
		return
	}

	getResponseData := dataArr[foundIndex]

	vpnvserverAppfwpolicyBindingSetAttrFromGet(ctx, data, getResponseData)
}
