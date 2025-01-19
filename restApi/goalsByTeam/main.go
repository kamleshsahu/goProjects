package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'getTotalGoals' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING team
 *  2. INTEGER year
 */

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

func getTotalGoals(team string, year int32) int32 {
	ans := 0

	resp := parsePage(team, int(year), 1, 1)
	totalPages := resp.TotalPages

	for page := 1; page <= totalPages; page++ {
		resp := parsePage(team, int(year), 1, page)

		for _, d := range resp.Data {
			val, _ := strconv.Atoi(d.Team1Goals)
			ans += val
		}

	}

	resp = parsePage(team, int(year), 2, 1)
	totalPages = resp.TotalPages

	for page := 1; page <= totalPages; page++ {
		resp := parsePage(team, int(year), 2, page)

		for _, d := range resp.Data {
			val, _ := strconv.Atoi(d.Team2Goals)
			ans += val
		}

	}

	return int32(ans)
}

func parsePage(teamName string, year int, team int, page int) Response {
	qp := fmt.Sprintf("year=%d&team%d=%s&page=%d", year, team, url.QueryEscape(teamName), page)
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

	team := readLine(reader)

	yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	year := int32(yearTemp)

	result := getTotalGoals(team, year)

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
