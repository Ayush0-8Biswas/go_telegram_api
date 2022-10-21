package controllers

import (
	"bytes"
	"fmt"
	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go_telegram_api/app/cmd/config"
	models2 "go_telegram_api/app/pkg/models"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func HandleText(mess *tgBotAPI.Message, dest string) error {
	data := url.Values{}
	data.Add("chat_id", dest)

	var name string = "*"
	if mess.From.UserName != "" {
		name += mess.From.UserName
	} else {
		name += mess.From.FirstName + mess.From.LastName
	}
	name += "*"

	data.Add("text", name+": "+mess.Text)

	resp, err := http.PostForm("http://localhost"+config.WhatsappPort+"/sendMessage/", data)
	if err != nil {
		return err
	}

	log.Println(resp)
	return nil
}

func HandlePhoto(mess *tgBotAPI.Message, dest string) error {
	client := &http.Client{Timeout: time.Minute * 60}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileUrl, err := models2.TgAPI.GetFileDirectURL(mess.Photo[len(mess.Photo)-1].FileID)
	if err != nil {
		log.Println("Error getting direct url.")
		return err
	}

	resp, err := client.Get(fileUrl)
	if err != nil {
		log.Println("Error getting file.")
		return err
	}

	var name string = "*"
	if mess.From.UserName != "" {
		name += mess.From.UserName
	} else {
		name += mess.From.FirstName + mess.From.LastName
	}
	name += "*"

	sendQuery := map[string]interface{}{
		"chat_id": dest,
		"caption": mess.Caption,
		"sender":  name,
	}

	for k, v := range sendQuery {
		fw, err := writer.CreateFormField(k)
		_, err = io.Copy(fw, strings.NewReader(fmt.Sprint(v)))
		if err != nil {
			return err
		}
	}
	fw, err := writer.CreateFormFile("image", fileUrl)
	_, err = io.Copy(fw, resp.Body)
	if err != nil {
		fmt.Println("Error while copying data")
		return err
	}
	err = writer.Close()
	if err != nil {
		fmt.Println("Error while closing writer.")
		return err
	}

	req, err := http.NewRequest("POST", "http://localhost"+config.WhatsappPort+"/sendPhoto/", body)
	if err != nil {
		fmt.Println("Error creating request.")
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err = client.Do(req)
	if err != nil {
		return err
	}

	log.Println(resp)
	return nil
}

func HandleVideo(mess *tgBotAPI.Message, dest string) error {
	client := &http.Client{Timeout: time.Minute * 60}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileUrl, err := models2.TgAPI.GetFileDirectURL(mess.Video.FileID)
	if err != nil {
		log.Println("Error getting direct url.")
		return err
	}

	resp, err := client.Get(fileUrl)
	if err != nil {
		log.Println("Error getting file.")
		return err
	}

	var name string = "*"
	if mess.From.UserName != "" {
		name += mess.From.UserName
	} else {
		name += mess.From.FirstName + mess.From.LastName
	}
	name += "*"

	sendQuery := map[string]interface{}{
		"chat_id": dest,
		"caption": mess.Caption,
		"sender":  name,
	}

	for k, v := range sendQuery {
		fw, err := writer.CreateFormField(k)
		_, err = io.Copy(fw, strings.NewReader(fmt.Sprint(v)))
		if err != nil {
			return err
		}
	}
	fw, err := writer.CreateFormFile("video", fileUrl)
	_, err = io.Copy(fw, resp.Body)
	if err != nil {
		fmt.Println("Error while copying data")
		return err
	}
	err = writer.Close()
	if err != nil {
		fmt.Println("Error while closing writer.")
		return err
	}

	req, err := http.NewRequest("POST", "http://localhost"+config.WhatsappPort+"/sendVideo/", body)
	if err != nil {
		fmt.Println("Error creating request.")
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err = client.Do(req)
	if err != nil {
		return err
	}

	log.Println(resp)
	return nil
}

func HandleDocument(mess *tgBotAPI.Message, dest string) error {
	client := &http.Client{Timeout: time.Minute * 60}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileUrl, err := models2.TgAPI.GetFileDirectURL(mess.Document.FileID)
	if err != nil {
		log.Println("Error getting direct url.")
		return err
	}

	resp, err := client.Get(fileUrl)
	if err != nil {
		log.Println("Error getting file.")
		return err
	}

	var name string = "*"
	if mess.From.UserName != "" {
		name += mess.From.UserName
	} else {
		name += mess.From.FirstName + mess.From.LastName
	}
	name += "*"

	sendQuery := map[string]interface{}{
		"chat_id": dest,
		"caption": mess.Caption,
		"sender":  name,
	}

	for k, v := range sendQuery {
		fw, err := writer.CreateFormField(k)
		_, err = io.Copy(fw, strings.NewReader(fmt.Sprint(v)))
		if err != nil {
			return err
		}
	}
	fw, err := writer.CreateFormFile("file", fileUrl)
	_, err = io.Copy(fw, resp.Body)
	if err != nil {
		fmt.Println("Error while copying data")
		return err
	}
	err = writer.Close()
	if err != nil {
		fmt.Println("Error while closing writer.")
		return err
	}

	req, err := http.NewRequest("POST", "http://localhost"+config.WhatsappPort+"/sendDocument/", body)
	if err != nil {
		fmt.Println("Error creating request.")
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err = client.Do(req)
	if err != nil {
		return err
	}

	log.Println(resp)
	return nil
}

func HandleAudio(mess *tgBotAPI.Message, dest string) error {
	client := &http.Client{Timeout: time.Minute * 60}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileUrl, err := models2.TgAPI.GetFileDirectURL(mess.Audio.FileID)
	if err != nil {
		log.Println("Error getting direct url.")
		return err
	}

	resp, err := client.Get(fileUrl)
	if err != nil {
		log.Println("Error getting file.")
		return err
	}

	var name string = "*"
	if mess.From.UserName != "" {
		name += mess.From.UserName
	} else {
		name += mess.From.FirstName + mess.From.LastName
	}
	name += "*"

	sendQuery := map[string]interface{}{
		"chat_id": dest,
		"sender":  name,
	}

	for k, v := range sendQuery {
		fw, err := writer.CreateFormField(k)
		_, err = io.Copy(fw, strings.NewReader(fmt.Sprint(v)))
		if err != nil {
			return err
		}
	}
	fw, err := writer.CreateFormFile("audio", fileUrl)
	_, err = io.Copy(fw, resp.Body)
	if err != nil {
		fmt.Println("Error while copying data")
		return err
	}
	err = writer.Close()
	if err != nil {
		fmt.Println("Error while closing writer.")
		return err
	}

	req, err := http.NewRequest("POST", "http://localhost"+config.WhatsappPort+"/sendAudio/", body)
	if err != nil {
		fmt.Println("Error creating request.")
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err = client.Do(req)
	if err != nil {
		return err
	}

	log.Println(resp)
	return nil
}

func HandleSticker(mess *tgBotAPI.Message, dest string) error {
	client := &http.Client{Timeout: time.Minute * 60}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileUrl, err := models2.TgAPI.GetFileDirectURL(mess.Sticker.FileID)
	if err != nil {
		log.Println("Error getting direct url.")
		return err
	}

	resp, err := client.Get(fileUrl)
	if err != nil {
		log.Println("Error getting file.")
		return err
	}

	var name string = "*"
	if mess.From.UserName != "" {
		name += mess.From.UserName
	} else {
		name += mess.From.FirstName + mess.From.LastName
	}
	name += "*"

	sendQuery := map[string]interface{}{
		"chat_id": dest,
		"sender":  name,
	}

	for k, v := range sendQuery {
		fw, err := writer.CreateFormField(k)
		_, err = io.Copy(fw, strings.NewReader(fmt.Sprint(v)))
		if err != nil {
			return err
		}
	}
	fw, err := writer.CreateFormFile("sticker", fileUrl)
	_, err = io.Copy(fw, resp.Body)
	if err != nil {
		fmt.Println("Error while copying data")
		return err
	}
	err = writer.Close()
	if err != nil {
		fmt.Println("Error while closing writer.")
		return err
	}

	req, err := http.NewRequest("POST", "http://localhost"+config.WhatsappPort+"/sendSticker/", body)
	if err != nil {
		fmt.Println("Error creating request.")
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err = client.Do(req)
	if err != nil {
		return err
	}

	log.Println(resp)
	return nil
}
