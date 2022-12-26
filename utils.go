package main

import (
    "fmt"
    "runtime"
    "log"
    "os"
    "golang.org/x/crypto/bcrypt"
)

// type JsonReturnStruct struct {
type SuccessStruct struct {
	// ErrorCode int
	SessionCode 					string 		`json:"session_code"`
	SessionCreatedDatetime 			string 		`json:"session_created_datetime"`
	Message 						string 		`json:"message"`
}

type ErrorStruct struct {
	ErrorCode 			int 		`json:"session_code"`
	Message 			string 		`json:"message"`
	Error 				error 		`json:"error"`
}

func (a ErrorStruct) String() string {
	return fmt.Sprintf("Error Code = %q, msg = %s.", a.ErrorCode, a.Message)
}
type ErrorInterface interface {
	String() string
}

func JSONSuccessReturn(JsonReturn SuccessStruct) SuccessStruct {
	if JsonReturn.Message != "" {
		JsonReturn.Message = "Success"
	}	
	return JsonReturn
}

func JSONErrorReturn(custom_err ErrorStruct) ErrorStruct {
	logger(custom_err.Error)
	msg := "Error Found"
	error_code := 250

	if custom_err.Message != "" {
		msg = custom_err.Message
	}
	if custom_err.ErrorCode != 0 {
		error_code = custom_err.ErrorCode
	}
	es := ErrorStruct{
		ErrorCode: error_code,
		Message: msg,
	}
	return es
}



func logger(msg interface{}) {
    // pc, fn, line, _ := runtime.Caller(1)
    _, fn, line, _ := runtime.Caller(1)
	f, err := os.OpenFile("errorlog.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
	    log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// log.Println("This is a test log entry abcdef")
	log.Println(msg)
    log.Printf("[ERROR] %s [%s:%d]", err, fn, line)
}


func loggerError(err error) {
	logger(err)
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}



func Contains(sl []string, name string) bool {
   for _, value := range sl {
      if value == name {
         return true
      }
   }
   return false
}



func isset(arr []string, index int) bool {
    return (len(arr) > index)
}

