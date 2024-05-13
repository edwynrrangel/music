{{/*
Expand the name of the chart.
*/}}
{{- define "grpc-server.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Validate the values provided by the user.
*/}}
{{- define "grpc-server.validateValues" -}}
{{- if not .Values.global.projectName }}
  {{- fail "global.projectName is required in values.yaml" }}
{{- end }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "grpc-server.fullname" -}}
{{- include "grpc-server.validateValues" . }}
{{- if .Values.global.fullnameOverride }}
{{- .Values.global.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := include "grpc-server.name" . }}
{{- $suffix := .Values.global.projectName | trimSuffix "-" }}
{{- printf "%s-%s" $name $suffix | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "grpc-server.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "grpc-server.labels" -}}
helm.sh/chart: {{ include "grpc-server.chart" . }}
{{ include "grpc-server.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "grpc-server.selectorLabels" -}}
app.kubernetes.io/name: {{ include "grpc-server.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}
