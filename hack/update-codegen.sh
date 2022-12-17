CODE_GENERATOR_EXEC_DIR="${GOPATH}/src/github.com/code-generator"
GENERATE_GROUPS="${CODE_GENERATOR_EXEC_DIR}/generate-groups.sh"

$GENERATE_GROUPS all \
 github.com/Dimss/hwc/pkg/generated \
 github.com/Dimss/hwc/apis hwc.dev:v1alpha1 \
 --go-header-file "${CODE_GENERATOR_EXEC_DIR}"/hack/boilerplate.go.txt