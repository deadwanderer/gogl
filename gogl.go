package gogl

import (
	"fmt"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type ShaderID uint32
type ProgramID uint32
type BufferID uint32
type TextureID uint32

func MglTest() {
	x := mgl32.NewVecN(2)
	fmt.Println(x)
}

func GetVersion() string {
	return gl.GoStr(gl.GetString(gl.VERSION))
}

func TriangleNormal(p1, p2, p3 mgl32.Vec3) mgl32.Vec3 {
	U := p2.Sub(p1)
	V := p2.Sub(p3)
	return U.Cross(V).Normalize()
}
