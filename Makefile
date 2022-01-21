mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir := $(dir $(mkfile_path))

version := v1.1.1
contour := v1.1.0

.PHONY: build-knative-helm
build-knative-helm:
	@echo "building knative helm"
	docker build -t build-knative-helm:v1 build

.PHONY: run-knative-helm
run-knative-helm: build-knative-helm helm-docs
	@echo "running knative helm"
	docker run -v ${mkfile_dir}/charts/knative/crds:/tmp/crds build-knative-helm:v1 ${version} crds
	docker run -v ${mkfile_dir}/charts/knative/templates:/tmp/templates build-knative-helm:v1 ${version} core
	docker run -v ${mkfile_dir}/charts/knative/templates:/tmp/templates build-knative-helm:v1 ${contour} contour

.PHONY: run-pgo-helm
run-pgo-helm:
	@rm -Rf /tmp/pgo
	git clone https://github.com/CrunchyData/postgres-operator-examples.git /tmp/pgo
	@cp -Rf /tmp/pgo/helm/install ${mkfile_dir}/charts/pgo

# https://github.com/CrunchyData/postgres-operator-examples/tree/main/helm/install

.PHONY: helm-docs
helm-docs: ## Generates helm documentation
helm-docs:
	GO111MODULE=on go get github.com/norwoodj/helm-docs/cmd/helm-docs
	helm-docs kubernetes/charts
