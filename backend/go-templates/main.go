package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	// "io"
	"bytes"

	"github.com/big-larry/mgo/bson"
)

type CourseMeta struct {
	Id        bson.ObjectId `bson:"_id"       json:"id,omitempty"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
	Title     string        `bson:"title"     json:"title"`
	Author    string        `bson:"author"    json:"author,omitempty"`
	ShortCard ShortCard     `bson:"shortCard" json:"shortCard"`
	FullCard  FullCard      `bson:"fullCard"  json:"fullCard"`
	Shown     bool          `bson:"shown"     json:"shown"`
	Deleted   bool          `bson:"deleted"   json:"deleted,omitempty"`
	Lessons   []Lesson
	Tests     map[bson.ObjectId]Test
}

type ShortCard struct {
	CourseId    bson.ObjectId `bson:"_id"            json:"courseId"`
	Description string        `bson:"description"    json:"description"`
	Cover       interface{}   `bson:"cover"          json:"cover"`
	City        string        `bson:"city,omitempty" json:"city,omitempty"`
	Time        string        `bson:"time,omitempty" json:"time,omitempty"`
	Type        string        `bson:"type,omitempty" json:"type,omitempty"`
}

type FullCard struct {
	CourseId          bson.ObjectId `bson:"_id" json:"courseId"`
	HeadDescription   string        `bson:"headDescription" json:"headDescription"`
	EducationForm     string        `bson:"educationForm"   json:"educationForm"`
	EndDocument       string        `bson:"endDocument"     json:"endDocument"`
	AvailableMaterial struct {
		Modules       uint `bson:"modules"       json:"modules"`
		TheoryHours   uint `bson:"theoryHours"   json:"theoryHours"`
		PracticeHours uint `bson:"practiceHours" json:"practiceHours"`
		Weeks         uint `bson:"weeks"         json:"weeks"`
	} `bson:"availableMaterial" json:"availableMaterial"`
	KnowledgeBlock struct {
		Header    string `bson:"header" json:"header"`
		Knowledge []struct {
			Head string `bson:"head" json:"head"`
			Body string `bson:"body" json:"body"`
		} `bson:"knowledge" json:"knowledge"`
	} `bson:"knowledgeBlock" json:"knowledgeBlock"`
	ForComfortBlock struct {
		Head       string `bson:"head" json:"head"`
		Instrument []struct {
			Icon interface{} `bson:"icon" json:"icon"`
			Text string      `bson:"text" json:"text"`
		} `bson:"instruments" json:"instruments"`
	} `bson:"forComfortBlock" json:"forComfortBlock"`
	Faq []struct {
		Question string `bson:"question" json:"question"`
		Answer   string `bson:"answer"   json:"answer"`
	} `bson:"faq" json:"faq"`
}

type Course struct {
	Days     string
	Location string
	Image    string
	Author   string
	Title    string
	Desc     string
	Link     string
	Show     string
	Deleted  bool
}

type MainPageData struct {
	BaseData
	HeaderType string
	ShowFooter bool
}

type SettingPageData struct {
	BaseData
	HeaderType string
	ShowFooter bool
	User       User
}

type PageDataShortCard struct {
	BaseData
	HeaderType string
	Courses    []CourseView
	ShowFooter bool
}

type PageDataCours struct {
	BaseData
	HeaderType     string
	Courses        CourseMeta
	CorrectAnswers map[string]interface{}
	ShowFooter     bool
}

type PageData struct {
	BaseData
	HeaderType string
	Courses    []Course
}

type Chat struct {
	Id      bson.ObjectId `bson:"_id"     json:"id,omitempty"`
	Title   string        `bson:"title"   json:"title"`
	Members []int         `bson:"members" json:"members"`
	IsGroup bool          `bson:"isGroup" json:"isGroup"`
	Deleted bool          `bson:"deleted" json:"deleted,omitempty"`
}

type Message struct {
	Id          bson.ObjectId `bson:"_id"                   json:"id,omitempty"`
	Chat        bson.ObjectId `bson:"chatId"                json:"chatId"`
	From        int           `bson:"from"                  json:"from"`
	Time        time.Time     `bson:"time"                  json:"time"`
	Text        string        `bson:"text,omitempty"        json:"text,omitempty"`
	Attachments []interface{} `bson:"attachments,omitempty" json:"attachments,omitmepty"`
	Deleted     bool          `bson:"deleted"               json:"deleted,omitempty"`
}

type ChatPageData struct {
	BaseData
	HeaderType string
	Chats      []ChatWithLastMsg
	ShowFooter bool
}

type ChatWithLastMsg struct {
	Chat        Chat
	LastMessage *Message
}

type Courses struct {
	CourseId bson.ObjectId `bson:"_id"     json:"courseId,omitempty"`
	Lessons  []Lesson      `bson:"lessons" json:"lessons"`
}

type Lesson struct {
	CourseId bson.ObjectId `bson:"courseId"         json:"courseId"`
	Type     string        `bson:"type"             json:"type"`
	Info     []Info        `bson:"info,omitempty"   json:"info,omitempty"`
	TestId   bson.ObjectId `bson:"testId,omitempty" json:"testId,omitempty"`
}

type Info struct {
	Type string `bson:"type" json:"type"`
	Data string `bson:"data" json:"data"`
}

type Test struct {
	Id        bson.ObjectId `bson:"_id"                   json:"id,omitempty"`
	Title     string        `bson:"title"                 json:"title"`
	Questions []Question    `bson:"questions"             json:"questions"`
	Deleted   bool          `bson:"deleted,omitempty"     json:"deleted,omitempty"`
}

type Question struct {
	Type     string                   `bson:"type"               json:"type"`
	Info     interface{}              `bson:"info"               json:"info"`
	Variants []map[string]interface{} `bson:"variants,omitempty" json:"variants,omitempty"`
	Answer   interface{}              `bson:"answer"             json:"answer"`
}

type TestUser struct {
	Id       bson.ObjectId `bson:"_id"      json:"id,omitempty"`
	UserId   int           `bson:"userId"   json:"userId"`
	TestId   bson.ObjectId `bson:"testId"   json:"testId"`
	Progress Progress      `bson:"progress" json:"progress"`
}

type Progress struct {
	Opened   bool `bson:"opened"   json:"opened"`
	Stage    int  `bson:"stage"    json:"stage"`
	Finished bool `bson:"finished" json:"finished"`
	Correct  int  `bson:"correct"  json:"correct"`
}

type CourseUser struct {
	Id        bson.ObjectId  `bson:"_id"               json:"id,omitempty"`
	UserId    int            `bson:"userId"            json:"userId"`
	CourseId  bson.ObjectId  `bson:"courseId"          json:"courseId"`
	Purchased bool           `bson:"purchased"         json:"purchased"`
	Progress  ProgressCourse `bson:"progress"          json:"progress"`
}

type ProgressCourse struct {
	Opened   bool `bson:"opened"   json:"opened"`
	Stage    int  `bson:"stage"    json:"stage"`
	Finished bool `bson:"finished" json:"finished"`
}

type CourseView struct {
	CourseMeta
	Started   bool
	Completed bool
}

type BaseData struct {
	IsAuthorized bool
	UserName     string
	UserId       int
}

type User struct {
	Id          int
	Name        string
	Email       string
	Unsubscribe bool
	Perm        int
}

func buildBaseData(r *http.Request) BaseData {
	userId := -1
	userName := ""
	isAuth := false

	cookies := r.CookiesNamed("token")
	for _, cookie := range cookies {
		if cookie.Value == "" {
			continue
		}
		isAuth = true
		url := fmt.Sprintf("http://localhost:8085/getMe")
		req, _ := http.NewRequest("GET", url, nil)
		req.AddCookie(cookie)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			var user User
			if json.NewDecoder(resp.Body).Decode(&user) == nil {
				userName = user.Name
			}
		}

	}

	return BaseData{
		IsAuthorized: isAuth,
		UserName:     userName,
		UserId:       userId,
	}
}

func main() {
	// обработка запроса
	http.HandleFunc("/courses", courseHandler)
	http.HandleFunc("/cours", coursHandler)
	http.HandleFunc("/chat", chatHandler)
	http.HandleFunc("/neurokids", mainPageHandler)
	http.HandleFunc("/diagnostic", diagnosticPageHandler)
	http.HandleFunc("/setting", settingPageHandler)
	http.HandleFunc("/userList", userListPageHandler)
	http.HandleFunc("/userCreate", userCreatePageHandler)
	http.HandleFunc("/learn", learnPageHandler)
	http.HandleFunc("/contacts", contactsPageHandler)
	http.HandleFunc("/dgComplecs", diagnosticInfoPageHandler)

	// отдача статики
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":6670", nil)
}

func coursHandler(w http.ResponseWriter, r *http.Request) {
	var cours CourseMeta
	var lesson Courses

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is missing", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("http://localhost:8085/getFullCourseMetaById?id=%s", id)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch courses", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&cours)
	if err != nil {
		http.Error(w, "Failed to decode course", http.StatusInternalServerError)
		return
	}

	url = fmt.Sprintf("http://localhost:8085//getCourseById?id=%s", id)
	resp, err = http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch courses", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&lesson)
	if err != nil {
		http.Error(w, "Failed to decode course", http.StatusInternalServerError)
		return
	}

	tests := make(map[bson.ObjectId]Test)

	for _, l := range lesson.Lessons {
		if l.Type == "test" && l.TestId.Valid() {
			testUrl := fmt.Sprintf("http://localhost:8085/getTestById?id=%s", l.TestId.Hex())
			resp, err := http.Get(testUrl)
			if err != nil {
				log.Printf("Failed to fetch test %s: %v", l.TestId.Hex(), err)
				continue
			}
			defer resp.Body.Close()

			var test Test
			err = json.NewDecoder(resp.Body).Decode(&test)
			if err != nil {
				log.Printf("Failed to decode test %s: %v", l.TestId.Hex(), err)
				continue
			}

			tests[l.TestId] = test
		}
	}

	correctAnswers := make(map[string]interface{})

	for testId, test := range tests {
		for i, q := range test.Questions {
			key := fmt.Sprintf("%s_q%d", testId.Hex(), i)
			correctAnswers[key] = q.Answer
		}
	}

	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"add1": func(i int) int { return i + 1 },
		"toJson": func(v interface{}) template.JS {
			b, err := json.Marshal(v)
			if err != nil {
				return template.JS("{}") // возвращаем пустой объект на ошибке
			}
			return template.JS(b) // безопасно вставляем JSON в JS
		},
	}
	tmpl := template.New("").Funcs(funcMap)
	tmpl = template.Must(tmpl.ParseFiles(
		"templates/index.html",
		"templates/header.html",
		"templates/footer.html",
		"templates/cours.html",
	))

	// fmt.Printf("CORRECT ANSWERS:\n%+v\n", correctAnswers)

	cours.Lessons = lesson.Lessons
	cours.Tests = tests
	data := PageDataCours{
		BaseData:       buildBaseData(r),
		HeaderType:     "header", // или "headerDefault"
		Courses:        cours,
		CorrectAnswers: correctAnswers,
		ShowFooter:     true,
	}

	// Выполняем шаблон и проверяем ошибку
	err = tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		// Логируем ошибку, но не вызываем http.Error, так как ответ уже мог быть частично записан
		log.Printf("Failed to execute template: %v", err)
		return
	}

	getURL := fmt.Sprintf("http://localhost:8085/getCourseUserByCourseId?courseId=%s", id)
	req, err := http.NewRequest("GET", getURL, nil)
	if err != nil {
		panic(err)
	}

	uidStr := r.Header.Get("x-user-id")
	var uid int
	if uid, err = strconv.Atoi(uidStr); err != nil {
		log.Println("Cannot get x-user-id")
		return
	}

	cookies := r.CookiesNamed("token")
	if len(cookies) < 1 || cookies[0].Value == "" {
		log.Println("Cannot get token")
		return
	}
	req.AddCookie(cookies[0])

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		// 404 — создаём
		fmt.Println("Пользователь не найден — создаём...")

		payload := map[string]interface{}{
			"userId":    uid,
			"courseId":  id,
			"purchased": true,
		}
		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			panic(err)
		}

		createURL := "http://localhost:8085/createCourseUser"
		postReq, err := http.NewRequest("POST", createURL, bytes.NewBuffer(jsonPayload))
		if err != nil {
			panic(err)
		}
		postReq.Header.Set("Content-Type", "application/json")
		postReq.AddCookie(cookies[0])

		postResp, err := client.Do(postReq)
		if err != nil {
			panic(err)
		}
		defer postResp.Body.Close()

		// if postResp.StatusCode == http.StatusOK {
		// 	fmt.Println("✅ Успешно создано.")
		// } else {
		// 	fmt.Println("❌ Ошибка при создании:", postResp.Status)
		// }
	}
	// else if resp.StatusCode != http.StatusOK {
	// 	body, _ := io.ReadAll(resp.Body)
	// }
	// else {
	// 	// Всё хорошо, пользователь найден
	// 	fmt.Println("Пользователь уже существует.")
	// }
	for _, lesson := range lesson.Lessons {
		if lesson.Type == "test" && lesson.TestId.Valid() {
			testId := lesson.TestId.Hex()

			getURL := fmt.Sprintf("http://localhost:8085/getTestUserByTestId?testId=%s", testId)
			req, err := http.NewRequest("GET", getURL, nil)
			if err != nil {
				log.Printf("Ошибка при создании запроса на TestUser: %v", err)
				continue
			}
			req.AddCookie(cookies[0])

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Printf("Ошибка при выполнении запроса на TestUser: %v", err)
				continue
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusNotFound {
				// TestUser не найден — создаём
				log.Printf("TestUser для теста %s не найден, создаём...", testId)

				payload := map[string]interface{}{
					"userId": uid,
					"testId": testId,
				}

				jsonPayload, err := json.Marshal(payload)
				if err != nil {
					log.Printf("Ошибка при сериализации payload для TestUser: %v", err)
					continue
				}

				createURL := "http://localhost:8085/createTestUser"
				postReq, err := http.NewRequest("POST", createURL, bytes.NewBuffer(jsonPayload))
				if err != nil {
					log.Printf("Ошибка при создании POST-запроса на TestUser: %v", err)
					continue
				}
				postReq.Header.Set("Content-Type", "application/json")
				postReq.AddCookie(cookies[0])

				postResp, err := client.Do(postReq)
				if err != nil {
					log.Printf("Ошибка при отправке POST-запроса на TestUser: %v", err)
					continue
				}
				defer postResp.Body.Close()

				// if postResp.StatusCode == http.StatusOK {
				// 	log.Printf("✅ TestUser для теста %s успешно создан.", testId)
				// } else {
				// 	body, _ := io.ReadAll(postResp.Body)
				// 	log.Printf("❌ Ошибка при создании TestUser: %s\nОтвет: %s", postResp.Status, string(body))
				// }
			}
			// else if resp.StatusCode != http.StatusOK {
			// 	body, _ := io.ReadAll(resp.Body)
			//  log.Printf("⚠️ Ошибка при получении TestUser: %d\nОтвет: %s", resp.StatusCode, string(body))
			// }
			// else {
			// 	log.Printf("TestUser для теста %s уже существует.", testId)
			// }
		}
	}

}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	var courses []CourseMeta

	resp, err := http.Get("http://localhost:8085/getAllShortCourseCards")
	if err != nil {
		http.Error(w, "Failed to fetch courses", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&courses)
	if err != nil {
		http.Error(w, "Failed to decode course list", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("GET", "http://localhost:8085/getAllCourseUser", nil)
	if err != nil {
		panic(err)
	}

	cookies := r.CookiesNamed("token")
	if len(cookies) < 1 || cookies[0].Value == "" {
		w.WriteHeader(401)
		return
	}
	req.AddCookie(cookies[0])

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var courseUsers []CourseUser
	if err := json.NewDecoder(resp.Body).Decode(&courseUsers); err != nil {
		panic(err)
	}

	var courseViews []CourseView

	for _, course := range courses {
		var started, completed bool

		// Ищем для этого курса все CourseUser записи
		for _, cu := range courseUsers {
			if cu.CourseId == course.Id {
				if cu.Progress.Opened {
					started = true
				}
				if cu.Progress.Finished {
					completed = true
				}
			}
		}

		courseViews = append(courseViews, CourseView{
			CourseMeta: course,
			Started:    started,
			Completed:  completed,
		})
	}

	tmpl := template.Must(template.New("").ParseFiles(
		"templates/index.html",
		"templates/header.html",
		"templates/footer.html",
		"templates/list.html",
	))

	data := PageDataShortCard{
		BaseData:   buildBaseData(r),
		HeaderType: "header", // или "headerDefault"
		Courses:    courseViews,
		ShowFooter: true,
	}

	// Выполняем шаблон и проверяем ошибку
	err = tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		// Логируем ошибку, но не вызываем http.Error, так как ответ уже мог быть частично записан
		log.Printf("Failed to execute template: %v", err)
		return
	}
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.New("").ParseFiles(
		"templates/index.html",
		"templates/header.html",
		"templates/footer.html",
		"templates/neurokids.html",
	))

	// Передаём данные в шаблон
	data := MainPageData{
		BaseData:   buildBaseData(r),
		HeaderType: "header",
		ShowFooter: true,
	}

	// Шаблон
	err := tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	var chats []Chat

	// --- Получаем токен ---
	cookies := r.CookiesNamed("token")
	if len(cookies) < 1 || cookies[0].Value == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	token := cookies[0]

	// --- Получаем текущего пользователя через /getMe ---
	userReq, _ := http.NewRequest("GET", "http://localhost:8085/getMe", nil)
	userReq.AddCookie(token)

	userResp, err := http.DefaultClient.Do(userReq)
	if err != nil || userResp.StatusCode != http.StatusOK {
		http.Error(w, "Не удалось получить текущего пользователя", http.StatusInternalServerError)
		return
	}
	defer userResp.Body.Close()

	var currentUser User
	err = json.NewDecoder(userResp.Body).Decode(&currentUser)
	if err != nil {
		http.Error(w, "Ошибка при декодировании пользователя", http.StatusInternalServerError)
		return
	}

	// --- Получаем чаты ---
	chatsReq, _ := http.NewRequest("GET", "http://localhost:8085/getChats", nil)
	chatsReq.AddCookie(token)

	chatsResp, err := http.DefaultClient.Do(chatsReq)
	if err != nil {
		http.Error(w, "Ошибка при получении чатов", http.StatusInternalServerError)
		return
	}
	defer chatsResp.Body.Close()

	err = json.NewDecoder(chatsResp.Body).Decode(&chats)
	if err != nil {
		http.Error(w, "Ошибка при декодировании чатов", http.StatusInternalServerError)
		return
	}

	// --- Собираем чаты с последними сообщениями ---
	var chatWithLastMsgList []ChatWithLastMsg

	for _, chat := range chats {
		// --- Если личный чат — подставляем имя собеседника ---
		if !chat.IsGroup {

			for _, memberID := range chat.Members {
				if memberID != currentUser.Id {

					userByIDURL := fmt.Sprintf("http://localhost:8085/getUserById?id=%d", memberID)
					userByIDReq, _ := http.NewRequest("GET", userByIDURL, nil)
					userByIDReq.AddCookie(token)

					userByIDResp, err := http.DefaultClient.Do(userByIDReq)
					if err == nil && userByIDResp.StatusCode == http.StatusOK {
						defer userByIDResp.Body.Close()
						var otherUser User
						if err := json.NewDecoder(userByIDResp.Body).Decode(&otherUser); err == nil {
							fmt.Print("menenne")
							chat.Title = otherUser.Name
						}
					}
					break
				}
			}
		}

		// --- Получаем последнее сообщение ---
		chatID := chat.Id.Hex()
		messagesURL := fmt.Sprintf("http://localhost:8085/getMessages?chatId=%s&amount=1", chatID)

		msgReq, err := http.NewRequest("GET", messagesURL, nil)
		if err != nil {
			fmt.Println("Ошибка при создании запроса к getMessages:", err)
			chatWithLastMsgList = append(chatWithLastMsgList, ChatWithLastMsg{Chat: chat})
			continue
		}
		msgReq.AddCookie(token)

		msgResp, err := http.DefaultClient.Do(msgReq)
		if err != nil {
			fmt.Println("Ошибка при запросе сообщений:", err)
			chatWithLastMsgList = append(chatWithLastMsgList, ChatWithLastMsg{Chat: chat})
			continue
		}
		defer msgResp.Body.Close()

		var raw json.RawMessage
		err = json.NewDecoder(msgResp.Body).Decode(&raw)
		if err != nil {
			fmt.Println("Ошибка при чтении JSON:", err)
			chatWithLastMsgList = append(chatWithLastMsgList, ChatWithLastMsg{Chat: chat})
			continue
		}

		var messages []Message
		if string(raw) != "null" {
			if err = json.Unmarshal(raw, &messages); err != nil {
				fmt.Println("Ошибка при разборе сообщений:", err)
			}
		}

		var lastMsg *Message
		if len(messages) > 0 {
			lastMsg = &messages[0]
		}

		chatWithLastMsgList = append(chatWithLastMsgList, ChatWithLastMsg{
			Chat:        chat,
			LastMessage: lastMsg,
		})
	}

	// --- Рендерим шаблон ---
	tmpl := template.Must(template.New("").ParseFiles(
		"templates/index.html",
		"templates/header.html",
		"templates/footer.html",
		"templates/chat.html",
	))

	data := ChatPageData{
		BaseData:   buildBaseData(r),
		HeaderType: "header",
		Chats:      chatWithLastMsgList,
		ShowFooter: false,
	}

	err = tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func diagnosticPageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.New("").ParseFiles(
		"templates/index.html",
		"templates/header.html",
		"templates/footer.html",
		"templates/diagnostic.html",
	))

	// Передаём данные в шаблон
	data := MainPageData{
		BaseData:   buildBaseData(r),
		HeaderType: "header",
		ShowFooter: true,
	}

	// Шаблон
	err := tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func settingPageHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем cookie "token"
	cookie, err := r.Cookie("token")
	if err != nil || cookie.Value == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Создаём запрос к /getMe с передачей cookie
	userReq, err := http.NewRequest("GET", "http://localhost:8085/getMe", nil)
	if err != nil {
		http.Error(w, "Ошибка создания запроса", http.StatusInternalServerError)
		return
	}
	userReq.AddCookie(cookie)

	// Выполняем запрос
	userResp, err := http.DefaultClient.Do(userReq)
	if err != nil || userResp.StatusCode != http.StatusOK {
		http.Error(w, "Не удалось получить текущего пользователя", http.StatusInternalServerError)
		return
	}
	defer userResp.Body.Close()

	// Декодируем пользователя
	var currentUser User
	err = json.NewDecoder(userResp.Body).Decode(&currentUser)
	if err != nil {
		http.Error(w, "Ошибка при декодировании пользователя", http.StatusInternalServerError)
		return
	}

	// Парсим шаблоны
	tmpl := template.Must(template.New("").ParseFiles(
		"templates/index.html",
		"templates/header.html",
		"templates/footer.html",
		"templates/setting.html",
	))

	// Формируем данные для шаблона
	data := SettingPageData{
		BaseData:   buildBaseData(r),
		HeaderType: "header",
		ShowFooter: false,
		User:       currentUser,
	}

	// Выполняем шаблон
	err = tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func userListPageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.New("").ParseFiles(
		"templates/index.html",
		"templates/header.html",
		"templates/footer.html",
		"templates/userList.html",
	))

	// Передаём данные в шаблон
	data := MainPageData{
		BaseData:   buildBaseData(r),
		HeaderType: "header",
		ShowFooter: true,
	}

	// Шаблон
	err := tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func userCreatePageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.New("").ParseFiles(
		"templates/index.html",
		"templates/header.html",
		"templates/footer.html",
		"templates/createUser.html",
	))

	// Передаём данные в шаблон
	data := MainPageData{
		BaseData:   buildBaseData(r),
		HeaderType: "header",
		ShowFooter: false,
	}

	// Шаблон
	err := tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func diagnosticInfoPageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.New("").ParseFiles(
		"templates/index.html",
		"templates/header.html",
		"templates/footer.html",
		"templates/diagnostic_info.html",
	))

	// Передаём данные в шаблон
	data := MainPageData{
		BaseData:   buildBaseData(r),
		HeaderType: "header",
		ShowFooter: true,
	}

	// Шаблон
	err := tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func contactsPageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.New("").ParseFiles(
		"templates/index.html",
		"templates/header.html",
		"templates/footer.html",
		"templates/contact.html",
	))

	// Передаём данные в шаблон
	data := MainPageData{
		BaseData:   buildBaseData(r),
		HeaderType: "header",
		ShowFooter: true,
	}

	// Шаблон
	err := tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func learnPageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.New("").ParseFiles(
		"templates/index.html",
		"templates/header.html",
		"templates/footer.html",
		"templates/learn.html",
	))

	// Передаём данные в шаблон
	data := MainPageData{
		BaseData:   buildBaseData(r),
		HeaderType: "header",
		ShowFooter: true,
	}

	// Шаблон
	err := tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
