package main

import (
	"log"

	"github.com/sleonovich1188/goav/avcodec"
	"github.com/sleonovich1188/goav/avdevice"
	"github.com/sleonovich1188/goav/avfilter"
	"github.com/sleonovich1188/goav/avformat"
	"github.com/sleonovich1188/goav/avutil"
	"github.com/sleonovich1188/goav/swresample"
	"github.com/sleonovich1188/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
