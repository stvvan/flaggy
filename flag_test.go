package flaggy

import "testing"

func TestParseArgWithValue(t *testing.T) {
	DebugMode = true
	defer func() { DebugMode = false }()
	testCases := make(map[string][]string)
	testCases["-f=test"] = []string{"f", "test"}
	testCases["--f=test"] = []string{"f", "test"}
	testCases["--flag=test"] = []string{"flag", "test"}
	testCases["-flag=test"] = []string{"flag", "test"}
	testCases["----flag=--test"] = []string{"--flag", "--test"}

	for arg, correctValues := range testCases {
		key, value := parseArgWithValue(arg)
		if key != correctValues[0] {
			t.Fatalf("Flag %s parsed key as %s but expected key %s", arg, key, correctValues[0])
		}
		if value != correctValues[1] {
			t.Fatalf("Flag %s parsed value as %s but expected value %s", arg, value, correctValues[1])
		}
		t.Logf("Flag %s parsed key as %s and value as %s correctly", arg, key, value)
	}
}

func TestDetermineArgType(t *testing.T) {
	testCases := make(map[string]string)
	testCases["-f"] = ArgIsFlagWithSpace
	testCases["--f"] = ArgIsFlagWithSpace
	testCases["-flag"] = ArgIsFlagWithSpace
	testCases["--flag"] = ArgIsFlagWithSpace
	testCases["positionalArg"] = ArgIsPositional
	testCases["subcommand"] = ArgIsPositional
	testCases["sub--+/\\324command"] = ArgIsPositional
	testCases["--flag=CONTENT"] = ArgIsFlagWithValue
	testCases["-flag=CONTENT"] = ArgIsFlagWithValue
	testCases["-anotherfl-ag=CONTENT"] = ArgIsFlagWithValue
	testCases["--anotherfl-ag=CONTENT"] = ArgIsFlagWithValue
	testCases["1--anotherfl-ag=CONTENT"] = ArgIsPositional

	for arg, correctArgType := range testCases {
		argType := determineArgType(arg)
		if argType != correctArgType {
			t.Fatalf("Flag %s determined to be type %s but expected type %s", arg, argType, correctArgType)
		} else {
			t.Logf("Flag %s correctly determined to be type %s", arg, argType)
		}
	}
}