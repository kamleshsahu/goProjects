package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'getNumDraws' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER year as parameter.
 */

func getNumDraws(year int32) int32 {
	ans := int32(0)

	for i := 0; i <= 10; i++ {
		resp := parsePage(i, i, int(year), 1)
		ans += int32(resp.Total)
	}
	return int32(ans)
}

type Response struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Data       []struct {
		Competition string `json:"competition"`
		Year        int    `json:"year"`
		Round       string `json:"round"`
		Team1       string `json:"team1"`
		Team2       string `json:"team2"`
		Team1Goals  string `json:"team1goals"`
		Team2Goals  string `json:"team2goals"`
	} `json:"data"`
}

func parsePage(t1g, t2g, year int, page int) Response {
	qp := fmt.Sprintf("year=%d&team1goals=%d&team2goals=%d&page=%d", year, t1g, t2g, page)
	urls := "http://jsonmock.hackerrank.com/api/football_matches?" + qp

	response, err := http.Get(urls)
	resp := Response{}
	if err != nil {
		fmt.Print(err.Error())
		return resp
	}

	data, _ := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
		return resp
	}

	err = json.Unmarshal([]byte(data), &resp)
	if err != nil {
		return Response{}
	}
	fmt.Println(urls)
	fmt.Println(resp)
	return resp
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	year := int32(yearTemp)

	result := getNumDraws(year)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
