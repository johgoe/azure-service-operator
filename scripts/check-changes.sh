#!/bin/bash
set -e

IGNORE_FILTERS=("docs/" "README.md" "hack/" "v2/" "Taskfile.yml" "scripts/check_samples.py" "scripts/make-multitenant-cluster.sh" "scripts/make-multitenant-tenant.sh" "scripts/check-changes.sh" "scripts/gen-api-docs.sh" "scripts/kustomize-build.sh" "scripts/generate-helm-manifest.sh" ".github/" ".devcontainer/" "workspace.code-workspace" ".golangci.yml" "scripts/wait-for-ca-bundles.sh" "scripts/wait-for-operator-ready.sh" "scripts/wrap-resource-urls.py" "dev.sh")
CHANGED_FILES=$(git diff HEAD HEAD~ --name-only)
IGNORED_COUNT=0
NON_IGNORED_COUNT=0
echo "Checking for file changes..."
for FILE in $CHANGED_FILES; do
  # Check if the file matches one of the ignore filters
  MATCHED=0
  for FILTER in ${IGNORE_FILTERS[@]}; do
    if [[ $FILE == *$FILTER* ]]; then
      MATCHED=1
      echo "${FILE} in ignore filter $FILTER"
    fi
  done

  if [[ $MATCHED -eq 0 ]]; then
    echo "Source code file ${FILE} changed"
    NON_IGNORED_COUNT=$(($NON_IGNORED_COUNT+1))
  else
    IGNORED_COUNT=$(($IGNORED_COUNT+1))

  fi
done

echo "" # Blank line for readability
echo "$IGNORED_COUNT match(es) for ignore filter '${IGNORE_FILTERS[*]}' found."
echo "$NON_IGNORED_COUNT match(es) for changed source code files found."
if [[ $NON_IGNORED_COUNT -gt 0 ]]; then
  echo "##vso[task.setvariable variable=SOURCE_CODE_CHANGED;isOutput=true]true"
else
  echo "##vso[task.setvariable variable=SOURCE_CODE_CHANGED;isOutput=true]false"
fi
