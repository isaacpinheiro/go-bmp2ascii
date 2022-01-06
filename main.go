package main

import (
    "fmt"
    "os"
    "log"
    "encoding/binary"
)

type BitmapHeader struct {
    Signature [2]byte
    Size uint32
    Reserved [4]int8
    Offset uint32
}

type DIB struct {
    SizeHeader uint32
    Width uint32
    Height uint32
    Planes uint16
    Bpp uint16
    Compression uint32
    ImgSize uint32
    Hppm uint32
    Vppm uint32
    NColors uint32
    NImportantColors uint32
    Redmask uint32
    Bluemask uint32
    Greenmask uint32
    Alphamask uint32
    CsType uint32
    Pad [36]int8
    GammaRed uint32
    GammaGreen uint32
    GammaBlue uint32
}

type Color32 struct {
    Blue byte
    Green byte
    Red byte
    Alpha byte
}

type Color struct {
    Blue byte
    Green byte
    Red byte
}

func generateImage(color []Color, width, height uint32) string {

    var pixel [10]string = [10]string{ " ", ".", ",", ":", ";", "o", "x", "%", "#", "@" }
    var i, j int32
    var res string = ""

    for i = int32(height) - 1; i >= 0; i-- {

        for j = 0; j < int32(width); j++ {

            var c Color = color[i*int32(width)+j]
            var intensity uint8 = uint8(9 - (((int(c.Blue) + int(c.Green) + int(c.Red))*9)/(255*3)))
            res += pixel[intensity]

        }

        res += "\n"

    }

    return res
}

func main() {

    var input, output, image string
    var inputFile, outputFile *os.File
    var err error
    var bh BitmapHeader
    var dib DIB
    var color []Color

    if len(os.Args) < 2 {

        fmt.Printf("\nbmp2ascii - version: 1.0\n")
        fmt.Printf("Author: Isaac Pinheiro <isaacpnhr@gmail.com>\n\n")
        fmt.Printf("Usage:\n\n")
        fmt.Printf("\tBasic usage:\n\n")
        fmt.Printf("\t\t$ bmp2ascii /path/to/your/file.bmp\n\n")
        fmt.Printf("\tGenerating an output file: \n\n")
        fmt.Printf("\t\t$ bmp2ascii /path/to/your/file.bmp /path/to/your/output_file.txt\n\n")
        fmt.Printf("\t\t\tor\n\n")
        fmt.Printf("\t\t$ bmp2ascii /path/to/your/file.bmp > output_file.txt\n\n")

    } else {

        input = os.Args[1]
        inputFile, err = os.Open(input)

        if err != nil {
            log.Fatal(err)
        }

        defer inputFile.Close()

        if len(os.Args) > 2 {

            output = os.Args[2]
            outputFile, err = os.Create(output)

            if err != nil {
                log.Fatal(err)
            }

            defer outputFile.Close()

        }

        binary.Read(inputFile, binary.LittleEndian, &bh)
        binary.Read(inputFile, binary.LittleEndian, &dib)

        color = make([]Color, dib.Width * dib.Height)
        binary.Read(inputFile, binary.LittleEndian, &color)

        image = generateImage(color, dib.Width, dib.Height)

        if len(os.Args) > 2 {
            outputFile.WriteString(image)
        } else {
            fmt.Printf("%s", image)
        }

    }

}

