// Originally from github.com/Benchkram/errz however I needed to change the code a bit to suit this application
// Package errz provides highlevel error handling helpers
//
//
package main

import (
	"log"
	"os"

	"github.com/pkg/errors"
)

func initErrLogger() *log.Logger {
	f, err := os.OpenFile(errorLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return log.New(f, "Error : ", log.LstdFlags)
}

//Fatal panics on error
//First parameter of msgs is used each following variadic arg is dropped
func (a App) Fatal(err error, msgs ...string) {
	if err != nil {
		var str string
		for _, msg := range msgs {
			str = msg
			break
		}
		panic(errors.Wrap(err, str))
	}
}

// Recover recovers a panic introduced by Fatal, any other function which calls panics
// or a memory corruption. Logs the error when called without args.
//
//Must be used at the top of the function deferred
//defer Recover()
//or
//defer Recover(&err)
func (a App) Recover(logger *log.Logger, errs ...*error) {
	var e *error
	for _, err := range errs {
		e = err
		break
	}

	//handle panic
	if r := recover(); r != nil {
		var errmsg error
		//Preserve error which might have happened before panic/recover
		//Check if a error ptr was passed + a error occurred
		if e != nil && *e != nil {
			//When error occurred before panic then Wrap panic error around it
			errmsg = errors.Wrap(*e, r.(error).Error())
		} else {
			//No error occurred just add a stacktrace
			errmsg = errors.Wrap(r.(error), "")
		}

		//If error cant't bubble up -> Log it
		if e != nil {
			*e = errmsg
		} else {
			a.ErrLogger.Printf("%+v", errmsg)
		}
	}
}

//Log logs an error + stack trace directly to console or file
//Use this at the top level to publish errors without creating a new error object
func (a App) Log(err error, msgs ...string) {
	if err != nil {
		var str string
		for _, msg := range msgs {
			str = msg
			break
		}
		a.ErrLogger.Printf("%+v", errors.Wrap(err, str))
	}
}

func (a App) LogFatal(err error, msgs ...string) {
	a.Log(err, msgs...)
	a.Fatal(err, msgs...)
}
