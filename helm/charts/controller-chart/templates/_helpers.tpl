{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "controller.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "controller.selectorLabels" -}}
app.kubernetes.io/name: {{ include "controller.fullname" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Build HTTP KAFKA URL
*/}}
{{- define "controller.buildHttpKafkaUrl" -}}
{{- "http://" }}
{{- .Values.global.processor.kafka.kafkaBridge.routeName }}{{ "-" }}
{{- .Values.global.namespace }}{{ "." }}
{{- .Values.global.domain }}
{{- end }}
