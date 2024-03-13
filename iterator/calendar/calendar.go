package calendar

import "fmt"

type MonthList struct {
	monthList []string
	index     int
}

func NewMonthList() *MonthList {
	return &MonthList{
		monthList: []string{"Jan", "Feb", "Mar"},
		index:     0,
	}
}

func (m *MonthList) HasNext() bool {
	return (m.index < len(m.monthList))
}

func (m *MonthList) Next() string {
	value := m.monthList[m.index]
	m.index += 1
	return value
}

type CustomList struct {
	index int
}

func NewCustomList() *CustomList {
	return &CustomList{
		index: 0,
	}
}

func (m *CustomList) HasNext() bool {
	return m.index < 5
}

func (m *CustomList) Next() string {
	return fmt.Sprintf("%d", m.index*m.index)
}
