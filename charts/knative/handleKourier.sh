#!/bin/bash

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

curl -L -H 'Cache-Control: no-cache' https://github.com/knative/net-kourier/releases/download/knative-v1.1.0/kourier.yaml > $dir/templates/kourier.yaml

# add helm labels
sed -i 's/^  labels:/  labels:\n    {{- include "knative.labels" . | nindent 4 }}/g' $dir/templates/kourier.yaml

# delete namespace
sed -i '1,25d' $dir/templates/kourier.yaml

# change namespace
sed -i 's/namespace: kourier-system/namespace: {{ .Release.Namespace }}/g' $dir/templates/kourier.yaml

# namespace for components goign in knative-serving
sed -i 's/namespace: knative-serving/namespace: {{ .Release.Namespace }}/g' $dir/templates/kourier.yaml

# namespace fenv for kourier
sed -i 's/value: "kourier-system"/value: {{ .Release.Namespace }}/g' $dir/templates/kourier.yaml
