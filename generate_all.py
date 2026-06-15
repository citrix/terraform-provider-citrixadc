#!/usr/bin/env python3
"""
Terraform Provider CitrixADC - Comprehensive Code Generator

This script combines all individual generators to create complete Terraform Framework
implementations for CitrixADC resources. It generates:

1. datasource_<resource>.go - Datasource implementation
2. datasource_schema.go - Datasource schema definition  
3. resource_<resource>.go - Resource implementation
4. resource_schema.go - Resource schema definition with utility functions

All files are generated for each resource listed in resources.txt.
"""

import json
import os
import sys
import re
from pathlib import Path
from typing import Dict, List, Any
from jinja2 import Environment, FileSystemLoader, Template

# Import the resource to module mapping
from resource_module_mapping import RESOURCE_MODULE_MAPPING

# Load legacy SDK v2 ID attribute order (used for backward-compatible ID parsing)
_RESOURCE_ID_MAPPING_PATH = Path(__file__).parent / 'resource_id_mapping.json'
if _RESOURCE_ID_MAPPING_PATH.exists():
    with open(_RESOURCE_ID_MAPPING_PATH) as _f:
        _LEGACY_ID_MAPPING: Dict[str, str] = json.load(_f)
else:
    _LEGACY_ID_MAPPING = {}


# ============================================================================
# UTILITY FUNCTIONS
# ============================================================================

def format_default_value_for_go(default_value, attr_type):
    """Format default values properly for Go code generation."""
    if default_value is None:
        return None
    
    if attr_type == 'string':
        # Handle already-quoted strings like '"200 OK"' or '"none"'
        str_value = str(default_value)
        if str_value.startswith('"') and str_value.endswith('"'):
            # Already quoted, use as-is but unescape for Go
            return str_value[1:-1]  # Remove outer quotes
        else:
            # Not quoted, return as-is to be quoted in template
            return str_value
    else:
        return default_value

# aaagroup_auditnslogpolicy_binding => AaagroupAuditnslogpolicyBinding
def convert_to_pascal_case(snake_str):
    """Convert snake_case to PascalCase."""
    components = snake_str.split('_')
    return ''.join(x.capitalize() for x in components)

# aaagroup_auditnslogpolicy_binding => Aaagroupauditnslogpolicybinding
def convert_to_sentence_case(snake_str):
    """Convert snake_case to sentence case (first letter uppercase)."""
    components = snake_str.split('_')
    return components[0].capitalize() + ''.join(x for x in components[1:])

# aaagroup_auditnslogpolicy_binding => Aaagroup_auditnslogpolicy_binding
def convert_to_sentence_case_and_retain_underscore(snake_str):
    """Convert snake_case to sentnce case and retain underscores."""
    components = snake_str.split('_')
    return  components[0].capitalize() + '_' + '_'.join(x for x in components[1:]) if len(components) > 1 else components[0].capitalize()

def convert_to_title_case(snake_str):
    """Convert snake_case to Title Case with spaces."""
    components = snake_str.split('_')
    return ' '.join(x.capitalize() for x in components)


def get_nitro_service_type(resource_name):
    """
    Map resource names to their corresponding NITRO service types.
    This mapping matches what's used in the NITRO Go client.
    """
    # Common mapping patterns for NITRO service types
    service_mapping = {
        'nsparam': 'Nsparam',
        'lbparameter': 'Lbparameter', 
        'lbvserver': 'Lbvserver',
        'csvserver': 'Csvserver',
        'sslvserver': 'Sslvserver',
        'gslbvserver': 'Gslbvserver',
        'server': 'Server',
        'service': 'Service',
        'servicegroup': 'Servicegroup',
        'lbmonitor': 'Lbmonitor',
        'sslcertkey': 'Sslcertkey',
        'systemfile': 'Systemfile',
        'rewritepolicy': 'Rewritepolicy',
        'rewriteaction': 'Rewriteaction',
        'responderpolicy': 'Responderpolicy',
        'responderaction': 'Responderaction',
        'tmtrafficpolicy': 'Tmtrafficpolicy',
        'tmtrafficaction': 'Tmtrafficaction',
        'authenticationvserver': 'Authenticationvserver',
        'authenticationpolicy': 'Authenticationpolicy',
        'authenticationaction': 'Authenticationaction',
        'appfwpolicy': 'Appfwpolicy',
        'appfwprofile': 'Appfwprofile',
        'transformpolicy': 'Transformpolicy',
        'transformaction': 'Transformaction',
        'cachepolicy': 'Cachepolicy',
        'cacheaction': 'Cacheaction',
        'compressionpolicy': 'Compressionpolicy',
        'compressionaction': 'Compressionaction',
    }
    
    # If not in mapping, use PascalCase version
    return service_mapping.get(resource_name, convert_to_sentence_case_and_retain_underscore(resource_name))


def get_nitro_package_info(resource_name: str) -> Dict[str, str]:
    """Get NITRO package import information based on resource name."""
    # Use the resource_module_mapping for accurate module lookup
    package = RESOURCE_MODULE_MAPPING.get(resource_name)
    
    # Fallback to basic if not found in mapping
    if not package:
        print(f"Warning: Module mapping not found for {resource_name}, defaulting to 'basic'")
        package = 'basic'
    
    return {
        'import_path': f'"github.com/citrix/adc-nitro-go/resource/config/{package}"',
        'type_name': f'{package}.{convert_to_sentence_case(resource_name)}'
    }


# ============================================================================
# SCHEMA PROCESSING FUNCTIONS
# ============================================================================

def get_terraform_type_info(metadata_type: str) -> Dict[str, str]:
    """
    Map metadata types to Terraform Plugin Framework types.
    Returns dict with 'framework_type', 'attribute_type', and 'element_type' (if applicable).
    """
    type_mapping = {
        'string': {
            'framework_type': 'types.String',
            'attribute_type': 'schema.StringAttribute',
            'element_type': None,
            'default_import': 'stringdefault'
        },
        'integer': {
            'framework_type': 'types.Int64',
            'attribute_type': 'schema.Int64Attribute', 
            'element_type': None,
            'default_import': 'int64default'
        },
        'boolean': {
            'framework_type': 'types.Bool',
            'attribute_type': 'schema.BoolAttribute',
            'element_type': None,
            'default_import': 'booldefault'
        },
        'number': {
            'framework_type': 'types.Float64',
            'attribute_type': 'schema.Float64Attribute',
            'element_type': None,
            'default_import': 'float64default'
        },
        'string[]': {
            'framework_type': 'types.List',
            'attribute_type': 'schema.ListAttribute',
            'element_type': 'types.StringType',
            'default_import': None
        },
        'integer[]': {
            'framework_type': 'types.List',
            'attribute_type': 'schema.ListAttribute',
            'element_type': 'types.Int64Type',
            'default_import': None
        },
        'boolean[]': {
            'framework_type': 'types.List',
            'attribute_type': 'schema.ListAttribute',
            'element_type': 'types.BoolType',
            'default_import': None
        },
        'number[]': {
            'framework_type': 'types.List',
            'attribute_type': 'schema.ListAttribute',
            'element_type': 'types.Float64Type',
            'default_import': None
        }
    }
    
    return type_mapping.get(metadata_type, {
        'framework_type': 'types.String',  # Default fallback
        'attribute_type': 'schema.StringAttribute',
        'element_type': None,
        'default_import': None
    })


