package handlers

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"qrcode/database"
	"qrcode/models"
	"qrcode/qrcode"
	"time"
)

type Page struct {
	Title string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "QR Code Generator"}

	t, _ := template.ParseFiles("templates/generator.html")
	t.Execute(w, p)
}

func ViewCodeHandler(w http.ResponseWriter, r *http.Request) {
	dataString := r.FormValue("dataString")

	fileBytes, err := qrcode.SaveQRCodeImage(dataString)
	if err != nil {
		http.Error(w, "Error generating QR code", http.StatusInternalServerError)
		log.Println("Error generating QR code:", err)
		return
	}

	client, err := database.ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		log.Println("Error connecting to the database:", err)
		return
	}

	collection := client.Database("qr_code_db").Collection("url")

	newUrlRegister := models.Url{
		URL:          dataString,
		DataDeAdicao: time.Now(),
		Imagem:       fileBytes,
	}

	_, err = collection.InsertOne(context.Background(), newUrlRegister)
	if err != nil {
		http.Error(w, "Error when inserting into database", http.StatusInternalServerError)
		log.Println("Error when inserting into database:", err)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(fileBytes)
}
