package dnssvcbrec

import (
	"context"
	"fmt"
	"net/url"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DnssvcbrecResource{}
var _ resource.ResourceWithConfigure = (*DnssvcbrecResource)(nil)
var _ resource.ResourceWithImportState = (*DnssvcbrecResource)(nil)

func NewDnssvcbrecResource() resource.Resource {
	return &DnssvcbrecResource{}
}

// DnssvcbrecResource defines the resource implementation.
type DnssvcbrecResource struct {
	client *service.NitroClient
}

// idAttrOrder is the composite-ID field order ("domain,targetname,priority,svcbtype").
// It matches the record identity used by the NITRO delete operation.
var idAttrOrder = []string{"domain", "targetname", "priority", "svcbtype"}

func (r *DnssvcbrecResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnssvcbrecResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnssvcbrec"
}

func (r *DnssvcbrecResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnssvcbrecResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnssvcbrecResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnssvcbrec resource")

	dnssvcbrec := dnssvcbrecGetThePayloadFromthePlan(ctx, &data)

	// Named/composite-key resource - use AddResource (POST /nitro/v1/config/dnssvcbrec).
	// The name argument mirrors the record identity "domain,targetname,priority,svcbtype".
	dnssvcbrecName := fmt.Sprintf("%s,%s,%d,%s",
		data.Domain.ValueString(),
		data.Targetname.ValueString(),
		data.Priority.ValueInt64(),
		data.Svcbtype.ValueString())
	_, err := r.client.AddResource(service.Dnssvcbrec.Type(), dnssvcbrecName, &dnssvcbrec)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnssvcbrec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dnssvcbrec resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(dnssvcbrecName)

	// Read the created state back
	if !r.readDnssvcbrecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnssvcbrec not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnssvcbrecResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnssvcbrecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnssvcbrec resource")

	found := r.readDnssvcbrecFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		// Self-healing: the record no longer exists on the appliance.
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnssvcbrecResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state DnssvcbrecResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (identity fields are RequiresReplace so it is stable)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating dnssvcbrec resource")

	// Check for changes in updateable attributes, and detect attributes removed
	// from config so they can be unset (reverted to their NITRO defaults).
	hasChange := false
	attributesToUnset := []string{}

	if !data.Nodeid.Equal(state.Nodeid) {
		tflog.Debug(ctx, "nodeid has changed for dnssvcbrec")
		hasChange = true
	}
	if !data.Alpn.Equal(state.Alpn) {
		if config.Alpn.IsNull() {
			attributesToUnset = append(attributesToUnset, "alpn")
		} else {
			hasChange = true
		}
	}
	if !data.Encryptedclienthello.Equal(state.Encryptedclienthello) {
		if config.Encryptedclienthello.IsNull() {
			attributesToUnset = append(attributesToUnset, "encryptedclienthello")
		} else {
			hasChange = true
		}
	}
	if !data.Ipv4hint.Equal(state.Ipv4hint) {
		if config.Ipv4hint.IsNull() {
			attributesToUnset = append(attributesToUnset, "ipv4hint")
		} else {
			hasChange = true
		}
	}
	if !data.Ipv6hint.Equal(state.Ipv6hint) {
		if config.Ipv6hint.IsNull() {
			attributesToUnset = append(attributesToUnset, "ipv6hint")
		} else {
			hasChange = true
		}
	}
	if !data.Mandatory.Equal(state.Mandatory) {
		if config.Mandatory.IsNull() {
			attributesToUnset = append(attributesToUnset, "mandatory")
		} else {
			hasChange = true
		}
	}
	if !data.Nodefaultalpn.Equal(state.Nodefaultalpn) {
		if config.Nodefaultalpn.IsNull() {
			attributesToUnset = append(attributesToUnset, "nodefaultalpn")
		} else {
			hasChange = true
		}
	}
	if !data.Port.Equal(state.Port) {
		if config.Port.IsNull() {
			attributesToUnset = append(attributesToUnset, "port")
		} else {
			hasChange = true
		}
	}
	if !data.Ttl.Equal(state.Ttl) {
		if config.Ttl.IsNull() {
			attributesToUnset = append(attributesToUnset, "ttl")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Identity fields (domain, targetname, priority, svcbtype) are always
		// included so NITRO can locate the record.
		dnssvcbrec := dnssvcbrecGetThePayloadFromthePlan(ctx, &data)
		// Update uses PUT /nitro/v1/config/dnssvcbrec (UpdateUnnamedResource).
		err := r.client.UpdateUnnamedResource(service.Dnssvcbrec.Type(), &dnssvcbrec)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dnssvcbrec, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated dnssvcbrec resource")
	} else {
		tflog.Debug(ctx, "No changes detected for dnssvcbrec resource, skipping update")
	}

	// Unset attributes removed from config so the appliance reverts them to
	// their NITRO defaults. The identity fields locate the record.
	unsetIdPayload := map[string]interface{}{
		"domain":     data.Domain.ValueString(),
		"targetname": data.Targetname.ValueString(),
		"priority":   int(state.Priority.ValueInt64()),
		"svcbtype":   data.Svcbtype.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Dnssvcbrec.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset dnssvcbrec attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readDnssvcbrecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnssvcbrec not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnssvcbrecResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnssvcbrecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnssvcbrec resource")

	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), idAttrOrder, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse dnssvcbrec ID, got error: %s", err))
		return
	}
	domain := idMap["domain"]

	// DELETE /nitro/v1/config/dnssvcbrec/<domain>?args=targetname:<>,priority:<>,svcbtype:<>
	// disambiguates the record when multiple exist for the same domain.
	argsMap := make(map[string]string)
	argsMap["targetname"] = url.QueryEscape(idMap["targetname"])
	argsMap["priority"] = url.QueryEscape(idMap["priority"])
	argsMap["svcbtype"] = url.QueryEscape(idMap["svcbtype"])

	err = r.client.DeleteResourceWithArgsMap(service.Dnssvcbrec.Type(), domain, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnssvcbrec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnssvcbrec resource")
}

// readDnssvcbrecFromApi reads dnssvcbrec data from the appliance. It returns
// false when the record no longer exists (so callers can drop it from state).
func (r *DnssvcbrecResource) readDnssvcbrecFromApi(ctx context.Context, data *DnssvcbrecResourceModel, diags *diag.Diagnostics) bool {
	// dnssvcbrec exposes only a "get (all)" endpoint, so fetch the array and
	// filter by the domain+targetname+priority+svcbtype identity.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), idAttrOrder, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse dnssvcbrec ID, got error: %s", err))
		return false
	}
	domain := idMap["domain"]
	targetname := idMap["targetname"]
	priority := idMap["priority"]
	svcbtype := idMap["svcbtype"]

	// dnssvcbrec's get REQUIRES the svcbtype key arg (get_keys=[svcbtype,type]); a
	// plain get-all with no args returns an empty set even when the record exists.
	// Pass ?args=svcbtype:<svcbtype> so NITRO returns the records of that type,
	// then match the remaining identity fields client-side.
	findParams := service.FindParams{
		ResourceType:             service.Dnssvcbrec.Type(),
		ArgsMap:                  map[string]string{"svcbtype": svcbtype},
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnssvcbrec, got error: %s", err))
		return false
	}

	if len(dataArr) == 0 {
		return false
	}

	foundIndex := -1
	for i, v := range dataArr {
		if d, ok := v["domain"].(string); !ok || d != domain {
			continue
		}
		if t, ok := v["targetname"].(string); !ok || t != targetname {
			continue
		}
		if s, ok := v["svcbtype"].(string); !ok || s != svcbtype {
			continue
		}
		// priority may be omitted from GET for AliasMode records (priority=0, which
		// NITRO drops from the response); treat a missing/unparseable value as "0".
		gotPriority := "0"
		if pv, ok := v["priority"]; ok && pv != nil {
			if intVal, cerr := utils.ConvertToInt64(pv); cerr == nil {
				gotPriority = fmt.Sprintf("%d", intVal)
			}
		}
		if gotPriority != priority {
			continue
		}
		foundIndex = i
		break
	}
	if foundIndex == -1 {
		return false
	}

	dnssvcbrecSetAttrFromGet(ctx, data, dataArr[foundIndex])

	return true
}