def get_planmodifier_info(metadata_type: str) -> Dict[str, str]:
    """Get plan modifier information based on attribute type."""
    planmodifier_mapping = {
        'string': {
            'import': 'stringplanmodifier',
            'modifier': 'stringplanmodifier.RequiresReplace()'
        },
        'integer': {
            'import': 'int64planmodifier',
            'modifier': 'int64planmodifier.RequiresReplace()'
        },
        'boolean': {
            'import': 'boolplanmodifier',
            'modifier': 'boolplanmodifier.RequiresReplace()'
        },
        'number': {
            'import': 'float64planmodifier',
            'modifier': 'float64planmodifier.RequiresReplace()'
        },
        'string[]': {
            'import': 'listplanmodifier',
            'modifier': 'listplanmodifier.RequiresReplace()'
        }
    }
    
    return planmodifier_mapping.get(metadata_type, {
        'import': 'stringplanmodifier',
        'modifier': 'stringplanmodifier.RequiresReplace()'
    })


def analyze_datasource_read_pattern(attributes: List[Dict[str, Any]]) -> Dict[str, Any]:
    """
    Analyze metadata attributes to determine the datasource Read method pattern.
    
    Returns a dict with:
    - pattern: 'simple_find' | 'find_with_id' | 'array_filter_no_id' | 'array_filter_with_id'
    - get_id_attr: attribute with is_get_id=true (if any)
    - unique_attrs: list of attributes with x-unique-attr=true (with type info)
    - filter_attrs: list of unique_attrs excluding get_id_attr (for filtering, with type info)
    """
    get_id_attrs = [attr for attr in attributes if attr.get('is_get_id', False)]
    unique_attrs = [attr for attr in attributes if attr.get('x-unique-attr', False)]
    
    # Extract attribute names and types for code generation
    get_id_attr = get_id_attrs[0] if get_id_attrs else None
    filter_attrs = [attr for attr in unique_attrs if attr != get_id_attr]
    
    # Add type information for filtering comparisons
    for attr in unique_attrs:
        attr_type = attr.get('type', 'string')
        attr['attr_type'] = attr_type
        # Determine value accessor and type assertion based on type
        if attr_type == 'integer':
            attr['value_accessor'] = 'ValueInt64()'
            attr['type_assertion'] = 'float64'  # JSON numbers are float64 in Go
            attr['comparison'] = 'int64(v["' + attr['option_name'] + '"].(float64))'
        elif attr_type == 'boolean':
            attr['value_accessor'] = 'ValueBool()'
            attr['type_assertion'] = 'bool'
            attr['comparison'] = 'v["' + attr['option_name'] + '"].(bool)'
        elif attr_type == 'number':
            attr['value_accessor'] = 'ValueFloat64()'
            attr['type_assertion'] = 'float64'
            attr['comparison'] = 'v["' + attr['option_name'] + '"].(float64)'
        else:  # string
            attr['value_accessor'] = 'ValueString()'
            attr['type_assertion'] = 'string'
            attr['comparison'] = 'v["' + attr['option_name'] + '"].(string)'
    
    # Determine pattern based on combinations
    if not get_id_attr and not unique_attrs:
        # Case 1: No is_get_id, no x-unique-attr
        pattern = 'simple_find'
    elif get_id_attr and len(unique_attrs) == 1 and get_id_attr in unique_attrs:
        # Case 2: One attribute with both is_get_id and x-unique-attr
        pattern = 'find_with_id'
    elif not get_id_attr and unique_attrs:
        # Case 3: No is_get_id, but has x-unique-attr
        pattern = 'array_filter_no_id'
    elif get_id_attr and len(unique_attrs) > 1:
        # Case 4: has is_get_id and multiple x-unique-attr
        pattern = 'array_filter_with_id'
    else:
        # Default fallback
        pattern = 'simple_find'
    
    return {
        'pattern': pattern,
        'get_id_attr': get_id_attr,
        'unique_attrs': unique_attrs,
        'filter_attrs': filter_attrs
    }


def process_attribute_for_datasource(attr: Dict[str, Any]) -> Dict[str, Any]:
    """Process a single attribute from metadata for datasource schema."""
    option_name = attr.get('option_name', '')
    attr_type = attr.get('type', 'string')
    description = attr.get('description', [])
    x_unique_attr = attr.get('x-unique-attr', False)
    x_secret_attr = attr.get('x-secret-attr', False)
    
    # Join description list into single string
    if isinstance(description, list):
        desc_text = '\n'.join(description)
    else:
        desc_text = str(description) if description else option_name
    
    # Escape quotes and newlines in description for Go string literals
    desc_text = desc_text.replace('\\', '\\\\')  # Escape backslashes first
    desc_text = desc_text.replace('"', '\\"')    # Escape quotes
    desc_text = desc_text.replace('\n', '\\n')   # Escape newlines
    
    # Get Terraform types
    tf_types = get_terraform_type_info(attr_type)
    
    return {
        'name': option_name,
        'option_name': option_name,
        'type_info': tf_types,
        'description': desc_text,
        'is_required': x_unique_attr,
        'is_optional': not x_unique_attr,
        'is_computed': not x_unique_attr,  # All datasource attributes except required are computed
        'x-secret-attr': x_secret_attr,
        'is_secret': x_secret_attr,
        'original_type': attr_type,
        'pascal_name': convert_to_pascal_case(option_name),
        'sentence_name': convert_to_sentence_case(option_name)
    }


