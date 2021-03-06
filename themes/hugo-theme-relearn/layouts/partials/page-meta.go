{{- $currentNode := . }}
{{- $currentNode.Scratch.Set "relearnIsSelfFound" nil }}
{{- $currentNode.Scratch.Set "relearnPrevPage" nil }}
{{- $currentNode.Scratch.Set "relearnNextPage" nil }}
{{- $currentNode.Scratch.Set "relearnSubPages" nil }}
{{- template "relearn-structure" dict "node" .Site.Home "currentnode" $currentNode "hiddenstem" false "hiddencurrent" false }}
{{- define "relearn-structure" }}
	{{- $currentNode := .currentnode }}
	{{- $isSelf := eq $currentNode.RelPermalink .node.RelPermalink }}
	{{- $isDescendant := and (not $isSelf) (.node.IsDescendant $currentNode) }}
	{{- $isAncestor := and (not $isSelf) (.node.IsAncestor $currentNode) }}
	{{- $isOther := and (not $isDescendant) (not $isSelf) (not $isAncestor) }}

	{{- if $isSelf }}
		{{- $currentNode.Scratch.Set "relearnIsSelfFound" true }}
	{{- end}}
	{{- $isSelfFound := eq ($currentNode.Scratch.Get "relearnIsSelfFound") true }}
	{{- $isPreSelf := and (not $isSelfFound) (not $isSelf) }}
	{{- $isPostSelf := and ($isSelfFound) (not $isSelf) }}

	{{- $hidden_node := or (.node.Params.hidden) (eq .node.Title "") }}
	{{- $hidden_stem:= or $hidden_node .hiddenstem }}
	{{- $hidden_current_stem:= or $hidden_node .hiddencurrent }}
	{{- $hidden_from_current := or (and $hidden_node (not $isAncestor) (not $isSelf) ) (and .hiddencurrent (or $isPreSelf $isPostSelf $isDescendant) ) }}
	{{- .node.Scratch.Set "relearnIsHiddenNode" $hidden_node}}
	{{- .node.Scratch.Set "relearnIsHiddenStem" $hidden_stem}}
	{{- .node.Scratch.Set "relearnIsHiddenFromCurrent" $hidden_current_stem}}

	{{- if not $hidden_from_current }}
		{{- if $isPreSelf }}
			{{- $currentNode.Scratch.Set "relearnPrevPage" .node }}
		{{- else if and $isPostSelf (eq ($currentNode.Scratch.Get "relearnNextPage") nil) }}
			{{- $currentNode.Scratch.Set "relearnNextPage" .node }}
		{{- end}}
	{{- end }}

	{{- $currentNode.Scratch.Set "relearnSubPages" .node.Pages }}
	{{- if .node.IsHome }}
		{{- $currentNode.Scratch.Set "relearnSubPages" .node.Sections }}
	{{- else if .node.Sections }}
		{{- $currentNode.Scratch.Set "relearnSubPages" (.node.Pages | union .node.Sections) }}
	{{- end }}
	{{- $pages := ($currentNode.Scratch.Get "relearnSubPages") }}

	{{- $defaultOrdersectionsby := .Site.Params.ordersectionsby | default "weight" }}
	{{- $currentOrdersectionsby := .node.Params.ordersectionsby | default $defaultOrdersectionsby }}
	{{- if eq $currentOrdersectionsby "title"}}
		{{- range $pages.ByTitle  }}
			{{- template "relearn-structure" dict "node" . "currentnode" $currentNode "hiddenstem" $hidden_stem "hiddencurrent" $hidden_from_current }}
		{{- end}}
	{{- else}}
		{{- range $pages.ByWeight  }}
			{{- template "relearn-structure" dict "node" . "currentnode" $currentNode "hiddenstem" $hidden_stem "hiddencurrent" $hidden_from_current }}
		{{- end}}
	{{- end }}
{{- end }}