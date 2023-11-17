#!/bin/bash

# Define the input and output directories
ARTIFACTS_DIR="./artifacts"
OUTPUT_DIR="./pkg"

rm -rf $OUTPUT_DIR

# Create the output directory if it doesn't exist
mkdir -p $OUTPUT_DIR

# Initialize error counter
errors=0

# Function to process JSON files and generate Go bindings
process_file() {
  local contract="$1"
  local subdir="$2"
  
  # Check if the JSON file contains "abi" and "bytecode" properties
  if ! jq 'has("abi") and has("bytecode")' "$contract" | grep -q 'true'; then
    return
  fi
  
  # Extract the contract name from the file name (without the .json extension)
  contract_name=$(basename "$contract" .json)
  contract_name_lowercase=$(echo "$contract_name" | tr '[:upper:]' '[:lower:]')

  # Define output subdirectory and create it if it doesn't exist
  output_subdir="$OUTPUT_DIR/${subdir/@/}"
  output_subdir_lowercase=$(echo "$output_subdir" | tr '[:upper:]' '[:lower:]')
  mkdir -p "$output_subdir_lowercase"

  package_name=$(basename "${subdir/@/}" | cut -d'.' -f1 | tr '[:upper:]' '[:lower:]')

  # Generate the Go binding for the contract
  echo "Compiling $contract_name..."
  cat "$contract" | jq .abi > "$output_subdir_lowercase/$contract_name.abi"
  cat "$contract" | jq .bytecode | tr -d '\"' > "$output_subdir_lowercase/$contract_name.bin"
  abigen --abi "$output_subdir_lowercase/$contract_name.abi" --bin "$output_subdir_lowercase/$contract_name.bin" --pkg "$package_name" --type "$contract_name" --out "$output_subdir_lowercase/$contract_name_lowercase.go" > /dev/null 2>&1
  # Check if there were errors during the compilation
  if [ $? -ne 0 ]; then
    echo "Error: Failed to compile $contract_name"
    errors=$((errors + 1))
  fi

  # Remove temporary .abi and .bin files
  rm "$output_subdir_lowercase/$contract_name.abi" "$output_subdir_lowercase/$contract_name.bin"
}

# Function to iterate through the artifacts directory
iterate_directory() {
  local parent="$1"
  local subdir="$2"

  for entry in "$parent"/*; do
    if [ -d "$entry" ]; then
      local new_subdir
      if [ -z "$subdir" ]; then
        new_subdir=$(basename "$entry")
      else
        new_subdir="$subdir/$(basename "$entry")"
      fi
      iterate_directory "$entry" "$new_subdir"
    elif [ -f "$entry" ] && [ "${entry##*.}" = "json" ]; then
      process_file "$entry" "$subdir"
    fi
  done
}

echo -e ""

iterate_directory "$ARTIFACTS_DIR"

echo -e ""

if [ $errors -eq 0 ]; then
  echo "All contracts have been compiled successfully."
else
  echo "There were $errors error(s) during the compilation process."
fi

echo -e ""