def process_attribute_for_resource(attr: Dict[str, Any]) -> Dict[str, Any]:
    """Process a single attribute from metadata for resource schema."""
    option_name = attr.get('option_name', '')
    attr_type = attr.get('type', 'string')
    description = attr.get('description', [])
    is_required = attr.get('is_required', False)
    is_get_id = attr.get('is_get_id', False)
    is_delete_id = attr.get('is_delete_id', False)
    x_unique_attr = attr.get('x-unique-attr', False)
    x_secret_attr = attr.get('x-secret-attr', False)
    default_value = attr.get('default', None)
    is_updateable = attr.get('is_updateable', True)
    
    # Join description list into single string
    if isinstance(description, list):
        desc_text = '\n'.join(description)
    else:
        desc_text = str(description) if description else option_name
    
    # Escape quotes and newlines in description for Go string literals
    desc_text = desc_text.replace('\\', '\\\\')  # Escape backslashes first
    desc_text = desc_text.replace('"', '\\"')    # Escape quotes
    desc_text = desc_text.replace('\n', '\\n')   # Escape newlines
    
    # Get Terraform types
    type_info = get_terraform_type_info(attr_type)
    
    # Determine attribute properties
    is_optional = not is_required
    # In Terraform Framework, attributes with default values must be computed
    # is_computed = bool(default_value is not None or not is_required)
    is_computed = not is_required
    
    # Get plan modifier info if needed
    needs_plan_modifier = not is_updateable
    plan_modifier_info = get_planmodifier_info(attr_type) if needs_plan_modifier else None
    
    return {
        'name': option_name,
        'pascal_name': convert_to_pascal_case(option_name),
        'sentence_name': convert_to_sentence_case(option_name),
        'type_info': type_info,
        'description': desc_text,
        'is_required': is_required,
        'is_optional': is_optional,
        'is_computed': is_computed,
        # 'default_value': format_default_value_for_go(default_value, attr_type),
        'default_value': None,
        'needs_plan_modifier': needs_plan_modifier,
        'plan_modifier_info': plan_modifier_info,
        'original_type': attr_type,
        'x-unique-attr': x_unique_attr,
        'is_secret': x_secret_attr,
        'is_get_id': is_get_id,
        'is_delete_id': is_delete_id
    }


def expand_secret_attributes(processed_attributes: List[Dict[str, Any]]) -> List[Dict[str, Any]]:
    """
    Expand attributes marked as secret (is_secret=True) into three attributes:
    1. Original attribute (e.g., password) - Optional, Sensitive
    2. Write-only variant (e.g., password_wo) - Optional, Sensitive, WriteOnly
    3. Version attribute (e.g., password_wo_version) - Optional, Computed, Default: 1

    This supports ephemeral credential handling in Terraform.
    When the original attribute is required, both original and _wo are marked Optional
    and a ValidateConfig check ensures at least one is provided.
    """
    expanded_attributes = []

    for attr in processed_attributes:
        if attr.get('is_secret', False):
            # Add the original attribute with Sensitive=true
            original_attr = attr.copy()
            original_attr['is_sensitive'] = True
            # If originally required, downgrade to optional (validation will ensure one is set)
            if original_attr.get('is_required', False):
                original_attr['was_required_secret'] = True
                original_attr['is_required'] = False
                original_attr['is_optional'] = True
                # is_computed stays as-is (optional attrs are computed in TF framework)
                original_attr['is_computed'] = True
            expanded_attributes.append(original_attr)
            
            # Add write-only variant
            wo_attr = attr.copy()
            wo_name = f"{attr['name']}_wo"
            wo_attr['name'] = wo_name
            wo_attr['pascal_name'] = convert_to_pascal_case(wo_name)
            wo_attr['sentence_name'] = convert_to_sentence_case(wo_name)
            wo_attr['is_sensitive'] = True
            wo_attr['is_write_only'] = True
            wo_attr['is_computed'] = False  # WriteOnly attributes cannot be computed
            wo_attr['is_secret'] = False  # Only mark the original as secret for template logic
            # Write-only variant is always optional (validation ensures one is set)
            wo_attr['is_required'] = False
            wo_attr['is_optional'] = True
            # Carry the was_required_secret flag for validation generation
            if attr.get('is_required', False):
                wo_attr['was_required_secret'] = True
            # Store reference to original attribute for payload mapping
            wo_attr['original_attr_name'] = attr['name']
            wo_attr['original_attr_pascal_name'] = attr['pascal_name']
            wo_attr['original_attr_sentence_name'] = attr['sentence_name']
            expanded_attributes.append(wo_attr)
            
            # Add version attribute (always integer type)
            # Inherit needs_plan_modifier from the secret attribute: if the write-only
            # attribute requires replacement, bumping the version must also trigger it.
            wo_needs_replace = attr.get('needs_plan_modifier', False)
            version_attr = {
                'name': f"{attr['name']}_wo_version",
                'pascal_name': convert_to_pascal_case(f"{attr['name']}_wo_version"),
                'sentence_name': convert_to_sentence_case(f"{attr['name']}_wo_version"),
                'type_info': get_terraform_type_info('integer'),
                'description': f"Increment this version to signal a {attr['name']}_wo update.",
                'is_required': False,
                'is_optional': True,
                'is_computed': True,
                'default_value': 1,
                'needs_plan_modifier': wo_needs_replace,
                'plan_modifier_info': get_planmodifier_info('integer') if wo_needs_replace else None,
                'original_type': 'integer',
                'x-unique-attr': False,
                'is_secret': False,
                'is_version_tracker': True,
                'tracks_secret': attr['name']  # Track which secret this is for
            }
            expanded_attributes.append(version_attr)
        else:
            expanded_attributes.append(attr)
    
    return expanded_attributes


def load_resource_metadata(tfdata_dir: Path, resource_name: str) -> List[Dict[str, Any]]:
    """Load and process metadata for a resource."""
    metadata_file = tfdata_dir / f"{resource_name}.json"
    
    if not metadata_file.exists():
        print(f"Warning: Metadata file not found: {metadata_file}")
        return []
    
    try:
        with open(metadata_file, 'r') as f:
            metadata = json.load(f)
        
        if not isinstance(metadata, list):
            print(f"Warning: Expected list in {metadata_file}, got {type(metadata)}")
            return []
            
        return metadata
    
    except json.JSONDecodeError as e:
        print(f"Error parsing JSON in {metadata_file}: {e}")
        return []
    except Exception as e:
        print(f"Error loading {metadata_file}: {e}")
        return []


