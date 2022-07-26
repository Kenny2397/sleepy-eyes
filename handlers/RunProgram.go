package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/Kenny2397/visual-programming/server"
)

type PythonCode struct {
	PythonCode string `json:"pythonCode"`
}

func RunProgram(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintln(w, err)
		}
		var python PythonCode
		json.Unmarshal(reqBody, &python)

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
		// fmt.Println(int)

		// execute
		cmd := exec.Command("C:\\Users\\Kenny\\AppData\\Local\\Programs\\Python\\Python39\\python.exe", "./new-script.py")
		out, err := cmd.Output()
		if err != nil {
			fmt.Fprint(w, err)
		}

		fmt.Println("output Code:", string(out))

		w.Header().Set("Content-Type", "application/json")
		// status 200 ok!
		w.WriteHeader(http.StatusOK)
		// fmt.Println(dg)
		// json.NewEncoder(w).Encode(string(out))
		w.Write(out)
	}
}
