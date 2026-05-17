package dnskey

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
var _ resource.Resource = &DnskeyResource{}
var _ resource.ResourceWithConfigure = (*DnskeyResource)(nil)
var _ resource.ResourceWithImportState = (*DnskeyResource)(nil)

func NewDnskeyResource() resource.Resource {
	return &DnskeyResource{}
}

// DnskeyResource defines the resource implementation.
type DnskeyResource struct {
	client *service.NitroClient
}

func (r *DnskeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnskeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnskey"
}

func (r *DnskeyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnskeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config DnskeyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnskey resource")
	// Get payload from plan (regular attributes)
	dnskey := dnskeyGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	dnskeyGetThePayloadFromtheConfig(ctx, &config, &dnskey)

	// Make API call
	// Named resource - use AddResource
	keyname_value := data.Keyname.ValueString()
	_, err := r.client.AddResource(service.Dnskey.Type(), keyname_value, &dnskey)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnskey, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dnskey resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Keyname.ValueString()))

	// Read the updated state back
	r.readDnskeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnskeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnskeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnskey resource")

	r.readDnskeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnskeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state DnskeyResourceModel

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

	tflog.Debug(ctx, "Updating dnskey resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Autorollover.Equal(state.Autorollover) {
		tflog.Debug(ctx, fmt.Sprintf("autorollover has changed for dnskey"))
		hasChange = true
	}
	if !data.Expires.Equal(state.Expires) {
		tflog.Debug(ctx, fmt.Sprintf("expires has changed for dnskey"))
		hasChange = true
	}
	if !data.Notificationperiod.Equal(state.Notificationperiod) {
		tflog.Debug(ctx, fmt.Sprintf("notificationperiod has changed for dnskey"))
		hasChange = true
	}
	// Check secret attribute password or its version tracker
	if !data.Password.Equal(state.Password) {
		tflog.Debug(ctx, fmt.Sprintf("password has changed for dnskey"))
		hasChange = true
	} else if !data.PasswordWoVersion.Equal(state.PasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("password_wo_version has changed for dnskey"))
		hasChange = true
	}
	if !data.Revoke.Equal(state.Revoke) {
		tflog.Debug(ctx, fmt.Sprintf("revoke has changed for dnskey"))
		hasChange = true
	}
	if !data.Rollovermethod.Equal(state.Rollovermethod) {
		tflog.Debug(ctx, fmt.Sprintf("rollovermethod has changed for dnskey"))
		hasChange = true
	}
	if !data.Ttl.Equal(state.Ttl) {
		tflog.Debug(ctx, fmt.Sprintf("ttl has changed for dnskey"))
		hasChange = true
	}
	if !data.Units1.Equal(state.Units1) {
		tflog.Debug(ctx, fmt.Sprintf("units1 has changed for dnskey"))
		hasChange = true
	}
	if !data.Units2.Equal(state.Units2) {
		tflog.Debug(ctx, fmt.Sprintf("units2 has changed for dnskey"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		dnskey := dnskeyGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		dnskeyGetThePayloadFromtheConfig(ctx, &config, &dnskey)
		// Make API call
		// Named resource - use UpdateResource
		keyname_value := data.Keyname.ValueString()
		_, err := r.client.UpdateResource(service.Dnskey.Type(), keyname_value, &dnskey)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dnskey, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated dnskey resource")
	} else {
		tflog.Debug(ctx, "No changes detected for dnskey resource, skipping update")
	}

	// Read the updated state back
	r.readDnskeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnskeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnskeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnskey resource")
	// Named resource - delete using DeleteResource
	keyname_value := data.Keyname.ValueString()
	err := r.client.DeleteResource(service.Dnskey.Type(), keyname_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnskey, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnskey resource")
}

// Helper function to read dnskey data from API
func (r *DnskeyResource) readDnskeyFromApi(ctx context.Context, data *DnskeyResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	keyname_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Dnskey.Type(), keyname_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnskey, got error: %s", err))
		return
	}

	dnskeySetAttrFromGet(ctx, data, getResponseData)

}
