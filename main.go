package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := r.ReadString('\n')
	fmt.Println(text)

	csvFile, _ := os.Open("krakenUSDDay.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	dataset := make([][][]float64, 0, 1000)
	//control := make([][][]float64, 0, 100)
	{

		sumts := 0.0
		sumtw := 0.0

		count := 0

		data := make([]float64, 0, 100)

		last := 0.0
		for i := 0; true; i++ {
			line, error := reader.Read()

			if error == io.EOF {
				break
			} else if error != nil {
				log.Fatal(error)
			}

			value, e := strconv.ParseFloat(line[1], 64)
			if i >= 2 && e == nil {
				data = append(data, (value-last)/last)
			}
			last = value
		}

		var multtw float64 = 2 / float64(13)
		var multts float64 = 2 / float64(27)
		var multni float64 = 2 / float64(10)

		sumDev := 0.0
		sumt := 0.0

		for i := 0; i < 26; i++ {
			if i < 20 {
				if i < 12 {
					sumtw += data[i]
				}
				sumt += data[i]
			}
			sumts += data[i]
		}

		var lastematw float64 = sumtw / 12
		var lastemats float64 = sumts / 26
		lastemani := 0.0
		for i := 26; i < 35; i++ {
			sumtw -= data[i-11]
			sumts -= data[i-25]
			sumt -= data[i-19]

			ematw := data[i]*multtw + lastematw*(1-multtw)
			lastematw = ematw
			emats := data[i]*multts + lastemats*(1-multts)
			lastemats = emats
			macd := ematw - emats
			emani := macd*multni + lastemani*(1-multni)
			lastemani = emani
		}
		for i := 35; i < len(data)-1; i++ {
			sumtw -= data[i-11]
			sumts -= data[i-25]
			sumt -= data[i-19]
			for i := 0; i < 20; i++ {
				sumDev += math.Abs(data[i] - (sumt / 20))
			}
			dev := sumDev / 20

			ematw := data[i]*multtw + lastematw*(1-multtw)
			lastematw = ematw
			emats := data[i]*multts + lastemats*(1-multts)
			lastemats = emats
			macd := ematw - emats
			emani := macd*multni + lastemani*(1-multni)
			lastemani = emani

			row := [][]float64{
				{data[i], sumtw / 12, sumts / 26, sumt / 20, ematw, emats, macd, emani - macd, data[0] - (sumt / 20) + 2*dev},
				{-1},
			}

			if i > 35 {
				val := -1.0
				if data[i] > data[i-1] {
					val = 1.0
				}
				dataset[count-1][1][0] = val
			}

			dataset = append(dataset, row)

			count++

			sumtw += data[i]
			sumts += data[i]
			sumt += data[i]

			/*people = append(people, Person{
				Firstname: line[0],
				Lastname:  line[1],
				Address: &Address{
					City:  line[2],
					State: line[3],
				},
			})*/
		}
	}
	//XOR data set
	/*data := [][][]float64{
		{
			{0, 1},
			{1},
		},
		{
			{1, 0},
			{1},
		},
		{
			{0, 0},
			{0},
		},
		{
			{1, 1},
			{0},
		},
	}*/

	var winner Network
	neat := GetNeatInstance(250, 9, 1, .3, .01)
	neat.initialize()

	winner = neat.start(dataset, 20, 1000)
	//neat.printNeat()

	fmt.Println()

	printNetwork(&winner)
	fmt.Println("best ", winner.fitness, "error", 1/winner.fitness)
	fmt.Println("result: ", winner.Process(dataset[0][0]), winner.Process(dataset[1][0]), winner.Process(dataset[2][0]), winner.Process(dataset[3][0])) //1 1 0 0
	fmt.Println("finsihed")
}
