package calculations

func IsGrandPrixKnown(racename string) bool {
	return racename != "" && racename != "Unknown"
}
