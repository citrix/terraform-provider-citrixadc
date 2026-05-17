package sslprofile

import (
	"context"
	"fmt"
	"strconv"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslprofileResource{}
var _ resource.ResourceWithConfigure = (*SslprofileResource)(nil)
var _ resource.ResourceWithImportState = (*SslprofileResource)(nil)

func NewSslprofileResource() resource.Resource {
	return &SslprofileResource{}
}

// SslprofileResource defines the resource implementation.
type SslprofileResource struct {
	client *service.NitroClient
}

func (r *SslprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslprofile"
}

func (r *SslprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config SslprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslprofile resource")
	// Get payload from plan (regular attributes)
	sslprofile := sslprofileGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	sslprofileGetThePayloadFromtheConfig(ctx, &config, &sslprofile)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Sslprofile.Type(), name_value, &sslprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Handle default binding deletions
	if !data.Nodefaultecccurvebindings.IsNull() && data.Nodefaultecccurvebindings.ValueBool() {
		if err := r.deleteDefaultEcccurveBindings(ctx, data.Name.ValueString()); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete default ECC curve bindings: %s", err))
			return
		}
	}
	if !data.Nodefaultcipherbindings.IsNull() && data.Nodefaultcipherbindings.ValueBool() {
		if err := r.deleteDefaultCipherBindings(ctx, data.Name.ValueString()); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete default cipher bindings: %s", err))
			return
		}
	}

	// Handle ECC curve bindings
	if !data.Ecccurvebindings.IsNull() {
		if err := r.createEcccurveBindings(ctx, &data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ECC curve bindings: %s", err))
			return
		}
	}

	// Handle cipher bindings
	if !data.Cipherbindings.IsNull() {
		if err := r.createCipherBindings(ctx, &data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cipher bindings: %s", err))
			return
		}
	}

	// Read the updated state back
	r.readSslprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslprofile resource")

	r.readSslprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SslprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Allowextendedmastersecret.Equal(state.Allowextendedmastersecret) {
		tflog.Debug(ctx, fmt.Sprintf("allowextendedmastersecret has changed for sslprofile"))
		hasChange = true
	}
	if !data.Allowunknownsni.Equal(state.Allowunknownsni) {
		tflog.Debug(ctx, fmt.Sprintf("allowunknownsni has changed for sslprofile"))
		hasChange = true
	}
	if !data.Alpnprotocol.Equal(state.Alpnprotocol) {
		tflog.Debug(ctx, fmt.Sprintf("alpnprotocol has changed for sslprofile"))
		hasChange = true
	}
	if !data.Ciphername.Equal(state.Ciphername) {
		tflog.Debug(ctx, fmt.Sprintf("ciphername has changed for sslprofile"))
		hasChange = true
	}
	if !data.Cipherpriority.Equal(state.Cipherpriority) {
		tflog.Debug(ctx, fmt.Sprintf("cipherpriority has changed for sslprofile"))
		hasChange = true
	}
	if !data.Cipherredirect.Equal(state.Cipherredirect) {
		tflog.Debug(ctx, fmt.Sprintf("cipherredirect has changed for sslprofile"))
		hasChange = true
	}
	if !data.Cipherurl.Equal(state.Cipherurl) {
		tflog.Debug(ctx, fmt.Sprintf("cipherurl has changed for sslprofile"))
		hasChange = true
	}
	if !data.Cleartextport.Equal(state.Cleartextport) {
		tflog.Debug(ctx, fmt.Sprintf("cleartextport has changed for sslprofile"))
		hasChange = true
	}
	if !data.Clientauth.Equal(state.Clientauth) {
		tflog.Debug(ctx, fmt.Sprintf("clientauth has changed for sslprofile"))
		hasChange = true
	}
	if !data.Clientauthuseboundcachain.Equal(state.Clientauthuseboundcachain) {
		tflog.Debug(ctx, fmt.Sprintf("clientauthuseboundcachain has changed for sslprofile"))
		hasChange = true
	}
	if !data.Clientcert.Equal(state.Clientcert) {
		tflog.Debug(ctx, fmt.Sprintf("clientcert has changed for sslprofile"))
		hasChange = true
	}
	if !data.Commonname.Equal(state.Commonname) {
		tflog.Debug(ctx, fmt.Sprintf("commonname has changed for sslprofile"))
		hasChange = true
	}
	if !data.Defaultsni.Equal(state.Defaultsni) {
		tflog.Debug(ctx, fmt.Sprintf("defaultsni has changed for sslprofile"))
		hasChange = true
	}
	if !data.Denysslreneg.Equal(state.Denysslreneg) {
		tflog.Debug(ctx, fmt.Sprintf("denysslreneg has changed for sslprofile"))
		hasChange = true
	}
	if !data.Dh.Equal(state.Dh) {
		tflog.Debug(ctx, fmt.Sprintf("dh has changed for sslprofile"))
		hasChange = true
	}
	if !data.Dhcount.Equal(state.Dhcount) {
		tflog.Debug(ctx, fmt.Sprintf("dhcount has changed for sslprofile"))
		hasChange = true
	}
	if !data.Dhekeyexchangewithpsk.Equal(state.Dhekeyexchangewithpsk) {
		tflog.Debug(ctx, fmt.Sprintf("dhekeyexchangewithpsk has changed for sslprofile"))
		hasChange = true
	}
	if !data.Dhfile.Equal(state.Dhfile) {
		tflog.Debug(ctx, fmt.Sprintf("dhfile has changed for sslprofile"))
		hasChange = true
	}
	if !data.Dhkeyexpsizelimit.Equal(state.Dhkeyexpsizelimit) {
		tflog.Debug(ctx, fmt.Sprintf("dhkeyexpsizelimit has changed for sslprofile"))
		hasChange = true
	}
	if !data.Dropreqwithnohostheader.Equal(state.Dropreqwithnohostheader) {
		tflog.Debug(ctx, fmt.Sprintf("dropreqwithnohostheader has changed for sslprofile"))
		hasChange = true
	}
	if !data.Encryptedclienthello.Equal(state.Encryptedclienthello) {
		tflog.Debug(ctx, fmt.Sprintf("encryptedclienthello has changed for sslprofile"))
		hasChange = true
	}
	if !data.Encrypttriggerpktcount.Equal(state.Encrypttriggerpktcount) {
		tflog.Debug(ctx, fmt.Sprintf("encrypttriggerpktcount has changed for sslprofile"))
		hasChange = true
	}
	if !data.Ersa.Equal(state.Ersa) {
		tflog.Debug(ctx, fmt.Sprintf("ersa has changed for sslprofile"))
		hasChange = true
	}
	if !data.Ersacount.Equal(state.Ersacount) {
		tflog.Debug(ctx, fmt.Sprintf("ersacount has changed for sslprofile"))
		hasChange = true
	}
	if !data.Hsts.Equal(state.Hsts) {
		tflog.Debug(ctx, fmt.Sprintf("hsts has changed for sslprofile"))
		hasChange = true
	}
	if !data.Includesubdomains.Equal(state.Includesubdomains) {
		tflog.Debug(ctx, fmt.Sprintf("includesubdomains has changed for sslprofile"))
		hasChange = true
	}
	if !data.Insertionencoding.Equal(state.Insertionencoding) {
		tflog.Debug(ctx, fmt.Sprintf("insertionencoding has changed for sslprofile"))
		hasChange = true
	}
	if !data.Maxage.Equal(state.Maxage) {
		tflog.Debug(ctx, fmt.Sprintf("maxage has changed for sslprofile"))
		hasChange = true
	}
	if !data.Maxrenegrate.Equal(state.Maxrenegrate) {
		tflog.Debug(ctx, fmt.Sprintf("maxrenegrate has changed for sslprofile"))
		hasChange = true
	}
	if !data.Nodefaultbindings.Equal(state.Nodefaultbindings) {
		tflog.Debug(ctx, fmt.Sprintf("nodefaultbindings has changed for sslprofile"))
		hasChange = true
	}
	if !data.Ocspstapling.Equal(state.Ocspstapling) {
		tflog.Debug(ctx, fmt.Sprintf("ocspstapling has changed for sslprofile"))
		hasChange = true
	}
	if !data.Preload.Equal(state.Preload) {
		tflog.Debug(ctx, fmt.Sprintf("preload has changed for sslprofile"))
		hasChange = true
	}
	if !data.Prevsessionkeylifetime.Equal(state.Prevsessionkeylifetime) {
		tflog.Debug(ctx, fmt.Sprintf("prevsessionkeylifetime has changed for sslprofile"))
		hasChange = true
	}
	if !data.Pushenctrigger.Equal(state.Pushenctrigger) {
		tflog.Debug(ctx, fmt.Sprintf("pushenctrigger has changed for sslprofile"))
		hasChange = true
	}
	if !data.Pushenctriggertimeout.Equal(state.Pushenctriggertimeout) {
		tflog.Debug(ctx, fmt.Sprintf("pushenctriggertimeout has changed for sslprofile"))
		hasChange = true
	}
	if !data.Pushflag.Equal(state.Pushflag) {
		tflog.Debug(ctx, fmt.Sprintf("pushflag has changed for sslprofile"))
		hasChange = true
	}
	if !data.Quantumsize.Equal(state.Quantumsize) {
		tflog.Debug(ctx, fmt.Sprintf("quantumsize has changed for sslprofile"))
		hasChange = true
	}
	if !data.Redirectportrewrite.Equal(state.Redirectportrewrite) {
		tflog.Debug(ctx, fmt.Sprintf("redirectportrewrite has changed for sslprofile"))
		hasChange = true
	}
	if !data.Sendclosenotify.Equal(state.Sendclosenotify) {
		tflog.Debug(ctx, fmt.Sprintf("sendclosenotify has changed for sslprofile"))
		hasChange = true
	}
	if !data.Serverauth.Equal(state.Serverauth) {
		tflog.Debug(ctx, fmt.Sprintf("serverauth has changed for sslprofile"))
		hasChange = true
	}
	if !data.Sessionkeylifetime.Equal(state.Sessionkeylifetime) {
		tflog.Debug(ctx, fmt.Sprintf("sessionkeylifetime has changed for sslprofile"))
		hasChange = true
	}
	if !data.Sessionticket.Equal(state.Sessionticket) {
		tflog.Debug(ctx, fmt.Sprintf("sessionticket has changed for sslprofile"))
		hasChange = true
	}
	// Check secret attribute sessionticketkeydata or its version tracker
	if !data.Sessionticketkeydata.Equal(state.Sessionticketkeydata) {
		tflog.Debug(ctx, fmt.Sprintf("sessionticketkeydata has changed for sslprofile"))
		hasChange = true
	} else if !data.SessionticketkeydataWoVersion.Equal(state.SessionticketkeydataWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("sessionticketkeydata_wo_version has changed for sslprofile"))
		hasChange = true
	}
	if !data.Sessionticketkeyrefresh.Equal(state.Sessionticketkeyrefresh) {
		tflog.Debug(ctx, fmt.Sprintf("sessionticketkeyrefresh has changed for sslprofile"))
		hasChange = true
	}
	if !data.Sessionticketlifetime.Equal(state.Sessionticketlifetime) {
		tflog.Debug(ctx, fmt.Sprintf("sessionticketlifetime has changed for sslprofile"))
		hasChange = true
	}
	if !data.Sessreuse.Equal(state.Sessreuse) {
		tflog.Debug(ctx, fmt.Sprintf("sessreuse has changed for sslprofile"))
		hasChange = true
	}
	if !data.Sesstimeout.Equal(state.Sesstimeout) {
		tflog.Debug(ctx, fmt.Sprintf("sesstimeout has changed for sslprofile"))
		hasChange = true
	}
	if !data.Skipclientcertpolicycheck.Equal(state.Skipclientcertpolicycheck) {
		tflog.Debug(ctx, fmt.Sprintf("skipclientcertpolicycheck has changed for sslprofile"))
		hasChange = true
	}
	if !data.Snienable.Equal(state.Snienable) {
		tflog.Debug(ctx, fmt.Sprintf("snienable has changed for sslprofile"))
		hasChange = true
	}
	if !data.Snihttphostmatch.Equal(state.Snihttphostmatch) {
		tflog.Debug(ctx, fmt.Sprintf("snihttphostmatch has changed for sslprofile"))
		hasChange = true
	}
	if !data.Ssl3.Equal(state.Ssl3) {
		tflog.Debug(ctx, fmt.Sprintf("ssl3 has changed for sslprofile"))
		hasChange = true
	}
	if !data.Sslclientlogs.Equal(state.Sslclientlogs) {
		tflog.Debug(ctx, fmt.Sprintf("sslclientlogs has changed for sslprofile"))
		hasChange = true
	}
	if !data.Sslimaxsessperserver.Equal(state.Sslimaxsessperserver) {
		tflog.Debug(ctx, fmt.Sprintf("sslimaxsessperserver has changed for sslprofile"))
		hasChange = true
	}
	if !data.Sslinterception.Equal(state.Sslinterception) {
		tflog.Debug(ctx, fmt.Sprintf("sslinterception has changed for sslprofile"))
		hasChange = true
	}
	if !data.Ssliocspcheck.Equal(state.Ssliocspcheck) {
		tflog.Debug(ctx, fmt.Sprintf("ssliocspcheck has changed for sslprofile"))
		hasChange = true
	}
	if !data.Sslireneg.Equal(state.Sslireneg) {
		tflog.Debug(ctx, fmt.Sprintf("sslireneg has changed for sslprofile"))
		hasChange = true
	}
	if !data.Ssllogprofile.Equal(state.Ssllogprofile) {
		tflog.Debug(ctx, fmt.Sprintf("ssllogprofile has changed for sslprofile"))
		hasChange = true
	}
	if !data.Sslredirect.Equal(state.Sslredirect) {
		tflog.Debug(ctx, fmt.Sprintf("sslredirect has changed for sslprofile"))
		hasChange = true
	}
	if !data.Ssltriggertimeout.Equal(state.Ssltriggertimeout) {
		tflog.Debug(ctx, fmt.Sprintf("ssltriggertimeout has changed for sslprofile"))
		hasChange = true
	}
	if !data.Strictcachecks.Equal(state.Strictcachecks) {
		tflog.Debug(ctx, fmt.Sprintf("strictcachecks has changed for sslprofile"))
		hasChange = true
	}
	if !data.Strictsigdigestcheck.Equal(state.Strictsigdigestcheck) {
		tflog.Debug(ctx, fmt.Sprintf("strictsigdigestcheck has changed for sslprofile"))
		hasChange = true
	}
	if !data.Tls1.Equal(state.Tls1) {
		tflog.Debug(ctx, fmt.Sprintf("tls1 has changed for sslprofile"))
		hasChange = true
	}
	if !data.Tls11.Equal(state.Tls11) {
		tflog.Debug(ctx, fmt.Sprintf("tls11 has changed for sslprofile"))
		hasChange = true
	}
	if !data.Tls12.Equal(state.Tls12) {
		tflog.Debug(ctx, fmt.Sprintf("tls12 has changed for sslprofile"))
		hasChange = true
	}
	if !data.Tls13.Equal(state.Tls13) {
		tflog.Debug(ctx, fmt.Sprintf("tls13 has changed for sslprofile"))
		hasChange = true
	}
	if !data.Tls13sessionticketsperauthcontext.Equal(state.Tls13sessionticketsperauthcontext) {
		tflog.Debug(ctx, fmt.Sprintf("tls13sessionticketsperauthcontext has changed for sslprofile"))
		hasChange = true
	}
	if !data.Zerorttearlydata.Equal(state.Zerorttearlydata) {
		tflog.Debug(ctx, fmt.Sprintf("zerorttearlydata has changed for sslprofile"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		sslprofile := sslprofileGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		sslprofileGetThePayloadFromtheConfig(ctx, &config, &sslprofile)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Sslprofile.Type(), name_value, &sslprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslprofile resource, skipping update")
	}

	// Handle ECC curve binding changes
	if !data.Ecccurvebindings.Equal(state.Ecccurvebindings) {
		if err := r.updateEcccurveBindings(ctx, &state, &data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ECC curve bindings: %s", err))
			return
		}
	}

	// Handle cipher binding changes
	if !data.Cipherbindings.Equal(state.Cipherbindings) {
		if err := r.updateCipherBindings(ctx, &state, &data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cipher bindings: %s", err))
			return
		}
	}

	// Read the updated state back
	r.readSslprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslprofile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Sslprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslprofile resource")
}

// Helper function to read sslprofile data from API
func (r *SslprofileResource) readSslprofileFromApi(ctx context.Context, data *SslprofileResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Sslprofile.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslprofile, got error: %s", err))
		return
	}

	sslprofileSetAttrFromGet(ctx, data, getResponseData)

	// Read ECC curve bindings only if configured
	if !data.Ecccurvebindings.IsNull() {
		r.readEcccurveBindings(ctx, data, diags)
	}

	// Read cipher bindings only if configured
	if !data.Cipherbindings.IsNull() {
		r.readCipherBindings(ctx, data, diags)
	}

}

