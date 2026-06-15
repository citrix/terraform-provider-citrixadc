package snmptrap_snmpuser_binding

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
var _ resource.Resource = &SnmptrapSnmpuserBindingResource{}
var _ resource.ResourceWithConfigure = (*SnmptrapSnmpuserBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SnmptrapSnmpuserBindingResource)(nil)

func NewSnmptrapSnmpuserBindingResource() resource.Resource {
	return &SnmptrapSnmpuserBindingResource{}
}

// SnmptrapSnmpuserBindingResource defines the resource implementation.
type SnmptrapSnmpuserBindingResource struct {
	client *service.NitroClient
}

func (r *SnmptrapSnmpuserBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SnmptrapSnmpuserBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmptrap_snmpuser_binding"
}

func (r *SnmptrapSnmpuserBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SnmptrapSnmpuserBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SnmptrapSnmpuserBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating snmptrap_snmpuser_binding resource")
	snmptrap_snmpuser_binding := snmptrap_snmpuser_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Snmptrap_snmpuser_binding.Type(), &snmptrap_snmpuser_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create snmptrap_snmpuser_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created snmptrap_snmpuser_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("td:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Td.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("trapclass:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Trapclass.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("trapdestination:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Trapdestination.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("username:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Username.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("version:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Version.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readSnmptrapSnmpuserBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmptrapSnmpuserBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SnmptrapSnmpuserBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading snmptrap_snmpuser_binding resource")

	r.readSnmptrapSnmpuserBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmptrapSnmpuserBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SnmptrapSnmpuserBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating snmptrap_snmpuser_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		snmptrap_snmpuser_binding := snmptrap_snmpuser_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Snmptrap_snmpuser_binding.Type(), &snmptrap_snmpuser_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update snmptrap_snmpuser_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated snmptrap_snmpuser_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for snmptrap_snmpuser_binding resource, skipping update")
	}

	// Read the updated state back
	r.readSnmptrapSnmpuserBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmptrapSnmpuserBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SnmptrapSnmpuserBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting snmptrap_snmpuser_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name
	// Multiple unique attributes - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"trapclass", "trapdestination", "username"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["trapclass"]; ok && val != "" {
		argsMap["trapclass"] = val
	}
	if val, ok := idMap["trapdestination"]; ok && val != "" {
		argsMap["trapdestination"] = val
	}
	if val, ok := idMap["username"]; ok && val != "" {
		argsMap["username"] = val
	}
	if val, ok := idMap["version"]; ok && val != "" {
		argsMap["version"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Snmptrap_snmpuser_binding.Type(), "", argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete snmptrap_snmpuser_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted snmptrap_snmpuser_binding binding")
}

// Helper function to read snmptrap_snmpuser_binding data from API
func (r *SnmptrapSnmpuserBindingResource) readSnmptrapSnmpuserBindingFromApi(ctx context.Context, data *SnmptrapSnmpuserBindingResourceModel, diags *diag.Diagnostics) {

	// Case 3: Array filter without parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"trapclass", "trapdestination", "username"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Snmptrap_snmpuser_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read snmptrap_snmpuser_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "snmptrap_snmpuser_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check td
		if idVal, ok := idMap["td"]; ok {
			if val, ok := v["td"]; ok {
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
		} else if _, ok := v["td"]; ok {
			match = false
			continue
		}

		// Check trapclass
		if idVal, ok := idMap["trapclass"]; ok {
			if val, ok := v["trapclass"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["trapclass"].(string); ok {
			match = false
			continue
		}

		// Check trapdestination
		if idVal, ok := idMap["trapdestination"]; ok {
			if val, ok := v["trapdestination"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["trapdestination"].(string); ok {
			match = false
			continue
		}

		// Check username
		if idVal, ok := idMap["username"]; ok {
			if val, ok := v["username"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["username"].(string); ok {
			match = false
			continue
		}

		// Check version
		if idVal, ok := idMap["version"]; ok {
			if val, ok := v["version"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["version"].(string); ok {
			match = false
			continue
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("snmptrap_snmpuser_binding not found with the provided ID attributes"))
		return
	}

	snmptrap_snmpuser_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