def analyze_id_generation_pattern(attributes: List[Dict[str, Any]]) -> Dict[str, Any]:
    """
    Analyze metadata attributes to determine the ID generation pattern.
    
    Returns a dict with:
    - pattern: 'static' | 'single_unique' | 'multiple_unique'
    - unique_attrs: list of attributes with x-unique-attr=true (with pascal_name and type info added)
    - needs_fmt: boolean indicating if fmt package is needed
    """
    unique_attrs = [attr for attr in attributes if attr.get('x-unique-attr', False)]
    
    # Add pascal_name and type info to unique attributes for template usage
    for attr in unique_attrs:
        if 'pascal_name' not in attr:
            attr['pascal_name'] = convert_to_pascal_case(attr.get('option_name', ''))
        # Add type information for proper formatting in ID generation
        attr_type = attr.get('type', 'string')
        attr['attr_type'] = attr_type
        # Determine format specifier and value accessor based on type
        if attr_type == 'integer':
            attr['format_spec'] = '%d'
            attr['value_accessor'] = 'ValueInt64()'
        elif attr_type == 'boolean':
            attr['format_spec'] = '%t'
            attr['value_accessor'] = 'ValueBool()'
        elif attr_type == 'number':
            attr['format_spec'] = '%f'
            attr['value_accessor'] = 'ValueFloat64()'
        else:  # string or default
            attr['format_spec'] = '%s'
            attr['value_accessor'] = 'ValueString()'
    
    if not unique_attrs:
        # Case 1: No unique attributes - static ID
        pattern = 'static'
        needs_fmt = False
        needs_strings = False
    elif len(unique_attrs) == 1:
        # Case 2: Single unique attribute - need fmt.Sprintf for key:base64(value) format
        pattern = 'single_unique'
        needs_fmt = True
        needs_strings = False
    else:
        # Case 3: Multiple unique attributes - need fmt.Sprintf and strings.Join
        pattern = 'multiple_unique'
        needs_fmt = True
        needs_strings = True
    
    return {
        'pattern': pattern,
        'unique_attrs': unique_attrs,
        'needs_fmt': needs_fmt,
        'needs_strings': needs_strings
    }


def get_required_imports_for_resource(attributes: List[Dict[str, Any]]) -> Dict[str, bool]:
    """Determine which imports are needed based on attributes for resource schema."""
    imports = {
        'types': True,  # Always needed
        'planmodifier': False,
        'stringplanmodifier': False,
        'int64planmodifier': False,
        'boolplanmodifier': False,
        'float64planmodifier': False,
        'listplanmodifier': False,
        'int64default': False,
        'booldefault': False,
        'stringdefault': False,
        'float64default': False,
        'fmt': False,  # For multiple unique attributes
        'strings': False
    }
    
    for attr in attributes:
        # Check for plan modifiers
        if attr.get('needs_plan_modifier') and attr.get('plan_modifier_info'):
            imports['planmodifier'] = True
            imports[attr['plan_modifier_info']['import']] = True
        
        # Check for default value imports (disabled for regular attrs, kept for _wo_version)
        # if attr.get('default_value') is not None:
        #     type_import = attr['type_info'].get('default_import')
        #     if type_import:
        #         imports[type_import] = True
        if attr.get('is_version_tracker') and attr.get('default_value') is not None:
            type_import = attr['type_info'].get('default_import')
            if type_import:
                imports[type_import] = True
    
    # Check if fmt and strings are needed for ID generation
    id_pattern = analyze_id_generation_pattern(attributes)
    if id_pattern['needs_fmt']:
        imports['fmt'] = True
    if id_pattern['needs_strings']:
        imports['strings'] = True
    
    return imports


def get_required_imports_for_resource_impl(read_pattern: Dict[str, Any], id_generation: Dict[str, Any] = None) -> Dict[str, bool]:
    """Determine which imports are needed for resource implementation based on read pattern and id_generation."""
    imports = {
        'strconv': False,
        'types': False,
        'strings': False,
        'utils': False
    }
    
    if not read_pattern:
        return imports
    
    # Check if strconv is needed - only for patterns that handle integer/boolean ID or filter attributes
    pattern = read_pattern.get('pattern')

    # For array_filter patterns, check if any unique or filter attrs are integer/boolean
    if pattern in ['array_filter_no_id', 'array_filter_with_id']:
        for attr in read_pattern.get('unique_attrs', []) + read_pattern.get('filter_attrs', []):
            if 'ValueInt64' in attr.get('value_accessor', '') or 'ValueBool' in attr.get('value_accessor', ''):
                imports['strconv'] = True
                break
    
    # Check if strings is needed for ID generation (case 3: multiple unique attributes)
    if id_generation and id_generation.get('pattern') == 'multiple_unique':
        imports['strings'] = True
    
    # types is needed for setting ID in Create function
    if id_generation:
        imports['types'] = True
    
    # utils is needed for:
    # 1. Parsing IDs: ParseIdString() - for array_filter patterns (multiple unique attrs) and binding deletes
    # 2. Encoding IDs: UrlEncode() - for multiple_unique ID generation only
    # 3. ConvertToInt64() - used in array filter comparisons
    utils_needed = False

    # Check if ID parsing is needed (ParseIdString for multi-value ID patterns only)
    # find_with_id uses a single_unique raw-value ID and does NOT need ParseIdString
    if pattern and pattern in ['array_filter_no_id', 'array_filter_with_id']:
        utils_needed = True

    # Check if ID URL encoding is needed (UrlEncode for multiple_unique only)
    if id_generation and id_generation.get('pattern') == 'multiple_unique':
        utils_needed = True

    imports['utils'] = utils_needed
    
    return imports


# ============================================================================
# TEMPLATE VARIABLE GENERATORS
# ============================================================================

def get_datasource_template_variables(resource_name, raw_attributes=None):
    """Generate template variables for datasource implementation."""
    pascal_case = convert_to_pascal_case(resource_name)
    title_case = convert_to_title_case(resource_name)
    
    # Special handling for lbparameter-style naming
    struct_name = f'{pascal_case}DataSource'
    constructor_name = f'{pascal_case[:2].upper()}{pascal_case[2:]}DataSource'
    schema_func_name = f'{pascal_case}DataSourceSchema'
    
    # Analyze read pattern if attributes provided
    read_pattern = None
    if raw_attributes:
        read_pattern = analyze_datasource_read_pattern(raw_attributes)
        # Add helper info for template
        if read_pattern['get_id_attr']:
            attr = read_pattern['get_id_attr']
            read_pattern['get_id_var_name'] = attr['option_name'].replace('_', '') + '_Name'
            read_pattern['get_id_field_name'] = convert_to_pascal_case(attr['option_name'])
            read_pattern['get_id_attr_name'] = attr['option_name']
        
        # Process filter attributes
        for attr in read_pattern['filter_attrs']:
            attr['var_name'] = attr['option_name'].replace('_', '') + '_Name'
            attr['field_name'] = convert_to_pascal_case(attr['option_name'])

        # Process unique attributes for array_filter_no_id pattern
        for attr in read_pattern['unique_attrs']:
            if 'var_name' not in attr:  # Only process if not already done
                attr['var_name'] = attr['option_name'].replace('_', '') + '_Name'
                attr['field_name'] = convert_to_pascal_case(attr['option_name'])
    
    return {
        'package_name': resource_name,
        'resource_name': resource_name,
        'datasource_struct_name': struct_name,
        'datasource_constructor_name': constructor_name,
        'datasource_schema_func_name': schema_func_name,
        'model_name': f'{pascal_case}ResourceModel',
        'nitro_service_type': get_nitro_service_type(resource_name),
        'set_attr_func': f'{resource_name}SetAttrFromGet',
        'title_case': title_case,
        'read_pattern': read_pattern
    }


