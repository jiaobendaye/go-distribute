package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "Nick",
			LastName:  "Carter",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 34,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 84,
				},
			},
		},
		{
			ID:        2,
			FirstName: "Fan",
			LastName:  "XiaoHao",
			Grades: []Grade{
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 89,
				},
				{
					Title: "Quiz 3",
					Type:  GradeQuiz,
					Score: 74,
				},
			},
		},
	}
}