// ECC curve binding helpers

func (r *SslprofileResource) deleteDefaultEcccurveBindings(ctx context.Context, profileName string) error {
	tflog.Debug(ctx, "Deleting default ECC curve bindings")
	bindings, _ := r.client.FindResourceArray(service.Sslprofile_ecccurve_binding.Type(), profileName)
	for _, binding := range bindings {
		ecccurvename, ok := binding["ecccurvename"].(string)
		if !ok {
			continue
		}
		args := []string{fmt.Sprintf("ecccurvename:%s", ecccurvename)}
		if err := r.client.DeleteResourceWithArgs(service.Sslprofile_ecccurve_binding.Type(), profileName, args); err != nil {
			tflog.Warn(ctx, fmt.Sprintf("Error deleting default ECC curve binding %s: %s", ecccurvename, err))
		}
	}
	return nil
}

func (r *SslprofileResource) createEcccurveBindings(ctx context.Context, data *SslprofileResourceModel) error {
	tflog.Debug(ctx, "Creating ECC curve bindings")
	profileName := data.Name.ValueString()

	// Delete default ECC curve bindings first
	defaultBindings, _ := r.client.FindResourceArray(service.Sslprofile_ecccurve_binding.Type(), profileName)
	for _, binding := range defaultBindings {
		ecccurvename, ok := binding["ecccurvename"].(string)
		if !ok {
			continue
		}
		args := []string{fmt.Sprintf("ecccurvename:%s", ecccurvename)}
		r.client.DeleteResourceWithArgs(service.Sslprofile_ecccurve_binding.Type(), profileName, args)
	}

	// Add configured bindings
	var ecccurves []string
	data.Ecccurvebindings.ElementsAs(context.Background(), &ecccurves, false)
	for _, ecccurvename := range ecccurves {
		bindingStruct := ssl.Sslprofileecccurvebinding{
			Name:         profileName,
			Ecccurvename: ecccurvename,
		}
		if _, err := r.client.UpdateResource(service.Sslprofile_ecccurve_binding.Type(), profileName, bindingStruct); err != nil {
			return err
		}
	}
	return nil
}