def get_resource_template_variables(resource_name, raw_attributes=None):
    """Generate template variables for resource implementation."""
    pascal_case = convert_to_pascal_case(resource_name)
    title_case = convert_to_title_case(resource_name)
    
    # Analyze read pattern for resource (same as datasource)
    read_pattern = None
    resource_impl_imports = {'strconv': False, 'types': False, 'strings': False, 'utils': False}
    id_generation = None
    processed_attributes = []
    
    if raw_attributes:
        read_pattern = analyze_datasource_read_pattern(raw_attributes)
        id_generation = analyze_id_generation_pattern(raw_attributes)
        resource_impl_imports = get_required_imports_for_resource_impl(read_pattern, id_generation)
        
        # Process attributes for the template
        processed_attributes = [process_attribute_for_resource(attr) for attr in raw_attributes]
        processed_attributes = expand_secret_attributes(processed_attributes)
        
        # Add field name and var name for code generation
        if read_pattern['get_id_attr']:
            attr = read_pattern['get_id_attr']
            read_pattern['get_id_var_name'] = attr['option_name'].replace('_', '') + '_Name'
            read_pattern['get_id_field_name'] = convert_to_pascal_case(attr['option_name'])
            read_pattern['get_id_attr_name'] = attr['option_name']
        
        # Add field names and var names for filter attributes
        for attr in read_pattern['filter_attrs']:
            attr['var_name'] = attr['option_name'].replace('_', '')
            attr['field_name'] = convert_to_pascal_case(attr['option_name'])
            
        for attr in read_pattern['unique_attrs']:
            attr['var_name'] = attr['option_name'].replace('_', '')
            attr['field_name'] = convert_to_pascal_case(attr['option_name'])
    
    # Legacy SDK v2 ID attribute order for backward-compatible ParseIdString calls
    legacy_id_str = _LEGACY_ID_MAPPING.get(resource_name, '')
    legacy_id_attrs_raw = [a for a in legacy_id_str.split(',') if a] if legacy_id_str else []
    legacy_id_attr_order = [a.rstrip('?') for a in legacy_id_attrs_raw]
    legacy_id_optional_attrs = [a.rstrip('?') for a in legacy_id_attrs_raw if a.endswith('?')]

    return {
        'package_name': resource_name,
        'resource_name': resource_name,
        'pascal_name': pascal_case,
        'resource_struct_name': f'{pascal_case}Resource',
        'model_name': f'{pascal_case}ResourceModel',
        'resource_var_name': resource_name.lower(),
        'nitro_service_type': get_nitro_service_type(resource_name),
        'get_payload_func': f'{resource_name}GetThePayloadFromthePlan',
        'set_attr_func': f'{resource_name}SetAttrFromGet',
        'title_case': title_case,
        'read_pattern': read_pattern,
        'resource_impl_imports': resource_impl_imports,
        'id_generation': id_generation,
        'attributes': processed_attributes,
        'legacy_id_attr_order': legacy_id_attr_order,
        'legacy_id_optional_attrs': legacy_id_optional_attrs
    }


# ============================================================================
# PROVIDER.GO UPDATE FUNCTIONS
# ============================================================================

def cleanup_unused_utils_imports(framework_dir: Path) -> int:
    """
    Remove utils import line from resource_schema.go files if utils is only mentioned once.
    
    Returns:
        Number of files modified
    """
    resource_schema_files = list(framework_dir.glob('*/resource_schema.go')) + list(framework_dir.glob('*/datasource_*.go'))

    modified_count = 0
    
    for file_path in resource_schema_files:
        with open(file_path, 'r') as f:
            content = f.read()
        
        # Count occurrences of "utils"
        utils_count = content.count('utils')
        
        # If utils appears only once, it's only in the import and not used
        if utils_count == 1:
            lines = content.split('\n')
            new_lines = []
            removed = False
            
            for line in lines:
                if 'github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils' in line:
                    removed = True
                    continue  # Skip this line
                new_lines.append(line)
            
            if removed:
                with open(file_path, 'w') as f:
                    f.write('\n'.join(new_lines))
                modified_count += 1
    
    return modified_count


def update_resources_go_file(framework_dir: Path, resources_go_path: Path) -> int:
    """
    Update resources.go to include all resources from citrixadc_framework.
    
    Returns:
        Number of new resources added
    """
    # Get all resources from framework (excluding special directories)
    skip_dirs = {'provider', 'utils', 'acctests'}
    framework_resources = []
    
    for subdir in framework_dir.iterdir():
        if subdir.is_dir() and subdir.name not in skip_dirs:
            framework_resources.append(subdir.name)
    
    framework_resources = sorted(framework_resources)
    
    if not framework_resources:
        return 0
    
    with open(resources_go_path, 'r') as f:
        content = f.read()
    
    # Parse existing enum entries
    const_pattern = r'const \(\n(.*?)\n\)'
    match = re.search(const_pattern, content, re.DOTALL)
    existing_enums = set()
    
    if match:
        const_block = match.group(1)
        for line in const_block.split('\n'):
            line = line.strip()
            if line and not line.startswith('//'):
                parts = line.split()
                if parts:
                    existing_enums.add(parts[0])
    
    # Parse existing string entries
    values_pattern = r'var resources = \[\]string\{(.*?)\n\}'
    match = re.search(values_pattern, content, re.DOTALL)
    existing_strings = set()
    
    if match:
        values_block = match.group(1)
        string_pattern = r'"([^"]+)"'
        existing_strings.update(re.findall(string_pattern, values_block))
    
    # Convert framework resources to PascalCase for enum names
    framework_enums = {convert_to_sentence_case_and_retain_underscore(r) for r in framework_resources}
    framework_strings = set(framework_resources)
    
    # Find missing entries
    missing_enums = framework_enums - existing_enums
    missing_strings = framework_strings - existing_strings
    
    if not missing_enums and not missing_strings:
        return 0
    
    # Create complete sorted lists
    all_enums = sorted(existing_enums | framework_enums)
    all_strings = sorted(existing_strings | framework_strings)
    
    # Create mapping for proper ordering
    pascal_to_snake = {convert_to_sentence_case_and_retain_underscore(r): r for r in framework_resources}
    for enum in existing_enums:
        if enum not in pascal_to_snake:
            snake = re.sub(r'([a-z0-9])([A-Z])', r'\1_\2', enum).lower()
            pascal_to_snake[enum] = snake
    
    sorted_enums = sorted(all_enums, key=lambda x: pascal_to_snake.get(x, x.lower()))
    
    # Generate new const block
    new_const_entries = []
    for i, enum in enumerate(sorted_enums):
        if i == 0:
            new_const_entries.append(f'\t{enum} Resource = iota')
        else:
            new_const_entries.append(f'\t{enum}')
    
    new_const_content = '\n'.join(new_const_entries)
    
    # Replace const block
    content = re.sub(
        const_pattern,
        f'const (\n{new_const_content}\n)',
        content,
        flags=re.DOTALL
    )
    
    # Generate new string slice
    new_values_entries = [f'\t"{s}",' for s in all_strings]
    new_values_content = '\n'.join(new_values_entries)
    
    # Replace values slice
    values_pattern = r'(var resources = \[\]string\{)(.*?)(\n\})'
    content = re.sub(
        values_pattern,
        f'var resources = []string{{\n{new_values_content}\n}}',
        content,
        flags=re.DOTALL
    )
    
    # Write back
    with open(resources_go_path, 'w') as f:
        f.write(content)
    
    return len(missing_enums)


