// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package efficiency_measure

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/erda-project/erda/internal/core/legacy/model"
)

var (
	allPersonalMetrics = []personalMetric{
		{
			name:      "personal_task_total",
			help:      "Total number of tasks",
			valueType: prometheus.CounterValue,
			getValues: func(personalInfo *PersonalPerformanceInfo) metricValues {
				var value float64
				if personalInfo.metricFields != nil {
					value = float64(personalInfo.metricFields.TaskTotal)
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
		{
			name:      "personal_working_task_total",
			help:      "Total number of working tasks",
			valueType: prometheus.CounterValue,
			getValues: func(personalInfo *PersonalPerformanceInfo) metricValues {
				var value float64
				if personalInfo.metricFields != nil {
					value = float64(personalInfo.metricFields.WorkingTaskTotal)
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
		{
			name:      "personal_pending_task_total",
			help:      "Total number of pending tasks",
			valueType: prometheus.CounterValue,
			getValues: func(personalInfo *PersonalPerformanceInfo) metricValues {
				var value float64
				if personalInfo.metricFields != nil {
					value = float64(personalInfo.metricFields.PendingTaskTotal)
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
		{
			name:      "personal_requirement_total",
			help:      "Total number of requirements that assignee is the user",
			valueType: prometheus.CounterValue,
			getValues: func(personalInfo *PersonalPerformanceInfo) metricValues {
				var value float64
				if personalInfo.metricFields != nil {
					value = float64(personalInfo.metricFields.RequirementTotal)
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
		{
			name:      "personal_working_requirement_total",
			help:      "Total number of working requirements",
			valueType: prometheus.CounterValue,
			getValues: func(personalInfo *PersonalPerformanceInfo) metricValues {
				var value float64
				if personalInfo.metricFields != nil {
					value = float64(personalInfo.metricFields.WorkingRequirementTotal)
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
		{
			name:      "personal_pending_requirement_total",
			help:      "Total number of pending requirement",
			valueType: prometheus.CounterValue,
			getValues: func(personalInfo *PersonalPerformanceInfo) metricValues {
				var value float64
				if personalInfo.metricFields != nil {
					value = float64(personalInfo.metricFields.PendingRequirementTotal)
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
		{
			name:      "personal_bug_total",
			help:      "Total number of bug that assignee is the user",
			valueType: prometheus.CounterValue,
			getValues: func(personalInfo *PersonalPerformanceInfo) metricValues {
				var value float64
				if personalInfo.metricFields != nil {
					value = float64(personalInfo.metricFields.BugTotal)
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
		{
			name:      "personal_owner_bug_total",
			help:      "Total number of bug that owner is the user",
			valueType: prometheus.CounterValue,
			getValues: func(personalInfo *PersonalPerformanceInfo) metricValues {
				var value float64
				if personalInfo.metricFields != nil {
					value = float64(personalInfo.metricFields.OwnerBugTotal)
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
		{
			name:      "personal_pending_bug_total",
			help:      "Total number of pending bugs",
			valueType: prometheus.CounterValue,
			getValues: func(personalInfo *PersonalPerformanceInfo) metricValues {
				var value float64
				if personalInfo.metricFields != nil {
					value = float64(personalInfo.metricFields.PendingBugTotal)
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
		{
			name:      "personal_working_bug_total",
			help:      "Total number of working bugs",
			valueType: prometheus.CounterValue,
			getValues: func(personalInfo *PersonalPerformanceInfo) metricValues {
				var value float64
				if personalInfo.metricFields != nil {
					value = float64(personalInfo.metricFields.WorkingBugTotal)
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
		{
			name:      "personal_demand_design_bug_total",
			help:      "Total number of personal demand design bugs that don't contain wontfix state",
			valueType: prometheus.CounterValue,
			getValues: func(p *PersonalPerformanceInfo) metricValues {
				var value float64
				if p.metricFields != nil {
					value = float64(p.metricFields.DemandDesignBugTotal)
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
		{
			name:      "personal_architecture_design_bug_total",
			help:      "Total number of personal architecture design bugs that don't contain wontfix state",
			valueType: prometheus.CounterValue,
			getValues: func(p *PersonalPerformanceInfo) metricValues {
				var value float64
				if p.metricFields != nil {
					value = float64(p.metricFields.ArchitectureDesignBugTotal)
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
		{
			name:      "personal_serious_bug_total",
			help:      "Total number of personal serious bugs that don't contain wontfix state",
			valueType: prometheus.CounterValue,
			getValues: func(p *PersonalPerformanceInfo) metricValues {
				var value float64
				if p.metricFields != nil {
					value = float64(p.metricFields.SeriousBugTotal)
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
		{
			name:      "personal_reopen_bug_total",
			help:      "Total number of personal reopen bugs that don't contain wontfix state",
			valueType: prometheus.CounterValue,
			getValues: func(p *PersonalPerformanceInfo) metricValues {
				var value float64
				if p.metricFields != nil {
					value = float64(p.metricFields.ReopenBugTotal)
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
		{
			name:      "personal_submit_bug_total",
			help:      "Total number of personal submitted bugs that don't contain wontfix state",
			valueType: prometheus.CounterValue,
			getValues: func(p *PersonalPerformanceInfo) metricValues {
				var value float64
				if p.metricFields != nil {
					value = float64(p.metricFields.SubmitBugTotal)
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
		{
			name:      "personal_test_case_total",
			help:      "Total number of test cases created",
			valueType: prometheus.CounterValue,
			getValues: func(p *PersonalPerformanceInfo) metricValues {
				var value float64
				if p.metricFields != nil {
					value = float64(p.metricFields.CreateTestCaseTotal)
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
		{
			name:      "personal_fix_bug_elapsed_minute",
			help:      "Average time spent fixing  bugs",
			valueType: prometheus.CounterValue,
			getValues: func(p *PersonalPerformanceInfo) metricValues {
				var value float64
				if p.metricFields != nil {
					value = p.metricFields.AvgFixBugElapsedMinute
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
		{
			name:      "personal_fix_bug_estimate_minute",
			help:      "Average time estimate fixing bugs",
			valueType: prometheus.CounterValue,
			getValues: func(p *PersonalPerformanceInfo) metricValues {
				var value float64
				if p.metricFields != nil {
					value = p.metricFields.AvgFixBugEstimateMinute
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
		{
			name:      "personal_fix_bug_estimate_minute_total",
			help:      "total time estimate fixing bugs",
			valueType: prometheus.CounterValue,
			getValues: func(p *PersonalPerformanceInfo) metricValues {
				var value float64
				if p.metricFields != nil {
					value = p.metricFields.TotalFixBugEstimateMinute
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
		{
			name:      "personal_fix_bug_elapsed_minute_total",
			help:      "total time elapsed fixing bugs",
			valueType: prometheus.CounterValue,
			getValues: func(p *PersonalPerformanceInfo) metricValues {
				var value float64
				if p.metricFields != nil {
					value = p.metricFields.TotalFixFixBugElapsedMinute
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
		{
			name:      "personal_resolved_bug_total",
			help:      "total resolved bug total",
			valueType: prometheus.CounterValue,
			getValues: func(p *PersonalPerformanceInfo) metricValues {
				var value float64
				if p.metricFields != nil {
					value = p.metricFields.ResolvedBugTotal
				}
				return metricValues{
					{
						value:     value,
						timestamp: time.Now(),
					},
				}
			},
		},
	}
)

type metricValue struct {
	value     float64
	labels    []string
	timestamp time.Time
}

type metricValues []metricValue

type personalMetric struct {
	name        string
	help        string
	valueType   prometheus.ValueType
	extraLabels []string
	getValues   func(p *PersonalPerformanceInfo) metricValues
}

func (pm *personalMetric) desc(baseLabels []string) *prometheus.Desc {
	return prometheus.NewDesc(pm.name, pm.help, append(baseLabels, pm.extraLabels...), nil)
}

type PersonalPerformanceInfo struct {
	userProject *model.UserJoinedProject

	metricFields *personalMetricField
}

type personalMetricField struct {
	CalculatedAt time.Time

	TaskTotal               uint64
	WorkingTaskTotal        uint64
	PendingTaskTotal        uint64
	RequirementTotal        uint64
	WorkingRequirementTotal uint64
	PendingRequirementTotal uint64
	BugTotal                uint64
	OwnerBugTotal           uint64
	WorkingBugTotal         uint64
	PendingBugTotal         uint64

	SeriousBugTotal             uint64
	DemandDesignBugTotal        uint64
	ArchitectureDesignBugTotal  uint64
	ReopenBugTotal              uint64
	SubmitBugTotal              uint64
	AvgFixBugElapsedMinute      float64
	AvgFixBugEstimateMinute     float64
	ResolvedBugTotal            float64
	TotalFixFixBugElapsedMinute float64
	TotalFixBugEstimateMinute   float64

	CreateTestCaseTotal uint64
}

// IsValid returns true if the personalMetricField is valid.
// we need to ensure that personalMetricField is the data of the day to avoid double calculation
func (p *personalMetricField) IsValid() bool {
	return p.CalculatedAt.Day() == time.Now().Day()
}

func (p *provider) personalLabelsFunc(info *PersonalPerformanceInfo) map[string]string {
	labels := map[string]string{
		"project_id":           strconv.FormatUint(info.userProject.ProjectID, 10),
		"user_id":              strconv.FormatUint(info.userProject.UserID, 10),
		"org_id":               strconv.FormatUint(info.userProject.OrgID, 10),
		"user_name":            info.userProject.UserName,
		"user_nickname":        info.userProject.UserNickName,
		"user_email":           info.userProject.UserEmail,
		"project_name":         info.userProject.ProjectName,
		"project_display_name": info.userProject.ProjectDisplayName,
		"org_name":             info.userProject.OrgName,
		"org_display_name":     info.userProject.OrgDisplayName,
	}
	projectLabels := strings.Split(info.userProject.ProjectLabels, ",")
	for _, label := range projectLabels {
		kvs := strings.Split(label, ":")
		if len(kvs) == 2 {
			labels[sanitizeLabelName(kvs[0])] = kvs[1]
		}
	}
	return labels
}

var invalidNameCharRE = regexp.MustCompile(`[^a-zA-Z0-9_]`)

func sanitizeLabelName(name string) string {
	return invalidNameCharRE.ReplaceAllString(name, "_")
}

func (p *provider) Collect(ch chan<- prometheus.Metric) {
	p.errors.Set(0)
	personalInfos, err := p.GetRequestedPersonalInfos()
	if err != nil {
		p.errors.Set(1)
		p.Log.Errorf("failed to get requested personal infos, err: %v", err)
	}
	for _, personalInfo := range personalInfos {
		rawLabels := map[string]struct{}{}
		for l := range p.personalLabelsFunc(personalInfo) {
			rawLabels[l] = struct{}{}
		}
		values := make([]string, 0, len(rawLabels))
		labels := make([]string, 0, len(rawLabels))
		personalLabels := p.personalLabelsFunc(personalInfo)
		for l := range rawLabels {
			duplicate := false
			for _, x := range labels {
				if l == x {
					duplicate = true
					break
				}
			}
			if !duplicate {
				labels = append(labels, l)
				values = append(values, personalLabels[l])
			}
		}
		for _, pm := range allPersonalMetrics {
			desc := pm.desc(labels)
			for _, metricVal := range pm.getValues(personalInfo) {
				ch <- prometheus.NewMetricWithTimestamp(
					metricVal.timestamp,
					prometheus.MustNewConstMetric(desc, pm.valueType, metricVal.value, append(values, metricVal.labels...)...),
				)
			}
		}
	}
	p.errors.Collect(ch)
}

func (p *provider) Describe(ch chan<- *prometheus.Desc) {
	p.errors.Describe(ch)
	for _, pm := range allPersonalMetrics {
		ch <- pm.desc([]string{})
	}
}
