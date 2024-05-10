{{/*
Expand the name of the chart.
*/}}
{{- define "grpc-service.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Validate the values provided by the user.
*/}}
{{- define "grpc-service.validateValues" -}}
{{- if not .Values.global.projectName }}
  {{- fail "global.projectName is required in values.yaml" }}
{{- end }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "grpc-service.fullname" -}}
{{- include "grpc-service.validateValues" . }}
{{- if .Values.global.fullnameOverride }}
{{- .Values.global.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.global.nameOverride | trunc 63 | trimSuffix "-" }}
{{- $suffix := .Values.global.projectName | trimSuffix "-" }}
{{- printf "%s-%s" $name $suffix | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "grpc-service.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "grpc-service.labels" -}}
helm.sh/chart: {{ include "grpc-service.chart" . }}
{{ include "grpc-service.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "grpc-service.selectorLabels" -}}
app.kubernetes.io/name: {{ include "grpc-service.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "grpc-service.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "grpc-service.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}
