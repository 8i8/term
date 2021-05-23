package term

// Colour represents one of the xterm's 256 colours.
type Colour uint16

// All of the colours.
//go:generate stringer -type=Colour
const (
	Black             Colour = iota // #000000 rgb(0,0,0)              hsl(0,0%,0%)
	Maroon                          // #800000 rgb(128,0,0)            hsl(0,100%,25%)
	Green                           // #008000 rgb(0,128,0)            hsl(120,100%,25%)
	Olive                           // #808000 rgb(128,128,0)          hsl(60,100%,25%)
	Navy                            // #000080 rgb(0,0,128)            hsl(240,100%,25%)
	Purple                          // #800080 rgb(128,0,128)          hsl(300,100%,25%)
	Teal                            // #008080 rgb(0,128,128)          hsl(180,100%,25%)
	Silver                          // #c1c0c0 rgb(192,192,192)        hsl(0,0%,75%)
	Grey                            // #808080 rgb(128,128,128)        hsl(0,0%,50%)
	Red                             // #ff0001 rgb(255,0,0)            hsl(0,100%,50%)
	Lime                            // #00ff00 rgb(0,255,0)            hsl(120,100%,50%)
	Yellow                          // #ffff00 rgb(255,255,0)          hsl(60,100%,50%)
	Blue                            // #0000ff rgb(0,0,255)            hsl(240,100%,50%)
	Fuchsia                         // #ff00ff rgb(255,0,255)          hsl(300,100%,50%)
	Aqua                            // #00ffff rgb(0,255,255)          hsl(180,100%,50%)
	White                           // #ffffff rgb(255,255,255)        hsl(0,0%,100%)
	Grey0                           // #000000 rgb(0,0,0)              hsl(0,0%,0%)
	NavyBlue                        // #00005f rgb(0,0,95)             hsl(240,100%,18%)
	DarkBlue                        // #000087 rgb(0,0,135)            hsl(240,100%,26%)
	Blue3                           // #0000af rgb(0,0,175)            hsl(240,100%,34%)
	Blue0                           // #0000d7 rgb(0,0,215)            hsl(240,100%,42%)
	Blue1                           // #0000ff rgb(0,0,255)            hsl(240,100%,50%)
	DarkGreen                       // #005f00 rgb(0,95,0)             hsl(120,100%,18%)
	DeepSkyBlue7                    // #005f5f rgb(0,95,95)            hsl(180,100%,18%)
	DeepSkyBlue6                    // #005f87 rgb(0,95,135)           hsl(97,100%,26%)
	DeepSkyBlue5                    // #005faf rgb(0,95,175)           hsl(07,100%,34%)
	DodgerBlue3                     // #005fd7 rgb(0,95,215)           hsl(13,100%,42%)
	DodgerBlue2                     // #005fff rgb(0,95,255)           hsl(17,100%,50%)
	Green4                          // #008700 rgb(0,135,0)            hsl(120,100%,26%)
	SpringGreen6                    // #00875f rgb(0,135,95)           hsl(62,100%,26%)
	Turquoise4                      // #008787 rgb(0,135,135)          hsl(180,100%,26%)
	DeepSkyBlue4                    // #0087af rgb(0,135,175)          hsl(93,100%,34%)
	DeepSkyBlue3                    // #0087d7 rgb(0,135,215)          hsl(02,100%,42%)
	DodgerBlue1                     // #0087ff rgb(0,135,255)          hsl(08,100%,50%)
	Green3                          // #00af00 rgb(0,175,0)            hsl(120,100%,34%)
	SpringGreen5                    // #00af5f rgb(0,175,95)           hsl(52,100%,34%)
	DarkCyan                        // #00af87 rgb(0,175,135)          hsl(66,100%,34%)
	LightSeaGreen                   // #00afaf rgb(0,175,175)          hsl(180,100%,34%)
	DeepSkyBlue2                    // #00afd7 rgb(0,175,215)          hsl(91,100%,42%)
	DeepSkyBlue1                    // #00afff rgb(0,175,255)          hsl(98,100%,50%)
	Green2                          // #00d700 rgb(0,215,0)            hsl(120,100%,42%)
	SpringGreen4                    // #00d75f rgb(0,215,95)           hsl(46,100%,42%)
	SpringGreen3                    // #00d787 rgb(0,215,135)          hsl(57,100%,42%)
	Cyan3                           // #00d7af rgb(0,215,175)          hsl(68,100%,42%)
	DarkTurquoise                   // #00d7d7 rgb(0,215,215)          hsl(180,100%,42%)
	Turquoise2                      // #00d7ff rgb(0,215,255)          hsl(89,100%,50%)
	Green1                          // #00ff00 rgb(0,255,0)            hsl(120,100%,50%)
	SpringGreen2                    // #00ff5f rgb(0,255,95)           hsl(42,100%,50%)
	SpringGreen1                    // #00ff87 rgb(0,255,135)          hsl(51,100%,50%)
	MediumSpringGreen               // #00ffaf rgb(0,255,175)          hsl(61,100%,50%)
	Cyan2                           // #00ffd7 rgb(0,255,215)          hsl(70,100%,50%)
	Cyan1                           // #00ffff rgb(0,255,255)          hsl(180,100%,50%)
	DarkRed1                        // #5f0000 rgb(95,0,0)             hsl(0,100%,18%)
	DeepPink8                       // #5f005f rgb(95,0,95)            hsl(300,100%,18%)
	Purple5                         // #5f0087 rgb(95,0,135)           hsl(82,100%,26%)
	Purple4                         // #5f00af rgb(95,0,175)           hsl(72,100%,34%)
	Purple3                         // #5f00d7 rgb(95,0,215)           hsl(66,100%,42%)
	BlueViolet                      // #5f00ff rgb(95,0,255)           hsl(62,100%,50%)
	Orange4                         // #5f5f00 rgb(95,95,0)            hsl(60,100%,18%)
	Grey37                          // #5f5f5f rgb(95,95,95)           hsl(0,0%,37%)
	MediumPurple7                   // #5f5f87 rgb(95,95,135)          hsl(240,17%,45%)
	SlateBlue3                      // #5f5faf rgb(95,95,175)          hsl(240,33%,52%)
	SlateBlue2                      // #5f5fd7 rgb(95,95,215)          hsl(240,60%,60%)
	RoyalBlue1                      // #5f5fff rgb(95,95,255)          hsl(240,100%,68%)
	Chartreuse6                     // #5f8700 rgb(95,135,0)           hsl(7,100%,26%)
	DarkSeaGreen9                   // #5f875f rgb(95,135,95)          hsl(120,17%,45%)
	PaleTurquoise4                  // #5f8787 rgb(95,135,135)         hsl(180,17%,45%)
	SteelBlue4                      // #5f87af rgb(95,135,175)         hsl(210,33%,52%)
	SteelBlue3                      // #5f87d7 rgb(95,135,215)         hsl(220,60%,60%)
	CornflowerBlue                  // #5f87ff rgb(95,135,255)         hsl(225,100%,68%)
	Chartreuse5                     // #5faf00 rgb(95,175,0)           hsl(7,100%,34%)
	DarkSeaGreen8                   // #5faf5f rgb(95,175,95)          hsl(120,33%,52%)
	CadetBlue1                      // #5faf87 rgb(95,175,135)         hsl(150,33%,52%)
	CadetBlue2                      // #5fafaf rgb(95,175,175)         hsl(180,33%,52%)
	SkyBlue3                        // #5fafd7 rgb(95,175,215)         hsl(200,60%,60%)
	SteelBlue2                      // #5fafff rgb(95,175,255)         hsl(210,100%,68%)
	Chartreuse4                     // #5fd700 rgb(95,215,0)           hsl(3,100%,42%)
	PaleGreen4                      // #5fd75f rgb(95,215,95)          hsl(120,60%,60%)
	SeaGreen4                       // #5fd787 rgb(95,215,135)         hsl(140,60%,60%)
	Aquamarine3                     // #5fd7af rgb(95,215,175)         hsl(160,60%,60%)
	MediumTurquoise                 // #5fd7d7 rgb(95,215,215)         hsl(180,60%,60%)
	SteelBlue1                      // #5fd7ff rgb(95,215,255)         hsl(195,100%,68%)
	Chartreuse3                     // #5fff00 rgb(95,255,0)           hsl(7,100%,50%)
	SeaGreen3                       // #5fff5f rgb(95,255,95)          hsl(120,100%,68%)
	SeaGreen2                       // #5fff87 rgb(95,255,135)         hsl(135,100%,68%)
	SeaGreen1                       // #5fffaf rgb(95,255,175)         hsl(150,100%,68%)
	Aquamarine2                     // #5fffd7 rgb(95,255,215)         hsl(165,100%,68%)
	DarkSlateGray2                  // #5fffff rgb(95,255,255)         hsl(180,100%,68%)
	DarkRed2                        // #870000 rgb(135,0,0)            hsl(0,100%,26%)
	DeepPink7                       // #87005f rgb(135,0,95)           hsl(17,100%,26%)
	DarkMagenta1                    // #870087 rgb(135,0,135)          hsl(300,100%,26%)
	DarkMagenta2                    // #8700af rgb(135,0,175)          hsl(86,100%,34%)
	DarkViolet1                     // #8700d7 rgb(135,0,215)          hsl(77,100%,42%)
	Purple2                         // #8700ff rgb(135,0,255)          hsl(71,100%,50%)
	Orange3                         // #875f00 rgb(135,95,0)           hsl(2,100%,26%)
	LightPink4                      // #875f5f rgb(135,95,95)          hsl(0,17%,45%)
	Plum4                           // #875f87 rgb(135,95,135)         hsl(300,17%,45%)
	MediumPurple6                   // #875faf rgb(135,95,175)         hsl(270,33%,52%)
	MediumPurple5                   // #875fd7 rgb(135,95,215)         hsl(260,60%,60%)
	SlateBlue1                      // #875fff rgb(135,95,255)         hsl(255,100%,68%)
	Yellow6                         // #878700 rgb(135,135,0)          hsl(60,100%,26%)
	Wheat4                          // #87875f rgb(135,135,95)         hsl(60,17%,45%)
	Grey53                          // #878787 rgb(135,135,135)        hsl(0,0%,52%)
	LightSlateGrey                  // #8787af rgb(135,135,175)        hsl(240,20%,60%)
	MediumPurple4                   // #8787d7 rgb(135,135,215)        hsl(240,50%,68%)
	LightSlateBlue                  // #8787ff rgb(135,135,255)        hsl(240,100%,76%)
	Yellow5                         // #87af00 rgb(135,175,0)          hsl(3,100%,34%)
	DarkOliveGreen6                 // #87af5f rgb(135,175,95)         hsl(90,33%,52%)
	DarkSeaGreen7                   // #87af87 rgb(135,175,135)        hsl(120,20%,60%)
	LightSkyBlue3                   // #87afaf rgb(135,175,175)        hsl(180,20%,60%)
	LightSkyBlue2                   // #87afd7 rgb(135,175,215)        hsl(210,50%,68%)
	SkyBlue2                        // #87afff rgb(135,175,255)        hsl(220,100%,76%)
	Chartreuse2                     // #87d700 rgb(135,215,0)          hsl(2,100%,42%)
	DarkOliveGreen5                 // #87d75f rgb(135,215,95)         hsl(100,60%,60%)
	PaleGreen3                      // #87d787 rgb(135,215,135)        hsl(120,50%,68%)
	DarkSeaGreen6                   // #87d7af rgb(135,215,175)        hsl(150,50%,68%)
	DarkSlateGray3                  // #87d7d7 rgb(135,215,215)        hsl(180,50%,68%)
	SkyBlue1                        // #87d7ff rgb(135,215,255)        hsl(200,100%,76%)
	Chartreuse1                     // #87ff00 rgb(135,255,0)          hsl(8,100%,50%)
	LightGreen1                     // #87ff5f rgb(135,255,95)         hsl(105,100%,68%)
	LightGreen2                     // #87ff87 rgb(135,255,135)        hsl(120,100%,76%)
	PaleGreen2                      // #87ffaf rgb(135,255,175)        hsl(140,100%,76%)
	Aquamarine1                     // #87ffd7 rgb(135,255,215)        hsl(160,100%,76%)
	DarkSlateGray1                  // #87ffff rgb(135,255,255)        hsl(180,100%,76%)
	Red3                            // #af0000 rgb(175,0,0)            hsl(0,100%,34%)
	DeepPink6                       // #af005f rgb(175,0,95)           hsl(27,100%,34%)
	MediumVioletRed                 // #af0087 rgb(175,0,135)          hsl(13,100%,34%)
	Magenta6                        // #af00af rgb(175,0,175)          hsl(300,100%,34%)
	DarkViolet2                     // #af00d7 rgb(175,0,215)          hsl(88,100%,42%)
	Purple1                         // #af00ff rgb(175,0,255)          hsl(81,100%,50%)
	DarkOrange3                     // #af5f00 rgb(175,95,0)           hsl(2,100%,34%)
	IndianRed4                      // #af5f5f rgb(175,95,95)          hsl(0,33%,52%)
	HotPink5                        // #af5f87 rgb(175,95,135)         hsl(330,33%,52%)
	MediumOrchid4                   // #af5faf rgb(175,95,175)         hsl(300,33%,52%)
	MediumOrchid3                   // #af5fd7 rgb(175,95,215)         hsl(280,60%,60%)
	MediumPurple3                   // #af5fff rgb(175,95,255)         hsl(270,100%,68%)
	DarkGoldenrod                   // #af8700 rgb(175,135,0)          hsl(6,100%,34%)
	LightSalmon3                    // #af875f rgb(175,135,95)         hsl(30,33%,52%)
	RosyBrown                       // #af8787 rgb(175,135,135)        hsl(0,20%,60%)
	Grey63                          // #af87af rgb(175,135,175)        hsl(300,20%,60%)
	MediumPurple2                   // #af87d7 rgb(175,135,215)        hsl(270,50%,68%)
	MediumPurple1                   // #af87ff rgb(175,135,255)        hsl(260,100%,76%)
	Gold3                           // #afaf00 rgb(175,175,0)          hsl(60,100%,34%)
	DarkKhaki                       // #afaf5f rgb(175,175,95)         hsl(60,33%,52%)
	NavajoWhite3                    // #afaf87 rgb(175,175,135)        hsl(60,20%,60%)
	Grey69                          // #afafaf rgb(175,175,175)        hsl(0,0%,68%)
	LightSteelBlue3                 // #afafd7 rgb(175,175,215)        hsl(240,33%,76%)
	LightSteelBlue                  // #afafff rgb(175,175,255)        hsl(240,100%,84%)
	Yellow4                         // #afd700 rgb(175,215,0)          hsl(1,100%,42%)
	DarkOliveGreen4                 // #afd75f rgb(175,215,95)         hsl(80,60%,60%)
	DarkSeaGreen5                   // #afd787 rgb(175,215,135)        hsl(90,50%,68%)
	DarkSeaGreen4                   // #afd7af rgb(175,215,175)        hsl(120,33%,76%)
	LightCyan3                      // #afd7d7 rgb(175,215,215)        hsl(180,33%,76%)
	LightSkyBlue1                   // #afd7ff rgb(175,215,255)        hsl(210,100%,84%)
	GreenYellow                     // #afff00 rgb(175,255,0)          hsl(8,100%,50%)
	DarkOliveGreen3                 // #afff5f rgb(175,255,95)         hsl(90,100%,68%)
	PaleGreen1                      // #afff87 rgb(175,255,135)        hsl(100,100%,76%)
	DarkSeaGreen3                   // #afffaf rgb(175,255,175)        hsl(120,100%,84%)
	DarkSeaGreen2                   // #afffd7 rgb(175,255,215)        hsl(150,100%,84%)
	PaleTurquoise1                  // #afffff rgb(175,255,255)        hsl(180,100%,84%)
	Red2                            // #d70000 rgb(215,0,0)            hsl(0,100%,42%)
	DeepPink5                       // #d7005f rgb(215,0,95)           hsl(33,100%,42%)
	DeepPink4                       // #d70087 rgb(215,0,135)          hsl(22,100%,42%)
	Magenta5                        // #d700af rgb(215,0,175)          hsl(11,100%,42%)
	Magenta4                        // #d700d7 rgb(215,0,215)          hsl(300,100%,42%)
	Magenta3                        // #d700ff rgb(215,0,255)          hsl(90,100%,50%)
	DarkOrange2                     // #d75f00 rgb(215,95,0)           hsl(6,100%,42%)
	IndianRed3                      // #d75f5f rgb(215,95,95)          hsl(0,60%,60%)
	HotPink4                        // #d75f87 rgb(215,95,135)         hsl(340,60%,60%)
	HotPink3                        // #d75faf rgb(215,95,175)         hsl(320,60%,60%)
	Orchid                          // #d75fd7 rgb(215,95,215)         hsl(300,60%,60%)
	MediumOrchid2                   // #d75fff rgb(215,95,255)         hsl(285,100%,68%)
	Orange2                         // #d78700 rgb(215,135,0)          hsl(7,100%,42%)
	LightSalmon2                    // #d7875f rgb(215,135,95)         hsl(20,60%,60%)
	LightPink3                      // #d78787 rgb(215,135,135)        hsl(0,50%,68%)
	Pink3                           // #d787af rgb(215,135,175)        hsl(330,50%,68%)
	Plum3                           // #d787d7 rgb(215,135,215)        hsl(300,50%,68%)
	Violet                          // #d787ff rgb(215,135,255)        hsl(280,100%,76%)
	Gold2                           // #d7af00 rgb(215,175,0)          hsl(8,100%,42%)
	LightGoldenrod5                 // #d7af5f rgb(215,175,95)         hsl(40,60%,60%)
	Tan                             // #d7af87 rgb(215,175,135)        hsl(30,50%,68%)
	MistyRose3                      // #d7afaf rgb(215,175,175)        hsl(0,33%,76%)
	Thistle3                        // #d7afd7 rgb(215,175,215)        hsl(300,33%,76%)
	Plum2                           // #d7afff rgb(215,175,255)        hsl(270,100%,84%)
	Yellow3                         // #d7d700 rgb(215,215,0)          hsl(60,100%,42%)
	Khaki3                          // #d7d75f rgb(215,215,95)         hsl(60,60%,60%)
	LightGoldenrod4                 // #d7d787 rgb(215,215,135)        hsl(60,50%,68%)
	LightYellow3                    // #d7d7af rgb(215,215,175)        hsl(60,33%,76%)
	Grey84                          // #d7d7d7 rgb(215,215,215)        hsl(0,0%,84%)
	LightSteelBlue1                 // #d7d7ff rgb(215,215,255)        hsl(240,100%,92%)
	Yellow2                         // #d7ff00 rgb(215,255,0)          hsl(9,100%,50%)
	DarkOliveGreen2                 // #d7ff5f rgb(215,255,95)         hsl(75,100%,68%)
	DarkOliveGreen1                 // #d7ff87 rgb(215,255,135)        hsl(80,100%,76%)
	DarkSeaGreen1                   // #d7ffaf rgb(215,255,175)        hsl(90,100%,84%)
	Honeydew2                       // #d7ffd7 rgb(215,255,215)        hsl(120,100%,92%)
	LightCyan1                      // #d7ffff rgb(215,255,255)        hsl(180,100%,92%)
	Red1                            // #ff0000 rgb(255,0,0)            hsl(0,100%,50%)
	DeepPink3                       // #ff005f rgb(255,0,95)           hsl(37,100%,50%)
	DeepPink2                       // #ff0087 rgb(255,0,135)          hsl(28,100%,50%)
	DeepPink1                       // #ff00af rgb(255,0,175)          hsl(18,100%,50%)
	Magenta2                        // #ff00d7 rgb(255,0,215)          hsl(09,100%,50%)
	Magenta1                        // #ff00ff rgb(255,0,255)          hsl(300,100%,50%)
	OrangeRed1                      // #ff5f00 rgb(255,95,0)           hsl(2,100%,50%)
	IndianRed2                      // #ff5f5f rgb(255,95,95)          hsl(0,100%,68%)
	IndianRed1                      // #ff5f87 rgb(255,95,135)         hsl(345,100%,68%)
	HotPink2                        // #ff5faf rgb(255,95,175)         hsl(330,100%,68%)
	HotPink1                        // #ff5fd7 rgb(255,95,215)         hsl(315,100%,68%)
	MediumOrchid1                   // #ff5fff rgb(255,95,255)         hsl(300,100%,68%)
	DarkOrange1                     // #ff8700 rgb(255,135,0)          hsl(1,100%,50%)
	Salmon1                         // #ff875f rgb(255,135,95)         hsl(15,100%,68%)
	LightCoral                      // #ff8787 rgb(255,135,135)        hsl(0,100%,76%)
	PaleVioletRed1                  // #ff87af rgb(255,135,175)        hsl(340,100%,76%)
	Orchid2                         // #ff87d7 rgb(255,135,215)        hsl(320,100%,76%)
	Orchid1                         // #ff87ff rgb(255,135,255)        hsl(300,100%,76%)
	Orange1                         // #ffaf00 rgb(255,175,0)          hsl(1,100%,50%)
	SandyBrown                      // #ffaf5f rgb(255,175,95)         hsl(30,100%,68%)
	LightSalmon1                    // #ffaf87 rgb(255,175,135)        hsl(20,100%,76%)
	LightPink1                      // #ffafaf rgb(255,175,175)        hsl(0,100%,84%)
	Pink1                           // #ffafd7 rgb(255,175,215)        hsl(330,100%,84%)
	Plum1                           // #ffafff rgb(255,175,255)        hsl(300,100%,84%)
	Gold1                           // #ffd700 rgb(255,215,0)          hsl(0,100%,50%)
	LightGoldenrod3                 // #ffd75f rgb(255,215,95)         hsl(45,100%,68%)
	LightGoldenrod2                 // #ffd787 rgb(255,215,135)        hsl(40,100%,76%)
	NavajoWhite1                    // #ffd7af rgb(255,215,175)        hsl(30,100%,84%)
	MistyRose1                      // #ffd7d7 rgb(255,215,215)        hsl(0,100%,92%)
	Thistle1                        // #ffd7ff rgb(255,215,255)        hsl(300,100%,92%)
	Yellow1                         // #ffff00 rgb(255,255,0)          hsl(60,100%,50%)
	LightGoldenrod1                 // #ffff5f rgb(255,255,95)         hsl(60,100%,68%)
	Khaki1                          // #ffff87 rgb(255,255,135)        hsl(60,100%,76%)
	Wheat1                          // #ffffaf rgb(255,255,175)        hsl(60,100%,84%)
	Cornsilk1                       // #ffffd7 rgb(255,255,215)        hsl(60,100%,92%)
	Grey100                         // #ffffff rgb(255,255,255)        hsl(0,0%,100%)
	Grey3                           // #080808 rgb(8,8,8)              hsl(0,0%,3%)
	Grey7                           // #121212 rgb(18,18,18)           hsl(0,0%,7%)
	Grey11                          // #1c1c1c rgb(28,28,28)           hsl(0,0%,10%)
	Grey15                          // #262626 rgb(38,38,38)           hsl(0,0%,14%)
	Grey19                          // #303030 rgb(48,48,48)           hsl(0,0%,18%)
	Grey23                          // #3a3a3a rgb(58,58,58)           hsl(0,0%,22%)
	Grey27                          // #444444 rgb(68,68,68)           hsl(0,0%,26%)
	Grey30                          // #4e4e4e rgb(78,78,78)           hsl(0,0%,30%)
	Grey35                          // #585858 rgb(88,88,88)           hsl(0,0%,34%)
	Grey39                          // #626262 rgb(98,98,98)           hsl(0,0%,37%)
	Grey42                          // #6c6c6c rgb(108,108,108)        hsl(0,0%,40%)
	Grey46                          // #767676 rgb(118,118,118)        hsl(0,0%,46%)
	Grey50                          // #808080 rgb(128,128,128)        hsl(0,0%,50%)
	Grey54                          // #8a8a8a rgb(138,138,138)        hsl(0,0%,54%)
	Grey58                          // #949494 rgb(148,148,148)        hsl(0,0%,58%)
	Grey62                          // #9e9e9e rgb(158,158,158)        hsl(0,0%,61%)
	Grey66                          // #a8a8a8 rgb(168,168,168)        hsl(0,0%,65%)
	Grey70                          // #b2b2b2 rgb(178,178,178)        hsl(0,0%,69%)
	Grey74                          // #bcbcbc rgb(188,188,188)        hsl(0,0%,73%)
	Grey78                          // #c6c6c6 rgb(198,198,198)        hsl(0,0%,77%)
	Grey82                          // #d0d0d0 rgb(208,208,208)        hsl(0,0%,81%)
	Grey85                          // #dadada rgb(218,218,218)        hsl(0,0%,85%)
	Grey89                          // #e4e4e4 rgb(228,228,228)        hsl(0,0%,89%)
	Grey93                          // #eeeeee rgb(238,238,238)        hsl(0,0%,93%)
)

