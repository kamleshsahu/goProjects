package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"sync"
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
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Data       []Data `json:"data"`
}

type Data struct {
	Name              string  `json:"name"`
	RuntimeOfSeries   string  `json:"runtime_of_series"`
	Certificate       string  `json:"certificate"`
	RuntimeOfEpisodes string  `json:"runtime_of_episodes"`
	Genre             string  `json:"genre"`
	ImdbRating        float64 `json:"imdb_rating"`
	Overview          string  `json:"overview"`
	NoOfVotes         int     `json:"no_of_votes"`
	ID                int     `json:"id"`
}

func getTotalGoals(searchGenre string) string {

	resp := parsePage(1)
	totalPages := resp.TotalPages

	allMovies := make([]Data, 0)
	mutex := sync.Mutex{}
	wait := sync.WaitGroup{}

	for page := 1; page <= totalPages; page++ {
		wait.Add(1)
		go func() {
			resp = parsePage(page)

			for _, movie := range resp.Data {
				genres := strings.Split(movie.Genre, ",")
				found := false
				for _, genre := range genres {
					if strings.EqualFold(strings.TrimSpace(genre), searchGenre) {
						found = true
						break
					}
				}
				if found {
					mutex.Lock() // Lock the mutex before appending to allMovies
					allMovies = append(allMovies, movie)
					mutex.Unlock() // Unlock the mutex after appending
				}
			}
			wait.Done()
		}()
	}
	wait.Wait()
	sort.Slice(allMovies, func(i, j int) bool {
		if allMovies[i].ImdbRating == allMovies[j].ImdbRating {
			return allMovies[i].Name < allMovies[j].Name
		}
		return allMovies[i].ImdbRating > allMovies[j].ImdbRating
	})

	if len(allMovies) == 0 {
		fmt.Println("some unexpected error")
	}

	return allMovies[0].Name
}

func parsePage(page int) Response {
	qp := fmt.Sprintf("page=%d", page)
	urls := "http://jsonmock.hackerrank.com/api/tvseries?" + qp

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
	//fmt.Println(urls)
	//fmt.Println(resp)
	return resp
}

func main() {

	Genre := "Adventure"
	fmt.Println(getTotalGoals(Genre))

	Genre1 := "Drama"
	fmt.Println(getTotalGoals(Genre1))

	Genre2 := "Horror"
	fmt.Println(getTotalGoals(Genre2))
	//reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	//
	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)
	//
	//defer stdout.Close()
	//
	//writer := bufio.NewWriterSize(stdout, 16*1024*1024)
	//
	//team := readLine(reader)
	//
	//yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//checkError(err)
	//year := int32(yearTemp)
	//
	//result := getTotalGoals(team, year)
	//
	//fmt.Fprintf(writer, "%d\n", result)
	//
	//writer.Flush()
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
