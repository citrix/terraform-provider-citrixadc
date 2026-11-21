package sslservice

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslserviceResource{}
var _ resource.ResourceWithConfigure = (*SslserviceResource)(nil)
var _ resource.ResourceWithImportState = (*SslserviceResource)(nil)

func NewSslserviceResource() resource.Resource {
	return &SslserviceResource{}
}

// SslserviceResource defines the resource implementation.
type SslserviceResource struct {
	client *service.NitroClient
}

func (r *SslserviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslserviceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslservice"
}

func (r *SslserviceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslserviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslserviceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslservice resource")

	sslservice := sslserviceGetThePayloadFromtheConfig(ctx, &data)

	// sslservice is configured on an existing SSL service (settings-style resource):
	// mirror SDK v2 which uses UpdateUnnamedResource for the initial push.
	err := r.client.UpdateUnnamedResource(service.Sslservice.Type(), &sslservice)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslservice, got error: %s", err))
		return
	}

	// GH #1166: when nodefaultcipherbindings is set, unbind the appliance's built-in
	// default cipher group(s) (DEFAULT / DEFAULT_BACKEND) auto-bound to the SSL
	// service, so a custom cipher group can take effect alone. Applied at create time
	// only (the attribute is RequiresReplace).
	if !data.Nodefaultcipherbindings.IsNull() && data.Nodefaultcipherbindings.ValueBool() {
		if err := r.removeDefaultCipherBindings(ctx, data.Servicename.ValueString()); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to remove default cipher bindings for sslservice %s, got error: %s", data.Servicename.ValueString(), err))
			return
		}
	}

	// GH #1166: when nodefaultecccurvebindings is set, unbind the appliance's built-in
	// default ECC curve bindings auto-bound to the SSL service. Applied at create time
	// only (the attribute is RequiresReplace).
	if !data.Nodefaultecccurvebindings.IsNull() && data.Nodefaultecccurvebindings.ValueBool() {
		if err := r.removeDefaultEcccurveBindings(ctx, data.Servicename.ValueString()); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to remove default ECC curve bindings for sslservice %s, got error: %s", data.Servicename.ValueString(), err))
			return
		}
	}

	// ID is the plain servicename value (matches SDK v2 d.SetId(servicename))
	data.Id = types.StringValue(data.Servicename.ValueString())

	tflog.Trace(ctx, "Created sslservice resource")

	// Read the updated state back
	if !r.readSslserviceFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslservice not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslserviceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslservice resource")

	found := r.readSslserviceFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SslserviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SslserviceResourceModel

	// Read Terraform prior state to preserve the ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (unset candidates)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (servicename is RequiresReplace so it never changes here)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslservice resource")

	// Build a DELTA payload (only changed fields), mirroring SDK v2
	// updateSslserviceFunc which used per-field d.HasChange(). Sending the full
	// config on every update violates NITRO prerequisites such as
	// "sessTimeout requires sessReuse==ENABLED" (errorcode 1093). servicename
	// cannot change (RequiresReplace), so use the current name for the update.
	sslservice, hasChange := sslserviceGetThePayloadForUpdate(ctx, &data, &state)

	// Collect attributes that were removed from config so they can be reverted
	// to their NITRO defaults via the unset action (mirrors analyticsprofile).
	attributesToUnset := []string{}
	appendUnset := func(name string, plan, prior, cfg types.String) {
		if !plan.Equal(prior) && cfg.IsNull() {
			attributesToUnset = append(attributesToUnset, name)
		}
	}
	appendUnset("cipherredirect", data.Cipherredirect, state.Cipherredirect, config.Cipherredirect)
	appendUnset("clientauth", data.Clientauth, state.Clientauth, config.Clientauth)
	appendUnset("dh", data.Dh, state.Dh, config.Dh)
	appendUnset("dhkeyexpsizelimit", data.Dhkeyexpsizelimit, state.Dhkeyexpsizelimit, config.Dhkeyexpsizelimit)
	appendUnset("ersa", data.Ersa, state.Ersa, config.Ersa)
	appendUnset("redirectportrewrite", data.Redirectportrewrite, state.Redirectportrewrite, config.Redirectportrewrite)
	appendUnset("serverauth", data.Serverauth, state.Serverauth, config.Serverauth)
	appendUnset("sessreuse", data.Sessreuse, state.Sessreuse, config.Sessreuse)
	appendUnset("snienable", data.Snienable, state.Snienable, config.Snienable)
	appendUnset("ssl2", data.Ssl2, state.Ssl2, config.Ssl2)
	appendUnset("ssl3", data.Ssl3, state.Ssl3, config.Ssl3)
	appendUnset("sslredirect", data.Sslredirect, state.Sslredirect, config.Sslredirect)
	appendUnset("sslv2redirect", data.Sslv2redirect, state.Sslv2redirect, config.Sslv2redirect)
	appendUnset("tls1", data.Tls1, state.Tls1, config.Tls1)
	appendUnset("tls11", data.Tls11, state.Tls11, config.Tls11)
	appendUnset("tls12", data.Tls12, state.Tls12, config.Tls12)
	appendUnset("tls13", data.Tls13, state.Tls13, config.Tls13)

	sslserviceName := data.Servicename.ValueString()
	if hasChange {
		_, err := r.client.UpdateResource(service.Sslservice.Type(), sslserviceName, &sslservice)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslservice %s, got error: %s", sslserviceName, err))
			return
		}
	}

	// Unset attributes removed from config so the appliance reverts them to
	// their defaults. Done after the update so any value the update payload
	// carried for a removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"servicename": sslserviceName,
	}
	if err := utils.ExecuteUnset(r.client, service.Sslservice.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset sslservice attributes, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated sslservice resource")

	// Read the updated state back
	if !r.readSslserviceFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslservice not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslserviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslserviceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslservice resource")

	// sslservice has no NITRO delete operation - the SSL settings live on the parent
	// SSL service. Mirror SDK v2 which only clears the ID (state removal only).
	tflog.Trace(ctx, "Deleted sslservice resource from state")
}

