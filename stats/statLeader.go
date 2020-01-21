package stats

import (
	"strings"

	"github.com/project-0/config"
	"github.com/project-0/dataframe"
)

func StatLeader(df *dataframe.Dataframe, flag string) string {
	for i := 1; i < len(df.Data); i++ {
		df.Data[i][0] = strings.Split(df.Data[i][0], "\\")[0]
	}
	header, _ := df.Data[0]
	stat := flag
	var statIdx int
	var found bool
	for i := 0; i < len(header); i++ {
		if header[i] == stat {
			statIdx = i
			found = true
			break
		}
	}
	if !found {
		return ("Couldn't find that stat")
	}
	df.Sort(statIdx)
	return df.PrettyString()
}

func StatLeaderJSON(flag string) string {
	df := dataframe.ReadCSV(config.FILE)
	df.DropCol(0)
	for i := 1; i < len(df.Data); i++ {
		df.Data[i][0] = strings.Split(df.Data[i][0], "\\")[0]
	}
	header, _ := df.Data[0]
	stat := flag
	var statIdx int
	var found bool
	for i := 0; i < len(header); i++ {
		if header[i] == stat {
			statIdx = i
			found = true
			break
		}
	}
	if !found {
		return ("Couldn't find that stat")
	}
	df.Sort(statIdx)
	return df.ToJSON()
}
