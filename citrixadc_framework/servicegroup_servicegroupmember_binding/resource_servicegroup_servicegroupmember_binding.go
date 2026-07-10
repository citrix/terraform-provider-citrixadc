package servicegroup_servicegroupmember_binding

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
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
var _ resource.Resource = &ServicegroupServicegroupmemberBindingResource{}
var _ resource.ResourceWithConfigure = (*ServicegroupServicegroupmemberBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ServicegroupServicegroupmemberBindingResource)(nil)

func NewServicegroupServicegroupmemberBindingResource() resource.Resource {
	return &ServicegroupServicegroupmemberBindingResource{}
}

// ServicegroupServicegroupmemberBindingResource defines the resource implementation.
type ServicegroupServicegroupmemberBindingResource struct {
	client *service.NitroClient
}

func (r *ServicegroupServicegroupmemberBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ServicegroupServicegroupmemberBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_servicegroup_servicegroupmember_binding"
}

func (r *ServicegroupServicegroupmemberBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ServicegroupServicegroupmemberBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ServicegroupServicegroupmemberBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating servicegroup_servicegroupmember_binding resource")
	servicegroup_servicegroupmember_binding := servicegroup_servicegroupmember_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Servicegroup_servicegroupmember_binding.Type(), &servicegroup_servicegroupmember_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create servicegroup_servicegroupmember_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created servicegroup_servicegroupmember_binding resource")

	// Set ID for the resource before reading state.
	// Mirrors the SDK v2 ID semantics: the member is identified by
	// servicegroupname + (servername OR ip) + optional port. When the user binds
	// by ip, the ADC creates a server named after the ip, so servername always
	// resolves to a valid search key — use ip as the effective servername here.
	effectiveServername := data.Servername.ValueString()
	if effectiveServername == "" {
		effectiveServername = data.Ip.ValueString()
	}
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(data.Servicegroupname.ValueString())))
	idParts = append(idParts, fmt.Sprintf("servername:%s", utils.UrlEncode(effectiveServername)))
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		idParts = append(idParts, fmt.Sprintf("port:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Port.ValueInt64()))))
	}
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back, unless disable_read is set (mirrors SDK v2,
	// which skips the read when disable_read is true). When skipping, resolve any
	// unknown Computed attributes to null so all values are known after apply.
	if data.DisableRead.ValueBool() {
		servicegroup_servicegroupmember_bindingResolveUnknownComputed(&data)
	} else {
		r.readServicegroupServicegroupmemberBindingFromApi(ctx, &data, &resp.Diagnostics)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServicegroupServicegroupmemberBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ServicegroupServicegroupmemberBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading servicegroup_servicegroupmember_binding resource")

	// disable_read (synthetic SDK v2 flag): when true, skip querying the ADC and
	// preserve the prior state unchanged (mirrors SDK v2's early return in Read).
	if data.DisableRead.ValueBool() {
		tflog.Debug(ctx, "disable_read is true; skipping API read and preserving prior state")
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}

	r.readServicegroupServicegroupmemberBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServicegroupServicegroupmemberBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ServicegroupServicegroupmemberBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating servicegroup_servicegroupmember_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		servicegroup_servicegroupmember_binding := servicegroup_servicegroupmember_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Servicegroup_servicegroupmember_binding.Type(), &servicegroup_servicegroupmember_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update servicegroup_servicegroupmember_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated servicegroup_servicegroupmember_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for servicegroup_servicegroupmember_binding resource, skipping update")
	}

	// Read the updated state back, unless disable_read is set (mirrors SDK v2,
	// which skips the read when disable_read is true). When skipping, resolve any
	// unknown Computed attributes to null so all values are known after apply.
	if data.DisableRead.ValueBool() {
		servicegroup_servicegroupmember_bindingResolveUnknownComputed(&data)
	} else {
		r.readServicegroupServicegroupmemberBindingFromApi(ctx, &data, &resp.Diagnostics)
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServicegroupServicegroupmemberBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ServicegroupServicegroupmemberBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting servicegroup_servicegroupmember_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// Mirrors SDK v2: delete by servername (== ip for ip-based members) and,
	// when present, port. ParseIdString handles new and legacy ID forms.
	idMap, optionalAbsent, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicegroupname", "servername", "port"}, []string{"servername", "port"})
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	servicegroupname_value, ok := idMap["servicegroupname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'servicegroupname' not found in ID")
		return
	}

	args := make([]string, 0, 2)
	if val, ok := idMap["servername"]; ok && val != "" {
		// URL-encode the value so slashy/special chars (e.g. IPv6) survive the
		// query string.
		args = append(args, fmt.Sprintf("servername:%s", url.QueryEscape(val)))
	}
	if val, ok := idMap["port"]; ok && val != "" && !optionalAbsent["port"] {
		args = append(args, fmt.Sprintf("port:%s", url.QueryEscape(val)))
	}

	err = r.client.DeleteResourceWithArgs(service.Servicegroup_servicegroupmember_binding.Type(), servicegroupname_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete servicegroup_servicegroupmember_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted servicegroup_servicegroupmember_binding binding")
}

// Helper function to read servicegroup_servicegroupmember_binding data from API
func (r *ServicegroupServicegroupmemberBindingResource) readServicegroupServicegroupmemberBindingFromApi(ctx context.Context, data *ServicegroupServicegroupmemberBindingResourceModel, diags *diag.Diagnostics) {

	// Binding with parent ID - parse from ID. Mirrors the SDK v2 identity:
	// servicegroupname + (servername OR ip, stored under "servername") + optional
	// port. ParseIdString handles both the new key:value form and the legacy
	// positional "servicegroupname,servername,port" form.
	idMap, optionalAbsent, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicegroupname", "servername", "port"}, []string{"servername", "port"})
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	servicegroupname_Name, ok := idMap["servicegroupname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'servicegroupname' not found in ID string")
		return
	}

	// The ADC creates a server named after the bound ip, so "servername" in the ID
	// is the effective member key whether the user bound by servername or by ip.
	idServername := idMap["servername"]

	idPort := 0
	if portStr, ok := idMap["port"]; ok && portStr != "" && !optionalAbsent["port"] {
		if idPort, err = strconv.Atoi(portStr); err != nil {
			diags.AddError("Parse Error", fmt.Sprintf("Unable to parse port from ID: %s", err))
			return
		}
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Servicegroup_servicegroupmember_binding.Type(),
		ResourceName:             servicegroupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read servicegroup_servicegroupmember_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "servicegroup_servicegroupmember_binding returned empty array.")
		return
	}

	// Iterate through results to find the matching member, matching on
	// servername (== ip for ip-based members) and, when supplied, port.
	foundIndex := -1
	for i, v := range dataArr {
		servernameVal, _ := v["servername"].(string)
		if servernameVal != idServername {
			continue
		}
		if idPort != 0 {
			if pv, ok := v["port"]; ok {
				portVal, _ := utils.ConvertToInt64(pv)
				if portVal != int64(idPort) {
					continue
				}
			} else {
				continue
			}
		}
		foundIndex = i
		break
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", "servicegroup_servicegroupmember_binding not found with the provided ID attributes")
		return
	}

	servicegroup_servicegroupmember_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}

