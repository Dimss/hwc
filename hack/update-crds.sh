CONTROLLER_GEN=~/.go/bin/controller-gen

$CONTROLLER_GEN crd \
 paths=github.com/Dimss/hwc/apis/hwc.dev/v1alpha1 \
 output:artifacts:config=manifests