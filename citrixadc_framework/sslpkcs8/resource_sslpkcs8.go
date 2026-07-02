package sslpkcs8

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
var _ resource.Resource = &Sslpkcs8Resource{}
var _ resource.ResourceWithConfigure = (*Sslpkcs8Resource)(nil)
var _ resource.ResourceWithImportState = (*Sslpkcs8Resource)(nil)

func NewSslpkcs8Resource() resource.Resource {
	return &Sslpkcs8Resource{}
}

// Sslpkcs8Resource defines the resource implementation.
type Sslpkcs8Resource struct {
	client *service.NitroClient
}

func (r *Sslpkcs8Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Sslpkcs8Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslpkcs8"
}

func (r *Sslpkcs8Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Sslpkcs8Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config Sslpkcs8ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslpkcs8 resource")
	// Get payload from plan (regular attributes)
	sslpkcs8 := sslpkcs8GetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	sslpkcs8GetThePayloadFromtheConfig(ctx, &config, &sslpkcs8)

	// Make API call
	// Action-only resource: sslpkcs8 exposes ONLY ?action=convert (no add/get/delete).
	// NOTE: this operation requires the source key file to be present on the appliance
	// filesystem and is NON-IDEMPOTENT - it writes the converted PKCS#8 output file.
	// There is no GET endpoint to read back the result.
	err := r.client.ActOnResource(service.Sslpkcs8.Type(), &sslpkcs8, "convert")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to convert sslpkcs8, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Converted sslpkcs8 resource")

	// Synthetic ID = pkcs8file value (there is no GET endpoint to derive an ID from).
	data.Id = types.StringValue(data.Pkcs8file.ValueString())

	// No Read-back: sslpkcs8 has no GET endpoint.

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Sslpkcs8Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Sslpkcs8ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op for sslpkcs8: NITRO exposes no GET endpoint for this
	// action-only resource. Preserve the prior state unchanged.
	tflog.Debug(ctx, "Read is a no-op for sslpkcs8; NITRO has no GET endpoint")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Sslpkcs8Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state Sslpkcs8ResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for sslpkcs8: NITRO exposes only ?action=convert and no
	// GET endpoint; every schema attribute is RequiresReplace, so changes force
	// recreation and Update is never reached for an attribute change.
	tflog.Debug(ctx, "Update is a no-op for sslpkcs8; all attributes are RequiresReplace")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Sslpkcs8Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Sslpkcs8ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslpkcs8 resource")
	// Action-only resource: ?action=convert has no inverse API. The converted
	// output file persists on the appliance; Delete only removes Terraform state.
	tflog.Trace(ctx, "Removed sslpkcs8 from Terraform state")
}
