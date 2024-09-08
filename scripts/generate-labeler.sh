#!/bin/bash

OUTPUT_FILE=".github/labeler.yml"

mkdir -p .github

echo "# Automatically generated labeler config for Go modules" > $OUTPUT_FILE
echo "" >> $OUTPUT_FILE

find . -name "go.mod" | while read -r go_mod; do
  module_dir=$(dirname "$go_mod")

  if [ "$module_dir" != "." ]; then
    module_name=$(basename "$module_dir")

    echo "$module_name:" >> $OUTPUT_FILE
    echo "  - changed-files:" >> $OUTPUT_FILE
    echo "      - any-glob-to-any-file: '$module_dir/**'" >> $OUTPUT_FILE
    echo "" >> $OUTPUT_FILE
  fi
done

cat <<EOL >> $OUTPUT_FILE
documentation:
  - changed-files:
      - any-glob-to-any-file: '**/*.md'

ci:
  - changed-files:
      - any-glob-to-any-file: '.github/**'
EOL

# Output a success message
echo "Labeler config generated in $OUTPUT_FILE"