// Black returns the given string wrapped in the appropriate escape
// codes to set the named colour, and then back to the base colour
// afterwards.
func (c Colour) Black(s string) string {
	return string(term[Black]) + s + string(term[c])
}

// Maroon returns the given string wrapped in the appropriate escape
// codes to set the named colour, and then back to the base colour
// afterwards.
func (c Colour) Maroon(s string) string {
	return string(term[Maroon]) + s + string(term[c])
}

// Green returns the given string wrapped in the appropriate escape
// codes to set the named colour, and then back to the base colour
// afterwards.
func (c Colour) Green(s string) string {
	return string(term[Green]) + s + string(term[c])
}

// Olive
// Navy
// Purple
// Teal
// Silver
// Grey
// Red

// Red returns the given string wrapped in the appropriate escape codes
// to set the named colour, and then back to the base colour afterwards.
func (c Colour) Red(s string) string {
	return string(term[Red]) + s + string(term[c])
}

// Lime
// Yellow
// Blue
// Fuchsia
// Aqua

// White returns the given string wrapped in the appropriate escape
// codes to set the named colour, and then back to the base colour
// afterwards.
func (c Colour) White(s string) string {
	return string(term[White]) + s + string(term[c])
}

// Grey0
// NavyBlue
// DarkBlue
// Blue3
// Blue0
// Blue1
// DarkGreen
// DeepSkyBlue7
// DeepSkyBlue6
// DeepSkyBlue5
// DodgerBlue3
// DodgerBlue2
// Green4
// SpringGreen6
// Turquoise4
// DeepSkyBlue4
// DeepSkyBlue3
// DodgerBlue1
// Green3
// SpringGreen5
// DarkCyan
// LightSeaGreen
// DeepSkyBlue2
// DeepSkyBlue1
// Green2
// SpringGreen4
// SpringGreen3
// Cyan3
// DarkTurquoise
// Turquoise2
// Green1
// SpringGreen2
// SpringGreen1
// MediumSpringGreen
// Cyan2
// Cyan1
// DarkRed1
// DeepPink8
// Purple5
// Purple4
// Purple3
// BlueViolet
// Orange4
// Grey37
// MediumPurple7
// SlateBlue3
// SlateBlue2
// RoyalBlue1
// Chartreuse6
// DarkSeaGreen9
// PaleTurquoise4
// SteelBlue4
// SteelBlue3
// CornflowerBlue
// Chartreuse5
// DarkSeaGreen8
// CadetBlue1
// CadetBlue2
// SkyBlue3
// SteelBlue2
// Chartreuse4
// PaleGreen4
// SeaGreen4
// Aquamarine3
// MediumTurquoise
// SteelBlue1
// Chartreuse3
// SeaGreen3
// SeaGreen2
// SeaGreen1
// Aquamarine2
// DarkSlateGray2
// DarkRed2
// DeepPink7
// DarkMagenta1
// DarkMagenta2
// DarkViolet1
// Purple2
// Orange3
// LightPink4
// Plum4
// MediumPurple6
// MediumPurple5
// SlateBlue1
// Yellow6
// Wheat4
// Grey53
// LightSlateGrey
// MediumPurple4
// LightSlateBlue
// Yellow5
// DarkOliveGreen6
// DarkSeaGreen7
// LightSkyBlue3
// LightSkyBlue2
// SkyBlue2
// Chartreuse2
// DarkOliveGreen5
// PaleGreen3
// DarkSeaGreen6
// DarkSlateGray3
// SkyBlue1
// Chartreuse1
// LightGreen1
// LightGreen2
// PaleGreen2
// Aquamarine1
// DarkSlateGray1
// Red3
// DeepPink6
// MediumVioletRed
// Magenta6
// DarkViolet2
// Purple1
// DarkOrange3
// IndianRed4
// HotPink5
// MediumOrchid4
// MediumOrchid3
// MediumPurple3
// DarkGoldenrod
// LightSalmon3
// RosyBrown
// Grey63
// MediumPurple2
// MediumPurple1
// Gold3
// DarkKhaki
// NavajoWhite3
// Grey69
// LightSteelBlue3
// LightSteelBlue
// Yellow4
// DarkOliveGreen4
// DarkSeaGreen5
// DarkSeaGreen4
// LightCyan3
// LightSkyBlue1
// GreenYellow
// DarkOliveGreen3
// PaleGreen1
// DarkSeaGreen3
// DarkSeaGreen2
// PaleTurquoise1
// Red2
// DeepPink5
// DeepPink4
// Magenta5
// Magenta4
// Magenta3
// DarkOrange2
// IndianRed3
// HotPink4
// HotPink3
// Orchid
// MediumOrchid2
// Orange2
// LightSalmon2
// LightPink3
// Pink3
// Plum3
// Violet
// Gold2
// LightGoldenrod5
// Tan
// MistyRose3
// Thistle3
// Plum2
// Yellow3
// Khaki3
// LightGoldenrod4
// LightYellow3
// Grey84
// LightSteelBlue1
// Yellow2
// DarkOliveGreen2
// DarkOliveGreen1
// DarkSeaGreen1
// Honeydew2
// LightCyan1
// Red1
// DeepPink3
// DeepPink2
// DeepPink1
// Magenta2
// Magenta1
// OrangeRed1
// IndianRed2
// IndianRed1
// HotPink2
// HotPink1
// MediumOrchid1
// DarkOrange1
// Salmon1
// LightCoral
// PaleVioletRed1
// Orchid2
// Orchid1

