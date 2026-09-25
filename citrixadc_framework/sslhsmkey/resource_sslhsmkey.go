package sslhsmkey

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslhsmkeyResource{}
var _ resource.ResourceWithUpgradeState = &SslhsmkeyResource{}
var _ resource.ResourceWithConfigure = (*SslhsmkeyResource)(nil)
var _ resource.ResourceWithImportState = (*SslhsmkeyResource)(nil)
var _ resource.ResourceWithValidateConfig = (*SslhsmkeyResource)(nil)

func NewSslhsmkeyResource() resource.Resource {
	return &SslhsmkeyResource{}
}

// SslhsmkeyResource defines the resource implementation.
type SslhsmkeyResource struct {
	client *service.NitroClient
}

func (r *SslhsmkeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslhsmkeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslhsmkey"
}

func (r *SslhsmkeyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// ValidateConfig enforces that the SafeNet HSM partition password is supplied via
// exactly one of the persisted "password" or the write-only "password_wo".
//
// The partition password applies only to SafeNet HSM (NITRO: "Applies only to
// SafeNet HSM"), so the requirement is scoped to hsmtype == SAFENET; THALES and
// KEYVAULT keys need no password and are left unconstrained. Prefer "password_wo":
// it is ephemeral (never written to state), so it is not available to the delete
// path and cannot leak the PIN into the delete-args URL or logs (CTXMYT-1123).
func (r *SslhsmkeyResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data SslhsmkeyResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Defer validation while any relevant value is unknown (e.g. sourced from
	// another resource); Terraform re-runs this once the values are known.
	if data.Hsmtype.IsUnknown() || data.Password.IsUnknown() || data.PasswordWo.IsUnknown() {
		return
	}

	// password/password_wo are meaningful only for SafeNet HSM. hsmtype is null in
	// config when unset (its THALES default is applied later), so only SAFENET matches.
	if !strings.EqualFold(data.Hsmtype.ValueString(), "SAFENET") {
		return
	}

	pwSet := !data.Password.IsNull() && data.Password.ValueString() != ""
	woSet := !data.PasswordWo.IsNull() && data.PasswordWo.ValueString() != ""

	switch {
	case !pwSet && !woSet:
		resp.Diagnostics.AddAttributeError(
			path.Root("password_wo"),
			"Missing Required Attribute",
			"A SafeNet HSM key requires a partition password: set exactly one of \"password\" or \"password_wo\". "+
				"Prefer \"password_wo\" (write-only) so the PIN is never persisted in state or sent in the delete-args URL.",
		)
	case pwSet && woSet:
		resp.Diagnostics.AddAttributeError(
			path.Root("password_wo"),
			"Conflicting Attributes",
			"Set only one of \"password\" or \"password_wo\" for sslhsmkey, not both.",
		)
	}
}

func (r *SslhsmkeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config SslhsmkeyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslhsmkey resource")
	// Get payload from plan (regular attributes)
	sslhsmkey := sslhsmkeyGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	sslhsmkeyGetThePayloadFromtheConfig(ctx, &config, &sslhsmkey)

	// Make API call
	// Named resource - use AddResource
	hsmkeyname_value := data.Hsmkeyname.ValueString()
	_, err := r.client.AddResource(service.Sslhsmkey.Type(), hsmkeyname_value, &sslhsmkey)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslhsmkey, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslhsmkey resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Hsmkeyname.ValueString()))

	// Read the updated state back
	if !r.readSslhsmkeyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslhsmkey not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslhsmkeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslhsmkeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslhsmkey resource")

	found := r.readSslhsmkeyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslhsmkeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslhsmkeyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslhsmkey resource")

	// No in-place update: every settable attribute of sslhsmkey is ForceNew
	// (RequiresReplace) — including the write-only password and its version tracker —
	// and NITRO exposes no "update" operation (only add/delete/get). Any change (e.g.
	// rotating password) triggers a destroy+recreate, so this method is never invoked
	// with a real diff. The prior UpdateResource(PUT) call was dead code (a PUT would
	// be rejected) and has been removed along with the now-unused config read. If a
	// future schema change makes an attribute non-ForceNew, add a proper update path
	// here.

	// Read the updated state back
	if !r.readSslhsmkeyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslhsmkey not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslhsmkeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslhsmkeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslhsmkey resource")
	// Named resource - delete using DeleteResourceWithArgsMap.
	// sslhsmkey delete requires the HSM identification args (hsmtype/key/keystore/
	// password/serialnum), matching the SDKv2 behavior. A bare DeleteResource omits
	// them, which can cause the delete to fail or misidentify the key.
	hsmkeyname_value := data.Hsmkeyname.ValueString()
	// URL-encode the delete-arg VALUES so a value containing args/URL delimiters
	// (, : ? & # =) or spaces — e.g. an HSM partition password — does not corrupt
	// the request. Escaping is done here at the resource layer (as the ~197 binding
	// resources do); the generic delete-with-args client must NOT also escape, or
	// values would be double-encoded.
	argsMap := make(map[string]string)
	if !data.Hsmtype.IsNull() && data.Hsmtype.ValueString() != "" {
		argsMap["hsmtype"] = url.QueryEscape(data.Hsmtype.ValueString())
	}
	if !data.Key.IsNull() && data.Key.ValueString() != "" {
		argsMap["key"] = url.QueryEscape(data.Key.ValueString())
	}
	if !data.Keystore.IsNull() && data.Keystore.ValueString() != "" {
		argsMap["keystore"] = url.QueryEscape(data.Keystore.ValueString())
	}
	if !data.Password.IsNull() && data.Password.ValueString() != "" {
		argsMap["password"] = url.QueryEscape(data.Password.ValueString())
	}
	if !data.Serialnum.IsNull() && data.Serialnum.ValueString() != "" {
		argsMap["serialnum"] = url.QueryEscape(data.Serialnum.ValueString())
	}
	err := r.client.DeleteResourceWithArgsMap(service.Sslhsmkey.Type(), hsmkeyname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslhsmkey, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslhsmkey resource")
}

// Helper function to read sslhsmkey data from API
func (r *SslhsmkeyResource) readSslhsmkeyFromApi(ctx context.Context, data *SslhsmkeyResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	hsmkeyname_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Sslhsmkey.Type(), hsmkeyname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslhsmkey, got error: %s", err))
		return false
	}

	sslhsmkeySetAttrFromGet(ctx, data, getResponseData)

	return true
}

// UpgradeState migrates pre-write-only state (GH #1441): it seeds the
// "*_wo_version" tracker attribute(s) to 1 when the stored state has no value
// for them, so the schema Default does not plan a spurious "null -> 1" update
// after upgrading the provider. Paired with the schema Version bump so the
// upgrade path actually runs. See utils.WoVersionUpgradeState.
func (r *SslhsmkeyResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	schemaResp := resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)
	return utils.WoVersionUpgradeState(schemaResp.Schema, func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
		var data SslhsmkeyResourceModel
		resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if data.PasswordWoVersion.IsNull() {
			data.PasswordWoVersion = types.Int64Value(1)
		}
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	})
}
