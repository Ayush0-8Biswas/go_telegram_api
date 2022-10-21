package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sunshineplan/imgconv"
	models2 "go_telegram_api/app/pkg/models"
	"image/png"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func SendMessage(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error while parsing form.")
		w.WriteHeader(http.StatusNotAcceptable)
		return err
	}

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
			return err
		}
	}
	return nil
}

func SendPhoto(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseMultipartForm(1024 << 20)
	if err != nil {
		fmt.Println("Error while parsing form.")
		w.WriteHeader(http.StatusNotAcceptable)
		return err
	}

	downFile := r.MultipartForm.File["image"][0]
	file, err := downFile.Open()
	defer file.Close()
	if err != nil {
		fmt.Println("Error opening file")
		w.WriteHeader(http.StatusNotAcceptable)
		return err
	}

	chatId, _ := strconv.ParseInt(r.MultipartForm.Value["chat_id"][0], 0, 0)

	var msg = tgBotAPI.NewPhoto(chatId, tgBotAPI.FileReader{
		Name:   downFile.Filename,
		Reader: file,
	})
	msg.Caption = r.FormValue("caption")

	rsp, err := models2.TgAPI.Send(msg)
	if err != nil {
		fmt.Println("Error sending message.")
		w.WriteHeader(http.StatusInternalServerError)
		b, _ := json.Marshal(err)
		w.Write(b)
	} else {
		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(rsp)
		w.Write(b)
	}
	return err
}

func SendVideo(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseMultipartForm(1024 << 20)
	if err != nil {
		fmt.Println("Error while parsing form.")
		w.WriteHeader(http.StatusNotAcceptable)
		return err
	}

	downFile := r.MultipartForm.File["video"][0]
	file, err := downFile.Open()
	if err != nil {
		fmt.Println("Error opening file.")
		return err
	}

	chatId, _ := strconv.ParseInt(r.FormValue("chat_id"), 0, 0)

	var msg = tgBotAPI.NewVideo(chatId, tgBotAPI.FileReader{
		Name:   downFile.Filename,
		Reader: file,
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
		return err
	} else {
		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(rsp)
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

func SendDocument(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseMultipartForm(1024 << 20)
	if err != nil {
		fmt.Println("Error while parsing form.")
		w.WriteHeader(http.StatusNotAcceptable)
		return err
	}

	downFile := r.MultipartForm.File["document"][0]
	file, err := downFile.Open()
	if err != nil {
		fmt.Println("Error opening file.")
		return err
	}

	chatId, _ := strconv.ParseInt(r.FormValue("chat_id"), 0, 0)

	var msg = tgBotAPI.NewDocument(chatId, tgBotAPI.FileReader{
		Name:   downFile.Filename,
		Reader: file,
	})
	msg.Caption = r.FormValue("caption")

	rsp, err := models2.TgAPI.Send(msg)

	if err != nil {
		fmt.Println("Error sending message.")
		w.WriteHeader(http.StatusNotModified)
		b, _ := json.Marshal(err)
		_, err = w.Write(b)
		if err != nil {
			return err
		}
		return err
	} else {
		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(rsp)
		_, err = w.Write(b)
		if err != nil {
			return err
		}
	}
	return nil
}

func SendAudio(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseMultipartForm(1024 << 20)
	if err != nil {
		fmt.Println("Error parsing form.")
		w.WriteHeader(http.StatusNotAcceptable)
		return err
	}

	downFile := r.MultipartForm.File["audio"][0]
	file, err := downFile.Open()
	if err != nil {
		fmt.Println("Error opening file.")
		return err
	}

	chatId, _ := strconv.ParseInt(r.FormValue("chat_id"), 0, 0)

	var msg = tgBotAPI.NewAudio(chatId, tgBotAPI.FileReader{
		Name:   downFile.Filename,
		Reader: file,
	})

	rsp, err := models2.TgAPI.Send(msg)
	if err != nil {
		fmt.Println("Error sending message.")
		w.WriteHeader(http.StatusNotModified)
		b, _ := json.Marshal(err)
		_, err = w.Write(b)
		if err != nil {
			return err
		}
		return err
	} else {
		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(rsp)
		_, err = w.Write(b)
		if err != nil {
			return err
		}
	}
	return nil
}

func SendSticker(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseMultipartForm(1024 << 20)
	if err != nil {
		fmt.Println("Error while parsing form.")
		w.WriteHeader(http.StatusNotAcceptable)
		return err
	}

	downFile := r.MultipartForm.File["sticker"][0]
	file, err := downFile.Open()
	if err != nil {
		fmt.Println("Error opening file.")
		return err
	}
	fmt.Println(downFile.Header.Get("sticker"))

	chatId, _ := strconv.ParseInt(r.FormValue("chat_id"), 0, 0)
	myImg, err := imgconv.Decode(file)
	if err != nil {
		fmt.Println("Error decoding file.")
		w.WriteHeader(http.StatusNotAcceptable)
		return err
	}

	fileBytes := new(bytes.Buffer)
	//if err := png.Encode(fileBytes, myImg); err != nil {
	//	fmt.Println("Error encoding file.")
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return err
	//}
	if err := imgconv.Write(fileBytes, myImg, &imgconv.FormatOption{Format: imgconv.PNG, EncodeOption: []imgconv.EncodeOption{imgconv.PNGCompressionLevel(png.BestSpeed)}}); err != nil {
		return err
	}

	//if err != nil {
	//	fmt.Println("Error while resizing")
	//	w.WriteHeader(http.StatusNotAcceptable)
	//	return err
	//}
	ioutil.WriteFile("sticker.png", fileBytes.Bytes(), fs.ModeAppend)

	var msg = tgBotAPI.NewSticker(chatId, tgBotAPI.FileBytes{
		Name:  downFile.Filename,
		Bytes: fileBytes.Bytes(),
	})

	rsp, err := models2.TgAPI.Send(msg)
	if err != nil {
		fmt.Println("Error sending message.")
		w.WriteHeader(http.StatusInternalServerError)
		b, _ := json.Marshal(err)
		w.Write(b)

	} else {
		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(rsp)
		w.Write(b)
	}
	return err
}

func SendContact(w http.ResponseWriter, r *http.Request) error {
	return nil
}
