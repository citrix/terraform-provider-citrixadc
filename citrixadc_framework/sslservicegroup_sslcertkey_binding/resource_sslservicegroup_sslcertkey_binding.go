package sslservicegroup_sslcertkey_binding

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
var _ resource.Resource = &SslservicegroupSslcertkeyBindingResource{}
var _ resource.ResourceWithConfigure = (*SslservicegroupSslcertkeyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SslservicegroupSslcertkeyBindingResource)(nil)

func NewSslservicegroupSslcertkeyBindingResource() resource.Resource {
	return &SslservicegroupSslcertkeyBindingResource{}
}

// SslservicegroupSslcertkeyBindingResource defines the resource implementation.
type SslservicegroupSslcertkeyBindingResource struct {
	client *service.NitroClient
}

func (r *SslservicegroupSslcertkeyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslservicegroupSslcertkeyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslservicegroup_sslcertkey_binding"
}

func (r *SslservicegroupSslcertkeyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslservicegroupSslcertkeyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslservicegroupSslcertkeyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslservicegroup_sslcertkey_binding resource")
	sslservicegroup_sslcertkey_binding := sslservicegroup_sslcertkey_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Sslservicegroup_sslcertkey_binding.Type(), &sslservicegroup_sslcertkey_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslservicegroup_sslcertkey_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslservicegroup_sslcertkey_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ca:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ca.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("certkeyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Certkeyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("crlcheck:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Crlcheck.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("snicert:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Snicert.ValueBool()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSslservicegroupSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupSslcertkeyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslservicegroupSslcertkeyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslservicegroup_sslcertkey_binding resource")

	r.readSslservicegroupSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupSslcertkeyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslservicegroupSslcertkeyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslservicegroup_sslcertkey_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		sslservicegroup_sslcertkey_binding := sslservicegroup_sslcertkey_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Sslservicegroup_sslcertkey_binding.Type(), &sslservicegroup_sslcertkey_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslservicegroup_sslcertkey_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslservicegroup_sslcertkey_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslservicegroup_sslcertkey_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSslservicegroupSslcertkeyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslservicegroupSslcertkeyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslservicegroupSslcertkeyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslservicegroup_sslcertkey_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicegroupname", "certkeyname", "snicert", "ca"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	servicegroupname_value, ok := idMap["servicegroupname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'servicegroupname' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["ca"]; ok && val != "" {
		argsMap["ca"] = val
	}
	if val, ok := idMap["certkeyname"]; ok && val != "" {
		argsMap["certkeyname"] = val
	}
	if val, ok := idMap["crlcheck"]; ok && val != "" {
		argsMap["crlcheck"] = val
	}
	if val, ok := idMap["snicert"]; ok && val != "" {
		argsMap["snicert"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Sslservicegroup_sslcertkey_binding.Type(), servicegroupname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslservicegroup_sslcertkey_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslservicegroup_sslcertkey_binding binding")
}

// Helper function to read sslservicegroup_sslcertkey_binding data from API
func (r *SslservicegroupSslcertkeyBindingResource) readSslservicegroupSslcertkeyBindingFromApi(ctx context.Context, data *SslservicegroupSslcertkeyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"servicegroupname", "certkeyname", "snicert", "ca"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	servicegroupname_Name, ok := idMap["servicegroupname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'servicegroupname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Sslservicegroup_sslcertkey_binding.Type(),
		ResourceName:             servicegroupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslservicegroup_sslcertkey_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "sslservicegroup_sslcertkey_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ca
		if idVal, ok := idMap["ca"]; ok {
			if val, ok := v["ca"].(bool); ok {
				idValBool, _ := strconv.ParseBool(idVal)
				if val != idValBool {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["ca"].(bool); ok {
			match = false
			continue
		}

		// Check certkeyname
		if idVal, ok := idMap["certkeyname"]; ok {
			if val, ok := v["certkeyname"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["certkeyname"].(string); ok {
			match = false
			continue
		}

		// Check crlcheck
		if idVal, ok := idMap["crlcheck"]; ok {
			if val, ok := v["crlcheck"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["crlcheck"].(string); ok {
			match = false
			continue
		}

		// Check snicert
		if idVal, ok := idMap["snicert"]; ok {
			if val, ok := v["snicert"].(bool); ok {
				idValBool, _ := strconv.ParseBool(idVal)
				if val != idValBool {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["snicert"].(bool); ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("sslservicegroup_sslcertkey_binding not found with the provided ID attributes"))
		return
	}

	sslservicegroup_sslcertkey_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
