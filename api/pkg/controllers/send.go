package controllers

import (
	"encoding/json"
	"fmt"
	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	models2 "go_telegram_api/app/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func SendMessage(w http.ResponseWriter, r *http.Request) {
	chatId, err := strconv.ParseInt(r.FormValue("chat_id"), 10, 64)
	var msg = tgBotAPI.NewMessage(chatId, r.FormValue("text"))

	resp, err := models2.TgAPI.Send(msg)

	if err != nil {
		fmt.Println("Error sending message")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(resp)
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
	}
}

func SendPhoto(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024 << 20)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	downFile, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error retrieving data from form-data")
		log.Println(err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	fileBytes, err := ioutil.ReadAll(downFile)

	chatId, _ := strconv.ParseInt(r.FormValue("chat_id"), 0, 0)

	var msg = tgBotAPI.NewPhoto(chatId, tgBotAPI.FileBytes{
		Name:  handler.Filename,
		Bytes: fileBytes,
	})
	msg.Caption = r.FormValue("caption")

	rsp, err := models2.TgAPI.Send(msg)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		b, _ := json.Marshal(err)
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
		return
	} else {
		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(rsp)
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
	}
}

func SendVideo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024 << 20)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	downFile, handler, err := r.FormFile("video")
	if err != nil {
		fmt.Println("Error retrieving data from form-data")
		log.Println(err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	fileBytes, err := ioutil.ReadAll(downFile)

	chatId, _ := strconv.ParseInt(r.FormValue("chat_id"), 0, 0)

	var msg = tgBotAPI.NewVideo(chatId, tgBotAPI.FileBytes{
		Name:  handler.Filename,
		Bytes: fileBytes,
	})
	msg.Caption = r.FormValue("caption")

	rsp, err := models2.TgAPI.Send(msg)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		b, _ := json.Marshal(err)
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
		return
	} else {
		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(rsp)
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
	}
}

func SendSticker(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024 << 20)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	downFile, handler, err := r.FormFile("sticker")
	if err != nil {
		fmt.Println("Error retrieving data from form-data")
		log.Println(err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	fileBytes, err := ioutil.ReadAll(downFile)

	chatId, _ := strconv.ParseInt(r.FormValue("chat_id"), 0, 0)

	var msg = tgBotAPI.NewSticker(chatId, tgBotAPI.FileBytes{
		Name:  handler.Filename,
		Bytes: fileBytes,
	})

	rsp, err := models2.TgAPI.Send(msg)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		b, _ := json.Marshal(err)
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
		return
	} else {
		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(rsp)
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
	}
}

func SendDocument(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024 << 20)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	downFile, handler, err := r.FormFile("document")
	if err != nil {
		fmt.Println("Error retrieving data from form-data")
		log.Println(err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	fileBytes, err := ioutil.ReadAll(downFile)

	chatId, _ := strconv.ParseInt(r.FormValue("chat_id"), 0, 0)

	var msg = tgBotAPI.NewDocument(chatId, tgBotAPI.FileBytes{
		Name:  handler.Filename,
		Bytes: fileBytes,
	})
	msg.Caption = r.FormValue("caption")

	rsp, err := models2.TgAPI.Send(msg)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotModified)
		b, _ := json.Marshal(err)
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
		return
	} else {
		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(rsp)
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
	}
}

func SendAudio(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024 << 20)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	downFile, handler, err := r.FormFile("audio")
	if err != nil {
		fmt.Println("Error retrieving data from form-data")
		log.Println(err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	fileBytes, err := ioutil.ReadAll(downFile)

	chatId, _ := strconv.ParseInt(r.FormValue("chat_id"), 0, 0)

	var msg = tgBotAPI.NewAudio(chatId, tgBotAPI.FileBytes{
		Name:  handler.Filename,
		Bytes: fileBytes,
	})

	rsp, err := models2.TgAPI.Send(msg)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotModified)
		b, _ := json.Marshal(err)
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
		return
	} else {
		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(rsp)
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
	}
}

func SendContact(w http.ResponseWriter, r *http.Request) {

}
