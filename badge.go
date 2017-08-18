package main

import (
	"github.com/ajstarks/svgo"
	"fmt"
	"errors"
	"strings"
	"regexp"
	"io"
)

type Banner struct {
	fontSize         int
	paddingWidth     int
	paddingHeight    int
	text1            string
	text2            string
	textColor1       string
	textColor2       string
	backgroundColor1 string
	backgroundColor2 string
}

const (
	fontSpacing = 2
)

var colors = map[string][]string{
	"1": {"#fff", "#fff", "#a7bfc1", "#16a085"},
	"2": {"#fff", "#fff", "#a7bfc1", "#27ae60"},
	"3": {"#fff", "#fff", "#a7bfc1", "#2980b9"},
	"4": {"#fff", "#fff", "#a7bfc1", "#8e44ad"},
	"5": {"#fff", "#fff", "#a7bfc1", "#2c3e50"},
	"6": {"#fff", "#fff", "#a7bfc1", "#f39c12"},
	"7": {"#fff", "#fff", "#a7bfc1", "#d35400"},
	"8": {"#fff", "#fff", "#a7bfc1", "#c0392b"},
	"9": {"#fff", "#fff", "#a7bfc1", "#7f8c8d"},

	"10": {"#fff", "#fff", "#1abc9c", "#16a085"},
	"11": {"#fff", "#fff", "#2ecc71", "#27ae60"},
	"12": {"#fff", "#fff", "#3498db", "#2980b9"},
	"13": {"#fff", "#fff", "#9b59b6", "#8e44ad"},
	"14": {"#fff", "#fff", "#34495e", "#2c3e50"},
	"15": {"#fff", "#fff", "#f1c40f", "#f39c12"},
	"16": {"#fff", "#fff", "#e67e22", "#d35400"},
	"17": {"#fff", "#fff", "#e74c3c", "#c0392b"},
	"18": {"#fff", "#fff", "#a7bfc1", "#7f8c8d"},

}

func Short(text string, i int) (string, error) {
	var short string
	if len(text) > i {
		short = text[:i]
	} else {
		short = text
	}

	reg, err := regexp.Compile("[^a-zA-Z0-9_\\-\\s]+")
	if err != nil {
		return "", err
	}

	return reg.ReplaceAllString(strings.ToUpper(short), ""), nil
}

func CreateBanner(text1 string, text2 string, size string, color string) (*Banner, error) {
	banner := &Banner{
		fontSize:         0,
		paddingWidth:     0,
		paddingHeight:    0,
		backgroundColor1: "#000",
		backgroundColor2: "#000",
		textColor1:       "#fff",
		textColor2:       "#fff",
		text1:            "",
		text2:            "",
	}

	text1, err := Short(text1, 15)
	if err != nil {
		return nil, errors.New("invalid text1")
	}

	text2, err = Short(text2, 15)
	if err != nil {
		return nil, errors.New("invalid text2")
	}

	banner.text1 = text1
	banner.text2 = text2

	if size == "sm" {
		banner.fontSize = 8
		banner.paddingWidth = 10
		banner.paddingHeight = 2
	} else if size == "md" {
		banner.fontSize = 15
		banner.paddingWidth = 15
		banner.paddingHeight = 4
	} else if size == "lg" {
		banner.fontSize = 30
		banner.paddingWidth = 30
		banner.paddingHeight = 5
	} else {
		return nil, errors.New("invalid size")
	}

	colorMap, ok := colors[color]
	if ok {
		banner.textColor1 = colorMap[0]
		banner.textColor2 = colorMap[1]
		banner.backgroundColor1 = colorMap[2]
		banner.backgroundColor2 = colorMap[3]
	} else {
		return nil, errors.New("invalid color")
	}

	return banner, nil
}

func generateBanner(w io.Writer, banner *Banner) {
	canvas := svg.New(w)

	var (
		fontOffset      = int(float32(banner.fontSize) * 0.25)
		charWidthFactor = float32(0.6)
		textWidthFactor = float32(banner.fontSize)*charWidthFactor + fontSpacing
		paddingWidth    = banner.paddingWidth * 2
		textOffsetY     = banner.fontSize + banner.paddingHeight
	)

	width1 := int(float32(len(banner.text1))*textWidthFactor) + paddingWidth
	width2 := int(float32(len(banner.text2))*textWidthFactor) + paddingWidth

	width := width1 + width2
	height := banner.paddingHeight*2 + int(banner.fontSize) + fontOffset

	canvas.Start(width, height)
	canvas.Rect(0, 0, width1, height, fmt.Sprintf("fill: %v", banner.backgroundColor1))
	canvas.Rect(width1, 0, width2, height, fmt.Sprintf("fill: %v", banner.backgroundColor2))
	canvas.Gstyle(fmt.Sprintf("font-family:Monaco,monospace;alignment-baseline:center;font-size:%vpx;text-anchor:middle;letter-spacing:%vpx;", banner.fontSize, fontSpacing))
	canvas.Text(width1/2, textOffsetY, banner.text1, fmt.Sprintf("fill:%v;font-weight:300;", banner.textColor1))
	canvas.Text(width1+width2/2, textOffsetY, banner.text2, fmt.Sprintf("fill:%v;font-weight:700", banner.textColor2))
	canvas.Gend()
	canvas.End()
}
