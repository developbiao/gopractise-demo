package main
import "fmt"

func main() {

    const lightSpeed = 299792
    const secondsPerDay = 86400

    var distance int64 = 41.3e12
    fmt.Println("Apha Centauri is", distance, "km away.")

    days := distance / lightSpeed / secondsPerDay
    years := days / 365
    fmt.Println("Thant is ", days, "days of travel at light speed.")
    fmt.Println("That is ", years, "years of travel at light speed.")
}

