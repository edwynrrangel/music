{{/*
Create a default suffix using the name of the chart.
*/}}
{{- define "pvc.suffix" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Validate the values provided by the user.
*/}}
{{- define "pvc.validateValues" -}}
{{- if not .Values.pvc.projectName }}
{{- fail "global.pvc.projectName is required in values.yaml" }}
{{- end }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "pvc.fullname" -}}
{{- include "pvc.validateValues" . }}
{{- $name := default .Values.pvc.projectName .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- $suffix := include "pvc.suffix" . }}
{{- printf "%s-%s" $name $suffix | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "pvc.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "pvc.labels" -}}
helm.sh/chart: {{ include "pvc.chart" . }}
{{ include "pvc.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "pvc.selectorLabels" -}}
app.kubernetes.io/name: {{ include "pvc.fullname" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}
