package fis_channel_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &FisChannelBindingResource{}
var _ resource.ResourceWithConfigure = (*FisChannelBindingResource)(nil)
var _ resource.ResourceWithImportState = (*FisChannelBindingResource)(nil)

func NewFisChannelBindingResource() resource.Resource {
	return &FisChannelBindingResource{}
}

// FisChannelBindingResource defines the resource implementation.
type FisChannelBindingResource struct {
	client *service.NitroClient
}

func (r *FisChannelBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *FisChannelBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fis_channel_binding"
}

func (r *FisChannelBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *FisChannelBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data FisChannelBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating fis_channel_binding resource")
	fis_channel_binding := fis_channel_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - NITRO add is HTTP PUT (bind), use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Fis_channel_binding.Type(), &fis_channel_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create fis_channel_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created fis_channel_binding resource")

	// Set ID for the resource exactly once here. Composite key: name,ifnum
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("ifnum:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ifnum.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back from the aggregate endpoint.
	r.readFisChannelBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FisChannelBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data FisChannelBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading fis_channel_binding resource")

	// IMPORTANT (verified live on NS VPX): fis_channel_binding has NO NITRO read path.
	// The bind PUT succeeds (for a channel or even a physical interface), but the
	// binding is not surfaced by any GET -- the aggregate fis_binding/<name> response
	// carries only {"name"} with no fis_channel_binding array, and the direct endpoint
	// returns a keyless empty body. Following the no-GET / action-only precedent
	// (fis_interface_binding), a "not found" read does NOT mean the binding is gone, so
	// we must NOT remove it from state (that would spuriously delete a valid binding on
	// every refresh). A best-effort read is still attempted for forward compatibility.
	r.readFisChannelBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FisChannelBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state FisChannelBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for fis_channel_binding: NITRO exposes only add (PUT) and delete
	// (no update/change endpoint), and all schema attributes are RequiresReplace, so Terraform
	// recreates the resource on any change rather than calling Update.
	tflog.Debug(ctx, "Update is a no-op for fis_channel_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readFisChannelBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// fis_channel_bindingAggregateRead queries the AGGREGATE parent endpoint
// (GET /nitro/v1/config/fis_binding/<name>) and flattens the nested
// "fis_channel_binding" arrays into a single slice of binding rows.
//
// The direct fis_channel_binding endpoint returns a keyless empty body, so the
// bound channels are only retrievable via the parent aggregate.
func fis_channel_bindingAggregateRead(client *service.NitroClient, name string) ([]map[string]interface{}, error) {
	findParams := service.FindParams{
		ResourceType:             "fis_binding",
		ResourceName:             name,
		ResourceMissingErrorCode: 258,
	}
	parentArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return nil, err
	}

	rows := make([]map[string]interface{}, 0)
	for _, parent := range parentArr {
		nested, ok := parent["fis_channel_binding"]
		if !ok || nested == nil {
			continue
		}
		nestedArr, ok := nested.([]interface{})
		if !ok {
			continue
		}
		for _, item := range nestedArr {
			if m, ok := item.(map[string]interface{}); ok {
				rows = append(rows, m)
			}
		}
	}
	return rows, nil
}

func (r *FisChannelBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FisChannelBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting fis_channel_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs with the parent (name) as the
	// resource (URL) name and ifnum (UrlEncoded, contains '/') passed as args.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "ifnum"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	args := make([]string, 0)
	if val, ok := idMap["ifnum"]; ok && val != "" {
		args = append(args, fmt.Sprintf("ifnum:%s", utils.UrlEncode(val)))
	}

	err = r.client.DeleteResourceWithArgs(service.Fis_channel_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete fis_channel_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted fis_channel_binding binding")
}

// readFisChannelBindingFromApi reads the binding from the appliance via the
// AGGREGATE parent endpoint (fis_binding/<name>) and matches the row by ifnum.
// It returns true when the binding is found and the model was populated, false
// when the binding is genuinely absent (drift). Hard errors (parse / transport)
// are reported via diags.
func (r *FisChannelBindingResource) readFisChannelBindingFromApi(ctx context.Context, data *FisChannelBindingResourceModel, diags *diag.Diagnostics) bool {

	// Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "ifnum"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return false
	}

	name_Name, ok := idMap["name"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'name' not found in ID string")
		return false
	}

	// The direct fis_channel_binding endpoint returns a keyless empty body; read the
	// bound channels from the aggregate parent endpoint instead.
	dataArr, err := fis_channel_bindingAggregateRead(r.client, name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read fis_channel_binding, got error: %s", err))
		return false
	}

	// Binding genuinely absent (parent missing or no nested rows): report drift.
	if len(dataArr) == 0 {
		return false
	}

	// Iterate through results to find the one with the right ifnum
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		if idVal, ok := idMap["ifnum"]; ok {
			if val, ok := v["ifnum"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["ifnum"].(string); ok {
			match = false
			continue
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Binding row not present in the aggregate response: drift.
	if foundIndex == -1 {
		return false
	}

	fis_channel_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
	return true
}
