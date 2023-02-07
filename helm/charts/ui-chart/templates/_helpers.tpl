{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "ui.fullname" -}}
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
{{- define "ui.selectorLabels" -}}
app.kubernetes.io/name: {{ include "ui.fullname" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}


{{/*
Build PROCESSOR SERVICE URL
*/}}
{{- define "ui.buildProcessorServiceUrl" -}}
{{- "ws://" }}
{{- .Values.global.processor.service.name }}{{ "." }}
{{- .Values.global.namespace }}{{ ".svc.cluster.local:" }}
{{- .Values.global.processor.service.ports.port }}
{{- end }}

{{/*
Build ENGINEERING SERVICE URL
*/}}
{{- define "ui.buildEngineeringServiceUrl" -}}
{{- "http://" }}
{{- .Values.engineering.serviceName }}{{ "." }}
{{- .Values.global.namespace }}{{ ".svc.cluster.local:" }}
{{- .Values.engineering.servicePort }}
{{- end }}