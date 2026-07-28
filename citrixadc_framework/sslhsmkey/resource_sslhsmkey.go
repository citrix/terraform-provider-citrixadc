package sslhsmkey

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
var _ resource.Resource = &SslhsmkeyResource{}
var _ resource.ResourceWithConfigure = (*SslhsmkeyResource)(nil)
var _ resource.ResourceWithImportState = (*SslhsmkeyResource)(nil)

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
	var data, config, state SslhsmkeyResourceModel

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

	tflog.Debug(ctx, "Updating sslhsmkey resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	// Check secret attribute password or its version tracker
	if !data.Password.Equal(state.Password) {
		tflog.Debug(ctx, fmt.Sprintf("password has changed for sslhsmkey"))
		hasChange = true
	} else if !data.PasswordWoVersion.Equal(state.PasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("password_wo_version has changed for sslhsmkey"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		sslhsmkey := sslhsmkeyGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		sslhsmkeyGetThePayloadFromtheConfig(ctx, &config, &sslhsmkey)
		// Make API call
		// Named resource - use UpdateResource
		hsmkeyname_value := data.Hsmkeyname.ValueString()
		_, err := r.client.UpdateResource(service.Sslhsmkey.Type(), hsmkeyname_value, &sslhsmkey)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslhsmkey, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslhsmkey resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslhsmkey resource, skipping update")
	}

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
	argsMap := make(map[string]string)
	if !data.Hsmtype.IsNull() && data.Hsmtype.ValueString() != "" {
		argsMap["hsmtype"] = data.Hsmtype.ValueString()
	}
	if !data.Key.IsNull() && data.Key.ValueString() != "" {
		argsMap["key"] = data.Key.ValueString()
	}
	if !data.Keystore.IsNull() && data.Keystore.ValueString() != "" {
		argsMap["keystore"] = data.Keystore.ValueString()
	}
	if !data.Password.IsNull() && data.Password.ValueString() != "" {
		argsMap["password"] = data.Password.ValueString()
	}
	if !data.Serialnum.IsNull() && data.Serialnum.ValueString() != "" {
		argsMap["serialnum"] = data.Serialnum.ValueString()
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
