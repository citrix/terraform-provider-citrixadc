package sslprofile

import (
	"context"
	"fmt"
	"strconv"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
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
	if !r.readSslprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslprofile not found immediately after create")
		}
		return
	}

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

	found := r.readSslprofileFromApi(ctx, &data, &resp.Diagnostics)
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

	// Delta-payload update. `full` is the complete payload built from the plan (handles
	// value extraction/conversion and write-only secrets); `sslprofile` is a fresh payload
	// seeded with only the identity (name). For each attribute that genuinely changed and is
	// still present in config we copy just that field from `full` into `sslprofile`, so the PUT
	// carries name + changed fields only. Rebuilding the whole struct instead re-sends
	// context-dependent args and breaks both the v2 -> Framework upgrade (ec1094 "too few
	// arguments" when computed attrs are unknown) and steady-state updates (ec1092/1093
	// prerequisite conflicts). Attributes removed from config are collected in attributesToUnset
	// and cleared via utils.ExecuteUnset below. sslprofiletype is create-only (ForceNew) and is
	// never in the change set, so it can never leak into the update (ec278).
	full := sslprofileGetThePayloadFromthePlan(ctx, &data)
	sslprofileGetThePayloadFromtheConfig(ctx, &config, &full)
	sslprofile := ssl.Sslprofile{Name: data.Name.ValueString()}

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Allowextendedmastersecret.Equal(state.Allowextendedmastersecret) {
		tflog.Debug(ctx, fmt.Sprintf("allowextendedmastersecret has changed for sslprofile"))
		if config.Allowextendedmastersecret.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "allowextendedmastersecret")
		} else if !data.Allowextendedmastersecret.IsUnknown() && !data.Allowextendedmastersecret.IsNull() {
			sslprofile.Allowextendedmastersecret = full.Allowextendedmastersecret
			hasChange = true
		}
	}
	if !data.Allowunknownsni.IsUnknown() && !data.Allowunknownsni.IsNull() && !data.Allowunknownsni.Equal(state.Allowunknownsni) {
		tflog.Debug(ctx, fmt.Sprintf("allowunknownsni has changed for sslprofile"))
		sslprofile.Allowunknownsni = full.Allowunknownsni
		hasChange = true
	}
	if !data.Alpnprotocol.IsUnknown() && !data.Alpnprotocol.IsNull() && !data.Alpnprotocol.Equal(state.Alpnprotocol) {
		tflog.Debug(ctx, fmt.Sprintf("alpnprotocol has changed for sslprofile"))
		sslprofile.Alpnprotocol = full.Alpnprotocol
		hasChange = true
	}
	if !data.Ciphername.IsUnknown() && !data.Ciphername.IsNull() && !data.Ciphername.Equal(state.Ciphername) {
		tflog.Debug(ctx, fmt.Sprintf("ciphername has changed for sslprofile"))
		sslprofile.Ciphername = full.Ciphername
		hasChange = true
	}
	if !data.Cipherpriority.IsUnknown() && !data.Cipherpriority.IsNull() && !data.Cipherpriority.Equal(state.Cipherpriority) {
		tflog.Debug(ctx, fmt.Sprintf("cipherpriority has changed for sslprofile"))
		sslprofile.Cipherpriority = full.Cipherpriority
		hasChange = true
	}
	if !data.Cipherredirect.Equal(state.Cipherredirect) {
		tflog.Debug(ctx, fmt.Sprintf("cipherredirect has changed for sslprofile"))
		if config.Cipherredirect.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "cipherredirect")
		} else if !data.Cipherredirect.IsUnknown() && !data.Cipherredirect.IsNull() {
			sslprofile.Cipherredirect = full.Cipherredirect
			hasChange = true
		}
	}
	if !data.Cipherurl.IsUnknown() && !data.Cipherurl.IsNull() && !data.Cipherurl.Equal(state.Cipherurl) {
		tflog.Debug(ctx, fmt.Sprintf("cipherurl has changed for sslprofile"))
		sslprofile.Cipherurl = full.Cipherurl
		hasChange = true
	}
	if !data.Cleartextport.IsUnknown() && !data.Cleartextport.IsNull() && !data.Cleartextport.Equal(state.Cleartextport) {
		tflog.Debug(ctx, fmt.Sprintf("cleartextport has changed for sslprofile"))
		sslprofile.Cleartextport = full.Cleartextport
		hasChange = true
	}
	if !data.Clientauth.Equal(state.Clientauth) {
		tflog.Debug(ctx, fmt.Sprintf("clientauth has changed for sslprofile"))
		if config.Clientauth.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "clientauth")
		} else if !data.Clientauth.IsUnknown() && !data.Clientauth.IsNull() {
			sslprofile.Clientauth = full.Clientauth
			hasChange = true
		}
	}
	if !data.Clientauthuseboundcachain.Equal(state.Clientauthuseboundcachain) {
		tflog.Debug(ctx, fmt.Sprintf("clientauthuseboundcachain has changed for sslprofile"))
		if config.Clientauthuseboundcachain.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "clientauthuseboundcachain")
		} else if !data.Clientauthuseboundcachain.IsUnknown() && !data.Clientauthuseboundcachain.IsNull() {
			sslprofile.Clientauthuseboundcachain = full.Clientauthuseboundcachain
			hasChange = true
		}
	}
	if !data.Clientcert.IsUnknown() && !data.Clientcert.IsNull() && !data.Clientcert.Equal(state.Clientcert) {
		tflog.Debug(ctx, fmt.Sprintf("clientcert has changed for sslprofile"))
		sslprofile.Clientcert = full.Clientcert
		hasChange = true
	}
	if !data.Commonname.IsUnknown() && !data.Commonname.IsNull() && !data.Commonname.Equal(state.Commonname) {
		tflog.Debug(ctx, fmt.Sprintf("commonname has changed for sslprofile"))
		sslprofile.Commonname = full.Commonname
		hasChange = true
	}
	if !data.Defaultsni.IsUnknown() && !data.Defaultsni.IsNull() && !data.Defaultsni.Equal(state.Defaultsni) {
		tflog.Debug(ctx, fmt.Sprintf("defaultsni has changed for sslprofile"))
		sslprofile.Defaultsni = full.Defaultsni
		hasChange = true
	}
	if !data.Denysslreneg.Equal(state.Denysslreneg) {
		tflog.Debug(ctx, fmt.Sprintf("denysslreneg has changed for sslprofile"))
		if config.Denysslreneg.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "denysslreneg")
		} else if !data.Denysslreneg.IsUnknown() && !data.Denysslreneg.IsNull() {
			sslprofile.Denysslreneg = full.Denysslreneg
			hasChange = true
		}
	}
	if !data.Dh.Equal(state.Dh) {
		tflog.Debug(ctx, fmt.Sprintf("dh has changed for sslprofile"))
		if config.Dh.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "dh")
		} else if !data.Dh.IsUnknown() && !data.Dh.IsNull() {
			sslprofile.Dh = full.Dh
			hasChange = true
		}
	}
	if !data.Dhcount.IsUnknown() && !data.Dhcount.IsNull() && !data.Dhcount.Equal(state.Dhcount) {
		tflog.Debug(ctx, fmt.Sprintf("dhcount has changed for sslprofile"))
		sslprofile.Dhcount = full.Dhcount
		hasChange = true
	}
	if !data.Dhekeyexchangewithpsk.Equal(state.Dhekeyexchangewithpsk) {
		tflog.Debug(ctx, fmt.Sprintf("dhekeyexchangewithpsk has changed for sslprofile"))
		if config.Dhekeyexchangewithpsk.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "dhekeyexchangewithpsk")
		} else if !data.Dhekeyexchangewithpsk.IsUnknown() && !data.Dhekeyexchangewithpsk.IsNull() {
			sslprofile.Dhekeyexchangewithpsk = full.Dhekeyexchangewithpsk
			hasChange = true
		}
	}
	if !data.Dhfile.IsUnknown() && !data.Dhfile.IsNull() && !data.Dhfile.Equal(state.Dhfile) {
		tflog.Debug(ctx, fmt.Sprintf("dhfile has changed for sslprofile"))
		sslprofile.Dhfile = full.Dhfile
		hasChange = true
	}
	if !data.Dhkeyexpsizelimit.Equal(state.Dhkeyexpsizelimit) {
		tflog.Debug(ctx, fmt.Sprintf("dhkeyexpsizelimit has changed for sslprofile"))
		if config.Dhkeyexpsizelimit.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "dhkeyexpsizelimit")
		} else if !data.Dhkeyexpsizelimit.IsUnknown() && !data.Dhkeyexpsizelimit.IsNull() {
			sslprofile.Dhkeyexpsizelimit = full.Dhkeyexpsizelimit
			hasChange = true
		}
	}
	if !data.Dropreqwithnohostheader.Equal(state.Dropreqwithnohostheader) {
		tflog.Debug(ctx, fmt.Sprintf("dropreqwithnohostheader has changed for sslprofile"))
		if config.Dropreqwithnohostheader.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "dropreqwithnohostheader")
		} else if !data.Dropreqwithnohostheader.IsUnknown() && !data.Dropreqwithnohostheader.IsNull() {
			sslprofile.Dropreqwithnohostheader = full.Dropreqwithnohostheader
			hasChange = true
		}
	}
	if !data.Encryptedclienthello.Equal(state.Encryptedclienthello) {
		tflog.Debug(ctx, fmt.Sprintf("encryptedclienthello has changed for sslprofile"))
		if config.Encryptedclienthello.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "encryptedclienthello")
		} else if !data.Encryptedclienthello.IsUnknown() && !data.Encryptedclienthello.IsNull() {
			sslprofile.Encryptedclienthello = full.Encryptedclienthello
			hasChange = true
		}
	}
	if !data.Encrypttriggerpktcount.Equal(state.Encrypttriggerpktcount) {
		tflog.Debug(ctx, fmt.Sprintf("encrypttriggerpktcount has changed for sslprofile"))
		if config.Encrypttriggerpktcount.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "encrypttriggerpktcount")
		} else if !data.Encrypttriggerpktcount.IsUnknown() && !data.Encrypttriggerpktcount.IsNull() {
			sslprofile.Encrypttriggerpktcount = full.Encrypttriggerpktcount
			hasChange = true
		}
	}
	if !data.Ersa.IsUnknown() && !data.Ersa.IsNull() && !data.Ersa.Equal(state.Ersa) {
		tflog.Debug(ctx, fmt.Sprintf("ersa has changed for sslprofile"))
		sslprofile.Ersa = full.Ersa
		hasChange = true
	}
	if !data.Ersacount.IsUnknown() && !data.Ersacount.IsNull() && !data.Ersacount.Equal(state.Ersacount) {
		tflog.Debug(ctx, fmt.Sprintf("ersacount has changed for sslprofile"))
		sslprofile.Ersacount = full.Ersacount
		hasChange = true
	}
	if !data.Hsts.Equal(state.Hsts) {
		tflog.Debug(ctx, fmt.Sprintf("hsts has changed for sslprofile"))
		if config.Hsts.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "hsts")
		} else if !data.Hsts.IsUnknown() && !data.Hsts.IsNull() {
			sslprofile.Hsts = full.Hsts
			hasChange = true
		}
	}
	if !data.Includesubdomains.Equal(state.Includesubdomains) {
		tflog.Debug(ctx, fmt.Sprintf("includesubdomains has changed for sslprofile"))
		if config.Includesubdomains.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "includesubdomains")
		} else if !data.Includesubdomains.IsUnknown() && !data.Includesubdomains.IsNull() {
			sslprofile.Includesubdomains = full.Includesubdomains
			hasChange = true
		}
	}
	if !data.Insertionencoding.IsUnknown() && !data.Insertionencoding.IsNull() && !data.Insertionencoding.Equal(state.Insertionencoding) {
		tflog.Debug(ctx, fmt.Sprintf("insertionencoding has changed for sslprofile"))
		sslprofile.Insertionencoding = full.Insertionencoding
		hasChange = true
	}
	if !data.Maxage.Equal(state.Maxage) {
		tflog.Debug(ctx, fmt.Sprintf("maxage has changed for sslprofile"))
		if config.Maxage.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "maxage")
		} else if !data.Maxage.IsUnknown() && !data.Maxage.IsNull() {
			sslprofile.Maxage = full.Maxage
			hasChange = true
		}
	}
	if !data.Maxrenegrate.IsUnknown() && !data.Maxrenegrate.IsNull() && !data.Maxrenegrate.Equal(state.Maxrenegrate) {
		tflog.Debug(ctx, fmt.Sprintf("maxrenegrate has changed for sslprofile"))
		sslprofile.Maxrenegrate = full.Maxrenegrate
		hasChange = true
	}
	if !data.Nodefaultbindings.IsUnknown() && !data.Nodefaultbindings.IsNull() && !data.Nodefaultbindings.Equal(state.Nodefaultbindings) {
		tflog.Debug(ctx, fmt.Sprintf("nodefaultbindings has changed for sslprofile"))
		sslprofile.Nodefaultbindings = full.Nodefaultbindings
		hasChange = true
	}
	if !data.Ocspstapling.Equal(state.Ocspstapling) {
		tflog.Debug(ctx, fmt.Sprintf("ocspstapling has changed for sslprofile"))
		if config.Ocspstapling.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "ocspstapling")
		} else if !data.Ocspstapling.IsUnknown() && !data.Ocspstapling.IsNull() {
			sslprofile.Ocspstapling = full.Ocspstapling
			hasChange = true
		}
	}
	if !data.Preload.Equal(state.Preload) {
		tflog.Debug(ctx, fmt.Sprintf("preload has changed for sslprofile"))
		if config.Preload.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "preload")
		} else if !data.Preload.IsUnknown() && !data.Preload.IsNull() {
			sslprofile.Preload = full.Preload
			hasChange = true
		}
	}
	if !data.Prevsessionkeylifetime.Equal(state.Prevsessionkeylifetime) {
		tflog.Debug(ctx, fmt.Sprintf("prevsessionkeylifetime has changed for sslprofile"))
		if config.Prevsessionkeylifetime.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "prevsessionkeylifetime")
		} else if !data.Prevsessionkeylifetime.IsUnknown() && !data.Prevsessionkeylifetime.IsNull() {
			sslprofile.Prevsessionkeylifetime = full.Prevsessionkeylifetime
			hasChange = true
		}
	}
	if !data.Pushenctrigger.IsUnknown() && !data.Pushenctrigger.IsNull() && !data.Pushenctrigger.Equal(state.Pushenctrigger) {
		tflog.Debug(ctx, fmt.Sprintf("pushenctrigger has changed for sslprofile"))
		sslprofile.Pushenctrigger = full.Pushenctrigger
		hasChange = true
	}
	if !data.Pushenctriggertimeout.Equal(state.Pushenctriggertimeout) {
		tflog.Debug(ctx, fmt.Sprintf("pushenctriggertimeout has changed for sslprofile"))
		if config.Pushenctriggertimeout.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "pushenctriggertimeout")
		} else if !data.Pushenctriggertimeout.IsUnknown() && !data.Pushenctriggertimeout.IsNull() {
			sslprofile.Pushenctriggertimeout = full.Pushenctriggertimeout
			hasChange = true
		}
	}
	if !data.Pushflag.IsUnknown() && !data.Pushflag.IsNull() && !data.Pushflag.Equal(state.Pushflag) {
		tflog.Debug(ctx, fmt.Sprintf("pushflag has changed for sslprofile"))
		sslprofile.Pushflag = full.Pushflag
		hasChange = true
	}
	if !data.Quantumsize.Equal(state.Quantumsize) {
		tflog.Debug(ctx, fmt.Sprintf("quantumsize has changed for sslprofile"))
		if config.Quantumsize.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "quantumsize")
		} else if !data.Quantumsize.IsUnknown() && !data.Quantumsize.IsNull() {
			sslprofile.Quantumsize = full.Quantumsize
			hasChange = true
		}
	}
	if !data.Redirectportrewrite.Equal(state.Redirectportrewrite) {
		tflog.Debug(ctx, fmt.Sprintf("redirectportrewrite has changed for sslprofile"))
		if config.Redirectportrewrite.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "redirectportrewrite")
		} else if !data.Redirectportrewrite.IsUnknown() && !data.Redirectportrewrite.IsNull() {
			sslprofile.Redirectportrewrite = full.Redirectportrewrite
			hasChange = true
		}
	}
	if !data.Sendclosenotify.Equal(state.Sendclosenotify) {
		tflog.Debug(ctx, fmt.Sprintf("sendclosenotify has changed for sslprofile"))
		if config.Sendclosenotify.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sendclosenotify")
		} else if !data.Sendclosenotify.IsUnknown() && !data.Sendclosenotify.IsNull() {
			sslprofile.Sendclosenotify = full.Sendclosenotify
			hasChange = true
		}
	}
	if !data.Serverauth.Equal(state.Serverauth) {
		tflog.Debug(ctx, fmt.Sprintf("serverauth has changed for sslprofile"))
		if config.Serverauth.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "serverauth")
		} else if !data.Serverauth.IsUnknown() && !data.Serverauth.IsNull() {
			sslprofile.Serverauth = full.Serverauth
			hasChange = true
		}
	}
	if !data.Sessionkeylifetime.IsUnknown() && !data.Sessionkeylifetime.IsNull() && !data.Sessionkeylifetime.Equal(state.Sessionkeylifetime) {
		tflog.Debug(ctx, fmt.Sprintf("sessionkeylifetime has changed for sslprofile"))
		sslprofile.Sessionkeylifetime = full.Sessionkeylifetime
		hasChange = true
	}
	if !data.Sessionticket.Equal(state.Sessionticket) {
		tflog.Debug(ctx, fmt.Sprintf("sessionticket has changed for sslprofile"))
		if config.Sessionticket.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sessionticket")
		} else if !data.Sessionticket.IsUnknown() && !data.Sessionticket.IsNull() {
			sslprofile.Sessionticket = full.Sessionticket
			hasChange = true
		}
	}
	// Check secret attribute sessionticketkeydata or its version tracker. Only send the
	// secret when the config actually supplies it (full.Sessionticketkeydata != ""): the
	// plain value must be known+non-null, or the _wo_version bumped while a _wo value is
	// present. This stops the legacy null/"" mismatch and the Default(1) _wo_version drift
	// from firing a spurious update on the v2 -> Framework upgrade.
	if !data.Sessionticketkeydata.IsUnknown() && !data.Sessionticketkeydata.IsNull() && !data.Sessionticketkeydata.Equal(state.Sessionticketkeydata) && full.Sessionticketkeydata != "" {
		tflog.Debug(ctx, fmt.Sprintf("sessionticketkeydata has changed for sslprofile"))
		sslprofile.Sessionticketkeydata = full.Sessionticketkeydata
		hasChange = true
	} else if !data.SessionticketkeydataWoVersion.Equal(state.SessionticketkeydataWoVersion) && full.Sessionticketkeydata != "" {
		tflog.Debug(ctx, fmt.Sprintf("sessionticketkeydata_wo_version has changed for sslprofile"))
		sslprofile.Sessionticketkeydata = full.Sessionticketkeydata
		hasChange = true
	}
	if !data.Sessionticketkeyrefresh.IsUnknown() && !data.Sessionticketkeyrefresh.IsNull() && !data.Sessionticketkeyrefresh.Equal(state.Sessionticketkeyrefresh) {
		tflog.Debug(ctx, fmt.Sprintf("sessionticketkeyrefresh has changed for sslprofile"))
		sslprofile.Sessionticketkeyrefresh = full.Sessionticketkeyrefresh
		hasChange = true
	}
	if !data.Sessionticketlifetime.IsUnknown() && !data.Sessionticketlifetime.IsNull() && !data.Sessionticketlifetime.Equal(state.Sessionticketlifetime) {
		tflog.Debug(ctx, fmt.Sprintf("sessionticketlifetime has changed for sslprofile"))
		sslprofile.Sessionticketlifetime = full.Sessionticketlifetime
		hasChange = true
	}
	if !data.Sessreuse.Equal(state.Sessreuse) {
		tflog.Debug(ctx, fmt.Sprintf("sessreuse has changed for sslprofile"))
		if config.Sessreuse.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sessreuse")
		} else if !data.Sessreuse.IsUnknown() && !data.Sessreuse.IsNull() {
			sslprofile.Sessreuse = full.Sessreuse
			hasChange = true
		}
	}
	if !data.Sesstimeout.IsUnknown() && !data.Sesstimeout.IsNull() && !data.Sesstimeout.Equal(state.Sesstimeout) {
		tflog.Debug(ctx, fmt.Sprintf("sesstimeout has changed for sslprofile"))
		sslprofile.Sesstimeout = full.Sesstimeout
		hasChange = true
	}
	if !data.Skipclientcertpolicycheck.IsUnknown() && !data.Skipclientcertpolicycheck.IsNull() && !data.Skipclientcertpolicycheck.Equal(state.Skipclientcertpolicycheck) {
		tflog.Debug(ctx, fmt.Sprintf("skipclientcertpolicycheck has changed for sslprofile"))
		sslprofile.Skipclientcertpolicycheck = full.Skipclientcertpolicycheck
		hasChange = true
	}
	if !data.Snienable.Equal(state.Snienable) {
		tflog.Debug(ctx, fmt.Sprintf("snienable has changed for sslprofile"))
		if config.Snienable.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "snienable")
		} else if !data.Snienable.IsUnknown() && !data.Snienable.IsNull() {
			sslprofile.Snienable = full.Snienable
			hasChange = true
		}
	}
	if !data.Snihttphostmatch.Equal(state.Snihttphostmatch) {
		tflog.Debug(ctx, fmt.Sprintf("snihttphostmatch has changed for sslprofile"))
		if config.Snihttphostmatch.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "snihttphostmatch")
		} else if !data.Snihttphostmatch.IsUnknown() && !data.Snihttphostmatch.IsNull() {
			sslprofile.Snihttphostmatch = full.Snihttphostmatch
			hasChange = true
		}
	}
	if !data.Ssl3.Equal(state.Ssl3) {
		tflog.Debug(ctx, fmt.Sprintf("ssl3 has changed for sslprofile"))
		if config.Ssl3.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "ssl3")
		} else if !data.Ssl3.IsUnknown() && !data.Ssl3.IsNull() {
			sslprofile.Ssl3 = full.Ssl3
			hasChange = true
		}
	}
	if !data.Sslclientlogs.Equal(state.Sslclientlogs) {
		tflog.Debug(ctx, fmt.Sprintf("sslclientlogs has changed for sslprofile"))
		if config.Sslclientlogs.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sslclientlogs")
		} else if !data.Sslclientlogs.IsUnknown() && !data.Sslclientlogs.IsNull() {
			sslprofile.Sslclientlogs = full.Sslclientlogs
			hasChange = true
		}
	}
	if !data.Sslimaxsessperserver.Equal(state.Sslimaxsessperserver) {
		tflog.Debug(ctx, fmt.Sprintf("sslimaxsessperserver has changed for sslprofile"))
		if config.Sslimaxsessperserver.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sslimaxsessperserver")
		} else if !data.Sslimaxsessperserver.IsUnknown() && !data.Sslimaxsessperserver.IsNull() {
			sslprofile.Sslimaxsessperserver = full.Sslimaxsessperserver
			hasChange = true
		}
	}
	if !data.Sslinterception.Equal(state.Sslinterception) {
		tflog.Debug(ctx, fmt.Sprintf("sslinterception has changed for sslprofile"))
		if config.Sslinterception.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sslinterception")
		} else if !data.Sslinterception.IsUnknown() && !data.Sslinterception.IsNull() {
			sslprofile.Sslinterception = full.Sslinterception
			hasChange = true
		}
	}
	if !data.Ssliocspcheck.Equal(state.Ssliocspcheck) {
		tflog.Debug(ctx, fmt.Sprintf("ssliocspcheck has changed for sslprofile"))
		if config.Ssliocspcheck.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "ssliocspcheck")
		} else if !data.Ssliocspcheck.IsUnknown() && !data.Ssliocspcheck.IsNull() {
			sslprofile.Ssliocspcheck = full.Ssliocspcheck
			hasChange = true
		}
	}
	if !data.Sslireneg.Equal(state.Sslireneg) {
		tflog.Debug(ctx, fmt.Sprintf("sslireneg has changed for sslprofile"))
		if config.Sslireneg.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sslireneg")
		} else if !data.Sslireneg.IsUnknown() && !data.Sslireneg.IsNull() {
			sslprofile.Sslireneg = full.Sslireneg
			hasChange = true
		}
	}
	if !data.Ssllogprofile.IsUnknown() && !data.Ssllogprofile.IsNull() && !data.Ssllogprofile.Equal(state.Ssllogprofile) {
		tflog.Debug(ctx, fmt.Sprintf("ssllogprofile has changed for sslprofile"))
		sslprofile.Ssllogprofile = full.Ssllogprofile
		hasChange = true
	}
	if !data.Sslredirect.Equal(state.Sslredirect) {
		tflog.Debug(ctx, fmt.Sprintf("sslredirect has changed for sslprofile"))
		if config.Sslredirect.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sslredirect")
		} else if !data.Sslredirect.IsUnknown() && !data.Sslredirect.IsNull() {
			sslprofile.Sslredirect = full.Sslredirect
			hasChange = true
		}
	}
	if !data.Ssltriggertimeout.Equal(state.Ssltriggertimeout) {
		tflog.Debug(ctx, fmt.Sprintf("ssltriggertimeout has changed for sslprofile"))
		if config.Ssltriggertimeout.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "ssltriggertimeout")
		} else if !data.Ssltriggertimeout.IsUnknown() && !data.Ssltriggertimeout.IsNull() {
			sslprofile.Ssltriggertimeout = full.Ssltriggertimeout
			hasChange = true
		}
	}
	if !data.Strictcachecks.Equal(state.Strictcachecks) {
		tflog.Debug(ctx, fmt.Sprintf("strictcachecks has changed for sslprofile"))
		if config.Strictcachecks.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "strictcachecks")
		} else if !data.Strictcachecks.IsUnknown() && !data.Strictcachecks.IsNull() {
			sslprofile.Strictcachecks = full.Strictcachecks
			hasChange = true
		}
	}
	if !data.Strictsigdigestcheck.IsUnknown() && !data.Strictsigdigestcheck.IsNull() && !data.Strictsigdigestcheck.Equal(state.Strictsigdigestcheck) {
		tflog.Debug(ctx, fmt.Sprintf("strictsigdigestcheck has changed for sslprofile"))
		sslprofile.Strictsigdigestcheck = full.Strictsigdigestcheck
		hasChange = true
	}
	if !data.Tls1.IsUnknown() && !data.Tls1.IsNull() && !data.Tls1.Equal(state.Tls1) {
		tflog.Debug(ctx, fmt.Sprintf("tls1 has changed for sslprofile"))
		sslprofile.Tls1 = full.Tls1
		hasChange = true
	}
	if !data.Tls11.IsUnknown() && !data.Tls11.IsNull() && !data.Tls11.Equal(state.Tls11) {
		tflog.Debug(ctx, fmt.Sprintf("tls11 has changed for sslprofile"))
		sslprofile.Tls11 = full.Tls11
		hasChange = true
	}
	if !data.Tls12.IsUnknown() && !data.Tls12.IsNull() && !data.Tls12.Equal(state.Tls12) {
		tflog.Debug(ctx, fmt.Sprintf("tls12 has changed for sslprofile"))
		sslprofile.Tls12 = full.Tls12
		hasChange = true
	}
	if !data.Tls13.IsUnknown() && !data.Tls13.IsNull() && !data.Tls13.Equal(state.Tls13) {
		tflog.Debug(ctx, fmt.Sprintf("tls13 has changed for sslprofile"))
		sslprofile.Tls13 = full.Tls13
		hasChange = true
	}
	if !data.Tls13sessionticketsperauthcontext.Equal(state.Tls13sessionticketsperauthcontext) {
		tflog.Debug(ctx, fmt.Sprintf("tls13sessionticketsperauthcontext has changed for sslprofile"))
		if config.Tls13sessionticketsperauthcontext.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "tls13sessionticketsperauthcontext")
		} else if !data.Tls13sessionticketsperauthcontext.IsUnknown() && !data.Tls13sessionticketsperauthcontext.IsNull() {
			sslprofile.Tls13sessionticketsperauthcontext = full.Tls13sessionticketsperauthcontext
			hasChange = true
		}
	}
	if !data.Zerorttearlydata.Equal(state.Zerorttearlydata) {
		tflog.Debug(ctx, fmt.Sprintf("zerorttearlydata has changed for sslprofile"))
		if config.Zerorttearlydata.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "zerorttearlydata")
		} else if !data.Zerorttearlydata.IsUnknown() && !data.Zerorttearlydata.IsNull() {
			sslprofile.Zerorttearlydata = full.Zerorttearlydata
			hasChange = true
		}
	}

	// sslprofiletype is create-only (ForceNew); ensure it never leaks into the update payload.
	sslprofile.Sslprofiletype = ""

	if hasChange {
		// Create API request body from the model
		// Get the SET (update) payload from plan. Per sslprofile.html, sslprofiletype
		// is create-only (present in the add payload, absent from the update payload)
		// and is EXCLUDED by the update builder; the "set sslprofile" call would
		// otherwise reject it (e.g. errorcode 278 "Invalid argument [sslprofiletype]").
		// A genuine change to sslprofiletype still forces a full destroy+recreate via
		// RequiresReplace. ciphername IS accepted by the update payload, so it is kept.
		sslprofile := sslprofileGetTheUpdatePayloadFromthePlan(ctx, &data)
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

	// Clear attributes that were removed from the configuration (update-then-unset
	// ordering, so any default carried by the update payload is superseded here).
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Sslprofile.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset sslprofile attributes, got error: %s", err))
		return
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
	if !r.readSslprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslprofile not found immediately after update")
		}
		return
	}

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
func (r *SslprofileResource) readSslprofileFromApi(ctx context.Context, data *SslprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Sslprofile.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslprofile, got error: %s", err))
		return false
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

	return true
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