// Orange1 returns the given string wrapped in the appropriate escape
// codes to set the named colour, and then back to the base colour
// afterwards.
func (c Colour) Orange1(s string) string {
	return string(term[Orange1]) + s + string(term[c])
}

// SandyBrown
// LightSalmon1
// LightPink1
// Pink1
// Plum1
// Gold1
// LightGoldenrod3
// LightGoldenrod2
// NavajoWhite1
// MistyRose1
// Thistle1
// Yellow1
// LightGoldenrod1
// Khaki1
// Wheat1
// Cornsilk1
// Grey100
// Grey3
// Grey7
// Grey11
// Grey15
// Grey19
// Grey23
// Grey27
// Grey30
// Grey35
// Grey39
// Grey42
// Grey46
// Grey50
// Grey54
// Grey58
// Grey62
// Grey66
// Grey70
// Grey74
// Grey78
// Grey82
// Grey85
// Grey89
// Grey93

var hexstring = [256]string{
	"#000000",
	"#800000",
	"#008000",
	"#808000",
	"#000080",
	"#800080",
	"#008080",
	"#c1c0c0",
	"#808080",
	"#ff0001",
	"#00ff00",
	"#ffff00",
	"#0000ff",
	"#ff00ff",
	"#00ffff",
	"#ffffff",
	"#000000",
	"#00005f",
	"#000087",
	"#0000af",
	"#0000d7",
	"#0000ff",
	"#005f00",
	"#005f5f",
	"#005f87",
	"#005faf",
	"#005fd7",
	"#005fff",
	"#008700",
	"#00875f",
	"#008787",
	"#0087af",
	"#0087d7",
	"#0087ff",
	"#00af00",
	"#00af5f",
	"#00af87",
	"#00afaf",
	"#00afd7",
	"#00afff",
	"#00d700",
	"#00d75f",
	"#00d787",
	"#00d7af",
	"#00d7d7",
	"#00d7ff",
	"#00ff00",
	"#00ff5f",
	"#00ff87",
	"#00ffaf",
	"#00ffd7",
	"#00ffff",
	"#5f0000",
	"#5f005f",
	"#5f0087",
	"#5f00af",
	"#5f00d7",
	"#5f00ff",
	"#5f5f00",
	"#5f5f5f",
	"#5f5f87",
	"#5f5faf",
	"#5f5fd7",
	"#5f5fff",
	"#5f8700",
	"#5f875f",
	"#5f8787",
	"#5f87af",
	"#5f87d7",
	"#5f87ff",
	"#5faf00",
	"#5faf5f",
	"#5faf87",
	"#5fafaf",
	"#5fafd7",
	"#5fafff",
	"#5fd700",
	"#5fd75f",
	"#5fd787",
	"#5fd7af",
	"#5fd7d7",
	"#5fd7ff",
	"#5fff00",
	"#5fff5f",
	"#5fff87",
	"#5fffaf",
	"#5fffd7",
	"#5fffff",
	"#870000",
	"#87005f",
	"#870087",
	"#8700af",
	"#8700d7",
	"#8700ff",
	"#875f00",
	"#875f5f",
	"#875f87",
	"#875faf",
	"#875fd7",
	"#875fff",
	"#878700",
	"#87875f",
	"#878787",
	"#8787af",
	"#8787d7",
	"#8787ff",
	"#87af00",
	"#87af5f",
	"#87af87",
	"#87afaf",
	"#87afd7",
	"#87afff",
	"#87d700",
	"#87d75f",
	"#87d787",
	"#87d7af",
	"#87d7d7",
	"#87d7ff",
	"#87ff00",
	"#87ff5f",
	"#87ff87",
	"#87ffaf",
	"#87ffd7",
	"#87ffff",
	"#af0000",
	"#af005f",
	"#af0087",
	"#af00af",
	"#af00d7",
	"#af00ff",
	"#af5f00",
	"#af5f5f",
	"#af5f87",
	"#af5faf",
	"#af5fd7",
	"#af5fff",
	"#af8700",
	"#af875f",
	"#af8787",
	"#af87af",
	"#af87d7",
	"#af87ff",
	"#afaf00",
	"#afaf5f",
	"#afaf87",
	"#afafaf",
	"#afafd7",
	"#afafff",
	"#afd700",
	"#afd75f",
	"#afd787",
	"#afd7af",
	"#afd7d7",
	"#afd7ff",
	"#afff00",
	"#afff5f",
	"#afff87",
	"#afffaf",
	"#afffd7",
	"#afffff",
	"#d70000",
	"#d7005f",
	"#d70087",
	"#d700af",
	"#d700d7",
	"#d700ff",
	"#d75f00",
	"#d75f5f",
	"#d75f87",
	"#d75faf",
	"#d75fd7",
	"#d75fff",
	"#d78700",
	"#d7875f",
	"#d78787",
	"#d787af",
	"#d787d7",
	"#d787ff",
	"#d7af00",
	"#d7af5f",
	"#d7af87",
	"#d7afaf",
	"#d7afd7",
	"#d7afff",
	"#d7d700",
	"#d7d75f",
	"#d7d787",
	"#d7d7af",
	"#d7d7d7",
	"#d7d7ff",
	"#d7ff00",
	"#d7ff5f",
	"#d7ff87",
	"#d7ffaf",
	"#d7ffd7",
	"#d7ffff",
	"#ff0000",
	"#ff005f",
	"#ff0087",
	"#ff00af",
	"#ff00d7",
	"#ff00ff",
	"#ff5f00",
	"#ff5f5f",
	"#ff5f87",
	"#ff5faf",
	"#ff5fd7",
	"#ff5fff",
	"#ff8700",
	"#ff875f",
	"#ff8787",
	"#ff87af",
	"#ff87d7",
	"#ff87ff",
	"#ffaf00",
	"#ffaf5f",
	"#ffaf87",
	"#ffafaf",
	"#ffafd7",
	"#ffafff",
	"#ffd700",
	"#ffd75f",
	"#ffd787",
	"#ffd7af",
	"#ffd7d7",
	"#ffd7ff",
	"#ffff00",
	"#ffff5f",
	"#ffff87",
	"#ffffaf",
	"#ffffd7",
	"#ffffff",
	"#080808",
	"#121212",
	"#1c1c1c",
	"#262626",
	"#303030",
	"#3a3a3a",
	"#444444",
	"#4e4e4e",
	"#585858",
	"#626262",
	"#6c6c6c",
	"#767676",
	"#808080",
	"#8a8a8a",
	"#949494",
	"#9e9e9e",
	"#a8a8a8",
	"#b2b2b2",
	"#bcbcbc",
	"#c6c6c6",
	"#d0d0d0",
	"#dadada",
	"#e4e4e4",
	"#eeeeee",
}

// HexString returns the colour as a string in the hex format.
func (c Colour) HexString() string {
	if Black <= c && c <= Grey93 {
		return hexstring[c]
	}
	var buf [3]byte
	n := fmtInt8(buf, uint8(c))
	return "!%HexString(" + string(buf[n:]) + ")"
}

