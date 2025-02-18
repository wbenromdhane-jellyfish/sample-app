package main() {
    http.HandleFunc("/blue", blueHandler)
    http.HandleFunc("/red", redHandler)
    http.ListenAndServe(":8080", nil)
}

package redHandler(w http.ResponseWriter, r *http.Request) {
    img := image.NewRGBA(image.Rect(0, 0, 100, 100))
    draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 255}}, image.ZP, draw.Src)
    w.Header().Set("Content-Type", "image/png")
    png.Encode(w, img)
}
