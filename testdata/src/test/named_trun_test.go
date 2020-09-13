package test

import "testing"

func tRun1(t *testing.T) { // want "tRun1's subtests should call t.Parallel"
	t.Parallel()

	t.Run("Named_tRun1_Sub", func(t *testing.T) {
		call("Named_tRun1_Sub")
	})
}

func tRun2(t *testing.T) { // want "tRun2's subtests should call t.Parallel"
	t.Run("Named_tRun2_Sub", func(t *testing.T) {
		call("Named_tRun2_Sub")
	})
}

func tRun3(t *testing.T) { // want "tRun3 should call t.Parallel on the top level as well as its subtests"
	t.Run("Named_tRun3_Sub", func(t *testing.T) {
		t.Parallel()
		call("Named_tRun3_Sub")
	})
}

func tRun4(t *testing.T) { // OK
	t.Run("Named_tRun4_Sub", func(t *testing.T) {
		t.Parallel()
		call("Named_tRun4_Sub")
	})
}

func tRun5(t *testing.T) { // OK
	t.Parallel()

	t.Run("Named_tRun5_Sub", func(t *testing.T) {
		t.Parallel()
		call("Named_tRun5_Sub")
	})
}

func Test_Named_tRun1(t *testing.T) {
	call("Test_Named_tRun1")

	tRun1(t)
}

func Test_Named_tRun2(t *testing.T) {
	t.Parallel()

	call("Test_Named_tRun2")

	tRun2(t)
}

func Test_Named_tRun3(t *testing.T) {
	call("Test_Named_tRun3")

	tRun3(t)
}

func Test_Named_tRun4(t *testing.T) {
	t.Parallel()

	call("Test_Named_tRun4")

	tRun4(t)
}

func Test_Named_tRun5(t *testing.T) {
	call("Test_Named_tRun5")

	tRun5(t)
}