var rgbstring = [256]string{
	"rgb(0,0,0)",
	"rgb(128,0,0)",
	"rgb(0,128,0)",
	"rgb(128,128,0)",
	"rgb(0,0,128)",
	"rgb(128,0,128)",
	"rgb(0,128,128)",
	"rgb(192,192,192)",
	"rgb(128,128,128)",
	"rgb(255,0,0)",
	"rgb(0,255,0)",
	"rgb(255,255,0)",
	"rgb(0,0,255)",
	"rgb(255,0,255)",
	"rgb(0,255,255)",
	"rgb(255,255,255)",
	"rgb(0,0,0)",
	"rgb(0,0,95)",
	"rgb(0,0,135)",
	"rgb(0,0,175)",
	"rgb(0,0,215)",
	"rgb(0,0,255)",
	"rgb(0,95,0)",
	"rgb(0,95,95)",
	"rgb(0,95,135)",
	"rgb(0,95,175)",
	"rgb(0,95,215)",
	"rgb(0,95,255)",
	"rgb(0,135,0)",
	"rgb(0,135,95)",
	"rgb(0,135,135)",
	"rgb(0,135,175)",
	"rgb(0,135,215)",
	"rgb(0,135,255)",
	"rgb(0,175,0)",
	"rgb(0,175,95)",
	"rgb(0,175,135)",
	"rgb(0,175,175)",
	"rgb(0,175,215)",
	"rgb(0,175,255)",
	"rgb(0,215,0)",
	"rgb(0,215,95)",
	"rgb(0,215,135)",
	"rgb(0,215,175)",
	"rgb(0,215,215)",
	"rgb(0,215,255)",
	"rgb(0,255,0)",
	"rgb(0,255,95)",
	"rgb(0,255,135)",
	"rgb(0,255,175)",
	"rgb(0,255,215)",
	"rgb(0,255,255)",
	"rgb(95,0,0)",
	"rgb(95,0,95)",
	"rgb(95,0,135)",
	"rgb(95,0,175)",
	"rgb(95,0,215)",
	"rgb(95,0,255)",
	"rgb(95,95,0)",
	"rgb(95,95,95)",
	"rgb(95,95,135)",
	"rgb(95,95,175)",
	"rgb(95,95,215)",
	"rgb(95,95,255)",
	"rgb(95,135,0)",
	"rgb(95,135,95)",
	"rgb(95,135,135)",
	"rgb(95,135,175)",
	"rgb(95,135,215)",
	"rgb(95,135,255)",
	"rgb(95,175,0)",
	"rgb(95,175,95)",
	"rgb(95,175,135)",
	"rgb(95,175,175)",
	"rgb(95,175,215)",
	"rgb(95,175,255)",
	"rgb(95,215,0)",
	"rgb(95,215,95)",
	"rgb(95,215,135)",
	"rgb(95,215,175)",
	"rgb(95,215,215)",
	"rgb(95,215,255)",
	"rgb(95,255,0)",
	"rgb(95,255,95)",
	"rgb(95,255,135)",
	"rgb(95,255,175)",
	"rgb(95,255,215)",
	"rgb(95,255,255)",
	"rgb(135,0,0)",
	"rgb(135,0,95)",
	"rgb(135,0,135)",
	"rgb(135,0,175)",
	"rgb(135,0,215)",
	"rgb(135,0,255)",
	"rgb(135,95,0)",
	"rgb(135,95,95)",
	"rgb(135,95,135)",
	"rgb(135,95,175)",
	"rgb(135,95,215)",
	"rgb(135,95,255)",
	"rgb(135,135,0)",
	"rgb(135,135,95)",
	"rgb(135,135,135)",
	"rgb(135,135,175)",
	"rgb(135,135,215)",
	"rgb(135,135,255)",
	"rgb(135,175,0)",
	"rgb(135,175,95)",
	"rgb(135,175,135)",
	"rgb(135,175,175)",
	"rgb(135,175,215)",
	"rgb(135,175,255)",
	"rgb(135,215,0)",
	"rgb(135,215,95)",
	"rgb(135,215,135)",
	"rgb(135,215,175)",
	"rgb(135,215,215)",
	"rgb(135,215,255)",
	"rgb(135,255,0)",
	"rgb(135,255,95)",
	"rgb(135,255,135)",
	"rgb(135,255,175)",
	"rgb(135,255,215)",
	"rgb(135,255,255)",
	"rgb(175,0,0)",
	"rgb(175,0,95)",
	"rgb(175,0,135)",
	"rgb(175,0,175)",
	"rgb(175,0,215)",
	"rgb(175,0,255)",
	"rgb(175,95,0)",
	"rgb(175,95,95)",
	"rgb(175,95,135)",
	"rgb(175,95,175)",
	"rgb(175,95,215)",
	"rgb(175,95,255)",
	"rgb(175,135,0)",
	"rgb(175,135,95)",
	"rgb(175,135,135)",
	"rgb(175,135,175)",
	"rgb(175,135,215)",
	"rgb(175,135,255)",
	"rgb(175,175,0)",
	"rgb(175,175,95)",
	"rgb(175,175,135)",
	"rgb(175,175,175)",
	"rgb(175,175,215)",
	"rgb(175,175,255)",
	"rgb(175,215,0)",
	"rgb(175,215,95)",
	"rgb(175,215,135)",
	"rgb(175,215,175)",
	"rgb(175,215,215)",
	"rgb(175,215,255)",
	"rgb(175,255,0)",
	"rgb(175,255,95)",
	"rgb(175,255,135)",
	"rgb(175,255,175)",
	"rgb(175,255,215)",
	"rgb(175,255,255)",
	"rgb(215,0,0)",
	"rgb(215,0,95)",
	"rgb(215,0,135)",
	"rgb(215,0,175)",
	"rgb(215,0,215)",
	"rgb(215,0,255)",
	"rgb(215,95,0)",
	"rgb(215,95,95)",
	"rgb(215,95,135)",
	"rgb(215,95,175)",
	"rgb(215,95,215)",
	"rgb(215,95,255)",
	"rgb(215,135,0)",
	"rgb(215,135,95)",
	"rgb(215,135,135)",
	"rgb(215,135,175)",
	"rgb(215,135,215)",
	"rgb(215,135,255)",
	"rgb(215,175,0)",
	"rgb(215,175,95)",
	"rgb(215,175,135)",
	"rgb(215,175,175)",
	"rgb(215,175,215)",
	"rgb(215,175,255)",
	"rgb(215,215,0)",
	"rgb(215,215,95)",
	"rgb(215,215,135)",
	"rgb(215,215,175)",
	"rgb(215,215,215)",
	"rgb(215,215,255)",
	"rgb(215,255,0)",
	"rgb(215,255,95)",
	"rgb(215,255,135)",
	"rgb(215,255,175)",
	"rgb(215,255,215)",
	"rgb(215,255,255)",
	"rgb(255,0,0)",
	"rgb(255,0,95)",
	"rgb(255,0,135)",
	"rgb(255,0,175)",
	"rgb(255,0,215)",
	"rgb(255,0,255)",
	"rgb(255,95,0)",
	"rgb(255,95,95)",
	"rgb(255,95,135)",
	"rgb(255,95,175)",
	"rgb(255,95,215)",
	"rgb(255,95,255)",
	"rgb(255,135,0)",
	"rgb(255,135,95)",
	"rgb(255,135,135)",
	"rgb(255,135,175)",
	"rgb(255,135,215)",
	"rgb(255,135,255)",
	"rgb(255,175,0)",
	"rgb(255,175,95)",
	"rgb(255,175,135)",
	"rgb(255,175,175)",
	"rgb(255,175,215)",
	"rgb(255,175,255)",
	"rgb(255,215,0)",
	"rgb(255,215,95)",
	"rgb(255,215,135)",
	"rgb(255,215,175)",
	"rgb(255,215,215)",
	"rgb(255,215,255)",
	"rgb(255,255,0)",
	"rgb(255,255,95)",
	"rgb(255,255,135)",
	"rgb(255,255,175)",
	"rgb(255,255,215)",
	"rgb(255,255,255)",
	"rgb(8,8,8)",
	"rgb(18,18,18)",
	"rgb(28,28,28)",
	"rgb(38,38,38)",
	"rgb(48,48,48)",
	"rgb(58,58,58)",
	"rgb(68,68,68)",
	"rgb(78,78,78)",
	"rgb(88,88,88)",
	"rgb(98,98,98)",
	"rgb(108,108,108)",
	"rgb(118,118,118)",
	"rgb(128,128,128)",
	"rgb(138,138,138)",
	"rgb(148,148,148)",
	"rgb(158,158,158)",
	"rgb(168,168,168)",
	"rgb(178,178,178)",
	"rgb(188,188,188)",
	"rgb(198,198,198)",
	"rgb(208,208,208)",
	"rgb(218,218,218)",
	"rgb(228,228,228)",
	"rgb(238,238,238)",
}

// RgbString returns the colour as a string in the hex format.
func (c Colour) RgbString() string {
	if Black <= c && c <= Grey93 {
		return rgbstring[c]
	}
	var buf [3]byte
	n := fmtInt8(buf, uint8(c))
	return "!%RgbString(" + string(buf[n:]) + ")"
}

