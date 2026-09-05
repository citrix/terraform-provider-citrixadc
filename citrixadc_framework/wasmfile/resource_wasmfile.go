package wasmfile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// wasmfileResourceType is the literal NITRO resource-type string. There is no
// adc-nitro-go SDK struct for wasmfile, so the generic NitroClient calls are
// driven with this string directly.
const wasmfileResourceType = "wasmfile"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &WasmfileResource{}
var _ resource.ResourceWithConfigure = (*WasmfileResource)(nil)
var _ resource.ResourceWithImportState = (*WasmfileResource)(nil)

func NewWasmfileResource() resource.Resource {
	return &WasmfileResource{}
}

// WasmfileResource defines the resource implementation.
type WasmfileResource struct {
	client *service.NitroClient
}

func (r *WasmfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wasmfile"
}

func (r *WasmfileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *WasmfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data WasmfileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating wasmfile resource")
	wasmfile := wasmfileGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes wasmfile create only via POST ?action=Import (no `add`).
	// Use ActOnResource with the case-sensitive "Import" verb and an untyped
	// map payload.
	err := r.client.ActOnResource(wasmfileResourceType, wasmfile, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create wasmfile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created wasmfile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readWasmfileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "wasmfile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WasmfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data WasmfileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading wasmfile resource")

	found := r.readWasmfileFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		// Self-healing: resource gone on the ADC, drop it from state.
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WasmfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO's wasmfile `change` (?action=update) accepts only `name` — there is
	// no in-place update path for any content field. Every writable attribute is
	// RequiresReplace, so Terraform forces destroy+recreate on any change and
	// never calls Update with a real field diff. This body is therefore a
	// documented no-op that preserves the prior ID and re-reads state.
	var data, state WasmfileResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for wasmfile; NITRO change only accepts name and all attributes are RequiresReplace")

	if !r.readWasmfileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "wasmfile not found immediately after update")
		}
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WasmfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data WasmfileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting wasmfile resource")
	// NITRO deletes wasmfile via DELETE /wasmfile?args=name:<name> (name is
	// passed as a query arg, not a URL path segment).
	name := data.Name.ValueString()
	args := []string{fmt.Sprintf("name:%s", name)}
	err := r.client.DeleteResourceWithArgs(wasmfileResourceType, "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete wasmfile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted wasmfile resource")
}

func (r *WasmfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// readWasmfileFromApi fetches the wasmfile matching the model's name via the
// NITRO get(all) endpoint filtered by name. Returns false (with no diagnostic)
// when the resource is absent, so callers can self-heal.
func (r *WasmfileResource) readWasmfileFromApi(ctx context.Context, data *WasmfileResourceModel, diags *diag.Diagnostics) bool {
	// Use the ID (which holds the name) so the read works on import too, where only
	// id is populated (data.Name would be empty -> an invalid empty NITRO filter).
	name := data.Id.ValueString()

	findParams := service.FindParams{
		ResourceType:             wasmfileResourceType,
		FilterMap:                map[string]string{"name": name},
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read wasmfile, got error: %s", err))
		return false
	}

	// Resource is missing.
	if len(dataArr) == 0 {
		return false
	}

	// The filter targets `name`, so match the exact entry defensively.
	for _, entry := range dataArr {
		if v, ok := entry["name"]; ok && v != nil {
			if v.(string) == name {
				wasmfileSetAttrFromGet(ctx, data, entry)
				return true
			}
		}
	}

	// Filter did not narrow to the exact name; fall back to the first entry.
	wasmfileSetAttrFromGet(ctx, data, dataArr[0])
	return true
}
