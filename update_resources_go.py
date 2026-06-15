#!/usr/bin/env python3
"""
Update resources.go to include all resources from citrixadc_framework.

This script:
1. Scans citrixadc_framework/ for resource directories
2. Parses resources.go to find existing entries
3. Adds missing entries to both the const enum and the string slice
4. Maintains alphabetical ordering
"""

from pathlib import Path
import re
from typing import List, Set


def get_citrixadc_resources(framework_dir: Path) -> List[str]:
    """Get list of all resource names from citrixadc_framework subdirectories."""
    resources = []
    
    # Skip these special directories
    skip_dirs = {'provider', 'utils', 'client'}
    
    for subdir in framework_dir.iterdir():
        if subdir.is_dir() and subdir.name not in skip_dirs:
            resources.append(subdir.name)
    
    return sorted(resources)


def convert_to_pascal_case(snake_str: str) -> str:
    """Convert snake_case to sentnce case and retain underscores."""
    components = snake_str.split('_')
    return  components[0].capitalize() + '_' + '_'.join(x for x in components[1:]) if len(components) > 1 else components[0].capitalize()


def parse_existing_enums(content: str) -> Set[str]:
    """Parse existing enum entries from const block."""
    existing = set()
    
    # Find the const block
    const_pattern = r'const \(\n(.*?)\n\)'
    match = re.search(const_pattern, content, re.DOTALL)
    
    if match:
        const_block = match.group(1)
        # Extract resource names (lines that define enum values)
        for line in const_block.split('\n'):
            line = line.strip()
            if line and not line.startswith('//'):
                # Extract the identifier (resource name)
                parts = line.split()
                if parts:
                    resource_name = parts[0]
                    existing.add(resource_name)
    
    return existing


def parse_existing_strings(content: str) -> Set[str]:
    """Parse existing string values from the resources slice."""
    existing = set()
    
    # Find the resources string slice
    values_pattern = r'var resources = \[\]string\{(.*?)\n\}'
    match = re.search(values_pattern, content, re.DOTALL)
    
    if match:
        values_block = match.group(1)
        # Extract strings
        string_pattern = r'"([^"]+)"'
        existing.update(re.findall(string_pattern, values_block))
    
    return existing


def generate_enum_entries(resources: List[str]) -> str:
    """Generate enum constant entries for resources."""
    lines = []
    
    for i, resource in enumerate(resources):
        pascal_name = convert_to_pascal_case(resource)
        
        if i == 0:
            # First entry uses iota
            lines.append(f'\t{pascal_name} Resource = iota')
        else:
            # Subsequent entries just have the name
            lines.append(f'\t{pascal_name}')
    
    return '\n'.join(lines)


def generate_string_slice(resources: List[str]) -> str:
    """Generate string slice entries for resources."""
    lines = []
    
    for resource in resources:
        lines.append(f'\t"{resource}",')
    
    return '\n'.join(lines)


def update_resources_go(resources_go_path: Path, framework_resources: List[str]):
    """Update resources.go with missing entries."""
    
    with open(resources_go_path, 'r') as f:
        content = f.read()
    
    # Parse existing entries
    existing_enums = parse_existing_enums(content)
    existing_strings = parse_existing_strings(content)
    
    # Convert framework resources to PascalCase for enum names
    framework_enums = {convert_to_pascal_case(r) for r in framework_resources}
    framework_strings = set(framework_resources)
    
    # Find missing entries
    missing_enums = framework_enums - existing_enums
    missing_strings = framework_strings - existing_strings
    
    print(f"Found {len(framework_resources)} resources in citrixadc_framework")
    print(f"Found {len(existing_enums)} existing enum entries in resources.go")
    print(f"Found {len(existing_strings)} existing string entries in resources.go")
    print()
    
    if not missing_enums and not missing_strings:
        print("✓ resources.go is already up to date!")
        return
    
    print(f"Missing enum entries: {len(missing_enums)}")
    print(f"Missing string entries: {len(missing_strings)}")
    print()
    
    # Create complete sorted lists
    all_enums = sorted(existing_enums | framework_enums)
    all_strings = sorted(existing_strings | framework_strings)
    
    # Generate new enum entries - need to convert back from PascalCase to snake_case for proper ordering
    # Create mapping of PascalCase to snake_case
    pascal_to_snake = {convert_to_pascal_case(r): r for r in framework_resources}
    for enum in existing_enums:
        if enum not in pascal_to_snake:
            # This is an existing enum not in our framework
            # Try to convert back to snake_case
            snake = enum.lower()
            # Handle multi-word cases
            snake = re.sub(r'([a-z0-9])([A-Z])', r'\1_\2', enum).lower()
            pascal_to_snake[enum] = snake
    
    # Sort by snake_case values
    sorted_enums = sorted(all_enums, key=lambda x: pascal_to_snake.get(x, x.lower()))
    
    # Update const block
    const_pattern = r'(const \(\n)(.*?)(\n\))'
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
    
    # Update string slice
    values_pattern = r'(var resources = \[\]string\{)(.*?)(\n\})'
    new_values_content = generate_string_slice(all_strings)
    
    # Replace values slice
    content = re.sub(
        values_pattern,
        f'var resources = []string{{\n{new_values_content}\n}}',
        content,
        flags=re.DOTALL
    )
    
    # Write back
    with open(resources_go_path, 'w') as f:
        f.write(content)
    
    print(f"✓ Updated resources.go")
    
    if missing_enums:
        print(f"\nAdded enum entries:")
        for enum in sorted(missing_enums):
            print(f"  - {enum}")
    
    if missing_strings:
        print(f"\nAdded string entries:")
        for string in sorted(missing_strings):
            print(f"  - {string}")


def main():
    """Main function."""
    script_dir = Path(__file__).parent
    framework_dir = script_dir / 'citrixadc_framework'
    resources_go_path = script_dir / 'vendor' / 'github.com' / 'citrix' / 'adc-nitro-go' / 'service' / 'resources.go'
    
    if not framework_dir.exists():
        print(f"Error: {framework_dir} not found")
        return 1
    
    if not resources_go_path.exists():
        print(f"Error: {resources_go_path} not found")
        return 1
    
    # Get all resources from framework
    framework_resources = get_citrixadc_resources(framework_dir)
    
    if not framework_resources:
        print("No resources found in citrixadc_framework")
        return 1
    
    # Update resources.go
    update_resources_go(resources_go_path, framework_resources)
    
    return 0


if __name__ == '__main__':
    exit(main())
