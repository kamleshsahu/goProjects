package sortingStrategy

import (
	"awesomeProject/lld/appointmentBooking/models"
	"sort"
)

type sortByRating struct {
}

func (s sortByRating) Sort(dl []models.Doctor) []models.Doctor {
	sort.Slice(dl, func(i, j int) bool {
		return dl[i].Rating > dl[j].Rating
	})
	return dl
}

func Default() ISort {
	return &sortByRating{}
}
