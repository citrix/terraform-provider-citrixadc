package vpnvserver_vpnepaprofile_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VpnvserverVpnepaprofileBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverVpnepaprofileBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverVpnepaprofileBindingResource)(nil)

func NewVpnvserverVpnepaprofileBindingResource() resource.Resource {
	return &VpnvserverVpnepaprofileBindingResource{}
}

// VpnvserverVpnepaprofileBindingResource defines the resource implementation.
type VpnvserverVpnepaprofileBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverVpnepaprofileBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverVpnepaprofileBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_vpnepaprofile_binding"
}

func (r *VpnvserverVpnepaprofileBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverVpnepaprofileBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverVpnepaprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_vpnepaprofile_binding resource")
	// NOTE (firmware deprecation): current NetScaler firmware REJECTS
	// `bind vpn vserver -epaprofile` with the error "design change in the
	// support of OPSWAT specific EPA scans. EPA Profile Configuration is no
	// longer needed." This binding is therefore non-functional on current
	// firmware and live applies will fail at the NITRO call below. The resource
	// is retained for older firmware / state compatibility.
	vpnvserver_vpnepaprofile_binding := vpnvserver_vpnepaprofile_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vpnvserver_vpnepaprofile_binding.Type(), &vpnvserver_vpnepaprofile_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_vpnepaprofile_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnvserver_vpnepaprofile_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("epaprofile:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Epaprofile.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readVpnvserverVpnepaprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverVpnepaprofileBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverVpnepaprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_vpnepaprofile_binding resource")

	r.readVpnvserverVpnepaprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverVpnepaprofileBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnvserverVpnepaprofileBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for vpnvserver_vpnepaprofile_binding: NITRO exposes no
	// update endpoint for this binding (only add/delete/get) and all schema
	// attributes use RequiresReplace, so Terraform will never reach this with an
	// in-place change. We simply re-read and persist state.
	tflog.Debug(ctx, "Update is a no-op for vpnvserver_vpnepaprofile_binding; no NITRO update endpoint, all attributes are RequiresReplace")

	// Read the updated state back
	r.readVpnvserverVpnepaprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverVpnepaprofileBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverVpnepaprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_vpnepaprofile_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["epaprofile"]; ok && val != "" {
		argsMap["epaprofile"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Vpnvserver_vpnepaprofile_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnvserver_vpnepaprofile_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnvserver_vpnepaprofile_binding binding")
}

// Helper function to read vpnvserver_vpnepaprofile_binding data from API
func (r *VpnvserverVpnepaprofileBindingResource) readVpnvserverVpnepaprofileBindingFromApi(ctx context.Context, data *VpnvserverVpnepaprofileBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	name_Name, ok := idMap["name"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'name' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Vpnvserver_vpnepaprofile_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_vpnepaprofile_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "vpnvserver_vpnepaprofile_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check epaprofile
		if idVal, ok := idMap["epaprofile"]; ok {
			if val, ok := v["epaprofile"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["epaprofile"].(string); ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("vpnvserver_vpnepaprofile_binding not found with the provided ID attributes"))
		return
	}

	vpnvserver_vpnepaprofile_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
