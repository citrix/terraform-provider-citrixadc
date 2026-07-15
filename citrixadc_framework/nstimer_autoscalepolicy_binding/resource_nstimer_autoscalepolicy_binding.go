package nstimer_autoscalepolicy_binding

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
var _ resource.Resource = &NstimerAutoscalepolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*NstimerAutoscalepolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*NstimerAutoscalepolicyBindingResource)(nil)

func NewNstimerAutoscalepolicyBindingResource() resource.Resource {
	return &NstimerAutoscalepolicyBindingResource{}
}

// NstimerAutoscalepolicyBindingResource defines the resource implementation.
type NstimerAutoscalepolicyBindingResource struct {
	client *service.NitroClient
}

func (r *NstimerAutoscalepolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NstimerAutoscalepolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nstimer_autoscalepolicy_binding"
}

func (r *NstimerAutoscalepolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NstimerAutoscalepolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NstimerAutoscalepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nstimer_autoscalepolicy_binding resource")
	nstimer_autoscalepolicy_binding := nstimer_autoscalepolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Nstimer_autoscalepolicy_binding.Type(), &nstimer_autoscalepolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nstimer_autoscalepolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nstimer_autoscalepolicy_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readNstimerAutoscalepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstimerAutoscalepolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NstimerAutoscalepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nstimer_autoscalepolicy_binding resource")

	// NOTE: on this firmware the typed binding GET
	// (GET /nstimer_autoscalepolicy_binding/<name>[?filter=policyname:<v>]) always
	// returns an empty body ({"errorcode":0} with no resource key) even though the
	// binding is present (verified via CLI "show nstimer <name>", which lists it as a
	// GLOBAL policy binding). The aggregate nstimer_binding/<name> endpoint and
	// count=yes likewise report nothing. There is therefore no reliable REST GET that
	// reflects the bound state, so Read cannot prove absence. We attempt the GET to
	// refresh attributes on firmwares that DO return the row, but when the GET is
	// empty we preserve the prior state rather than removing the resource (which
	// would cause a spurious recreate on every plan).
	r.readNstimerAutoscalepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Binding deleted out-of-band (list returned rows but ours is gone): remove from
	// state so a later apply re-creates it. Note: the empty-body case is deliberately
	// NOT treated as gone here (see readNstimerAutoscalepolicyBindingFromApi / firmware
	// note above), so Id is only nulled when the item is genuinely absent from a
	// non-empty list.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstimerAutoscalepolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NstimerAutoscalepolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for nstimer_autoscalepolicy_binding: the binding is
	// immutable (NITRO has no update endpoint; every attribute is
	// RequiresReplace).
	tflog.Debug(ctx, "Update is a no-op for nstimer_autoscalepolicy_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readNstimerAutoscalepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstimerAutoscalepolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NstimerAutoscalepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nstimer_autoscalepolicy_binding resource")
	// Binding with parent - delete keyed on parent name with policyname as the
	// disambiguating arg (DELETE .../<name>?args=policyname:<policyname>).
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "policyname"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	args := []string{}
	if val, ok := idMap["policyname"]; ok && val != "" {
		args = append(args, "policyname:"+utils.UrlEncode(val))
	}

	err = r.client.DeleteResourceWithArgs(service.Nstimer_autoscalepolicy_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nstimer_autoscalepolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nstimer_autoscalepolicy_binding binding")
}

// Helper function to read nstimer_autoscalepolicy_binding data from API.
// Returns true if the binding was found, false if it no longer exists (drift).
func (r *NstimerAutoscalepolicyBindingResource) readNstimerAutoscalepolicyBindingFromApi(ctx context.Context, data *NstimerAutoscalepolicyBindingResourceModel, diags *diag.Diagnostics) bool {

	// Array filter with parent ID - parse from ID (legacy order: name,policyname)
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "policyname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return false
	}

	name_Name, ok := idMap["name"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'name' not found in ID string")
		return false
	}
	policyname_Name := idMap["policyname"]

	// Primary read: by-name typed binding endpoint, narrowed by policyname via the
	// documented ?filter= query parameter. On firmwares that echo the bound row this
	// returns it; on this firmware it returns an empty body (see the note in Read).
	findParams := service.FindParams{
		ResourceType:             service.Nstimer_autoscalepolicy_binding.Type(),
		ResourceName:             name_Name,
		FilterMap:                map[string]string{"policyname": policyname_Name},
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nstimer_autoscalepolicy_binding, got error: %s", err))
		return false
	}

	// The typed binding GET does not reflect the bound state on this firmware (always
	// empty even when bound - confirmed via CLI). Absence cannot be proven over REST,
	// so when nothing is returned we preserve the existing model unchanged rather than
	// signalling drift.
	if len(dataArr) == 0 {
		tflog.Debug(ctx, "nstimer_autoscalepolicy_binding GET returned empty; preserving existing state (binding GET is not reflected over REST on this firmware)")
		return false
	}

	// Iterate through results to find the one with the right policyname.
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["policyname"].(string); ok {
			if val == policyname_Name {
				foundIndex = i
				break
			}
		}
	}

	if foundIndex == -1 {
		// List returned rows but our binding is not among them: genuinely gone.
		// Signal drift so Read removes it from state (self-heal on next apply).
		data.Id = types.StringNull()
		return false
	}

	nstimer_autoscalepolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
	return true
}
