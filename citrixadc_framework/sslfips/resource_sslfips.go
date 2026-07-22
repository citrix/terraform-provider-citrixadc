package sslfips

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
var _ resource.Resource = &SslfipsResource{}
var _ resource.ResourceWithConfigure = (*SslfipsResource)(nil)
var _ resource.ResourceWithImportState = (*SslfipsResource)(nil)
var _ resource.ResourceWithValidateConfig = (*SslfipsResource)(nil)

func NewSslfipsResource() resource.Resource {
	return &SslfipsResource{}
}

// SslfipsResource defines the resource implementation.
type SslfipsResource struct {
	client *service.NitroClient
}

func (r *SslfipsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslfipsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslfips"
}

func (r *SslfipsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslfipsResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data SslfipsResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that either oldsopassword or oldsopassword_wo is specified
	if data.Oldsopassword.IsNull() && data.OldsopasswordWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("oldsopassword"),
			"Missing Required Attribute",
			"Either \"oldsopassword\" or \"oldsopassword_wo\" must be specified.",
		)
	}

	// Validate that either sopassword or sopassword_wo is specified
	if data.Sopassword.IsNull() && data.SopasswordWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("sopassword"),
			"Missing Required Attribute",
			"Either \"sopassword\" or \"sopassword_wo\" must be specified.",
		)
	}

	// Validate that either userpassword or userpassword_wo is specified
	if data.Userpassword.IsNull() && data.UserpasswordWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("userpassword"),
			"Missing Required Attribute",
			"Either \"userpassword\" or \"userpassword_wo\" must be specified.",
		)
	}
}

func (r *SslfipsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config SslfipsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslfips resource")
	// Get payload from plan (regular attributes)
	sslfips := sslfipsGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	sslfipsGetThePayloadFromtheConfig(ctx, &config, &sslfips)

	// Make API call
	// sslfips is a singleton configured via the PUT `set` endpoint (HTTP Method: PUT
	// for the `add`/set operation). There is no NITRO `add` or `delete` verb.
	// WARNING: DISRUPTIVE. Setting inithsm performs HSM initialization which ERASES
	// all FIPS key/cert data on the appliance and requires dedicated FIPS hardware.
	// This is unsupported on non-FIPS VPX appliances.
	_, err := r.client.UpdateResource(service.Sslfips.Type(), "", &sslfips)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslfips, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslfips resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("sslfips-config")

	// Read the updated state back
	r.readSslfipsFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslfipsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslfips resource")

	r.readSslfipsFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SslfipsResourceModel

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

	tflog.Debug(ctx, "Updating sslfips resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Fipsfw.Equal(state.Fipsfw) {
		tflog.Debug(ctx, fmt.Sprintf("fipsfw has changed for sslfips"))
		hasChange = true
	}
	if !data.Hsmlabel.Equal(state.Hsmlabel) {
		tflog.Debug(ctx, fmt.Sprintf("hsmlabel has changed for sslfips"))
		hasChange = true
	}
	if !data.Inithsm.Equal(state.Inithsm) {
		tflog.Debug(ctx, fmt.Sprintf("inithsm has changed for sslfips"))
		hasChange = true
	}
	// Check secret attribute oldsopassword or its version tracker
	if !data.Oldsopassword.Equal(state.Oldsopassword) {
		tflog.Debug(ctx, fmt.Sprintf("oldsopassword has changed for sslfips"))
		hasChange = true
	} else if !data.OldsopasswordWoVersion.Equal(state.OldsopasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("oldsopassword_wo_version has changed for sslfips"))
		hasChange = true
	}
	// Check secret attribute sopassword or its version tracker
	if !data.Sopassword.Equal(state.Sopassword) {
		tflog.Debug(ctx, fmt.Sprintf("sopassword has changed for sslfips"))
		hasChange = true
	} else if !data.SopasswordWoVersion.Equal(state.SopasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("sopassword_wo_version has changed for sslfips"))
		hasChange = true
	}
	// Check secret attribute userpassword or its version tracker
	if !data.Userpassword.Equal(state.Userpassword) {
		tflog.Debug(ctx, fmt.Sprintf("userpassword has changed for sslfips"))
		hasChange = true
	} else if !data.UserpasswordWoVersion.Equal(state.UserpasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("userpassword_wo_version has changed for sslfips"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		sslfips := sslfipsGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		sslfipsGetThePayloadFromtheConfig(ctx, &config, &sslfips)
		// Make API call
		// Singleton set via PUT (HTTP Method: PUT for the sslfips set endpoint).
		_, err := r.client.UpdateResource(service.Sslfips.Type(), "", &sslfips)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslfips, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslfips resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslfips resource, skipping update")
	}

	// Read the updated state back
	r.readSslfipsFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslfipsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslfips resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed sslfips from Terraform state")
}

// Helper function to read sslfips data from API
func (r *SslfipsResource) readSslfipsFromApi(ctx context.Context, data *SslfipsResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Sslfips.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslfips, got error: %s", err))
		return
	}

	sslfipsSetAttrFromGet(ctx, data, getResponseData)

}
