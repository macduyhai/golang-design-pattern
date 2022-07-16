package proxy

import "log"

// Mục đích
// Kiểm soát quyền truy suất các phương thức của đối tượng
// Bổ xung thêm chức năng trước khi thực thi phương thức gốc
// Tạo ra đối tượng mới có chức năng cao hơn đối tượng ban đầu
// Giảm chi phí khi có nhiều truy cập vào đối tượng có chi phí khởi tạo ban đầu lớn
type Image interface {
	Display()
}
type HighResolutionImage struct {
	imageFilePath string
}
type ImageProxy struct {
	imageFilePath string
	realImage     Image
}

func (this *HighResolutionImage) loadImage(path string) {
	log.Printf("Load images %s from disk\n", path)
}
func (this *HighResolutionImage) Display() {

	this.loadImage(this.imageFilePath)
	log.Println("Display high resolution image")
}
func (this *ImageProxy) Display() {
	this.realImage = &HighResolutionImage{imageFilePath: this.imageFilePath}
	this.realImage.Display()
}
func NewImageProxy(path string) Image {
	return &ImageProxy{imageFilePath: path}
}
