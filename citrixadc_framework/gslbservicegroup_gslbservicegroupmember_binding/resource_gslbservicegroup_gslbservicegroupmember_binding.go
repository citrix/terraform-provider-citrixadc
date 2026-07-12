package gslbservicegroup_gslbservicegroupmember_binding

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
var _ resource.Resource = &GslbservicegroupGslbservicegroupmemberBindingResource{}
var _ resource.ResourceWithConfigure = (*GslbservicegroupGslbservicegroupmemberBindingResource)(nil)
var _ resource.ResourceWithImportState = (*GslbservicegroupGslbservicegroupmemberBindingResource)(nil)

func NewGslbservicegroupGslbservicegroupmemberBindingResource() resource.Resource {
	return &GslbservicegroupGslbservicegroupmemberBindingResource{}
}

// GslbservicegroupGslbservicegroupmemberBindingResource defines the resource implementation.
type GslbservicegroupGslbservicegroupmemberBindingResource struct {
	client *service.NitroClient
}

func (r *GslbservicegroupGslbservicegroupmemberBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbservicegroupGslbservicegroupmemberBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbservicegroup_gslbservicegroupmember_binding"
}

func (r *GslbservicegroupGslbservicegroupmemberBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbservicegroupGslbservicegroupmemberBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbservicegroupGslbservicegroupmemberBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbservicegroup_gslbservicegroupmember_binding resource")
	gslbservicegroup_gslbservicegroupmember_binding := gslbservicegroup_gslbservicegroupmember_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Build the NITRO binding id exactly as the SDK v2 resource does:
	// servicegroupname + (servername OR ip) + optional port, joined by commas.
	// When the user binds by ip, the ADC creates a server named after the ip, so
	// servername is always a valid search key.
	bindingIdSlice := make([]string, 0, 3)
	bindingIdSlice = append(bindingIdSlice, data.Servicegroupname.ValueString())
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() && data.Servername.ValueString() != "" {
		bindingIdSlice = append(bindingIdSlice, data.Servername.ValueString())
	} else if !data.Ip.IsNull() && !data.Ip.IsUnknown() && data.Ip.ValueString() != "" {
		bindingIdSlice = append(bindingIdSlice, data.Ip.ValueString())
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		bindingIdSlice = append(bindingIdSlice, strconv.Itoa(int(data.Port.ValueInt64())))
	}
	bindingId := strings.Join(bindingIdSlice, ",")

	// Make API call
	// Binding resource with parent - use AddResource with the binding id (SDK v2 parity).
	_, err := r.client.AddResource(service.Gslbservicegroup_gslbservicegroupmember_binding.Type(), bindingId, &gslbservicegroup_gslbservicegroupmember_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create gslbservicegroup_gslbservicegroupmember_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created gslbservicegroup_gslbservicegroupmember_binding resource")

	// Set ID for the resource before reading state.
	// Mirrors the SDK v2 ID semantics: the member is identified by
	// servicegroupname + (servername OR ip) + optional port. When the user binds
	// by ip, the ADC creates a server named after the ip, so servername always
	// resolves to a valid search key -- use ip as the effective servername here.
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

	// Read the updated state back.
	r.readGslbservicegroupGslbservicegroupmemberBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "gslbservicegroup_gslbservicegroupmember_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbservicegroupGslbservicegroupmemberBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbservicegroupGslbservicegroupmemberBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading gslbservicegroup_gslbservicegroupmember_binding resource")

	r.readGslbservicegroupGslbservicegroupmemberBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *GslbservicegroupGslbservicegroupmemberBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state GslbservicegroupGslbservicegroupmemberBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating gslbservicegroup_gslbservicegroupmember_binding resource")

	// All attributes are ForceNew (RequiresReplace) mirroring SDK v2, so there are no
	// in-place updatable attributes; changes trigger a replace instead of Update.
	hasChange := false

	if hasChange {
		gslbservicegroup_gslbservicegroupmember_binding := gslbservicegroup_gslbservicegroupmember_bindingGetThePayloadFromtheConfig(ctx, &data)
		_, err := r.client.AddResource(service.Gslbservicegroup_gslbservicegroupmember_binding.Type(), data.Id.ValueString(), &gslbservicegroup_gslbservicegroupmember_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update gslbservicegroup_gslbservicegroupmember_binding, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated gslbservicegroup_gslbservicegroupmember_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for gslbservicegroup_gslbservicegroupmember_binding resource, skipping update")
	}

	// Read the updated state back.
	r.readGslbservicegroupGslbservicegroupmemberBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "gslbservicegroup_gslbservicegroupmember_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbservicegroupGslbservicegroupmemberBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbservicegroupGslbservicegroupmemberBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting gslbservicegroup_gslbservicegroupmember_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// Mirrors SDK v2: delete by servername (== ip for ip-based members) and,
	// when present, port. ParseIdString handles both the new key:value form and the
	// legacy positional "servicegroupname,servername,port" form.
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
		// URL-encode the value so slashy/special chars survive the query string.
		args = append(args, fmt.Sprintf("servername:%s", url.QueryEscape(val)))
	}
	if val, ok := idMap["port"]; ok && val != "" && !optionalAbsent["port"] {
		args = append(args, fmt.Sprintf("port:%s", url.QueryEscape(val)))
	}

	err = r.client.DeleteResourceWithArgs(service.Gslbservicegroup_gslbservicegroupmember_binding.Type(), servicegroupname_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete gslbservicegroup_gslbservicegroupmember_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted gslbservicegroup_gslbservicegroupmember_binding binding")
}

// Helper function to read gslbservicegroup_gslbservicegroupmember_binding data from API
func (r *GslbservicegroupGslbservicegroupmemberBindingResource) readGslbservicegroupGslbservicegroupmemberBindingFromApi(ctx context.Context, data *GslbservicegroupGslbservicegroupmemberBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Gslbservicegroup_gslbservicegroupmember_binding.Type(),
		ResourceName:             servicegroupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbservicegroup_gslbservicegroupmember_binding, got error: %s", err))
		return
	}

	// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
	// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
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

	// Binding not present in the returned set: signal removal via a null Id (see above).
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	gslbservicegroup_gslbservicegroupmember_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
