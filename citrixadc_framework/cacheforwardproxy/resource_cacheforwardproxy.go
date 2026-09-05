package cacheforwardproxy

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
var _ resource.Resource = &CacheforwardproxyResource{}
var _ resource.ResourceWithConfigure = (*CacheforwardproxyResource)(nil)
var _ resource.ResourceWithImportState = (*CacheforwardproxyResource)(nil)

func NewCacheforwardproxyResource() resource.Resource {
	return &CacheforwardproxyResource{}
}

// CacheforwardproxyResource defines the resource implementation.
type CacheforwardproxyResource struct {
	client *service.NitroClient
}

func (r *CacheforwardproxyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CacheforwardproxyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cacheforwardproxy"
}

func (r *CacheforwardproxyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CacheforwardproxyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CacheforwardproxyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cacheforwardproxy resource")

	cacheforwardproxy := cacheforwardproxyGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource keyed on ipaddress (matches SDK v2 behavior)
	ipaddress_value := data.Ipaddress.ValueString()
	_, err := r.client.AddResource(service.Cacheforwardproxy.Type(), ipaddress_value, &cacheforwardproxy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cacheforwardproxy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created cacheforwardproxy resource")

	// Set ID (new key:value format; ParseIdString also accepts the legacy "ipaddress,port" form)
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ipaddress:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ipaddress.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("port:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Port.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readCacheforwardproxyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "cacheforwardproxy not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheforwardproxyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CacheforwardproxyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cacheforwardproxy resource")

	r.readCacheforwardproxyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	// Resource is gone on the ADC (readFromApi nulled the Id): drop it from state so a
	// subsequent apply recreates it, matching the SDK v2 provider's behaviour.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheforwardproxyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CacheforwardproxyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating cacheforwardproxy resource")

	// cacheforwardproxy has no NITRO-updatable attributes: both ipaddress and port are
	// ForceNew/RequiresReplace, so any change triggers a destroy+create rather than an
	// in-place update. There is nothing to push here; just re-read current state.

	// Read the updated state back
	r.readCacheforwardproxyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "cacheforwardproxy not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheforwardproxyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CacheforwardproxyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cacheforwardproxy resource")

	// Multi-key resource - delete keyed on ipaddress with port passed as an arg
	// (matches SDK v2 DeleteResourceWithArgs). ParseIdString handles both the new
	// key:value ID and the legacy "ipaddress,port" positional form.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"ipaddress", "port"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	ipaddress_value, ok := idMap["ipaddress"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "attribute 'ipaddress' not found in ID")
		return
	}

	args := make([]string, 0)
	if val, ok := idMap["port"]; ok && val != "" {
		args = append(args, fmt.Sprintf("port:%s", val))
	}

	err = r.client.DeleteResourceWithArgs(service.Cacheforwardproxy.Type(), ipaddress_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete cacheforwardproxy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted cacheforwardproxy resource")
}

// Helper function to read cacheforwardproxy data from API
func (r *CacheforwardproxyResource) readCacheforwardproxyFromApi(ctx context.Context, data *CacheforwardproxyResourceModel, diags *diag.Diagnostics) {

	// Multi-key resource: parse ipaddress and port from the ID (handles both the new
	// key:value format and the legacy SDK v2 "ipaddress,port" positional format).
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"ipaddress", "port"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	ipaddress_Name, ok := idMap["ipaddress"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'ipaddress' not found in ID string")
		return
	}
	port_value, ok := idMap["port"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'port' not found in ID string")
		return
	}

	// cacheforwardproxy is retrieved as a collection (GET /cacheforwardproxy), then
	// filtered on ipaddress+port -- the same approach as the SDK v2 FindAllResources loop.
	findParams := service.FindParams{
		ResourceType:             service.Cacheforwardproxy.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cacheforwardproxy, got error: %s", err))
		return
	}

	// Resource no longer exists on the ADC: signal removal via a null Id (matches SDK v2 d.SetId("")).
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate to find the entry matching both ipaddress and port
	foundIndex := -1
	for i, v := range dataArr {
		ipVal, ipOk := v["ipaddress"].(string)
		if !ipOk || ipVal != ipaddress_Name {
			continue
		}
		if portFloat, portOk := v["port"].(float64); portOk {
			if fmt.Sprintf("%v", int(portFloat)) == port_value {
				foundIndex = i
				break
			}
		}
	}

	// Entry not present in the returned set: signal removal via a null Id (see above).
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	cacheforwardproxySetAttrFromGet(ctx, data, dataArr[foundIndex])
}