def update_provider_go(resources: List[str], provider_file_path: Path):
    """
    Update provider.go to add imports, datasource registrations, and resource registrations.
    Only adds entries that don't already exist to avoid duplicates.
    """
    if not provider_file_path.exists():
        print(f"Warning: provider.go not found at {provider_file_path}")
        return
    
    with open(provider_file_path, 'r') as f:
        content = f.read()
    
    # Track what we need to add
    imports_to_add = []
    datasources_to_add = []
    resources_to_add = []
    
    for resource_name in resources:
        # Generate import path
        import_path = f'"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/{resource_name}"'
        
        # Get the correct datasource constructor name using the same logic as generation
        template_vars = get_datasource_template_variables(resource_name)
        datasource_constructor = template_vars['datasource_constructor_name']
        datasource_func = f'{resource_name}.{datasource_constructor}'
        
        # Get the correct resource constructor name
        resource_template_vars = get_resource_template_variables(resource_name)
        resource_struct_name = resource_template_vars['resource_struct_name']
        resource_func = f'{resource_name}.New{resource_struct_name}'
        
        # Check if import already exists
        if import_path not in content:
            imports_to_add.append((resource_name, import_path))
        
        # Check if datasource registration already exists
        if datasource_func not in content:
            datasources_to_add.append((resource_name, datasource_func))
        
        # Check if resource registration already exists
        if resource_func not in content:
            resources_to_add.append((resource_name, resource_func))
    
    # If nothing to add, return early
    if not imports_to_add and not datasources_to_add and not resources_to_add:
        print("Provider.go is already up to date")
        return
    
    # Add imports
    if imports_to_add:
        # Find the last import line before the closing paren
        import_section_pattern = r'(import \(\n.*?"github\.com/citrix/terraform-provider-citrixadc/citrixadc_framework/[^"]+"\n)'
        matches = list(re.finditer(import_section_pattern, content, re.DOTALL))
        
        if matches:
            last_import_match = matches[-1]
            insert_pos = last_import_match.end()
            
            # Generate new import lines
            new_imports = '\n'.join([f'\t{import_path}' for _, import_path in sorted(imports_to_add)])
            content = content[:insert_pos] + new_imports + '\n' + content[insert_pos:]
            
            print(f"Added {len(imports_to_add)} new imports to provider.go")
    
    # Add datasource registrations
    if datasources_to_add:
        # Find the DataSources function and the return statement
        datasources_pattern = r'(func \(p \*CitrixAdcFrameworkProvider\) DataSources\(ctx context\.Context\) \[\]func\(\) datasource\.DataSource \{\n\treturn \[\]func\(\) datasource\.DataSource\{\n(?:.*?\n)*?)(\t\}\n\})'
        
        match = re.search(datasources_pattern, content, re.DOTALL)
        if match:
            # Insert before the closing braces
            insert_pos = match.end(1)
            
            # Generate new datasource lines
            new_datasources = '\n'.join([f'\t\t{datasource_func},' for _, datasource_func in sorted(datasources_to_add)])
            content = content[:insert_pos] + new_datasources + '\n' + content[insert_pos:]
            
            print(f"Added {len(datasources_to_add)} new datasource registrations to provider.go")
    
    # Add resource registrations
    if resources_to_add:
        # Find the Resources function and the return statement
        resources_pattern = r'(func \(p \*CitrixAdcFrameworkProvider\) Resources\(ctx context\.Context\) \[\]func\(\) resource\.Resource \{\n\treturn \[\]func\(\) resource\.Resource\{\n(?:.*?\n)*?)(\t\}\n\})'
        
        match = re.search(resources_pattern, content, re.DOTALL)
        if match:
            # Insert before the closing braces
            insert_pos = match.end(1)
            
            # Generate new resource lines
            new_resources = '\n'.join([f'\t\t{resource_func},' for _, resource_func in sorted(resources_to_add)])
            content = content[:insert_pos] + new_resources + '\n' + content[insert_pos:]
            
            print(f"Added {len(resources_to_add)} new resource registrations to provider.go")
    
    # Write back the modified content
    with open(provider_file_path, 'w') as f:
        f.write(content)
    
    print(f"Updated provider.go successfully")


# ============================================================================
# FILE GENERATION FUNCTIONS
# ============================================================================

def generate_datasource_implementation(resource_name, template, output_dir, raw_attributes=None):
    """Generate datasource_<resource>.go file."""
    template_vars = get_datasource_template_variables(resource_name, raw_attributes)
    content = template.render(**template_vars)
    
    resource_dir = output_dir / resource_name
    resource_dir.mkdir(parents=True, exist_ok=True)
    
    output_file = resource_dir / f'datasource_{resource_name}.go'
    with open(output_file, 'w') as f:
        f.write(content)
    
    return output_file


def generate_datasource_schema(resource_name, template, attributes, output_dir):
    """Generate datasource_schema.go file."""
    package_name = resource_name
    function_name = ''.join(word.capitalize() for word in resource_name.split('_')) + 'DataSourceSchema'
            
    # Process attributes for datasource
    processed_attributes = [process_attribute_for_datasource(attr) for attr in attributes]
    
    # Expand secret attributes to include _wo and _wo_version variants (for model consistency)
    processed_attributes = expand_secret_attributes(processed_attributes)
    
    content = template.render(
        package_name=package_name,
        function_name=function_name,
        attributes=processed_attributes,
        resource_name=resource_name
    )
    
    resource_dir = output_dir / resource_name
    resource_dir.mkdir(parents=True, exist_ok=True)
    
    output_file = resource_dir / 'datasource_schema.go'
    with open(output_file, 'w') as f:
        f.write(content)
    
    return output_file


