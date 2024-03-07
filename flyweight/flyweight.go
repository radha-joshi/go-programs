package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

type Format struct {
	Style string
	Size  int
	Color string
}

type TextBuffer struct {
	text   string
	format *Format
}

type Editor struct {
	textBufferList []TextBuffer
}

var formatList = make([]*Format, 0)

func GetFormat(style string, size int, color string) *Format {
	for _, item := range formatList {
		if item.Style == style && item.Size == size && item.Color == color {
			log.Info().Msg(fmt.Sprintf("Same Format returned %v", item))
			return item
		}
	}
	format := &Format{
		Style: style,
		Size:  size,
		Color: color,
	}
	formatList = append(formatList, format)
	log.Info().Msg(fmt.Sprintf("New Format created %v", format))
	return format
}

func (e *Editor) AddText(text string, style string, size int, color string) {
	currentFormat := GetFormat(style, size, color)
	textBuffer := TextBuffer{
		text:   text,
		format: currentFormat,
	}
	e.textBufferList = append(e.textBufferList, textBuffer)
}

func (e *Editor) String() string {
	value := ""
	for _, item := range e.textBufferList {
		format := item.format
		if format.Style == "BOLD" {
			value += "<b>"
		}
		value += fmt.Sprintf("<span style=\"size: %d color: %s\">", format.Size, format.Color)
		value += item.text
		value += "</span>"
		if format.Style == "BOLD" {
			value += "</b>"
		}
	}
	return value
}
