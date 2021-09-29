package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime int = 40
// TODO: define the 'RemainingOvenTime()' function
func RemainingOvenTime (time int) int {
    return OvenTime - time
}
// TODO: define the 'PreparationTime()' function
func PreparationTime (time int) int {
    return 2*time
}
// TODO: define the 'ElapsedTime()' function
func ElapsedTime (layer int, time int) int {
	return 2*layer + time
}