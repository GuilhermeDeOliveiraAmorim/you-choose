package usecases

import (
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
	"github.com/google/uuid"
)

type FileUseCase struct {
	FileRepository entity.FileRepositoryInterface
}

func NewFileUseCase(fileRepository entity.FileRepositoryInterface) *FileUseCase {
	return &FileUseCase{
		FileRepository: fileRepository,
	}
}

func (fileUseCase *FileUseCase) Create(input InputCreateFileDto) (OutputCreateFileDto, error) {
	output := OutputCreateFileDto{}

	_, name, size, extension, err := MoveFile(input.File, input.Handler)
	if err != nil {
		return output, errors.New(err.Error())
	}

	colorAverage, err := PictureAverageColor(name, extension)
	if err != nil {
		return output, errors.New(err.Error())
	}

	fileEntity, err := entity.NewFile(name, input.EntityId, size, extension, colorAverage)
	if err != nil {
		return output, errors.New(err.Error())
	}

	if err := fileUseCase.FileRepository.Create(fileEntity); err != nil {
		return output, errors.New(err.Error())
	}

	output.File.ID = fileEntity.ID
	output.File.EntityId = fileEntity.EntityId
	output.File.Name = fileEntity.Name
	output.File.Size = fileEntity.Size
	output.File.Extension = fileEntity.Extension
	output.File.IsDeleted = fileEntity.IsDeleted
	output.File.CreatedAt = fileEntity.CreatedAt
	output.File.UpdatedAt = fileEntity.UpdatedAt
	output.File.DeletedAt = fileEntity.DeletedAt

	return output, nil
}

func (fileUseCase *FileUseCase) Find(input InputFindFileDto) (OutputFindFileDto, error) {
	output := OutputFindFileDto{}

	file, err := fileUseCase.FileRepository.Find(input.FileId)
	if err != nil {
		return output, errors.New(err.Error())
	}

	output.File.ID = file.ID
	output.File.EntityId = file.EntityId
	output.File.Name = file.Name
	output.File.Size = file.Size
	output.File.Extension = file.Extension
	output.File.IsDeleted = file.IsDeleted
	output.File.CreatedAt = file.CreatedAt
	output.File.UpdatedAt = file.UpdatedAt
	output.File.DeletedAt = file.DeletedAt

	return output, nil
}

// func (fileUseCase *FileUseCase) Delete(input InputDeleteFileDto) (OutputDeleteFileDto, error) {
// 	timeNow := time.Now().Local().String()
// 	output := OutputDeleteFileDto{}

// 	file, err := fileUseCase.FileRepository.Find(input.FileId)
// 	if err != nil {
// 		return output, errors.New(err.Error())
// 	}

// 	if file.IsDeleted {
// 		return output, errors.New("file previously deleted")
// 	}

// 	file.IsDeleted = true
// 	file.DeletedAt = timeNow

// 	output.IsDeleted = file.IsDeleted

// 	return output, errors.New(err.Error())
// }

// func (fileUseCase *FileUseCase) Update(input InputUpdateFileDto) (OutputUpdateFileDto, error) {
// 	timeNow := time.Now().Local().String()
// 	output := OutputUpdateFileDto{}

// 	file, err := fileUseCase.FileRepository.Find(input.FileId)
// 	if err != nil {
// 		return output, errors.New(err.Error())
// 	}

// 	file.Name = input.Name

// 	isValid, err := file.Validate()
// 	if !isValid {
// 		return output, errors.New(err.Error())
// 	}

// 	file.UpdatedAt = timeNow

// 	err = fileUseCase.FileRepository.Update(&file)
// 	if err != nil {
// 		return output, errors.New(err.Error())
// 	}

// 	output.File.ID = file.ID
// 	output.File.EntityId = file.EntityId
// 	output.File.Name = file.Name
// 	output.File.Size = file.Size
// 	output.File.Extension = file.Extension
// 	output.File.IsDeleted = file.IsDeleted
// 	output.File.CreatedAt = file.CreatedAt
// 	output.File.UpdatedAt = file.UpdatedAt
// 	output.File.DeletedAt = file.DeletedAt

// 	return output, nil
// }

// func (fileUseCase *FileUseCase) IsDeleted(input InputIsDeletedFileDto) (OutputIsDeletedFileDto, error) {
// 	output := OutputIsDeletedFileDto{}

// 	file, err := fileUseCase.FileRepository.Find(input.FileId)
// 	if err != nil {
// 		return output, errors.New(err.Error())
// 	}

// 	output.IsDeleted = false

// 	if file.IsDeleted {
// 		output.IsDeleted = true
// 	}

// 	return output, nil
// }

// func (fileUseCase *FileUseCase) FindAll() (OutputFindAllFileDto, error) {
// 	output := OutputFindAllFileDto{}

// 	files, err := fileUseCase.FileRepository.FindAll()
// 	if err != nil {
// 		return output, errors.New(err.Error())
// 	}

// 	for _, file := range files {
// 		output.Files = append(output.Files, FileDto{
// 			ID:        file.ID,
// 			EntityId:  file.EntityId,
// 			Name:      file.Name,
// 			Size:      file.Size,
// 			Extension: file.Extension,
// 			IsDeleted: file.IsDeleted,
// 			CreatedAt: file.CreatedAt,
// 			UpdatedAt: file.UpdatedAt,
// 			DeletedAt: file.DeletedAt,
// 		})
// 	}

