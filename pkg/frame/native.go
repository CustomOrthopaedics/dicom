package frame

import (
	"image"
	"image/color"
)

// NativeFrame represents a native image frame. This is an interface to encapsulate underlying data storage type.
type NativeFrame interface {
	Rows() int
	Cols() int
	BitsPerSample() int
	GetPixel(x, y int) int
	SetPixel(idx int, value []int)
	IsEncapsulated() bool
	GetNativeFrame() (*NativeFrame, error)
	GetEncapsulatedFrame() (*EncapsulatedFrame, error)
	GetImage() (image.Image, error)
}

func NewNativeFrame(bitsPerSample, rows, cols int) *NativeFrame {
	if bitsPerSample == 8 {
		return &NativeFrame8{
			bitsPerSample: bitsPerSample,
			rows:          rows,
			cols:          cols,
			data:          make([][]uint8, rows*cols),
		}
	}

	return &NativeFrame8{}

	// if bitsPerSample == 16 {
	// 	return nativeFrame16{
	// 		bitsPerSample: bitsPerSample,
	// 		rows:          rows,
	// 		cols:          cols,
	// 		data:          make([][]uint16, rows*cols),
	// 	}
	// }

	// return nativeFrame32{
	// 	bitsPerSample: bitsPerSample,
	// 	rows:          rows,
	// 	cols:          cols,
	// 	data:          make([][]uint32, rows*cols),
	// }
}

type NativeFrame8 struct {
	data          [][]uint8
	rows          int
	cols          int
	bitsPerSample int
}

func (n *NativeFrame8) Rows() int {
	return n.rows
}

func (n *NativeFrame8) Cols() int {
	return n.cols
}

func (n *NativeFrame8) BitsPerSample() int {
	return n.bitsPerSample
}

func (n *NativeFrame8) GetPixel(x, y int) int {
	// TODO: implement
	return n.bitsPerSample
}

func (n *NativeFrame8) SetPixel(idx int, value []int) {
	for i := 0; i < len(value); i++ {
		n.data[idx][i] = uint8(value[i])
	}
}

// IsEncapsulated indicates if the frame is encapsulated or not.
func (n *NativeFrame8) IsEncapsulated() bool { return false }

// GetNativeFrame returns a NativeFrame from this frame. If the underlying frame
// is not a NativeFrame, ErrorFrameTypeNotPresent will be returned.
func (n *NativeFrame8) GetNativeFrame() (*NativeFrame, error) {
	return n, nil
}

// GetEncapsulatedFrame returns ErrorFrameTypeNotPresent, because this struct
// does not hold encapsulated frame data.
func (n *NativeFrame8) GetEncapsulatedFrame() (*EncapsulatedFrame, error) {
	return nil, ErrorFrameTypeNotPresent
}

// GetImage returns an image.Image representation the frame, using default
// processing. This default processing is basic at the moment, and does not
// autoscale pixel values or use window width or level info.
func (n *NativeFrame8) GetImage() (image.Image, error) {
	i := image.NewGray16(image.Rect(0, 0, n.cols, n.rows))
	for j := 0; j < len(n.data); j++ {
		i.SetGray16(j%n.cols, j/n.cols, color.Gray16{Y: uint16(n.data[j][0])}) // for now, assume we're not overflowing uint16, assume gray image
	}
	return i, nil
}
