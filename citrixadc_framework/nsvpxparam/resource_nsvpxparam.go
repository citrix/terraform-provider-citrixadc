package nsvpxparam

import (
	"context"
	"fmt"
	"strconv"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NsvpxparamResource{}
var _ resource.ResourceWithConfigure = (*NsvpxparamResource)(nil)
var _ resource.ResourceWithImportState = (*NsvpxparamResource)(nil)

func NewNsvpxparamResource() resource.Resource {
	return &NsvpxparamResource{}
}

// NsvpxparamResource defines the resource implementation.
type NsvpxparamResource struct {
	client *service.NitroClient
}

func (r *NsvpxparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsvpxparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsvpxparam"
}

func (r *NsvpxparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsvpxparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsvpxparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsvpxparam resource")

	nsvpxparam := nsvpxparamGetThePayloadFromtheConfig(ctx, &data)

	// nsvpxparam is a singleton/unnamed configuration resource.
	err := r.client.UpdateUnnamedResource(service.Nsvpxparam.Type(), &nsvpxparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsvpxparam, got error: %s", err))
		return
	}

	// ID scheme mirrors SDK v2: on a cluster, ownernode selects the node the
	// settings apply to and is the resource's identity; on a standalone VPX
	// ownernode is unconfigured and there is a single implicit entry. Encoding
	// the decision in the ID lets Read select the correct node during a refresh,
	// when the raw config is not available.
	if !data.Ownernode.IsNull() && !data.Ownernode.IsUnknown() {
		data.Id = types.StringValue(strconv.Itoa(int(data.Ownernode.ValueInt64())))
	} else {
		data.Id = types.StringValue("nsvpxparam-config")
	}

	tflog.Trace(ctx, "Created nsvpxparam resource")

	// Read the updated state back
	r.readNsvpxparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsvpxparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsvpxparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsvpxparam resource")

	r.readNsvpxparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsvpxparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NsvpxparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nsvpxparam resource")

	// All configurable attributes are RequiresReplaceIfConfigured (SDK v2
	// ForceNew), so a configured change is handled by recreate rather than
	// Update. This branch still pushes the config to keep any computed-value
	// resolution consistent.
	nsvpxparam := nsvpxparamGetThePayloadFromtheConfig(ctx, &data)

	err := r.client.UpdateUnnamedResource(service.Nsvpxparam.Type(), &nsvpxparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsvpxparam, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated nsvpxparam resource")

	// Read the updated state back (data.Id carried over from plan/prior state)
	r.readNsvpxparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsvpxparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsvpxparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsvpxparam resource")

	// nsvpxparam is a global configuration resource: it cannot actually be
	// deleted, only reset. Matching SDK v2, Delete just removes the reference
	// from Terraform state (the framework drops the resource after this returns).
	tflog.Trace(ctx, "Deleted nsvpxparam resource from state")
}

// readNsvpxparamFromApi reads the nsvpxparam config from the ADC and populates
// data. It mirrors the SDK v2 Read: fetch the array, then select the entry that
// matches the ownernode encoded in the ID (numeric ID => cluster node match;
// non-numeric ID => standalone single entry at index 0).
func (r *NsvpxparamResource) readNsvpxparamFromApi(ctx context.Context, data *NsvpxparamResourceModel, diags *diag.Diagnostics) {
	findParams := service.FindParams{
		ResourceType: "nsvpxparam",
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsvpxparam, got error: %s", err))
		return
	}

	if len(dataArr) == 0 {
		return
	}

	foundIndex := -1
	if node, convErr := strconv.Atoi(data.Id.ValueString()); convErr == nil {
		// Cluster mode: match by ownernode. NITRO returns ownernode as a number
		// (float64) or string; compare via %v.
		target := strconv.Itoa(node)
		for index, value := range dataArr {
			if fmt.Sprintf("%v", value["ownernode"]) == target {
				foundIndex = index
				break
			}
		}
	} else {
		// Standalone VPX: single implicit entry.
		foundIndex = 0
	}

	if foundIndex == -1 {
		return
	}

	nsvpxparamSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
