package clustersync

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cluster"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// clustersync_force is an ACTION-ONLY, ZERO-ATTRIBUTE resource.
//
//   - NITRO exposes only the Force action:
//     POST /nitro/v1/config/clustersync?action=Force, which forces a
//     synchronization of the cluster configuration across nodes.
//   - There is NO add/set/get/delete endpoint, so:
//     Create performs the Force action, Read/Update are no-ops (there is nothing
//     to reconcile), and Delete is a state-only removal.
//   - Because there is no GET endpoint, there is NO datasource for clustersync_force.
var _ resource.Resource = &ClustersyncForceResource{}
var _ resource.ResourceWithConfigure = (*ClustersyncForceResource)(nil)

func NewClustersyncForceResource() resource.Resource {
	return &ClustersyncForceResource{}
}

// ClustersyncForceResource defines the resource implementation.
type ClustersyncForceResource struct {
	client *service.NitroClient
}

// ClustersyncForceResourceModel describes the resource data model.
//
// clustersync_force is a ZERO-ATTRIBUTE, ACTION-ONLY resource: the NITRO
// "clustersync" object exposes no read/write properties and only the Force action
// (POST /nitro/v1/config/clustersync?action=Force). The model therefore carries
// only the synthetic id.
type ClustersyncForceResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *ClustersyncForceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clustersync_force"
}

func (r *ClustersyncForceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClustersyncForceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the clustersync_force resource.",
			},
		},
	}
}

func (r *ClustersyncForceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClustersyncForceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating clustersync_force resource (Force action)")
	clustersync := clustersync_forceGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - NITRO exposes only POST ?action=Force
	err := r.client.ActOnResource(service.Clustersync.Type(), &clustersync, "Force")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to force clustersync, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Forced clustersync")

	// Synthetic ID - no GET endpoint exists to derive it from
	data.Id = types.StringValue("clustersync_force")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. clustersync_force has no GET endpoint; there is nothing to reconcile.
func (r *ClustersyncForceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClustersyncForceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for clustersync_force; NITRO exposes no GET endpoint (action=Force only)")

	// Preserve prior state unchanged - no GET endpoint to reconcile against
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. clustersync_force has no attributes and no set endpoint.
func (r *ClustersyncForceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ClustersyncForceResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for clustersync_force; it has no attributes and no set endpoint")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. clustersync_force has no delete endpoint; the action is not
// reversible and there is no persistent object to remove.
func (r *ClustersyncForceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for clustersync_force; NITRO has no delete endpoint")
}

// clustersync_forceGetThePayloadFromthePlan builds the (empty) NITRO payload for
// the Force action. clustersync_force has no read/write attributes, so the payload
// is an empty cluster.Clustersync struct.
func clustersync_forceGetThePayloadFromthePlan(ctx context.Context, data *ClustersyncForceResourceModel) cluster.Clustersync {
	tflog.Debug(ctx, "In clustersync_forceGetThePayloadFromthePlan Function")
	clustersync := cluster.Clustersync{}
	return clustersync
}
