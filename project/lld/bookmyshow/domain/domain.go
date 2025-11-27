package domain

/*
Movie Booking System - Low Level Design

You're tasked with designing a movie booking system for a city with multiple multiplexes.
Each multiplex has several screens, each showing different movies at specific times. Each screen has a set of seats available for booking.
Once a seat is selected, it should be temporarily held for 15 minutes while the user completes the booking.
If the user doesn’t finish the booking within 15 minutes, the seat should automatically be marked as available again.


    - User can view all the current showing movies
    - User can select a movie and view the details of the movie(rating, actors)
    - A multiplex will have multiple screens
    - User can view the shows(time slot) in each screen
    - User can select seats of the movie show(seats, screen, slot, multiplex)
    - User can reserve a seat for some window period, and other user have to wait


Classes and methods
    - User(id,name, dob, email, phoneNo)
    - Multiplex(id, name, location:string, []*Screen)
        - NewMultiplex
        - addScreens
    - Screen(id, capacity, name, type: 3D/4D, multiplexID, []*Seat)
        - CreateScreen()
        - GetScreen(type, name, multiplexID)
        - addSeats()
    - Seat(id, type=premium/general, row, number, screen_id)
        - CreateSeat()
        - GetSeat()
    - SeatInventory(id, seat_id, status:HOLD, show_id, holdExpiresAt) - it’s a per-show snapshot of each seat’s booking state
        - createSeatInventory
    - Movie(id, name, rating, description, actors, language)
        - CreateMovie()
        - GetMovieDetails()
    - Show(id, screen_id, movie_id, duration)
        - ListShows(movie_name/rating/language) -> []*Shows
        - CreateShow()
    - Booking(id, show_id, user_id, []seat_ids, payment_status)
        - book() -> id
        - getBookingDetails(id)
    - BookingApp(orchestrator/apis)
        - currentShows() -> []shows
        - addMovie()
        - addMultiplex
        - addSeats
        - addShow
        - createUser
  		- holdSeats()
        - bookTicket()
*/

type User struct {
	userId int
	location map[string]string
}