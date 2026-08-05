{{- define "config-service.name" -}}
config-service
{{- end }}

{{- define "config-service.fullname" -}}
{{include "config-service.name" .}}
{{- end }}