def generate_resource_implementation(resource_name, template, output_dir, raw_attributes=None):
    """Generate resource_<resource>.go file."""
    template_vars = get_resource_template_variables(resource_name, raw_attributes)
    content = template.render(**template_vars)
    
    resource_dir = output_dir / resource_name
    resource_dir.mkdir(parents=True, exist_ok=True)
    
    output_file = resource_dir / f'resource_{resource_name}.go'
    with open(output_file, 'w') as f:
        f.write(content)
    
    return output_file


def generate_resource_schema(resource_name, template, attributes, output_dir):
    """Generate resource_schema.go file."""
    package_name = resource_name
    model_name = ''.join(word.capitalize() for word in resource_name.split('_')) + 'ResourceModel'
    resource_type_name = ''.join(word.capitalize() for word in resource_name.split('_')) + 'Resource'
    
    # Process attributes for resource
    processed_attributes = [process_attribute_for_resource(attr) for attr in attributes]
    
    # Expand secret attributes to include _wo and _wo_version variants
    processed_attributes = expand_secret_attributes(processed_attributes)
    
    # Get additional template variables
    nitro_info = get_nitro_package_info(resource_name)
    function_names = {
        'get_payload_func': f'{resource_name}GetThePayloadFromthePlan',
        'set_attr_func': f'{resource_name}SetAttrFromGet'
    }
    required_imports = get_required_imports_for_resource(processed_attributes)
    id_generation = analyze_id_generation_pattern(attributes)
    
    content = template.render(
        package_name=package_name,
        model_name=model_name,
        resource_type_name=resource_type_name,
        attributes=processed_attributes,
        resource_name=resource_name,
        required_imports=required_imports,
        nitro_info=nitro_info,
        function_names=function_names,
        id_generation=id_generation
    )
    
    resource_dir = output_dir / resource_name
    resource_dir.mkdir(parents=True, exist_ok=True)
    
    output_file = resource_dir / 'resource_schema.go'
    with open(output_file, 'w') as f:
        f.write(content)
    
    return output_file


def generate_all_files_for_resource(resource_name, templates, tfdata_dir, output_dir):
    """Generate all files for a single resource."""
    print(f"Processing: {resource_name}")
    
    generated_files = []
    
    try:
        # Load metadata
        raw_attributes = load_resource_metadata(tfdata_dir, resource_name)
        if not raw_attributes:
            print(f"  Skipping {resource_name} - no valid attributes found")
            return []

        # processed_attributes = []
        # for attr in raw_attributes:
        #     if "appfwprofile_" in resource_name and "_binding" in resource_name and resource_name != "appfwprofile_sqlinjection_binding " and attr["option_name"] == "ruletype":
        #         continue
        #     if "authenticationvserver_" in resource_name and "_binding" in resource_name and resource_name not in ["authenticationvserver_cachepolicy_binding", "authenticationvserver_rewritepolicy_binding"] and attr["option_name"] == "bindpoint":
        #         continue
        #     if "crvserver_" in resource_name and "_binding" in resource_name and resource_name != "crvserver_rewritepolicy_binding" and attr["option_name"] == "bindpoint":
        #         continue
        #     if "csvserver_" in resource_name and "_binding" in resource_name and resource_name not in ["csvserver_cachepolicy_binding", "csvserver_rewritepolicy_binding", "csvserver_cmppolicy_binding", "csvserver_contentinspectionpolicy_binding"] and attr["option_name"] == "bindpoint":
        #         continue
        #     if "lbvserver_" in resource_name and "_binding" in resource_name and resource_name not in ["lbvserver_cachepolicy_binding", "lbvserver_rewritepolicy_binding", "lbvserver_cmppolicy_binding", "lbvserver_contentinspectionpolicy_binding", "lbvserver_videooptimizationdetectionpolicy_binding", "lbvserver_videooptimizationpacingpolicy_binding"] and attr["option_name"] == "bindpoint":
        #         continue
        #     if "vpnvserver_" in resource_name and "_binding" in resource_name and resource_name not in ["vpnvserver_cachepolicy_binding", "vpnvserver_rewritepolicy_binding"] and attr["option_name"] == "bindpoint":
        #         continue
            # if "vpnglobal_" in resource_name and "_binding" in resource_name and  attr["option_name"] in ["groupextraction", "secondary"]:
            #     attr["x-unique-attr"] = False
            #     processed_attributes.append(attr)
            #     continue
            # if "vpnvserver_" in resource_name and "_binding" in resource_name and attr["option_name"] in ["groupextraction", "secondary"]:
            #     attr["x-unique-attr"] = False
            #     processed_attributes.append(attr)
            #     continue
            # if "aaagroup_" in resource_name and "_binding" in resource_name and attr["option_name"] in ["type", "netmask", "numaddr"]:
            #     attr["x-unique-attr"] = False
            #     processed_attributes.append(attr)
            #     continue
            # if "aaauser_" in resource_name and "_binding" in resource_name and attr["option_name"] in ["type", "netmask", "numaddr"]:
            #     attr["x-unique-attr"] = False
            #     processed_attributes.append(attr)
            #     continue
            # if "appfwglobal_" in resource_name and "_binding" in resource_name and attr["option_name"] == "priority":
            #     attr["x-unique-attr"] = False
            #     processed_attributes.append(attr)
            #     continue
            # if "crvserver_" in resource_name and "_binding" in resource_name and attr["option_name"] == "priority":
            #     attr["x-unique-attr"] = False
            #     processed_attributes.append(attr)
            #     continue
            # if "csvserver_" in resource_name and "_binding" in resource_name and attr["option_name"] == "priority":
            #     attr["x-unique-attr"] = False
            #     processed_attributes.append(attr)
            #     continue
            # if "lbvserver_" in resource_name and "_binding" in resource_name and attr["option_name"] == "priority":
            #     attr["x-unique-attr"] = False
            #     processed_attributes.append(attr)
            #     continue
            # if resource_name in ["appfwpolicylabel_appfwpolicy_binding", "appflowpolicylabel_appflowpolicy_binding", "authenticationpolicylabel_authenticationpolicy_binding", "authorizationpolicylabel_authorizationpolicy_binding", "botglobal_botpolicy_binding", "botpolicylabel_botpolicy_binding", "cacheglobal_cachepolicy_binding", "cachepolicylabel_cachepolicy_binding", "cmpglobal_cmppolicy_binding", "cmppolicylabel_cmppolicy_binding", "contentinspectionglobal_contentinspectionpolicy_binding", "contentinspectionpolicylabel_contentinspectionpolicy_binding", "dnsglobal_dnspolicy_binding", "dnspolicylabel_dnspolicy_binding", "feoglobal_feopolicy_binding", "icaglobal_icapolicy_binding", "transformglobal_transformpolicy_binding", "transformpolicylabel_transformpolicy_binding", "tunnelglobal_tunneltrafficpolicy_binding"] and attr["option_name"] == "priority":
            #     attr["x-unique-attr"] = False
            #     processed_attributes.append(attr)
            #     continue
        #     processed_attributes.append(attr)


        # raw_attributes = processed_attributes
        
        # Generate datasource implementation (pass attributes for pattern analysis)
        file_path = generate_datasource_implementation(resource_name, templates['datasource'], output_dir, raw_attributes)
        generated_files.append(file_path)
        print(f"  Generated: {file_path}")
        
        # Generate datasource schema
        file_path = generate_datasource_schema(resource_name, templates['datasource_schema'], raw_attributes, output_dir)
        generated_files.append(file_path)
        print(f"  Generated: {file_path}")
        
        # Generate resource implementation
        file_path = generate_resource_implementation(resource_name, templates['resource'], output_dir, raw_attributes)
        generated_files.append(file_path)
        print(f"  Generated: {file_path}")
        
        # Generate resource schema
        file_path = generate_resource_schema(resource_name, templates['resource_schema'], raw_attributes, output_dir)
        generated_files.append(file_path)
        print(f"  Generated: {file_path}")
        
    except Exception as e:
        print(f"  Error processing {resource_name}: {e}")
        return []
    
    return generated_files


