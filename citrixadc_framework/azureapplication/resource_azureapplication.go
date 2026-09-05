package azureapplication

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
var _ resource.Resource = &AzureapplicationResource{}
var _ resource.ResourceWithUpgradeState = &AzureapplicationResource{}
var _ resource.ResourceWithConfigure = (*AzureapplicationResource)(nil)
var _ resource.ResourceWithImportState = (*AzureapplicationResource)(nil)

func NewAzureapplicationResource() resource.Resource {
	return &AzureapplicationResource{}
}

// AzureapplicationResource defines the resource implementation.
type AzureapplicationResource struct {
	client *service.NitroClient
}

func (r *AzureapplicationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AzureapplicationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_azureapplication"
}

func (r *AzureapplicationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AzureapplicationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AzureapplicationResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating azureapplication resource")
	// Get payload from plan (regular attributes)
	azureapplication := azureapplicationGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	azureapplicationGetThePayloadFromtheConfig(ctx, &config, &azureapplication)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Azureapplication.Type(), name_value, &azureapplication)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create azureapplication, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created azureapplication resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAzureapplicationFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "azureapplication not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AzureapplicationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AzureapplicationResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading azureapplication resource")

	found := r.readAzureapplicationFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AzureapplicationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AzureapplicationResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating azureapplication resource")

	// No in-place update: every settable attribute of azureapplication is ForceNew
	// (RequiresReplace) — including the write-only clientsecret and its version
	// tracker — and NITRO exposes no "update" operation (only add/delete/get). Any
	// change (e.g. rotating clientsecret) triggers a destroy+recreate, so this method
	// is never invoked with a real diff. The prior UpdateResource(PUT) call was dead
	// code (a PUT would be rejected) and has been removed along with the now-unused
	// config read. If a future schema change makes an attribute non-ForceNew, add a
	// proper update path here.

	// Read the updated state back
	if !r.readAzureapplicationFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "azureapplication not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AzureapplicationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AzureapplicationResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting azureapplication resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Azureapplication.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete azureapplication, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted azureapplication resource")
}

// Helper function to read azureapplication data from API
func (r *AzureapplicationResource) readAzureapplicationFromApi(ctx context.Context, data *AzureapplicationResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Azureapplication.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read azureapplication, got error: %s", err))
		return false
	}

	azureapplicationSetAttrFromGet(ctx, data, getResponseData)

	return true
}

// UpgradeState migrates pre-write-only state (GH #1441): it seeds the
// "*_wo_version" tracker attribute(s) to 1 when the stored state has no value
// for them, so the schema Default does not plan a spurious "null -> 1" update
// after upgrading the provider. Paired with the schema Version bump so the
// upgrade path actually runs. See utils.WoVersionUpgradeState.
func (r *AzureapplicationResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	schemaResp := resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)
	return utils.WoVersionUpgradeState(schemaResp.Schema, func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
		var data AzureapplicationResourceModel
		resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if data.ClientsecretWoVersion.IsNull() {
			data.ClientsecretWoVersion = types.Int64Value(1)
		}
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	})
}