// servicegroup_servicegroupmember_bindingResolveUnknownComputed sets any unknown
// Computed attribute to null. Used on the disable_read path where we skip the API
// read but Terraform still requires every value to be known after apply.
func servicegroup_servicegroupmember_bindingResolveUnknownComputed(data *ServicegroupServicegroupmemberBindingResourceModel) {
	if data.Customserverid.IsUnknown() {
		data.Customserverid = types.StringNull()
	}
	if data.Dbsttl.IsUnknown() {
		data.Dbsttl = types.Int64Null()
	}
	if data.Hashid.IsUnknown() {
		data.Hashid = types.Int64Null()
	}
	if data.Ip.IsUnknown() {
		data.Ip = types.StringNull()
	}
	if data.Nameserver.IsUnknown() {
		data.Nameserver = types.StringNull()
	}
	if data.Order.IsUnknown() {
		data.Order = types.Int64Null()
	}
	if data.Port.IsUnknown() {
		data.Port = types.Int64Null()
	}
	if data.Serverid.IsUnknown() {
		data.Serverid = types.Int64Null()
	}
	if data.Servername.IsUnknown() {
		data.Servername = types.StringNull()
	}
	if data.State.IsUnknown() {
		data.State = types.StringNull()
	}
	if data.Weight.IsUnknown() {
		data.Weight = types.Int64Null()
	}
}
