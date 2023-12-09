const express = require("express")
const Controller = require("../Controllers/controller.js")

const router = express.Router()

router.post("/register",Controller.register)
router.put("/buycid",Controller.buyCID)
router.put("/addbal",Controller.addBal)
router.put("/delbal",Controller.delBal)

module.exports = router