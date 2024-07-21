package main

import (
	"fmt"
	arraysandhashing "neetcode/arraysAndHashing"
	"neetcode/binarysearch"
	dynamicprogramming1d "neetcode/dynamicProgramming1d"
	twopointers "neetcode/twoPointers"
	"os"
	"reflect"
	"strconv"
)

type FunctionWrapper struct {
	fn       interface{}
	argTypes []reflect.Type
}

// convertArgs converts command line arguments to a slice of integers
func convertArgs(args []string, types []reflect.Type) []reflect.Value {
	var convertedArgs []reflect.Value

	argIndex := 0
	for _, typ := range types {
		var convertedValue reflect.Value
		switch typ.Kind() {
		case reflect.Int:
			if argIndex >= len(args) {
				panic("not enough arguments")
			}
			value, err := strconv.Atoi(args[argIndex])
			if err != nil {
				panic(fmt.Sprintf("cannot convert argument to int: %v", args[argIndex]))
			}
			convertedValue = reflect.ValueOf(value)
			argIndex++

		case reflect.String:
			if argIndex >= len(args) {
				panic("not enough arguments")
			}
			convertedValue = reflect.ValueOf(args[argIndex])
			argIndex++

		case reflect.Slice:
			// In case of slice, read all remaining args
			if typ.Elem().Kind() == reflect.Int {
				var nums []int
				for argIndex < len(args) {
					value, err := strconv.Atoi(args[argIndex])
					if err != nil {
						panic(fmt.Sprintf("cannot convert argument to int: %v", args[argIndex]))
					}
					nums = append(nums, value)
					argIndex++
				}
				convertedValue = reflect.ValueOf(nums)
			}
		default:
			panic(fmt.Sprintf("unsupported argument type: %v", typ.Kind()))
		}
		if convertedValue.IsValid() {
			convertedArgs = append(convertedArgs, convertedValue)
		}
	}
	return convertedArgs
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <function> [args...]")
		return
	}

	functionNumber := os.Args[1]
	args := os.Args[2:]

	functions := map[string]FunctionWrapper{
		"15":  {twopointers.ThreeSum, []reflect.Type{reflect.TypeOf([]int{})}},
		"70":  {dynamicprogramming1d.ClimbStairs, []reflect.Type{reflect.TypeOf(0)}},
		"153": {binarysearch.FindMin, []reflect.Type{reflect.TypeOf([]int{})}},
		"198": {dynamicprogramming1d.Rob, []reflect.Type{reflect.TypeOf([]int{})}},
		"217": {arraysandhashing.ContainsDuplicate, []reflect.Type{reflect.TypeOf([]int{})}},
		"238": {arraysandhashing.ProductExceptSelf, []reflect.Type{reflect.TypeOf([]int{})}},
		"242": {arraysandhashing.IsAnagram, []reflect.Type{reflect.TypeOf(""), reflect.TypeOf("")}},
	}

	functionWrapper, exists := functions[functionNumber]
	if !exists {
		fmt.Printf("Function %s not found\n", functionNumber)
		return
	}

	convertedArgs := convertArgs(args, functionWrapper.argTypes)

	results := reflect.ValueOf(functionWrapper.fn).Call(convertedArgs)

	for _, result := range results {
		fmt.Printf("Result: %v\n", result.Interface())
	}
}
