package clusterfiles

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cluster"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ClusterfilesSyncerResource{}
var _ resource.ResourceWithConfigure = (*ClusterfilesSyncerResource)(nil)

func NewClusterfilesSyncerResource() resource.Resource {
	return &ClusterfilesSyncerResource{}
}

// ClusterfilesSyncerResource defines the resource implementation.
//
// This resource models the NITRO clusterfiles `?action=sync` action. sync is a
// one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. It mirrors the SDK v2 resource
// citrixadc_clusterfiles_syncer exactly: the sync payload carries the required
// `mode` list, and `timestamp` is a required, provider-supplied input whose
// value becomes the Terraform resource ID (SDK v2 does d.SetId(timestamp)).
// Both attributes are ForceNew in SDK v2 -> RequiresReplace here.
type ClusterfilesSyncerResource struct {
	client *service.NitroClient
}

// ClusterfilesSyncerResourceModel describes the resource data model.
type ClusterfilesSyncerResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Timestamp types.String `tfsdk:"timestamp"`
	Mode      types.Set    `tfsdk:"mode"`
}

func (r *ClusterfilesSyncerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusterfiles_syncer"
}

func (r *ClusterfilesSyncerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusterfilesSyncerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the clusterfiles_syncer resource. Set to the value of `timestamp`.",
			},
			// timestamp is a Required provider-supplied input in SDK v2 (not
			// generated/Computed). Its value is used as the resource ID via
			// d.SetId(timestamp). ForceNew -> RequiresReplace.
			"timestamp": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "A timestamp used to trigger and identify the sync operation. Used as the resource ID.",
			},
			// mode is the only field the NITRO sync action accepts. Required Set
			// of strings. ForceNew -> RequiresReplace.
			"mode": schema.SetAttribute{
				ElementType: types.StringType,
				Required:    true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.RequiresReplace(),
				},
				Description: "The directories and files to be synchronized. Possible values = all, bookmarks, ssl, imports, misc, dns, krb, AAA, app_catalog, all_plus_misc.",
			},
		},
	}
}

func (r *ClusterfilesSyncerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusterfilesSyncerResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Syncing clusterfiles (action-only resource)")
	payload := clusterfiles_syncerGetThePayloadFromthePlan(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// sync is a POST ?action=sync action. There is no add endpoint; the verb
	// casing is lower-case per the NITRO URL. Mirrors SDK v2 ActOnResource.
	err := r.client.ActOnResource(service.Clusterfiles.Type(), payload, "sync")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to sync clusterfiles, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Synced clusterfiles")

	// SDK v2 sets the ID to the supplied timestamp: d.SetId(timestamp).
	data.Id = types.StringValue(data.Timestamp.ValueString())

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterfilesSyncerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// sync is a one-shot action. NITRO has no GET endpoint that reports
	// sync-state, so Read is a pure preserve-state no-op (SDK v2 schema.Noop).
	var data ClusterfilesSyncerResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for clusterfiles_syncer; sync has no stable GET-backed object")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterfilesSyncerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for sync; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state ClusterfilesSyncerResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for clusterfiles_syncer; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterfilesSyncerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// sync is a one-shot side-effect action. There is no inverse NITRO API.
	// Delete simply removes the resource from Terraform state (SDK v2 schema.Noop).
	tflog.Debug(ctx, "Delete is a no-op for clusterfiles_syncer; NITRO has no inverse of the sync action")
}

// clusterfiles_syncerGetThePayloadFromthePlan builds the body for the sync
// action. The NITRO sync action accepts ONLY `mode`; `timestamp` is a
// provider-side value used for the resource ID and is intentionally excluded
// from the payload.
func clusterfiles_syncerGetThePayloadFromthePlan(ctx context.Context, data *ClusterfilesSyncerResourceModel, diags *diag.Diagnostics) cluster.Clusterfiles {
	tflog.Debug(ctx, "In clusterfiles_syncerGetThePayloadFromthePlan Function")

	clusterfiles := cluster.Clusterfiles{}
	if !data.Mode.IsNull() && !data.Mode.IsUnknown() {
		mode := make([]string, 0, len(data.Mode.Elements()))
		diags.Append(data.Mode.ElementsAs(ctx, &mode, false)...)
		clusterfiles.Mode = mode
	}

	return clusterfiles
}
