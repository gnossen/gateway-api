#!/bin/bash

set -euxo pipefail

GCP_PROJECT=rbellevi-gke-dev
BASE_TAG=gcr.io/${GCP_PROJECT}/gateway-conformance

go test -c conformance/conformance_test.go

HASH=$(shasum conformance.test | awk '{print $1;}')
TAG="${BASE_TAG}:${HASH}"

TMPDIR=$(mktemp -d)

cp conformance.test "${TMPDIR}"
cp Dockerfile.conformance "${TMPDIR}"

docker build -t "${TAG}" -f Dockerfile.conformance "${TMPDIR}"

docker push "${TAG}"

sed "s|IMAGEGOESHERE|${TAG}|g" conformance-job.yaml >"${TMPDIR}/conformance-job.yaml"

kubectl delete job istio-conformance || true

kubectl apply -f "${TMPDIR}/conformance-job.yaml"

# Doesn't work for jobs.
# kubectl wait --for=condition=ready job istio-conformance

while ! kubectl describe job istio-conformance | grep SuccessfulCreate; do sleep 1; done

while ! kubectl logs -f job/istio-conformance; do sleep 2; done
