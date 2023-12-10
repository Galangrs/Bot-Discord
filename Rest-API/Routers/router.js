const express = require("express")
const Controller = require("../Controllers/controller.js")
const { Authentication , Authorization } = require("../Middleware/auth.js")
const router = express.Router()

router.post("/register",Controller.register)

router.use(Authentication)

router.put("/buycid",Controller.buyCID)
router.get("/balance",Controller.getBal)
router.put("/addbal",Authorization,Controller.addBal)
router.put("/delbal",Authorization,Controller.delBal)

module.exports = router