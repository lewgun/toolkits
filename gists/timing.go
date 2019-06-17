//https://stackoverflow.com/questions/45766572/is-there-an-efficient-way-to-calculate-execution-time-in-golang
func SomeFunction(list *[]string) {
    defer TimeTrack(time.Now())
    // Do whatever you want.
}

func TimeTrack(start time.Time) {
    elapsed := time.Since(start)

    // Skip this function, and fetch the PC and file for its parent.
    pc, _, _, _ := runtime.Caller(1)

    // Retrieve a function object this functions parent.
    funcObj := runtime.FuncForPC(pc)

    // Regex to extract just the function name (and not the module path).
    runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
    name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")

    log.Println(fmt.Sprintf("%s took %s", name, elapsed))
}
