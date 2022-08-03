package sample

import (
	"grpc_project/pb"

	"github.com/golang/protobuf/ptypes"
)

func NewKeyboard() *pb.Keyboard {
	//
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

//return New CPU object
func NewCPU() *pb.CPU {
	//
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numberOfCores := randomInt(2, 8)
	numberOfThreads := randomInt(numberOfCores, 12)
	min_Ghz := randomFloat64(2.0, 3.5)
	max_Ghz := randomFloat64(min_Ghz, 5.0)

	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberOfCores),
		NumberThreads: uint32(numberOfThreads),
		MinGhz:        min_Ghz,
		MaxGhz:        max_Ghz,
	}
	return cpu
}

// return New GPU object
func NewGPU() *pb.GPU {
	//
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	min_Ghz := randomFloat64(2.0, 3.5)
	max_Ghz := randomFloat64(min_Ghz, 5.0)
	memory := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}

	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: min_Ghz,
		MaxGhz: max_Ghz,
		Memory: memory,
	}

	return gpu
}

//return memory object

func NewRAM() *pb.Memory {
	//
	memory := &pb.Memory{
		Value: uint64(randomInt(4, 64)),
		Unit:  pb.Memory_GIGABYTE,
	}

	return memory
}

//return SSD object

func NewSSD() *pb.Storage {
	//
	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	return ssd
}

//return HDD object
func NewHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}

	return hdd
}

//return Screen object
func NewScreen() *pb.Screen {
	//

	panel := randomPanel()
	resolutio := randomResolution()

	screen := &pb.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: resolutio,
		Panel:      panel,
		Multitouch: randomBool(),
	}

	return screen
}

//return laptop object

func NewLaptop() *pb.Laptop {
	//
	id := randomID()
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &pb.Laptop{
		Id:       id,
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRAM(),
		Gpus:     []*pb.GPU{NewGPU()},
		Storages: []*pb.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: float64(randomFloat32(1.0, 3.0)),
		},
		PriceUsd:    randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdatedAt:   ptypes.TimestampNow(),
	}
	return laptop
}
