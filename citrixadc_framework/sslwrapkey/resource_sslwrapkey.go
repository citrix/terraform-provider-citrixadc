package sslwrapkey

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
var _ resource.Resource = &SslwrapkeyResource{}
var _ resource.ResourceWithConfigure = (*SslwrapkeyResource)(nil)
var _ resource.ResourceWithImportState = (*SslwrapkeyResource)(nil)
var _ resource.ResourceWithValidateConfig = (*SslwrapkeyResource)(nil)

func NewSslwrapkeyResource() resource.Resource {
	return &SslwrapkeyResource{}
}

// SslwrapkeyResource defines the resource implementation.
type SslwrapkeyResource struct {
	client *service.NitroClient
}

func (r *SslwrapkeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslwrapkeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslwrapkey"
}

func (r *SslwrapkeyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslwrapkeyResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data SslwrapkeyResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// password and salt are mandatory for the create action (Pattern 8 / Pattern 17).
	// Each is expanded into a secret triple whose value attributes are both Optional,
	// so enforce at-least-one-of at plan time.
	if data.Password.IsNull() && data.PasswordWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing Required Attribute",
			"Either \"password\" or \"password_wo\" must be specified.",
		)
	}
	if data.Salt.IsNull() && data.SaltWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("salt"),
			"Missing Required Attribute",
			"Either \"salt\" or \"salt_wo\" must be specified.",
		)
	}
}

func (r *SslwrapkeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config SslwrapkeyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslwrapkey resource")
	// Get payload from plan (regular attributes)
	sslwrapkey := sslwrapkeyGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	sslwrapkeyGetThePayloadFromtheConfig(ctx, &config, &sslwrapkey)

	// Make API call
	// sslwrapkey is created via the NITRO `create` action (?action=create, POST),
	// NOT a plain add. There is no update endpoint.
	// NOTE: may require the FIPS/crypto subsystem to be available - gate the
	// acceptance test accordingly.
	err := r.client.ActOnResource(service.Sslwrapkey.Type(), &sslwrapkey, "create")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslwrapkey, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslwrapkey resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(data.Wrapkeyname.ValueString())

	// Read the updated state back
	r.readSslwrapkeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslwrapkeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslwrapkeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslwrapkey resource")

	r.readSslwrapkeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslwrapkeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslwrapkeyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for sslwrapkey; the NITRO doc exposes no update endpoint and
	// all attributes are RequiresReplace (Pattern 5).
	tflog.Debug(ctx, "Update is a no-op for sslwrapkey; all attributes are RequiresReplace")

	// Read the updated state back
	r.readSslwrapkeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslwrapkeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslwrapkeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslwrapkey resource")
	// NITRO supports DELETE /config/sslwrapkey/<wrapkeyname> (Pattern 4).
	err := r.client.DeleteResource(service.Sslwrapkey.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslwrapkey, got error: %s", err))
		return
	}
	tflog.Trace(ctx, "Deleted sslwrapkey resource")
}

// Helper function to read sslwrapkey data from API
func (r *SslwrapkeyResource) readSslwrapkeyFromApi(ctx context.Context, data *SslwrapkeyResourceModel, diags *diag.Diagnostics) {

	// Named resource: read by wrapkeyname (the ID holds the plain key value).
	getResponseData, err := r.client.FindResource(service.Sslwrapkey.Type(), data.Id.ValueString())
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslwrapkey, got error: %s", err))
		return
	}

	sslwrapkeySetAttrFromGet(ctx, data, getResponseData)
}
