package gslbldnsentry

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GslbldnsentryResource{}
var _ resource.ResourceWithConfigure = (*GslbldnsentryResource)(nil)
var _ resource.ResourceWithImportState = (*GslbldnsentryResource)(nil)

func NewGslbldnsentryResource() resource.Resource {
	return &GslbldnsentryResource{}
}

// GslbldnsentryResource defines the resource implementation.
type GslbldnsentryResource struct {
	client *service.NitroClient
}

func (r *GslbldnsentryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbldnsentryResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbldnsentry"
}

func (r *GslbldnsentryResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// Create is a delete-as-create action resource.
//
// gslbldnsentry is an UNUSUAL "delete-only" NITRO resource: the NITRO API
// exposes ONLY the `delete` verb (no add/get/update/count/clear). The only
// thing you can do with it is remove a single runtime-learned LDNS entry by
// its IP address (CLI: `rm gslb ldnsentry <ipaddress>`).
//
// We therefore model it as "delete-as-create": APPLYING this resource
// performs the NITRO HTTP DELETE that REMOVES the learned LDNS entry with the
// given ipaddress. The delete is keyed solely by the query arg
// `args=ipaddress:<value>` (there is no URL name segment). Because the action
// IS an HTTP DELETE, Create calls DeleteResourceWithArgs directly (NOT
// AddResource/UpdateUnnamedResource/ActOnResource).
func (r *GslbldnsentryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbldnsentryResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ipaddress := data.Ipaddress.ValueString()
	tflog.Debug(ctx, fmt.Sprintf("Creating gslbldnsentry resource (delete-as-create): removing learned LDNS entry %s", ipaddress))

	// The NITRO "action" for this resource is an HTTP DELETE keyed solely by
	// args=ipaddress:<value>. Call DeleteResourceWithArgs directly.
	args := []string{"ipaddress:" + utils.UrlEncode(ipaddress)}
	err := r.client.DeleteResourceWithArgs(service.Gslbldnsentry.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to remove gslbldnsentry %s, got error: %s", ipaddress, err))
		return
	}

	tflog.Trace(ctx, "Removed learned LDNS entry via gslbldnsentry resource")

	// Set ID for the resource (equals the ipaddress acted upon).
	data.Id = types.StringValue(ipaddress)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbldnsentryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Read is a no-op: NITRO exposes no GET endpoint for gslbldnsentry
	// (the entry is removed at create time and there is nothing to read back).
	// Preserve prior state unchanged.
	var data GslbldnsentryResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for gslbldnsentry; no GET endpoint on NITRO side")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbldnsentryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update is a no-op: ipaddress is the only attribute and it is
	// RequiresReplace, so a change forces a destroy/recreate and Update is
	// never reached with an actual change.
	var data, state GslbldnsentryResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for gslbldnsentry; ipaddress is RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbldnsentryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Delete (Terraform destroy) is a no-op: the learned LDNS entry was already
	// removed at create time. Re-deleting a non-existent learned entry would
	// error, so we simply drop the resource from state.
	var data GslbldnsentryResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Delete is a no-op for gslbldnsentry; entry was already removed at create. Dropping from state only.")
}
