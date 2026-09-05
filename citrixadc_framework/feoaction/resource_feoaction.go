package feoaction

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
var _ resource.Resource = &FeoactionResource{}
var _ resource.ResourceWithConfigure = (*FeoactionResource)(nil)
var _ resource.ResourceWithImportState = (*FeoactionResource)(nil)

func NewFeoactionResource() resource.Resource {
	return &FeoactionResource{}
}

// FeoactionResource defines the resource implementation.
type FeoactionResource struct {
	client *service.NitroClient
}

func (r *FeoactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *FeoactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_feoaction"
}

func (r *FeoactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *FeoactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data FeoactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating feoaction resource")

	// Create API request body from the model
	feoaction := feoactionGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	feoactionName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Feoaction.Type(), feoactionName, &feoaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create feoaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created feoaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(feoactionName)

	// Read the updated state back
	if !r.readFeoactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "feoaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FeoactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data FeoactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading feoaction resource")

	found := r.readFeoactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *FeoactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state FeoactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (-> unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating feoaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Cachemaxage.Equal(state.Cachemaxage) {
		tflog.Debug(ctx, "cachemaxage has changed for feoaction")
		hasChange = true
	}
	if !data.Clientsidemeasurements.Equal(state.Clientsidemeasurements) {
		tflog.Debug(ctx, "clientsidemeasurements has changed for feoaction")
		if config.Clientsidemeasurements.IsNull() {
			attributesToUnset = append(attributesToUnset, "clientsidemeasurements")
		} else {
			hasChange = true
		}
	}
	if !data.Convertimporttolink.Equal(state.Convertimporttolink) {
		tflog.Debug(ctx, "convertimporttolink has changed for feoaction")
		if config.Convertimporttolink.IsNull() {
			attributesToUnset = append(attributesToUnset, "convertimporttolink")
		} else {
			hasChange = true
		}
	}
	if !data.Csscombine.Equal(state.Csscombine) {
		tflog.Debug(ctx, "csscombine has changed for feoaction")
		if config.Csscombine.IsNull() {
			attributesToUnset = append(attributesToUnset, "csscombine")
		} else {
			hasChange = true
		}
	}
	if !data.Cssimginline.Equal(state.Cssimginline) {
		tflog.Debug(ctx, "cssimginline has changed for feoaction")
		if config.Cssimginline.IsNull() {
			attributesToUnset = append(attributesToUnset, "cssimginline")
		} else {
			hasChange = true
		}
	}
	if !data.Cssinline.Equal(state.Cssinline) {
		tflog.Debug(ctx, "cssinline has changed for feoaction")
		if config.Cssinline.IsNull() {
			attributesToUnset = append(attributesToUnset, "cssinline")
		} else {
			hasChange = true
		}
	}
	if !data.Cssminify.Equal(state.Cssminify) {
		tflog.Debug(ctx, "cssminify has changed for feoaction")
		if config.Cssminify.IsNull() {
			attributesToUnset = append(attributesToUnset, "cssminify")
		} else {
			hasChange = true
		}
	}
	if !data.Cssmovetohead.Equal(state.Cssmovetohead) {
		tflog.Debug(ctx, "cssmovetohead has changed for feoaction")
		if config.Cssmovetohead.IsNull() {
			attributesToUnset = append(attributesToUnset, "cssmovetohead")
		} else {
			hasChange = true
		}
	}
	if !data.Dnsshards.Equal(state.Dnsshards) {
		tflog.Debug(ctx, "dnsshards has changed for feoaction")
		hasChange = true
	}
	if !data.Domainsharding.Equal(state.Domainsharding) {
		tflog.Debug(ctx, "domainsharding has changed for feoaction")
		if config.Domainsharding.IsNull() {
			attributesToUnset = append(attributesToUnset, "domainsharding")
		} else {
			hasChange = true
		}
	}
	if !data.Htmlminify.Equal(state.Htmlminify) {
		tflog.Debug(ctx, "htmlminify has changed for feoaction")
		if config.Htmlminify.IsNull() {
			attributesToUnset = append(attributesToUnset, "htmlminify")
		} else {
			hasChange = true
		}
	}
	if !data.Imggiftopng.Equal(state.Imggiftopng) {
		tflog.Debug(ctx, "imggiftopng has changed for feoaction")
		if config.Imggiftopng.IsNull() {
			attributesToUnset = append(attributesToUnset, "imggiftopng")
		} else {
			hasChange = true
		}
	}
	if !data.Imginline.Equal(state.Imginline) {
		tflog.Debug(ctx, "imginline has changed for feoaction")
		if config.Imginline.IsNull() {
			attributesToUnset = append(attributesToUnset, "imginline")
		} else {
			hasChange = true
		}
	}
	if !data.Imglazyload.Equal(state.Imglazyload) {
		tflog.Debug(ctx, "imglazyload has changed for feoaction")
		if config.Imglazyload.IsNull() {
			attributesToUnset = append(attributesToUnset, "imglazyload")
		} else {
			hasChange = true
		}
	}
	if !data.Imgshrinktoattrib.Equal(state.Imgshrinktoattrib) {
		tflog.Debug(ctx, "imgshrinktoattrib has changed for feoaction")
		if config.Imgshrinktoattrib.IsNull() {
			attributesToUnset = append(attributesToUnset, "imgshrinktoattrib")
		} else {
			hasChange = true
		}
	}
	if !data.Imgtojpegxr.Equal(state.Imgtojpegxr) {
		tflog.Debug(ctx, "imgtojpegxr has changed for feoaction")
		if config.Imgtojpegxr.IsNull() {
			attributesToUnset = append(attributesToUnset, "imgtojpegxr")
		} else {
			hasChange = true
		}
	}
	if !data.Imgtowebp.Equal(state.Imgtowebp) {
		tflog.Debug(ctx, "imgtowebp has changed for feoaction")
		if config.Imgtowebp.IsNull() {
			attributesToUnset = append(attributesToUnset, "imgtowebp")
		} else {
			hasChange = true
		}
	}
	if !data.Jpgoptimize.Equal(state.Jpgoptimize) {
		tflog.Debug(ctx, "jpgoptimize has changed for feoaction")
		if config.Jpgoptimize.IsNull() {
			attributesToUnset = append(attributesToUnset, "jpgoptimize")
		} else {
			hasChange = true
		}
	}
	if !data.Jsinline.Equal(state.Jsinline) {
		tflog.Debug(ctx, "jsinline has changed for feoaction")
		if config.Jsinline.IsNull() {
			attributesToUnset = append(attributesToUnset, "jsinline")
		} else {
			hasChange = true
		}
	}
	if !data.Jsminify.Equal(state.Jsminify) {
		tflog.Debug(ctx, "jsminify has changed for feoaction")
		if config.Jsminify.IsNull() {
			attributesToUnset = append(attributesToUnset, "jsminify")
		} else {
			hasChange = true
		}
	}
	if !data.Jsmovetoend.Equal(state.Jsmovetoend) {
		tflog.Debug(ctx, "jsmovetoend has changed for feoaction")
		if config.Jsmovetoend.IsNull() {
			attributesToUnset = append(attributesToUnset, "jsmovetoend")
		} else {
			hasChange = true
		}
	}
	if !data.Pageextendcache.Equal(state.Pageextendcache) {
		tflog.Debug(ctx, "pageextendcache has changed for feoaction")
		if config.Pageextendcache.IsNull() {
			attributesToUnset = append(attributesToUnset, "pageextendcache")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		feoaction := feoactionGetThePayloadFromthePlan(ctx, &data)
		// Named resource - use UpdateResource
		feoactionName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Feoaction.Type(), feoactionName, &feoaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update feoaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated feoaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for feoaction resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Feoaction.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset feoaction attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readFeoactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "feoaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FeoactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FeoactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting feoaction resource")

	// Named resource - delete using DeleteResource keyed on the resource ID (name)
	feoactionName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Feoaction.Type(), feoactionName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete feoaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted feoaction resource")
}

// Helper function to read feoaction data from API.
// Returns false (without error) when the resource no longer exists on the ADC.
func (r *FeoactionResource) readFeoactionFromApi(ctx context.Context, data *FeoactionResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain value (name)
	feoactionName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Feoaction.Type(), feoactionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read feoaction, got error: %s", err))
		return false
	}

	feoactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
