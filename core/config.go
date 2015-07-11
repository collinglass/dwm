package core

import (
	"os/exec"
)

/* appearance */
const (
	font            = "-*-terminus-medium-r-*-*-16-*-*-*-*-*-*-*"
	normbordercolor = "#444444"
	normbgcolor     = "#222222"
	normfgcolor     = "#bbbbbb"
	selbordercolor  = "#005577"
	selbgcolor      = "#005577"
	selfgcolor      = "#eeeeee"

	borderpx uint = 1    /* border pixel of windows */
	snap     uint = 32   /* snap pixel */
	showbar       = true /* false means no bar */
	topbar        = true /* false means bottom bar */

	/* layout(s) */
	mfact       = 0.55 /* factor of master area size [0.05..0.95] */
	nmaster     = 1    /* number of clients in master area */
	resizehints = true /* true means respect size hints in tiled resizals */

	// X keyboard events
	XK_Return = 1 + iota
	XK_Tab
	XK_space
	XK_comma
	XK_period

	XK_p
	XK_b
	XK_j
	XK_k
	XK_i
	XK_d
	XK_h
	XK_l
	XK_c
	XK_t
	XK_f
	XK_m
	XK_q

	XK_0
	XK_1
	XK_2
	XK_3
	XK_4
	XK_5
	XK_6
	XK_7
	XK_8
	XK_9

	// when ctrl or shift are being held
	ControlMask
	ShiftMask

	Button1
	Button2
	Button3
)

func TAGKEYS(KEY KeySym, TAG uint) []Key {
	return []Key{
		Key{MODKEY, KEY, view, uint(1 << TAG)},
		Key{MODKEY | ControlMask, KEY, toggleview, uint(1 << TAG)},
		Key{MODKEY | ShiftMask, KEY, tag, uint(1 << TAG)},
		Key{MODKEY | ControlMask | ShiftMask, KEY, toggletag, uint(1 << TAG)},
	}
}

/* helper for spawning shell commands in the pre dwm-5.0 fashion */
func SHCMD(cmd exec.Cmd) interface{} {
	return []interface{}{"/bin/sh", "-c", cmd, nil}
}

var (
	/* tagging */
	tags = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	rules = []Rule{
		/* class      instance    title       tags mask     isfloating   monitor */
		Rule{"Gimp", "", "", 0, true, -1},
		Rule{"Firefox", "", "", 1 << 8, false, -1},
	}

	layouts = []Layout{
		/* symbol     arrange function */
		Layout{"[]=", tile}, /* first entry is default */
		Layout{"><>", nil},  /* no layout function means floating behavior */
		Layout{"[M]", monocle},
	}

	/* commands */
	dmenucmd = []string{"dmenu_run", "-fn", font, "-nb", normbgcolor, "-nf", normfgcolor, "-sb", selbgcolor, "-sf", selfgcolor, ""}
	termcmd  = []string{"uxterm", ""}

	/* key definitions */
	MODKEY uint

	key = setupKeys()

	/* button definitions */
	/* click can be ClkLtSymbol, ClkStatusText, ClkWinTitle, ClkClientWin, or ClkRootWin */
	buttons = []Button{
		/* click          event mask      button          function        argument */
		Button{ClkLtSymbol, 0, Button1, setlayout, 0},
		Button{ClkLtSymbol, 0, Button3, setlayout, &layouts[2]},
		Button{ClkWinTitle, 0, Button2, zoom, 0},
		Button{ClkStatusText, 0, Button2, spawn, termcmd},
		Button{ClkClientWin, MODKEY, Button1, movemouse, 0},
		Button{ClkClientWin, MODKEY, Button2, togglefloating, 0},
		Button{ClkClientWin, MODKEY, Button3, resizemouse, 0},
		Button{ClkTagBar, 0, Button1, view, 0},
		Button{ClkTagBar, 0, Button3, toggleview, 0},
		Button{ClkTagBar, MODKEY, Button1, tag, 0},
		Button{ClkTagBar, MODKEY, Button3, toggletag, 0},
	}
)

func setupKeys() []Key {
	keys := []Key{
		/* modifier                     key        function        argument */
		Key{MODKEY, XK_p, spawn, dmenucmd},
		Key{MODKEY | ShiftMask, XK_Return, spawn, termcmd},
		Key{MODKEY, XK_b, togglebar, 0},
		Key{MODKEY, XK_j, focusstack, 1},
		Key{MODKEY, XK_k, focusstack, -1},
		Key{MODKEY, XK_i, incnmaster, 1},
		Key{MODKEY, XK_d, incnmaster, -1},
		Key{MODKEY, XK_h, setmfact, -0.05},
		Key{MODKEY, XK_l, setmfact, 0.05},
		Key{MODKEY, XK_Return, zoom, 0},
		Key{MODKEY, XK_Tab, view, 0},
		Key{MODKEY | ShiftMask, XK_c, killclient, 0},
		Key{MODKEY, XK_t, setlayout, &layouts[0]},
		Key{MODKEY, XK_f, setlayout, &layouts[1]},
		Key{MODKEY, XK_m, setlayout, &layouts[2]},
		Key{MODKEY, XK_space, setlayout, 0},
		Key{MODKEY | ShiftMask, XK_space, togglefloating, 0},
		Key{MODKEY, XK_0, view, -0},
		Key{MODKEY | ShiftMask, XK_0, tag, -0},
		Key{MODKEY, XK_comma, focusmon, -1},
		Key{MODKEY, XK_period, focusmon, 1},
		Key{MODKEY | ShiftMask, XK_comma, tagmon, -1},
		Key{MODKEY | ShiftMask, XK_period, tagmon, 1},
		Key{MODKEY | ShiftMask, XK_q, quit, 0},
	}

	keys = append(keys, TAGKEYS(XK_1, 0)...)
	keys = append(keys, TAGKEYS(XK_2, 1)...)
	keys = append(keys, TAGKEYS(XK_3, 2)...)
	keys = append(keys, TAGKEYS(XK_4, 3)...)
	keys = append(keys, TAGKEYS(XK_5, 4)...)
	keys = append(keys, TAGKEYS(XK_6, 5)...)
	keys = append(keys, TAGKEYS(XK_7, 6)...)
	keys = append(keys, TAGKEYS(XK_8, 7)...)
	keys = append(keys, TAGKEYS(XK_9, 8)...)

	return keys
}
