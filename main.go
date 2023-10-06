package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Inisiasi Constructor Data
type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade string `json:"grade"`
}

// Inisiasi Database
var students []Student

func main() {
	// Inisiasi Echo Framework
	e := echo.New()

	students = []Student{}
	// Routing
	e.GET("/students", getStudents)          // Mengambil semua data
	e.GET("/students/:id", getStudentsByID)  // Mengambil data berdasarkan ID
	e.POST("/students", createStudent)       // Menambahkan data
	e.PUT("/students/:id", updateStudent)    // Mengubah data
	e.DELETE("/students/:id", deleteStudent) // Menghapus data

	e.Logger.Fatal(e.Start(":1323")) // Menjalankan Echo Framework
}

// Mengambil semua data yang ada dan mengembalikan HTTP Status
func getStudents(c echo.Context) error {
	return c.JSON(http.StatusOK, students)
}

// Mengambil data berdasarkan ID dan mengembalikan HTTP Status
func getStudentsByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, student := range students {
		if student.ID == id {
			return c.JSON(http.StatusOK, student)
		}
	}
	return c.JSON(http.StatusNotFound, "Student Not Found")
}

// Menambahkan data dan mengembalikan HTTP Status
func createStudent(c echo.Context) error {
	student := new(Student)
	if err := c.Bind(student); err != nil {
		return err
	}

	student.ID = len(students) + 1
	students = append(students, *student)

	return c.JSON(http.StatusCreated, student)
}

// Mengubah data berdasarkan ID dan mengembalikan HTTP Status
func updateStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, student := range students {
		if student.ID == id {
			// Konversi nilai string ke int
			age, err := strconv.Atoi(c.FormValue("age"))
			if err != nil {
				return c.JSON(http.StatusBadRequest, "Invalid Age")
			}

			students[i].Name = c.FormValue("name")
			students[i].Age = age // Menggunakan nilai yang telah dikonversi
			students[i].Grade = c.FormValue("grade")
			return c.JSON(http.StatusOK, students[i])
		}
	}
	return c.JSON(http.StatusNotFound, "Student Not Found")
}

// Menghapus data berdasarkan ID dan mengembalikan HTTP Status
func deleteStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			return c.JSON(http.StatusOK, "Student Deleted")
		}
	}
	return c.JSON(http.StatusNotFound, "Student Not Found")
}
