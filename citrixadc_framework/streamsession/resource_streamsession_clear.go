package streamsession

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/stream"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &StreamsessionClearResource{}
var _ resource.ResourceWithConfigure = (*StreamsessionClearResource)(nil)
var _ resource.ResourceWithImportState = (*StreamsessionClearResource)(nil)

func NewStreamsessionClearResource() resource.Resource {
	return &StreamsessionClearResource{}
}

// StreamsessionClearResource defines the resource implementation.
type StreamsessionClearResource struct {
	client *service.NitroClient
}

// StreamsessionClearResourceModel describes the resource data model.
//
// This resource models the NITRO streamsession `?action=clear` action. clear is a
// one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops.
type StreamsessionClearResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (r *StreamsessionClearResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *StreamsessionClearResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_streamsession_clear"
}

func (r *StreamsessionClearResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *StreamsessionClearResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the streamsession_clear resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the stream identifier.",
			},
		},
	}
}

func (r *StreamsessionClearResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data StreamsessionClearResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating streamsession resource (clear action)")
	streamsession := streamsession_clearGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - NITRO exposes only POST ?action=clear
	err := r.client.ActOnResource(service.Streamsession.Type(), streamsession, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear streamsession, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared streamsession")

	// Synthetic ID - no GET endpoint exists to derive it from
	data.Id = types.StringValue("streamsession_clear")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StreamsessionClearResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// clear is a one-shot action. NITRO has no GET endpoint that reports
	// clear-state, so Read is a pure preserve-state no-op.
	var data StreamsessionClearResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for streamsession_clear; NITRO exposes no GET endpoint (action=clear only)")

	// Preserve prior state unchanged - no GET endpoint to reconcile against
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StreamsessionClearResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for clear; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state StreamsessionClearResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for streamsession_clear; the only attribute (name) is
	// RequiresReplace and NITRO exposes no update endpoint (action=clear only).
	tflog.Debug(ctx, "Update is a no-op for streamsession_clear")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StreamsessionClearResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// clear is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-clear"). Delete simply removes the resource from Terraform state.
	var data StreamsessionClearResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Action-only resource - no delete operation on NITRO (action=clear only), just remove from state
	tflog.Debug(ctx, "Delete is a no-op for streamsession_clear; removed from Terraform state")
}

func streamsession_clearGetThePayloadFromthePlan(ctx context.Context, data *StreamsessionClearResourceModel) stream.Streamsession {
	tflog.Debug(ctx, "In streamsession_clearGetThePayloadFromthePlan Function")

	// Create API request body from the model
	streamsession := stream.Streamsession{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		streamsession.Name = data.Name.ValueString()
	}

	return streamsession
}
