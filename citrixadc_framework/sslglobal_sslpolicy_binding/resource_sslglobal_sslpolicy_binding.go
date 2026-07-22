package sslglobal_sslpolicy_binding

import (
	"context"
	"fmt"
	"strconv"
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
var _ resource.Resource = &SslglobalSslpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*SslglobalSslpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslglobalSslpolicyBindingResource)(nil)

func NewSslglobalSslpolicyBindingResource() resource.Resource {
	return &SslglobalSslpolicyBindingResource{}
}

// SslglobalSslpolicyBindingResource defines the resource implementation.
type SslglobalSslpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *SslglobalSslpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslglobalSslpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslglobal_sslpolicy_binding"
}

func (r *SslglobalSslpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslglobalSslpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslglobalSslpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslglobal_sslpolicy_binding resource")
	sslglobal_sslpolicy_binding := sslglobal_sslpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Sslglobal_sslpolicy_binding.Type(), &sslglobal_sslpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslglobal_sslpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslglobal_sslpolicy_binding resource")

	// Set ID for the resource before reading state
	// Composite ID = policyname,type,priority
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Type.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("priority:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Priority.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSslglobalSslpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslglobalSslpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslglobalSslpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslglobal_sslpolicy_binding resource")

	r.readSslglobalSslpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// If the binding was deleted out-of-band, remove it from state so a
	// subsequent apply re-creates it instead of erroring.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslglobalSslpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslglobalSslpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Pattern 5: no NITRO update endpoint for this binding; all attributes are
	// RequiresReplace, so Update is a no-op that just reconciles state.
	tflog.Debug(ctx, "Update is a no-op for sslglobal_sslpolicy_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readSslglobalSslpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslglobalSslpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslglobalSslpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslglobal_sslpolicy_binding resource")
	// Global keyless binding - delete using DeleteResourceWithArgs with empty resource
	// name; all three discriminators (policyname, type, priority) go in the args slice.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"policyname", "type", "priority"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	args := make([]string, 0, 3)
	if val, ok := idMap["policyname"]; ok && val != "" {
		args = append(args, "policyname:"+utils.UrlEncode(val))
	}
	if val, ok := idMap["type"]; ok && val != "" {
		args = append(args, "type:"+val)
	}
	if val, ok := idMap["priority"]; ok && val != "" {
		args = append(args, "priority:"+val)
	}

	err = r.client.DeleteResourceWithArgs(service.Sslglobal_sslpolicy_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslglobal_sslpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslglobal_sslpolicy_binding binding")
}

// Helper function to read sslglobal_sslpolicy_binding data from API
func (r *SslglobalSslpolicyBindingResource) readSslglobalSslpolicyBindingFromApi(ctx context.Context, data *SslglobalSslpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 3: Array filter without parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"policyname", "type", "priority"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	var dataArr []map[string]interface{}
	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["type"]; ok && val != "" {
		argsMap["type"] = val
	}

	findParams := service.FindParams{
		ResourceType:             service.Sslglobal_sslpolicy_binding.Type(),
		ArgsMap:                  argsMap,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslglobal_sslpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing - signal removal via null Id.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check policyname
		if idVal, ok := idMap["policyname"]; ok {
			if val, ok := v["policyname"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["policyname"].(string); ok {
			match = false
			continue
		}

		// Check priority
		if idVal, ok := idMap["priority"]; ok {
			if val, ok := v["priority"]; ok {
				val, _ = utils.ConvertToInt64(val)
				idValInt64, _ := strconv.ParseInt(idVal, 10, 64)
				if val != idValInt64 {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["priority"]; ok {
			match = false
			continue
		}
		// Check type
		if val, ok := idMap["type"]; ok && val != "" {
			if v, ok := v["type"]; ok {
				if v.(string) != val {
					match = false
				}
			}
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing - signal removal via null Id.
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	sslglobal_sslpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
