package strategy

import "strings"

type csvStrategy struct {
}

func (c csvStrategy) Start(builder *strings.Builder) {
	builder.WriteString(" |")
}

func (c csvStrategy) End(builder *strings.Builder) {
	//TODO implement me
	//panic("implement me")
}

func (c csvStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString(" " + item + " |")
}

func CSVStrategy() ListStrategy {
	return &csvStrategy{}
}
