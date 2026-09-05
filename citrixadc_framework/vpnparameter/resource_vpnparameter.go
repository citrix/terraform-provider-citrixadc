package vpnparameter

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
var _ resource.Resource = &VpnparameterResource{}
var _ resource.ResourceWithConfigure = (*VpnparameterResource)(nil)
var _ resource.ResourceWithImportState = (*VpnparameterResource)(nil)

func NewVpnparameterResource() resource.Resource {
	return &VpnparameterResource{}
}

// VpnparameterResource defines the resource implementation.
type VpnparameterResource struct {
	client *service.NitroClient
}

func (r *VpnparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnparameter"
}

func (r *VpnparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnparameter resource")

	// Create API request body from the model
	vpnparameter := vpnparameterGetThePayloadFromtheConfig(ctx, &data)

	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vpnparameter.Type(), &vpnparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnparameter, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnparameter-config")

	tflog.Trace(ctx, "Created vpnparameter resource")

	// Read the updated state back
	r.readVpnparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnparameter resource")

	r.readVpnparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state VpnparameterResourceModel

	// Read Terraform prior state, plan, and config into the models
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnparameter resource")

	// Check if there are any changes in updateable attributes and collect
	// attributes that were removed from config so they can be unset. Only
	// attributes that carry a server default (and are unset-eligible per the
	// NITRO spec on this appliance) are routed to the unset path; the
	// config.IsNull() check keys off the schema Default value planned when such
	// an attribute is removed from config.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Accessrestrictedpageredirect.Equal(state.Accessrestrictedpageredirect) {
		if !config.Accessrestrictedpageredirect.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Advancedclientlessvpnmode.Equal(state.Advancedclientlessvpnmode) {
		if !config.Advancedclientlessvpnmode.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Allowedlogingroups.Equal(state.Allowedlogingroups) {
		if !config.Allowedlogingroups.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Allprotocolproxy.Equal(state.Allprotocolproxy) {
		if !config.Allprotocolproxy.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Alwaysonprofilename.Equal(state.Alwaysonprofilename) {
		if !config.Alwaysonprofilename.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Apptokentimeout.Equal(state.Apptokentimeout) {
		if !config.Apptokentimeout.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Authorizationgroup.Equal(state.Authorizationgroup) {
		if !config.Authorizationgroup.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Autoproxyurl.Equal(state.Autoproxyurl) {
		if !config.Autoproxyurl.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Backendcertvalidation.Equal(state.Backendcertvalidation) {
		if config.Backendcertvalidation.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "backendcertvalidation")
		} else {
			hasChange = true
		}
	}
	if !data.Backenddtls12.Equal(state.Backenddtls12) {
		if config.Backenddtls12.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "backenddtls12")
		} else {
			hasChange = true
		}
	}
	if !data.Backendserversni.Equal(state.Backendserversni) {
		if config.Backendserversni.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "backendserversni")
		} else {
			hasChange = true
		}
	}
	if !data.Citrixreceiverhome.Equal(state.Citrixreceiverhome) {
		if !config.Citrixreceiverhome.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Clientchoices.Equal(state.Clientchoices) {
		if !config.Clientchoices.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Clientcleanupprompt.Equal(state.Clientcleanupprompt) {
		if config.Clientcleanupprompt.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "clientcleanupprompt")
		} else {
			hasChange = true
		}
	}
	if !data.Clientconfiguration.Equal(state.Clientconfiguration) {
		if !config.Clientconfiguration.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Clientdebug.Equal(state.Clientdebug) {
		if !config.Clientdebug.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Clientidletimeout.Equal(state.Clientidletimeout) {
		if !config.Clientidletimeout.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Clientlessmodeurlencoding.Equal(state.Clientlessmodeurlencoding) {
		if !config.Clientlessmodeurlencoding.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Clientlesspersistentcookie.Equal(state.Clientlesspersistentcookie) {
		if !config.Clientlesspersistentcookie.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Clientlessvpnmode.Equal(state.Clientlessvpnmode) {
		if !config.Clientlessvpnmode.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Clientoptions.Equal(state.Clientoptions) {
		if !config.Clientoptions.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Clientsecurity.Equal(state.Clientsecurity) {
		if !config.Clientsecurity.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Clientsecuritygroup.Equal(state.Clientsecuritygroup) {
		if !config.Clientsecuritygroup.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Clientsecuritylog.Equal(state.Clientsecuritylog) {
		if !config.Clientsecuritylog.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Clientsecuritymessage.Equal(state.Clientsecuritymessage) {
		if !config.Clientsecuritymessage.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Clientversions.Equal(state.Clientversions) {
		if !config.Clientversions.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Defaultauthorizationaction.Equal(state.Defaultauthorizationaction) {
		if !config.Defaultauthorizationaction.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Deviceposture.Equal(state.Deviceposture) {
		if config.Deviceposture.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "deviceposture")
		} else {
			hasChange = true
		}
	}
	if !data.Dnsvservername.Equal(state.Dnsvservername) {
		if !config.Dnsvservername.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Emailhome.Equal(state.Emailhome) {
		if !config.Emailhome.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Encryptcsecexp.Equal(state.Encryptcsecexp) {
		if config.Encryptcsecexp.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "encryptcsecexp")
		} else {
			hasChange = true
		}
	}
	if !data.Epaclienttype.Equal(state.Epaclienttype) {
		if !config.Epaclienttype.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Forcecleanup.Equal(state.Forcecleanup) {
		if !config.Forcecleanup.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Forcedtimeout.Equal(state.Forcedtimeout) {
		if !config.Forcedtimeout.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Forcedtimeoutwarning.Equal(state.Forcedtimeoutwarning) {
		if !config.Forcedtimeoutwarning.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Fqdnspoofedip.Equal(state.Fqdnspoofedip) {
		if !config.Fqdnspoofedip.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Ftpproxy.Equal(state.Ftpproxy) {
		if !config.Ftpproxy.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Gopherproxy.Equal(state.Gopherproxy) {
		if !config.Gopherproxy.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Homepage.Equal(state.Homepage) {
		if !config.Homepage.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Httpport.Equal(state.Httpport) {
		if !config.Httpport.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Httpproxy.Equal(state.Httpproxy) {
		if !config.Httpproxy.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Httptrackconnproxy.Equal(state.Httptrackconnproxy) {
		if config.Httptrackconnproxy.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "httptrackconnproxy")
		} else {
			hasChange = true
		}
	}
	if !data.Icaproxy.Equal(state.Icaproxy) {
		if !config.Icaproxy.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Icasessiontimeout.Equal(state.Icasessiontimeout) {
		if !config.Icasessiontimeout.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Icauseraccounting.Equal(state.Icauseraccounting) {
		if !config.Icauseraccounting.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Iconwithreceiver.Equal(state.Iconwithreceiver) {
		if !config.Iconwithreceiver.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Iipdnssuffix.Equal(state.Iipdnssuffix) {
		if !config.Iipdnssuffix.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Kcdaccount.Equal(state.Kcdaccount) {
		if !config.Kcdaccount.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Killconnections.Equal(state.Killconnections) {
		if config.Killconnections.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "killconnections")
		} else {
			hasChange = true
		}
	}
	if !data.Linuxpluginupgrade.Equal(state.Linuxpluginupgrade) {
		if !config.Linuxpluginupgrade.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Locallanaccess.Equal(state.Locallanaccess) {
		if config.Locallanaccess.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "locallanaccess")
		} else {
			hasChange = true
		}
	}
	if !data.Loginscript.Equal(state.Loginscript) {
		if !config.Loginscript.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Logoutscript.Equal(state.Logoutscript) {
		if !config.Logoutscript.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Macpluginupgrade.Equal(state.Macpluginupgrade) {
		if !config.Macpluginupgrade.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Maxiipperuser.Equal(state.Maxiipperuser) {
		if config.Maxiipperuser.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "maxiipperuser")
		} else {
			hasChange = true
		}
	}
	if !data.Mdxtokentimeout.Equal(state.Mdxtokentimeout) {
		if !config.Mdxtokentimeout.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Netmask.Equal(state.Netmask) {
		if !config.Netmask.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Ntdomain.Equal(state.Ntdomain) {
		if !config.Ntdomain.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Pcoipprofilename.Equal(state.Pcoipprofilename) {
		if !config.Pcoipprofilename.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Proxy.Equal(state.Proxy) {
		if !config.Proxy.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Proxyexception.Equal(state.Proxyexception) {
		if !config.Proxyexception.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Proxylocalbypass.Equal(state.Proxylocalbypass) {
		if config.Proxylocalbypass.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "proxylocalbypass")
		} else {
			hasChange = true
		}
	}
	if !data.Rdpclientprofilename.Equal(state.Rdpclientprofilename) {
		if !config.Rdpclientprofilename.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Rfc1918.Equal(state.Rfc1918) {
		if config.Rfc1918.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "rfc1918")
		} else {
			hasChange = true
		}
	}
	if !data.Samesite.Equal(state.Samesite) {
		if !config.Samesite.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Securebrowse.Equal(state.Securebrowse) {
		if config.Securebrowse.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "securebrowse")
		} else {
			hasChange = true
		}
	}
	if !data.Secureprivateaccess.Equal(state.Secureprivateaccess) {
		if !config.Secureprivateaccess.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Secureprivateaccessprofile.Equal(state.Secureprivateaccessprofile) {
		if !config.Secureprivateaccessprofile.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Sesstimeout.Equal(state.Sesstimeout) {
		if !config.Sesstimeout.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Smartgroup.Equal(state.Smartgroup) {
		if !config.Smartgroup.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Socksproxy.Equal(state.Socksproxy) {
		if !config.Socksproxy.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Splitdns.Equal(state.Splitdns) {
		if !config.Splitdns.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Splittunnel.Equal(state.Splittunnel) {
		if config.Splittunnel.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "splittunnel")
		} else {
			hasChange = true
		}
	}
	if !data.Spoofiip.Equal(state.Spoofiip) {
		if config.Spoofiip.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "spoofiip")
		} else {
			hasChange = true
		}
	}
	if !data.Sslproxy.Equal(state.Sslproxy) {
		if !config.Sslproxy.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Sso.Equal(state.Sso) {
		if config.Sso.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "sso")
		} else {
			hasChange = true
		}
	}
	if !data.Ssocredential.Equal(state.Ssocredential) {
		if config.Ssocredential.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "ssocredential")
		} else {
			hasChange = true
		}
	}
	if !data.Storefronturl.Equal(state.Storefronturl) {
		if !config.Storefronturl.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Transparentinterception.Equal(state.Transparentinterception) {
		if config.Transparentinterception.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "transparentinterception")
		} else {
			hasChange = true
		}
	}
	if !data.Uitheme.Equal(state.Uitheme) {
		if !config.Uitheme.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Useiip.Equal(state.Useiip) {
		if config.Useiip.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "useiip")
		} else {
			hasChange = true
		}
	}
	if !data.Usemip.Equal(state.Usemip) {
		if config.Usemip.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "usemip")
		} else {
			hasChange = true
		}
	}
	if !data.Userdomains.Equal(state.Userdomains) {
		if !config.Userdomains.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Wihome.Equal(state.Wihome) {
		if !config.Wihome.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Wihomeaddresstype.Equal(state.Wihomeaddresstype) {
		if !config.Wihomeaddresstype.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Windowsautologon.Equal(state.Windowsautologon) {
		if config.Windowsautologon.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "windowsautologon")
		} else {
			hasChange = true
		}
	}
	if !data.Windowsclienttype.Equal(state.Windowsclienttype) {
		if config.Windowsclienttype.IsNull() { // removed from config -> unset it (attribute has a server default)
			attributesToUnset = append(attributesToUnset, "windowsclienttype")
		} else {
			hasChange = true
		}
	}
	if !data.Windowspluginupgrade.Equal(state.Windowspluginupgrade) {
		if !config.Windowspluginupgrade.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Winsip.Equal(state.Winsip) {
		if !config.Winsip.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}
	if !data.Wiportalmode.Equal(state.Wiportalmode) {
		if !config.Wiportalmode.IsNull() { // only a real configured change triggers an update
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		vpnparameter := vpnparameterGetThePayloadFromtheConfig(ctx, &data)

		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Vpnparameter.Type(), &vpnparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnparameter, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vpnparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnparameter resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. vpnparameter is a singleton, so the unset payload
	// carries no identifying key.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Vpnparameter.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset vpnparameter attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readVpnparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnparameter resource")

	// For vpnparameter, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnparameter resource from state")
}

// Helper function to read vpnparameter data from API
func (r *VpnparameterResource) readVpnparameterFromApi(ctx context.Context, data *VpnparameterResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnparameter, got error: %s", err))
		return
	}

	vpnparameterSetAttrFromGet(ctx, data, getResponseData)

}
