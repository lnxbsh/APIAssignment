package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
	"strconv"
	"params"

    "github.com/gorilla/mux"
    _ "github.com/mattn/go-sqlite3"
)

func setupDB() *sql.DB {
    db, err := sql.Open("sqlite3", "./BlogDB.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

    return db
}


type Posts struct {
    id int
	Title   string  
	Content string
	Username string 
	Tags []string
	comment []string
}


type Users struct {
	id []int
	Username []string
}



type Tags struct {
	id []int
	Tagname []string
}



type JsonResponse struct {
    Type    string `json:"type"`
    Data    []Posts `json:"data"`
    Message string `json:"message"`
}


// Main function
func main() {

    // Init the mux router
    router := mux.NewRouter()

// Route handles & endpoints

    // Get all Posts
    router.HandleFunc("/posts/", GetPosts).Methods("GET")

	// Delete a specific Post by the postid 
	router.HandleFunc("/posts/{postid}", DeletePost).Methods("DELETE")

	// Update a specfic post by the postid
	router.HandleFunc("/posts/", updatePost).Methods("PUT")
    
	//Get single post by postid
	router.HandleFunc("/post/{postid}",GetPost).Methods("GET")

	//Get all comments by post id
	router.HandleFunc("/comments/{postid}",GetComment).Methods("GET")
	
	//Post a comment on post using json field and post methods
	router.HandleFunc("/comments/",PostComment).Methods("POST")
	//Get all tag from data base
	router.HandleFunc("/tags/",GetTags).Methods("GET") 
	//Get all tag of the post
	router.HandleFunc("/tags/{postid}",PostTags).Methods("GET")
	//Updating tag by tag id 
	router.HandleFunc("/tags/{tagid}",PostComment).Methods("PUT")


	
	// serve the app
    fmt.Println("Server at 8080")
    log.Fatal(http.ListenAndServe(":8000", router))
}

// Function for handling messages
func printMessage(message string) {
    fmt.Println("")
    fmt.Println(message)
    fmt.Println("")
}


func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}



func GetPost(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	postid := params["postid"]
		var id int
        var title string
        var content string
		var tag int
		var tagname string
		var comment int
		var comm string
		var user int
		var username string
		var Tgs []string
		var cmts []string

	db, err := sql.Open("sqlite3", "BlogDB.db")
    // Get all posts from posts table 

    row, err := db.Prepare("SELECT * FROM Posts where id = ?")
	checkErr()
	row.Query(strconv.Itoa(postid)).Scan(&id, &title, &content,&tag,&comment,&user)
	row,err := db.Query("SELECT content from comments where id =?",&comment).Scan(&comm)
	checkErr()
	for row.Next()
		row.Scan(&comm)
			cmts = append(cmts, comm)
	row,err := db.Query("SELECT tagname from tags where tagid =? ",&tag)
	for row.Next()
		row.Scan(&tagname)
			tgs = append(tgs, tagname)
	
	
	checkErr() 

	defer db.Close()
	post = { _title: title, _content: content, _tags: tgs, _comments:cmts}
	var response = JsonResponse{Type: "success", Data: posts}

    json.NewEncoder(w).Encode(response)

	

}

func GetComment(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	postid := params["postid"]
	var cmtid int
	var comment string
	var comments  []string
	db, err := sql.Open("sqlite3", "BlogDB.db")
	checkErr()
	rows, err := db.Query("Select cid from POSTS where id = ?", postid)
	checkErr()
	for row.Next()
		row.Scan(&cmtid)
		row,err := db.Query("Select content from comments where cmtid =?", cmtid).Scan(&comment)
		checkErr()
		comments = append(comments, comment)
		var response = JsonResponse{Type: "success", Data: comments}
		json.NewEncoder(w).Encode(response)
}

func PostComment(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	postid := params["id"]
	comment := params["comment"]
	
	row,err := db.Prepare("INSERT INTO COMMENTS VALUE(?)")
	checkErr()
	res = row.Exec(&comment)
	affected = res.RowsAffected()
	
	row,err := db.Prepare("INSERT INTO POSTS (CID) value(?) where postid = ?")
	checkErr()
	res, err := db.Exec(&affected,&postid)

	var response = JsonResponse{type:  "success"}
	json.NewEncoder(w).Encode(response)

}

func GetTags(w http.ResponseWriter, r *http.Request){
	var tgname string
	var tagnames []string
	rows,err := db.Query("Select tagname from tags")
	
	for rows.Next()
		rows.Scan(&tgname)
		tagnames = append(tagnames, tgname)

	var response = JsonResponse(type: "success", Data:tagnames)
	json.NewEncoder(w).Encode(response)

}

func PostTags(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	postid := params["id"]
	
}




func updatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

    idToUpdate := params["id"]
	title := params["title"]
	content := params["content"]
	stmt, err := db.Prepare("UPDATE post set title = ?, content = ? where id = ?")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(title,content,idToUpdate)
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)
	var response = JsonResponse{Type: "success", Data: posts}

    json.NewEncoder(w).Encode(res)

func DeletePost(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

    idToDelete := params["id"]
	

	stmt, err := db.Prepare("DELETE FROM people where id = ?")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(strconv.Itoa (idToDelete))
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)
	var response = JsonResponse{Type: "success", Data: posts}

    json.NewEncoder(w).Encode(res)
	
}



// Get all posts

// response and request handlers
func GetPosts(w http.ResponseWriter, r *http.Request) {
    //db := setupDB()

    printMessage("Getting posts...")
	db, err := sql.Open("sqlite3", "BlogDB.db")
    // Get all posts from posts table 
    rows, err := db.Query("SELECT * FROM Posts")
	defer db.Close()

	fmt.Println(rows)
    // check errors
    checkErr(err)

    // var response []JsonResponse
    var posts []Posts

    // Foreach movie
    for rows.Next() {
        var id int
        var title string
        var content string
		var tag int
		var tagname string
		var comment int
		var comm string
		var user int
		var username string
		var Tgs []string
		var cmts []string

        err = rows.Scan(&id, &title, &content,&tag,&comment,&user)
		tagsrow, err := db.Prepare("Select Tagname from Tags where Tagid = ?")
		checkErr(err)
		fmt.Println(tag)
		fmt.Println(comment)
		err = tagsrow.QueryRow(strconv.Itoa(tag)).Scan(&tagname)
        // check errors
        checkErr(err)
		Tgs = append(Tgs,tagname)


		cmtrows, err := db.Prepare("Select Content from Comments where Cmtid = ?")
		checkErr(err)
		err = cmtrows.QueryRow(&comment).Scan(&comm)
        // check errors
        checkErr(err)
		cmts = append(cmts,comm)
		usrrows, err := db.Prepare("Select Username where Userid = ?")
		checkErr(err)
		err = usrrows.QueryRow(&user).Scan(&username)
        // check errors
        checkErr(err)
		
		
        posts = append(posts, Posts{id: id, Title: title, Content:content, Username:username, Tags:Tgs, comment:cmts})
   
	}

    var response = JsonResponse{Type: "success", Data: posts}

    json.NewEncoder(w).Encode(response)
}