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
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ip:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ip.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("port:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Port.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("servername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servername.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

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

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicegroupname", "servername", "ip", "port"}, []string{"servername", "ip", "port"})
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
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
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbservicegroup_gslbservicegroupmember_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "gslbservicegroup_gslbservicegroupmember_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
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

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("gslbservicegroup_gslbservicegroupmember_binding not found with the provided ID attributes"))
		return
	}

	gslbservicegroup_gslbservicegroupmember_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
