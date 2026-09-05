package hafailover

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	sdkid "github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &HafailoverResource{}
var _ resource.ResourceWithConfigure = (*HafailoverResource)(nil)

func NewHafailoverResource() resource.Resource {
	return &HafailoverResource{}
}

// HafailoverResource defines the resource implementation.
type HafailoverResource struct {
	client *service.NitroClient
}

func (r *HafailoverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hafailover"
}

func (r *HafailoverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *HafailoverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data HafailoverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating hafailover resource")

	// Build the Force action payload from the plan
	hafailover := hafailoverGetThePayloadFromtheConfig(ctx, &data)

	// Read the current state of the HA node identified by ipaddress
	curState, found, err := r.readHaNodeState(data.Ipaddress.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read ha node for hafailover, got error: %s", err))
		return
	}
	if !found {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Cannot find ipaddress %v in ha node", data.Ipaddress.ValueString()))
		return
	}

	// Only force a failover when the current state differs from the desired state
	if curState != data.State.ValueString() {
		err := r.client.ActOnResource(service.Hafailover.Type(), &hafailover, "Force")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to force hafailover, got error: %s", err))
			return
		}
	}

	// Generate a unique ID for this resource (matches SDK v2 id format)
	data.Id = types.StringValue(sdkid.PrefixedUniqueId("tf-hafailover-"))

	// force is Optional+Computed and never returned by any GET; give it a concrete value
	if data.Force.IsNull() || data.Force.IsUnknown() {
		data.Force = types.BoolValue(false)
	}

	tflog.Trace(ctx, "Created hafailover resource")

	// Read the resulting HA node state back into the model
	r.readHafailoverFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HafailoverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data HafailoverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading hafailover resource")

	r.readHafailoverFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HafailoverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// All attributes are RequiresReplace, so Terraform never reaches Update with a
	// real change. This implementation preserves the identity and refreshes the
	// observed HA node state defensively.
	var data, state HafailoverResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating hafailover resource")

	// Preserve ID from prior state
	data.Id = state.Id

	// force is Optional+Computed and never returned by any GET; give it a concrete value
	if data.Force.IsNull() || data.Force.IsUnknown() {
		data.Force = types.BoolValue(false)
	}

	tflog.Trace(ctx, "Updated hafailover resource")

	// Read the resulting HA node state back into the model
	r.readHafailoverFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HafailoverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data HafailoverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting hafailover resource")

	// hafailover is a one-shot action resource with no NITRO delete verb; removing
	// it from state (done automatically by the framework) is the correct behavior,
	// mirroring the SDK v2 no-op delete.
	tflog.Trace(ctx, "Deleted hafailover resource from state")
}

// readHaNodeState returns the current HA state of the node identified by ipaddress.
// The returned bool indicates whether a matching node was found.
func (r *HafailoverResource) readHaNodeState(ipaddress string) (string, bool, error) {
	findParams := service.FindParams{
		ResourceType:             service.Hanode.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return "", false, err
	}

	for _, v := range dataArr {
		if v["ipaddress"] == ipaddress {
			if s, ok := v["state"].(string); ok {
				return s, true, nil
			}
			return "", true, nil
		}
	}

	return "", false, nil
}

// readHafailoverFromApi refreshes the observed HA node state into the model.
// It intentionally only sets the "state" attribute (mirroring the SDK v2 read):
// force, ipaddress and id are user/identity inputs and must be preserved.
func (r *HafailoverResource) readHafailoverFromApi(ctx context.Context, data *HafailoverResourceModel, diags *diag.Diagnostics) {
	tflog.Debug(ctx, "In readHafailoverFromApi Function")

	nodeState, found, err := r.readHaNodeState(data.Ipaddress.ValueString())
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read ha node for hafailover, got error: %s", err))
		return
	}
	if !found {
		diags.AddError("Client Error", fmt.Sprintf("Cannot find ipaddress %v in ha node", data.Ipaddress.ValueString()))
		return
	}

	data.State = types.StringValue(nodeState)
}
