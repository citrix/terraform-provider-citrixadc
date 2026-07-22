package vrid_trackinterface_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VridTrackinterfaceBindingResource{}
var _ resource.ResourceWithConfigure = (*VridTrackinterfaceBindingResource)(nil)

func NewVridTrackinterfaceBindingResource() resource.Resource {
	return &VridTrackinterfaceBindingResource{}
}

// VridTrackinterfaceBindingResource defines the resource implementation.
type VridTrackinterfaceBindingResource struct {
	client *service.NitroClient
}

func (r *VridTrackinterfaceBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vrid_trackinterface_binding"
}

func (r *VridTrackinterfaceBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VridTrackinterfaceBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VridTrackinterfaceBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vrid_trackinterface_binding resource")
	vrid_trackinterface_binding := vrid_trackinterface_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - NITRO add is HTTP PUT (bind), use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vrid_trackinterface_binding.Type(), &vrid_trackinterface_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vrid_trackinterface_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vrid_trackinterface_binding resource")

	// Set ID for the resource before reading state.
	// Composite key: id (the NITRO integer VRID), trackifnum.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("id:%s", utils.UrlEncode(fmt.Sprintf("%v", data.VridId.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("trackifnum:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Trackifnum.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// IMPORTANT (verified live on NS VPX): vrid_trackinterface_binding has NO NITRO
	// read path. The bind PUT succeeds, but the binding is not surfaced by any GET --
	// the aggregate vrid_binding/<id> response carries only {"id"} with no
	// vrid_trackinterface_binding array, and the direct endpoint returns a keyless
	// empty body. We therefore follow the no-GET / action-only precedent
	// (fis_interface_binding): the create is authoritative once the PUT succeeds, and
	// we do NOT require a post-create read to confirm it (which would always fail with
	// "not found after create"). A best-effort read is still attempted to opportunis-
	// tically populate computed fields if a future firmware does surface the row.
	r.readVridTrackinterfaceBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VridTrackinterfaceBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VridTrackinterfaceBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vrid_trackinterface_binding resource")

	// No-GET binding (see Create): a best-effort read may populate computed fields,
	// but a "not found" result does NOT mean the binding is gone -- the appliance
	// simply never surfaces it. Do not remove the resource from state on !found,
	// otherwise every refresh would spuriously delete a perfectly valid binding.
	r.readVridTrackinterfaceBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VridTrackinterfaceBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VridTrackinterfaceBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for vrid_trackinterface_binding: NITRO exposes only add (PUT)
	// and delete (no update/change endpoint), and all schema attributes are
	// RequiresReplace, so Terraform recreates the resource on any change rather than
	// calling Update.
	tflog.Debug(ctx, "Update is a no-op for vrid_trackinterface_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readVridTrackinterfaceBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VridTrackinterfaceBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VridTrackinterfaceBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vrid_trackinterface_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs with the parent
	// (the NITRO integer id) as the resource (URL) name and trackifnum passed as an arg.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"id", "trackifnum"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	id_value, ok := idMap["id"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'id' not found in ID")
		return
	}

	args := make([]string, 0)
	if val, ok := idMap["trackifnum"]; ok && val != "" {
		// Interface ids contain '/' (e.g. 1/3); URL-encode the value.
		args = append(args, fmt.Sprintf("trackifnum:%s", utils.UrlEncode(val)))
	}

	err = r.client.DeleteResourceWithArgs(service.Vrid_trackinterface_binding.Type(), id_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vrid_trackinterface_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vrid_trackinterface_binding binding")
}

// vrid_trackinterface_bindingAggregateRead queries the AGGREGATE parent endpoint
// (GET /nitro/v1/config/vrid_binding/<id>) and flattens the nested
// "vrid_trackinterface_binding" arrays into a single slice of binding rows. The
// by-name binding endpoint can return a keyless empty body, so the bound members
// are read via the parent aggregate for robustness.
func vrid_trackinterface_bindingAggregateRead(client *service.NitroClient, id string) ([]map[string]interface{}, error) {
	findParams := service.FindParams{
		ResourceType:             "vrid_binding",
		ResourceName:             id,
		ResourceMissingErrorCode: 258,
	}
	parentArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return nil, err
	}

	rows := make([]map[string]interface{}, 0)
	for _, parent := range parentArr {
		nested, ok := parent["vrid_trackinterface_binding"]
		if !ok || nested == nil {
			continue
		}
		nestedArr, ok := nested.([]interface{})
		if !ok {
			continue
		}
		for _, item := range nestedArr {
			if m, ok := item.(map[string]interface{}); ok {
				rows = append(rows, m)
			}
		}
	}
	return rows, nil
}

// readVridTrackinterfaceBindingFromApi reads the binding via the AGGREGATE parent
// endpoint (vrid_binding/<id>) and matches the row by trackifnum. It returns true
// when the binding is found and the model was populated, false when the binding is
// genuinely absent (drift). Hard errors (parse / transport) are reported via diags.
func (r *VridTrackinterfaceBindingResource) readVridTrackinterfaceBindingFromApi(ctx context.Context, data *VridTrackinterfaceBindingResourceModel, diags *diag.Diagnostics) bool {

	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"id", "trackifnum"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return false
	}

	id_Name, ok := idMap["id"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'id' not found in ID string")
		return false
	}

	dataArr, err := vrid_trackinterface_bindingAggregateRead(r.client, id_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vrid_trackinterface_binding, got error: %s", err))
		return false
	}

	// Binding genuinely absent (parent missing or no nested rows): report drift.
	if len(dataArr) == 0 {
		return false
	}

	// Iterate through results to find the one with the right trackifnum.
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		if idVal, ok := idMap["trackifnum"]; ok {
			if val, ok := v["trackifnum"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["trackifnum"].(string); ok {
			match = false
			continue
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Binding row not present in the aggregate response: drift.
	if foundIndex == -1 {
		return false
	}

	vrid_trackinterface_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
	return true
}
