package controller

import (
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/whatsauth/whatsauth"
	"github.com/Ardivadiva/dipa"
	"github.com/Ardivadiva/letme/config"
)

func WsWhatsAuthQR(c *websocket.Conn) {
	whatsauth.RunSocket(c, config.PublicKey, config.Usertables[:], config.Ulbimariaconn)
}

func PostWhatsAuthRequest(c *fiber.Ctx) error {
	if string(c.Request().Host()) == config.Internalhost {
		var req whatsauth.WhatsauthRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}
		ntfbtn := whatsauth.RunModuleLegacy(req, config.PrivateKey, config.Usertables[:], config.Ulbimariaconn)
		return c.JSON(ntfbtn)
	} else {
		var ws whatsauth.WhatsauthStatus
		ws.Status = string(c.Request().Host())
		return c.JSON(ws)
	}

}

func GetHome(c *fiber.Ctx) error {
	getip := musik.GetIPaddress()
	return c.JSON(getip)
}

func GetDataUndanganRapat(c *fiber.Ctx) error {
	getun := dipa.GetDataUndanganRapat("Auditorium")
	return c.JSON(getun)
}

func GetDataListTamu(c *fiber.Ctx) error {
	getlis := dipa.GetDataListTamu("GABYAZANA")
	return c.JSON(getlis)
}

func GetDataPesertaRapat(c *fiber.Ctx) error {
	getpes := dipa.GetDataPesertaRapat("Gaby")
	return c.JSON(getpes)
}

func GetDataRapatMulai(c *fiber.Ctx) error {
	getra := dipa.GetDataRapatMulai("Jokowi")
	return c.JSON(getra)
}

func GetDataWaktuRapat(c *fiber.Ctx) error {
	getwa := dipa.GetDataWaktuRapat("rapat")
	return c.JSON(getwa)
}