{{/*
Expand the name of the chart.
*/}}
{{- define "banking-service.name" -}}
{{- default .Chart.Name .Values.global.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Validate the values provided by the user.
*/}}
{{- define "banking-service.validateValues" -}}
{{- if not .Values.global.projectName }}
  {{- fail "global.projectName is required in values.yaml" }}
{{- end }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "banking-service.fullname" -}}
{{- include "banking-service.validateValues" . }}
{{- if .Values.global.fullnameOverride }}
{{- .Values.global.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := include "banking-service.name" . }}
{{- $suffix := .Values.global.projectName | trimSuffix "-" }}
{{- printf "%s-%s" $name $suffix | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "banking-service.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "banking-service.labels" -}}
helm.sh/chart: {{ include "banking-service.chart" . }}
{{ include "banking-service.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "banking-service.selectorLabels" -}}
app.kubernetes.io/name: {{ include "banking-service.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "banking-service.serviceAccountName" -}}
{{- if .Values.app.serviceAccount.enabled }}
{{- default (include "banking-service.fullname" .) .Values.app.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.app.serviceAccount.name }}
{{- end }}
{{- end }}