// Helper function to read sslservice data from API.
// Returns false when the resource no longer exists on the ADC.
func (r *SslserviceResource) readSslserviceFromApi(ctx context.Context, data *SslserviceResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain servicename value.
	sslserviceName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Sslservice.Type(), sslserviceName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslservice, got error: %s", err))
		return false
	}

	sslserviceSetAttrFromGet(ctx, data, getResponseData)

	return true
}

// defaultCipherGroupNames are the appliance's built-in cipher groups that NetScaler
// auto-binds to an SSL service. nodefaultcipherbindings unbinds these.
var defaultCipherGroupNames = map[string]bool{
	"DEFAULT":         true,
	"DEFAULT_BACKEND": true,
}

// removeDefaultCipherBindings unbinds the built-in default cipher group(s)
// (DEFAULT / DEFAULT_BACKEND) from the SSL service (GH #1166). The bound ciphers are
// NOT reflected by the typed sslservice_sslcipher_binding GET on 14.1-73.x (it
// returns an empty body); they are only visible via the umbrella sslservice_binding
// endpoint under sslservice_sslciphersuite_binding[]. So enumerate via the umbrella
// and unbind each built-in default by ciphername. Only the names in
// defaultCipherGroupNames are touched, so a custom cipher bound earlier in the same
// apply is never removed.
func (r *SslserviceResource) removeDefaultCipherBindings(ctx context.Context, servicename string) error {
	umbrella, err := r.client.FindResourceArrayWithParams(service.FindParams{
		ResourceType:             service.Sslservice_binding.Type(),
		ResourceName:             servicename,
		ResourceMissingErrorCode: 258,
	})
	if err != nil {
		return err
	}
	for _, svc := range umbrella {
		raw, ok := svc["sslservice_sslciphersuite_binding"]
		if !ok {
			continue
		}
		rows, ok := raw.([]interface{})
		if !ok {
			continue
		}
		for _, cb := range rows {
			m, ok := cb.(map[string]interface{})
			if !ok {
				continue
			}
			name := ""
			if v, ok := m["ciphername"].(string); ok && v != "" {
				name = v
			} else if v, ok := m["cipheraliasname"].(string); ok && v != "" {
				name = v
			}
			if !defaultCipherGroupNames[name] {
				continue
			}
			argsMap := map[string]string{"ciphername": name}
			if err := r.client.DeleteResourceWithArgsMap(service.Sslservice_sslcipher_binding.Type(), servicename, argsMap); err != nil {
				return fmt.Errorf("unbind default cipher %q: %w", name, err)
			}
			tflog.Debug(ctx, fmt.Sprintf("Unbound default cipher %q from sslservice %s", name, servicename))
		}
	}
	return nil
}

// removeDefaultEcccurveBindings unbinds the appliance's built-in default ECC curve
// bindings from the SSL service (GH #1166). Unlike the cipher case, the default ECC
// curves ARE reflected by the typed sslservice_ecccurve_binding GET on 14.1-73.x, so
// enumerate it directly (mirrors sslprofile.deleteDefaultEcccurveBindings) and unbind
// each by ecccurvename. All currently-bound curves at create time are the appliance
// defaults.
func (r *SslserviceResource) removeDefaultEcccurveBindings(ctx context.Context, servicename string) error {
	bindings, err := r.client.FindResourceArray(service.Sslservice_ecccurve_binding.Type(), servicename)
	if err != nil {
		return err
	}
	for _, binding := range bindings {
		ecccurvename, ok := binding["ecccurvename"].(string)
		if !ok || ecccurvename == "" {
			continue
		}
		args := []string{fmt.Sprintf("ecccurvename:%s", ecccurvename)}
		if err := r.client.DeleteResourceWithArgs(service.Sslservice_ecccurve_binding.Type(), servicename, args); err != nil {
			return fmt.Errorf("unbind default ECC curve %q: %w", ecccurvename, err)
		}
		tflog.Debug(ctx, fmt.Sprintf("Unbound default ECC curve %q from sslservice %s", ecccurvename, servicename))
	}
	return nil
}
