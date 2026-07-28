package clusterpropstatus

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ClusterpropstatusClearResource{}
var _ resource.ResourceWithConfigure = (*ClusterpropstatusClearResource)(nil)

func NewClusterpropstatusClearResource() resource.Resource {
	return &ClusterpropstatusClearResource{}
}

// ClusterpropstatusClearResource defines the resource implementation.
//
// This resource models the NITRO clusterpropstatus `?action=clear` action. clear
// is a one-shot side-effect action (POST) that resets the property-propagation
// status counters. NITRO exposes no add, update/set, or delete endpoint for
// clusterpropstatus, so Read/Update/Delete are no-ops.
//
// The clear verb takes NO arguments: the NITRO clear payload is empty
// ({"clusterpropstatus":{}}) and the live CLI (`clear cluster propstatus`)
// rejects `-nodeid` with "No such argument". nodeid is a GET-only filter
// (Pattern 15) that belongs to the get/count datasource side, not this action,
// so it is intentionally excluded here.
type ClusterpropstatusClearResource struct {
	client *service.NitroClient
}

// ClusterpropstatusClearResourceModel describes the resource data model.
type ClusterpropstatusClearResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *ClusterpropstatusClearResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusterpropstatus_clear"
}

func (r *ClusterpropstatusClearResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusterpropstatusClearResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the clusterpropstatus_clear resource.",
			},
			// The clear action accepts no arguments (verified against the live CLI:
			// `clear cluster propstatus` takes zero arguments). No config attributes
			// are exposed.
		},
	}
}

func (r *ClusterpropstatusClearResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusterpropstatusClearResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Clearing clusterpropstatus (action-only resource)")
	payload := clusterpropstatus_clearGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes clear as POST ?action=clear. Use ActOnResource with the
	// case-sensitive "clear" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Clusterpropstatus.Type(), payload, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear clusterpropstatus, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared clusterpropstatus")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("clusterpropstatus_clear")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterpropstatusClearResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// clear is a one-shot action. NITRO has no GET endpoint that reports
	// clear-state, so Read is a pure preserve-state no-op.
	var data ClusterpropstatusClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for clusterpropstatus_clear; NITRO has no query endpoint for clear state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterpropstatusClearResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for clear; the resource has no config
	// attributes, so Terraform never invokes Update for a real change.
	var data, state ClusterpropstatusClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for clusterpropstatus_clear; NITRO has no update endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterpropstatusClearResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// clear is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-clear"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for clusterpropstatus_clear; NITRO has no inverse of the clear action")
}

// clusterpropstatus_clearGetThePayloadFromthePlan builds the body for the clear
// action. The clear verb accepts no arguments (confirmed by the live CLI and the
// empty NITRO clear payload {"clusterpropstatus":{}}), so the returned map is
// always empty. Returning map[string]interface{} matches the original
// implementation's payload type.
func clusterpropstatus_clearGetThePayloadFromthePlan(ctx context.Context, data *ClusterpropstatusClearResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In clusterpropstatus_clearGetThePayloadFromthePlan Function")

	clusterpropstatus := map[string]interface{}{}

	return clusterpropstatus
}
