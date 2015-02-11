package termstyle

import (
	"fmt"
	"log"
	"os/exec"
)

type Color int

func init() {
	if len(FG) > 0 {
		return
	}

	FG = make([]string, 256)
	BG = make([]string, 256)

	tput("bold", &B)
	tput("smul", &U)
	tput("sgr0", &C)

	for i := 0; i < 256; i++ {
		b, err := exec.Command("tput", "setaf", fmt.Sprint(i)).CombinedOutput()
		if err != nil {
			log.Println(err.Error())
			return
		}
		FG[i] = string(b)
		b, err = exec.Command("tput", "setab", fmt.Sprint(i)).CombinedOutput()
		if err != nil {
			log.Println(err.Error())
			return
		}
		BG[i] = string(b)
	}
}

func tput(name string, load *string) {
	b, err := exec.Command("tput", name).CombinedOutput()
	if err != nil {
		log.Println(err.Error())
		return
	}
	*load = string(b)
}

const (
	System0 Color = iota
	System1
	System2
	System3
	System4
	System5
	System6
	System7
	System8
	System9
	System10
	System11
	System12
	System13
	System14
	System15
	Grey0
	NavyBlue
	DarkBlue
	Blue3
	Blue31
	Blue1
	DarkGreen
	DeepSkyBlue4
	DeepSkyBlue41
	DeepSkyBlue42
	DodgerBlue3
	DodgerBlue2
	Green4
	SpringGreen4
	Turquoise4
	DeepSkyBlue3
	DeepSkyBlue31
	DodgerBlue1
	Green3
	SpringGreen3
	DarkCyan
	LightSeaGreen
	DeepSkyBlue2
	DeepSkyBlue1
	Green31
	SpringGreen31
	SpringGreen2
	Cyan3
	DarkTurquoise
	Turquoise2
	Green1
	SpringGreen21
	SpringGreen1
	MediumSpringGreen
	Cyan2
	Cyan1
	DarkRed
	DeepPink4
	Purple4
	Purple41
	Purple3
	BlueViolet
	Orange4
	Grey37
	MediumPurple4
	SlateBlue3
	SlateBlue31
	RoyalBlue1
	Chartreuse4
	DarkSeaGreen4
	PaleTurquoise4
	SteelBlue
	SteelBlue3
	CornflowerBlue
	Chartreuse3
	DarkSeaGreen41
	CadetBlue
	CadetBlue1
	SkyBlue3
	SteelBlue1
	Chartreuse31
	PaleGreen3
	SeaGreen3
	Aquamarine3
	MediumTurquoise
	SteelBlue11
	Chartreuse2
	SeaGreen2
	SeaGreen1
	SeaGreen11
	Aquamarine1
	DarkSlateGray2
	DarkRed1
	DeepPink41
	DarkMagenta
	DarkMagenta1
	DarkViolet
	Purple
	Orange41
	LightPink4
	Plum4
	MediumPurple3
	MediumPurple31
	SlateBlue1
	Yellow4
	Wheat4
	Grey53
	LightSlateGrey
	MediumPurple
	LightSlateBlue
	Yellow41
	DarkOliveGreen3
	DarkSeaGreen
	LightSkyBlue3
	LightSkyBlue31
	SkyBlue2
	Chartreuse21
	DarkOliveGreen31
	PaleGreen31
	DarkSeaGreen3
	DarkSlateGray3
	SkyBlue1
	Chartreuse1
	LightGreen
	LightGreen1
	PaleGreen1
	Aquamarine11
	DarkSlateGray1
	Red3
	DeepPink411
	MediumVioletRed
	Magenta3
	DarkViolet1
	Purple1
	DarkOrange3
	IndianRed
	HotPink3
	MediumOrchid3
	MediumOrchid
	MediumPurple2
	DarkGoldenrod
	LightSalmon3
	RosyBrown
	Grey63
	MediumPurple21
	MediumPurple1
	Gold3
	DarkKhaki
	NavajoWhite3
	Grey69
	LightSteelBlue3
	LightSteelBlue
	Yellow3
	DarkOliveGreen311
	DarkSeaGreen31
	DarkSeaGreen2
	LightCyan3
	LightSkyBlue1
	GreenYellow
	DarkOliveGreen2
	PaleGreen11
	DarkSeaGreen21
	DarkSeaGreen1
	PaleTurquoise1
	Red31
	DeepPink3
	DeepPink31
	Magenta31
	Magenta311
	Magenta2
	DarkOrange31
	IndianRed1
	HotPink31
	HotPink2
	Orchid
	MediumOrchid1
	Orange3
	LightSalmon31
	LightPink3
	Pink3
	Plum3
	Violet
	Gold31
	LightGoldenrod3
	Tan
	MistyRose3
	Thistle3
	Plum2
	Yellow31
	Khaki3
	LightGoldenrod2
	LightYellow3
	Grey84
	LightSteelBlue1
	Yellow2
	DarkOliveGreen1
	DarkOliveGreen11
	DarkSeaGreen11
	Honeydew2
	LightCyan1
	Red1
	DeepPink2
	DeepPink1
	DeepPink11
	Magenta21
	Magenta1
	OrangeRed1
	IndianRed11
	IndianRed111
	HotPink
	HotPink1
	MediumOrchid11
	DarkOrange
	Salmon1
	LightCoral
	PaleVioletRed1
	Orchid2
	Orchid1
	Orange1
	SandyBrown
	LightSalmon1
	LightPink1
	Pink1
	Plum1
	Gold1
	LightGoldenrod21
	LightGoldenrod211
	NavajoWhite1
	MistyRose1
	Thistle1
	Yellow1
	LightGoldenrod1
	Khaki1
	Wheat1
	Cornsilk1
	Grey100
	Grey3
	Grey7
	Grey11
	Grey15
	Grey19
	Grey23
	Grey27
	Grey30
	Grey35
	Grey39
	Grey42
	Grey46
	Grey50
	Grey54
	Grey58
	Grey62
	Grey66
	Grey70
	Grey74
	Grey78
	Grey82
	Grey85
	Grey89
	Grey93
)