var hslstring = [256]string{
	"hsl(0,0%,0%)",
	"hsl(0,100%,25%)",
	"hsl(120,100%,25%)",
	"hsl(60,100%,25%)",
	"hsl(240,100%,25%)",
	"hsl(300,100%,25%)",
	"hsl(180,100%,25%)",
	"hsl(0,0%,75%)",
	"hsl(0,0%,50%)",
	"hsl(0,100%,50%)",
	"hsl(120,100%,50%)",
	"hsl(60,100%,50%)",
	"hsl(240,100%,50%)",
	"hsl(300,100%,50%)",
	"hsl(180,100%,50%)",
	"hsl(0,0%,100%)",
	"hsl(0,0%,0%)",
	"hsl(240,100%,18%)",
	"hsl(240,100%,26%)",
	"hsl(240,100%,34%)",
	"hsl(240,100%,42%)",
	"hsl(240,100%,50%)",
	"hsl(120,100%,18%)",
	"hsl(180,100%,18%)",
	"hsl(97,100%,26%)",
	"hsl(07,100%,34%)",
	"hsl(13,100%,42%)",
	"hsl(17,100%,50%)",
	"hsl(120,100%,26%)",
	"hsl(62,100%,26%)",
	"hsl(180,100%,26%)",
	"hsl(93,100%,34%)",
	"hsl(02,100%,42%)",
	"hsl(08,100%,50%)",
	"hsl(120,100%,34%)",
	"hsl(52,100%,34%)",
	"hsl(66,100%,34%)",
	"hsl(180,100%,34%)",
	"hsl(91,100%,42%)",
	"hsl(98,100%,50%)",
	"hsl(120,100%,42%)",
	"hsl(46,100%,42%)",
	"hsl(57,100%,42%)",
	"hsl(68,100%,42%)",
	"hsl(180,100%,42%)",
	"hsl(89,100%,50%)",
	"hsl(120,100%,50%)",
	"hsl(42,100%,50%)",
	"hsl(51,100%,50%)",
	"hsl(61,100%,50%)",
	"hsl(70,100%,50%)",
	"hsl(180,100%,50%)",
	"hsl(0,100%,18%)",
	"hsl(300,100%,18%)",
	"hsl(82,100%,26%)",
	"hsl(72,100%,34%)",
	"hsl(66,100%,42%)",
	"hsl(62,100%,50%)",
	"hsl(60,100%,18%)",
	"hsl(0,0%,37%)",
	"hsl(240,17%,45%)",
	"hsl(240,33%,52%)",
	"hsl(240,60%,60%)",
	"hsl(240,100%,68%)",
	"hsl(7,100%,26%)",
	"hsl(120,17%,45%)",
	"hsl(180,17%,45%)",
	"hsl(210,33%,52%)",
	"hsl(220,60%,60%)",
	"hsl(225,100%,68%)",
	"hsl(7,100%,34%)",
	"hsl(120,33%,52%)",
	"hsl(150,33%,52%)",
	"hsl(180,33%,52%)",
	"hsl(200,60%,60%)",
	"hsl(210,100%,68%)",
	"hsl(3,100%,42%)",
	"hsl(120,60%,60%)",
	"hsl(140,60%,60%)",
	"hsl(160,60%,60%)",
	"hsl(180,60%,60%)",
	"hsl(195,100%,68%)",
	"hsl(7,100%,50%)",
	"hsl(120,100%,68%)",
	"hsl(135,100%,68%)",
	"hsl(150,100%,68%)",
	"hsl(165,100%,68%)",
	"hsl(180,100%,68%)",
	"hsl(0,100%,26%)",
	"hsl(17,100%,26%)",
	"hsl(300,100%,26%)",
	"hsl(86,100%,34%)",
	"hsl(77,100%,42%)",
	"hsl(71,100%,50%)",
	"hsl(2,100%,26%)",
	"hsl(0,17%,45%)",
	"hsl(300,17%,45%)",
	"hsl(270,33%,52%)",
	"hsl(260,60%,60%)",
	"hsl(255,100%,68%)",
	"hsl(60,100%,26%)",
	"hsl(60,17%,45%)",
	"hsl(0,0%,52%)",
	"hsl(240,20%,60%)",
	"hsl(240,50%,68%)",
	"hsl(240,100%,76%)",
	"hsl(3,100%,34%)",
	"hsl(90,33%,52%)",
	"hsl(120,20%,60%)",
	"hsl(180,20%,60%)",
	"hsl(210,50%,68%)",
	"hsl(220,100%,76%)",
	"hsl(2,100%,42%)",
	"hsl(100,60%,60%)",
	"hsl(120,50%,68%)",
	"hsl(150,50%,68%)",
	"hsl(180,50%,68%)",
	"hsl(200,100%,76%)",
	"hsl(8,100%,50%)",
	"hsl(105,100%,68%)",
	"hsl(120,100%,76%)",
	"hsl(140,100%,76%)",
	"hsl(160,100%,76%)",
	"hsl(180,100%,76%)",
	"hsl(0,100%,34%)",
	"hsl(27,100%,34%)",
	"hsl(13,100%,34%)",
	"hsl(300,100%,34%)",
	"hsl(88,100%,42%)",
	"hsl(81,100%,50%)",
	"hsl(2,100%,34%)",
	"hsl(0,33%,52%)",
	"hsl(330,33%,52%)",
	"hsl(300,33%,52%)",
	"hsl(280,60%,60%)",
	"hsl(270,100%,68%)",
	"hsl(6,100%,34%)",
	"hsl(30,33%,52%)",
	"hsl(0,20%,60%)",
	"hsl(300,20%,60%)",
	"hsl(270,50%,68%)",
	"hsl(260,100%,76%)",
	"hsl(60,100%,34%)",
	"hsl(60,33%,52%)",
	"hsl(60,20%,60%)",
	"hsl(0,0%,68%)",
	"hsl(240,33%,76%)",
	"hsl(240,100%,84%)",
	"hsl(1,100%,42%)",
	"hsl(80,60%,60%)",
	"hsl(90,50%,68%)",
	"hsl(120,33%,76%)",
	"hsl(180,33%,76%)",
	"hsl(210,100%,84%)",
	"hsl(8,100%,50%)",
	"hsl(90,100%,68%)",
	"hsl(100,100%,76%)",
	"hsl(120,100%,84%)",
	"hsl(150,100%,84%)",
	"hsl(180,100%,84%)",
	"hsl(0,100%,42%)",
	"hsl(33,100%,42%)",
	"hsl(22,100%,42%)",
	"hsl(11,100%,42%)",
	"hsl(300,100%,42%)",
	"hsl(90,100%,50%)",
	"hsl(6,100%,42%)",
	"hsl(0,60%,60%)",
	"hsl(340,60%,60%)",
	"hsl(320,60%,60%)",
	"hsl(300,60%,60%)",
	"hsl(285,100%,68%)",
	"hsl(7,100%,42%)",
	"hsl(20,60%,60%)",
	"hsl(0,50%,68%)",
	"hsl(330,50%,68%)",
	"hsl(300,50%,68%)",
	"hsl(280,100%,76%)",
	"hsl(8,100%,42%)",
	"hsl(40,60%,60%)",
	"hsl(30,50%,68%)",
	"hsl(0,33%,76%)",
	"hsl(300,33%,76%)",
	"hsl(270,100%,84%)",
	"hsl(60,100%,42%)",
	"hsl(60,60%,60%)",
	"hsl(60,50%,68%)",
	"hsl(60,33%,76%)",
	"hsl(0,0%,84%)",
	"hsl(240,100%,92%)",
	"hsl(9,100%,50%)",
	"hsl(75,100%,68%)",
	"hsl(80,100%,76%)",
	"hsl(90,100%,84%)",
	"hsl(120,100%,92%)",
	"hsl(180,100%,92%)",
	"hsl(0,100%,50%)",
	"hsl(37,100%,50%)",
	"hsl(28,100%,50%)",
	"hsl(18,100%,50%)",
	"hsl(09,100%,50%)",
	"hsl(300,100%,50%)",
	"hsl(2,100%,50%)",
	"hsl(0,100%,68%)",
	"hsl(345,100%,68%)",
	"hsl(330,100%,68%)",
	"hsl(315,100%,68%)",
	"hsl(300,100%,68%)",
	"hsl(1,100%,50%)",
	"hsl(15,100%,68%)",
	"hsl(0,100%,76%)",
	"hsl(340,100%,76%)",
	"hsl(320,100%,76%)",
	"hsl(300,100%,76%)",
	"hsl(1,100%,50%)",
	"hsl(30,100%,68%)",
	"hsl(20,100%,76%)",
	"hsl(0,100%,84%)",
	"hsl(330,100%,84%)",
	"hsl(300,100%,84%)",
	"hsl(0,100%,50%)",
	"hsl(45,100%,68%)",
	"hsl(40,100%,76%)",
	"hsl(30,100%,84%)",
	"hsl(0,100%,92%)",
	"hsl(300,100%,92%)",
	"hsl(60,100%,50%)",
	"hsl(60,100%,68%)",
	"hsl(60,100%,76%)",
	"hsl(60,100%,84%)",
	"hsl(60,100%,92%)",
	"hsl(0,0%,100%)",
	"hsl(0,0%,3%)",
	"hsl(0,0%,7%)",
	"hsl(0,0%,10%)",
	"hsl(0,0%,14%)",
	"hsl(0,0%,18%)",
	"hsl(0,0%,22%)",
	"hsl(0,0%,26%)",
	"hsl(0,0%,30%)",
	"hsl(0,0%,34%)",
	"hsl(0,0%,37%)",
	"hsl(0,0%,40%)",
	"hsl(0,0%,46%)",
	"hsl(0,0%,50%)",
	"hsl(0,0%,54%)",
	"hsl(0,0%,58%)",
	"hsl(0,0%,61%)",
	"hsl(0,0%,65%)",
	"hsl(0,0%,69%)",
	"hsl(0,0%,73%)",
	"hsl(0,0%,77%)",
	"hsl(0,0%,81%)",
	"hsl(0,0%,85%)",
	"hsl(0,0%,89%)",
	"hsl(0,0%,93%)",
}

// HslString returns the colour as a string in the hex format.
func (c Colour) HslString() string {
	if Black <= c && c <= Grey93 {
		return hslstring[c]
	}
	var buf [3]byte
	n := fmtInt8(buf, uint8(c))
	return "!%HslString(" + string(buf[n:]) + ")"
}

var term = [256][]byte{
	{'\033', '[', '3', '8', ';', '5', ';', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '3', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '3', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '3', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '3', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '3', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '3', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '3', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '3', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '3', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '3', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '4', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '4', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '4', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '4', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '4', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '4', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '4', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '4', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '4', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '4', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '5', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '5', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '5', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '5', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '5', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '5', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '5', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '5', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '5', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '5', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '6', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '6', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '6', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '6', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '6', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '6', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '6', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '6', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '6', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '6', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '7', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '7', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '7', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '7', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '7', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '7', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '7', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '7', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '7', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '7', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '8', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '8', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '8', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '8', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '8', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '8', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '8', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '8', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '8', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '8', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '9', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '9', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '9', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '9', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '9', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '9', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '9', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '9', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '9', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '9', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '0', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '0', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '0', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '0', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '0', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '0', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '0', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '0', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '0', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '0', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '1', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '1', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '1', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '1', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '1', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '1', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '1', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '1', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '1', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '1', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '2', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '2', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '2', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '2', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '2', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '2', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '2', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '2', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '2', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '2', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '3', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '3', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '3', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '3', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '3', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '3', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '3', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '3', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '3', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '3', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '4', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '4', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '4', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '4', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '4', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '4', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '4', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '4', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '4', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '4', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '5', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '5', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '5', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '5', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '5', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '5', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '5', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '5', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '5', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '5', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '6', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '6', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '6', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '6', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '6', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '6', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '6', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '6', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '6', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '6', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '7', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '7', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '7', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '7', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '7', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '7', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '7', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '7', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '7', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '7', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '8', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '8', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '8', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '8', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '8', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '8', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '8', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '8', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '8', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '8', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '9', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '9', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '9', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '9', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '9', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '9', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '9', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '9', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '9', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '1', '9', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '0', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '0', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '0', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '0', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '0', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '0', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '0', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '0', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '0', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '0', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '1', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '1', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '1', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '1', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '1', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '1', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '1', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '1', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '1', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '1', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '2', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '2', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '2', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '2', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '2', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '2', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '2', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '2', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '2', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '2', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '3', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '3', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '3', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '3', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '3', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '3', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '3', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '3', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '3', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '3', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '4', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '4', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '4', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '4', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '4', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '4', '5', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '4', '6', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '4', '7', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '4', '8', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '4', '9', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '5', '0', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '5', '1', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '5', '2', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '5', '3', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '5', '4', 'm'},
	{'\033', '[', '3', '8', ';', '5', ';', '2', '5', '5', 'm'},
}

// Bytes returns a byteslice of the color in the hex format.
func (c Colour) Bytes() []byte {
	if Black <= c && c <= Grey93 {
		return term[c]
	}
	var buf [3]byte
	n := fmtInt8(buf, uint8(c))
	return []byte("!%Bytes(" + string(buf[n:]) + ")")
}

