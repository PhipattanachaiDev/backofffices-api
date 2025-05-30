package certificatesController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

func CreatePDF(c *gin.Context) {
	// สร้าง PDF ใหม่
	pdf := gofpdf.New("P", "mm", "A4", "")

	// เพิ่มหน้าของ PDF
	pdf.AddPage()

	// ตั้งค่า font
	pdf.SetFont("Arial", "B", 16)

	// เขียนข้อความใน PDF
	pdf.Cell(40, 10, "Hello, World!")

	// กำหนด headers สำหรับ response PDF
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=example.pdf")

	// ส่ง PDF ไปยัง client
	if err := pdf.Output(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

}

// ฟังก์ชันสำหรับสร้าง PDF และแสดงเป็น preview
// ฟังก์ชันสำหรับสร้าง PDF และแสดงเป็น preview
func PreviewPDF(c *gin.Context) {
	// สร้าง PDF ใหม่
	pdf := gofpdf.New("P", "mm", "A4", "")

	// โหลดฟอนต์ภาษาไทย (ตรวจสอบว่าไฟล์ฟอนต์ .ttf อยู่ในโฟลเดอร์ที่ถูกต้อง)
	// ถ้าใช้ฟอนต์ภาษาไทยที่รองรับเช่น NotoSansThai
	pdf.AddUTF8Font("NotoSansThai", "", "./static/fonts/NotoSansThai-Regular.ttf")
	// โหลดฟอนต์ภาษาอังกฤษ
	pdf.AddUTF8Font("NotoSans", "", "./static/fonts/NotoSans-Regular.ttf")

	// เพิ่มหน้าใหม่
	pdf.AddPage()

	// ตั้งค่า font เป็นฟอนต์ภาษาไทย
	pdf.SetFont("NotoSansThai", "", 12)

	// แยกส่วนของข้อความแต่ละส่วนที่ต้องการแสดง
	// ส่วนที่ 1: ข้อความหัวข้อ
	pdf.CellFormat(0, 10, "หนังสือรับรองการติดตั้งเครื่องบันทึกข้อมูลการเดินทางของรถ", "", 1, "C", false, 0, "")
	pdf.Ln(5) // เพิ่มระยะห่าง (line break)

	// ส่วนที่ 2: ข้อความที่ 1
	text1 := `เลขที่หนังสือ 0001/2564`
	pdf.MultiCell(0, 10, text1, "", "", false)
	pdf.Ln(5) // เพิ่มระยะห่างระหว่างข้อความ

	// ส่วนที่ 3: ข้อความที่ 2 (ภาษาอังกฤษ)
	text2 := `Teletec Co., Ltd.`
	pdf.SetFont("NotoSans", "", 12) // เปลี่ยนฟอนต์เป็นภาษาอังกฤษ
	pdf.MultiCell(0, 10, text2, "", "", false)
	pdf.Ln(5) // เพิ่มระยะห่างระหว่างข้อความ

	// ส่วนที่ 4: ข้อความที่ 3 (ภาษาไทย)
	text3 := `ขอให้ท่านตรวจสอบและยืนยันการติดตั้งเครื่องบันทึกตามรายละเอียดในเอกสารนี้`
	pdf.SetFont("NotoSansThai", "", 12) // กลับมาฟอนต์ภาษาไทย
	pdf.MultiCell(0, 10, text3, "", "", false)

	// ตั้งค่า header สำหรับแสดงผล PDF โดยตรงใน browser (Inline)
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "inline; filename=cetificate.pdf")

	// ส่ง PDF ไปยัง client
	if err := pdf.Output(c.Writer); err != nil {
		// ถ้ามีข้อผิดพลาดในการสร้างหรือส่ง PDF
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถสร้าง PDF ได้: " + err.Error()})
		return
	}
}