var Names = [...]string{
	System0:           "System0",
	System1:           "System1",
	System2:           "System2",
	System3:           "System3",
	System4:           "System4",
	System5:           "System5",
	System6:           "System6",
	System7:           "System7",
	System8:           "System8",
	System9:           "System9",
	System10:          "System10",
	System11:          "System11",
	System12:          "System12",
	System13:          "System13",
	System14:          "System14",
	System15:          "System15",
	Grey0:             "Grey0",
	NavyBlue:          "NavyBlue",
	DarkBlue:          "DarkBlue",
	Blue3:             "Blue3",
	Blue31:            "Blue31",
	Blue1:             "Blue1",
	DarkGreen:         "DarkGreen",
	DeepSkyBlue4:      "DeepSkyBlue4",
	DeepSkyBlue41:     "DeepSkyBlue41",
	DeepSkyBlue42:     "DeepSkyBlue42",
	DodgerBlue3:       "DodgerBlue3",
	DodgerBlue2:       "DodgerBlue2",
	Green4:            "Green4",
	SpringGreen4:      "SpringGreen4",
	Turquoise4:        "Turquoise4",
	DeepSkyBlue3:      "DeepSkyBlue3",
	DeepSkyBlue31:     "DeepSkyBlue31",
	DodgerBlue1:       "DodgerBlue1",
	Green3:            "Green3",
	SpringGreen3:      "SpringGreen3",
	DarkCyan:          "DarkCyan",
	LightSeaGreen:     "LightSeaGreen",
	DeepSkyBlue2:      "DeepSkyBlue2",
	DeepSkyBlue1:      "DeepSkyBlue1",
	Green31:           "Green31",
	SpringGreen31:     "SpringGreen31",
	SpringGreen2:      "SpringGreen2",
	Cyan3:             "Cyan3",
	DarkTurquoise:     "DarkTurquoise",
	Turquoise2:        "Turquoise2",
	Green1:            "Green1",
	SpringGreen21:     "SpringGreen21",
	SpringGreen1:      "SpringGreen1",
	MediumSpringGreen: "MediumSpringGreen",
	Cyan2:             "Cyan2",
	Cyan1:             "Cyan1",
	DarkRed:           "DarkRed",
	DeepPink4:         "DeepPink4",
	Purple4:           "Purple4",
	Purple41:          "Purple41",
	Purple3:           "Purple3",
	BlueViolet:        "BlueViolet",
	Orange4:           "Orange4",
	Grey37:            "Grey37",
	MediumPurple4:     "MediumPurple4",
	SlateBlue3:        "SlateBlue3",
	SlateBlue31:       "SlateBlue31",
	RoyalBlue1:        "RoyalBlue1",
	Chartreuse4:       "Chartreuse4",
	DarkSeaGreen4:     "DarkSeaGreen4",
	PaleTurquoise4:    "PaleTurquoise4",
	SteelBlue:         "SteelBlue",
	SteelBlue3:        "SteelBlue3",
	CornflowerBlue:    "CornflowerBlue",
	Chartreuse3:       "Chartreuse3",
	DarkSeaGreen41:    "DarkSeaGreen41",
	CadetBlue:         "CadetBlue",
	CadetBlue1:        "CadetBlue1",
	SkyBlue3:          "SkyBlue3",
	SteelBlue1:        "SteelBlue1",
	Chartreuse31:      "Chartreuse31",
	PaleGreen3:        "PaleGreen3",
	SeaGreen3:         "SeaGreen3",
	Aquamarine3:       "Aquamarine3",
	MediumTurquoise:   "MediumTurquoise",
	SteelBlue11:       "SteelBlue11",
	Chartreuse2:       "Chartreuse2",
	SeaGreen2:         "SeaGreen2",
	SeaGreen1:         "SeaGreen1",
	SeaGreen11:        "SeaGreen11",
	Aquamarine1:       "Aquamarine1",
	DarkSlateGray2:    "DarkSlateGray2",
	DarkRed1:          "DarkRed1",
	DeepPink41:        "DeepPink41",
	DarkMagenta:       "DarkMagenta",
	DarkMagenta1:      "DarkMagenta1",
	DarkViolet:        "DarkViolet",
	Purple:            "Purple",
	Orange41:          "Orange41",
	LightPink4:        "LightPink4",
	Plum4:             "Plum4",
	MediumPurple3:     "MediumPurple3",
	MediumPurple31:    "MediumPurple31",
	SlateBlue1:        "SlateBlue1",
	Yellow4:           "Yellow4",
	Wheat4:            "Wheat4",
	Grey53:            "Grey53",
	LightSlateGrey:    "LightSlateGrey",
	MediumPurple:      "MediumPurple",
	LightSlateBlue:    "LightSlateBlue",
	Yellow41:          "Yellow41",
	DarkOliveGreen3:   "DarkOliveGreen3",
	DarkSeaGreen:      "DarkSeaGreen",
	LightSkyBlue3:     "LightSkyBlue3",
	LightSkyBlue31:    "LightSkyBlue31",
	SkyBlue2:          "SkyBlue2",
	Chartreuse21:      "Chartreuse21",
	DarkOliveGreen31:  "DarkOliveGreen31",
	PaleGreen31:       "PaleGreen31",
	DarkSeaGreen3:     "DarkSeaGreen3",
	DarkSlateGray3:    "DarkSlateGray3",
	SkyBlue1:          "SkyBlue1",
	Chartreuse1:       "Chartreuse1",
	LightGreen:        "LightGreen",
	LightGreen1:       "LightGreen1",
	PaleGreen1:        "PaleGreen1",
	Aquamarine11:      "Aquamarine11",
	DarkSlateGray1:    "DarkSlateGray1",
	Red3:              "Red3",
	DeepPink411:       "DeepPink411",
	MediumVioletRed:   "MediumVioletRed",
	Magenta3:          "Magenta3",
	DarkViolet1:       "DarkViolet1",
	Purple1:           "Purple1",
	DarkOrange3:       "DarkOrange3",
	IndianRed:         "IndianRed",
	HotPink3:          "HotPink3",
	MediumOrchid3:     "MediumOrchid3",
	MediumOrchid:      "MediumOrchid",
	MediumPurple2:     "MediumPurple2",
	DarkGoldenrod:     "DarkGoldenrod",
	LightSalmon3:      "LightSalmon3",
	RosyBrown:         "RosyBrown",
	Grey63:            "Grey63",
	MediumPurple21:    "MediumPurple21",
	MediumPurple1:     "MediumPurple1",
	Gold3:             "Gold3",
	DarkKhaki:         "DarkKhaki",
	NavajoWhite3:      "NavajoWhite3",
	Grey69:            "Grey69",
	LightSteelBlue3:   "LightSteelBlue3",
	LightSteelBlue:    "LightSteelBlue",
	Yellow3:           "Yellow3",
	DarkOliveGreen311: "DarkOliveGreen311",
	DarkSeaGreen31:    "DarkSeaGreen31",
	DarkSeaGreen2:     "DarkSeaGreen2",
	LightCyan3:        "LightCyan3",
	LightSkyBlue1:     "LightSkyBlue1",
	GreenYellow:       "GreenYellow",
	DarkOliveGreen2:   "DarkOliveGreen2",
	PaleGreen11:       "PaleGreen11",
	DarkSeaGreen21:    "DarkSeaGreen21",
	DarkSeaGreen1:     "DarkSeaGreen1",
	PaleTurquoise1:    "PaleTurquoise1",
	Red31:             "Red31",
	DeepPink3:         "DeepPink3",
	DeepPink31:        "DeepPink31",
	Magenta31:         "Magenta31",
	Magenta311:        "Magenta311",
	Magenta2:          "Magenta2",
	DarkOrange31:      "DarkOrange31",
	IndianRed1:        "IndianRed1",
	HotPink31:         "HotPink31",
	HotPink2:          "HotPink2",
	Orchid:            "Orchid",
	MediumOrchid1:     "MediumOrchid1",
	Orange3:           "Orange3",
	LightSalmon31:     "LightSalmon31",
	LightPink3:        "LightPink3",
	Pink3:             "Pink3",
	Plum3:             "Plum3",
	Violet:            "Violet",
	Gold31:            "Gold31",
	LightGoldenrod3:   "LightGoldenrod3",
	Tan:               "Tan",
	MistyRose3:        "MistyRose3",
	Thistle3:          "Thistle3",
	Plum2:             "Plum2",
	Yellow31:          "Yellow31",
	Khaki3:            "Khaki3",
	LightGoldenrod2:   "LightGoldenrod2",
	LightYellow3:      "LightYellow3",
	Grey84:            "Grey84",
	LightSteelBlue1:   "LightSteelBlue1",
	Yellow2:           "Yellow2",
	DarkOliveGreen1:   "DarkOliveGreen1",
	DarkOliveGreen11:  "DarkOliveGreen11",
	DarkSeaGreen11:    "DarkSeaGreen11",
	Honeydew2:         "Honeydew2",
	LightCyan1:        "LightCyan1",
	Red1:              "Red1",
	DeepPink2:         "DeepPink2",
	DeepPink1:         "DeepPink1",
	DeepPink11:        "DeepPink11",
	Magenta21:         "Magenta21",
	Magenta1:          "Magenta1",
	OrangeRed1:        "OrangeRed1",
	IndianRed11:       "IndianRed11",
	IndianRed111:      "IndianRed111",
	HotPink:           "HotPink",
	HotPink1:          "HotPink1",
	MediumOrchid11:    "MediumOrchid11",
	DarkOrange:        "DarkOrange",
	Salmon1:           "Salmon1",
	LightCoral:        "LightCoral",
	PaleVioletRed1:    "PaleVioletRed1",
	Orchid2:           "Orchid2",
	Orchid1:           "Orchid1",
	Orange1:           "Orange1",
	SandyBrown:        "SandyBrown",
	LightSalmon1:      "LightSalmon1",
	LightPink1:        "LightPink1",
	Pink1:             "Pink1",
	Plum1:             "Plum1",
	Gold1:             "Gold1",
	LightGoldenrod21:  "LightGoldenrod21",
	LightGoldenrod211: "LightGoldenrod211",
	NavajoWhite1:      "NavajoWhite1",
	MistyRose1:        "MistyRose1",
	Thistle1:          "Thistle1",
	Yellow1:           "Yellow1",
	LightGoldenrod1:   "LightGoldenrod1",
	Khaki1:            "Khaki1",
	Wheat1:            "Wheat1",
	Cornsilk1:         "Cornsilk1",
	Grey100:           "Grey100",
	Grey3:             "Grey3",
	Grey7:             "Grey7",
	Grey11:            "Grey11",
	Grey15:            "Grey15",
	Grey19:            "Grey19",
	Grey23:            "Grey23",
	Grey27:            "Grey27",
	Grey30:            "Grey30",
	Grey35:            "Grey35",
	Grey39:            "Grey39",
	Grey42:            "Grey42",
	Grey46:            "Grey46",
	Grey50:            "Grey50",
	Grey54:            "Grey54",
	Grey58:            "Grey58",
	Grey62:            "Grey62",
	Grey66:            "Grey66",
	Grey70:            "Grey70",
	Grey74:            "Grey74",
	Grey78:            "Grey78",
	Grey82:            "Grey82",
	Grey85:            "Grey85",
	Grey89:            "Grey89",
	Grey93:            "Grey93",
}
