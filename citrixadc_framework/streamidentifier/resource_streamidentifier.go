package streamidentifier

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
var _ resource.Resource = &StreamidentifierResource{}
var _ resource.ResourceWithConfigure = (*StreamidentifierResource)(nil)
var _ resource.ResourceWithImportState = (*StreamidentifierResource)(nil)

func NewStreamidentifierResource() resource.Resource {
	return &StreamidentifierResource{}
}

// StreamidentifierResource defines the resource implementation.
type StreamidentifierResource struct {
	client *service.NitroClient
}

func (r *StreamidentifierResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *StreamidentifierResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_streamidentifier"
}

func (r *StreamidentifierResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *StreamidentifierResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data StreamidentifierResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating streamidentifier resource")

	// Create API request body from the plan
	streamidentifier := streamidentifierGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	streamidentifierName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Streamidentifier.Type(), streamidentifierName, &streamidentifier)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create streamidentifier, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created streamidentifier resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(streamidentifierName)

	// Read the updated state back
	if !r.readStreamidentifierFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "streamidentifier not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StreamidentifierResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data StreamidentifierResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading streamidentifier resource")

	found := r.readStreamidentifierFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *StreamidentifierResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state StreamidentifierResourceModel

	// Read Terraform prior state to preserve ID and perform change detection
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating streamidentifier resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Acceptancethreshold.Equal(state.Acceptancethreshold) {
		tflog.Debug(ctx, "acceptancethreshold has changed for streamidentifier")
		hasChange = true
	}
	if !data.Appflowlog.Equal(state.Appflowlog) {
		tflog.Debug(ctx, "appflowlog has changed for streamidentifier")
		hasChange = true
	}
	if !data.Breachthreshold.Equal(state.Breachthreshold) {
		tflog.Debug(ctx, "breachthreshold has changed for streamidentifier")
		hasChange = true
	}
	if !data.Interval.Equal(state.Interval) {
		tflog.Debug(ctx, "interval has changed for streamidentifier")
		hasChange = true
	}
	if !data.Log.Equal(state.Log) {
		tflog.Debug(ctx, "log has changed for streamidentifier")
		hasChange = true
	}
	if !data.Loginterval.Equal(state.Loginterval) {
		tflog.Debug(ctx, "loginterval has changed for streamidentifier")
		hasChange = true
	}
	if !data.Loglimit.Equal(state.Loglimit) {
		tflog.Debug(ctx, "loglimit has changed for streamidentifier")
		hasChange = true
	}
	if !data.Maxtransactionthreshold.Equal(state.Maxtransactionthreshold) {
		tflog.Debug(ctx, "maxtransactionthreshold has changed for streamidentifier")
		hasChange = true
	}
	if !data.Mintransactionthreshold.Equal(state.Mintransactionthreshold) {
		tflog.Debug(ctx, "mintransactionthreshold has changed for streamidentifier")
		hasChange = true
	}
	if !data.Samplecount.Equal(state.Samplecount) {
		tflog.Debug(ctx, "samplecount has changed for streamidentifier")
		hasChange = true
	}
	if !data.Selectorname.Equal(state.Selectorname) {
		tflog.Debug(ctx, "selectorname has changed for streamidentifier")
		hasChange = true
	}
	if !data.Snmptrap.Equal(state.Snmptrap) {
		tflog.Debug(ctx, "snmptrap has changed for streamidentifier")
		hasChange = true
	}
	if !data.Sort.Equal(state.Sort) {
		tflog.Debug(ctx, "sort has changed for streamidentifier")
		hasChange = true
	}
	if !data.Trackackonlypackets.Equal(state.Trackackonlypackets) {
		tflog.Debug(ctx, "trackackonlypackets has changed for streamidentifier")
		hasChange = true
	}
	if !data.Tracktransactions.Equal(state.Tracktransactions) {
		tflog.Debug(ctx, "tracktransactions has changed for streamidentifier")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the plan
		streamidentifier := streamidentifierGetThePayloadFromthePlan(ctx, &data)
		// streamidentifier update is an unnamed PUT (SDK v2 parity: UpdateUnnamedResource)
		err := r.client.UpdateUnnamedResource(service.Streamidentifier.Type(), &streamidentifier)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update streamidentifier, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated streamidentifier resource")
	} else {
		tflog.Debug(ctx, "No changes detected for streamidentifier resource, skipping update")
	}

	// Read the updated state back
	if !r.readStreamidentifierFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "streamidentifier not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StreamidentifierResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data StreamidentifierResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting streamidentifier resource")

	// Named resource - delete using DeleteResource
	streamidentifierName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Streamidentifier.Type(), streamidentifierName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete streamidentifier, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted streamidentifier resource")
}

// Helper function to read streamidentifier data from API
func (r *StreamidentifierResource) readStreamidentifierFromApi(ctx context.Context, data *StreamidentifierResourceModel, diags *diag.Diagnostics) bool {

	// Named resource - find by its single unique attribute (the ID holds the plain name value)
	streamidentifierName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Streamidentifier.Type(), streamidentifierName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read streamidentifier, got error: %s", err))
		return false
	}

	streamidentifierSetAttrFromGet(ctx, data, getResponseData)

	return true
}