var hexbytes = [256][]byte{
	{'#', '0', '0', '0', '0', '0', '0'},
	{'#', '8', '0', '0', '0', '0', '0'},
	{'#', '0', '0', '8', '0', '0', '0'},
	{'#', '8', '0', '8', '0', '0', '0'},
	{'#', '0', '0', '0', '0', '8', '0'},
	{'#', '8', '0', '0', '0', '8', '0'},
	{'#', '0', '0', '8', '0', '8', '0'},
	{'#', 'c', '1', 'c', '0', 'c', '0'},
	{'#', '8', '0', '8', '0', '8', '0'},
	{'#', 'f', 'f', '0', '0', '0', '1'},
	{'#', '0', '0', 'f', 'f', '0', '0'},
	{'#', 'f', 'f', 'f', 'f', '0', '0'},
	{'#', '0', '0', '0', '0', 'f', 'f'},
	{'#', 'f', 'f', '0', '0', 'f', 'f'},
	{'#', '0', '0', 'f', 'f', 'f', 'f'},
	{'#', 'f', 'f', 'f', 'f', 'f', 'f'},
	{'#', '0', '0', '0', '0', '0', '0'},
	{'#', '0', '0', '0', '0', '5', 'f'},
	{'#', '0', '0', '0', '0', '8', '7'},
	{'#', '0', '0', '0', '0', 'a', 'f'},
	{'#', '0', '0', '0', '0', 'd', '7'},
	{'#', '0', '0', '0', '0', 'f', 'f'},
	{'#', '0', '0', '5', 'f', '0', '0'},
	{'#', '0', '0', '5', 'f', '5', 'f'},
	{'#', '0', '0', '5', 'f', '8', '7'},
	{'#', '0', '0', '5', 'f', 'a', 'f'},
	{'#', '0', '0', '5', 'f', 'd', '7'},
	{'#', '0', '0', '5', 'f', 'f', 'f'},
	{'#', '0', '0', '8', '7', '0', '0'},
	{'#', '0', '0', '8', '7', '5', 'f'},
	{'#', '0', '0', '8', '7', '8', '7'},
	{'#', '0', '0', '8', '7', 'a', 'f'},
	{'#', '0', '0', '8', '7', 'd', '7'},
	{'#', '0', '0', '8', '7', 'f', 'f'},
	{'#', '0', '0', 'a', 'f', '0', '0'},
	{'#', '0', '0', 'a', 'f', '5', 'f'},
	{'#', '0', '0', 'a', 'f', '8', '7'},
	{'#', '0', '0', 'a', 'f', 'a', 'f'},
	{'#', '0', '0', 'a', 'f', 'd', '7'},
	{'#', '0', '0', 'a', 'f', 'f', 'f'},
	{'#', '0', '0', 'd', '7', '0', '0'},
	{'#', '0', '0', 'd', '7', '5', 'f'},
	{'#', '0', '0', 'd', '7', '8', '7'},
	{'#', '0', '0', 'd', '7', 'a', 'f'},
	{'#', '0', '0', 'd', '7', 'd', '7'},
	{'#', '0', '0', 'd', '7', 'f', 'f'},
	{'#', '0', '0', 'f', 'f', '0', '0'},
	{'#', '0', '0', 'f', 'f', '5', 'f'},
	{'#', '0', '0', 'f', 'f', '8', '7'},
	{'#', '0', '0', 'f', 'f', 'a', 'f'},
	{'#', '0', '0', 'f', 'f', 'd', '7'},
	{'#', '0', '0', 'f', 'f', 'f', 'f'},
	{'#', '5', 'f', '0', '0', '0', '0'},
	{'#', '5', 'f', '0', '0', '5', 'f'},
	{'#', '5', 'f', '0', '0', '8', '7'},
	{'#', '5', 'f', '0', '0', 'a', 'f'},
	{'#', '5', 'f', '0', '0', 'd', '7'},
	{'#', '5', 'f', '0', '0', 'f', 'f'},
	{'#', '5', 'f', '5', 'f', '0', '0'},
	{'#', '5', 'f', '5', 'f', '5', 'f'},
	{'#', '5', 'f', '5', 'f', '8', '7'},
	{'#', '5', 'f', '5', 'f', 'a', 'f'},
	{'#', '5', 'f', '5', 'f', 'd', '7'},
	{'#', '5', 'f', '5', 'f', 'f', 'f'},
	{'#', '5', 'f', '8', '7', '0', '0'},
	{'#', '5', 'f', '8', '7', '5', 'f'},
	{'#', '5', 'f', '8', '7', '8', '7'},
	{'#', '5', 'f', '8', '7', 'a', 'f'},
	{'#', '5', 'f', '8', '7', 'd', '7'},
	{'#', '5', 'f', '8', '7', 'f', 'f'},
	{'#', '5', 'f', 'a', 'f', '0', '0'},
	{'#', '5', 'f', 'a', 'f', '5', 'f'},
	{'#', '5', 'f', 'a', 'f', '8', '7'},
	{'#', '5', 'f', 'a', 'f', 'a', 'f'},
	{'#', '5', 'f', 'a', 'f', 'd', '7'},
	{'#', '5', 'f', 'a', 'f', 'f', 'f'},
	{'#', '5', 'f', 'd', '7', '0', '0'},
	{'#', '5', 'f', 'd', '7', '5', 'f'},
	{'#', '5', 'f', 'd', '7', '8', '7'},
	{'#', '5', 'f', 'd', '7', 'a', 'f'},
	{'#', '5', 'f', 'd', '7', 'd', '7'},
	{'#', '5', 'f', 'd', '7', 'f', 'f'},
	{'#', '5', 'f', 'f', 'f', '0', '0'},
	{'#', '5', 'f', 'f', 'f', '5', 'f'},
	{'#', '5', 'f', 'f', 'f', '8', '7'},
	{'#', '5', 'f', 'f', 'f', 'a', 'f'},
	{'#', '5', 'f', 'f', 'f', 'd', '7'},
	{'#', '5', 'f', 'f', 'f', 'f', 'f'},
	{'#', '8', '7', '0', '0', '0', '0'},
	{'#', '8', '7', '0', '0', '5', 'f'},
	{'#', '8', '7', '0', '0', '8', '7'},
	{'#', '8', '7', '0', '0', 'a', 'f'},
	{'#', '8', '7', '0', '0', 'd', '7'},
	{'#', '8', '7', '0', '0', 'f', 'f'},
	{'#', '8', '7', '5', 'f', '0', '0'},
	{'#', '8', '7', '5', 'f', '5', 'f'},
	{'#', '8', '7', '5', 'f', '8', '7'},
	{'#', '8', '7', '5', 'f', 'a', 'f'},
	{'#', '8', '7', '5', 'f', 'd', '7'},
	{'#', '8', '7', '5', 'f', 'f', 'f'},
	{'#', '8', '7', '8', '7', '0', '0'},
	{'#', '8', '7', '8', '7', '5', 'f'},
	{'#', '8', '7', '8', '7', '8', '7'},
	{'#', '8', '7', '8', '7', 'a', 'f'},
	{'#', '8', '7', '8', '7', 'd', '7'},
	{'#', '8', '7', '8', '7', 'f', 'f'},
	{'#', '8', '7', 'a', 'f', '0', '0'},
	{'#', '8', '7', 'a', 'f', '5', 'f'},
	{'#', '8', '7', 'a', 'f', '8', '7'},
	{'#', '8', '7', 'a', 'f', 'a', 'f'},
	{'#', '8', '7', 'a', 'f', 'd', '7'},
	{'#', '8', '7', 'a', 'f', 'f', 'f'},
	{'#', '8', '7', 'd', '7', '0', '0'},
	{'#', '8', '7', 'd', '7', '5', 'f'},
	{'#', '8', '7', 'd', '7', '8', '7'},
	{'#', '8', '7', 'd', '7', 'a', 'f'},
	{'#', '8', '7', 'd', '7', 'd', '7'},
	{'#', '8', '7', 'd', '7', 'f', 'f'},
	{'#', '8', '7', 'f', 'f', '0', '0'},
	{'#', '8', '7', 'f', 'f', '5', 'f'},
	{'#', '8', '7', 'f', 'f', '8', '7'},
	{'#', '8', '7', 'f', 'f', 'a', 'f'},
	{'#', '8', '7', 'f', 'f', 'd', '7'},
	{'#', '8', '7', 'f', 'f', 'f', 'f'},
	{'#', 'a', 'f', '0', '0', '0', '0'},
	{'#', 'a', 'f', '0', '0', '5', 'f'},
	{'#', 'a', 'f', '0', '0', '8', '7'},
	{'#', 'a', 'f', '0', '0', 'a', 'f'},
	{'#', 'a', 'f', '0', '0', 'd', '7'},
	{'#', 'a', 'f', '0', '0', 'f', 'f'},
	{'#', 'a', 'f', '5', 'f', '0', '0'},
	{'#', 'a', 'f', '5', 'f', '5', 'f'},
	{'#', 'a', 'f', '5', 'f', '8', '7'},
	{'#', 'a', 'f', '5', 'f', 'a', 'f'},
	{'#', 'a', 'f', '5', 'f', 'd', '7'},
	{'#', 'a', 'f', '5', 'f', 'f', 'f'},
	{'#', 'a', 'f', '8', '7', '0', '0'},
	{'#', 'a', 'f', '8', '7', '5', 'f'},
	{'#', 'a', 'f', '8', '7', '8', '7'},
	{'#', 'a', 'f', '8', '7', 'a', 'f'},
	{'#', 'a', 'f', '8', '7', 'd', '7'},
	{'#', 'a', 'f', '8', '7', 'f', 'f'},
	{'#', 'a', 'f', 'a', 'f', '0', '0'},
	{'#', 'a', 'f', 'a', 'f', '5', 'f'},
	{'#', 'a', 'f', 'a', 'f', '8', '7'},
	{'#', 'a', 'f', 'a', 'f', 'a', 'f'},
	{'#', 'a', 'f', 'a', 'f', 'd', '7'},
	{'#', 'a', 'f', 'a', 'f', 'f', 'f'},
	{'#', 'a', 'f', 'd', '7', '0', '0'},
	{'#', 'a', 'f', 'd', '7', '5', 'f'},
	{'#', 'a', 'f', 'd', '7', '8', '7'},
	{'#', 'a', 'f', 'd', '7', 'a', 'f'},
	{'#', 'a', 'f', 'd', '7', 'd', '7'},
	{'#', 'a', 'f', 'd', '7', 'f', 'f'},
	{'#', 'a', 'f', 'f', 'f', '0', '0'},
	{'#', 'a', 'f', 'f', 'f', '5', 'f'},
	{'#', 'a', 'f', 'f', 'f', '8', '7'},
	{'#', 'a', 'f', 'f', 'f', 'a', 'f'},
	{'#', 'a', 'f', 'f', 'f', 'd', '7'},
	{'#', 'a', 'f', 'f', 'f', 'f', 'f'},
	{'#', 'd', '7', '0', '0', '0', '0'},
	{'#', 'd', '7', '0', '0', '5', 'f'},
	{'#', 'd', '7', '0', '0', '8', '7'},
	{'#', 'd', '7', '0', '0', 'a', 'f'},
	{'#', 'd', '7', '0', '0', 'd', '7'},
	{'#', 'd', '7', '0', '0', 'f', 'f'},
	{'#', 'd', '7', '5', 'f', '0', '0'},
	{'#', 'd', '7', '5', 'f', '5', 'f'},
	{'#', 'd', '7', '5', 'f', '8', '7'},
	{'#', 'd', '7', '5', 'f', 'a', 'f'},
	{'#', 'd', '7', '5', 'f', 'd', '7'},
	{'#', 'd', '7', '5', 'f', 'f', 'f'},
	{'#', 'd', '7', '8', '7', '0', '0'},
	{'#', 'd', '7', '8', '7', '5', 'f'},
	{'#', 'd', '7', '8', '7', '8', '7'},
	{'#', 'd', '7', '8', '7', 'a', 'f'},
	{'#', 'd', '7', '8', '7', 'd', '7'},
	{'#', 'd', '7', '8', '7', 'f', 'f'},
	{'#', 'd', '7', 'a', 'f', '0', '0'},
	{'#', 'd', '7', 'a', 'f', '5', 'f'},
	{'#', 'd', '7', 'a', 'f', '8', '7'},
	{'#', 'd', '7', 'a', 'f', 'a', 'f'},
	{'#', 'd', '7', 'a', 'f', 'd', '7'},
	{'#', 'd', '7', 'a', 'f', 'f', 'f'},
	{'#', 'd', '7', 'd', '7', '0', '0'},
	{'#', 'd', '7', 'd', '7', '5', 'f'},
	{'#', 'd', '7', 'd', '7', '8', '7'},
	{'#', 'd', '7', 'd', '7', 'a', 'f'},
	{'#', 'd', '7', 'd', '7', 'd', '7'},
	{'#', 'd', '7', 'd', '7', 'f', 'f'},
	{'#', 'd', '7', 'f', 'f', '0', '0'},
	{'#', 'd', '7', 'f', 'f', '5', 'f'},
	{'#', 'd', '7', 'f', 'f', '8', '7'},
	{'#', 'd', '7', 'f', 'f', 'a', 'f'},
	{'#', 'd', '7', 'f', 'f', 'd', '7'},
	{'#', 'd', '7', 'f', 'f', 'f', 'f'},
	{'#', 'f', 'f', '0', '0', '0', '0'},
	{'#', 'f', 'f', '0', '0', '5', 'f'},
	{'#', 'f', 'f', '0', '0', '8', '7'},
	{'#', 'f', 'f', '0', '0', 'a', 'f'},
	{'#', 'f', 'f', '0', '0', 'd', '7'},
	{'#', 'f', 'f', '0', '0', 'f', 'f'},
	{'#', 'f', 'f', '5', 'f', '0', '0'},
	{'#', 'f', 'f', '5', 'f', '5', 'f'},
	{'#', 'f', 'f', '5', 'f', '8', '7'},
	{'#', 'f', 'f', '5', 'f', 'a', 'f'},
	{'#', 'f', 'f', '5', 'f', 'd', '7'},
	{'#', 'f', 'f', '5', 'f', 'f', 'f'},
	{'#', 'f', 'f', '8', '7', '0', '0'},
	{'#', 'f', 'f', '8', '7', '5', 'f'},
	{'#', 'f', 'f', '8', '7', '8', '7'},
	{'#', 'f', 'f', '8', '7', 'a', 'f'},
	{'#', 'f', 'f', '8', '7', 'd', '7'},
	{'#', 'f', 'f', '8', '7', 'f', 'f'},
	{'#', 'f', 'f', 'a', 'f', '0', '0'},
	{'#', 'f', 'f', 'a', 'f', '5', 'f'},
	{'#', 'f', 'f', 'a', 'f', '8', '7'},
	{'#', 'f', 'f', 'a', 'f', 'a', 'f'},
	{'#', 'f', 'f', 'a', 'f', 'd', '7'},
	{'#', 'f', 'f', 'a', 'f', 'f', 'f'},
	{'#', 'f', 'f', 'd', '7', '0', '0'},
	{'#', 'f', 'f', 'd', '7', '5', 'f'},
	{'#', 'f', 'f', 'd', '7', '8', '7'},
	{'#', 'f', 'f', 'd', '7', 'a', 'f'},
	{'#', 'f', 'f', 'd', '7', 'd', '7'},
	{'#', 'f', 'f', 'd', '7', 'f', 'f'},
	{'#', 'f', 'f', 'f', 'f', '0', '0'},
	{'#', 'f', 'f', 'f', 'f', '5', 'f'},
	{'#', 'f', 'f', 'f', 'f', '8', '7'},
	{'#', 'f', 'f', 'f', 'f', 'a', 'f'},
	{'#', 'f', 'f', 'f', 'f', 'd', '7'},
	{'#', 'f', 'f', 'f', 'f', 'f', 'f'},
	{'#', '0', '8', '0', '8', '0', '8'},
	{'#', '1', '2', '1', '2', '1', '2'},
	{'#', '1', 'c', '1', 'c', '1', 'c'},
	{'#', '2', '6', '2', '6', '2', '6'},
	{'#', '3', '0', '3', '0', '3', '0'},
	{'#', '3', 'a', '3', 'a', '3', 'a'},
	{'#', '4', '4', '4', '4', '4', '4'},
	{'#', '4', 'e', '4', 'e', '4', 'e'},
	{'#', '5', '8', '5', '8', '5', '8'},
	{'#', '6', '2', '6', '2', '6', '2'},
	{'#', '6', 'c', '6', 'c', '6', 'c'},
	{'#', '7', '6', '7', '6', '7', '6'},
	{'#', '8', '0', '8', '0', '8', '0'},
	{'#', '8', 'a', '8', 'a', '8', 'a'},
	{'#', '9', '4', '9', '4', '9', '4'},
	{'#', '9', 'e', '9', 'e', '9', 'e'},
	{'#', 'a', '8', 'a', '8', 'a', '8'},
	{'#', 'b', '2', 'b', '2', 'b', '2'},
	{'#', 'b', 'c', 'b', 'c', 'b', 'c'},
	{'#', 'c', '6', 'c', '6', 'c', '6'},
	{'#', 'd', '0', 'd', '0', 'd', '0'},
	{'#', 'd', 'a', 'd', 'a', 'd', 'a'},
	{'#', 'e', '4', 'e', '4', 'e', '4'},
	{'#', 'e', 'e', 'e', 'e', 'e', 'e'},
}

