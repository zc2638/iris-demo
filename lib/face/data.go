package face

type GetSetResult struct {
	TimeUsed int `json:"time_used"`
	Facesets []struct {
		FacesetToken string `json:"faceset_token"`
		OuterID      string `json:"outer_id"`
		DisplayName  string `json:"display_name"`
		Tags         string `json:"tags"`
	} `json:"facesets"`
	RequestID    string `json:"request_id"`
	ErrorMessage string `json:"error_message"`
}

type CreateSetData struct {
	ApiKey      string `json:"api_key"`
	ApiSecret   string `json:"api_secret"`
	DisplayName string `json:"display_name"`
	OuterID     string `json:"outer_id"`
}

type CreateSetResult struct {
	FacesetToken  string `json:"faceset_token"`
	TimeUsed      int    `json:"time_used"`
	FaceCount     int    `json:"face_count"`
	FaceAdded     int    `json:"face_added"`
	RequestID     string `json:"request_id"`
	OuterID       string `json:"outer_id"`
	FailureDetail []struct {
		Reason    string `json:"reason"`
		FaceToken string `json:"face_token"`
	} `json:"failure_detail"`
	ErrorMessage string `json:"error_message"`
}

type DetectImageData struct {
	ApiKey      string `json:"api_key"`
	ApiSecret   string `json:"api_secret"`
	ImageBase64 string `json:"image_base64"`
	ImageUrl    string `json:"image_url"`
}

// 如果没有检测出人脸 faces 为空
type DetectImageResult struct {
	ImageID   string `json:"image_id"`
	RequestID string `json:"request_id"`
	TimeUsed  int    `json:"time_used"`
	Faces     []struct {
		Landmark struct {
			MouthUpperLipLeftContour2 struct {
				Y int `json:"y"`
				X int `json:"x"`
			} `json:"mouth_upper_lip_left_contour2"`
			ContourChin struct {
				Y int `json:"y"`
				X int `json:"x"`
			} `json:"contour_chin"`
			RightEyePupil struct {
				Y int `json:"y"`
				X int `json:"x"`
			} `json:"right_eye_pupil"`
			MouthUpperLipBottom struct {
				Y int `json:"y"`
				X int `json:"x"`
			} `json:"mouth_upper_lip_bottom"`
		} `json:"landmark"`
		Attributes struct {
			Gender struct {
				Value string `json:"value"`
			} `json:"gender"`
			Age struct {
				Value int `json:"value"`
			} `json:"age"`
			Glass struct {
				Value string `json:"value"`
			} `json:"glass"`
			Headpose struct {
				YawAngle   float64 `json:"yaw_angle"`
				PitchAngle float64 `json:"pitch_angle"`
				RollAngle  float64 `json:"roll_angle"`
			} `json:"headpose"`
			Smile struct {
				Threshold float64 `json:"threshold"`
				Value     float64 `json:"value"`
			} `json:"smile"`
		} `json:"attributes"`
		FaceRectangle struct {
			Width  int `json:"width"`
			Top    int `json:"top"`
			Left   int `json:"left"`
			Height int `json:"height"`
		} `json:"face_rectangle"`
		FaceToken string `json:"face_token"`
	} `json:"faces"`
	ErrorMessage string `json:"error_message"`
}

type AddFaceData struct {
	ApiKey     string `json:"api_key"`
	ApiSecret  string `json:"api_secret"`
	OuterID    string `json:"outer_id"`
	FaceTokens string `json:"face_tokens"`
}

type AddFaceResult struct {
	FacesetToken  string `json:"faceset_token"`
	TimeUsed      int    `json:"time_used"`
	FaceCount     int    `json:"face_count"`
	FaceAdded     int    `json:"face_added"`
	RequestID     string `json:"request_id"`
	OuterID       string `json:"outer_id"`
	FailureDetail []struct {
		Reason    string `json:"reason"`
		FaceToken string `json:"face_token"`
	} `json:"failure_detail"`
	ErrorMessage string `json:"error_message"`
}

type SearchData struct {
	ApiKey            string `json:"api_key"`
	ApiSecret         string `json:"api_secret"`
	OuterID           string `json:"outer_id"`
	ReturnResultCount int    `json:"return_result_count"`
	ImageBase64       string `json:"image_base64"`
	ImageUrl          string `json:"image_url"`
}

type SearchResult struct {
	RequestID  string `json:"request_id"`
	TimeUsed   int    `json:"time_used"`
	Thresholds struct {
		OneE3 float64 `json:"1e-3"`
		OneE5 float64 `json:"1e-5"`
		OneE4 float64 `json:"1e-4"`
	} `json:"thresholds"`
	Results []struct {
		Confidence float64 `json:"confidence"`
		UserID     string  `json:"user_id"`
		FaceToken  string  `json:"face_token"`
	} `json:"results"`
	ErrorMessage string `json:"error_message"`
}
