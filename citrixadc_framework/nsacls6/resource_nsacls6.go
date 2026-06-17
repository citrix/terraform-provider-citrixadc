package nsacls6

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Nsacls6Resource{}
var _ resource.ResourceWithConfigure = (*Nsacls6Resource)(nil)
var _ resource.ResourceWithImportState = (*Nsacls6Resource)(nil)

func NewNsacls6Resource() resource.Resource {
	return &Nsacls6Resource{}
}

// Nsacls6Resource defines the resource implementation.
type Nsacls6Resource struct {
	client *service.NitroClient
}

func (r *Nsacls6Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Nsacls6Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsacls6"
}

func (r *Nsacls6Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Nsacls6Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Nsacls6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Applying nsacls6 (action-only resource)")
	payload := nsacls6GetThePayloadFromthePlan(ctx, &data)

	// nsacls6 is an action-only resource (apply/clear/renumber, no add/get).
	// The Create maps to the "apply" action.
	err := r.client.ActOnResource(service.Nsacls6.Type(), payload, "apply")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to apply nsacls6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Applied nsacls6 resource")

	// Generate a synthetic ID; nsacls6 has no GET endpoint.
	data.Id = types.StringValue("nsacls6-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nsacls6Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Nsacls6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op: nsacls6 is an action-only resource with no GET endpoint.
	tflog.Debug(ctx, "Read is a no-op for nsacls6 (no GET endpoint); preserving state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nsacls6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state Nsacls6ResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op for nsacls6; all attributes are RequiresReplace.
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nsacls6; all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nsacls6Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Nsacls6ResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete is a no-op: nsacls6 is action-only with no inverse API; just remove from state.
	tflog.Debug(ctx, "Delete is a no-op for nsacls6; removing from Terraform state")
}