// HexBytes returns a byteslice of the color in the hex format.
func (c Colour) HexBytes() []byte {
	if Black <= c && c <= Grey93 {
		return hexbytes[c]
	}
	var buf [3]byte
	n := fmtInt8(buf, uint8(c))
	return []byte("!%HexBytes(" + string(buf[n:]) + ")")
}

var rgbbytes = [256][]byte{
	{'r', 'g', 'b', '(', '0', ',', '0', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '1', '2', '8', ',', '0', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '1', '2', '8', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '1', '2', '8', ',', '1', '2', '8', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '0', ',', '1', '2', '8', ',', ')'},
	{'r', 'g', 'b', '(', '1', '2', '8', ',', '0', ',', '1', '2', '8', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '1', '2', '8', ',', '1', '2', '8', ',', ')'},
	{'r', 'g', 'b', '(', '1', '9', '2', ',', '1', '9', '2', ',', '1', '9', '2', ',', ')'},
	{'r', 'g', 'b', '(', '1', '2', '8', ',', '1', '2', '8', ',', '1', '2', '8', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '0', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '2', '5', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '2', '5', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '0', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '0', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '2', '5', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '2', '5', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '0', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '0', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '0', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '0', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '0', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '0', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '9', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '9', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '9', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '9', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '9', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '9', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '1', '3', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '1', '3', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '1', '3', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '1', '3', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '1', '3', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '1', '3', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '1', '7', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '1', '7', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '1', '7', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '1', '7', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '1', '7', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '1', '7', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '2', '1', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '2', '1', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '2', '1', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '2', '1', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '2', '1', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '2', '1', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '2', '5', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '2', '5', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '2', '5', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '2', '5', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '2', '5', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '0', ',', '2', '5', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '0', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '0', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '0', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '0', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '0', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '0', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '9', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '9', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '9', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '9', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '9', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '9', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '1', '3', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '1', '3', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '1', '3', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '1', '3', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '1', '3', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '1', '3', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '1', '7', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '1', '7', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '1', '7', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '1', '7', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '1', '7', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '1', '7', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '2', '1', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '2', '1', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '2', '1', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '2', '1', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '2', '1', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '2', '1', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '2', '5', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '2', '5', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '2', '5', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '2', '5', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '2', '5', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '9', '5', ',', '2', '5', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '0', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '0', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '0', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '0', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '0', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '0', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '9', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '9', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '9', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '9', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '9', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '9', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '1', '3', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '1', '3', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '1', '3', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '1', '3', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '1', '3', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '1', '3', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '1', '7', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '1', '7', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '1', '7', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '1', '7', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '1', '7', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '1', '7', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '2', '1', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '2', '1', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '2', '1', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '2', '1', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '2', '1', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '2', '1', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '2', '5', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '2', '5', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '2', '5', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '2', '5', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '2', '5', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '5', ',', '2', '5', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '0', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '0', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '0', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '0', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '0', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '0', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '9', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '9', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '9', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '9', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '9', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '9', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '1', '3', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '1', '3', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '1', '3', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '1', '3', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '1', '3', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '1', '3', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '1', '7', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '1', '7', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '1', '7', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '1', '7', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '1', '7', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '1', '7', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '2', '1', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '2', '1', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '2', '1', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '2', '1', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '2', '1', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '2', '1', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '2', '5', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '2', '5', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '2', '5', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '2', '5', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '2', '5', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '5', ',', '2', '5', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '0', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '0', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '0', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '0', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '0', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '0', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '9', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '9', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '9', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '9', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '9', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '9', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '1', '3', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '1', '3', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '1', '3', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '1', '3', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '1', '3', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '1', '3', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '1', '7', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '1', '7', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '1', '7', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '1', '7', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '1', '7', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '1', '7', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '2', '1', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '2', '1', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '2', '1', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '2', '1', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '2', '1', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '2', '1', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '2', '5', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '2', '5', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '2', '5', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '2', '5', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '2', '5', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '5', ',', '2', '5', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '0', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '0', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '0', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '0', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '0', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '0', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '9', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '9', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '9', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '9', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '9', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '9', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '1', '3', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '1', '3', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '1', '3', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '1', '3', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '1', '3', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '1', '3', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '1', '7', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '1', '7', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '1', '7', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '1', '7', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '1', '7', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '1', '7', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '2', '1', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '2', '1', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '2', '1', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '2', '1', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '2', '1', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '2', '1', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '2', '5', '5', ',', '0', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '2', '5', '5', ',', '9', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '2', '5', '5', ',', '1', '3', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '2', '5', '5', ',', '1', '7', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '2', '5', '5', ',', '2', '1', '5', ',', ')'},
	{'r', 'g', 'b', '(', '2', '5', '5', ',', '2', '5', '5', ',', '2', '5', '5', ',', ')'},
	{'r', 'g', 'b', '(', '8', ',', '8', ',', '8', ',', ')'},
	{'r', 'g', 'b', '(', '1', '8', ',', '1', '8', ',', '1', '8', ',', ')'},
	{'r', 'g', 'b', '(', '2', '8', ',', '2', '8', ',', '2', '8', ',', ')'},
	{'r', 'g', 'b', '(', '3', '8', ',', '3', '8', ',', '3', '8', ',', ')'},
	{'r', 'g', 'b', '(', '4', '8', ',', '4', '8', ',', '4', '8', ',', ')'},
	{'r', 'g', 'b', '(', '5', '8', ',', '5', '8', ',', '5', '8', ',', ')'},
	{'r', 'g', 'b', '(', '6', '8', ',', '6', '8', ',', '6', '8', ',', ')'},
	{'r', 'g', 'b', '(', '7', '8', ',', '7', '8', ',', '7', '8', ',', ')'},
	{'r', 'g', 'b', '(', '8', '8', ',', '8', '8', ',', '8', '8', ',', ')'},
	{'r', 'g', 'b', '(', '9', '8', ',', '9', '8', ',', '9', '8', ',', ')'},
	{'r', 'g', 'b', '(', '1', '0', '8', ',', '1', '0', '8', ',', '1', '0', '8', ',', ')'},
	{'r', 'g', 'b', '(', '1', '1', '8', ',', '1', '1', '8', ',', '1', '1', '8', ',', ')'},
	{'r', 'g', 'b', '(', '1', '2', '8', ',', '1', '2', '8', ',', '1', '2', '8', ',', ')'},
	{'r', 'g', 'b', '(', '1', '3', '8', ',', '1', '3', '8', ',', '1', '3', '8', ',', ')'},
	{'r', 'g', 'b', '(', '1', '4', '8', ',', '1', '4', '8', ',', '1', '4', '8', ',', ')'},
	{'r', 'g', 'b', '(', '1', '5', '8', ',', '1', '5', '8', ',', '1', '5', '8', ',', ')'},
	{'r', 'g', 'b', '(', '1', '6', '8', ',', '1', '6', '8', ',', '1', '6', '8', ',', ')'},
	{'r', 'g', 'b', '(', '1', '7', '8', ',', '1', '7', '8', ',', '1', '7', '8', ',', ')'},
	{'r', 'g', 'b', '(', '1', '8', '8', ',', '1', '8', '8', ',', '1', '8', '8', ',', ')'},
	{'r', 'g', 'b', '(', '1', '9', '8', ',', '1', '9', '8', ',', '1', '9', '8', ',', ')'},
	{'r', 'g', 'b', '(', '2', '0', '8', ',', '2', '0', '8', ',', '2', '0', '8', ',', ')'},
	{'r', 'g', 'b', '(', '2', '1', '8', ',', '2', '1', '8', ',', '2', '1', '8', ',', ')'},
	{'r', 'g', 'b', '(', '2', '2', '8', ',', '2', '2', '8', ',', '2', '2', '8', ',', ')'},
	{'r', 'g', 'b', '(', '2', '3', '8', ',', '2', '3', '8', ',', '2', '3', '8', ',', ')'},
}

// RgbBytes returns a byteslice of the color in the hex format.
func (c Colour) RgbBytes() []byte {
	if Black <= c && c <= Grey93 {
		return rgbbytes[c]
	}
	var buf [3]byte
	n := fmtInt8(buf, uint8(c))
	return []byte("!%RgbBytes(" + string(buf[n:]) + ")")
}

