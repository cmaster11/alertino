package config

import "testing"

func TestInputInput(t *testing.T) {

	srcTest1Key := "test1"

	sampleConfig := `
inputs:
  test1:
    hashTemplate: "{{.name}}-2"
`

	ioConfig, err := parseIOConfigFromBytes([]byte(sampleConfig))
	if err != nil {
		t.Fatal(err)
	}

	result, err := ioConfig.Inputs[srcTest1Key].HashTemplate.Execute(map[string]interface{}{
		"name": "Bla",
	})
	if err != nil {
		t.Fatal(err)
	}

	if result != "Bla-2" {
		t.Fatal("bad template execution")
	}

}
