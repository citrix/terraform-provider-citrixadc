package nsfeature

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NsfeatureResource{}
var _ resource.ResourceWithConfigure = (*NsfeatureResource)(nil)
var _ resource.ResourceWithImportState = (*NsfeatureResource)(nil)

func NewNsfeatureResource() resource.Resource {
	return &NsfeatureResource{}
}

// NsfeatureResource defines the resource implementation.
type NsfeatureResource struct {
	client *service.NitroClient
}

func (r *NsfeatureResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsfeatureResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsfeature"
}

func (r *NsfeatureResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsfeatureResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsfeatureResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsfeature resource")

	// Enable/disable the features the practitioner explicitly configured.
	r.applyFeatures(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate the (static) ID for this configuration resource.
	data.Id = types.StringValue("nsfeature-config")

	tflog.Trace(ctx, "Created nsfeature resource")

	// Read the updated state back so every (Optional+Computed) feature attribute
	// reflects the actual state on the ADC.
	r.readNsfeatureFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsfeatureResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsfeatureResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsfeature resource")

	r.readNsfeatureFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsfeatureResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NsfeatureResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read prior state to preserve the ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nsfeature resource")

	// Enable/disable the features the practitioner explicitly configured.
	r.applyFeatures(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "Updated nsfeature resource")

	// Read the updated state back
	r.readNsfeatureFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsfeatureResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsfeatureResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsfeature resource")

	// nsfeature is a global toggle set on the ADC; matching the SDK v2 behavior we
	// do not disable any features on delete, we simply drop the resource from
	// state (the framework removes it automatically once Delete returns).
	tflog.Trace(ctx, "Deleted nsfeature resource from state")
}

// applyFeatures enables/disables the features that were explicitly configured
// (known values) in the plan, leaving features the practitioner omitted
// (null/unknown, i.e. Computed) untouched on the ADC.
func (r *NsfeatureResource) applyFeatures(ctx context.Context, data *NsfeatureResourceModel, diags *diag.Diagnostics) {
	featureMap := nsfeatureModelToMap(data)

	enableList := make([]string, 0, len(featureList))
	disableList := make([]string, 0, len(featureList))

	for _, featureName := range featureList {
		val := featureMap[featureName]
		// Only act on features the practitioner explicitly configured. Omitted
		// Optional+Computed features are null (on create) or unknown (on update)
		// and must be left as-is on the ADC.
		if val.IsNull() || val.IsUnknown() {
			continue
		}
		if val.ValueBool() {
			enableList = append(enableList, featureName)
		} else {
			disableList = append(disableList, featureName)
		}
	}

	if len(enableList) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("Enabling nsfeatures %v", enableList))
		if err := r.client.EnableFeatures(enableList); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to enable nsfeatures %v, got error: %s", enableList, err))
			return
		}
	}

	if len(disableList) > 0 {
		tflog.Debug(ctx, fmt.Sprintf("Disabling nsfeatures %v", disableList))
		if err := r.client.DisableFeatures(disableList); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to disable nsfeatures %v, got error: %s", disableList, err))
			return
		}
	}
}

// readNsfeatureFromApi reads the currently-enabled features from the ADC and
// populates every feature attribute (true if enabled, false otherwise).
func (r *NsfeatureResource) readNsfeatureFromApi(ctx context.Context, data *NsfeatureResourceModel, diags *diag.Diagnostics) {
	tflog.Debug(ctx, "Reading nsfeature state from ADC")

	featuresData, err := r.client.ListEnabledFeatures()
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsfeature, got error: %s", err))
		return
	}

	// Normalize to lowercase to match the NITRO feature tokens.
	enabledFeatures := make([]string, len(featuresData))
	for i, val := range featuresData {
		enabledFeatures[i] = strings.ToLower(val)
	}

	nsfeatureSetAttrFromGet(ctx, data, enabledFeatures)
}
