package subscriberprofile

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
var _ resource.Resource = &SubscriberprofileResource{}
var _ resource.ResourceWithConfigure = (*SubscriberprofileResource)(nil)
var _ resource.ResourceWithImportState = (*SubscriberprofileResource)(nil)

func NewSubscriberprofileResource() resource.Resource {
	return &SubscriberprofileResource{}
}

// SubscriberprofileResource defines the resource implementation.
type SubscriberprofileResource struct {
	client *service.NitroClient
}

func (r *SubscriberprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SubscriberprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_subscriberprofile"
}

func (r *SubscriberprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SubscriberprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SubscriberprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating subscriberprofile resource")

	// Build the create payload and push it to the ADC.
	subscriberprofile := subscriberprofileGetThePayloadFromthePlan(ctx, &data)

	// Named resource (keyed by ip) - use AddResource (matches SDK v2).
	subscriberprofileName := data.Ip.ValueString()
	_, err := r.client.AddResource(service.Subscriberprofile.Type(), subscriberprofileName, &subscriberprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create subscriberprofile, got error: %s", err))
		return
	}

	// ID is the plain ip value (matches SDK v2 d.SetId(ip)).
	data.Id = types.StringValue(subscriberprofileName)

	tflog.Trace(ctx, "Created subscriberprofile resource")

	// Read the updated state back
	if !r.readSubscriberprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "subscriberprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SubscriberprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SubscriberprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading subscriberprofile resource")

	found := r.readSubscriberprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SubscriberprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SubscriberprofileResourceModel

	// Read Terraform prior state to preserve ID and to detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating subscriberprofile resource")

	// Change detection on the updateable (non-ForceNew) attributes.
	hasChange := false
	if !data.Servicepath.Equal(state.Servicepath) {
		tflog.Debug(ctx, "servicepath has changed for subscriberprofile")
		hasChange = true
	}
	if !data.Subscriberrules.Equal(state.Subscriberrules) {
		tflog.Debug(ctx, "subscriberrules has changed for subscriberprofile")
		hasChange = true
	}
	if !data.Subscriptionidtype.Equal(state.Subscriptionidtype) {
		tflog.Debug(ctx, "subscriptionidtype has changed for subscriberprofile")
		hasChange = true
	}
	if !data.Subscriptionidvalue.Equal(state.Subscriptionidvalue) {
		tflog.Debug(ctx, "subscriptionidvalue has changed for subscriberprofile")
		hasChange = true
	}

	if hasChange {
		subscriberprofile := subscriberprofileGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// SDK v2 updates subscriberprofile via UpdateUnnamedResource.
		err := r.client.UpdateUnnamedResource(service.Subscriberprofile.Type(), &subscriberprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update subscriberprofile, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated subscriberprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for subscriberprofile resource, skipping update")
	}

	// Read the updated state back
	if !r.readSubscriberprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "subscriberprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SubscriberprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SubscriberprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting subscriberprofile resource")

	// Delete keyed by ip, disambiguated by vlan (matches SDK v2 DeleteResourceWithArgs).
	subscriberprofileName := data.Id.ValueString()
	vlan := int64(0)
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		vlan = data.Vlan.ValueInt64()
	}
	args := []string{fmt.Sprintf("vlan:%d", vlan)}
	err := r.client.DeleteResourceWithArgs(service.Subscriberprofile.Type(), subscriberprofileName, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete subscriberprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted subscriberprofile resource")
}

// Helper function to read subscriberprofile data from API.
// Mirrors SDK v2 Read: array query filtered by vlan (ArgsMap) then matched by ip.
// Returns false when the resource no longer exists on the ADC.
func (r *SubscriberprofileResource) readSubscriberprofileFromApi(ctx context.Context, data *SubscriberprofileResourceModel, diags *diag.Diagnostics) bool {
	// The ip is carried in the ID (plain value), which also works for import.
	subscriberprofileName := data.Id.ValueString()
	vlan := int64(0)
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		vlan = data.Vlan.ValueInt64()
	}

	findParams := service.FindParams{
		ResourceType:             service.Subscriberprofile.Type(),
		ArgsMap:                  map[string]string{"vlan": strconv.FormatInt(vlan, 10)},
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		// Resource is gone (mirrors SDK v2 which clears state on read error).
		return false
	}
	if len(dataArr) == 0 {
		return false
	}

	foundIndex := -1
	for i, v := range dataArr {
		if ipVal, ok := v["ip"]; ok && ipVal != nil && ipVal.(string) == subscriberprofileName {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		return false
	}

	subscriberprofileSetAttrFromGet(ctx, data, dataArr[foundIndex])
	return true
}
