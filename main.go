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

// Inisiasi Database (Simpan di dalam slice)
var students []Student

func main() {
	// Inisiasi Echo Framework
	e := echo.New()

	// Routing
	e.GET("/students", getStudents)          // Mendapatkan semua data siswa
	e.GET("/students/:id", getStudentByID)   // Mendapatkan data siswa berdasarkan ID
	e.POST("/students", createStudent)       // Menambahkan data siswa
	e.PUT("/students/:id", updateStudent)    // Mengubah data siswa berdasarkan ID
	e.DELETE("/students/:id", deleteStudent) // Menghapus data siswa berdasarkan ID

	// Menjalankan Echo Framework di port 1323
	e.Logger.Fatal(e.Start(":1323"))
}

// Mendapatkan semua data siswa dan mengembalikan HTTP Status
func getStudents(c echo.Context) error {
	return c.JSON(http.StatusOK, students)
}

// Mendapatkan data siswa berdasarkan ID dan mengembalikan HTTP Status
func getStudentByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, student := range students {
		if student.ID == id {
			return c.JSON(http.StatusOK, student)
		}
	}
	return c.JSON(http.StatusNotFound, "Student Not Found")
}

// Menambahkan data siswa dan mengembalikan HTTP Status
func createStudent(c echo.Context) error {
	student := new(Student)
	if err := c.Bind(student); err != nil {
		return err
	}

	student.ID = len(students) + 1
	students = append(students, *student)

	return c.JSON(http.StatusCreated, student)
}

// Mengubah data siswa berdasarkan ID dan mengembalikan HTTP Status
func updateStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, student := range students {
		if student.ID == id {
			updatedStudent := new(Student)
			if err := c.Bind(updatedStudent); err != nil {
				return err
			}

			updatedStudent.ID = student.ID
			students[i] = *updatedStudent

			return c.JSON(http.StatusOK, updatedStudent)
		}
	}
	return c.JSON(http.StatusNotFound, "Student Not Found")
}

// Menghapus data siswa berdasarkan ID dan mengembalikan HTTP Status
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
