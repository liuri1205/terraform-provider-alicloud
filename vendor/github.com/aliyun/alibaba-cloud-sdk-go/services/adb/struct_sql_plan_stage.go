package adb

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// SqlPlanStage is a nested struct in adb response
type SqlPlanStage struct {
	StageId      int    `json:"StageId" xml:"StageId"`
	State        string `json:"State" xml:"State"`
	OperatorCost int64  `json:"OperatorCost" xml:"OperatorCost"`
	PeakMemory   int64  `json:"PeakMemory" xml:"PeakMemory"`
	CPUTimeMin   int64  `json:"CPUTimeMin" xml:"CPUTimeMin"`
	CPUTimeMax   int64  `json:"CPUTimeMax" xml:"CPUTimeMax"`
	CPUTimeAvg   int64  `json:"CPUTimeAvg" xml:"CPUTimeAvg"`
	InputSizeMin int64  `json:"InputSizeMin" xml:"InputSizeMin"`
	InputSizeMax int64  `json:"InputSizeMax" xml:"InputSizeMax"`
	InputSizeAvg int64  `json:"InputSizeAvg" xml:"InputSizeAvg"`
	ScanSizeMin  int64  `json:"ScanSizeMin" xml:"ScanSizeMin"`
	ScanSizeMax  int64  `json:"ScanSizeMax" xml:"ScanSizeMax"`
	ScanSizeAvg  int64  `json:"ScanSizeAvg" xml:"ScanSizeAvg"`
	ScanTimeMin  int64  `json:"ScanTimeMin" xml:"ScanTimeMin"`
	ScanTimeMax  int64  `json:"ScanTimeMax" xml:"ScanTimeMax"`
	ScanTimeAvg  int64  `json:"ScanTimeAvg" xml:"ScanTimeAvg"`
}