func (r *SslprofileResource) updateEcccurveBindings(ctx context.Context, state *SslprofileResourceModel, plan *SslprofileResourceModel) error {
	tflog.Debug(ctx, "Updating ECC curve bindings")
	profileName := plan.Name.ValueString()

	var oldCurves, newCurves []string
	if !state.Ecccurvebindings.IsNull() {
		state.Ecccurvebindings.ElementsAs(context.Background(), &oldCurves, false)
	}
	if !plan.Ecccurvebindings.IsNull() {
		plan.Ecccurvebindings.ElementsAs(context.Background(), &newCurves, false)
	}

	oldSet := make(map[string]bool)
	for _, v := range oldCurves {
		oldSet[v] = true
	}
	newSet := make(map[string]bool)
	for _, v := range newCurves {
		newSet[v] = true
	}

	// Remove curves no longer in the set
	for _, v := range oldCurves {
		if !newSet[v] {
			args := []string{fmt.Sprintf("ecccurvename:%s", v)}
			if err := r.client.DeleteResourceWithArgs(service.Sslprofile_ecccurve_binding.Type(), profileName, args); err != nil {
				return err
			}
		}
	}

	// Add new curves
	for _, v := range newCurves {
		if !oldSet[v] {
			bindingStruct := ssl.Sslprofileecccurvebinding{
				Name:         profileName,
				Ecccurvename: v,
			}
			if _, err := r.client.UpdateResource(service.Sslprofile_ecccurve_binding.Type(), profileName, bindingStruct); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *SslprofileResource) readEcccurveBindings(ctx context.Context, data *SslprofileResourceModel, diags *diag.Diagnostics) {
	tflog.Debug(ctx, "Reading ECC curve bindings")
	profileName := data.Name.ValueString()

	bindings, _ := r.client.FindResourceArray(service.Sslprofile_ecccurve_binding.Type(), profileName)

	ecccurves := make([]attr.Value, 0, len(bindings))
	for _, binding := range bindings {
		if ecccurvename, ok := binding["ecccurvename"].(string); ok {
			ecccurves = append(ecccurves, types.StringValue(ecccurvename))
		}
	}
	setValue, d := types.SetValue(types.StringType, ecccurves)
	diags.Append(d...)
	data.Ecccurvebindings = setValue
}

// Cipher binding helpers

func (r *SslprofileResource) deleteDefaultCipherBindings(ctx context.Context, profileName string) error {
	tflog.Debug(ctx, "Deleting default cipher bindings")
	bindings, _ := r.client.FindResourceArray(service.Sslprofile_sslcipher_binding.Type(), profileName)
	for _, binding := range bindings {
		ciphername, ok := binding["cipheraliasname"].(string)
		if !ok {
			continue
		}
		args := []string{fmt.Sprintf("ciphername:%s", ciphername)}
		if err := r.client.DeleteResourceWithArgs(service.Sslprofile_sslcipher_binding.Type(), profileName, args); err != nil {
			tflog.Warn(ctx, fmt.Sprintf("Error deleting default cipher binding %s: %s", ciphername, err))
		}
	}
	return nil
}

func (r *SslprofileResource) createCipherBindings(ctx context.Context, data *SslprofileResourceModel) error {
	tflog.Debug(ctx, "Creating cipher bindings")
	profileName := data.Name.ValueString()

	// Delete default cipher bindings first
	defaultBindings, _ := r.client.FindResourceArray(service.Sslprofile_sslcipher_binding.Type(), profileName)
	for _, binding := range defaultBindings {
		ciphername, ok := binding["cipheraliasname"].(string)
		if !ok {
			continue
		}
		args := []string{fmt.Sprintf("ciphername:%s", ciphername)}
		r.client.DeleteResourceWithArgs(service.Sslprofile_sslcipher_binding.Type(), profileName, args)
	}

	// Add configured bindings
	var cipherBindings []CipherbindingModel
	data.Cipherbindings.ElementsAs(context.Background(), &cipherBindings, false)
	for _, cb := range cipherBindings {
		bindingStruct := ssl.Sslprofilecipherbinding{
			Name:       profileName,
			Ciphername: cb.Ciphername.ValueString(),
		}
		if !cb.Cipherpriority.IsNull() && !cb.Cipherpriority.IsUnknown() {
			bindingStruct.Cipherpriority = uint32(cb.Cipherpriority.ValueInt64())
		}
		if _, err := r.client.UpdateResource(service.Sslprofile_sslcipher_binding.Type(), profileName, bindingStruct); err != nil {
			return err
		}
	}
	return nil
}

func (r *SslprofileResource) updateCipherBindings(ctx context.Context, state *SslprofileResourceModel, plan *SslprofileResourceModel) error {
	tflog.Debug(ctx, "Updating cipher bindings")
	profileName := plan.Name.ValueString()

	var oldBindings, newBindings []CipherbindingModel
	if !state.Cipherbindings.IsNull() {
		state.Cipherbindings.ElementsAs(context.Background(), &oldBindings, false)
	}
	if !plan.Cipherbindings.IsNull() {
		plan.Cipherbindings.ElementsAs(context.Background(), &newBindings, false)
	}

	type cipherKey struct {
		name     string
		priority int64
	}

	oldSet := make(map[cipherKey]bool)
	for _, b := range oldBindings {
		key := cipherKey{name: b.Ciphername.ValueString(), priority: b.Cipherpriority.ValueInt64()}
		oldSet[key] = true
	}
	newSet := make(map[cipherKey]bool)
	for _, b := range newBindings {
		key := cipherKey{name: b.Ciphername.ValueString(), priority: b.Cipherpriority.ValueInt64()}
		newSet[key] = true
	}

	// Remove old bindings not in new set
	for _, b := range oldBindings {
		key := cipherKey{name: b.Ciphername.ValueString(), priority: b.Cipherpriority.ValueInt64()}
		if !newSet[key] {
			args := []string{fmt.Sprintf("ciphername:%s", b.Ciphername.ValueString())}
			if err := r.client.DeleteResourceWithArgs(service.Sslprofile_sslcipher_binding.Type(), profileName, args); err != nil {
				return err
			}
		}
	}

	// Add new bindings not in old set
	for _, b := range newBindings {
		key := cipherKey{name: b.Ciphername.ValueString(), priority: b.Cipherpriority.ValueInt64()}
		if !oldSet[key] {
			bindingStruct := ssl.Sslprofilecipherbinding{
				Name:       profileName,
				Ciphername: b.Ciphername.ValueString(),
			}
			if !b.Cipherpriority.IsNull() && !b.Cipherpriority.IsUnknown() {
				bindingStruct.Cipherpriority = uint32(b.Cipherpriority.ValueInt64())
			}
			if _, err := r.client.UpdateResource(service.Sslprofile_sslcipher_binding.Type(), profileName, bindingStruct); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *SslprofileResource) readCipherBindings(ctx context.Context, data *SslprofileResourceModel, diags *diag.Diagnostics) {
	tflog.Debug(ctx, "Reading cipher bindings")
	profileName := data.Name.ValueString()

	bindings, _ := r.client.FindResourceArray(service.Sslprofile_sslcipher_binding.Type(), profileName)

	cipherBindingAttrTypes := map[string]attr.Type{
		"ciphername":     types.StringType,
		"cipherpriority": types.Int64Type,
	}

	cipherBindings := make([]attr.Value, 0, len(bindings))
	for _, binding := range bindings {
		ciphername, _ := binding["cipheraliasname"].(string)
		var cipherpriority int64
		if cpStr, ok := binding["cipherpriority"].(string); ok {
			cp, err := strconv.Atoi(cpStr)
			if err == nil {
				cipherpriority = int64(cp)
			}
		}

		bindingObj, d := types.ObjectValue(cipherBindingAttrTypes, map[string]attr.Value{
			"ciphername":     types.StringValue(ciphername),
			"cipherpriority": types.Int64Value(cipherpriority),
		})
		diags.Append(d...)
		cipherBindings = append(cipherBindings, bindingObj)
	}

	setValue, d := types.SetValue(types.ObjectType{AttrTypes: cipherBindingAttrTypes}, cipherBindings)
	diags.Append(d...)
	data.Cipherbindings = setValue
}
