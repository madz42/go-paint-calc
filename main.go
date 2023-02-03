package main

import (
	"fmt"
	"math"
	"os"
	"painter-calc/utils"
)

func main() {
	fmt.Println("PAINTING COSTS CALCULATOR\n")
	// enter rooms count
	var roomsCount int
	for {
		val, err := utils.GetDataInt("How many rooms do you want to paint? (0-10)")
		if err == nil && val >= 0 && val <= 10 {
			if val == 0 {
				fmt.Println("Bye!")
				os.Exit(0)
			}
			roomsCount = int(val)
			fmt.Println("Rooms:", roomsCount)
			break
		}
		fmt.Println("Check your input data")
	}
	var roomsSurface, wallsFinish []float64
	var layers []int
	var paintType []string
	paints := map[string][]float64{
		// name: price, usage
		"acrylic":  {5.99, 0.15},
		"emulsion": {3.19, 0.1},
	}

	// enter rooms data
	for i := 1; i <= roomsCount; i++ {
		var wallL, wallH float64
		var apertures int
		fmt.Printf("\nStart collecting info for Room #%d\n\n", i)
		//room length
		for {
			fmt.Printf("[ROOM %d/%d] ", i, roomsCount)
			val, err := utils.GetDataFloat("Enter total length of walls (in meters)")
			if err == nil && val > 0 && val <= 100 {
				wallL = val
				fmt.Println("Length:", wallL)
				break
			}
			fmt.Println("Check your input data")
		}
		//room height
		for {
			fmt.Printf("[ROOM %d/%d] ", i, roomsCount)
			val, err := utils.GetDataFloat("Enter height of walls (in meters)")
			if err == nil && val > 0 && val <= 10 {
				wallH = val
				fmt.Println("Height:", wallH)
				break
			}
			fmt.Println("Check your input data")
		}
		//room surface calc
		roomsSurface = append(roomsSurface, wallL*wallH)

		//get number of doors and windows
		for {
			fmt.Printf("[ROOM %d/%d] ", i, roomsCount)
			val, err := utils.GetDataInt("Enter number of doors/windows (1-100)")
			if err == nil && val > 0 && val <= 100 {
				apertures = int(val)
				fmt.Println("Apertures:", apertures)
				break
			}
			fmt.Println("Check your input data")
		}
		//doors and windows loop
		for j := 1; j <= apertures; j++ {
			var apertureH, apertureW float64
			//width loop
			for {
				fmt.Printf("[ROOM %d/%d][Aperture %d/%d] ", i, roomsCount, j, apertures)
				val, err := utils.GetDataFloat("Enter aperture width (in meters)")
				if err == nil && val > 0 && val <= 10 {
					apertureW = val
					fmt.Println("Width:", apertureW)
					break
				}
				fmt.Println("Check your input data")
			}
			for {
				fmt.Printf("[ROOM %d/%d][Aperture %d/%d] ", i, roomsCount, j, apertures)
				val, err := utils.GetDataFloat("Enter aperture height (in meters)")
				if err == nil && val > 0 && val <= 10 {
					apertureH = val
					fmt.Println("Height:", apertureH)
					break
				}
				fmt.Println("Check your input data")
			}
			roomsSurface[i-1] -= apertureH * apertureW
		}
		// fmt.Println(roomsSurface)
		// paint loop
		for {
			fmt.Printf("[ROOM %d/%d] Choose paint type:\n", i, roomsCount)
			for name, nums := range paints {
				fmt.Printf(" - '%s' paint (%.2fL per sqm, %.2f€ per L\n", name, nums[1], nums[0])
			}
			val, err := utils.GetDataString("What paint do you want? (type name)")
			if err == nil {
				_, ok := paints[val]
				if ok {
					// fmt.Println(chkPaint)
					paintType = append(paintType, val)
					break
				}
			}
			fmt.Println("Check your input data")
		}
		//finish loop
		for {
			fmt.Printf("[ROOM %d/%d] Choose current walls finish:\n", i, roomsCount)
			fmt.Println("0 - smooth \n... \n10 - 10mm diameter bumps")
			val, err := utils.GetDataInt("Enter walls finish (0-10)")
			if err == nil && val >= 0 && val <= 10 {
				wallsFinish = append(wallsFinish, float64(val)/10+1)
				fmt.Printf("Bumps: %dmm\n", val)
				break
			}
			fmt.Println("Check your input data")
		}
		//layers loop
		for {
			fmt.Printf("[ROOM %d/%d] ", i, roomsCount)
			val, err := utils.GetDataInt("How many layers of paint is needed (1-5)")
			if err == nil && val >= 1 && val <= 5 {
				layers = append(layers, int(val))
				fmt.Println("Layers:", val)
				break
			}
			fmt.Println("Check your input data")
		}

	}
	//ask for painters
	var workers int
	for {
		fmt.Println("\nSERVICE: ")
		fmt.Println("  0 - You will do it by your self")
		fmt.Println("  1 - Hire our 'BrushCo Ltd' painters")
		fmt.Println("  2 - Get our highly trained monkeys")
		val, err := utils.GetDataInt("Choose who is going to paint your house (0-2)")
		if err == nil && val >= 0 && val <= 2 {
			workers = int(val)
			break
		}
		fmt.Println("Check your input data")
	}

	//perform calculations
	var totalSurface, totalPaint []float64

	fmt.Println("\nHOW MUCH WILL IT COST")
	for i := 0; i < roomsCount; i++ {
		fmt.Printf("\nROOM #%d\n", i+1)
		if roomsSurface[i] <= 0 {
			roomsSurface[i] = 0
		}
		resSurface := roomsSurface[i] * float64(layers[i])
		totalSurface = append(totalSurface, resSurface)
		fmt.Printf("Walls' surface: %.2f sqm\n", roomsSurface[i])
		fmt.Printf("Total surface: %.2f sqm (%d layers of %s paint)\n", totalSurface[i], layers[i], paintType[i])
		paintUsage := paints[paintType[i]][1] * wallsFinish[i]
		fmt.Printf("Paint usage per sqm: %.2fL\n", paintUsage)
		paintForRoom := resSurface * paintUsage
		fmt.Printf("Paint needed: %.2fL\n", paintForRoom)
		totalPaint = append(totalPaint, paintForRoom*paints[paintType[i]][0])
		fmt.Printf("Paint costs: %.2f€\n", totalPaint[i])
	}

	var surfaceToPaint float64
	for _, x := range totalSurface {
		surfaceToPaint += x
	}
	var paintCost float64
	for _, x := range totalPaint {
		paintCost += x
	}
	const workersPrice float64 = 4.25
	const monkeysPrice float64 = 0.7
	bananas := int(math.Floor(surfaceToPaint * monkeysPrice))
	switch workers {
	case 1:
		fmt.Println("\nSERVICE")
		fmt.Printf("\n'BrushCo Ltd' painters: %.2f€\n", surfaceToPaint*workersPrice)

	case 2:
		fmt.Println("\nSERVICE")
		fmt.Printf("\nHighly trained monkeys: %d bananas\n", bananas)
	}
	// fmt.Println(totalSurface)
	// fmt.Println(totalPaint)

	fmt.Println("\n\n  ===  TOTAL  ===")
	fmt.Printf("\nPaint: %.2f€\n", paintCost)
	switch workers {
	case 0:
		fmt.Println("\nService: 0€")
		fmt.Printf("\nTotal: %.2f€\n", paintCost)
	case 1:
		fmt.Printf("\nService: %.2f€\n\n", surfaceToPaint*workersPrice)
		fmt.Printf("\nTotal: %.2f€\n", paintCost+surfaceToPaint*workersPrice)
	case 2:
		fmt.Printf("\nService: %d bananas\n\n", bananas)
		fmt.Printf("\nTotal: %.2f€ + %d bananas\n", paintCost, bananas)
	}
}