# ============================================================================
# MAIN FUNCTION
# ============================================================================

def main():
    """Main function to process all resources and generate all files."""
    
    # Get script directory
    script_dir = Path(__file__).parent
    
    # Define paths
    resources_file = script_dir / 'resources.txt'
    tfdata_dir = script_dir / 'tfdata'
    output_dir = script_dir / 'citrixadc_framework'
    log_file = script_dir / 'code_generation.log'
    
    # Template files
    template_files = {
        'datasource': 'datasource.go.j2',
        'datasource_schema': 'datasource_schema.go.j2',
        'resource': 'resource.go.j2',
        'resource_schema': 'resource_schema.go.j2'
    }
    
    # Check required files exist
    if not resources_file.exists():
        print(f"Error: {resources_file} not found")
        sys.exit(1)
    
    if not tfdata_dir.exists():
        print(f"Error: {tfdata_dir} directory not found")
        sys.exit(1)
    
    # Check all template files exist
    template_dir = script_dir / 'jinja_templates'
    for template_name, template_file in template_files.items():
        template_path = template_dir / template_file
        if not template_path.exists():
            print(f"Error: {template_file} template not found")
            sys.exit(1)
    
    # Load all Jinja templates
    env = Environment(
        loader=FileSystemLoader(template_dir),
        trim_blocks=False,
        lstrip_blocks=False
    )
    
    templates = {}
    for template_name, template_file in template_files.items():
        templates[template_name] = env.get_template(template_file)
    
    # Read resources list
    try:
        with open(resources_file, 'r') as f:
            resources = [line.strip() for line in f if line.strip()]
            # resources = [line.strip() for line in f if line.strip() and '_binding' in line.strip()]
    except Exception as e:
        print(f"Error reading {resources_file}: {e}")
        sys.exit(1)
    
    if not resources:
        print("No resources found in resources.txt")
        sys.exit(1)
    
    print(f"Processing {len(resources)} resources...")
    print(f"Generating 4 files per resource (datasource implementation, datasource schema, resource implementation, resource schema)")
    print()
    
    # Open log file
    with open(log_file, 'w') as log:
        log.write("CitrixADC Terraform Provider Code Generation Log\n")
        log.write("=" * 80 + "\n\n")
        
        # Process each resource
        successful_resources = 0
        total_files_generated = 0
        failed_resource_list = []
    
        for resource_name in resources:
            generated_files = generate_all_files_for_resource(resource_name, templates, tfdata_dir, output_dir)
            
            if generated_files:
                successful_resources += 1
                total_files_generated += len(generated_files)
                log.write(f"[SUCCESS] {resource_name}\n")
                for file_path in generated_files:
                    log.write(f"  - {file_path}\n")
            else:
                failed_resource_list.append(resource_name)
                log.write(f"[FAILED] {resource_name}\n")
            
            log.flush()  # Ensure log is written immediately
            print()  # Add spacing between resources
    
        # Summary
        log.write("\n" + "=" * 80 + "\n")
        log.write("Summary\n")
        log.write("=" * 80 + "\n")
        log.write(f"Resources processed: {successful_resources}/{len(resources)}\n")
        log.write(f"Total files generated: {total_files_generated}\n")
        log.write(f"Expected files per resource: 4\n")
        
        if successful_resources < len(resources):
            failed_resources = len(resources) - successful_resources
            log.write(f"Failed resources: {failed_resources}\n\n")
            log.write("Failed resource list:\n")
            for failed_resource in failed_resource_list:
                log.write(f"  - {failed_resource}\n")
        else:
            log.write("\nAll resources processed successfully!\n")
    
    # Print summary to console
    print("=== Summary ===")
    print(f"Resources processed: {successful_resources}/{len(resources)}")
    print(f"Total files generated: {total_files_generated}")
    print(f"Expected files per resource: 4")
    
    if successful_resources < len(resources):
        failed_resources = len(resources) - successful_resources
        print(f"Failed resources: {failed_resources}")
        print("\nFailed resource list:")
        for failed_resource in failed_resource_list:
            print(f"  - {failed_resource}")
        print(f"\nDetailed log written to: {log_file}")
    
    if successful_resources > 0:
        print("\nAll resources processed successfully!")
        print(f"Detailed log written to: {log_file}")
        
        # Clean up unused utils imports
        print("\n=== Cleaning up unused utils imports ===")
        modified_count = cleanup_unused_utils_imports(output_dir)
        print(f"Removed unused utils imports from {modified_count} files")
        
        # Update resources.go with all framework resources
        print("\n=== Updating vendor resources.go ===")
        resources_go_path = script_dir / 'vendor' / 'github.com' / 'citrix' / 'adc-nitro-go' / 'service' / 'resources.go'
        if resources_go_path.exists():
            new_resources_added = update_resources_go_file(output_dir, resources_go_path)
            if new_resources_added > 0:
                print(f"Added {new_resources_added} new resource entries to resources.go")
            else:
                print("resources.go is already up to date")
        else:
            print(f"Warning: resources.go not found at {resources_go_path}")
        
        # Update provider.go with successful resources
        successful_resource_list = [res for res in resources if res not in failed_resource_list]
        provider_file = script_dir / 'citrixadc_framework' / 'provider' / 'provider.go'
        print("\n=== Updating provider.go ===")
        update_provider_go(successful_resource_list, provider_file)
    
    if successful_resources < len(resources):
        sys.exit(1)


if __name__ == '__main__':
    main()