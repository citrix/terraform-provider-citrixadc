package vpnvserver_secureprivateaccessurl_binding

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
var _ resource.Resource = &VpnvserverSecureprivateaccessurlBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverSecureprivateaccessurlBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverSecureprivateaccessurlBindingResource)(nil)

func NewVpnvserverSecureprivateaccessurlBindingResource() resource.Resource {
	return &VpnvserverSecureprivateaccessurlBindingResource{}
}

// VpnvserverSecureprivateaccessurlBindingResource defines the resource implementation.
type VpnvserverSecureprivateaccessurlBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverSecureprivateaccessurlBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverSecureprivateaccessurlBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_secureprivateaccessurl_binding"
}

func (r *VpnvserverSecureprivateaccessurlBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverSecureprivateaccessurlBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverSecureprivateaccessurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_secureprivateaccessurl_binding resource")
	vpnvserver_secureprivateaccessurl_binding := vpnvserver_secureprivateaccessurl_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vpnvserver_secureprivateaccessurl_binding.Type(), &vpnvserver_secureprivateaccessurl_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_secureprivateaccessurl_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnvserver_secureprivateaccessurl_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("secureprivateaccessurl:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Secureprivateaccessurl.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readVpnvserverSecureprivateaccessurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverSecureprivateaccessurlBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverSecureprivateaccessurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_secureprivateaccessurl_binding resource")

	r.readVpnvserverSecureprivateaccessurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverSecureprivateaccessurlBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnvserverSecureprivateaccessurlBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// No NITRO update endpoint exists for this binding; both attributes are
	// RequiresReplace, so Terraform never invokes Update with changed values.
	tflog.Debug(ctx, "Update is a no-op for vpnvserver_secureprivateaccessurl_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readVpnvserverSecureprivateaccessurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverSecureprivateaccessurlBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverSecureprivateaccessurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_secureprivateaccessurl_binding resource")
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
	if val, ok := idMap["secureprivateaccessurl"]; ok && val != "" {
		// The value is a URL containing reserved characters (':' and '/').
		// nitro-go appends the arg value to the query string without encoding,
		// so encode the value (not the "key:" separator) to avoid a 400 from NITRO.
		argsMap["secureprivateaccessurl"] = utils.UrlEncode(val)
	}

	err = r.client.DeleteResourceWithArgsMap(service.Vpnvserver_secureprivateaccessurl_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnvserver_secureprivateaccessurl_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnvserver_secureprivateaccessurl_binding binding")
}

// Helper function to read vpnvserver_secureprivateaccessurl_binding data from API
func (r *VpnvserverSecureprivateaccessurlBindingResource) readVpnvserverSecureprivateaccessurlBindingFromApi(ctx context.Context, data *VpnvserverSecureprivateaccessurlBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Vpnvserver_secureprivateaccessurl_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_secureprivateaccessurl_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "vpnvserver_secureprivateaccessurl_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check secureprivateaccessurl
		if idVal, ok := idMap["secureprivateaccessurl"]; ok {
			if val, ok := v["secureprivateaccessurl"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["secureprivateaccessurl"].(string); ok {
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
		diags.AddError("Client Error", fmt.Sprintf("vpnvserver_secureprivateaccessurl_binding not found with the provided ID attributes"))
		return
	}

	vpnvserver_secureprivateaccessurl_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
