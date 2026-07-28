package gslbservicegroup_gslbservicegroupmember_binding

import (
	"context"
	"fmt"
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
var _ resource.ResourceWithValidateConfig = (*GslbservicegroupGslbservicegroupmemberBindingResource)(nil)

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

// ValidateConfig enforces the NITRO mutually-exclusive mandatory choice between the IP-path member
// (ip) and the server-name-path member (servername): exactly one must be set.
func (r *GslbservicegroupGslbservicegroupmemberBindingResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data GslbservicegroupGslbservicegroupmemberBindingResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// At-least-one-of(ip, servername)
	if data.Ip.IsNull() && data.Servername.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("ip"),
			"Missing Required Attribute",
			"Exactly one of \"ip\" or \"servername\" must be specified.",
		)
	}

	// Mutually exclusive: ip and servername cannot both be set
	if !data.Ip.IsNull() && !data.Servername.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("ip"),
			"Conflicting Attributes",
			"Only one of \"ip\" or \"servername\" may be specified, not both.",
		)
	}
}

func (r *GslbservicegroupGslbservicegroupmemberBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbservicegroupGslbservicegroupmemberBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbservicegroup_gslbservicegroupmember_binding resource")
	gslbservicegroup_gslbservicegroupmember_binding := gslbservicegroup_gslbservicegroupmember_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Gslbservicegroup_gslbservicegroupmember_binding.Type(), &gslbservicegroup_gslbservicegroupmember_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create gslbservicegroup_gslbservicegroupmember_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created gslbservicegroup_gslbservicegroupmember_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(gslbservicegroup_gslbservicegroupmember_bindingBuildId(
		data.Servicegroupname.ValueString(),
		data.Servername.ValueString(),
		data.Ip.ValueString(),
		data.Port.ValueInt64(),
	))

	// Read the updated state back
	r.readGslbservicegroupGslbservicegroupmemberBindingFromApi(ctx, &data, &resp.Diagnostics)

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

	// Update is a no-op for this binding: NITRO exposes only add (bind) / delete (unbind),
	// there is no update/set endpoint, and every schema attribute is RequiresReplace.
	tflog.Debug(ctx, "Update is a no-op for gslbservicegroup_gslbservicegroupmember_binding; all attributes are RequiresReplace (bind/unbind only)")

	// Read the current state back
	r.readGslbservicegroupGslbservicegroupmemberBindingFromApi(ctx, &data, &resp.Diagnostics)

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

	// Legacy SDK v2 positional id ("servicegroupname,<servername-or-ip>,port"): delete the
	// way SDK v2 did (servername:token,port). This works for both servername- and ip-bound
	// members because ADC auto-names an ip-bound server == ip. Handled explicitly because the
	// 4-element ParseIdString order mis-maps a 3-token id (port -> ip). (A prior Read normally
	// rewrites the id to the new format first; this is defense-in-depth for destroy -refresh=false.)
	if utils.IsLegacyIdFormat(data.Id.ValueString()) {
		parts := strings.SplitN(data.Id.ValueString(), ",", 3)
		if len(parts) < 1 || parts[0] == "" {
			resp.Diagnostics.AddError("Parse Error", "Parent attribute 'servicegroupname' not found in ID")
			return
		}
		var legacyArgs []string
		if len(parts) > 1 && parts[1] != "" {
			legacyArgs = append(legacyArgs, fmt.Sprintf("servername:%s", utils.UrlEncode(parts[1])))
		}
		if len(parts) > 2 && parts[2] != "" {
			legacyArgs = append(legacyArgs, fmt.Sprintf("port:%s", utils.UrlEncode(parts[2])))
		}
		if err := r.client.DeleteResourceWithArgs(service.Gslbservicegroup_gslbservicegroupmember_binding.Type(), parts[0], legacyArgs); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete gslbservicegroup_gslbservicegroupmember_binding, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Deleted gslbservicegroup_gslbservicegroupmember_binding binding")
		return
	}

	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicegroupname", "servername", "ip", "port"}, []string{"servername", "ip", "port"})
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	servicegroupname_value, ok := idMap["servicegroupname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'servicegroupname' not found in ID")
		return
	}

	// Build delete args; UrlEncode the values since ip may be an IPv6 address containing colons,
	// which would otherwise break the args=ip:..,servername:..,port:.. parsing.
	var args []string
	if val, ok := idMap["ip"]; ok && val != "" {
		args = append(args, fmt.Sprintf("ip:%s", utils.UrlEncode(val)))
	}
	if val, ok := idMap["servername"]; ok && val != "" {
		args = append(args, fmt.Sprintf("servername:%s", utils.UrlEncode(val)))
	}
	if val, ok := idMap["port"]; ok && val != "" {
		args = append(args, fmt.Sprintf("port:%s", utils.UrlEncode(val)))
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

	idStr := data.Id.ValueString()

	// legacy marks an SDK v2 positional id ("servicegroupname,<servername-or-ip>,port").
	// The 4-element ParseIdString order [servicegroupname,servername,ip,port] mis-maps a
	// 3-token legacy id (the port lands in the ip slot), so parse it explicitly here and
	// resolve the ambiguous middle token to servername/ip after the GET.
	var idMap map[string]string
	var legacyToken string
	legacy := utils.IsLegacyIdFormat(idStr)
	if legacy {
		parts := strings.SplitN(idStr, ",", 3)
		idMap = map[string]string{}
		if len(parts) > 0 {
			idMap["servicegroupname"] = parts[0]
		}
		if len(parts) > 1 {
			legacyToken = parts[1]
		}
		if len(parts) > 2 {
			idMap["port"] = parts[2]
		}
	} else {
		// Case 4: Array filter with parent ID - parse from ID
		var err error
		idMap, _, err = utils.ParseIdString(idStr, []string{"servicegroupname", "servername", "ip", "port"}, []string{"servername", "ip", "port"})
		if err != nil {
			diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
			return
		}
	}

	servicegroupname_Name, ok := idMap["servicegroupname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'servicegroupname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Gslbservicegroup_gslbservicegroupmember_binding.Type(),
		ResourceName:             servicegroupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbservicegroup_gslbservicegroupmember_binding, got error: %s", err))
		return
	}

	// Resource is missing (deleted out-of-band) - signal removal from state.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if legacy {
			// Legacy match: port must match and the ambiguous token must equal either the
			// member's servername or its ip. (ADC auto-names an ip-bound server == ip, so a
			// servername match covers most ip bindings; the ip clause covers pure ip/autoscale
			// members whose servername is empty.)
			if portStr, ok := idMap["port"]; ok && portStr != "" {
				vp, ok := v["port"]
				if !ok {
					continue
				}
				vpInt, _ := utils.ConvertToInt64(vp)
				tokenPort, _ := strconv.ParseInt(portStr, 10, 64)
				if vpInt != tokenPort {
					continue
				}
			}
			vServername, _ := v["servername"].(string)
			vIp, _ := v["ip"].(string)
			if vServername == legacyToken || vIp == legacyToken {
				foundIndex = i
				break
			}
			continue
		}

		match := true

		// Check ip (ip/servername are a mutually-exclusive choice; the unused one is empty in the ID,
		// so only filter on ip when the ID actually carries a non-empty ip value)
		if idVal, ok := idMap["ip"]; ok && idVal != "" {
			if val, ok := v["ip"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		// Check port
		if idVal, ok := idMap["port"]; ok {
			if val, ok := v["port"]; ok {
				val, _ = utils.ConvertToInt64(val)
				idValInt64, _ := strconv.ParseInt(idVal, 10, 64)
				if val != idValInt64 {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["port"]; ok {
			match = false
			continue
		}

		// Check servername (mutually-exclusive with ip; only filter when the ID carries a non-empty value)
		if idVal, ok := idMap["servername"]; ok && idVal != "" {
			if val, ok := v["servername"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing (deleted out-of-band) - signal removal from state.
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	member := dataArr[foundIndex]
	gslbservicegroup_gslbservicegroupmember_bindingSetAttrFromGet(ctx, data, member)

	if legacy {
		// Normalize legacy SDK v2 (v2.2.x) state to the Framework model and rewrite the id
		// to the new key:value format, so subsequent plans take the fast (non-legacy) path.
		var memberPort int64
		if portStr, ok := idMap["port"]; ok && portStr != "" {
			memberPort, _ = strconv.ParseInt(portStr, 10, 64)
		} else if vp, ok := member["port"]; ok {
			memberPort, _ = utils.ConvertToInt64(vp)
		}

		// Resolve the ambiguous token: if it equals the member's ip, the user bound by ip
		// (ADC created a server named == ip); otherwise they bound by servername. Populate
		// exactly one of servername/ip (mirroring Create) so the refreshed state matches a
		// config that sets only that field (both are RequiresReplace).
		memberIp, _ := member["ip"].(string)
		if memberIp == legacyToken {
			data.Ip = types.StringValue(legacyToken)
			data.Servername = types.StringNull()
		} else {
			data.Servername = types.StringValue(legacyToken)
			data.Ip = types.StringNull()
		}
		data.Port = types.Int64Value(memberPort)
		data.Servicegroupname = types.StringValue(servicegroupname_Name)

		// SDK v2 marked several optional attrs Computed and stored server-echoed zero/empty
		// values; the Framework treats them as plain Optional (RequiresReplace). Coerce those
		// empties to null so an upgrade against a config that omits them does not force a
		// spurious replace.
		if data.Hashid.ValueInt64() == 0 {
			data.Hashid = types.Int64Null()
		}
		if data.Order.ValueInt64() == 0 {
			data.Order = types.Int64Null()
		}
		if data.Publicport.ValueInt64() == 0 {
			data.Publicport = types.Int64Null()
		}
		if data.Publicip.ValueString() == "" {
			data.Publicip = types.StringNull()
		}
		if data.Siteprefix.ValueString() == "" {
			data.Siteprefix = types.StringNull()
		}

		data.Id = types.StringValue(gslbservicegroup_gslbservicegroupmember_bindingBuildId(
			data.Servicegroupname.ValueString(),
			data.Servername.ValueString(),
			data.Ip.ValueString(),
			memberPort,
		))
	}
}
