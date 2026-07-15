package nsaptlicense

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NsaptlicenseResource{}
var _ resource.ResourceWithConfigure = (*NsaptlicenseResource)(nil)
var _ resource.ResourceWithImportState = (*NsaptlicenseResource)(nil)

func NewNsaptlicenseResource() resource.Resource {
	return &NsaptlicenseResource{}
}

// NsaptlicenseResource defines the resource implementation.
type NsaptlicenseResource struct {
	client *service.NitroClient
}

func (r *NsaptlicenseResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsaptlicenseResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsaptlicense"
}

func (r *NsaptlicenseResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsaptlicenseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsaptlicenseResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsaptlicense resource")
	nsaptlicense := nsaptlicenseGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource: NITRO exposes only the `change` verb
	// (POST ?action=update). NOTE: this allocates license counts and is
	// DISRUPTIVE / non-idempotent on the appliance.
	err := r.client.ActOnResource(service.Nsaptlicense.Type(), &nsaptlicense, "update")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsaptlicense, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nsaptlicense resource")

	// The NITRO License ID (data.Id) is the resource identifier and is already
	// populated from the plan; set once here (Pattern 6).

	// Read the updated state back
	r.readNsaptlicenseFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsaptlicenseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsaptlicenseResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsaptlicense resource")

	r.readNsaptlicenseFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Resource was deleted out-of-band; remove it from state so it can be recreated
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsaptlicenseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NsaptlicenseResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state.
	data.Id = state.Id

	// Update is a no-op for nsaptlicense: id and serialno are RequiresReplace,
	// and re-running the license allocation action on every plan would be
	// disruptive. Just read state back.
	tflog.Debug(ctx, "Update is a no-op for nsaptlicense")
	r.readNsaptlicenseFromApi(ctx, &data, &resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsaptlicenseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsaptlicenseResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Action-only resource: NITRO exposes no delete endpoint. Just remove from
	// Terraform state (the allocated licenses remain on the appliance).
	tflog.Trace(ctx, "Removed nsaptlicense from Terraform state (no delete endpoint on NITRO)")
}

// Helper function to read nsaptlicense data from API. GET filters by serialno
// (the GET-only key); the matching record is identified by its License ID.
func (r *NsaptlicenseResource) readNsaptlicenseFromApi(ctx context.Context, data *NsaptlicenseResourceModel, diags *diag.Diagnostics) {

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Nsaptlicense.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsaptlicense, got error: %s", err))
		return
	}

	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Match the record by License ID, falling back to serialno.
	foundIndex := -1
	for i, v := range dataArr {
		if idVal, ok := v["id"].(string); ok && idVal == data.Id.ValueString() {
			foundIndex = i
			break
		}
		if snVal, ok := v["serialno"].(string); ok && !data.Serialno.IsNull() && snVal == data.Serialno.ValueString() {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	nsaptlicenseSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
