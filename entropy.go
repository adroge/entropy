// Package entropy is used to calculate the entropy of a string.
//
// This can be used as one of the means to help determine how safe a password is in backend code.
//
// See entropy_test.go for usage.
package entropy

//go:generate mockgen -source entropy.go -destination mock_entropy/mock_entropy.go

import (
	"errors"
	"math"
)

const (
	Invalid = iota
	VeryWeak
	Weak
	Reasonable
	Strong
	VeryStrong
)

var (
	ErrInvalidAlphabet    = errors.New("invalid alphabet")
	ErrInvalidEntropy     = errors.New("invalid entropy")
	ErrInvalidDescription = errors.New("invalid description")
	ErrUnexpectedRune     = errors.New("unexpected rune in input")

	_Invalid    = "invalid"
	_Weak       = "weak"
	_VeryWeak   = "very weak"
	_Reasonable = "reasonable"
	_Strong     = "strong"
	_VeryStrong = "very strong"

	evaluationWords = map[int]*string{
		Invalid:    &_Invalid,
		VeryWeak:   &_VeryWeak,
		Weak:       &_Weak,
		Reasonable: &_Reasonable,
		Strong:     &_Strong,
		VeryStrong: &_VeryStrong,
	}

	runeTypes map[rune]int

	alphabets = []string{
		`abcdefghijklmnopqrstuvwxyz`,
		`ABCDEFGHIJKLMNOPQRSTUVWXYZ`,
		`1234567890`,
		`!@#$%^&*`,
		`()[]{}<>`,
		`~-_=+|;:',./? \"` + "`",
	}

	entropyBounds = [...]float64{
		30.0,
		40.0,
		60.0,
		127.0,
	}
)

// EntropyResult contains the bits, and the evaluation that maps to exported constants
type EntropyResult struct {
	Bits       float64
	Evaluation int
}

// DescriptionTags can be used to change the string representation of an entropy value
type DescriptionTags struct {
	Invalid    string
	VeryWeak   string
	Weak       string
	Reasonable string
	Strong     string
	VeryStrong string
}

// String implements the stringer interface
func (er EntropyResult) String() string {
	return *evaluationWords[er.Evaluation]
}

// Methods defines an interface that is used for testing
// with the generated mock_entropy code.
type Methods interface {
	Alphabets(newAlphabets []string) (err error)
	Bounds(veryWeak, weak, reasonable, strong float64) (err error)
	Calculate(input string) (result EntropyResult, err error)
	Descriptions(tags DescriptionTags) (err error)
	EntropyBounds() (bounds []float64)
}

type method struct{}

var call Methods = method{}

func init() {
	loadAlphabets()
}

func loadAlphabets() {
	runeTypes = make(map[rune]int)
	for alphabetIndex, alphabet := range alphabets {
		for _, runeCharacter := range alphabet {
			if _, found := runeTypes[runeCharacter]; !found {
				runeTypes[runeCharacter] = alphabetIndex
			}
		}
	}
}

// Alphabets replaces the currently defined alphabets.
func Alphabets(newAlphabets []string) (err error) {
	return call.Alphabets(newAlphabets)
}
func (method) Alphabets(newAlphabets []string) (err error) {
	if len(newAlphabets) == 0 {
		return ErrInvalidAlphabet
	}
	alphabets = newAlphabets
	loadAlphabets()
	return
}

// Bounds changes the upper ranges for evaluating the strength
// of the calculated entropy.
func Bounds(veryWeak, weak, reasonable, strong float64) (err error) {
	return call.Bounds(veryWeak, weak, reasonable, strong)
}
func (method) Bounds(veryWeak, weak, reasonable, strong float64) (err error) {
	if veryWeak >= weak || weak >= reasonable || reasonable >= strong {
		return ErrInvalidEntropy
	}

	entropyBounds[0] = veryWeak
	entropyBounds[1] = weak
	entropyBounds[2] = reasonable
	entropyBounds[3] = strong

	return
}

// Calculate analyzes the entropy of a string.
func Calculate(input string) (result EntropyResult, err error) {
	return call.Calculate(input)
}
func (method) Calculate(input string) (result EntropyResult, err error) {
	result.Bits, err = calculateEntropy(input)
	if err != nil {
		return
	}
	result.Evaluation, err = evaluateEntropy(result.Bits)
	return
}

func calculateEntropy(input string) (entropy float64, err error) {
	if len(input) == 0 {
		return
	}

	characterPool := make(map[int]int)
	for _, character := range input {
		if _, found := runeTypes[character]; !found {
			err = ErrUnexpectedRune
			return
		}
		characterPool[runeTypes[character]] = len(alphabets[runeTypes[character]])
	}

	poolSize := 0
	for _, length := range characterPool {
		poolSize += length
	}

	entropy = math.Log2(math.Pow(float64(poolSize), float64(len(input))))

	return
}

func evaluateEntropy(entropy float64) (evaluation int, err error) {
	if math.IsNaN(entropy) || entropy < 0.0 {
		err = ErrInvalidEntropy
		return
	}
	if entropy <= entropyBounds[0] {
		evaluation = VeryWeak
		return
	}
	if entropy <= entropyBounds[1] {
		evaluation = Weak
		return
	}
	if entropy <= entropyBounds[2] {
		evaluation = Reasonable
		return
	}
	if entropy <= entropyBounds[3] {
		evaluation = Strong
		return
	}
	evaluation = VeryStrong
	return
}

// Descriptions changes the string descriptions of the String() evaluation
// of the entropy value.
func Descriptions(tags DescriptionTags) (err error) {
	return call.Descriptions(tags)
}
func (method) Descriptions(tags DescriptionTags) (err error) {
	if len(tags.Invalid) == 0 || len(tags.VeryWeak) == 0 ||
		len(tags.Weak) == 0 || len(tags.Reasonable) == 0 ||
		len(tags.Strong) == 0 || len(tags.VeryStrong) == 0 {
		return ErrInvalidDescription
	}
	_Invalid = tags.Invalid
	_VeryWeak = tags.VeryWeak
	_Weak = tags.Weak
	_Reasonable = tags.Reasonable
	_Strong = tags.Strong
	_VeryStrong = tags.VeryStrong
	return
}

// EntropyBounds returns an array of bound values used internally.
func EntropyBounds() (bounds []float64) {
	return call.EntropyBounds()
}

func (method) EntropyBounds() (bounds []float64) {
	bounds = make([]float64, 4)
	copy(bounds, entropyBounds[:])
	return
}
