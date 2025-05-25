package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

// CSV Reader liest die CSV und gibt sie als Slice wieder
func readCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("fehler beim Öffnen der CSV-Datei: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("fehler beim Lesen der CSV-Datei: %v", err)
	}

	return records, nil
}

// createEvenColLayout berechnet eine Aufteilung für bis zu 12 Spalten.
// Wenn numCols > 12, sollte man ein anderes Layout-Konzept nutzen, ich glaube aber das geht in Maroto nicht gut
func createEvenColLayout(numCols int) []int {
	layout := make([]int, numCols)

	if numCols == 0 { // gibt eine leeres PDF aus sollte die CSV leer sein
		return layout
	}

	// Sollte die CSV mehr als 12 spalten wieder geben, wird diese automatisch auf 12 zurück gesetzt

	if numCols > 12 {
		numCols = 12
		layout = make([]int, numCols)
	}

	// passt die Grundbreite der PDF auf col 12 an

	baseSize := 12 / numCols          // Bsp. wenn die CSV 5 spalten gibt rechnet baseSize 10 / 5 = 2 und gibt so jeder spalte die basesize von 2
	leftover := 12 - baseSize*numCols // Hier werden die restspalten verteilt. 12 - (2  *  5)  in unserem besipiel. gibt 2 wieder und diese 2 restlichen spalten werden dann aufgeteilt

	// Jede Spalte kriegt eine basis größe, abhängig von der anzahl der spalten

	for i := 0; i < numCols; i++ {
		layout[i] = baseSize //Spalte layout[i] kriegt gleich die spaltengröße 12 / numCols
	}
	// Hier wird von Links nach Rechts die restlichen Spalten verteilt.

	for i := 0; i < leftover; i++ {

		if i == 2 { // Sorgt dafür, dass leftover nur +1 geht wenn MCCMNC übersprungen wird
			leftover++
		}

		if i != 2 { // Überspringt die dritte spalte, was MCCMNC ist und ein Col von 2 einfach zu groß wäre
			layout[i]++ // sind es z.B. 5 spalten werden 5 spalten mit der größe 2 generiert und spalte 1 und 2 werden demnach +1 gesetzt
		}
	}

	return layout
}

// tablePDF baut aus den CSV-Daten eine Tabelle in der PDF.
// dass die 12 Rasterspalten möglichst komplett genutzt werden.

func tablePDF(m pdf.Maroto, records [][]string) {
	if len(records) == 0 {
		log.Println("Keine Daten in der CSV-Datei gefunden!")
		return
	}
	

	for _, row := range records {
		m.Row(6, func() { // Die zahl gibt die höhe der spalten in PT an
			layout := createEvenColLayout(len(row)) // z.B. 5 Spalten -> [3,3,2,2,2] = 12

			for i, col := range row {

				if col == "PASS" || col == "PASS +" || col == "PASS -" || col == "PASS x" || col == "MANUAL PASSED" {
					m.SetBackgroundColor(greenColorForPASS())
				} else if col == "Blocked" || col == "No 2G/3G" || col == "No SIGOS support" || col == "x" || col == "RLU Fail" || col == "Not in Orange Hub" || strings.Contains(col, "Pool") {
					m.SetBackgroundColor(whiteForBlockNoSig())
				} else if col == "FAIL" || col == "FAIL +" || col == "FAIL -" || col == "FAIL x" || col == "TEF PASSED" {
					m.SetBackgroundColor(redForFail())
				} else {
					// Setze eine Standardfarbe für alle anderen Zellen
					m.SetBackgroundColor(color.Color{Red: 68, Green: 114, Blue: 196})
				}

				// Falls row mehr Spalten hat als 12, wird layout kleiner sein als row.
				// D.h. wir ignorieren zusätzliche Spalten (oder man bricht hier ab).
				if i >= len(layout) {
					break
				}
				colSize := layout[i]

				m.Col(uint(colSize), func() {
					m.SetBorder(true)

					// Die IF's sollen dafür sorgen, dass die Farbe der schrift die selbe ist wie der Hintergrund
					// Bei denen wo kein Text in der Zelle sein soll. Konnte noch keine bessere Lösung finden

					if col == "PASS" || col == "PASS +" || col == "PASS -" || col == "PASS x" || col == "MANUAL PASSED" {
						m.Text(col, props.Text{
							Color: greenColorForPASS(),
						})
					} else if col == "FAIL" || col == "FAIL +" || col == "FAIL -" || col == "FAIL x" || col == "TEF PASSED" {
						m.Text(col, props.Text{
							Color: redForFail(),
						})
					} else if col == "Blocked" || col == "No 2G/3G" || col == "No SIGOS support" || col == "RLU Fail"  ||strings.Contains(col, "Pool") {
						m.Text(col, props.Text{
							Size:  6,
							Style: consts.Bold,
							Align: consts.Center,
							Top:   1,
							Color: color.NewBlack(),
						})
					} else {
						m.Text(col, props.Text{
							Size:  6,
							Style: consts.Bold,
							Align: consts.Center,
							Top:   1,

							Color: color.NewWhite(),
						})

					}
				})
			}

		})
	}
}

// Funktionen für die Fabgebungen die in anderen Funktionen verwendet werden

func greenColorForPASS() color.Color {
	return color.Color{
		Red:   146,
		Green: 208,
		Blue:  80,
	}
}

func redForFail() color.Color {
	return color.Color{
		Red:   255,
		Green: 0,
		Blue:  0,
	}
}

func whiteForBlockNoSig() color.Color {
	return color.Color{
		Red:   255,
		Green: 255,
		Blue:  255,
	}
}

func RemoveLastNLines(filename string, n int) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Prüfen, ob die letzten n Zeilen mit den bestimmten Wörtern beginnen
	count := 0
	for i := len(lines) - 1; i >= 0 && count < n; i-- {
		if strings.HasPrefix(lines[i], "Statistics") ||
			strings.HasPrefix(lines[i], "Global Roamer") ||
			strings.HasPrefix(lines[i], "Segron") ||
			len(strings.TrimSpace(lines[i])) == 0 {
			count++
		}
	}

	// Falls weniger als n relevante Zeilen gefunden wurden, nichts ändern
	if count < n {
		return nil
	}

	// Datei überschreiben und nur die relevanten Zeilen behalten
	file, err = os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < len(lines)-n; i++ {
		_, err := writer.WriteString(lines[i] + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func main() {

	// Pfade anpassen an deine Umgebung
	outputPath := "C:/Users/Marcel Helmer/Desktop/testpdf/testPDF_maroto.pdf"
	csvFilePath := "C:/Users/Marcel Helmer/Desktop/testpdf/TestCSV.csv"

	// Maroto-PDF erstellen
	m := pdf.NewMaroto(consts.Landscape, consts.A4)
	m.SetPageMargins(10, 0, 10)

	//CSV Normalisieren

	err := RemoveLastNLines(csvFilePath, 4)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	// CSV laden
	records, err := readCSV(csvFilePath)
	if err != nil {
		log.Fatalf("Fehler beim Lesen der CSV-Datei: %v", err)
	}

	// Tabelle in die PDF einfügen
	tablePDF(m, records)

	// PDF speichern
	err = m.OutputFileAndClose(outputPath)
	if err != nil {
		log.Fatalf("Fehler beim Speichern der PDF-Datei: %v", err)
	}

	fmt.Println("PDF wurde erfolgreich erstellt!")
}