// 	return output, nil
// }

func MoveFile(file multipart.File, handler *multipart.FileHeader) (int64, string, int64, string, error) {
	path := "upload/"

	extension := filepath.Ext(handler.Filename)

	name := uuid.New().String()

	size := handler.Size

	fileCreate, err := os.Create(path + name + extension)
	if err != nil {
		return 0, "", 0, "", errors.New(err.Error())
	}

	defer file.Close()
	defer fileCreate.Close()

	fileWritten, err := io.Copy(fileCreate, file)
	if err != nil {
		return 0, "", 0, "", errors.New(err.Error())
	}

	extension = strings.Replace(filepath.Ext(handler.Filename), ".", "", -1)

	return fileWritten, name, size, extension, nil
}

func PictureToBase64(path string, name string, extension string) (string, error) {
	pictureBytes, err := ioutil.ReadFile(path + name + "." + extension)
	if err != nil {
		return "", errors.New(err.Error())
	}

	var pictureBase64 string

	mimeType := http.DetectContentType(pictureBytes)

	switch mimeType {
	case "image/jpeg":
		pictureBase64 += "data:image/jpeg;base64,"
	case "image/png":
		pictureBase64 += "data:image/png;base64,"
	case "image/jpg":
		pictureBase64 += "data:image/jpg;base64,"
	}

	pictureBase64 += base64.StdEncoding.EncodeToString(pictureBytes)

	return pictureBase64, nil
}

func PictureAverageColor(name string, extension string) (string, error) {
	file, err := os.Open("/home/guilherme/Workspace/you-choose/cmd/upload/" + name + "." + extension)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}

	imgSize := img.Bounds().Size()

	var redSum float64
	var greenSum float64
	var blueSum float64

	for x := 0; x < imgSize.X; x++ {
		for y := 0; y < imgSize.Y; y++ {
			pixel := img.At(x, y)
			col := color.RGBAModel.Convert(pixel).(color.RGBA)

			redSum += float64(col.R)
			greenSum += float64(col.G)
			blueSum += float64(col.B)
		}
	}

	imgArea := float64(imgSize.X * imgSize.Y)

	red := redSum / imgArea
	green := greenSum / imgArea
	blue := blueSum / imgArea

	rgbToHex, err := RgbToHex(red, green, blue)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return rgbToHex, nil
}

func RgbToHex(R float64, G float64, B float64) (string, error) {
	a := fmt.Sprintf("%f", R/16)
	aa, err := strconv.ParseFloat(strings.Split(a, ".")[0], 64)
	if err != nil {
		return "", errors.New("error during conversion")
	}

	b := fmt.Sprintf("%f", ((R/16)-aa)*16)
	bb, err := strconv.ParseFloat(strings.Split(b, ".")[0], 64)
	if err != nil {
		return "", errors.New("error during conversion")
	}

	c := fmt.Sprintf("%f", G/16)
	cc, err := strconv.ParseFloat(strings.Split(c, ".")[0], 64)
	if err != nil {
		return "", errors.New("error during conversion")
	}

	d := fmt.Sprintf("%f", ((G/16)-cc)*16)
	dd, err := strconv.ParseFloat(strings.Split(d, ".")[0], 64)
	if err != nil {
		return "", errors.New("error during conversion")
	}

	e := fmt.Sprintf("%f", (B / 16))
	ee, err := strconv.ParseFloat(strings.Split(e, ".")[0], 64)
	if err != nil {
		return "", errors.New("error during conversion")
	}

	f := fmt.Sprintf("%f", ((B/16)-ee)*16)
	ff, err := strconv.ParseFloat(strings.Split(f, ".")[0], 64)
	if err != nil {
		return "", errors.New("error during conversion")
	}

	g := fmt.Sprintf("%f", aa)
	h := fmt.Sprintf("%f", bb)
	i := fmt.Sprintf("%f", cc)
	j := fmt.Sprintf("%f", dd)
	k := fmt.Sprintf("%f", ee)
	l := fmt.Sprintf("%f", ff)

	var decimalsToConvert []string

	decimalsToConvert = append(decimalsToConvert, strings.Split(g, ".")[0])
	decimalsToConvert = append(decimalsToConvert, strings.Split(h, ".")[0])
	decimalsToConvert = append(decimalsToConvert, strings.Split(i, ".")[0])
	decimalsToConvert = append(decimalsToConvert, strings.Split(j, ".")[0])
	decimalsToConvert = append(decimalsToConvert, strings.Split(k, ".")[0])
	decimalsToConvert = append(decimalsToConvert, strings.Split(l, ".")[0])

	decimalsToConvert = DecToHexTable(decimalsToConvert)

	return "#" + decimalsToConvert[0] + decimalsToConvert[1] + decimalsToConvert[2] + decimalsToConvert[3] + decimalsToConvert[4] + decimalsToConvert[5], nil
}

func DecToHexTable(decimalsToConvert []string) []string {
	var dec = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", "62", "63"}
	var hex = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "1a", "1b", "1c", "1d", "1e", "1f", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "2a", "2b", "2c", "2d", "2e", "2f", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "3a", "3b", "3c", "3d", "3e", "3f"}

	for x, decimal := range decimalsToConvert {
		for y, d := range dec {
			if decimal == d {
				decimalsToConvert[x] = hex[y]
			}
		}
	}

	return decimalsToConvert
}