var hslbytes = [256][]byte{
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '1', '0', '0', '%', ',', '2', '5', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '2', '0', ',', '1', '0', '0', '%', ',', '2', '5', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '0', ',', '1', '0', '0', '%', ',', '2', '5', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '4', '0', ',', '1', '0', '0', '%', ',', '2', '5', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', '0', ',', '1', '0', '0', '%', ',', '2', '5', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '8', '0', ',', '1', '0', '0', '%', ',', '2', '5', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '7', '5', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '2', '0', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '0', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '4', '0', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', '0', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '8', '0', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '1', '0', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '4', '0', ',', '1', '0', '0', '%', ',', '1', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '4', '0', ',', '1', '0', '0', '%', ',', '2', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '4', '0', ',', '1', '0', '0', '%', ',', '3', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '4', '0', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '4', '0', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '2', '0', ',', '1', '0', '0', '%', ',', '1', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '8', '0', ',', '1', '0', '0', '%', ',', '1', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '9', '7', ',', '1', '0', '0', '%', ',', '2', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', '7', ',', '1', '0', '0', '%', ',', '3', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '3', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '7', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '2', '0', ',', '1', '0', '0', '%', ',', '2', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '2', ',', '1', '0', '0', '%', ',', '2', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '8', '0', ',', '1', '0', '0', '%', ',', '2', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '9', '3', ',', '1', '0', '0', '%', ',', '3', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', '2', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', '8', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '2', '0', ',', '1', '0', '0', '%', ',', '3', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '5', '2', ',', '1', '0', '0', '%', ',', '3', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '6', ',', '1', '0', '0', '%', ',', '3', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '8', '0', ',', '1', '0', '0', '%', ',', '3', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '9', '1', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '9', '8', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '2', '0', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '4', '6', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '5', '7', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '8', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '8', '0', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '8', '9', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '2', '0', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '4', '2', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '5', '1', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '1', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '7', '0', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '8', '0', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '1', '0', '0', '%', ',', '1', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', '0', ',', '1', '0', '0', '%', ',', '1', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '8', '2', ',', '1', '0', '0', '%', ',', '2', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '7', '2', ',', '1', '0', '0', '%', ',', '3', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '6', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '2', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '0', ',', '1', '0', '0', '%', ',', '1', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '3', '7', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '4', '0', ',', '1', '7', '%', ',', '4', '5', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '4', '0', ',', '3', '3', '%', ',', '5', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '4', '0', ',', '6', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '4', '0', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '7', ',', '1', '0', '0', '%', ',', '2', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '2', '0', ',', '1', '7', '%', ',', '4', '5', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '8', '0', ',', '1', '7', '%', ',', '4', '5', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '1', '0', ',', '3', '3', '%', ',', '5', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '2', '0', ',', '6', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '2', '5', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '7', ',', '1', '0', '0', '%', ',', '3', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '2', '0', ',', '3', '3', '%', ',', '5', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '5', '0', ',', '3', '3', '%', ',', '5', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '8', '0', ',', '3', '3', '%', ',', '5', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '0', '0', ',', '6', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '1', '0', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '2', '0', ',', '6', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '4', '0', ',', '6', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '6', '0', ',', '6', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '8', '0', ',', '6', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '9', '5', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '7', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '2', '0', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '3', '5', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '5', '0', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '6', '5', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '8', '0', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '1', '0', '0', '%', ',', '2', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '7', ',', '1', '0', '0', '%', ',', '2', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', '0', ',', '1', '0', '0', '%', ',', '2', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '8', '6', ',', '1', '0', '0', '%', ',', '3', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '7', '7', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '7', '1', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', ',', '1', '0', '0', '%', ',', '2', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '1', '7', '%', ',', '4', '5', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', '0', ',', '1', '7', '%', ',', '4', '5', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '7', '0', ',', '3', '3', '%', ',', '5', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '6', '0', ',', '6', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '5', '5', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '0', ',', '1', '0', '0', '%', ',', '2', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '0', ',', '1', '7', '%', ',', '4', '5', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '5', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '4', '0', ',', '2', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '4', '0', ',', '5', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '4', '0', ',', '1', '0', '0', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', ',', '1', '0', '0', '%', ',', '3', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '9', '0', ',', '3', '3', '%', ',', '5', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '2', '0', ',', '2', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '8', '0', ',', '2', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '1', '0', ',', '5', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '2', '0', ',', '1', '0', '0', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '0', '0', ',', '6', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '2', '0', ',', '5', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '5', '0', ',', '5', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '8', '0', ',', '5', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '0', '0', ',', '1', '0', '0', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '8', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '0', '5', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '2', '0', ',', '1', '0', '0', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '4', '0', ',', '1', '0', '0', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '6', '0', ',', '1', '0', '0', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '8', '0', ',', '1', '0', '0', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '1', '0', '0', '%', ',', '3', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '7', ',', '1', '0', '0', '%', ',', '3', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '3', ',', '1', '0', '0', '%', ',', '3', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', '0', ',', '1', '0', '0', '%', ',', '3', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '8', '8', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '8', '1', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', ',', '1', '0', '0', '%', ',', '3', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '3', '3', '%', ',', '5', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '3', '0', ',', '3', '3', '%', ',', '5', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', '0', ',', '3', '3', '%', ',', '5', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '8', '0', ',', '6', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '7', '0', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', ',', '1', '0', '0', '%', ',', '3', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', ',', '3', '3', '%', ',', '5', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '2', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', '0', ',', '2', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '7', '0', ',', '5', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '6', '0', ',', '1', '0', '0', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '0', ',', '1', '0', '0', '%', ',', '3', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '0', ',', '3', '3', '%', ',', '5', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '0', ',', '2', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '4', '0', ',', '3', '3', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '4', '0', ',', '1', '0', '0', '%', ',', '8', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '8', '0', ',', '6', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '9', '0', ',', '5', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '2', '0', ',', '3', '3', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '8', '0', ',', '3', '3', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '1', '0', ',', '1', '0', '0', '%', ',', '8', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '8', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '9', '0', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '0', '0', ',', '1', '0', '0', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '2', '0', ',', '1', '0', '0', '%', ',', '8', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '5', '0', ',', '1', '0', '0', '%', ',', '8', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '8', '0', ',', '1', '0', '0', '%', ',', '8', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '3', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '2', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '1', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', '0', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '9', '0', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '6', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '4', '0', ',', '6', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '2', '0', ',', '6', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', '0', ',', '6', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '8', '5', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '7', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '0', ',', '6', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '5', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '3', '0', ',', '5', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', '0', ',', '5', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '8', '0', ',', '1', '0', '0', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '8', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '4', '0', ',', '6', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', ',', '5', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '3', '3', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', '0', ',', '3', '3', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '7', '0', ',', '1', '0', '0', '%', ',', '8', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '0', ',', '1', '0', '0', '%', ',', '4', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '0', ',', '6', '0', '%', ',', '6', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '0', ',', '5', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '0', ',', '3', '3', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '8', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '4', '0', ',', '1', '0', '0', '%', ',', '9', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '9', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '7', '5', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '8', '0', ',', '1', '0', '0', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '9', '0', ',', '1', '0', '0', '%', ',', '8', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '2', '0', ',', '1', '0', '0', '%', ',', '9', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '8', '0', ',', '1', '0', '0', '%', ',', '9', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '7', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '8', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '8', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', '9', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', '0', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '4', '5', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '3', '0', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '1', '5', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', '0', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', '5', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '1', '0', '0', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '4', '0', ',', '1', '0', '0', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '2', '0', ',', '1', '0', '0', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', '0', ',', '1', '0', '0', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '1', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '2', '0', ',', '1', '0', '0', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '1', '0', '0', '%', ',', '8', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '3', '0', ',', '1', '0', '0', '%', ',', '8', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', '0', ',', '1', '0', '0', '%', ',', '8', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '4', '5', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '4', '0', ',', '1', '0', '0', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', ',', '1', '0', '0', '%', ',', '8', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '1', '0', '0', '%', ',', '9', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '3', '0', '0', ',', '1', '0', '0', '%', ',', '9', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '0', ',', '1', '0', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '0', ',', '1', '0', '0', '%', ',', '6', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '0', ',', '1', '0', '0', '%', ',', '7', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '0', ',', '1', '0', '0', '%', ',', '8', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '6', '0', ',', '1', '0', '0', '%', ',', '9', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '1', '0', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '3', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '7', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '1', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '1', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '1', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '2', '2', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '2', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '3', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '3', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '3', '7', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '4', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '4', '6', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '5', '0', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '5', '4', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '5', '8', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '6', '1', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '6', '5', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '6', '9', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '7', '3', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '7', '7', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '8', '1', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '8', '5', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '8', '9', '%', ')'},
	{'h', 's', 'l', '(', ',', '0', ',', '0', '%', ',', '9', '3', '%', ')'},
}

// HslBytes returns a byteslice of the color in the hex format.
func (c Colour) HslBytes() []byte {
	if Black <= c && c <= Grey93 {
		return hslbytes[c]
	}
	var buf [3]byte
	n := fmtInt8(buf, uint8(c))
	return []byte("!%HslBytes(" + string(buf[n:]) + ")")
}

func fmtInt8(buf [3]byte, v uint8) int {
	w := len(buf)
	if v == 0 {
		w--
		buf[w] = '0'
		return w
	}
	for v > 0 {
		w--
		buf[w] = byte(v%10) + '0'
		v /= 10
	}
	return w
}

// Black
// Maroon
// Green
// Olive
// Navy
// Purple
// Teal
// Silver
// Grey
// Red
// Lime
// Yellow
// Blue
// Fuchsia
// Aqua
// White
// Grey0
// NavyBlue
// DarkBlue
// Blue3
// Blue0
// Blue1
// DarkGreen
// DeepSkyBlue7
// DeepSkyBlue6
// DeepSkyBlue5
// DodgerBlue3
// DodgerBlue2
// Green4
// SpringGreen6
// Turquoise4
// DeepSkyBlue4
// DeepSkyBlue3
// DodgerBlue1
// Green3
// SpringGreen5
// DarkCyan
// LightSeaGreen
// DeepSkyBlue2
// DeepSkyBlue1
// Green2
// SpringGreen4
// SpringGreen3
// Cyan3
// DarkTurquoise
// Turquoise2
// Green1
// SpringGreen2
// SpringGreen1
// MediumSpringGreen
// Cyan2
// Cyan1
// DarkRed1
// DeepPink8
// Purple5
// Purple4
// Purple3
// BlueViolet
// Orange4
// Grey37
// MediumPurple7
// SlateBlue3
// SlateBlue2
// RoyalBlue1
// Chartreuse6
// DarkSeaGreen9
// PaleTurquoise4
// SteelBlue4
// SteelBlue3
// CornflowerBlue
// Chartreuse5
// DarkSeaGreen8
// CadetBlue1
// CadetBlue2
// SkyBlue3
// SteelBlue2
// Chartreuse4
// PaleGreen4
// SeaGreen4
// Aquamarine3
// MediumTurquoise
// SteelBlue1
// Chartreuse3
// SeaGreen3
// SeaGreen2
// SeaGreen1
// Aquamarine2
// DarkSlateGray2
// DarkRed2
// DeepPink7
// DarkMagenta1
// DarkMagenta2
// DarkViolet1
// Purple2
// Orange3
// LightPink4
// Plum4
// MediumPurple6
// MediumPurple5
// SlateBlue1
// Yellow6
// Wheat4
// Grey53
// LightSlateGrey
// MediumPurple4
// LightSlateBlue
// Yellow5
// DarkOliveGreen6
// DarkSeaGreen7
// LightSkyBlue3
// LightSkyBlue2
// SkyBlue2
// Chartreuse2
// DarkOliveGreen5
// PaleGreen3
// DarkSeaGreen6
// DarkSlateGray3
// SkyBlue1
// Chartreuse1
// LightGreen1
// LightGreen2
// PaleGreen2
// Aquamarine1
// DarkSlateGray1
// Red3
// DeepPink6
// MediumVioletRed
// Magenta6
// DarkViolet2
// Purple1
// DarkOrange3
// IndianRed4
// HotPink5
// MediumOrchid4
// MediumOrchid3
// MediumPurple3
// DarkGoldenrod
// LightSalmon3
// RosyBrown
// Grey63
// MediumPurple2
// MediumPurple1
// Gold3
// DarkKhaki
// NavajoWhite3
// Grey69
// LightSteelBlue3
// LightSteelBlue
// Yellow4
// DarkOliveGreen4
// DarkSeaGreen5
// DarkSeaGreen4
// LightCyan3
// LightSkyBlue1
// GreenYellow
// DarkOliveGreen3
// PaleGreen1
// DarkSeaGreen3
// DarkSeaGreen2
// PaleTurquoise1
// Red2
// DeepPink5
// DeepPink4
// Magenta5
// Magenta4
// Magenta3
// DarkOrange2
// IndianRed3
// HotPink4
// HotPink3
// Orchid
// MediumOrchid2
// Orange2
// LightSalmon2
// LightPink3
// Pink3
// Plum3
// Violet
// Gold2
// LightGoldenrod5
// Tan
// MistyRose3
// Thistle3
// Plum2
// Yellow3
// Khaki3
// LightGoldenrod4
// LightYellow3
// Grey84
// LightSteelBlue1
// Yellow2
// DarkOliveGreen2
// DarkOliveGreen1
// DarkSeaGreen1
// Honeydew2
// LightCyan1
// Red1
// DeepPink3
// DeepPink2
// DeepPink1
// Magenta2
// Magenta1
// OrangeRed1
// IndianRed2
// IndianRed1
// HotPink2
// HotPink1
// MediumOrchid1
// DarkOrange1
// Salmon1
// LightCoral
// PaleVioletRed1
// Orchid2
// Orchid1
// Orange1
// SandyBrown
// LightSalmon1
// LightPink1
// Pink1
// Plum1
// Gold1
// LightGoldenrod3
// LightGoldenrod2
// NavajoWhite1
// MistyRose1
// Thistle1
// Yellow1
// LightGoldenrod1
// Khaki1
// Wheat1
// Cornsilk1
// Grey100
// Grey3
// Grey7
// Grey11
// Grey15
// Grey19
// Grey23
// Grey27
// Grey30
// Grey35
// Grey39
// Grey42
// Grey46
// Grey50
// Grey54
// Grey58
// Grey62
// Grey66
// Grey70
// Grey74
// Grey78
// Grey82
// Grey85
// Grey89
// Grey93
