package hasync

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &HasyncForceResource{}
var _ resource.ResourceWithConfigure = (*HasyncForceResource)(nil)

func NewHasyncForceResource() resource.Resource {
	return &HasyncForceResource{}
}

// HasyncForceResource defines the resource implementation.
type HasyncForceResource struct {
	client *service.NitroClient
}

// HasyncForceResourceModel describes the resource data model.
//
// This resource models the NITRO hasync `?action=Force` action. Force is a
// one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The Force payload carries the optional
// attributes force and save.
type HasyncForceResourceModel struct {
	Id    types.String `tfsdk:"id"`
	Force types.Bool   `tfsdk:"force"`
	Save  types.String `tfsdk:"save"`
}

func (r *HasyncForceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hasync_force"
}

func (r *HasyncForceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *HasyncForceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the hasync_force resource.",
			},
			"force": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Force synchronization regardless of the state of HA propagation and HA synchronization on either node.",
			},
			"save": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "After synchronization, automatically save the configuration in the secondary node configuration file (ns.conf) without prompting for confirmation.",
			},
		},
	}
}

func (r *HasyncForceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data HasyncForceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating hasync_force resource")

	// hasync exposes only the POST ?action=Force action on NITRO (capital F,
	// case-sensitive). There is no add/get/update/delete endpoint. Use
	// ActOnResource with the exact "Force" verb.
	payload := hasync_forceGetThePayloadFromthePlan(ctx, &data)

	err := r.client.ActOnResource(service.Hasync.Type(), payload, "Force")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to Force sync hasync, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Force synced hasync_force resource")

	// Synthetic constant ID - there is no NITRO identity for this action resource.
	data.Id = types.StringValue("hasync_force")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HasyncForceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data HasyncForceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op for hasync_force: NITRO exposes no GET endpoint for this
	// action-only resource, so drift detection is impossible. Preserve state.
	tflog.Debug(ctx, "Read is a no-op for hasync_force; no GET endpoint on NITRO side")

	// Save (unchanged) data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HasyncForceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state HasyncForceResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for hasync_force; force/save are RequiresReplace, so
	// Terraform re-creates (re-Force-syncs) on change instead.
	tflog.Debug(ctx, "Update is a no-op for hasync_force; all attributes are RequiresReplace")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HasyncForceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data HasyncForceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete is a no-op for hasync_force: NITRO exposes no DELETE endpoint for
	// this action-only resource. Just remove from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for hasync_force; no DELETE endpoint on NITRO side")
	tflog.Trace(ctx, "Removed hasync_force from Terraform state")
}

func hasync_forceGetThePayloadFromthePlan(ctx context.Context, data *HasyncForceResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In hasync_forceGetThePayloadFromthePlan Function")

	// Build the Force action payload. force/save are included only when set.
	hasync := make(map[string]interface{})
	if !data.Force.IsNull() && !data.Force.IsUnknown() {
		hasync["force"] = data.Force.ValueBool()
	}
	if !data.Save.IsNull() && !data.Save.IsUnknown() {
		hasync["save"] = data.Save.ValueString()
	}

	return hasync
}
