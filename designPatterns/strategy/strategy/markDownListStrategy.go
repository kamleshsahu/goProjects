package strategy

import "strings"

type MarkdownListStrategy struct{}

func (m *MarkdownListStrategy) Start(builder *strings.Builder) {

}

func (m *MarkdownListStrategy) End(builder *strings.Builder) {

}

func (m *MarkdownListStrategy) AddListItem(
	builder *strings.Builder, item string) {
	builder.WriteString(" * " + item + "\n")
}
