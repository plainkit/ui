package handlers

import (
	"github.com/plainkit/html"
	"github.com/plainkit/ui/avatar"
)

func RenderAvatarsContent() html.Node {
	return html.Div(
		html.AClass("space-y-10"),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("User avatars")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Fallbacks automatically show initials when no image is available.")),
			),
			html.Div(
				html.AClass("flex flex-wrap items-center gap-6"),
				avatar.Avatar(
					avatar.Props{},
					avatar.Image(avatar.ImageProps{Src: "https://i.pravatar.cc/64?img=12", Alt: "Ryan"}),
					avatar.Fallback(avatar.FallbackProps{}, html.Text("RY")),
				),
				avatar.Avatar(
					avatar.Props{Size: avatar.SizeSm},
					avatar.Fallback(avatar.FallbackProps{}, html.Text("DC")),
				),
			),
		),
		html.Section(
			html.AClass("space-y-4"),
			html.Div(
				html.AClass("space-y-1"),
				html.H2(html.AClass("text-2xl font-semibold"), html.Text("Avatar group")),
				html.P(html.AClass("text-sm text-slate-400"), html.Text("Stack avatars with spacing and overflow indicator.")),
			),
			html.Div(
				html.AClass("flex items-center gap-8"),
				avatar.Group(
					avatar.GroupProps{Spacing: avatar.GroupSpacingMd},
					avatar.Avatar(avatar.Props{InGroup: true}, avatar.Image(avatar.ImageProps{Src: "https://i.pravatar.cc/64?img=32", Alt: "Yara"})),
					avatar.Avatar(avatar.Props{InGroup: true}, avatar.Image(avatar.ImageProps{Src: "https://i.pravatar.cc/64?img=19", Alt: "Lee"})),
					avatar.Avatar(avatar.Props{InGroup: true}, avatar.Fallback(avatar.FallbackProps{}, html.Text("MW"))),
					avatar.GroupOverflow(4, avatar.Props{InGroup: true, Size: avatar.SizeSm}),
				),
			),
		),
	)
}
