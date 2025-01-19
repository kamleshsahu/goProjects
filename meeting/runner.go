package meeting

import (
	"awesomeProject/meeting/service"
	"fmt"
	"time"
)

func Run() {

	bs := service.NewBookingService(2, nil)

	meetingId, roomId, err := bs.Book(time.Now(), time.Now().Add(time.Minute*30))

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(meetingId, roomId)

	userEmails := []string{"kamlesh@gmail.com"}

	err = bs.InviteUser(meetingId, userEmails)
	if err != nil {
		fmt.Println(err)
	}
}
