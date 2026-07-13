package sslcrl

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
var _ resource.Resource = &SslcrlResource{}
var _ resource.ResourceWithConfigure = (*SslcrlResource)(nil)
var _ resource.ResourceWithImportState = (*SslcrlResource)(nil)

func NewSslcrlResource() resource.Resource {
	return &SslcrlResource{}
}

// SslcrlResource defines the resource implementation.
type SslcrlResource struct {
	client *service.NitroClient
}

func (r *SslcrlResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslcrlResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcrl"
}

func (r *SslcrlResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslcrlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config SslcrlResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslcrl resource")
	// Get payload from plan (regular attributes)
	sslcrl := sslcrlGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	sslcrlGetThePayloadFromtheConfig(ctx, &config, &sslcrl)

	// Make API call
	// Named resource - use AddResource
	crlname_value := data.Crlname.ValueString()
	_, err := r.client.AddResource(service.Sslcrl.Type(), crlname_value, &sslcrl)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslcrl, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslcrl resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Crlname.ValueString()))

	// Read the updated state back
	if !r.readSslcrlFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslcrl not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcrlResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslcrlResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslcrl resource")

	found := r.readSslcrlFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SslcrlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SslcrlResourceModel

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

	tflog.Debug(ctx, "Updating sslcrl resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Basedn.Equal(state.Basedn) {
		tflog.Debug(ctx, fmt.Sprintf("basedn has changed for sslcrl"))
		hasChange = true
	}
	if !data.Binary.Equal(state.Binary) {
		tflog.Debug(ctx, fmt.Sprintf("binary has changed for sslcrl"))
		hasChange = true
	}
	if !data.Binddn.Equal(state.Binddn) {
		tflog.Debug(ctx, fmt.Sprintf("binddn has changed for sslcrl"))
		hasChange = true
	}
	if !data.Cacert.Equal(state.Cacert) {
		tflog.Debug(ctx, fmt.Sprintf("cacert has changed for sslcrl"))
		hasChange = true
	}
	if !data.Day.Equal(state.Day) {
		tflog.Debug(ctx, fmt.Sprintf("day has changed for sslcrl"))
		hasChange = true
	}
	if !data.Interval.Equal(state.Interval) {
		tflog.Debug(ctx, fmt.Sprintf("interval has changed for sslcrl"))
		hasChange = true
	}
	if !data.Method.Equal(state.Method) {
		tflog.Debug(ctx, fmt.Sprintf("method has changed for sslcrl"))
		hasChange = true
	}
	// Check secret attribute password or its version tracker
	if !data.Password.Equal(state.Password) {
		tflog.Debug(ctx, fmt.Sprintf("password has changed for sslcrl"))
		hasChange = true
	} else if !data.PasswordWoVersion.Equal(state.PasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("password_wo_version has changed for sslcrl"))
		hasChange = true
	}
	if !data.Port.Equal(state.Port) {
		tflog.Debug(ctx, fmt.Sprintf("port has changed for sslcrl"))
		hasChange = true
	}
	if !data.Refresh.Equal(state.Refresh) {
		tflog.Debug(ctx, fmt.Sprintf("refresh has changed for sslcrl"))
		hasChange = true
	}
	if !data.Scope.Equal(state.Scope) {
		tflog.Debug(ctx, fmt.Sprintf("scope has changed for sslcrl"))
		hasChange = true
	}
	if !data.Server.Equal(state.Server) {
		tflog.Debug(ctx, fmt.Sprintf("server has changed for sslcrl"))
		hasChange = true
	}
	if !data.Time.Equal(state.Time) {
		tflog.Debug(ctx, fmt.Sprintf("time has changed for sslcrl"))
		hasChange = true
	}
	if !data.Url.Equal(state.Url) {
		tflog.Debug(ctx, fmt.Sprintf("url has changed for sslcrl"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		sslcrl := sslcrlGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		sslcrlGetThePayloadFromtheConfig(ctx, &config, &sslcrl)
		// Make API call
		// Named resource - use UpdateResource
		crlname_value := data.Crlname.ValueString()
		_, err := r.client.UpdateResource(service.Sslcrl.Type(), crlname_value, &sslcrl)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslcrl, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslcrl resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslcrl resource, skipping update")
	}

	// Read the updated state back
	if !r.readSslcrlFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslcrl not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcrlResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslcrlResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslcrl resource")
	// Named resource - delete using DeleteResource
	crlname_value := data.Crlname.ValueString()
	err := r.client.DeleteResource(service.Sslcrl.Type(), crlname_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslcrl, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslcrl resource")
}

// Helper function to read sslcrl data from API
func (r *SslcrlResource) readSslcrlFromApi(ctx context.Context, data *SslcrlResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	crlname_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Sslcrl.Type(), crlname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslcrl, got error: %s", err))
		return false
	}

	sslcrlSetAttrFromGet(ctx, data, getResponseData)

	return true
}
