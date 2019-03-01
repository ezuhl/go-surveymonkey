package response

/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.txt', which is part of this source code package.
 */

type SurveyListQuestion struct {
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
	Data    []struct {
		Position int    `json:"position"`
		Href     string `json:"href"`
		Heading  string `json:"heading"`
		ID       string `json:"id"`
	} `json:"data"`
	Page  int `json:"page"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

type SurveyQuestion struct {
	Id            string              `json:"id"`
	Heading       []*SurveyHeading    `json:"headings"`
	Position      int                 `json:"position,omitempty"`
	Family        string              `json:"family"`
	Subtype       string              `json:"subtype"`
	Sorting       *Sorting            `json:"sorting,omitempty"`
	Required      *Required           `json:"required,omitempty"`
	Validation    *Validation         `json:"validation,omitempty"`
	ForcedRanking bool                `json:"forced_ranking,omitempty"`
	Visible       bool                `json:"visible"`
	Answers       AnswerTypeInterface `json:"answers"`
}

type SurveyHeading struct {
	Heading          string            `json:"heading"`
	Description      string            `json:"description,omitempty"`
	Image            *Image            `json:"image,omitempty"`
	RandomAssignment *RandomAssignment `json:"random_assignment,omitempty"`
}

type RandomAssignment struct {
	Percent      int    `json:"percent"`
	Position     int    `json:"position"`
	VariableName string `json:"variable_name"`
}

type Image struct {
	Url int `json:"url"`
}

type Sorting struct {
	Type       string `json:"type"` //default, textasc, textdesc, resp_count_asc, resp_count_desc, random, flip
	IgnoreLast bool   `json:"ignore_last"`
}

type Required struct {
	Text   string `json:"text"`
	Type   string `json:"type"` //all , at_least, at_most, exactly, or range
	Amount string `json:"amount"`
}

type Validation struct {
	Type    string `json:"type"` //any, integer, decimal, date_us, date_intl, regex, email, or text_length
	Text    string `json:"text"`
	Min     string `json:"min"`
	Max     string `json:"max"`
	Sum     int    `json:"sum"`
	SumText string `json:"sum_text"`
}

type QuizOptions struct {
	ScoringEnabled bool      `json:"scoring_enabled"`
	FeedBack       *FeedBack `json:"feedback"`
}

type FeedBack struct {
	CorrectText   string `json:"correct_text"`
	IncorrectText string `json:"incorrect_text"`
	PartialText   string `json:"partial_text"`
}

type AnswerTypeInterface interface {
}

type AnswerSinglechoice struct {
	AnswerTypeInterface
	Headings []struct {
		Heading string `json:"heading"`
	} `json:"headings"`
	Position int    `json:"position"`
	Family   string `json:"family"`
	Subtype  string `json:"subtype"`
	Answers  struct {
		Choices []struct {
			Text string `json:"text"`
		} `json:"choices"`
		Other []struct {
			Text     string `json:"text"`
			NumChars int    `json:"num_chars"`
			NumLines int    `json:"num_lines"`
		} `json:"other"`
	} `json:"answers"`
}

type AnswerMultipleChoice struct {
	AnswerTypeInterface
	Headings []struct {
		Heading string `json:"heading"`
	} `json:"headings"`
	Position int    `json:"position"`
	Family   string `json:"family"`
	Subtype  string `json:"subtype"`
	Answers  struct {
		Choices []struct {
			Text string `json:"text"`
		} `json:"choices"`
	} `json:"answers"`
}

type AnswerImageChoice struct {
	AnswerTypeInterface
	Headings []struct {
		Heading string `json:"heading"`
	} `json:"headings"`
	Position       int    `json:"position"`
	Family         string `json:"family"`
	Subtype        string `json:"subtype"`
	DisplayOptions struct {
		DisplayType string `json:"display_type"`
	} `json:"display_options"`
	Answers struct {
		Choices []struct {
			Text  string `json:"text"`
			Image struct {
				URL string `json:"url"`
			} `json:"image"`
		} `json:"choices"`
	} `json:"answers"`
}

type AnswerMatrixSingle struct {
	AnswerTypeInterface
	Headings []struct {
		Heading string `json:"heading"`
	} `json:"headings"`
	Position      int    `json:"position"`
	Family        string `json:"family"`
	Subtype       string `json:"subtype"`
	ForcedRanking bool   `json:"forced_ranking"`
	Answers       struct {
		Rows []struct {
			Text string `json:"text"`
		} `json:"rows"`
		Choices []struct {
			Text string `json:"text"`
		} `json:"choices"`
	} `json:"answers"`
}

type AnswerMatrixRating struct {
	AnswerTypeInterface
	Headings []struct {
		Heading string `json:"heading"`
	} `json:"headings"`
	Position      int    `json:"position"`
	Family        string `json:"family"`
	Subtype       string `json:"subtype"`
	ForcedRanking bool   `json:"forced_ranking"`
	Answers       struct {
		Rows []struct {
			Text string `json:"text"`
		} `json:"rows"`
		Choices []struct {
			Text   string `json:"text"`
			Weight int    `json:"weight"`
		} `json:"choices"`
	} `json:"answers"`
}

type AnswerMatrixRanking struct {
	AnswerTypeInterface
	Headings []struct {
		Heading string `json:"heading"`
	} `json:"headings"`
	Position int    `json:"position"`
	Family   string `json:"family"`
	Subtype  string `json:"subtype"`
	Answers  struct {
		Rows []struct {
			Text string `json:"text"`
		} `json:"rows"`
	} `json:"answers"`
}

type AnswerMatrixMenu struct {
	AnswerTypeInterface
	Headings []struct {
		Heading string `json:"heading"`
	} `json:"headings"`
	Position int    `json:"position"`
	Family   string `json:"family"`
	Subtype  string `json:"subtype"`
	Answers  struct {
		Rows []struct {
			Text string `json:"text"`
		} `json:"rows"`
		Cols []struct {
			Text    string `json:"text"`
			Choices []struct {
				Text     string `json:"text"`
				Visible  bool   `json:"visible"`
				Position int    `json:"position"`
			} `json:"choices"`
		} `json:"cols"`
	} `json:"answers"`
}

type AnswerOpenEnded struct {
	AnswerTypeInterface
	Headings []struct {
		Heading string `json:"heading"`
	} `json:"headings"`
	Position int    `json:"position"`
	Family   string `json:"family"`
	Subtype  string `json:"subtype"`
}

type AnswerOpenEndedMulti struct {
	AnswerTypeInterface
	Headings []struct {
		Heading string `json:"heading"`
	} `json:"headings"`
	Position int    `json:"position"`
	Family   string `json:"family"`
	Subtype  string `json:"subtype"`
	Answers  struct {
		Rows []struct {
			Text string `json:"text"`
		} `json:"rows"`
	} `json:"answers"`
}

type AnswerOpenEndedNumerical struct {
	AnswerTypeInterface
	Headings []struct {
		Heading string `json:"heading"`
	} `json:"headings"`
	Position int    `json:"position"`
	Family   string `json:"family"`
	Subtype  string `json:"subtype"`
	Answers  struct {
		Rows []struct {
			Text string `json:"text"`
		} `json:"rows"`
	} `json:"answers"`
	Validation struct {
		Text    string `json:"text"`
		Type    string `json:"type"`
		Sum     int    `json:"sum"`
		SumText string `json:"sum_text"`
	} `json:"validation"`
}

type AnswerDemographic struct {
	AnswerTypeInterface
	Headings []struct {
		Heading string `json:"heading"`
	} `json:"headings"`
	Position int    `json:"position"`
	Visible  bool   `json:"visible"`
	Family   string `json:"family"`
	Subtype  string `json:"subtype"`
	Answers  struct {
		Rows []struct {
			Visible  bool   `json:"visible"`
			Required bool   `json:"required"`
			Type     string `json:"type"`
			Text     string `json:"text,omitempty"`
		} `json:"rows"`
	} `json:"answers"`
}

type AnswerDateTime struct {
	AnswerTypeInterface
	Headings []struct {
		Heading string `json:"heading"`
	} `json:"headings"`
	Position int    `json:"position"`
	Visible  bool   `json:"visible"`
	Family   string `json:"family"`
	Subtype  string `json:"subtype"`
	Answers  struct {
		Rows []struct {
			Text     string `json:"text"`
			Position int    `json:"position"`
		} `json:"rows"`
	} `json:"answers"`
}

type AnswerPresentation struct {
	AnswerTypeInterface
	Headings []struct {
		Heading string `json:"heading"`
		Image   struct {
			URL string `json:"url"`
		} `json:"image"`
	} `json:"headings"`
	Position int    `json:"position"`
	Family   string `json:"family"`
	Subtype  string `json:"subtype"`
}

type AnswerFileUpload struct {
	AnswerTypeInterface
	Headings []struct {
		Heading string `json:"heading"`
	} `json:"headings"`
	Position       int    `json:"position"`
	Family         string `json:"family"`
	Subtype        string `json:"subtype"`
	DisplayOptions struct {
		DisplayType string `json:"display_type"`
	} `json:"display_options"`
}

type AnswerSlider struct {
	AnswerTypeInterface
	Headings []struct {
		Heading string `json:"heading"`
	} `json:"headings"`
	Position       int    `json:"position"`
	Family         string `json:"family"`
	Subtype        string `json:"subtype"`
	DisplayOptions struct {
		DisplayType string `json:"display_type"`
	} `json:"display_options"`
}

type AnswerSliderComplex struct {
	AnswerTypeInterface
	Headings []struct {
		Heading string `json:"heading"`
	} `json:"headings"`
	Position       int    `json:"position"`
	Family         string `json:"family"`
	Subtype        string `json:"subtype"`
	DisplayOptions struct {
		RightLabel    string `json:"right_label"`
		DisplayType   string `json:"display_type"`
		CustomOptions struct {
			StartingPosition int `json:"starting_position"`
			StepSize         int `json:"step_size"`
		} `json:"custom_options"`
		LeftLabel string `json:"left_label"`
	} `json:"display_options"`
	Validation struct {
		SumText string `json:"sum_text"`
		Min     int    `json:"min"`
		Text    string `json:"text"`
		Max     int    `json:"max"`
		Type    string `json:"type"`
	} `json:"validation"`
}

type AnswerEmoji struct {
	AnswerTypeInterface
	Headings []struct {
		Heading string `json:"heading"`
	} `json:"headings"`
	Position       int    `json:"position"`
	Family         string `json:"family"`
	Subtype        string `json:"subtype"`
	DisplayOptions struct {
		DisplayType    string `json:"display_type"`
		DisplaySubtype string `json:"display_subtype"`
	} `json:"display_options"`
	ForcedRanking bool `json:"forced_ranking"`
	Answers       struct {
		Rows []struct {
			Visible  bool   `json:"visible"`
			Text     string `json:"text"`
			Position int    `json:"position"`
		} `json:"rows"`
		Choices []struct {
			Weight int    `json:"weight"`
			Text   string `json:"text"`
		} `json:"choices"`
	} `json:"answers"`
}

type AnswerQuizQuestions struct {
	AnswerTypeInterface
	Headings []struct {
		Heading string `json:"heading"`
	} `json:"headings"`
	Position int    `json:"position"`
	Family   string `json:"family"`
	Subtype  string `json:"subtype"`
	Answers  struct {
		Choices []struct {
			Text        string `json:"text"`
			QuizOptions struct {
				Score int `json:"score"`
			} `json:"quiz_options"`
		} `json:"choices"`
	} `json:"answers"`
	QuizOptions struct {
		ScoringEnabled bool `json:"scoring_enabled"`
		Feedback       struct {
			CorrectText   string `json:"correct_text"`
			IncorrectText string `json:"incorrect_text"`
		} `json:"feedback"`
	} `json:"quiz_options"`
}

type AnswerQuestionsBankQuestions struct {
	AnswerTypeInterface
	QuestionBank struct {
		QuestionBankQuestionID string `json:"question_bank_question_id"`
		ModifierOptions        struct {
			Num36628 interface{} `json:"36628"`
		} `json:"modifier_options"`
	} `json:"question_bank"`
}
