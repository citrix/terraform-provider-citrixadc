package botprofile

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
var _ resource.Resource = &BotprofileResource{}
var _ resource.ResourceWithConfigure = (*BotprofileResource)(nil)
var _ resource.ResourceWithImportState = (*BotprofileResource)(nil)

func NewBotprofileResource() resource.Resource {
	return &BotprofileResource{}
}

// BotprofileResource defines the resource implementation.
type BotprofileResource struct {
	client *service.NitroClient
}

func (r *BotprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BotprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botprofile"
}

func (r *BotprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BotprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BotprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating botprofile resource")

	// Create API request body from the model
	botprofile := botprofileGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	botprofileName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Botprofile.Type(), botprofileName, &botprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create botprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created botprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(botprofileName)

	// Read the updated state back
	if !r.readBotprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "botprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BotprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading botprofile resource")

	found := r.readBotprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *BotprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state BotprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating botprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Addcookieflags.Equal(state.Addcookieflags) {
		hasChange = true
	}
	if !data.BotEnableBlackList.Equal(state.BotEnableBlackList) {
		hasChange = true
	}
	if !data.BotEnableIpReputation.Equal(state.BotEnableIpReputation) {
		hasChange = true
	}
	if !data.BotEnableRateLimit.Equal(state.BotEnableRateLimit) {
		hasChange = true
	}
	if !data.BotEnableTps.Equal(state.BotEnableTps) {
		hasChange = true
	}
	if !data.BotEnableWhiteList.Equal(state.BotEnableWhiteList) {
		hasChange = true
	}
	if !data.Clientipexpression.Equal(state.Clientipexpression) {
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		hasChange = true
	}
	if !data.Devicefingerprint.Equal(state.Devicefingerprint) {
		hasChange = true
	}
	if !data.Devicefingerprintaction.Equal(state.Devicefingerprintaction) {
		hasChange = true
	}
	if !data.Devicefingerprintmobile.Equal(state.Devicefingerprintmobile) {
		hasChange = true
	}
	if !data.Dfprequestlimit.Equal(state.Dfprequestlimit) {
		hasChange = true
	}
	if !data.Errorurl.Equal(state.Errorurl) {
		hasChange = true
	}
	if !data.Headlessbrowserdetection.Equal(state.Headlessbrowserdetection) {
		hasChange = true
	}
	if !data.Kmdetection.Equal(state.Kmdetection) {
		hasChange = true
	}
	if !data.Kmeventspostbodylimit.Equal(state.Kmeventspostbodylimit) {
		hasChange = true
	}
	if !data.Kmjavascriptname.Equal(state.Kmjavascriptname) {
		hasChange = true
	}
	if !data.Sessioncookiename.Equal(state.Sessioncookiename) {
		hasChange = true
	}
	if !data.Sessiontimeout.Equal(state.Sessiontimeout) {
		hasChange = true
	}
	if !data.Signature.Equal(state.Signature) {
		hasChange = true
	}
	if !data.Signaturemultipleuseragentheaderaction.Equal(state.Signaturemultipleuseragentheaderaction) {
		hasChange = true
	}
	if !data.Signaturenouseragentheaderaction.Equal(state.Signaturenouseragentheaderaction) {
		hasChange = true
	}
	if !data.Spoofedreqaction.Equal(state.Spoofedreqaction) {
		hasChange = true
	}
	if !data.Trap.Equal(state.Trap) {
		hasChange = true
	}
	if !data.Trapaction.Equal(state.Trapaction) {
		hasChange = true
	}
	if !data.Trapurl.Equal(state.Trapurl) {
		hasChange = true
	}
	if !data.Verboseloglevel.Equal(state.Verboseloglevel) {
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		botprofile := botprofileGetThePayloadFromthePlan(ctx, &data)
		// name is the identifier for the PUT body
		botprofile.Name = data.Name.ValueString()

		// Named resource - use UpdateResource
		botprofileName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Botprofile.Type(), botprofileName, &botprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update botprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated botprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for botprofile resource, skipping update")
	}

	// Read the updated state back
	if !r.readBotprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "botprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BotprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting botprofile resource")

	// Named resource - delete using DeleteResource
	botprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Botprofile.Type(), botprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete botprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted botprofile resource")
}

// Helper function to read botprofile data from API
func (r *BotprofileResource) readBotprofileFromApi(ctx context.Context, data *BotprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (name)
	botprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Botprofile.Type(), botprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read botprofile, got error: %s", err))
		return false
	}

	botprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
