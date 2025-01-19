package main

import (
	"awesomeProject/designPatterns/strategy/strategy"
	"strings"
)

type OutputFormat int

const (
	Markdown OutputFormat = iota
	Html
)

type TextProcessor struct {
	builder      strings.Builder
	listStrategy strategy.ListStrategy
}

func NewTextProcessor(listStrategy strategy.ListStrategy) *TextProcessor {
	return &TextProcessor{strings.Builder{}, listStrategy}
}

func (t *TextProcessor) SetOutputFormat(fmt OutputFormat) {
	switch fmt {
	case Markdown:
		t.listStrategy = &strategy.MarkdownListStrategy{}
	case Html:
		t.listStrategy = &strategy.HtmlListStrategy{}
	}
}

func (t *TextProcessor) AppendList(items []string) {
	t.listStrategy.Start(&t.builder)
	for _, item := range items {
		t.listStrategy.AddListItem(&t.builder, item)
	}
	t.listStrategy.End(&t.builder)
}

func (t *TextProcessor) Reset() {
	t.builder.Reset()
}

func (t *TextProcessor) String() string {
	return t.builder.String()
}
