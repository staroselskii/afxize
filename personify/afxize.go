package personify

import (
	"image"
	"image/draw"
	"github.com/staroselskii/afxize/facefinder"

	"github.com/disintegration/imaging"
)

func Personify(baseImage image.Image, haarCascade *string, personFaces facefinder.FaceList) image.Image {
	finder := facefinder.NewFinder(*haarCascade)

	faces := finder.Detect(baseImage)

	bounds := baseImage.Bounds()

	canvas := facefinder.CanvasFromImage(baseImage)

	for _, face := range faces {
        rect := facefinder.RectMargin(30.0, face)

		newFace := personFaces.Random()
		if newFace == nil {
			panic("nil face")
		}
		chrisFace := imaging.Fit(newFace, rect.Dx(), rect.Dy(), imaging.Lanczos)

		draw.Draw(
			canvas,
			rect,
			chrisFace,
			bounds.Min,
			draw.Over,
		)
	}

	if len(faces) == 0 {
		face := imaging.Resize(
			personFaces[0],
			bounds.Dx()/3,
			0,
			imaging.Lanczos,
		)
		face_bounds := face.Bounds()
		draw.Draw(
			canvas,
			bounds,
			face,
			bounds.Min.Add(image.Pt(-bounds.Max.X/2+face_bounds.Max.X/2, -bounds.Max.Y+int(float64(face_bounds.Max.Y)/1.9))),
			draw.Over,
		)
	}

    return canvas
}
