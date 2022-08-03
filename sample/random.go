package sample

import (
	"grpc_project/pb"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func randomKeyboardLayout() pb.Keyboard_Layout {
	//
	rKeyboard := rand.Intn(3)
	switch rKeyboard {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}
func randomCPUBrand() string {
	//
	return randomStringFromSet("Intel", "AMD")
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"COre i7-9750H",
			"Core i5-9400F",
			"Core i3-1005G1",
		)
	} else {
		return randomStringFromSet(
			"Ryzen 7 PRO 2700U",
			"Ryzen 5 PRO 3500U",
			"Ryzen 3 PRO 3200GE",
		)
	}
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1660-Ti",
			"GTX 1070",
		)
	} else {
		return randomStringFromSet(
			"RX 590",
			"RTX 580",
			"RX 5700-XT",
			"RX Vega-56",
		)
	}
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomStringFromSet(str ...string) string {
	//
	n := len(str)
	if n == 0 {
		return ""
	}
	return str[rand.Intn(n)]
}

func randomBool() bool {
	//
	return rand.Intn(2) == 1
}

func randomPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomResolution() *pb.Screen_Resolution {
	//

	height := randomInt(1080, 4320)
	width := height * 16 / 9
	screen := &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}
	return screen
}

func randomID() string {
	id := uuid.New()
	return id.String()
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "HP", "Dell", "Lenovo")
}

func randomLaptopName(brand string) string {

	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Pro", "MacBook Air")
	case "HP":
		return randomStringFromSet("Pavillion", "Elite Book")
	case "Dell":
		return randomStringFromSet("Latitude", "Vostro", "XPS", "Alienware")
	case "lenovo":
		return randomStringFromSet("ThinkPad X1", "Thinkpad P1", "ThinkPad P53")
	}
	return brand
}
