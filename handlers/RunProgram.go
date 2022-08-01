package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/Kenny2397/visual-programming/server"
)

type PythonCode struct {
	PythonCode string `json:"pythonCode"`
	Output     string
}

// var python PythonCode

// post
func RunProgram(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		python := PythonCode{}

		err := json.NewDecoder(r.Body).Decode(&python)
		if err != nil {
			panic(err)
		}

		osFile, err := os.Create("new-script.py")
		if err != nil {
			fmt.Println(err)
		}
		defer osFile.Close()
		// fmt.Println(osFile)

		len, err := osFile.WriteString(python.PythonCode)
		if err != nil {
			fmt.Println(err, len)
		}

		// execute
		cmd := exec.Command("C:\\Users\\Kenny\\AppData\\Local\\Programs\\Python\\Python39\\python.exe", "./new-script.py")
		out, err := cmd.Output()
		if err != nil {
			fmt.Fprint(w, err)
		}

		fmt.Println("output Code:", string(out))

		python.Output = string(out)

		pythonJson, err := json.Marshal(python)
		if err != nil {
			panic(err)
		}

		w.Header().Set("context-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(pythonJson)
		// reqBody, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	fmt.Fprintln(w, err)
		// }

		// json.Unmarshal(reqBody, &python)

		// // ------------------------------------
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// w.Write([]byte("ok"))
	}
}
