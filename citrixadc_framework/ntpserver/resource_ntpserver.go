package ntpserver

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
var _ resource.Resource = &NtpserverResource{}
var _ resource.ResourceWithConfigure = (*NtpserverResource)(nil)
var _ resource.ResourceWithImportState = (*NtpserverResource)(nil)

func NewNtpserverResource() resource.Resource {
	return &NtpserverResource{}
}

// NtpserverResource defines the resource implementation.
type NtpserverResource struct {
	client *service.NitroClient
}

func (r *NtpserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NtpserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ntpserver"
}

func (r *NtpserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NtpserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NtpserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ntpserver resource")

	// Build the identifier (serverip preferred, otherwise servername) - SDK v2 parity.
	identifier := ntpserverIdentifier(&data)
	if identifier == "" {
		resp.Diagnostics.AddError("Configuration Error", "At least one of serverip or servername must be specified")
		return
	}

	ntpserver := ntpserverGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	_, err := r.client.AddResource(service.Ntpserver.Type(), identifier, &ntpserver)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ntpserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created ntpserver resource")

	// SDK v2 ID scheme: plain identifier value.
	data.Id = types.StringValue(identifier)

	// Read the updated state back
	if !r.readNtpserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "ntpserver not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NtpserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NtpserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading ntpserver resource")

	found := r.readNtpserverFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NtpserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NtpserverResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config so attributes removed from config can be unset
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (serverip/servername are ForceNew and unchanged).
	data.Id = state.Id

	tflog.Debug(ctx, "Updating ntpserver resource")

	// Detect changes and, for attributes removed from config, mark them for unset
	// so the appliance reverts them to their NITRO defaults.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Autokey.Equal(state.Autokey) {
		if config.Autokey.IsNull() {
			attributesToUnset = append(attributesToUnset, "autokey")
		} else {
			hasChange = true
		}
	}
	if !data.Maxpoll.Equal(state.Maxpoll) {
		if config.Maxpoll.IsNull() {
			attributesToUnset = append(attributesToUnset, "maxpoll")
		} else {
			hasChange = true
		}
	}
	if !data.Minpoll.Equal(state.Minpoll) {
		if config.Minpoll.IsNull() {
			attributesToUnset = append(attributesToUnset, "minpoll")
		} else {
			hasChange = true
		}
	}
	if !data.Preferredntpserver.Equal(state.Preferredntpserver) {
		if config.Preferredntpserver.IsNull() {
			attributesToUnset = append(attributesToUnset, "preferredntpserver")
		} else {
			hasChange = true
		}
	}
	if !data.Key.Equal(state.Key) {
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		ntpserver := ntpserverGetThePayloadFromtheConfig(ctx, &data)

		// Singleton-style update: PUT with serverip/servername in the payload (SDK v2 parity).
		err := r.client.UpdateUnnamedResource(service.Ntpserver.Type(), &ntpserver)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ntpserver, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated ntpserver resource")
	} else {
		tflog.Debug(ctx, "No changes detected for ntpserver resource, skipping update")
	}

	// Unset attributes removed from config so the appliance reverts them to
	// their defaults. The unset key is serverip/servername (the resource key).
	unsetIdPayload := map[string]interface{}{}
	if !data.Serverip.IsNull() && !data.Serverip.IsUnknown() && data.Serverip.ValueString() != "" {
		unsetIdPayload["serverip"] = data.Serverip.ValueString()
	}
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() && data.Servername.ValueString() != "" {
		unsetIdPayload["servername"] = data.Servername.ValueString()
	}
	if err := utils.ExecuteUnset(r.client, service.Ntpserver.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset ntpserver attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readNtpserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "ntpserver not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NtpserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NtpserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting ntpserver resource")

	// The ID holds the plain serverip/servername value (SDK v2 parity).
	identifier := data.Id.ValueString()
	if identifier == "" {
		identifier = ntpserverIdentifier(&data)
	}

	if identifier != "" {
		err := r.client.DeleteResource(service.Ntpserver.Type(), identifier)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ntpserver, got error: %s", err))
			return
		}
	}

	tflog.Trace(ctx, "Deleted ntpserver resource")
}

// ntpserverIdentifier returns the plain-value identifier: serverip if set,
// otherwise servername (matches the SDK v2 ID scheme).
func ntpserverIdentifier(data *NtpserverResourceModel) string {
	if !data.Serverip.IsNull() && !data.Serverip.IsUnknown() && data.Serverip.ValueString() != "" {
		return data.Serverip.ValueString()
	}
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() && data.Servername.ValueString() != "" {
		return data.Servername.ValueString()
	}
	return ""
}

// readNtpserverFromApi reads ntpserver state from the ADC. It uses FindAllResources
// + a serverip/servername filter (SDK v2 parity, since ntpserver only exposes
// "get (all)"). Returns false when the resource is absent so the caller can remove
// it from state.
func (r *NtpserverResource) readNtpserverFromApi(ctx context.Context, data *NtpserverResourceModel, diags *diag.Diagnostics) bool {
	// The ID is the plain serverip/servername value.
	identifier := data.Id.ValueString()
	if identifier == "" {
		identifier = ntpserverIdentifier(data)
	}
	if identifier == "" {
		diags.AddError("Configuration Error", "At least one of serverip or servername must be specified")
		return false
	}

	dataArr, err := r.client.FindAllResources(service.Ntpserver.Type())
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read ntpserver, got error: %s", err))
		return false
	}

	foundIndex := -1
	for i, v := range dataArr {
		if v["serverip"] == identifier || v["servername"] == identifier {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		return false
	}

	ntpserverSetAttrFromGet(ctx, data, dataArr[foundIndex])
	return true
}
