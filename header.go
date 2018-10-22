/**
 * Copyright 2018 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy of
 * the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */

package main

import (
	"myitcv.io/react"
)

// HeaderDef is the definition for the Header component
type HeaderDef struct {
	react.ComponentDef
}

// HeaderProps is the prop types for the Header component
type HeaderProps struct {
	path string
}

// Header creates instances of the Header component
func Header(h HeaderProps) *HeaderElem {
	return buildHeaderElem(HeaderProps{path: h.path})
}

// Render renders the Header component
func (f HeaderDef) Render() react.Element {
	return react.Header(&react.HeaderProps{ClassName: "header"},
		f.renderNav(),
	)
}

func (f HeaderDef) genLink(name string, link string, storyType string) react.Element {
	props := f.Props()

	class := ""

	if storyType == props.path {
		class = "active"
	}

	return react.A(
		&react.AProps{Href: link, ClassName: class},
		react.S(name),
	)
}

func (f HeaderDef) renderNav() *react.NavElem {
	links := []react.Element{
		react.A(
			&react.AProps{Href: "#/"},
			react.Img(&react.ImgProps{Src: "assets/logo.png", ClassName: "logo", Alt: "Golang logo"}),
		),
		f.genLink("Top", "#/news/1", "news"),
		f.genLink("New", "#/newest/1", "newest"),
		f.genLink("Show", "#/show/1", "show"),
		f.genLink("Ask", "#/ask/1", "ask"),
		f.genLink("Jobs", "#/jobs/1", "jobs"),
	}

	return react.Nav(nil,
		react.Div(&react.DivProps{ClassName: "inner"},
			links...,
		),
		react.A(
			&react.AProps{Href: "https://github.com/GoogleChromeLabs/go-hackernews", ClassName: "github"},
			react.S("Built with GopherJS"),
		),
	)
}
