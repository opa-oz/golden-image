package golden_image

import (
	"fmt"
	"github.com/gkampitakis/ciinfo"
	"github.com/opa-oz/hikaku"
	"image"
	_ "image/jpeg"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

var testsRegistry = newRegistry()
var isCI = ciinfo.IsCI
var shouldUpdate = getEnvBool("UPDATE_SNAPS", false) && !isCI

func ToGildImage(t *testing.T, threshold float64, values ...image.Image) {
	t.Helper()

	if len(values) == 0 {
		t.Log("[Warning] ToGildImage called without params")
		return
	}

	for _, img := range values {
		toGildImage(t, threshold, img)
	}
}

func toGildImage(t *testing.T, threshold float64, value image.Image) {
	t.Helper()

	dir, goldenPath := getPaths(goldenDir)
	d, diffPath := getPaths(path.Join(goldenDir, "__diffs__"))
	testID := testsRegistry.getTestID(t.Name(), goldenPath)

	prevSnapshot, err := getPrevSnapshot(testID, goldenPath)

	if err != nil {
		if isCI {
			t.Error(err)
			return
		}

		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			t.Error(err)
			return
		}

		err := saveSnapshot(testID, goldenPath, value)
		if err != nil {
			t.Error(err)
			return
		}

		t.Log("› New snapshot written.\n")
		return
	}

	isIdenticalByBounds := hikaku.CompareByParams(prevSnapshot, value)

	if !isIdenticalByBounds {
		if shouldUpdate {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				t.Error(err)
				return
			}

			err := saveSnapshot(testID, goldenPath, value)
			cleanDiffPath(buildPath(diffPath, testID))

			if err != nil {
				t.Error(err)
				return
			}

			t.Log("› Snapshot updated.\n")
			return
		}

		prevBounds := prevSnapshot.Bounds().Max
		valueBounds := value.Bounds().Max
		t.Error(fmt.Sprintf("• Images' shapes are different. \n\tOriginal is %dx%d\n\tNew is %dx%d", prevBounds.X, prevBounds.Y, valueBounds.X, valueBounds.Y))
		return
	}

	params := hikaku.ComparisonParameters{Threshold: threshold, BinsCount: 32}

	goldenHist := hikaku.PrepareHistogram(prevSnapshot, params)
	copperHist := hikaku.PrepareHistogram(value, params)

	params.NormalizedGoldHist = goldenHist
	params.NormalizedCopperHist = copperHist

	isIdentical, diff := hikaku.CompareHistogramsOnly(params)

	if isIdentical {
		cleanDiffPath(buildPath(diffPath, testID))
		return
	}

	if shouldUpdate {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			t.Error(err)
			return
		}

		err := saveSnapshot(testID, goldenPath, value)
		cleanDiffPath(buildPath(diffPath, testID))
		if err != nil {
			t.Error(err)
			return
		}

		t.Log("› Snapshot updated.\n")
		return
	}

	if err := os.MkdirAll(d, os.ModePerm); err != nil {
		t.Error(err)
		return
	}

	diffMask := hikaku.FindDiffShapesMask(prevSnapshot, value, hikaku.ContextParameters{})
	diffImage := hikaku.ApplyDiff(value, diffMask, 128)

	err = saveSnapshot(testID, diffPath, diffImage)
	if err != nil {
		t.Log(err)
		return
	}

	t.Error(fmt.Sprintf("• Found diff: %F. \n\tHighlights are stored at %s.\n", diff, buildPath(diffPath, testID)))
}

func getPaths(folder string) (dir, goldenPath string) {
	callerPath := findRoot()
	base := filepath.Base(callerPath)

	dir = filepath.Join(filepath.Dir(callerPath), folder)
	goldenPath = filepath.Join(dir, strings.TrimSuffix(base, filepath.Ext(base)))

	return
}
