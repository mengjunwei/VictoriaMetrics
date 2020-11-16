// Code generated by qtc from "tags_autocomplete_tags_response.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// TagsAutoCompleteTagsResponse generates response for /tags/autoComplete/tags handler in Graphite Tags API.See https://graphite.readthedocs.io/en/stable/tags.html#auto-complete-support

//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:5
package graphite

//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:5
func StreamTagsAutoCompleteTagsResponse(qw422016 *qt422016.Writer, isPartial bool, labels []string, jsonp string) {
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:6
	if jsonp != "" {
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:6
		qw422016.N().S(jsonp)
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:6
		qw422016.N().S(`(`)
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:6
	}
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:6
	qw422016.N().S(`[`)
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:8
	for i, label := range labels {
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:9
		qw422016.N().Q(label)
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:10
		if i+1 < len(labels) {
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:10
			qw422016.N().S(`,`)
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:10
		}
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:11
	}
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:11
	qw422016.N().S(`]`)
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:13
	if jsonp != "" {
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:13
		qw422016.N().S(`)`)
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:13
	}
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:14
}

//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:14
func WriteTagsAutoCompleteTagsResponse(qq422016 qtio422016.Writer, isPartial bool, labels []string, jsonp string) {
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:14
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:14
	StreamTagsAutoCompleteTagsResponse(qw422016, isPartial, labels, jsonp)
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:14
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:14
}

//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:14
func TagsAutoCompleteTagsResponse(isPartial bool, labels []string, jsonp string) string {
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:14
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:14
	WriteTagsAutoCompleteTagsResponse(qb422016, isPartial, labels, jsonp)
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:14
	qs422016 := string(qb422016.B)
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:14
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:14
	return qs422016
//line app/vmselect/graphite/tags_autocomplete_tags_response.qtpl:14
}