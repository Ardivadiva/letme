package url

import (
	"github.com/Ardivadiva/letme/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsaut
	page.Get("/", controller.GetHome)
	page.Post("/insert", controller.InsertListTamu)
	page.Get("/tamu", controller.GetDataListTamu)
	page.Get("/undangan", controller.GetDataUndanganRapat)
	page.Get("/peserta", controller.GetDataPesertaRapat)
	page.Get("/waktu", controller.GetDataWaktuRapat)
	page.Get("/rapat", controller.GetDataRapatMulai)

}