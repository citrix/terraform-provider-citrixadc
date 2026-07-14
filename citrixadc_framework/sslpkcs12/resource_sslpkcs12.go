package sslpkcs12

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
var _ resource.Resource = &Sslpkcs12Resource{}
var _ resource.ResourceWithConfigure = (*Sslpkcs12Resource)(nil)
var _ resource.ResourceWithImportState = (*Sslpkcs12Resource)(nil)
var _ resource.ResourceWithValidateConfig = (*Sslpkcs12Resource)(nil)

func NewSslpkcs12Resource() resource.Resource {
	return &Sslpkcs12Resource{}
}

// Sslpkcs12Resource defines the resource implementation.
type Sslpkcs12Resource struct {
	client *service.NitroClient
}

// ValidateConfig enforces that the mandatory secret attribute password is
// supplied via one of its value/write-only variants (Pattern 17). password
// protects the PKCS#12 material and is required for the convert action.
//
// pempassphrase is intentionally NOT required: it is only needed to unlock an
// ENCRYPTED input PEM key when exporting. Exporting an unencrypted key needs no
// pass phrase, so requiring it unconditionally would reject a valid config. When
// it is actually needed but missing/wrong, the NITRO convert call returns a
// clear error.
func (r *Sslpkcs12Resource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data Sslpkcs12ResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Password.IsNull() && data.PasswordWo.IsNull() {
		resp.Diagnostics.AddError(
			"Missing required attribute",
			"One of \"password\" or \"password_wo\" must be set for sslpkcs12.",
		)
	}
}

func (r *Sslpkcs12Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Sslpkcs12Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslpkcs12"
}

func (r *Sslpkcs12Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Sslpkcs12Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config Sslpkcs12ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslpkcs12 resource")
	// Get payload from plan (regular attributes)
	sslpkcs12 := sslpkcs12GetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	sslpkcs12GetThePayloadFromtheConfig(ctx, &config, &sslpkcs12)

	// Make API call
	// Action-only resource: sslpkcs12 exposes ONLY ?action=convert (no add/get/delete).
	// NOTE: this operation is DISRUPTIVE and NON-IDEMPOTENT - it requires the source
	// certificate/key files to be present on the appliance filesystem and writes the
	// converted output file. There is no GET endpoint to read back the result.
	err := r.client.ActOnResource(service.Sslpkcs12.Type(), &sslpkcs12, "convert")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to convert sslpkcs12, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Converted sslpkcs12 resource")

	// Synthetic ID = outfile value (there is no GET endpoint to derive an ID from).
	data.Id = types.StringValue(data.Outfile.ValueString())

	// No Read-back: sslpkcs12 has no GET endpoint.

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Sslpkcs12Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Sslpkcs12ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op for sslpkcs12: NITRO exposes no GET endpoint for this
	// action-only resource. Preserve the prior state unchanged.
	tflog.Debug(ctx, "Read is a no-op for sslpkcs12; NITRO has no GET endpoint")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Sslpkcs12Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state Sslpkcs12ResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for sslpkcs12: NITRO exposes only ?action=convert and no
	// GET endpoint; every schema attribute is RequiresReplace, so changes force
	// recreation and Update is never reached for an attribute change.
	tflog.Debug(ctx, "Update is a no-op for sslpkcs12; all attributes are RequiresReplace")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Sslpkcs12Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Sslpkcs12ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslpkcs12 resource")
	// Action-only resource: ?action=convert has no inverse API. The converted
	// output file persists on the appliance; Delete only removes Terraform state.
	tflog.Trace(ctx, "Removed sslpkcs12 from Terraform state")
}
