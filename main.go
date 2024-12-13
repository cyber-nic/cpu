package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/cpu"
)

type CPUInfo struct {
	name           string
	speed          float64
	cores          int
	threadsPerCore int
}

func main() {
	cpu := getCPUInfo()
	drawCPU(cpu)
}

func getCPUInfo() CPUInfo {
	cpuInfo, err := cpu.Info()
	if err != nil || len(cpuInfo) == 0 {
		log.Fatalf("Failed to fetch CPU info: %v", err)
	}

	// Logical CPU count
	logicalCPUs := runtime.NumCPU()

	// Count unique cores using the CoreID field
	coreMap := make(map[string]bool)
	for _, ci := range cpuInfo {
		coreMap[ci.CoreID] = true
	}
	physicalCores := len(coreMap)

	// Calculate threads per core
	var threadsPerCore int
	if physicalCores > 0 {
		threadsPerCore = logicalCPUs / physicalCores
	}

	return CPUInfo{
		name:           cpuInfo[0].ModelName,
		speed:          cpuInfo[0].Mhz,
		cores:          physicalCores,
		threadsPerCore: threadsPerCore,
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func drawCPU(cpu CPUInfo) {
	// Constants for sizing
	cpuWidth := 49
	coreWidth := 22
	threadWidth := 16

	// Create CPU outer box
	fmt.Println(strings.Repeat(" ", (cpuWidth-cpuWidth)/2) + "┌" + strings.Repeat("─", cpuWidth-2) + "┐")
	fmt.Println(strings.Repeat(" ", (cpuWidth-cpuWidth)/2) + "│" + centerText(cpu.name, cpuWidth-2) + "│")
	fmt.Println(strings.Repeat(" ", (cpuWidth-cpuWidth)/2) + "│" + centerText(fmt.Sprintf("%2.fMHz", cpu.speed), cpuWidth-2) + "│")

	// Dynamic drawing for cores based on total cores and threads
	rows := (cpu.cores + 1) / 2 // Calculate rows required (2 cores per row)
	for row := 0; row < rows; row++ {
		coresInRow := min(cpu.cores-row*2, 2) // Handle remaining cores
		drawCoreRow(row, coresInRow, cpu.threadsPerCore, coreWidth, threadWidth, cpuWidth)
	}

	// Close CPU box
	fmt.Println(strings.Repeat(" ", (cpuWidth-cpuWidth)/2) + "└" + strings.Repeat("─", cpuWidth-2) + "┘")
}

func drawCoreRow(row, numCores, threadsPerCore, coreWidth, threadWidth, cpuWidth int) {
	coreBoxes := make([][]string, numCores)
	maxHeight := 0

	// Generate core boxes
	for i := 0; i < numCores; i++ {
		coreNum := row*2 + i
		coreBoxes[i] = createCoreBox(coreNum, threadsPerCore, coreWidth, threadWidth)
		if len(coreBoxes[i]) > maxHeight {
			maxHeight = len(coreBoxes[i])
		}
	}

	// Print core boxes side by side
	for lineNum := 0; lineNum < maxHeight; lineNum++ {
		line := "│"
		for i := 0; i < numCores; i++ {
			if lineNum < len(coreBoxes[i]) {
				if i == 0 {
					line += " "
				}
				line += coreBoxes[i][lineNum]
				if i < numCores-1 {
					line += " "
				}
			}
		}
		// Pad to full CPU width
		for len(line) < cpuWidth-1 {
			line += " "
		}
		line += " │"
		fmt.Println(line)
	}
}

func createCoreBox(coreNum, threadsPerCore, width, threadWidth int) []string {
	var box []string

	// Core top
	box = append(box, fmt.Sprintf("┌%s┐", strings.Repeat("─", width-2)))
	box = append(box, fmt.Sprintf("│%s│", centerText(fmt.Sprintf("Core %d", coreNum), width-2)))

	// For each thread in the core
	for t := 0; t < threadsPerCore; t++ {
		threadNum := coreNum*threadsPerCore + t
		threadBox := createThreadBox(threadNum, threadWidth)

		// Add padding to thread box
		paddedThreadBox := make([]string, len(threadBox))
		for i, line := range threadBox {
			paddedSpace := (width - threadWidth - 2) / 2
			padding := strings.Repeat(" ", paddedSpace)
			paddedThreadBox[i] = "│" + padding + line + padding + "│"
		}
		box = append(box, paddedThreadBox...)
	}

	// Core bottom
	box = append(box, fmt.Sprintf("└%s┘", strings.Repeat("─", width-2)))
	return box
}

func createThreadBox(threadNum, width int) []string {
	var box []string

	// Thread top
	box = append(box, fmt.Sprintf("┌%s┐", strings.Repeat("─", width-2)))
	box = append(box, fmt.Sprintf("│%s│", centerText(fmt.Sprintf("Thread %d", threadNum), width-2)))

	// Thread content
	box = append(box, fmt.Sprintf("│%s│", strings.Repeat("█", width-2)))

	// Thread bottom
	box = append(box, fmt.Sprintf("└%s┘", strings.Repeat("─", width-2)))
	return box
}

func centerText(text string, width int) string {
	textLen := len(text)
	if textLen >= width {
		return text
	}
	leftPad := (width - textLen) / 2
	rightPad := width - textLen - leftPad
	return strings.Repeat(" ", leftPad) + text + strings.Repeat(" ", rightPad)
}